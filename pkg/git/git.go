package git

import (
	"os/exec"
	"log"
)


// GetCurrentHash will return the hash of the checkedout git commit
func GetCurrentHash() string {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("Failed with %s\n", err)
    }
	return string(out)
}

// GetCurrentHashShort will return the short version of the hash
func GetCurrentHashShort() string {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("Failed with %s\n", err)
    }
	return string(out)
}

func tagRepository() {
	// TODO tag the current repository
}

func fetchTags(){
	// TODO fetch git tags
}

func pushTags(){
	// TODO push the git tags
}