VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GOFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"
GO ?= go

.PHONY: build run clean test json fast force verbose all markdown html docs endpointsecurity compact accessibility graphql yaml yaml-all yaml-security

build: appledocs appledocs-gql

appledocs:
	$(GO) build $(GOFLAGS) -o appledocs .

appledocs-gql:
	$(GO) build $(GOFLAGS) -o appledocs-gql ./cmd/appledocs-gql

run: build
	./appledocs -mode crawl

clean:
	rm -f appledocs appledocs-gql
	rm -rf output .cache markdown yaml-output

test:
	$(GO) test ./...

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

# Crawl only accessibility documentation
accessibility: build
	./appledocs -mode crawl -entry-point "/tutorials/data/index/accessibility"

# Build and run the GraphQL server
graphql:
	cd cmd/appledocs-gql && go build -o appledocs-gql
	cd cmd/appledocs-gql && ./appledocs-gql

# Run the most comprehensive mirror
all: clean build
	./appledocs -mode all -concurrency 20 -force
	go run es-md-test.go

# Convert specific JSON files from cache to YAML
yaml:
	@echo "Converting JSON files to YAML..."
	@mkdir -p yaml-output
	@find .cache -name "*.json" -type f -print0 | xargs -0 -I{} bash -c 'mkdir -p yaml-output/$$(dirname {}) && cat {} | yq -P > yaml-output/$${1%.json}.yaml' - {}
	@echo "YAML conversion complete. Files are in yaml-output directory."

# Convert all JSON files from cache to YAML with directory structure preserved
yaml-all:
	@echo "Converting all JSON files to YAML..."
	@mkdir -p yaml-output
	@find .cache -name "*.json" -type f -print0 | while read -d $$'\0' file; do \
		dir=$$(dirname "$$file"); \
		mkdir -p "yaml-output/$$dir"; \
		filename=$$(basename "$$file" .json); \
		yq -P "$$file" > "yaml-output/$$dir/$$filename.yaml"; \
	done
	@echo "YAML conversion complete. Files are in yaml-output directory with original structure preserved."

# Convert only SecurityFoundation and EndpointSecurity JSON files to YAML
yaml-security:
	@echo "Converting SecurityFoundation and EndpointSecurity JSON files to YAML..."
	@mkdir -p yaml-output
	@find .cache -path "*/SecurityFoundation*.json" -o -path "*/EndpointSecurity*.json" -type f -print0 | while read -d $$'\0' file; do \
		dir=$$(dirname "$$file"); \
		mkdir -p "yaml-output/$$dir"; \
		filename=$$(basename "$$file" .json); \
		echo "Converting $$file"; \
		yq -P "$$file" > "yaml-output/$$dir/$$filename.yaml"; \
	done
	@echo "YAML conversion complete. Security-related files are in yaml-output directory with original structure preserved."