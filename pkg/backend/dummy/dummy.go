package dummy

import (
	"log"
	"io/ioutil"
	"fmt"
	"os"
	"time"
)


// InitFile will create a basic file to when none exists.
func InitFile(buildNum int, prefix, dateTime, gitHash string)  {
	combinedMeta := fmt.Sprintf("%d.%s.%s.%s", buildNum, prefix, dateTime, gitHash)
	byteMeta := []byte(combinedMeta)

	err := ioutil.WriteFile("version.data", byteMeta, 0777)
	if err != nil {
		log.Println(err)
	}
}
func fileExists(path string) bool {
	// path to file does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// LocalBackend will be running all operations locally.
func LocalBackend(prefix, path string) {
	dateTime := time.Now().Format("2006-01-02_15:04:05")
	gitHash := git.GetCurrentHashShort()

	if !fileExists(path) {
		initBuildNum := 1
		local.InitFile(initBuildNum, prefix, dateTime, gitHash)
	} else {
		
	}
}

// ParseStorageFile will read in the storage file
func ParseStorageFile() {

}

func saveArtefact() {
}