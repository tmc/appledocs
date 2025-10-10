.PHONY: all clean run

SWIFT_LIB = libCoreGraphicsUI.dylib

all: $(SWIFT_LIB)

$(SWIFT_LIB): AppKitUI.swift
	swiftc -emit-library AppKitUI.swift -o $(SWIFT_LIB) \
		-framework Cocoa \
		-framework CoreGraphics \
		-framework ImageIO

run: $(SWIFT_LIB)
	go run main.go

clean:
	rm -f $(SWIFT_LIB)

.PHONY: all clean run
