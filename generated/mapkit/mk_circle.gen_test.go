// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKCircle

// ExampleNewMKCircleWithCenterCoordinateRadius demonstrates how to create a MKCircle instance using NewMKCircleWithCenterCoordinateRadius.
// Creates and returns a circle object using the specified coordinate and radius.
func ExampleNewMKCircleWithCenterCoordinateRadius() {
	_ = mapkit.NewMKCircleWithCenterCoordinateRadius(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coord LocationCoordinate2D /* not a class type */
		mapkit.LocationDistance /* not a class type */{}, // radius LocationDistance /* not a class type */
	)
	// Output:
}
