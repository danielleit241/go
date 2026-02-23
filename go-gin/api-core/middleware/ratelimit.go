package middleware

import (
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
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

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips:   make(map[string]*client),
		r:     r,
		burst: b,
	}
	go i.cleanup()
	return i
}

func (i *IPRateLimiter) cleanup() {
	for range time.Tick(1 * time.Minute) {
		i.mu.Lock()
		for ip, c := range i.ips {
			if time.Since(c.blockedAt) > 3*penaltyDuration && c.blockedAt.IsZero() == false {
				log.Printf("[Cleanup] Removing inactive IP: %s", ip)
				delete(i.ips, ip)
			}
		}
		i.mu.Unlock()
	}
}

// ab -n 20 -c 1 http://localhost:8080/api/v1/users
func RateLimit() gin.HandlerFunc {
	limiter := NewIPRateLimiter(5, 10)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		limiter.mu.Lock()
		s, exists := limiter.ips[ip]
		if !exists {
			s = &client{limiter: rate.NewLimiter(limiter.r, limiter.burst)}
			limiter.ips[ip] = s
		}

		// 1. Check nếu đang trong thời gian bị phạt
		if !s.blockedAt.IsZero() && now.Sub(s.blockedAt) < penaltyDuration {
			retry := int(penaltyDuration.Seconds() - now.Sub(s.blockedAt).Seconds())
			limiter.mu.Unlock()
			c.AbortWithStatusJSON(429, gin.H{"error": "Too many requests", "retry_after": max(retry, 1)})
			return
		}

		// 2. Check token bucket
		if !s.limiter.Allow() {
			s.blockedAt = now // Đánh dấu bị chặn
			limiter.mu.Unlock()
			c.AbortWithStatusJSON(429, gin.H{"error": "Rate limit exceeded, blocked for 1m"})
			return
		}
		limiter.mu.Unlock()
		c.Next()
	}
}
