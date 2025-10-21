// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewBundle

// ExampleNewBundleWithURL demonstrates how to create a Bundle instance using NewBundleWithURL.
// Returns an   object initialized to correspond to the specified file URL.
func ExampleNewBundleWithURL() {
	_ = foundation.NewBundleWithURL(
		foundation.URL{}, // url URL
	)
	// Output:
}
