# Appledocs Benchmarking Framework

A comprehensive performance benchmarking framework designed specifically for measuring and validating the appledocs refactoring phases. This framework provides detailed performance metrics, statistical analysis, and regression detection to ensure each refactoring phase improves performance without introducing regressions.

## Features

### 1. Pre-Refactoring Baseline Establishment
- **Comprehensive performance baselines** for the current codebase
- **Memory usage profiling** across different file sizes and concurrency levels
- **JSON parsing performance** measurements
- **Network I/O and cache efficiency** metrics
- **End-to-end crawling performance** benchmarks

### 2. Refactoring Phase Validation
- **Before/after comparison** framework for each refactoring phase
- **Regression detection** with configurable thresholds
- **Performance impact isolation** to identify which changes caused improvements or degradations
- **Memory leak detection** during extended runs
- **Statistical significance testing** using t-tests and effect size analysis

### 3. Continuous Performance Monitoring
- **CI/CD integration** for automated performance regression testing
- **Real-time performance alerts** during development
- **Historical trend analysis** and performance trajectory modeling
- **A/B testing framework** for comparing different implementation approaches

### 4. Implementation Features
- **Go benchmarking best practices** integration
- **Statistical significance testing** for performance comparisons
- **Multi-platform benchmarking** support
- **Automated report generation** with visualizations
- **Memory profiling** with heap, stack, and GC analysis
- **CPU profiling** with hotspot identification

## Quick Start

### Installation

```bash
# Install the benchmark command
go install github.com/tmc/appledocs/cmd/benchmark@latest
```

### Basic Usage

```bash
# Run all benchmark phases
benchmark -phase all

# Run baseline benchmark only
benchmark -phase baseline

# Compare two phases
benchmark -compare baseline,json-optimize

# Run in CI mode (fewer iterations, critical scenarios only)
benchmark -ci

# Generate test data
benchmark -generate-data

# Run continuous monitoring
benchmark -continuous -interval 5m
```

### Generate Reports

```bash
# Generate HTML report
benchmark -report html

# Generate Markdown report
benchmark -report markdown

# Custom output directory
benchmark -output ./results/2024-01-15
```

## Benchmark Scenarios

The framework includes appledocs-specific scenarios:

### 1. JSON Parsing (Small Files)
- Files < 10KB
- Tests basic parsing performance
- URL extraction efficiency

### 2. JSON Parsing (Large Files)
- Files > 1MB (technologies.json, UIKit.json)
- Memory efficiency testing
- Streaming vs. full-load comparison

### 3. Concurrent Crawling
- Different concurrency levels (1, 5, 10, 20, 50)
- Network I/O performance
- Rate limiting effectiveness

### 4. Cache Operations
- Cache read/write performance
- Validation overhead
- Concurrent access patterns

### 5. Markdown Generation
- Transformation performance
- Memory usage during generation
- Parallel processing benefits

### 6. Memory Intensive Operations
- Large dataset handling
- Memory leak detection
- GC pressure analysis

### 7. Real-World Simulation
- Full crawl simulation
- Mixed operation patterns
- End-to-end performance

## Performance Thresholds

Each scenario has configurable thresholds:

```json
{
  "json_parsing_small": {
    "max_duration": "100ms",
    "max_memory": "10MB",
    "max_allocations": 1000,
    "max_gc_pause": "10ms",
    "regression_margin": 0.1
  }
}
```

## Refactoring Phases

### Phase 1: Baseline
Current implementation performance baseline

### Phase 2: JSON Optimization
- Optimized JSON parsing with streaming
- Reduced allocations
- Better memory efficiency

### Phase 3: Streaming
- Full streaming JSON processing
- Reduced memory footprint
- Incremental processing

### Phase 4: Concurrency
- Optimized worker pools
- Better synchronization
- Reduced contention

### Phase 5: Memory Optimization
- Object pooling
- Buffer reuse
- GC pressure reduction

## Integration with CI/CD

### GitHub Actions Example

