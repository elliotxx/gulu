//go:generate go run -mod vendor gen.go
//go:generate go fmt

package version

import (
	"fmt"

	_ "go-utils/gitutil"
)

var (
	ShortVersion   = "(dev)"
	GitSha1Version = "git-sha1"
	BuildDate      = "2020-01-01"
)

func PrintVersionInfo(printer func(string, ...interface{})) {
	printer("Release Version: %s", ShortVersion)
	printer("Git Commit Hash: %s", GitSha1Version)
	printer("Build Time: %s", BuildDate)
}

func GetVersionString() string {
	return fmt.Sprintf("%s; git: %s; build time: %s", ShortVersion, GitSha1Version, BuildDate)
}
