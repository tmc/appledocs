// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewDisplayCriteria

// ExampleNewDisplayCriteriaWithRefreshRateFormatDescription demonstrates how to create a DisplayCriteria instance using NewDisplayCriteriaWithRefreshRateFormatDescription.
// Creates a display criteria object with the specified refresh rate and format description.
func ExampleNewDisplayCriteriaWithRefreshRateFormatDescription() {
	_ = avfoundation.NewDisplayCriteriaWithRefreshRateFormatDescription(
		0.0, // refreshRate float32
		avfoundation.FormatDescriptionRef /* not a class type */{}, // formatDescription FormatDescriptionRef /* not a class type */
	)
	// Output:
}
