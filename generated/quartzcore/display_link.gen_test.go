// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewDisplayLink


// ExampleNewDisplayLinkWithTargetSelector demonstrates how to create a DisplayLink instance using NewDisplayLinkWithTargetSelector.
// Creates a display link for a target that calls its selector.
func ExampleNewDisplayLinkWithTargetSelector() {
	_ = quartzcore.NewDisplayLinkWithTargetSelector(
		0, // target objc.ID
		0, // sel objc.SEL
	)
	// Output:
}


