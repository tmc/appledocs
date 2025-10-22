package main

import (
	"context"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available Apple frameworks",
	Long: `List all available Apple frameworks that can be crawled.

This command fetches the list of frameworks from Apple's technologies.json
and displays them with their descriptions.

Examples:
  appledocs list
  appledocs list --json
  appledocs list --refresh`,
	RunE: runList,
}

var listRefresh bool

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&listRefresh, "refresh", false,
		"refresh the cached framework list from technologies.json")
}

func runList(cmd *cobra.Command, args []string) error {
	// Initialize logger
	initLogger(logLevel, verbose)

	ctx := context.Background()
	err := listAvailableFrameworks(ctx, cacheDir, baseURL, jsonOutput, listRefresh)
	if err != nil {
		return err
	}

	return nil
}
