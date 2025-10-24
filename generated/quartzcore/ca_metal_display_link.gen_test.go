// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewMetalDisplayLink

// ExampleMetalDisplayLink_Invalidate demonstrates using Invalidate on a MetalDisplayLink instance.
// Removes the display link from all run loops for all modes.
func ExampleMetalDisplayLink_Invalidate() {
	obj := quartzcore.NewMetalDisplayLink()
	obj.Invalidate()
	// Output:
	}

