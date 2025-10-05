//go:build ignore

// Package main provides streaming JSON parsing implementations
// optimized for large Apple documentation files
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/tidwall/gjson"
	"github.com/valyala/fastjson"
)

// StreamingJSONProcessor defines the interface for streaming JSON processors
type StreamingJSONProcessor interface {
	ProcessStream(reader io.Reader, callback URLCallback) error
	ProcessFile(filename string, callback URLCallback) error
	Name() string
}

// URLCallback is called for each URL found during streaming
type URLCallback func(url string, source string) error

// StreamingResult contains the results of streaming processing
type StreamingResult struct {
	ProcessorName  string
	URLsFound      int
	BytesProcessed int64
	Error          error
}

// StandardStreamingProcessor uses encoding/json with streaming
type StandardStreamingProcessor struct{}

func (p *StandardStreamingProcessor) Name() string {
	return "encoding/json-streaming"
}

func (p *StandardStreamingProcessor) ProcessStream(reader io.Reader, callback URLCallback) error {
	decoder := json.NewDecoder(reader)
	
	// Use decoder.Token() to parse incrementally
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("decode token: %v", err)
		}
		
		// Handle different token types
		switch t := token.(type) {
		case json.Delim:
			if t == '{' {
				// Start of object - we could handle this more efficiently
				// by parsing only the parts we need
			}
		case string:
			// Check if this string looks like a URL
			if strings.HasSuffix(t, ".json") {
				if err := callback(t, "token"); err != nil {
					return err
				}
			}
		}
	}
	
	return nil
}

func (p *StandardStreamingProcessor) ProcessFile(filename string, callback URLCallback) error {
	// For files, we can still use the full document approach
	// but with streaming decoder
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	var doc DocJSONData
	
	if err := decoder.Decode(&doc); err != nil {
		return fmt.Errorf("decode document: %v", err)
	}
	
	// Extract URLs using the structured approach
	return p.extractURLsFromDoc(&doc, callback)
}

func (p *StandardStreamingProcessor) extractURLsFromDoc(doc *DocJSONData, callback URLCallback) error {
	// Extract from references
	for _, ref := range doc.References {
		if ref.URL != "" && strings.HasSuffix(ref.URL, ".json") {
			if err := callback(ref.URL, "references"); err != nil {
				return err
			}
		}
	}
	
	// Extract from topic sections
	for _, section := range doc.TopicSections {
		for _, id := range section.Identifiers {
			if err := callback(id, "topicSections"); err != nil {
				return err
			}
		}
	}
	
	return nil
}

// FastJSONStreamingProcessor uses fastjson with streaming approach
type FastJSONStreamingProcessor struct {
	parser fastjson.Parser
}

func (p *FastJSONStreamingProcessor) Name() string {
	return "fastjson-streaming"
}

func (p *FastJSONStreamingProcessor) ProcessStream(reader io.Reader, callback URLCallback) error {
	// Read in chunks to avoid loading entire file into memory
	const chunkSize = 64 * 1024 // 64KB chunks
	
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, chunkSize), chunkSize*2)
	
	var jsonBuffer bytes.Buffer
	bracketLevel := 0
	inString := false
	escaped := false
	
	for scanner.Scan() {
		line := scanner.Bytes()
		
		for _, b := range line {
			jsonBuffer.WriteByte(b)
			
			if !escaped && b == '"' {
				inString = !inString
			}
			
			if !inString {
				switch b {
				case '{', '[':
					bracketLevel++
				case '}', ']':
					bracketLevel--
					
					// If we've closed all brackets, we have a complete JSON object
					if bracketLevel == 0 {
						// Parse the complete object
						if err := p.parseChunk(jsonBuffer.Bytes(), callback); err != nil {
							return err
						}
						jsonBuffer.Reset()
					}
				}
			}
			
			escaped = !escaped && b == '\\'
			if b != '\\' {
				escaped = false
			}
			
			// Prevent buffer from growing too large
			if jsonBuffer.Len() > 10*1024*1024 { // 10MB limit
				return fmt.Errorf("JSON object too large to process in streaming mode")
			}
		}
	}
	
	// Process any remaining data
	if jsonBuffer.Len() > 0 {
		return p.parseChunk(jsonBuffer.Bytes(), callback)
	}
	
	return scanner.Err()
}

