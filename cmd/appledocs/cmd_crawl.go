package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/tmc/appledocs/internal/crawler"
)

var crawlCmd = &cobra.Command{
	Use:   "crawl [framework]",
	Short: "Crawl and cache Apple documentation",
	Long: `Crawl Apple's developer documentation and cache it locally.

You can crawl a specific framework by name:
  appledocs crawl Foundation
  appledocs crawl AppKit

Or crawl from a custom entry point:
  appledocs crawl --entry-point /tutorials/data/documentation/technologies.json`,
	Args: cobra.MaximumNArgs(1),
	RunE: runCrawl,
}

var (
	crawlEntryPoint   string
	crawlOutputDir    string
	crawlBadURLsFile  string
	crawlExcludePaths string
	crawlDelay        int
	crawlRateLimit    float64
	crawlMaxTime      int
	crawlSkipSymbols  bool
	crawlPrintURLs    bool
	crawlValidate     bool
	crawlChecksum     bool
	crawlFetchBoth    bool
)

func init() {
	rootCmd.AddCommand(crawlCmd)

	// Crawl-specific flags
	crawlCmd.Flags().StringVar(&crawlEntryPoint, "entry-point", "/tutorials/data/documentation/technologies.json",
		"entry point to start crawling from")
	crawlCmd.Flags().StringVar(&crawlOutputDir, "output", "output",
		"directory to store mirrored content")
	crawlCmd.Flags().StringVar(&crawlBadURLsFile, "bad-urls-file", "",
		"file containing URLs to skip (default: $cache/known-bad-urls.txt)")
	crawlCmd.Flags().StringVar(&crawlExcludePaths, "exclude-paths", "en-US/docs/Mozilla",
		"comma-separated list of paths to exclude from crawling")
	crawlCmd.Flags().IntVar(&crawlDelay, "delay", 0,
		"delay between requests in milliseconds")
	crawlCmd.Flags().Float64Var(&crawlRateLimit, "rate-limit", 10.0,
		"requests per second rate limit (0 = no limit)")
	crawlCmd.Flags().IntVar(&crawlMaxTime, "max-time", 3600,
		"maximum time to run in seconds")
	crawlCmd.Flags().BoolVar(&crawlSkipSymbols, "skip-symbols", false,
		"skip individual symbol level documentation")
	crawlCmd.Flags().BoolVar(&crawlPrintURLs, "print-urls", false,
		"only print discovered URLs from entry point and exit")
	crawlCmd.Flags().BoolVar(&crawlValidate, "validate-cache", false,
		"validate cache integrity on startup")
	crawlCmd.Flags().BoolVar(&crawlChecksum, "checksum-validation", false,
		"enable enhanced checksum-based cache validation")
	crawlCmd.Flags().BoolVar(&crawlFetchBoth, "fetch-both-languages", false,
		"fetch both Swift and Objective-C variants")
}

func runCrawl(cmd *cobra.Command, args []string) error {
	// Initialize logger
	initLogger(logLevel, verbose)

	// Resolve framework name to entry point if provided as argument
	var frameworkName string
	if len(args) > 0 {
		frameworkName = args[0]
		ctx := context.Background()
		resolvedName, err := resolveFrameworkEntryPoint(ctx, cacheDir, baseURL, frameworkName)
		if err != nil {
			return fmt.Errorf("failed to resolve framework %q: %w", frameworkName, err)
		}
		crawlEntryPoint = resolvedName
		logger.Info("Resolved framework to entry point", "framework", frameworkName, "entry_point", resolvedName)
	}

	// Set default bad URLs file if not specified
	if crawlBadURLsFile == "" {
		crawlBadURLsFile = filepath.Join(cacheDir, "known-bad-urls.txt")
	}

	// Setup cancellable context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(crawlMaxTime)*time.Second)
	defer cancel()

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		logger.Info("Received interrupt signal, shutting down gracefully...")
		cancel()
	}()

	// Create output directory
	if err := os.MkdirAll(crawlOutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Run the crawler
	logger.Info("Starting crawler",
		"entry_point", crawlEntryPoint,
		"cache_dir", cacheDir,
		"output_dir", crawlOutputDir,
		"concurrency", concurrency,
		"rate_limit", crawlRateLimit,
	)

	cfg := &crawler.Config{
		BaseURL:            baseURL,
		EntryPoint:         crawlEntryPoint,
		OutputDir:          crawlOutputDir,
		CacheDir:           cacheDir,
		BadURLsFile:        crawlBadURLsFile,
		ExcludePaths:       crawlExcludePaths,
		Concurrency:        concurrency,
		RateLimit:          crawlRateLimit,
		ForceRefresh:       forceRefresh,
		Timeout:            time.Duration(timeout) * time.Second,
		SkipSymbols:        crawlSkipSymbols,
		PrettyJSON:         prettyJSON,
		FetchBothLanguages: crawlFetchBoth,
		ChecksumValidation: crawlChecksum,
		Verbose:            verbose,
		Logger:             logger,
	}

	c := crawler.New(cfg)
	err := c.Run(ctx, cfg)

	if err != nil {
		return fmt.Errorf("crawl failed: %w", err)
	}

	logger.Info("Crawl completed successfully")
	return nil
}
