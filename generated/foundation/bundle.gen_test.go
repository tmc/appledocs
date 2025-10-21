// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewBundle

// ExampleNewBundleWithIdentifier demonstrates how to create a Bundle instance using NewBundleWithIdentifier.
// Returns the   instance that has the specified bundle identifier.
func ExampleNewBundleWithIdentifier() {
	_ = foundation.NewBundleWithIdentifier(
		"identifier", // identifier string
	)
	// Output:
}

// ExampleNewBundleWithPath demonstrates how to create a Bundle instance using NewBundleWithPath.
// Returns an   object initialized to correspond to the specified directory.
func ExampleNewBundleWithPath() {
	_ = foundation.NewBundleWithPath(
		"/tmp/test", // path string
	)
	// Output:
}




