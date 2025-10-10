.PHONY: all clean run test build-lib

SWIFT_LIB_DIR = ../../generated/swift/frameworks/coregraphics
SWIFT_LIB = libCoreGraphicsSwift.dylib

all: build-lib link-lib

build-lib:
	@echo "Building shared CoreGraphics Swift library..."
	cd $(SWIFT_LIB_DIR) && swift build -c release
	@echo "✓ Library built at $(SWIFT_LIB_DIR)/.build/release/$(SWIFT_LIB)"

link-lib: build-lib
	@echo "Linking library to local directory..."
	ln -sf $(SWIFT_LIB_DIR)/.build/release/$(SWIFT_LIB) ./$(SWIFT_LIB)

run: link-lib
	go run main.go

clean:
	rm -f $(SWIFT_LIB) output.png
	cd $(SWIFT_LIB_DIR) && swift package clean

test: run
	@echo "Checking if output.png was created..."
	@test -f output.png && echo "✓ PNG file created successfully" || (echo "✗ PNG file not created" && exit 1)
	@echo "Checking PNG file size..."
	@test -s output.png && echo "✓ PNG file has content" || (echo "✗ PNG file is empty" && exit 1)
