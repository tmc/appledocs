package crawler

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"
)

// MetricsSnapshot represents a comprehensive snapshot of all metrics
type MetricsSnapshot struct {
	// Basic metrics
	Processed      int `json:"processed"`
	CacheHits      int `json:"cache_hits"`
	CacheMisses    int `json:"cache_misses"`
	Errors         int `json:"errors"`
	SkippedURLs    int `json:"skipped_urls"`
	SkippedSymbols int `json:"skipped_symbols"`

	// Enhanced metrics
	StartTime            time.Time     `json:"start_time"`
	RuntimeDuration      time.Duration `json:"runtime_duration"`
	TotalBytesDownloaded int64         `json:"total_bytes_downloaded"`
	TotalBytesFromCache  int64         `json:"total_bytes_from_cache"`
	AvgResponseTime      time.Duration `json:"avg_response_time"`
	RequestCount         int           `json:"request_count"`
	HTTPErrors           map[int]int   `json:"http_errors"`
	RetryCount           int           `json:"retry_count"`
	FrameworkCount       int           `json:"framework_count"`
	ClassCount           int           `json:"class_count"`
	MethodCount          int           `json:"method_count"`

	// Calculated metrics
	CacheHitRate           float64       `json:"cache_hit_rate"`
	DownloadRate           float64       `json:"download_rate_mbps"`
	ProcessingRate         float64       `json:"processing_rate_per_sec"`
	EstimatedTimeRemaining time.Duration `json:"estimated_time_remaining"`
}

// getEnhancedMetrics returns comprehensive metrics snapshot
func (c *Crawler) getEnhancedMetrics() MetricsSnapshot {
	c.statsMutex.Lock()
	defer c.statsMutex.Unlock()
	c.entriesMutex.Lock()
	processed := len(c.jsonEntries)
	c.entriesMutex.Unlock()

	now := time.Now()
	runtime := now.Sub(c.startTime)

	// Calculate derived metrics
	var cacheHitRate float64
	totalRequests := c.cacheHits + c.cacheMisses
	if totalRequests > 0 {
		cacheHitRate = float64(c.cacheHits) / float64(totalRequests) * 100
	}

	// Download rate in MB/s
	var downloadRate float64
	if runtime.Seconds() > 0 {
		totalMB := float64(c.totalBytesDownloaded) / 1024 / 1024
		downloadRate = totalMB / runtime.Seconds()
	}

	// Processing rate (files per second)
	var processingRate float64
	if runtime.Seconds() > 0 {
		processingRate = float64(processed) / runtime.Seconds()
	}

	// Estimate time remaining
	var estimatedTimeRemaining time.Duration
	totalURLs := 0
	c.visitedURLs.Range(func(_, _ interface{}) bool {
		totalURLs++
		return true
	})

	if processed > 0 && totalURLs > processed && processingRate > 0 {
		remaining := totalURLs - processed
		estimatedTimeRemaining = time.Duration(float64(remaining)/processingRate) * time.Second
	}

	return MetricsSnapshot{
		Processed:              processed,
		CacheHits:              c.cacheHits,
		CacheMisses:            c.cacheMisses,
		Errors:                 c.errors,
		SkippedURLs:            c.skippedURLs,
		SkippedSymbols:         c.skippedSymbols,
		StartTime:              c.startTime,
		RuntimeDuration:        runtime,
		TotalBytesDownloaded:   c.totalBytesDownloaded,
		TotalBytesFromCache:    c.totalBytesFromCache,
		AvgResponseTime:        c.avgResponseTime,
		RequestCount:           c.requestCount,
		HTTPErrors:             c.httpErrors,
		RetryCount:             c.retryCount,
		FrameworkCount:         c.frameworkCount,
		ClassCount:             c.classCount,
		MethodCount:            c.methodCount,
		CacheHitRate:           cacheHitRate,
		DownloadRate:           downloadRate,
		ProcessingRate:         processingRate,
		EstimatedTimeRemaining: estimatedTimeRemaining,
	}
}

// reportProgress reports progress periodically
func (c *Crawler) reportProgress(ctx context.Context, ticker *time.Ticker, cfg *Config) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			metrics := c.getEnhancedMetrics()

			log.Printf("Progress: %d files processed | Cache: %d hits (%.1f%%), %d misses | %.2f MB/s | Avg: %v/req | Errors: %d | Retries: %d",
				metrics.Processed,
				metrics.CacheHits,
				metrics.CacheHitRate,
				metrics.CacheMisses,
				metrics.DownloadRate,
				metrics.AvgResponseTime,
				metrics.Errors,
				metrics.RetryCount)

			// Additional metrics when verbose
			if cfg.Verbose && metrics.EstimatedTimeRemaining > 0 {
				log.Printf("Content: %d frameworks, %d classes, %d methods | ETA: %v",
					metrics.FrameworkCount,
					metrics.ClassCount,
					metrics.MethodCount,
					metrics.EstimatedTimeRemaining.Round(time.Second))
			}
		}
	}
}

// reportFinalStats reports final statistics
func (c *Crawler) reportFinalStats(cfg *Config) {
	metrics := c.getEnhancedMetrics()
	log.Printf("Final statistics:")
	log.Printf("  - Runtime: %v", metrics.RuntimeDuration.Round(time.Second))
	log.Printf("  - Processed: %d JSON files (%.2f files/sec)", metrics.Processed, metrics.ProcessingRate)
	log.Printf("  - Cache: %d hits (%.1f%%), %d misses", metrics.CacheHits, metrics.CacheHitRate, metrics.CacheMisses)
	log.Printf("  - Data: %.2f MB downloaded, %.2f MB from cache (%.2f MB/s)",
		float64(metrics.TotalBytesDownloaded)/1024/1024,
		float64(metrics.TotalBytesFromCache)/1024/1024,
		metrics.DownloadRate)
	log.Printf("  - Network: %d requests, avg %v/req, %d retries",
		metrics.RequestCount, metrics.AvgResponseTime, metrics.RetryCount)
	log.Printf("  - Content: %d frameworks, %d classes, %d methods",
		metrics.FrameworkCount, metrics.ClassCount, metrics.MethodCount)
	log.Printf("  - Errors: %d", metrics.Errors)
	log.Printf("  - Skipped URLs: %d", metrics.SkippedURLs)
	if cfg.SkipSymbols {
		log.Printf("  - Skipped symbol URLs: %d", metrics.SkippedSymbols)
	}

	// Report HTTP errors if any
	if len(metrics.HTTPErrors) > 0 {
		log.Printf("  - HTTP Errors by status code:")
		for statusCode, count := range metrics.HTTPErrors {
			log.Printf("    - %d: %d errors", statusCode, count)
		}
	}
}

// ExportMetrics exports metrics to a JSON file
func (c *Crawler) ExportMetrics(metricsPath string) error {
	metrics := c.getEnhancedMetrics()
	metricsJSON, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(metricsPath, metricsJSON, 0644)
}
