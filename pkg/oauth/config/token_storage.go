package config

import (
	"github.com/go-oauth2/oauth2/v4/store"
)

func defaultMemoryTokenStorage() Option {
	return func(config *Config) {
		var err error

		config.TokenStorage, err = store.NewMemoryTokenStore()
		if err != nil {
			panic(err)
		}
	}
}
