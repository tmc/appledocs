package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary <framework>[.symbol]",
	Short: "Generate a framework or symbol summary",
	Long: `Generate a detailed summary of a cached framework or specific symbol.

Framework summary includes:
  - Symbol statistics (classes, protocols, enums, etc.)
  - Content statistics (methods, properties, API coverage)
  - Platform information
  - Top-level categories

Symbol summary includes:
  - Symbol metadata (kind, role, platforms)
  - Methods and properties
  - Inheritance hierarchy
  - Protocol conformance

Examples:
  appledocs summary Foundation
  appledocs summary AppKit --json
  appledocs summary ObjectiveC.NSObject
  appledocs summary Foundation.NSString`,
	Args: cobra.ExactArgs(1),
	RunE: runSummary,
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}

func runSummary(cmd *cobra.Command, args []string) error {
	// Initialize logger
	initLogger(logLevel, verbose)

	arg := args[0]

	// Check if argument contains a dot (Framework.Symbol syntax)
	if strings.Contains(arg, ".") {
		parts := strings.SplitN(arg, ".", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid format: expected Framework.Symbol, got %q", arg)
		}
		frameworkName := parts[0]
		symbolName := parts[1]

		logger.Info("Generating symbol summary", "framework", frameworkName, "symbol", symbolName)

		err := generateSymbolSummary(cacheDir, frameworkName, symbolName, jsonOutput)
		if err != nil {
			return fmt.Errorf("failed to generate symbol summary: %w", err)
		}
	} else {
		frameworkName := arg

		logger.Info("Generating framework summary", "framework", frameworkName)

		err := generateFrameworkSummary(cacheDir, frameworkName, jsonOutput)
		if err != nil {
			return fmt.Errorf("failed to generate framework summary: %w", err)
		}
	}

	return nil
}
