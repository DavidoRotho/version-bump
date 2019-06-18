package config

import (
	"fmt"
	"github.com/urfave/cli"
)

// Flags stores the data relevant to initialize and run the App.
type Flags struct {
	Prefix          string
	Backend         string
	Localstorage 	string
}

// ParseCLi will receive the entrypoint flags that are defined when the tool runs.
func ParseCLi(urfaveCli *cli.App) {
	var prefix string
	var manualPrefix string
	var backend string
	var localStorage string
	// var bucketStorage string
	// var googleProjectID string

	urfaveCli.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "prefix, p",
			Usage:       "Specify the location of the prefix file for our version scheme.",
			Destination: &prefix,
		},
		cli.StringFlag{
			Name:        "manual-prefix, m",
			Usage:       "Manually set the prefix.",
			Destination: &manualPrefix,
		},
		cli.StringFlag{
			Name:        "backend, b",
			Usage:       "Specify the backend driver applied to this tool.",
			Destination: &backend,
		},
		cli.StringFlag{
			Name:        "dummy-location, d",
			Usage:       "Local storage can be specified for with the 'dummy' driver.",
			Destination: &localStorage,
		},
		// cli.StringFlag{
		// 	Name:        "gcp-project, g",
		// 	Usage:       "The project ID of the google project you are working with.",
		// 	Destination: &googleProjectID,
		// },
		// cli.StringFlag{
		// 	Name:        "bucket-storage, b",
		// 	Usage:       "Specify the bucket location as a part of the 'gcs' backend.",
		// 	Destination: &bucketStorage,
		// },
	}

	urfaveCli.Action = func(cliContext *cli.Context) error {
		// Parse the prefix flag
		prefixString := prefix
		if len(prefix) == 0 {
			if len(manualPrefix) != 0 {
				prefixString = manualPrefix
			} else {
				return fmt.Errorf("The flag '--prefix' needs the location of the text file, you can manually override it with '--manual-prefix'") 
			}
		}

		// Parse the backend flag
		switch backend {
		case "":
			return fmt.Errorf("\nPlease specify a backend driver with '--backend':\n\t- gcs (google cloud storage)\n\t- dummy (local storage)")
		case "dummy":
			fmt.Println("dummy")
			// config.backend = new DummyBackend();
		case "gcs":
			fmt.Println("google cloud")
		}

		// TODO Continue fleshing out flag parsing
		fmt.Print(prefixString)
		// // Run the app
		// flags := Flags {
		// 	Prefix:         prefixString,
		// 	Backend:        backend,
		// 	Localstorage: 	localStorage,
		// }
		
		return nil
	}
}
