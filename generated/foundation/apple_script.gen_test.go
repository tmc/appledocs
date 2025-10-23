// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewAppleScript

// ExampleNewAppleScriptWithContentsOfURLError demonstrates how to create a AppleScript instance using NewAppleScriptWithContentsOfURLError.
// Initializes a newly allocated script instance from the source identified by the passed URL.
func ExampleNewAppleScriptWithContentsOfURLError() {
	_ = foundation.NewAppleScriptWithContentsOfURLError(
		foundation.URL{}, // url URL
		foundation.IDictionary{}, // errorInfo IDictionary
	)
	// Output:
}
// ExampleNewAppleScriptWithSource demonstrates how to create a AppleScript instance using NewAppleScriptWithSource.
// Initializes a newly allocated script instance from the passed source.
func ExampleNewAppleScriptWithSource() {
	_ = foundation.NewAppleScriptWithSource(
		"source", // source string
	)
	// Output:
}
