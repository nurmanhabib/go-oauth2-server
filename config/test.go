package config

// TestMode is a function to enable test mode.
func TestMode() Option {
	return func(config *Config) {
		config.TestMode = true
	}
}
