package config

import (
	"fmt"
	"io"
	"runtime"
)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	GitTag      string
	GitBranch   string
	GitHash     string
	GoBuildTime string
)

func PrintVersion(w io.Writer) {
	fmt.Fprintf(w, "\nBuild:\n")
	if GitTag != "" {
		fmt.Fprintf(w, "  Build: %v (branch: %q hash:%q)\n", GitTag, GitBranch, GitHash)
	}
	if GoBuildTime != "" {
		fmt.Fprintf(w, "   Time: %v\n", GoBuildTime)
	}
	fmt.Fprintf(w, "     Go: %v (%v/%v)\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
