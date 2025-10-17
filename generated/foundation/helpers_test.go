package foundation_test

import (
	"fmt"
	"testing"

	"github.com/tmc/appledocs/generated/foundation"
)

// TestStringWithGoString tests creating NSString from Go string.
func TestStringWithGoString(t *testing.T) {
	str := foundation.StringWithGoString("Hello, World!")
	if str.Length() == 0 {
		t.Error("String length should not be 0")
	}
}

// TestString_GoString tests converting NSString to Go string.
func TestString_GoString(t *testing.T) {
	nsStr := foundation.StringWithGoString("Test String")
	goStr := nsStr.GoString()

	if goStr != "Test String" {
		t.Errorf("Expected 'Test String', got '%s'", goStr)
	}
}

// ExampleStringWithGoString demonstrates creating an NSString from a Go string.
func ExampleStringWithGoString() {
	str := foundation.StringWithGoString("Hello, Foundation!")
	fmt.Printf("Created NSString with length: %d\n", str.Length())

	// Output:
	// Created NSString with length: 18
}

// ExampleString_GoString demonstrates converting an NSString to a Go string.
func ExampleString_GoString() {
	nsStr := foundation.StringWithGoString("Hello from NSString")
	goStr := nsStr.GoString()
	fmt.Println(goStr)

	// Output:
	// Hello from NSString
}

// ExampleString_Length demonstrates getting the length of an NSString.
func ExampleString_Length() {
	str := foundation.StringWithGoString("Test")
	length := str.Length()
	fmt.Printf("String length: %d\n", length)

	// Output:
	// String length: 4
}
