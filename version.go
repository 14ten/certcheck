package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// Set via -ldflags="-X main.version=..."
var (
	version = "dev"
	commit  = ""
)

func fullVersion() string {
	c := commit
	if c == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, s := range info.Settings {
				if s.Key == "vcs.revision" {
					c = s.Value
					break
				}
			}
		}
	}
	if c == "" {
		c = "unknown"
	}
	if len(c) > 7 {
		c = c[:7]
	}
	return fmt.Sprintf("certcheck %s (%s, %s/%s, %s)",
		version, c, runtime.GOOS, runtime.GOARCH, runtime.Version())
}
