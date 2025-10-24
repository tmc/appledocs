// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURL

// ExampleURL_RemoveAllCachedResourceValues demonstrates using RemoveAllCachedResourceValues on a URL instance.
// Removes all cached resource values and temporary resource values from the URL object.
func ExampleURL_RemoveAllCachedResourceValues() {
	obj := foundation.NewURL()
	obj.RemoveAllCachedResourceValues()
	// Output:
	}

