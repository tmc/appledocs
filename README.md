# appledocs

A Go tool that mirrors Apple documentation JSON files to disk with on-disk HTTP caching, with options to generate HTML and Markdown versions for offline reading.

## Overview

This tool crawls the Apple documentation site starting from the technologies.json endpoint (or a custom entry point) and saves all discovered JSON files to disk. It uses an on-disk HTTP cache to avoid unnecessary requests and supports concurrent downloads.

The tool recursively extracts URLs from JSON files by looking for fields that look like URLs and point to other JSON files. This allows it to discover and download the entire graph of JSON documentation files. You can specify an alternate entry point to crawl specific sections of the documentation.

Key features:
- Recursive crawling of Apple's documentation with customizable entry points
- Fast concurrent downloads with adjustable concurrency level
- On-disk HTTP caching to avoid redundant requests
- HTML index generation for easy browsing
- Markdown generation for readable documentation
- Operation modes for separate crawling and generating output formats

## Tools

This project includes two main tools:

1. **appledocs** - The main documentation crawler and converter
2. **appledocs-gql** - A GraphQL API server for the documentation

### appledocs

The main `appledocs` tool can be used to:

- Mirror Apple's documentation JSON files to local disk
- Convert the JSON to Markdown documentation
- Generate HTML documentation browser

### appledocs-gql

The `appledocs-gql` tool provides:

- A GraphQL API for querying the documentation
- REST endpoints for common operations
- Interactive GraphQL playground
- Apollo Sandbox for exploring the API

## Building

```bash
# Build both tools
make build

# Build only the main tool
make appledocs

# Build only the GraphQL server
make appledocs-gql
```

## Running the Tools

```bash
# Run the crawler to mirror documentation
make crawl
# or
./appledocs -mode=crawl

# Generate Markdown
make markdown
# or
./appledocs -mode=markdown

# Generate HTML
make html
# or
./appledocs -mode=html

# Generate all documentation formats
make all-docs
# or
./appledocs -mode=all

# Start the GraphQL server
make gql-server
# or
./appledocs-gql
```

## Usage

```
# Build and run with default settings (crawl mode)
make run

# Clean build artifacts and downloaded content
make clean

# Force refresh all content (ignores cache)
make force

# Run with higher concurrency for faster downloads
make fast

# Skip symbol-level documentation (methods, properties) for smaller output
./appledocs -mode crawl -skip-symbols

# Crawl a specific entry point (e.g., accessibility documentation)
./appledocs -mode crawl -entry-point "/tutorials/data/index/accessibility"

# Generate HTML index only (requires existing JSON files)
make html

# Generate Markdown documentation only (requires existing JSON files)
make markdown

# Generate both HTML and Markdown (requires existing JSON files)
make docs

# Generate specialized EndpointSecurity reference with collapsible sections
make endpointsecurity

# Do everything: crawl, generate HTML index, Markdown docs, and EndpointSecurity reference
make all
```

## Command-line options

```
./appledocs -h
  # Directories and URLs
  -output string
        directory to store mirrored content (default "output")
  -cache string
        directory to store HTTP cache (default ".cache")
  -base string
        base URL for Apple docs (default "https://developer.apple.com")
  -entry-point string
        path to start crawling from (default "/tutorials/data/documentation/technologies.json")
  
  # Operation mode
  -mode string
        operation mode: crawl, html, markdown, or all (default "crawl")
  
  # Crawling options
  -concurrency int
        number of concurrent downloads (default 10)
  -exclude-paths string
        comma-separated list of paths to exclude from crawling (default "en-US/docs/Mozilla")
  -force
        force refresh all content
  -timeout duration
        HTTP request timeout (default 30s)
  -max-time duration
        maximum time to run the program (default 1h)
  -skip-symbols
        skip individual symbol-level documentation (methods, properties)
  
  # Markdown-specific options
  -md-output string
        directory to store Markdown documentation (default "markdown")
```

## Output Formats

### JSON
The mirrored JSON files are saved to the `output` directory, preserving the original URL paths. The HTTP cache is stored in the `.cache` directory.

### HTML
An interactive HTML tree view is generated with the JSON files organized in a hierarchical structure. This makes it easy to browse the documentation. Open `output/index.html` in your browser to access it.

### Markdown
Human-readable documentation is generated in Markdown format, organized by framework and category. The Markdown files include properly formatted headings, code examples with syntax highlighting, and cross-references between related topics.

## Operation Modes

The tool can operate in four different modes:

1. **crawl**: (Default) Only downloads JSON files from Apple's site
2. **html**: Only generates the HTML index from existing JSON files
3. **markdown**: Only generates Markdown documentation from existing JSON files
4. **all**: Performs all operations: crawl, generate HTML index, and generate Markdown