func (p *FastJSONStreamingProcessor) parseChunk(data []byte, callback URLCallback) error {
	v, err := p.parser.ParseBytes(data)
	if err != nil {
		return fmt.Errorf("parse JSON chunk: %v", err)
	}
	
	return p.extractURLsFromFastJSON(v, callback)
}

func (p *FastJSONStreamingProcessor) extractURLsFromFastJSON(v *fastjson.Value, callback URLCallback) error {
	// Extract from references
	refs := v.Get("references")
	if refs != nil {
		refs.GetObject().Visit(func(key []byte, v *fastjson.Value) {
			if url := v.Get("url"); url != nil {
				if urlStr := string(url.GetStringBytes()); strings.HasSuffix(urlStr, ".json") {
					callback(urlStr, "references")
				}
			}
		})
	}
	
	// Extract from topic sections
	topics := v.Get("topicSections")
	if topics != nil {
		for _, topic := range topics.GetArray() {
			identifiers := topic.Get("identifiers")
			if identifiers != nil {
				for _, id := range identifiers.GetArray() {
					callback(string(id.GetStringBytes()), "topicSections")
				}
			}
		}
	}
	
	return nil
}

func (p *FastJSONStreamingProcessor) ProcessFile(filename string, callback URLCallback) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	return p.ProcessStream(file, callback)
}

// GJSONStreamingProcessor uses gjson with streaming approach
type GJSONStreamingProcessor struct{}

func (p *GJSONStreamingProcessor) Name() string {
	return "gjson-streaming"
}

func (p *GJSONStreamingProcessor) ProcessStream(reader io.Reader, callback URLCallback) error {
	// GJSON works best with complete JSON, so we read in larger chunks
	const maxChunkSize = 1024 * 1024 // 1MB chunks
	
	buffer := make([]byte, maxChunkSize)
	var jsonData strings.Builder
	
	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			jsonData.Write(buffer[:n])
		}
		
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read chunk: %v", err)
		}
		
		// Check if we have a complete JSON object
		data := jsonData.String()
		if gjson.Valid(data) {
			return p.extractURLsFromGJSON(data, callback)
		}
	}
	
	// Process the complete JSON
	data := jsonData.String()
	if !gjson.Valid(data) {
		return fmt.Errorf("invalid JSON data")
	}
	
	return p.extractURLsFromGJSON(data, callback)
}

func (p *GJSONStreamingProcessor) extractURLsFromGJSON(data string, callback URLCallback) error {
	// Extract from references using path-based queries
	result := gjson.Get(data, "references")
	result.ForEach(func(key, value gjson.Result) bool {
		if url := value.Get("url"); url.Exists() && strings.HasSuffix(url.String(), ".json") {
			callback(url.String(), "references")
		}
		return true
	})
	
	// Extract from topic sections
	gjson.Get(data, "topicSections").ForEach(func(key, value gjson.Result) bool {
		value.Get("identifiers").ForEach(func(key, value gjson.Result) bool {
			callback(value.String(), "topicSections")
			return true
		})
		return true
	})
	
	return nil
}

func (p *GJSONStreamingProcessor) ProcessFile(filename string, callback URLCallback) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	return p.ProcessStream(file, callback)
}

// JSONParserStreamingProcessor uses jsonparser for streaming
type JSONParserStreamingProcessor struct{}

func (p *JSONParserStreamingProcessor) Name() string {
	return "jsonparser-streaming"
}

