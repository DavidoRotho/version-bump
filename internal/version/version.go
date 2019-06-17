package version

import (
	"context"
	"fmt"
)

// Version stores the data relevant to initialize and run the App.
type Version struct {
	Repository string
	Prefix     string
}

// Run initializes the configuration.
func (version *Version) Run(ctx context.Context) error {
	if len(version.Repository) == 0 {
		fmt.Println("repository not set, defaulting to local directory. './'")
		version.Repository = "."
	} else {
		fmt.Print("set repository to '" + version.Repository + "'")
	}
	return nil
}
