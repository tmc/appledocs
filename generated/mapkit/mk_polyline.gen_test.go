// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKPolyline

// ExampleNewMKPolylineWithCoordinatesCount demonstrates how to create a MKPolyline instance using NewMKPolylineWithCoordinatesCount.
// Creates a polyline object from the specified set of coordinates.
func ExampleNewMKPolylineWithCoordinatesCount() {
	_ = mapkit.NewMKPolylineWithCoordinatesCount(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coords LocationCoordinate2D /* not a class type */
		10, // count uint
	)
	// Output:
}