```yaml
name: Performance Benchmarks

on:
  pull_request:
    paths:
      - '**.go'
      - 'go.mod'

jobs:
  benchmark:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run Benchmarks
        run: |
          go run cmd/benchmark/main.go -ci -fail-on-regression
      
      - name: Upload Results
        uses: actions/upload-artifact@v3
        with:
          name: benchmark-results
          path: benchmark_results/
```

### GitLab CI Example

```yaml
benchmark:
  stage: test
  script:
    - go run cmd/benchmark/main.go -ci
  artifacts:
    reports:
      performance: benchmark_results/benchmark_report.json
    paths:
      - benchmark_results/
  only:
    - merge_requests
```

## Analyzing Results

### Understanding the Report

The benchmark report includes:

1. **Phase Results**: Performance metrics for each phase
2. **Comparisons**: Statistical comparison between phases
3. **Recommendations**: Actionable insights
4. **Critical Findings**: Performance regressions or issues

### Key Metrics

- **Duration**: Operation execution time
- **Memory**: Heap allocation and usage
- **Allocations**: Number of memory allocations
- **GC Pause**: Garbage collection pause times
- **Throughput**: Operations per second

### Statistical Analysis

- **T-Test**: Determines if performance changes are statistically significant
- **Effect Size**: Measures the magnitude of performance changes
- **P-Value**: Probability that changes are due to chance
- **Confidence Intervals**: Range of expected performance

## Advanced Usage

### Custom Configuration

Create a `benchmark.json` file:

```json
{
  "iterations": 10,
  "warmup_iterations": 3,
  "timeout": "5m",
  "cpu_profile": true,
  "mem_profile": true,
  "scenarios": [
    {
      "name": "custom_scenario",
      "operations": [
        {
          "name": "parse_json",
          "type": "parse",
          "function": "json.Unmarshal"
        }
      ]
    }
  ]
}
```

Run with custom config:

```bash
benchmark -config benchmark.json
```

### Memory Leak Detection

```bash
# Run extended benchmark with memory profiling
benchmark -phase baseline -iterations 100 -mem-profile

# Analyze memory growth
go tool pprof benchmark_results/*/profiles/*_mem_*.prof
```

### CPU Profiling

```bash
# Run with CPU profiling
benchmark -cpu-profile

# Analyze CPU hotspots
go tool pprof benchmark_results/*/profiles/*_cpu_*.prof
```

## Best Practices

1. **Warm-up Iterations**: Always include warm-up to stabilize performance
2. **Multiple Iterations**: Run at least 5 iterations for statistical significance
3. **Isolated Environment**: Run benchmarks on dedicated hardware when possible
4. **Consistent State**: Ensure consistent system state between runs
5. **Version Control**: Track benchmark results over time

## Troubleshooting

### High Variance in Results

- Increase iterations
- Check for background processes
- Ensure consistent system load
- Use dedicated benchmark hardware

### Memory Leaks

- Enable memory profiling
- Check for increasing heap usage
- Analyze allocation patterns
- Review goroutine counts

### Performance Regressions

- Compare with baseline
- Check statistical significance
- Review code changes
- Profile specific operations

## Contributing

When adding new benchmarks:

1. Create scenario in `appledocs_scenarios.go`
2. Add runner implementation in `appledocs_runners.go`
3. Define thresholds
4. Add integration tests
5. Update documentation

## Architecture

```
benchmark/
├── framework.go          # Core framework
├── runner.go            # Benchmark runners
├── profiler.go          # Profiling support
├── comparator.go        # Result comparison
├── validator.go         # Validation logic
├── reporter.go          # Report generation
├── storage.go           # Result storage
├── appledocs_scenarios.go # Appledocs scenarios
└── appledocs_runners.go   # Implementation runners
```

## Future Enhancements

1. **Cloud Integration**: Store results in cloud storage
2. **Dashboards**: Real-time performance dashboards
3. **ML Analysis**: Anomaly detection and trend prediction
4. **Distributed Benchmarking**: Run on multiple machines
5. **Custom Metrics**: Extensible metric collection

## License

Same as the appledocs project.