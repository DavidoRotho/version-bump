# version-bump

```java

class BuildMeta {
	int x, y;
}

class Build {
	BuildMeta meta;
	int buildNumber;
}

abstract class Backend() {
	void lock();
	void unlock();
	List<Build> getHistory();
	Build addBuild(BuildMeta meta);
}

class GCSBackend() {
	GCSBackend(Map<String, String> flags)  {
		if (!flags.contains("bucket")) {
			throw error;
		}
	}

	void lock () {
		// do GCS lock
	}

	void unlock() {
		// do GCS unlock
	}

	List<Build> getHistory() {
		if (!locked) {
			throw
		}

		read history file
		for each line
			convert line into Build

		return a list of builds
	}

	...
}

class DummyBackend() {
	void lock () {
		lock = true
	}

	void unlock() {
		lock = false;
	}

	List<Build> getHistory() {
		if (!locked) {
			throw
		}

		read history file locally
		for each line
			convert line into Build

		return a list of builds
	}

	...
}

class Configuration {
	Backend backend;
	VersionCalculator calculator;
	// something else 
}

Configuration parseCLI(args) {
	Configuration config;

	// VERSION 1
	if (args.contains("gcs")) {
		config.backend = new GCSBackend();
	} else {
		config.backend = new DummyBackend();
	}

	// VERSION 2
	Map<String, String> flags = args.parseFlagsAsMap();
	if (!flags.contains("backend")) {
		throw error
	}
	if (flags["backend"] == "gcs") {
		config.backend = new GCSBackned(flags);
	}

	// VERSION 3
	config.backend = BackendFactory.createBackend(flags)

	return config;
}

int main() {
	Configuration config = parseCLI(args);

	config.backend.lock();
	
	val history = config.backend.getHistory();
	val nextHumanVersion = config.caclulator.calculate(history);
	val meta = BuildMeta(config.git.commitHash, nextHumanVersion)
	val build = config.backend.addBuild(meta);

	config.backend.unlock();

	// ALMOST LAST PR
	SuperGit git = SuperGit()
	git.pushTags(build.buildNumer, build.meta.humanVersion);

	File file = new File("result.txt");
	file.write(build.meta.humanVersion);
	
	file.saveToTheProperPlaceToBeVisibleInCloudBuild();
	return 1;
}
```