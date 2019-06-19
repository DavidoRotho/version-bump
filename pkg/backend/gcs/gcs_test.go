package gcs 

import (
	"testing"
	"fmt"
)

// func TestLocalBackend(t *testing.T) {
// 	prefix := "1.3"
// 	path := "./version.data"
// 	LocalBackend(prefix, path)
// }
	// // Run the app
		// flags := Flags {
		// 	Prefix:         prefixString,
		// 	Backend:        backend,
		// 	Localstorage: 	localStorage,
		// }

func TestLock(t *testing.T) {
	gcsStruct := Gcs {} 

	fmt.Print(gcsStruct.lock())
}
