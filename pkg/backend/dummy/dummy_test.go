package dummy

import "testing"

func TestLocalBackend(t *testing.T) {
	prefix := "1.3"
	path := "./version.data"
	LocalBackend(prefix, path)
}