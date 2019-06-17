package backend

import (
	"os"
	"time"

	"github.com/MichaelDao/version-bump/pkg/git"
	"github.com/MichaelDao/version-bump/pkg/backend/local"
)

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {	
		// path/to/whatever does not exist
		return false
	  }
	return true
}

// LocalBackend will be running all operations locally.
func LocalBackend(prefix, path string) {
	dt := time.Now().Format("2006-01-02_15:04:05")
	gitHash := git.GetCurrentHashShort()

	if !fileExists(path) {
		initBuildNum := 1
		local.InitFile(initBuildNum, prefix, dt, gitHash)
	} else {
		 
	}
}


// GCSBackend will run all functions related to the gcs backend.
func GCSBackend(){

}

func lock() {
	// TODO setup lock mechanism
}

func unlock() {
	// TODO setup lock mechanism
}

func getHistory() {
	// TODO get history of all the logs we have stored in our build file
}

func addBuild()  {
	// TODO create a build from the metabuild we provide
}
