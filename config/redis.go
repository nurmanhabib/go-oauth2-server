package config

// Redis is a container for configuration related redis connection.
type Redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// WithRedis is a function to apply database configuration options.
func WithRedis() Option {
	return func(config *Config) {
		config.Redis = Redis{
			Host:     GetEnv("REDIS_HOST", "127.0.0.1"),
			Port:     GetEnv("REDIS_PORT", "6379"),
			Password: GetEnv("REDIS_PASSWORD", ""),
			DB:       GetEnvAsInt("REDIS_DB", 0),
		}
	}
}
