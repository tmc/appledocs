.PHONY: all clean run test

all: libcoregraphics_swift.dylib

libcoregraphics_swift.dylib: coregraphics_swift.swift
	swiftc -emit-library coregraphics_swift.swift -o libcoregraphics_swift.dylib -framework CoreGraphics -framework Foundation -framework ImageIO

run: libcoregraphics_swift.dylib
	go run main.go

clean:
	rm -f libcoregraphics_swift.dylib output.png

test: run
	@echo "Checking if output.png was created..."
	@test -f output.png && echo "✓ PNG file created successfully" || (echo "✗ PNG file not created" && exit 1)
	@echo "Checking PNG file size..."
	@test -s output.png && echo "✓ PNG file has content" || (echo "✗ PNG file is empty" && exit 1)
