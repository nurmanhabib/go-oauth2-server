package config

import "time"

func loadTimezone() Option {
	return func(config *Config) {
		timeLoc, _ := time.LoadLocation(config.Application.Timezone)
		time.Local = timeLoc
	}
}
