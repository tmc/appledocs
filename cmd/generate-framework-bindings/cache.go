package main

import (
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
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
	// Hash all JSON files in the framework directory for proper cache invalidation
	// This ensures cache is invalidated when any documentation file changes
	frameworkPath := filepath.Join(inputDir, framework)

	// Check if framework directory exists
	if _, err := os.Stat(frameworkPath); err != nil {
		return "", err
	}

	// Collect all JSON files with their mtimes
	type fileInfo struct {
		path  string
		mtime time.Time
		size  int64
	}
	var files []fileInfo

	err := filepath.WalkDir(frameworkPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".json" {
			info, err := d.Info()
			if err != nil {
				return err
			}
			files = append(files, fileInfo{
				path:  path,
				mtime: info.ModTime(),
				size:  info.Size(),
			})
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	// Sort files by path for consistent hashing
	sort.Slice(files, func(i, j int) bool {
		return files[i].path < files[j].path
	})

	// Hash the sorted file information
	h := sha256.New()
	h.Write([]byte(framework))
	h.Write([]byte(cacheVersion))
	for _, f := range files {
		// Include relative path, mtime, and size in hash
		relPath, _ := filepath.Rel(frameworkPath, f.path)
		h.Write([]byte(relPath))
		h.Write([]byte(f.mtime.Format(time.RFC3339Nano)))
		h.Write([]byte(fmt.Sprintf("%d", f.size)))
	}

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
