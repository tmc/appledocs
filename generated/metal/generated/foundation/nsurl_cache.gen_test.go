// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLCache

// ExampleURLCache_RemoveAllCachedResponses demonstrates using RemoveAllCachedResponses on a URLCache instance.
// Clears the receiver’s cache, removing all stored cached URL responses.
func ExampleURLCache_RemoveAllCachedResponses() {
	obj := foundation.NewURLCache()
	obj.RemoveAllCachedResponses()
	// Output:
	}

