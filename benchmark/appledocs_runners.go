// Package benchmark provides specialized runners for appledocs operations
package benchmark

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// AppledocsBaselineRunner implements the current appledocs behavior
type AppledocsBaselineRunner struct {
	client      *http.Client
	cache       map[string][]byte
	cacheMutex  sync.RWMutex
	visitedURLs map[string]bool
	urlMutex    sync.RWMutex
}

func NewAppledocsBaselineRunner() *AppledocsBaselineRunner {
	return &AppledocsBaselineRunner{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		cache:       make(map[string][]byte),
		visitedURLs: make(map[string]bool),
	}
}

func (r *AppledocsBaselineRunner) Name() string {
	return "appledocs-baseline"
}

func (r *AppledocsBaselineRunner) Setup() error {
	// Initialize HTTP client with appledocs settings
	r.client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}
	return nil
}

func (r *AppledocsBaselineRunner) Teardown() error {
	r.client.CloseIdleConnections()
	return nil
}

func (r *AppledocsBaselineRunner) Run(ctx context.Context, op Operation, data []byte) (interface{}, error) {
	switch op.Function {
	case "json.Unmarshal":
		return r.parseJSON(data)
	case "extractJSONURLs":
		return r.extractURLs(data)
	case "processURLBatch":
		return r.processURLBatch(ctx, op.Parameters)
	case "generateMarkdown":
		return r.generateMarkdown(data)
	case "validateJSONStructure":
		return r.validateStructure(data)
	default:
		return nil, fmt.Errorf("unknown operation: %s", op.Function)
	}
}

func (r *AppledocsBaselineRunner) parseJSON(data []byte) (interface{}, error) {
	var result map[string]interface{}
	return result, json.Unmarshal(data, &result)
}

func (r *AppledocsBaselineRunner) extractURLs(data []byte) (interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	urls := make([]string, 0, 100)
	r.extractURLsFromValue(result, &urls)
	return urls, nil
}

func (r *AppledocsBaselineRunner) extractURLsFromValue(v interface{}, urls *[]string) {
	switch val := v.(type) {
	case map[string]interface{}:
		// Check for destination objects with identifiers
		if _, hasType := val["type"]; hasType {
			if identifier, ok := val["identifier"].(string); ok {
				if strings.HasPrefix(identifier, "doc://") {
					docPath := strings.TrimPrefix(identifier, "doc://")
					parts := strings.SplitN(docPath, "/", 2)
					if len(parts) > 1 {
						jsonURL := parts[1] + ".json"
						*urls = append(*urls, jsonURL)
					}
				}
			}
		}

		for k, v := range val {
			if (k == "url" || strings.HasSuffix(k, "URL") || strings.HasSuffix(k, "Uri")) && v != nil {
				if urlStr, ok := v.(string); ok && strings.HasSuffix(urlStr, ".json") {
					*urls = append(*urls, urlStr)
				}
			}
			r.extractURLsFromValue(v, urls)
		}
	case []interface{}:
		for _, item := range val {
			r.extractURLsFromValue(item, urls)
		}
	}
}

func (r *AppledocsBaselineRunner) processURLBatch(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	batchSize := 10
	if bs, ok := params["batch_size"].(int); ok {
		batchSize = bs
	}

	processed := 0
	errors := 0

	// Simulate processing a batch of URLs
	for i := 0; i < batchSize; i++ {
		select {
		case <-ctx.Done():
			return map[string]int{"processed": processed, "errors": errors}, ctx.Err()
		default:
			// Simulate URL processing
			time.Sleep(10 * time.Millisecond)
			processed++
		}
	}

	return map[string]int{"processed": processed, "errors": errors}, nil
}

func (r *AppledocsBaselineRunner) generateMarkdown(data []byte) (interface{}, error) {
	var doc map[string]interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	
	// Generate markdown header
	if title, ok := doc["title"].(string); ok {
		buf.WriteString(fmt.Sprintf("# %s\n\n", title))
	}

	// Generate abstract
	if abstract, ok := doc["abstract"].([]interface{}); ok {
		for _, item := range abstract {
			if m, ok := item.(map[string]interface{}); ok {
				if text, ok := m["text"].(string); ok {
					buf.WriteString(text + "\n\n")
				}
			}
		}
	}

	return buf.String(), nil
}

