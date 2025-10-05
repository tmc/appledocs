# AppLeDocs Utility Tools

This directory contains standalone analysis and benchmarking tools for the appledocs project. These files are excluded from the main build using `//go:build ignore` tags.

## Available Tools

### Benchmark Tools

#### `benchmark_standalone.go`
Comprehensive standalone benchmark for JSON parsing performance.

**Run:**
```bash
go run ../benchmark_standalone.go
```

**Features:**
- System information reporting
- Performance comparison across multiple JSON parsers
- Memory analysis
- Streaming analysis
- Final recommendations

#### `run_analysis.go`
Standalone runner for JSON parsing analysis suite.

**Run:**
```bash
go run ../run_analysis.go
```

**Features:**
- Basic performance tests
- Memory profiling
- Parser selection recommendations
- Detailed analysis reports

#### `benchmark_runner.go`
Orchestrates all benchmark and analysis operations.

**Run:**
```bash
go run ../benchmark_runner.go [options]
```

**Options:**
- `-html` - Generate HTML output
- `-markdown` - Generate markdown output
- `-memory` - Run memory profiling
- `-streaming` - Run streaming tests
- `-guide` - Generate migration guide
- `-verbose` - Enable verbose logging

### Analysis Tools

#### `json_memory_analysis.go`
Comprehensive memory analysis for JSON parsing in appledocs.

**Run:**
```bash
go run ../json_memory_analysis.go
```

**Features:**
- Memory profiling across different parsers
- Allocation tracking
- Memory usage comparisons
- Optimization recommendations

#### `enhanced_benchmarks.go`
Enhanced benchmarking for JSON parsing libraries with comprehensive analysis.

**Run:**
```bash
go run ../enhanced_benchmarks.go
```

**Features:**
- Comparative parser analysis
- Real-world document testing
- Performance metrics
- Detailed reports

#### `migration_guide.go`
Migration strategies and implementation guides for JSON parsing optimization.

**Run:**
```bash
go run ../migration_guide.go
```

**Features:**
- Migration strategies
- Code examples
- Before/after comparisons
- Implementation best practices

#### `streaming_examples.go`
Streaming JSON parsing implementations optimized for large Apple documentation files.

**Run:**
```bash
go run ../streaming_examples.go
```

**Features:**
- Streaming parser demonstrations
- Large file handling
- Memory-efficient parsing
- Performance comparisons

## Dependencies

These tools use additional JSON parsing libraries not required by the main binary:

- `github.com/buger/jsonparser` - Fast JSON parser
- `github.com/json-iterator/go` - High-performance JSON library
- `github.com/tidwall/gjson` - Get JSON values quickly
- `github.com/valyala/fastjson` - Fast JSON parser and validator

To install dependencies:
```bash
go get github.com/buger/jsonparser
go get github.com/json-iterator/go
go get github.com/tidwall/gjson
go get github.com/valyala/fastjson
```

Or use:
```bash
go mod download
```

## Running All Benchmarks

To run a comprehensive benchmark suite:

```bash
cd tools/
go run ../benchmark_standalone.go > benchmark_report.txt
```

## Integration with Main Project

These tools are intentionally separated from the main build to:
1. Keep the main binary small
2. Avoid unnecessary dependencies in production
3. Provide specialized analysis capabilities
4. Allow experimental features

## Notes

- All tools are standalone and can be run independently
- Some tools may take several minutes to complete
- Output is typically written to stdout (redirect to save)
- Some tools create temporary test data in memory

## Related

See `ROADMAP.md` for the project roadmap and `README.md` for the main project documentation.
