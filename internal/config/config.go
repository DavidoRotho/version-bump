package config

import (
	"context"
	"fmt"

	"github.com/MichaelDao/version-bump/internal/version"
	"github.com/urfave/cli"
)

// SetConfiguration will receive the entrypoint flags that are defined when the tool runs.
func SetConfiguration(urfaveCli *cli.App) {
	var repository string
	var prefix string
	var googleProjectID string
	var backend string

	urfaveCli.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "repo, r",
			Usage:       "The directory of the repository to bump.",
			Destination: &repository,
		},
		cli.StringFlag{
			Name:        "prefix, p",
			Usage:       "The location of the file with the prefix for our version.",
			Destination: &prefix,
		},
		cli.StringFlag{
			Name:        "gcp-project, g",
			Usage:       "The project ID of the google project you are working with.",
			Destination: &googleProjectID,
		},
		cli.StringFlag{
			Name:        "backend, b",
			Usage:       "The backend driver for this tool.",
			Destination: &backend,
		},
	}

	urfaveCli.Action = func(cliContext *cli.Context) error {
		// Repository flag defaults to local directory if unspecified
		// TODO figure out whether we need this flag for git tagging, maybe disable git tagging flag?
		if len(repository) == 0 {
			fmt.Println("repository defaulting to local directory.")
			repository = "."
		}

		// Prefix flag must be set
		// TODO validate prefix file
		if len(prefix) == 0 {
			return fmt.Errorf("The flag '--prefix' requires the location of a text file with a string")
		}

		// Backend flag must be set
		if len(backend) == 0 {
			return fmt.Errorf("\nPlease specify a backend driver with '--backend':\n\t- gcs (google cloud storage)\n\t- dummy (test)")
		}

		// Run the app
		version := version.Version{
			Repository:      repository,
			Prefix:          prefix,
			GoogleProjectID: googleProjectID,
			Backend:         backend,
		}
		return version.Run(context.Background())
	}
}
