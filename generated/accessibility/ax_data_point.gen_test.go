// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility_test

import (
	"github.com/tmc/appledocs/generated/accessibility"
)

// Suppress unused import errors
var _ = accessibility.NewAXDataPoint

// ExampleNewAXDataPointWithXY demonstrates how to create a AXDataPoint instance using NewAXDataPointWithXY.
// Creates a data point with the specified x- and y-values.
func ExampleNewAXDataPointWithXY() {
	_ = accessibility.NewAXDataPointWithXY(
		accessibility.AXDataPointValue{}, // xValue AXDataPointValue
		accessibility.AXDataPointValue{}, // yValue AXDataPointValue
	)
	// Output:
}
// ExampleNewAXDataPointWithXYAdditionalValues demonstrates how to create a AXDataPoint instance using NewAXDataPointWithXYAdditionalValues.
// Creates a data point with the specified x-value, y-value, and additional values.
func ExampleNewAXDataPointWithXYAdditionalValues() {
	_ = accessibility.NewAXDataPointWithXYAdditionalValues(
		accessibility.AXDataPointValue{}, // xValue AXDataPointValue
		accessibility.AXDataPointValue{}, // yValue AXDataPointValue
		[]accessibility.AXDataPointValue{}, // additionalValues []AXDataPointValue
	)
	// Output:
}
// ExampleNewAXDataPointWithXYAdditionalValuesLabel demonstrates how to create a AXDataPoint instance using NewAXDataPointWithXYAdditionalValuesLabel.
// Creates a data point with the specified x-value, y-value, additional values, and   label.
func ExampleNewAXDataPointWithXYAdditionalValuesLabel() {
	_ = accessibility.NewAXDataPointWithXYAdditionalValuesLabel(
		accessibility.AXDataPointValue{}, // xValue AXDataPointValue
		accessibility.AXDataPointValue{}, // yValue AXDataPointValue
		[]accessibility.AXDataPointValue{}, // additionalValues []AXDataPointValue
		"label", // label string
	)
	// Output:
}
