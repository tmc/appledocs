package main

import (
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/tmc/appledocs/occ2go"
)

const cacheVersion = "v1" // Increment when cache format changes

type ParsedCache struct {
	Framework string
	Version   string
	Timestamp time.Time
	InputHash string
	Classes   []*occ2go.ParsedClass
	Protocols []*occ2go.ParsedProtocol
	Enums     []*occ2go.ParsedEnum
	Functions []*occ2go.ParsedFunction
	Typedefs  []*occ2go.ParsedTypedef
	Constants []*occ2go.ParsedConstant
}

func getCacheDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cacheDir := filepath.Join(home, ".appledocs", "cache", "parsed")
	return cacheDir, os.MkdirAll(cacheDir, 0755)
}

func getCachePath(framework string) (string, error) {
	dir, err := getCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, fmt.Sprintf("%s-%s.gob", framework, cacheVersion)), nil
}

func getInputHash(inputDir, framework string) (string, error) {
	// Fast cache validation using directory stats
	// This is much faster than walking all files while still catching most changes
	frameworkPath := filepath.Join(inputDir, framework)

	// Get directory info
	dirInfo, err := os.Stat(frameworkPath)
	if err != nil {
		return "", err
	}

	// Quick count of JSON files (not full walk)
	entries, err := os.ReadDir(frameworkPath)
	if err != nil {
		return "", err
	}

	fileCount := 0
	var totalSize int64
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			info, err := entry.Info()
			if err == nil {
				fileCount++
				totalSize += info.Size()
			}
		}
	}

	// Hash: directory mtime + file count + total size + cache version
	// This catches:
	// - Files added/removed (count changes)
	// - Files modified (size usually changes, or dir mtime updates)
	// - Directory structure changes (dir mtime updates)
	h := sha256.New()
	h.Write([]byte(framework))
	h.Write([]byte(cacheVersion))
	h.Write([]byte(dirInfo.ModTime().Format(time.RFC3339Nano)))
	h.Write([]byte(fmt.Sprintf("%d", fileCount)))
	h.Write([]byte(fmt.Sprintf("%d", totalSize)))

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func loadCache(framework, inputDir string, verbose bool) (*ParsedCache, error) {
	cachePath, err := getCachePath(framework)
	if err != nil {
		return nil, err
	}

	// Check if cache exists
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		if verbose {
			fmt.Fprintf(os.Stderr, "[%s] No cache found at %s\n", framework, cachePath)
		}
		return nil, nil
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "[%s] Found cache file at %s\n", framework, cachePath)
	}

	// Load cache
	f, err := os.Open(cachePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cache ParsedCache
	dec := gob.NewDecoder(f)
	if err := dec.Decode(&cache); err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "[%s] Cache decode error: %v\n", framework, err)
		}
		return nil, nil // Treat decode errors as cache miss
	}

	// Validate cache freshness
	inputHash, err := getInputHash(inputDir, framework)
	if err != nil {
		return nil, err
	}

	if cache.InputHash != inputHash {
		if verbose {
			fmt.Fprintf(os.Stderr, "[%s] Cache stale (input changed)\n", framework)
		}
		return nil, nil
	}

	if verbose {
		fmt.Fprintf(os.Stderr, "[%s] Loaded cache from %s (age: %v)\n",
			framework, cachePath, time.Since(cache.Timestamp).Round(time.Second))
	}

	return &cache, nil
}

func saveCache(cache *ParsedCache, framework string, verbose bool) error {
	cachePath, err := getCachePath(framework)
	if err != nil {
		return err
	}

	f, err := os.Create(cachePath)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(cache); err != nil {
		return err
	}

	if verbose {
		info, _ := f.Stat()
		fmt.Fprintf(os.Stderr, "[%s] Saved cache to %s (size: %d bytes)\n",
			framework, cachePath, info.Size())
	}

	return nil
}

func countMethods(classes []*occ2go.ParsedClass) int {
	count := 0
	for _, c := range classes {
		count += len(c.Methods)
	}
	return count
}
