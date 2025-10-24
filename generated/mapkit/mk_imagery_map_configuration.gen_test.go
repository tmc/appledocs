// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKImageryMapConfiguration

// ExampleNewMKImageryMapConfiguration demonstrates how to create a MKImageryMapConfiguration instance.
// Creates a new imagery based map configuration.
func ExampleNewMKImageryMapConfiguration() {
	_ = mapkit.NewMKImageryMapConfiguration()
	// Output:
}
// ExampleNewMKImageryMapConfigurationWithElevationStyle demonstrates how to create a MKImageryMapConfiguration instance using NewMKImageryMapConfigurationWithElevationStyle.
// Creates a new imagery based map configuration with the specified elevation style.
func ExampleNewMKImageryMapConfigurationWithElevationStyle() {
	_ = mapkit.NewMKImageryMapConfigurationWithElevationStyle(
		mapkit.MKMapElevationStyle{}, // elevationStyle MKMapElevationStyle
	)
	// Output:
}
