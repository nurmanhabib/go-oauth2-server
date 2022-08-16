package config

// DefaultAppXXX is the default constant of the application configuration.
const (
	DefaultAppEnvironment = "local"
	DefaultAppLanguage    = "en"
	DefaultAppTimezone    = "UTC"
	DefaultAppPort        = 8000
)

// Application is a container for configuration related applications.
type Application struct {
	Environment string
	Language    string
	Timezone    string
	Port        int
}

func withApplication() Option {
	return func(config *Config) {
		config.Application = Application{
			Environment: GetEnv("APP_ENV", DefaultAppEnvironment),
			Language:    GetEnv("APP_LANG", DefaultAppLanguage),
			Timezone:    GetEnv("APP_TIMEZONE", DefaultAppTimezone),
			Port:        GetEnvAsInt("APP_PORT", DefaultAppPort),
		}
	}
}