func (r *AppledocsBaselineRunner) validateStructure(data []byte) (interface{}, error) {
	var doc map[string]interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return false, err
	}

	// Check required fields
	required := []string{"metadata", "abstract", "identifier"}
	missing := []string{}
	
	for _, field := range required {
		if _, exists := doc[field]; !exists {
			missing = append(missing, field)
		}
	}

	return map[string]interface{}{
		"valid":   len(missing) == 0,
		"missing": missing,
	}, nil
}

// AppledocsOptimizedRunner implements optimized JSON parsing
type AppledocsOptimizedRunner struct {
	*AppledocsBaselineRunner
	decoder *json.Decoder
	buffer  *bytes.Buffer
}

func NewAppledocsOptimizedRunner() *AppledocsOptimizedRunner {
	return &AppledocsOptimizedRunner{
		AppledocsBaselineRunner: NewAppledocsBaselineRunner(),
		buffer:                  &bytes.Buffer{},
	}
}

func (r *AppledocsOptimizedRunner) Name() string {
	return "appledocs-optimized"
}

func (r *AppledocsOptimizedRunner) parseJSON(data []byte) (interface{}, error) {
	// Use decoder for better memory efficiency
	r.buffer.Reset()
	r.buffer.Write(data)
	
	decoder := json.NewDecoder(r.buffer)
	decoder.UseNumber() // Avoid float64 for numbers
	
	var result map[string]interface{}
	return result, decoder.Decode(&result)
}

// AppledocsStreamingRunner implements streaming JSON processing
type AppledocsStreamingRunner struct {
	*AppledocsBaselineRunner
}

func NewAppledocsStreamingRunner() *AppledocsStreamingRunner {
	return &AppledocsStreamingRunner{
		AppledocsBaselineRunner: NewAppledocsBaselineRunner(),
	}
}

func (r *AppledocsStreamingRunner) Name() string {
	return "appledocs-streaming"
}

func (r *AppledocsStreamingRunner) Run(ctx context.Context, op Operation, data []byte) (interface{}, error) {
	switch op.Function {
	case "extractJSONURLs":
		return r.extractURLsStreaming(data)
	default:
		return r.AppledocsBaselineRunner.Run(ctx, op, data)
	}
}

func (r *AppledocsStreamingRunner) extractURLsStreaming(data []byte) (interface{}, error) {
	urls := make([]string, 0, 100)
	decoder := json.NewDecoder(bytes.NewReader(data))
	
	// Stream through the JSON looking for URLs
	if err := r.streamExtractURLs(decoder, &urls); err != nil && err != io.EOF {
		return nil, err
	}
	
	return urls, nil
}

func (r *AppledocsStreamingRunner) streamExtractURLs(decoder *json.Decoder, urls *[]string) error {
	t, err := decoder.Token()
	if err != nil {
		return err
	}
	
	switch token := t.(type) {
	case json.Delim:
		switch token {
		case '{':
			// Process object
			for decoder.More() {
				// Read key
				keyToken, err := decoder.Token()
				if err != nil {
					return err
				}
				
				key, ok := keyToken.(string)
				if !ok {
					continue
				}
				
				// Check if this is a URL field
				if key == "url" || strings.HasSuffix(key, "URL") || strings.HasSuffix(key, "Uri") {
					// Read value
					valueToken, err := decoder.Token()
					if err != nil {
						return err
					}
					
					if urlStr, ok := valueToken.(string); ok && strings.HasSuffix(urlStr, ".json") {
						*urls = append(*urls, urlStr)
					}
				} else {
					// Recurse into value
					if err := r.streamExtractURLs(decoder, urls); err != nil {
						return err
					}
				}
			}
			// Read closing }
			_, err := decoder.Token()
			return err
			
		case '[':
			// Process array
			for decoder.More() {
				if err := r.streamExtractURLs(decoder, urls); err != nil {
					return err
				}
			}
			// Read closing ]
			_, err := decoder.Token()
			return err
		}
	}
	
	return nil
}

