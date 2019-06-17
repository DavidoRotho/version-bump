package config

import (
	"context"

	"github.com/MichaelDao/version-bump/internal/version"
	"github.com/urfave/cli"
)

// SetConfiguration will receive the entrypoint flags that are defined when the tool runs.
func SetConfiguration(urfaveCli *cli.App) {
	var repository string
	var prefix string

	urfaveCli.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "repo, r",
			Usage:       "The directory of the repository to bump.",
			Destination: &repository,
		},
		cli.StringFlag{
			Name: "prefix, p",
			Usage: 	"The location of the file with the prefix for our version.",
			Destination: &prefix,
		},
	}

	urfaveCli.Action = func(cliContext *cli.Context) error {
		version := version.Version{
			Repository: repository,
			Prefix: prefix,
		}
		return version.Run(context.Background())
	}
}