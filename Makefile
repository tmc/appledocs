.PHONY: build run clean test json fast force verbose all markdown html docs endpointsecurity compact

build:
	go build -o appledocs

run: build
	./appledocs -mode crawl

clean:
	rm -f appledocs
	rm -rf output .cache markdown

test:
	go test ./...

fmt:
	gofmt -w .

# Force refresh all content when mirroring
force: build
	./appledocs -mode crawl -force

# Mirror with higher concurrency
fast: build
	./appledocs -mode crawl -concurrency 20

# Just mirror the json files (default behavior)
json: build
	./appledocs -mode crawl

# Mirror with verbose output
verbose: build
	./appledocs -mode crawl -concurrency 10 -verbose

# Generate HTML index only (requires existing JSON files)
html: build
	./appledocs -mode html

# Generate Markdown documentation only (requires existing JSON files)
markdown: build
	./appledocs -mode markdown

# Generate both HTML and Markdown (requires existing JSON files)
docs: build
	./appledocs -mode html
	./appledocs -mode markdown

# Generate specialized EndpointSecurity reference
endpointsecurity:
	go run es-md-test.go

# Skip symbol-level documentation (methods, properties) for smaller output
compact: build
	./appledocs -mode crawl -skip-symbols

# Run the most comprehensive mirror
all: clean build
	./appledocs -mode all -concurrency 20 -force
	go run es-md-test.go