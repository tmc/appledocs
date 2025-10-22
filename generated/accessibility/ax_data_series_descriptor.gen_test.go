// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXDataSeriesDescriptor

// ExampleNewAXDataSeriesDescriptorWithNameIsContinuousDataPoints demonstrates how to create a AXDataSeriesDescriptor instance using NewAXDataSeriesDescriptorWithNameIsContinuousDataPoints.
// Creates a data series with the specified name, a Boolean value that indicates whether   the series is continuous, and data points.
func ExampleNewAXDataSeriesDescriptorWithNameIsContinuousDataPoints() {
	_ = accessibility.NewAXDataSeriesDescriptorWithNameIsContinuousDataPoints(
		"name", // name string
		false, // isContinuous bool
		[]accessibility.AXDataPoint{}, // dataPoints []AXDataPoint
	)
	// Output:
}
