// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLComponents

// ExampleNewURLComponents demonstrates how to create a URLComponents instance.
// Creates a URL components object with all components left undefined.
func ExampleNewURLComponents() {
	_ = foundation.NewURLComponents()
	// Output:
}

// ExampleNewURLComponentsWithString demonstrates how to create a URLComponents instance using NewURLComponentsWithString.
// Creates a URL components object by parsing a URL in string form.
func ExampleNewURLComponentsWithString() {
	_ = foundation.NewURLComponentsWithString(
		"https://example.com", // URLString string
	)
	// Output:
}

// ExampleNewURLComponentsWithStringEncodingInvalidCharacters demonstrates how to create a URLComponents instance using NewURLComponentsWithStringEncodingInvalidCharacters.
// Creates a URL components instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
func ExampleNewURLComponentsWithStringEncodingInvalidCharacters() {
	_ = foundation.NewURLComponentsWithStringEncodingInvalidCharacters(
		"https://example.com", // URLString string
		false, // encodingInvalidCharacters bool
	)
	// Output:
}




