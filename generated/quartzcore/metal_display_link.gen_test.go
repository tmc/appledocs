// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewMetalDisplayLink

// ExampleNewMetalDisplayLinkWithMetalLayer demonstrates how to create a MetalDisplayLink instance using NewMetalDisplayLinkWithMetalLayer.
// Creates a display link for Metal from a Core Animation layer.
func ExampleNewMetalDisplayLinkWithMetalLayer() {
	_ = quartzcore.NewMetalDisplayLinkWithMetalLayer(
		quartzcore.CAMetalLayer{}, // layer CAMetalLayer
	)
	// Output:
}
