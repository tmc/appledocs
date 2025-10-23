// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLSession

// ExampleNewURLSession demonstrates how to create a URLSession instance.
func ExampleNewURLSession() {
	_ = foundation.NewURLSession()
	// Output:
}
// ExampleNewURLSessionWithConfiguration demonstrates how to create a URLSession instance using NewURLSessionWithConfiguration.
// Creates a session with the specified session configuration.
func ExampleNewURLSessionWithConfiguration() {
	_ = foundation.NewURLSessionWithConfiguration(
		foundation.NSURLSessionConfiguration{}, // configuration NSURLSessionConfiguration
	)
	// Output:
}
