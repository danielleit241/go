package middleware

import (
	"strconv"
	"sync"
	"time"

	"github.com/danielleit241/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

const penaltyDuration = time.Minute

type IPRateLimiter struct {
	mu    sync.Mutex
	ips   map[string]*client
	r     rate.Limit
	burst int
}

type client struct {
	limiter   *rate.Limiter
	blockedAt time.Time
}

func NewIPRateLimiter(r rate.Limit, b int, rateLimitLogger *zerolog.Logger) *IPRateLimiter {
	i := &IPRateLimiter{
		ips:   make(map[string]*client),
		r:     r,
		burst: b,
	}
	go i.cleanup(rateLimitLogger)
	return i
}

func (i *IPRateLimiter) cleanup(rateLimitLogger *zerolog.Logger) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		i.mu.Lock()
		for ip, c := range i.ips {
			if time.Since(c.blockedAt) > 3*penaltyDuration && c.blockedAt.IsZero() == false {
				rateLimitLogger.Warn().Str("ip", ip).Msg("[Cleanup] Removing inactive IP")
				delete(i.ips, ip)
			}
		}
		i.mu.Unlock()
	}
}

// ab -n 20 -c 1 http://localhost:8080/api/v1/users
func RateLimit(rateLimitLogger *zerolog.Logger) gin.HandlerFunc {
	burstLimit := getEnvLimit("BURST_LIMIT", 10)
	rateLimit := getEnvLimit("RATE_LIMIT", 5)

	limiter := NewIPRateLimiter(rate.Limit(rateLimit), burstLimit, rateLimitLogger)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		limiter.mu.Lock()
		s, exists := limiter.ips[ip]
		if !exists {
			s = &client{limiter: rate.NewLimiter(limiter.r, limiter.burst)}
			limiter.ips[ip] = s
		}

		if !s.blockedAt.IsZero() && now.Sub(s.blockedAt) < penaltyDuration {
			retry := int(penaltyDuration.Seconds() - now.Sub(s.blockedAt).Seconds())
			limiter.mu.Unlock()

			if shouldLogRateLimit(ip) {
				logRateLimitWarn(rateLimitLogger, c, ip).
					Int("retry_after_seconds", max(retry, 1)).
					Msg("Rate limit exceeded, currently blocked")
			}

			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests", "retry_after": max(retry, 1)})
			return
		}

		if !s.limiter.Allow() {
			s.blockedAt = now
			limiter.mu.Unlock()

			if shouldLogRateLimit(ip) {
				logRateLimitWarn(rateLimitLogger, c, ip).
					Msg("Rate limit exceeded, blocked for 1m")
			}

			c.AbortWithStatusJSON(429, gin.H{"error": "Rate limit exceeded, blocked for 1m"})
			return
		}
		limiter.mu.Unlock()
		c.Next()
	}
}

func logRateLimitWarn(rateLimitLogger *zerolog.Logger, c *gin.Context, ip string) *zerolog.Event {
	return rateLimitLogger.Warn().
		Str("client_ip", ip).
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Str("user_agent", c.GetHeader("User-Agent")).
		Str("referer", c.GetHeader("Referer"))
}

func getEnvLimit(key string, defaultVal int) int {
	valStr := utils.GetEnvOrDefault(key, strconv.Itoa(defaultVal))
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}

var rateLimitCache = sync.Map{}

const rateLimitLogTTL = 10 * time.Second

func shouldLogRateLimit(ip string) bool {
	now := time.Now()
	if val, exists := rateLimitCache.Load(ip); exists {
		if t, ok := val.(time.Time); ok && now.Sub(t) < rateLimitLogTTL {
			return false
		}
	}

	rateLimitCache.Store(ip, now)
	return true
}
