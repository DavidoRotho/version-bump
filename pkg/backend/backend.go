package backend

/*
abstract class Backend() {
	void lock();
	void unlock();
	List<Build> getHistory();
	Build addBuild(BuildMeta meta);
}
*/

import (
	"github.com/MichaelDao/version-bump/pkg/build"
	"github.com/MichaelDao/version-bump/pkg/buildmeta"
)

// Backend will define how we should do our backend implementations
type Backend interface {
	lock()
	unlock()
	getHistory() []Build
	addBuild(meta BuildMeta) Build
}

func lock() {}

func unlock() {}

func getHistory() []Build {}

func addBuild(meta BuildMeta) Build {}
