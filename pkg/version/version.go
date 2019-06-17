package version

import (
	"context"
	"github.com/MichaelDao/version-bump/pkg/backend"
)

// Version stores the data relevant to initialize and run the App.
type Version struct {
	Prefix          string
	GoogleProjectID string
	Backend         string
	Localstorage 	string
}

// Run initializes the configuration.
func (version *Version) Run(ctx context.Context) error {

	if version.Backend == "local" {
		backend.LocalBackend(Prefix, Localstorage)		
	}

	return nil
}
