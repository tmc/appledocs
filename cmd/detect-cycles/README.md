# detect-cycles

Detects circular import dependencies in generated Go packages.

## Two Implementations

### main.go (File-based)
Parses Go source files directly using `go/parser`.

**Pros:**
- Fast
- No compilation needed
- Works even with build errors

**Cons:**
- Doesn't respect build tags
- May miss some edge cases

```bash
go build -o detect-cycles main.go
./detect-cycles -v -dir ../../generated
```

### main_golist.go (go list-based)
Uses `go list -json` to get accurate package information.

**Pros:**
- More accurate
- Respects build tags and go.mod
- Uses Go's own dependency resolution

**Cons:**
- Requires packages to be valid Go modules
- Slower (runs go list)
- Fails if there are build errors

```bash
go build -o detect-cycles-golist main_golist.go
./detect-cycles-golist -v -dir ../../generated
```

## Usage

### Basic usage
```bash
# Detect cycles
./detect-cycles -dir ../../generated

# Verbose output
./detect-cycles -v -dir ../../generated

# Save to JSON
./detect-cycles -output cycles.json -dir ../../generated
```

### Exit codes
- `0` - No cycles found
- `1` - Cycles found (or error)

### Output format
```json
{
  "cycles_count": 2,
  "cycles": [
    {
      "cycle": ["vision", "coreimage", "vision"],
      "edges": [
        {
          "from": "vision",
          "to": "coreimage",
          "files": ["vision/vn_barcode_observation.gen.go"]
        },
        {
          "from": "coreimage",
          "to": "vision",
          "files": ["coreimage/ci_image.gen.go"]
        }
      ]
    }
  ]
}
```

## Integration with CI

### GitHub Actions
```yaml
- name: Check for circular imports
  run: |
    cd cmd/detect-cycles
    go build
    ./detect-cycles -dir ../../generated
```

### Pre-commit hook
```bash
#!/bin/bash
cd cmd/detect-cycles
go build
if ! ./detect-cycles -dir ../../generated; then
  echo "❌ Circular dependencies detected!"
  exit 1
fi
```

## Comparison with go list

You can also use `go list` directly to find cycles:

```bash
# Try to build all packages and look for cycle errors
cd generated
go list ./... 2>&1 | grep "import cycle"

# Get import information for specific packages
go list -f '{{.ImportPath}}: {{join .Imports ", "}}' ./vision ./coreimage

# JSON output for programmatic analysis
go list -json ./... > packages.json
```

## Advantages of this tool over go list

1. **Works with invalid packages** - Can detect cycles even when packages have other build errors
2. **Detailed cycle information** - Shows exact files and locations
3. **Normalized cycle output** - Deduplicates equivalent cycles
4. **Fast** - No compilation required (file-based version)
5. **Structured output** - JSON format for automation

## See also

- `cycle-breaking-rules.yaml` - Configuration for breaking rules
- `funcs_imports.go` - Import detection logic in generator
