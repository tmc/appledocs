// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKStandardMapConfiguration

// ExampleNewMKStandardMapConfiguration demonstrates how to create a MKStandardMapConfiguration instance.
// Creates a new standard map configuration.
func ExampleNewMKStandardMapConfiguration() {
	_ = mapkit.NewMKStandardMapConfiguration()
	// Output:
}
// ExampleNewMKStandardMapConfigurationWithElevationStyle demonstrates how to create a MKStandardMapConfiguration instance using NewMKStandardMapConfigurationWithElevationStyle.
// Creates a new standard map configuration with the specified elevation style.
func ExampleNewMKStandardMapConfigurationWithElevationStyle() {
	_ = mapkit.NewMKStandardMapConfigurationWithElevationStyle(
		mapkit.MKMapElevationStyle{}, // elevationStyle MKMapElevationStyle
	)
	// Output:
}
// ExampleNewMKStandardMapConfigurationWithElevationStyleEmphasisStyle demonstrates how to create a MKStandardMapConfiguration instance using NewMKStandardMapConfigurationWithElevationStyleEmphasisStyle.
// Creates a standard map configuration with the specified elevation and emphasis styles.
func ExampleNewMKStandardMapConfigurationWithElevationStyleEmphasisStyle() {
	_ = mapkit.NewMKStandardMapConfigurationWithElevationStyleEmphasisStyle(
		mapkit.MKMapElevationStyle{}, // elevationStyle MKMapElevationStyle
		mapkit.MKStandardMapEmphasisStyle{}, // emphasisStyle MKStandardMapEmphasisStyle
	)
	// Output:
}
// ExampleNewMKStandardMapConfigurationWithEmphasisStyle demonstrates how to create a MKStandardMapConfiguration instance using NewMKStandardMapConfigurationWithEmphasisStyle.
// Creates a standard map configuration with the specified emphasis style.
func ExampleNewMKStandardMapConfigurationWithEmphasisStyle() {
	_ = mapkit.NewMKStandardMapConfigurationWithEmphasisStyle(
		mapkit.MKStandardMapEmphasisStyle{}, // emphasisStyle MKStandardMapEmphasisStyle
	)
	// Output:
}
