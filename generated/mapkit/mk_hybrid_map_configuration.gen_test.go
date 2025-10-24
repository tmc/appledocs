// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKHybridMapConfiguration

// ExampleNewMKHybridMapConfiguration demonstrates how to create a MKHybridMapConfiguration instance.
// Creates a new hybrid map configuration.
func ExampleNewMKHybridMapConfiguration() {
	_ = mapkit.NewMKHybridMapConfiguration()
	// Output:
}
// ExampleNewMKHybridMapConfigurationWithElevationStyle demonstrates how to create a MKHybridMapConfiguration instance using NewMKHybridMapConfigurationWithElevationStyle.
// Creates a new hybrid map configuration with the specified elevation style.
func ExampleNewMKHybridMapConfigurationWithElevationStyle() {
	_ = mapkit.NewMKHybridMapConfigurationWithElevationStyle(
		mapkit.MKMapElevationStyle{}, // elevationStyle MKMapElevationStyle
	)
	// Output:
}
