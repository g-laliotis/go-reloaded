// Package version provides build and version information for go-reloaded.
package version

import (
	"fmt"
	"runtime"
)

const (
	// Version is the current version of go-reloaded
	Version = "1.0.0"
	
	// Name is the application name
	Name = "go-reloaded"
	
	// Description is a brief description of the application
	Description = "A command-line text transformation tool with pipeline architecture"
)

// BuildInfo contains build and runtime information
type BuildInfo struct {
	Version   string
	GoVersion string
	OS        string
	Arch      string
}

// GetBuildInfo returns current build information
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version:   Version,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}

// String returns a formatted version string
func (b BuildInfo) String() string {
	return fmt.Sprintf("%s v%s (built with %s for %s/%s)", 
		Name, b.Version, b.GoVersion, b.OS, b.Arch)
}

// Short returns a short version string
func Short() string {
	return fmt.Sprintf("%s v%s", Name, Version)
}