package main

import (
	"context"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list-frameworks",
	Short: "List available Apple frameworks",
	Long: `List all available Apple frameworks that can be crawled.

This command fetches the list of frameworks from Apple's technologies.json
and displays them with their descriptions and platform availability.

Examples:
  appledocs list-frameworks
  appledocs list-frameworks --json
  appledocs list-frameworks --refresh
  appledocs list-frameworks --platform macOS
  appledocs list-frameworks --platform iOS --min-version 15.0
  appledocs list-frameworks --pattern "^Core"
  appledocs list-frameworks --beta`,
	RunE: runList,
}

var (
	listRefresh    bool
	listPlatform   string
	listMinVersion string
	listPattern    string
	listShowAll    bool
	listBeta       bool
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&listRefresh, "refresh", false,
		"refresh the cached framework list from technologies.json")
	listCmd.Flags().StringVar(&listPlatform, "platform", "",
		"filter by platform (iOS, macOS, tvOS, watchOS, visionOS, etc.)")
	listCmd.Flags().StringVar(&listMinVersion, "min-version", "",
		"filter by minimum platform version (e.g., 15.0)")
	listCmd.Flags().StringVar(&listPattern, "pattern", "",
		"filter by name pattern (regex)")
	listCmd.Flags().BoolVar(&listShowAll, "all", false,
		"show all details including deprecated and beta frameworks")
	listCmd.Flags().BoolVar(&listBeta, "beta", false,
		"show only beta APIs")
}

func runList(cmd *cobra.Command, args []string) error {
	// Initialize logger
	initLogger(logLevel, verbose)

	ctx := context.Background()
	filters := ListFilters{
		Platform:   listPlatform,
		MinVersion: listMinVersion,
		Pattern:    listPattern,
		ShowAll:    listShowAll,
		Beta:       listBeta,
	}
	err := listAvailableFrameworks(ctx, cacheDir, baseURL, jsonOutput, listRefresh, filters)
	if err != nil {
		return err
	}

	return nil
}
