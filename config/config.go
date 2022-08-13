package config

// Config is a container for all configurations.
type Config struct {
	Application

	Database Database
	Redis    Redis

	TestMode bool
}

// New is a new configuration constructor
// with Option parameters to manipulate as it builds.
func New(options ...Option) *Config {
	config := &Config{}

	// Default Config
	withApplication()(config)

	for _, apply := range options {
		apply(config)
	}

	// Load Timezone
	loadTimezone()(config)

	return config
}

// NewTestMode is a new configuration constructor with test mode enabled
// and Option parameters to manipulate at build.
func NewTestMode(options ...Option) *Config {
	options = append([]Option{
		TestMode(),
	}, options...)

	return New(options...)
}
