// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKGeodesicPolyline

// ExampleNewMKGeodesicPolylineWithCoordinatesCount demonstrates how to create a MKGeodesicPolyline instance using NewMKGeodesicPolylineWithCoordinatesCount.
// Creates and returns a geodesic polyline using the specified coordinates.
func ExampleNewMKGeodesicPolylineWithCoordinatesCount() {
	_ = mapkit.NewMKGeodesicPolylineWithCoordinatesCount(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coords LocationCoordinate2D /* not a class type */
		10, // count uint
	)
	// Output:
}
