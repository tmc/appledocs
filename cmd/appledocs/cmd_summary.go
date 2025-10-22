package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary <framework>",
	Short: "Generate a framework summary",
	Long: `Generate a detailed summary of a cached framework including:
  - Symbol statistics (classes, protocols, enums, etc.)
  - Content statistics (methods, properties, API coverage)
  - Platform information
  - Top-level categories

Examples:
  appledocs summary Foundation
  appledocs summary AppKit --json`,
	Args: cobra.ExactArgs(1),
	RunE: runSummary,
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}

func runSummary(cmd *cobra.Command, args []string) error {
	// Initialize logger
	initLogger(logLevel, verbose)

	frameworkName := args[0]

	logger.Info("Generating framework summary", "framework", frameworkName)

	err := generateFrameworkSummary(cacheDir, frameworkName, jsonOutput)
	if err != nil {
		return fmt.Errorf("failed to generate summary: %w", err)
	}

	return nil
}
