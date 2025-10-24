// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKPlacemark

// ExampleNewMKPlacemarkWithCoordinate demonstrates how to create a MKPlacemark instance using NewMKPlacemarkWithCoordinate.
// Creates and returns a placemark object using the specified coordinate.
func ExampleNewMKPlacemarkWithCoordinate() {
	_ = mapkit.NewMKPlacemarkWithCoordinate(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coordinate LocationCoordinate2D /* not a class type */
	)
	// Output:
}
