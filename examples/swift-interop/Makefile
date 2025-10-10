.PHONY: all clean run test

all: libhello.dylib

libhello.dylib: hello.swift
	swiftc -emit-library hello.swift -o libhello.dylib

run: libhello.dylib
	go run main.go

clean:
	rm -f libhello.dylib

test: run
	@echo "Test passed!"
