package gcs


/*
List<Build> getHistory();
void lock();
void unlock();
Build addNewBuild(BuildMeta buildMeta);
*/

type Backend interface {
	Lock() string
	Unlock() string
}

// Gcs is the struct for Google cloud storage
type Gcs struct {}

// Lock google cloud storage.
func (p Gcs) Lock() string {
	return "Lock GCS"
}

// Unlock google cloud storage.
func (p Gcs) Unlock() string {
	return "Unlock GCS"
}
