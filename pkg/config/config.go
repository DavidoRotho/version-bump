package config

import (
	"context"
	"fmt"

	"github.com/MichaelDao/version-bump/pkg/version"
	"github.com/urfave/cli"
)

// SetConfiguration will receive the entrypoint flags that are defined when the tool runs.
func SetConfiguration(urfaveCli *cli.App) {
	var prefix string
	var manualPrefix string
	var googleProjectID string
	var backend string
	var localStorage string
	var bucketStorage string

	urfaveCli.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "prefix, p",
			Usage:       "The location of the file with the prefix for our version.",
			Destination: &prefix,
		},
		cli.StringFlag{
			Name:        "manual-prefix, m",
			Usage:       "Manually set the prefix through the tool.",
			Destination: &manualPrefix,
		},
		cli.StringFlag{
			Name:        "backend, b",
			Usage:       "The backend driver for this tool.",
			Destination: &backend,
		},
		cli.StringFlag{
			Name:        "local-storage, l",
			Usage:       "Local storage that must be set as a part of the 'local' backend.",
			Destination: &localStorage,
		},
		cli.StringFlag{
			Name:        "gcp-project, g",
			Usage:       "The project ID of the google project you are working with.",
			Destination: &googleProjectID,
		},
		cli.StringFlag{
			Name:        "bucket-storage, b",
			Usage:       "Specify the bucket location as a part of the 'gcs' backend.",
			Destination: &bucketStorage,
		},
	}

	urfaveCli.Action = func(cliContext *cli.Context) error {		
		// Prefix or readPrfix flag must be set
		// TODO validate prefix file when reading it in
		prefixString := prefix
		if len(prefix) == 0 {
			// override the string with the manual prefix
			if len(manualPrefix) != 0 {
				prefixString = manualPrefix
			} else {
				return fmt.Errorf("The flag '--prefix' requires the location of a text file with a string") 
			}
		}

		// Backend flag must be set
		if len(backend) == 0 {
			return fmt.Errorf("\nPlease specify a backend driver with '--backend':\n\t- gcs (google cloud storage)\n\t- local")
		}

		// Run the app
		version := version.Version{
			Prefix:          prefixString,
			GoogleProjectID: googleProjectID,
			Backend:         backend,
			LocalStorage: localStorage
		}
		return version.Run(context.Background())
	}
}
