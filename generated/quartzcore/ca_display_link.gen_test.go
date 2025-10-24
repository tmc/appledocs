// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewDisplayLink

// ExampleDisplayLink_Invalidate demonstrates using Invalidate on a DisplayLink instance.
// Removes the display link from all run loop modes.
func ExampleDisplayLink_Invalidate() {
	obj := quartzcore.NewDisplayLink()
	obj.Invalidate()
	// Output:
	}

