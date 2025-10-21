// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewMediaTimingFunction

// ExampleNewMediaTimingFunctionWithName demonstrates how to create a MediaTimingFunction instance using NewMediaTimingFunctionWithName.
// Creates and returns a new instance of   configured with the predefined timing function specified by  .
func ExampleNewMediaTimingFunctionWithName() {
	_ = quartzcore.NewMediaTimingFunctionWithName(
		quartzcore.MediaTimingFunctionName{}, // name MediaTimingFunctionName
	)
	// Output:
}
