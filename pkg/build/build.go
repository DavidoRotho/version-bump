package build

import 	"github.com/MichaelDao/version-bump/pkg/buildmeta"

// Build contains all the data about the current version
type Build struct {
	meta buildmeta.BuildMeta
	buildNumber int
}
