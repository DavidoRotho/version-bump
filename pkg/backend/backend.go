package backend

/*
abstract class Backend() {
	void lock();
	void unlock();
	List<Build> getHistory();
	Build addBuild(BuildMeta meta);
}
*/

// Backend will define how we should do our backend implementations
type Backend interface {
	Lock() string
	Unlock() string
	// getHistory() []Build
	// addBuild(meta BuildMeta) Build
}