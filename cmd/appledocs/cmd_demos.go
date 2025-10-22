package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var demosCmd = &cobra.Command{
	Use:   "demos [framework]",
	Short: "List and download demo code",
	Long: `List available demo code from cached documentation.
Optionally filter by framework name.

Examples:
  appledocs demos                    # List all demos
  appledocs demos Foundation         # List Foundation demos
  appledocs demos --download         # Download all demos
  appledocs demos Foundation --download --output ~/demos`,
	Args: cobra.MaximumNArgs(1),
	RunE: runDemos,
}

var (
	demosDownload  bool
	demosOutputDir string
)

func init() {
	rootCmd.AddCommand(demosCmd)

	demosCmd.Flags().BoolVar(&demosDownload, "download", false,
		"download demo code")
	demosCmd.Flags().StringVar(&demosOutputDir, "output", getDefaultDemosDir(),
		"directory to store downloaded demo code")
}

func getDefaultDemosDir() string {
	home := os.Getenv("HOME")
	if home == "" {
		home = "."
	}
	return filepath.Join(home, "go", "src", "github.com", "tmc", "appledocs-examples")
}

func runDemos(cmd *cobra.Command, args []string) error {
	// Initialize logger
	initLogger(logLevel, verbose)

	var frameworkName string
	if len(args) > 0 {
		frameworkName = args[0]
	}

	logger.Info("Listing demo code", "framework", frameworkName, "download", demosDownload)

	err := listDemoCode(cacheDir, frameworkName, demosDownload, demosOutputDir)
	if err != nil {
		return fmt.Errorf("failed to list demo code: %w", err)
	}

	return nil
}
