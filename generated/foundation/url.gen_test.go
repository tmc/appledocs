// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURL






// ExampleNewURLFileURLWithPath demonstrates how to create a URL instance using NewURLFileURLWithPath.
// Initializes a newly created NSURL referencing the local file or directory at  .
func ExampleNewURLFileURLWithPath() {
	_ = foundation.NewURLFileURLWithPath(
		"/tmp/test", // path string
	)
	// Output:
}

// ExampleNewURLFileURLWithPathIsDirectory demonstrates how to create a URL instance using NewURLFileURLWithPathIsDirectory.
// Initializes a newly created NSURL referencing the local file or directory at  .
func ExampleNewURLFileURLWithPathIsDirectory() {
	_ = foundation.NewURLFileURLWithPathIsDirectory(
		"/tmp/test", // path string
		false, // isDir bool
	)
	// Output:
}





// ExampleNewURLWithSchemeHostPath demonstrates how to create a URL instance using NewURLWithSchemeHostPath.
// Initializes a newly created NSURL with a specified scheme, host, and path.
func ExampleNewURLWithSchemeHostPath() {
	_ = foundation.NewURLWithSchemeHostPath(
		"scheme", // scheme string
		"host", // host string
		"/tmp/test", // path string
	)
	// Output:
}

// ExampleNewURLWithString demonstrates how to create a URL instance using NewURLWithString.
// Initializes an NSURL object with a provided URL string.
func ExampleNewURLWithString() {
	_ = foundation.NewURLWithString(
		"https://example.com", // URLString string
	)
	// Output:
}

// ExampleNewURLWithStringEncodingInvalidCharacters demonstrates how to create a URL instance using NewURLWithStringEncodingInvalidCharacters.
// Creates an instance from the provided string, optionally IDNA- and percent-encoding any invalid characters.
func ExampleNewURLWithStringEncodingInvalidCharacters() {
	_ = foundation.NewURLWithStringEncodingInvalidCharacters(
		"https://example.com", // URLString string
		false, // encodingInvalidCharacters bool
	)
	// Output:
}



