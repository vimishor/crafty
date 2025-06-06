package version

import (
	"fmt"
	"runtime"
)

// Version number
var Version string = "0.1.0"

var (
	// GitCommit returns the git commit that was compiled. This will be filled in by the compiler.
	GitCommit string

	// BuildDate returns the date the binary was built
	BuildDate = ""

	// GoVersion returns the version of the go runtime used to compile the binary
	GoVersion = runtime.Version()

	// OsArch returns the os and arch used to build the binary
	OsArch = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

// String returns current version as string.
func String() string {
	return fmt.Sprintf("crafty %s", Version)
}

func Short() string {
	return Version
}

func Long() string {
	return fmt.Sprintf("version: %s, commit: %s, build: %s, %s", Version, GitCommit, BuildDate, OsArch)
}
