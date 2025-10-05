VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GOFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"
GO ?= go

.PHONY: build run clean test json fast force verbose all markdown html docs endpointsecurity compact accessibility graphql yaml yaml-all yaml-security appledocs appledocs-gql

build: appledocs

appledocs:
	$(GO) install $(GOFLAGS) ./cmd/appledocs

appledocs-gql:
	cd cmd/appledocs-gql && $(GO) install $(GOFLAGS) .

run: build
	appledocs -mode crawl

clean:
	rm -rf output .cache markdown yaml-output

test:
	$(GO) test ./...

fmt:
	gofmt -w .

# Force refresh all content when mirroring
force: build
	appledocs -mode crawl -force

# Mirror with higher concurrency
fast: build
	appledocs -mode crawl -concurrency 20

# Just mirror the json files (default behavior)
json: build
	appledocs -mode crawl

# Mirror with verbose output
verbose: build
	appledocs -mode crawl -concurrency 10 -verbose

# Generate HTML index only (requires existing JSON files)
html: build
	appledocs -mode html

# Generate Markdown documentation only (requires existing JSON files)
markdown: build
	appledocs -mode markdown

# Generate both HTML and Markdown (requires existing JSON files)
docs: build
	appledocs -mode html
	appledocs -mode markdown

# Generate specialized EndpointSecurity reference
endpointsecurity:
	go run es-md-test.go

# Skip symbol-level documentation (methods, properties) for smaller output
compact: build
	appledocs -mode crawl -skip-symbols

# Crawl only accessibility documentation
accessibility: build
	appledocs -mode crawl -entry-point "/tutorials/data/index/accessibility"

# Build and run the GraphQL server
graphql: appledocs-gql
	appledocs-gql

# Run the most comprehensive mirror
all: clean build
	appledocs -mode all -concurrency 20 -force
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

# Security-related targets
.PHONY: docker-secure docker-distroless docker-alpine security-scan security-check security-all

# Build secure Docker images
docker-secure: build
	docker build -f Dockerfile.secure -t appledocs:secure .

docker-distroless: build
	docker build -f Dockerfile.distroless -t appledocs:distroless .

docker-alpine: build
	docker build -f Dockerfile.alpine -t appledocs:alpine .

# Run security scans
security-scan: docker-secure
	./security/security-scan.sh appledocs:secure

# Quick security check
security-check:
	@echo "Running quick security checks..."
	@go mod verify
	@go vet ./...
	@if command -v gosec >/dev/null 2>&1; then \
		gosec -quiet -fmt text ./...; \
	else \
		echo "gosec not found. Install with: go install github.com/securego/gosec/v2/cmd/gosec@latest"; \
	fi
	@if command -v govulncheck >/dev/null 2>&1; then \
		govulncheck ./...; \
	else \
		echo "govulncheck not found. Install with: go install golang.org/x/vuln/cmd/govulncheck@latest"; \
	fi

# Run all security operations
security-all: docker-secure docker-distroless docker-alpine security-scan security-check
	@echo "All security checks completed. Check security-reports/ for detailed results."