func (p *JSONParserStreamingProcessor) ProcessStream(reader io.Reader, callback URLCallback) error {
	// Read the entire stream into memory for jsonparser
	// jsonparser requires the complete JSON data
	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("read stream: %v", err)
	}
	
	return p.extractURLsFromJSONParser(data, callback)
}

func (p *JSONParserStreamingProcessor) extractURLsFromJSONParser(data []byte, callback URLCallback) error {
	// Extract from references
	err := jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if url, err := jsonparser.GetString(value, "url"); err == nil && strings.HasSuffix(url, ".json") {
			return callback(url, "references")
		}
		return nil
	}, "references")
	
	if err != nil {
		return fmt.Errorf("extract from references: %v", err)
	}
	
	// Extract from topic sections
	_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			callback(string(value), "topicSections")
		}, "identifiers")
	}, "topicSections")
	
	if err != nil {
		return fmt.Errorf("extract from topic sections: %v", err)
	}
	
	return nil
}

func (p *JSONParserStreamingProcessor) ProcessFile(filename string, callback URLCallback) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	return p.ProcessStream(file, callback)
}

// ChunkedStreamingProcessor processes JSON in fixed-size chunks
type ChunkedStreamingProcessor struct {
	chunkSize int
	processor StreamingJSONProcessor
}

func NewChunkedStreamingProcessor(chunkSize int, processor StreamingJSONProcessor) *ChunkedStreamingProcessor {
	return &ChunkedStreamingProcessor{
		chunkSize: chunkSize,
		processor: processor,
	}
}

func (p *ChunkedStreamingProcessor) Name() string {
	return fmt.Sprintf("chunked-%s", p.processor.Name())
}

func (p *ChunkedStreamingProcessor) ProcessStream(reader io.Reader, callback URLCallback) error {
	buffer := make([]byte, p.chunkSize)
	var remainder []byte
	
	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			chunk := append(remainder, buffer[:n]...)
			
			// Find complete JSON objects in the chunk
			objects := p.findCompleteJSONObjects(chunk)
			
			for _, obj := range objects {
				objReader := bytes.NewReader(obj)
				if err := p.processor.ProcessStream(objReader, callback); err != nil {
					// Log error but continue processing
					fmt.Printf("Error processing chunk: %v\n", err)
				}
			}
			
			// Keep any incomplete data for the next iteration
			remainder = p.findIncompleteData(chunk)
		}
		
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read chunk: %v", err)
		}
	}
	
	// Process any remaining data
	if len(remainder) > 0 {
		objReader := bytes.NewReader(remainder)
		return p.processor.ProcessStream(objReader, callback)
	}
	
	return nil
}

// findCompleteJSONObjects finds complete JSON objects in a byte slice
func (p *ChunkedStreamingProcessor) findCompleteJSONObjects(data []byte) [][]byte {
	var objects [][]byte
	var start int
	bracketLevel := 0
	inString := false
	escaped := false
	
	for i, b := range data {
		if !escaped && b == '"' {
			inString = !inString
		}
		
		if !inString {
			switch b {
			case '{':
				if bracketLevel == 0 {
					start = i
				}
				bracketLevel++
			case '}':
				bracketLevel--
				if bracketLevel == 0 {
					objects = append(objects, data[start:i+1])
				}
			}
		}
		
		escaped = !escaped && b == '\\'
		if b != '\\' {
			escaped = false
		}
	}
	
	return objects
}

// findIncompleteData returns any incomplete JSON data at the end of a chunk
func (p *ChunkedStreamingProcessor) findIncompleteData(data []byte) []byte {
	bracketLevel := 0
	inString := false
	escaped := false
	lastCompleteEnd := -1
	
	for i, b := range data {
		if !escaped && b == '"' {
			inString = !inString
		}
		
		if !inString {
			switch b {
			case '{':
				bracketLevel++
			case '}':
				bracketLevel--
				if bracketLevel == 0 {
					lastCompleteEnd = i
				}
			}
		}
		
		escaped = !escaped && b == '\\'
		if b != '\\' {
			escaped = false
		}
	}
	
	if lastCompleteEnd >= 0 && lastCompleteEnd < len(data)-1 {
		return data[lastCompleteEnd+1:]
	}
	
	return nil
}

