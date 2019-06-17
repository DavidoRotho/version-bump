package local

import (
	"log"
	"io/ioutil"
	"fmt"
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

// ParseStorageFile will read in the storage file
func ParseStorageFile() {

}

func saveArtefact() {
}