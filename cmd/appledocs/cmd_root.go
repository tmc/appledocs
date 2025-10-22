package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Global flags
	cacheDir     string
	baseURL      string
	verbose      bool
	logLevel     string
	jsonOutput   bool
	prettyJSON   bool
	timeout      int
	concurrency  int
	forceRefresh bool
)

var rootCmd = &cobra.Command{
	Use:   "appledocs",
	Short: "Apple Developer Documentation crawler and analyzer",
	Long: `appledocs is a tool for crawling, caching, and analyzing Apple's developer documentation.
It provides several subcommands for different operations:

  crawl          - Crawl and cache Apple documentation
  list           - List available frameworks
  summary        - Generate framework summaries
  demos          - List and download demo code
  analyze        - Analyze documentation structure
  gentypes       - Generate Go types from documentation
  html           - Generate HTML documentation
  markdown       - Generate Markdown documentation`,
	Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, buildTime),
}

func init() {
	// Global flags available to all subcommands
	rootCmd.PersistentFlags().StringVar(&cacheDir, "cache", getDefaultCacheDir(), "directory to store HTTP cache")
	rootCmd.PersistentFlags().StringVar(&baseURL, "base", "https://developer.apple.com", "base URL for Apple docs")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbose logging")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level: debug, info, warn, error")
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "output in JSON format")
	rootCmd.PersistentFlags().BoolVar(&prettyJSON, "pretty", true, "pretty-print JSON files")
	rootCmd.PersistentFlags().IntVar(&timeout, "timeout", 30, "HTTP request timeout in seconds")
	rootCmd.PersistentFlags().IntVar(&concurrency, "concurrency", 1, "number of concurrent operations")
	rootCmd.PersistentFlags().BoolVar(&forceRefresh, "force", false, "force refresh cached content")
}

func getDefaultCacheDir() string {
	home := os.Getenv("HOME")
	if home == "" {
		home = "."
	}
	return home + "/.appledocs/cache"
}

func Execute() error {
	return rootCmd.Execute()
}
