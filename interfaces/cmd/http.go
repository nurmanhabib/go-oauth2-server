package cmd

import (
	"github.com/urfave/cli/v2"
)

func WithHTTPService() cli.Commands {
	return cli.Commands{
		{
			Name: "http",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
}
