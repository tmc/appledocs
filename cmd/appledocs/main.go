// Command appledocs mirrors Apple documentation JSON files to disk.
package main

import (
	"log/slog"
	"os"
)

// Build-time variables (set via -ldflags)
var (
	version   = "dev"
	commit    = "none"
	buildTime = "unknown"
)

// Global structured logger
var logger *slog.Logger

func main() {
	if err := Execute(); err != nil {
		os.Exit(1)
	}
}
