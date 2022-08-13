package config

// DefaultAppXXX is the default constant of the application configuration.
const (
	DefaultAppEnvironment = "local"
	DefaultAppLanguage    = "en"
	DefaultAppTimezone    = "UTC"
)

// Application is a container for configuration related applications.
type Application struct {
	Environment string
	Language    string
	Timezone    string
}

func withApplication() Option {
	return func(config *Config) {
		config.Application = Application{
			Environment: GetEnv("APP_ENV", DefaultAppEnvironment),
			Language:    GetEnv("APP_LANG", DefaultAppLanguage),
			Timezone:    GetEnv("APP_TIMEZONE", DefaultAppTimezone),
		}
	}
}
