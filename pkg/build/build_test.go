package build

import (
	"fmt"
	"testing"

	// "github.com/MichaelDao/version-bump/pkg/buildmeta"
)

func TestGetBuild(t *testing.T) {
	build := getBuild()
	fmt.Print(build)
}

func getBuild() Build {
	// Run the app
	var testMeta = buildmeta.BuildMeta{
		CommitHash:   "test",
		HumanVersion: "test",
		Timestamp:    "test",
	}

	var newBuild = Build{
		meta:        testMeta,
		buildNumber: 1,
	}
	return newBuild
}
