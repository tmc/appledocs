# appledocs

A Go tool that mirrors Apple documentation JSON files to disk with on-disk HTTP caching.

## Overview

This tool crawls the Apple documentation site starting from the technologies.json endpoint and saves all discovered JSON files to disk. It uses an on-disk HTTP cache to avoid unnecessary requests and supports concurrent downloads.

The tool recursively extracts URLs from JSON files by looking for fields that look like URLs and point to other JSON files. This allows it to discover and download the entire graph of JSON documentation files.

## Usage

```
# Build and run with default settings
make run

# Clean build artifacts and downloaded content
make clean

# Force refresh all content (ignores cache)
make force

# Run with higher concurrency for faster downloads
make fast
```

## Command-line options

```
./appledocs -h
  -base string
    	base URL for Apple docs (default "https://developer.apple.com")
  -cache string
    	directory to store HTTP cache (default ".cache")
  -concurrency int
    	number of concurrent downloads (default 10)
  -force
    	force refresh all content
  -output string
    	directory to store mirrored content (default "output")
```

## Output

The mirrored JSON files are saved to the `output` directory, preserving the original URL paths. The HTTP cache is stored in the `.cache` directory.