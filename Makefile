.PHONY: build run clean test json fast force verbose all

build:
	go build -o appledocs

run: build
	./appledocs

clean:
	rm -f appledocs
	rm -rf output .cache

test:
	go test ./...

fmt:
	gofmt -w .

# Force refresh all content when mirroring
force: build
	./appledocs -force

# Mirror with higher concurrency
fast: build
	./appledocs -concurrency 20

# Just mirror the json files (default behavior)
json: build
	./appledocs

# Mirror with verbose output
verbose: build
	./appledocs -concurrency 10 -output output -cache .cache

# Run the most comprehensive mirror
all: clean build
	./appledocs -concurrency 20 -force