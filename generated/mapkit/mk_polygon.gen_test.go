// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKPolygon

// ExampleNewMKPolygonWithCoordinatesCount demonstrates how to create a MKPolygon instance using NewMKPolygonWithCoordinatesCount.
// Creates and returns a polygon object from the specified set of coordinates.
func ExampleNewMKPolygonWithCoordinatesCount() {
	_ = mapkit.NewMKPolygonWithCoordinatesCount(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coords LocationCoordinate2D /* not a class type */
		10, // count uint
	)
	// Output:
}
// ExampleNewMKPolygonWithCoordinatesCountInteriorPolygons demonstrates how to create a MKPolygon instance using NewMKPolygonWithCoordinatesCountInteriorPolygons.
// Creates and returns a polygon object from the specified set of coordinates and interior polygons.
func ExampleNewMKPolygonWithCoordinatesCountInteriorPolygons() {
	_ = mapkit.NewMKPolygonWithCoordinatesCountInteriorPolygons(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coords LocationCoordinate2D /* not a class type */
		10, // count uint
		[]mapkit.MKPolygon{}, // interiorPolygons []MKPolygon
	)
	// Output:
}
