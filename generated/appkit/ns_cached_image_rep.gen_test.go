// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCachedImageRep

// ExampleNewCachedImageRepWithSizeDepthSeparateAlpha demonstrates how to create a CachedImageRep instance using NewCachedImageRepWithSizeDepthSeparateAlpha.
// Returns a cached image representation initialized with the specified image characteristics.
func ExampleNewCachedImageRepWithSizeDepthSeparateAlpha() {
	_ = appkit.NewCachedImageRepWithSizeDepthSeparateAlpha(
		appkit.Size /* not a class type */{}, // size Size /* not a class type */
		appkit.WindowDepth{}, // depth WindowDepth
		false, // flag bool
		false, // alpha bool
	)
	// Output:
}
