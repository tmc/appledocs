// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewSpatialVideoConfiguration

// ExampleNewSpatialVideoConfiguration demonstrates how to create a SpatialVideoConfiguration instance.
func ExampleNewSpatialVideoConfiguration() {
	_ = avfoundation.NewSpatialVideoConfiguration()
	// Output:
}
// ExampleNewSpatialVideoConfigurationWithFormatDescription demonstrates how to create a SpatialVideoConfiguration instance using NewSpatialVideoConfigurationWithFormatDescription.
// Initializes an AVSpatialVideoConfiguration with a format description.
func ExampleNewSpatialVideoConfigurationWithFormatDescription() {
	_ = avfoundation.NewSpatialVideoConfigurationWithFormatDescription(
		avfoundation.FormatDescriptionRef /* not a class type */{}, // formatDescription FormatDescriptionRef /* not a class type */
	)
	// Output:
}