func (p *ChunkedStreamingProcessor) ProcessFile(filename string, callback URLCallback) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	return p.ProcessStream(file, callback)
}

// MemoryEfficientExtractor provides memory-efficient URL extraction
type MemoryEfficientExtractor struct {
	processors []StreamingJSONProcessor
}

func NewMemoryEfficientExtractor() *MemoryEfficientExtractor {
	return &MemoryEfficientExtractor{
		processors: []StreamingJSONProcessor{
			&StandardStreamingProcessor{},
			&FastJSONStreamingProcessor{},
			&GJSONStreamingProcessor{},
			&JSONParserStreamingProcessor{},
		},
	}
}

// ExtractURLsEfficiently extracts URLs using the most memory-efficient method
func (e *MemoryEfficientExtractor) ExtractURLsEfficiently(filename string) ([]string, error) {
	// Try processors in order of memory efficiency
	for _, processor := range e.processors {
		var urls []string
		
		callback := func(url, source string) error {
			urls = append(urls, url)
			return nil
		}
		
		if err := processor.ProcessFile(filename, callback); err != nil {
			continue
		}
		
		return urls, nil
	}
	
	return nil, fmt.Errorf("all processors failed")
}

// CompareStreamingPerformance compares different streaming approaches
func CompareStreamingPerformance(filename string) ([]StreamingResult, error) {
	extractors := []StreamingJSONProcessor{
		&StandardStreamingProcessor{},
		&FastJSONStreamingProcessor{},
		&GJSONStreamingProcessor{},
		&JSONParserStreamingProcessor{},
		NewChunkedStreamingProcessor(64*1024, &StandardStreamingProcessor{}),
	}
	
	var results []StreamingResult
	
	for _, extractor := range extractors {
		result := StreamingResult{
			ProcessorName: extractor.Name(),
		}
		
		var urls []string
		callback := func(url, source string) error {
			urls = append(urls, url)
			return nil
		}
		
		if err := extractor.ProcessFile(filename, callback); err != nil {
			result.Error = err
		} else {
			result.URLsFound = len(urls)
			
			// Get file size
			if info, err := os.Stat(filename); err == nil {
				result.BytesProcessed = info.Size()
			}
		}
		
		results = append(results, result)
	}
	
	return results, nil
}

// Example usage function
func demonstrateStreamingUsage() {
	fmt.Println("Streaming JSON Processing Examples")
	fmt.Println("==================================")
	
	// Example with different processors
	processors := []StreamingJSONProcessor{
		&StandardStreamingProcessor{},
		&FastJSONStreamingProcessor{},
		&GJSONStreamingProcessor{},
		&JSONParserStreamingProcessor{},
	}
	
	// Simulate processing a file
	testJSON := `{
		"metadata": {"title": "Test"},
		"references": {
			"ref1": {"url": "test1.json", "title": "Test 1"},
			"ref2": {"url": "test2.json", "title": "Test 2"}
		},
		"topicSections": [
			{"identifiers": ["ref1", "ref2"]}
		]
	}`
	
	reader := strings.NewReader(testJSON)
	
	for _, processor := range processors {
		fmt.Printf("\nTesting %s:\n", processor.Name())
		
		var urls []string
		callback := func(url, source string) error {
			urls = append(urls, fmt.Sprintf("%s (from %s)", url, source))
			return nil
		}
		
		// Reset reader for each processor
		reader.Seek(0, 0)
		
		if err := processor.ProcessStream(reader, callback); err != nil {
			fmt.Printf("  Error: %v\n", err)
		} else {
			fmt.Printf("  Found %d URLs:\n", len(urls))
			for _, url := range urls {
				fmt.Printf("    - %s\n", url)
			}
		}
	}
}