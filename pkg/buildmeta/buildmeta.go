package buildmeta

// // this is a struct
// /*

// struct BuildMeta {
// 	string commitHash;
// 	string humanVersion;
// 	DateTime timeStamp;
// 	... as discussed
// }

// meta = BuildMeta("blah", "blah", today);
// print meta.commitHash

// may be you don't need get methods
// */

// BuildMeta will store all meta information about our version
type BuildMeta struct {
	CommitHash string
	HumanVersion string
	Timestamp string
}
