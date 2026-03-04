package config

type Config struct {
	ServerPort string
	ApiPrefix  string
}

func NewConfig() *Config {
	return &Config{
		ServerPort: "127.0.0.1:8080",
		ApiPrefix:  "/api/v1",
	}
}
