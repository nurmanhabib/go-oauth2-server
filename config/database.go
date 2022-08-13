package config

// Database is a container for configuration related database connection.
type Database struct {
	DBDriver                    string
	DBHost                      string
	DBPort                      string
	DBUser                      string
	DBName                      string
	DBPassword                  string
	DBTimeZone                  string
	DBLog                       bool
	DisableForeignKeyConstraint bool
}

// WithDatabase is a function to apply database configuration options.
func WithDatabase() Option {
	return func(config *Config) {
		config.Database = Database{
			DBDriver:                    GetEnv("DB_DRIVER", "mysql"),
			DBHost:                      GetEnv("DB_HOST", "localhost"),
			DBPort:                      GetEnv("DB_PORT", "3306"),
			DBUser:                      GetEnv("DB_USER", "root"),
			DBName:                      GetEnv("DB_NAME", "skeleton"),
			DBPassword:                  GetEnv("DB_PASSWORD", ""),
			DBTimeZone:                  GetEnv("APP_TIMEZONE", "Asia/Jakarta"),
			DBLog:                       GetEnvAsBool("ENABLE_LOGGER", true),
			DisableForeignKeyConstraint: GetEnvAsBool("DISABLE_FOREIGN_KEY_CONSTRAINT", false),
		}
	}
}
