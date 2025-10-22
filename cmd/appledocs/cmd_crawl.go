package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
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
	crawlPrune        bool
	crawlMaxDepth     int
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
	crawlCmd.Flags().BoolVar(&crawlPrune, "prune", false,
		"prune old cache files (HTML files and pre-index.json structure)")
	crawlCmd.Flags().IntVar(&crawlMaxDepth, "max-depth", 0,
		"maximum link depth to follow from entry point (0 = unlimited)")
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

	// Prune old cache if requested
	if crawlPrune {
		logger.Info("Pruning old cache files...")
		if err := pruneOldCache(cacheDir, verbose); err != nil {
			logger.Warn("Failed to prune cache", "error", err)
		}
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
		MaxDepth:           crawlMaxDepth,
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

// pruneOldCache removes old cache files:
// 1. HTML files (should be JSON)
// 2. Files in directories where index.json now exists (old structure)
func pruneOldCache(cacheDir string, verbose bool) error {
	var htmlCount, oldStructureCount, languageVariantCount int
	var totalBytesRemoved int64

	// Walk the cache directory
	err := filepath.Walk(cacheDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Skip files that were removed during walk
			if os.IsNotExist(err) {
				return nil
			}
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Skip if this is an index.json or .etag file (new structure)
		if filepath.Base(path) == "index.json" || filepath.Ext(path) == ".etag" {
			return nil
		}

		// Check if this file should be pruned
		shouldPrune := false
		reason := ""

		// Check if it's an HTML file
		if isHTMLFile(path) {
			shouldPrune = true
			reason = "HTML file"
			htmlCount++
		}

		// Check if there's an index.json in the same directory
		// (indicates this is old structure being replaced)
		dir := filepath.Dir(path)
		indexPath := filepath.Join(dir, "index.json")
		if _, err := os.Stat(indexPath); err == nil {
			// index.json exists, this is old structure
			if !shouldPrune { // Don't double-count
				shouldPrune = true
				reason = "old structure (index.json exists)"
				oldStructureCount++
			}
		}

		// Check if this is a language variant file (has .language%3D in the name)
		if !shouldPrune && (strings.Contains(path, ".language%3D") || strings.Contains(path, "?language=")) {
			shouldPrune = true
			reason = "language variant duplicate"
			languageVariantCount++
		}

		// Prune the file if needed
		if shouldPrune {
			size := info.Size()
			if verbose {
				logger.Info("Pruning cache file", "path", path, "reason", reason, "size", size)
			}

			// Remove the file
			if err := os.Remove(path); err != nil {
				logger.Warn("Failed to remove file", "path", path, "error", err)
				return nil // Continue walking
			}

			// Also remove associated .etag file if it exists
			etagPath := path + ".etag"
			if _, err := os.Stat(etagPath); err == nil {
				os.Remove(etagPath)
			}

			// Remove query-parameter variants if they exist
			// (e.g., file.json.language%3Dobjc, file.json.language%3Dswift)
			baseWithoutExt := path
			if ext := filepath.Ext(path); ext != "" {
				baseWithoutExt = path[:len(path)-len(ext)]
			}

			// Look for files matching base.json.language%3D*
			if matches, err := filepath.Glob(baseWithoutExt + ".*"); err == nil {
				for _, match := range matches {
					if match != path && (filepath.Ext(match) != ".etag") {
						os.Remove(match)
						os.Remove(match + ".etag")
						if fi, err := os.Stat(match); err == nil {
							totalBytesRemoved += fi.Size()
						}
					}
				}
			}

			totalBytesRemoved += size
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk cache directory: %w", err)
	}

	logger.Info("Cache pruning complete",
		"html_files_removed", htmlCount,
		"old_structure_files_removed", oldStructureCount,
		"language_variant_files_removed", languageVariantCount,
		"total_bytes_removed", totalBytesRemoved,
		"total_mb_removed", float64(totalBytesRemoved)/(1024*1024))

	return nil
}

// isHTMLFile checks if a file contains HTML content
func isHTMLFile(path string) bool {
	// Read first 512 bytes to detect file type
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil && n == 0 {
		return false
	}

	// Check for HTML markers
	content := string(buf[:n])
	return contains(content, "<!DOCTYPE html") ||
	       contains(content, "<html") ||
	       contains(content, "<HTML")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s[:len(substr)] == substr ||
		len(s) > len(substr) && stringContains(s, substr))
}

func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