// AppledocsConcurrentRunner implements concurrent processing optimizations
type AppledocsConcurrentRunner struct {
	*AppledocsBaselineRunner
	workerPool *sync.Pool
	semaphore  chan struct{}
}

func NewAppledocsConcurrentRunner() *AppledocsConcurrentRunner {
	return &AppledocsConcurrentRunner{
		AppledocsBaselineRunner: NewAppledocsBaselineRunner(),
		workerPool: &sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
		semaphore: make(chan struct{}, 10), // Limit concurrent operations
	}
}

func (r *AppledocsConcurrentRunner) Name() string {
	return "appledocs-concurrent"
}

func (r *AppledocsConcurrentRunner) processURLBatch(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	batchSize := 10
	if bs, ok := params["batch_size"].(int); ok {
		batchSize = bs
	}
	
	// Process URLs concurrently
	results := make(chan int, batchSize)
	errors := make(chan error, batchSize)
	
	var wg sync.WaitGroup
	for i := 0; i < batchSize; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			
			select {
			case r.semaphore <- struct{}{}:
				defer func() { <-r.semaphore }()
			case <-ctx.Done():
				errors <- ctx.Err()
				return
			}
			
			// Simulate URL processing
			time.Sleep(10 * time.Millisecond)
			results <- 1
		}(i)
	}
	
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()
	
	processed := 0
	errorCount := 0
	
	for {
		select {
		case <-ctx.Done():
			return map[string]int{"processed": processed, "errors": errorCount}, ctx.Err()
		case res, ok := <-results:
			if !ok {
				return map[string]int{"processed": processed, "errors": errorCount}, nil
			}
			processed += res
		case err, ok := <-errors:
			if ok && err != nil {
				errorCount++
			}
		}
	}
}

// AppledocsMemoryOptimizedRunner implements memory-optimized operations
type AppledocsMemoryOptimizedRunner struct {
	*AppledocsBaselineRunner
	bufferPool *sync.Pool
}

func NewAppledocsMemoryOptimizedRunner() *AppledocsMemoryOptimizedRunner {
	return &AppledocsMemoryOptimizedRunner{
		AppledocsBaselineRunner: NewAppledocsBaselineRunner(),
		bufferPool: &sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}

func (r *AppledocsMemoryOptimizedRunner) Name() string {
	return "appledocs-memory-optimized"
}

func (r *AppledocsMemoryOptimizedRunner) generateMarkdown(data []byte) (interface{}, error) {
	// Use pooled buffer to reduce allocations
	buf := r.bufferPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		r.bufferPool.Put(buf)
	}()
	
	// Parse JSON with streaming to avoid full load
	decoder := json.NewDecoder(bytes.NewReader(data))
	
	// Generate markdown incrementally
	if err := r.generateMarkdownStreaming(decoder, buf); err != nil {
		return nil, err
	}
	
	// Return a copy to avoid buffer reuse issues
	result := make([]byte, buf.Len())
	copy(result, buf.Bytes())
	
	return string(result), nil
}

func (r *AppledocsMemoryOptimizedRunner) generateMarkdownStreaming(decoder *json.Decoder, buf *bytes.Buffer) error {
	// Implementation would stream through JSON and generate markdown incrementally
	// This is a simplified version
	var doc map[string]interface{}
	if err := decoder.Decode(&doc); err != nil {
		return err
	}
	
	if title, ok := doc["title"].(string); ok {
		buf.WriteString(fmt.Sprintf("# %s\n\n", title))
	}
	
	return nil
}

// RunnerFactory creates benchmark runners for different phases
func RunnerFactory(phase Phase) BenchmarkRunner {
	switch phase {
	case PhaseBaseline:
		return NewAppledocsBaselineRunner()
	case PhaseJSONOptimize:
		return NewAppledocsOptimizedRunner()
	case PhaseStreaming:
		return NewAppledocsStreamingRunner()
	case PhaseConcurrency:
		return NewAppledocsConcurrentRunner()
	case PhaseMemoryOpt:
		return NewAppledocsMemoryOptimizedRunner()
	default:
		return NewAppledocsBaselineRunner()
	}
}