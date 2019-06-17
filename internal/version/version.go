package version

import (
	"context"
)

// Version stores the data relevant to initialize and run the App.
type Version struct {
	Repository      string
	Prefix          string
	GoogleProjectID string
	Backend         string
}

// Run initializes the configuration.
func (version *Version) Run(ctx context.Context) error {
	
	return nil
}
