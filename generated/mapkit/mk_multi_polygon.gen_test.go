// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKMultiPolygon

// ExampleNewMKMultiPolygonWithPolygons demonstrates how to create a MKMultiPolygon instance using NewMKMultiPolygonWithPolygons.
// Creates a multipolygon object using the provided polygons.
func ExampleNewMKMultiPolygonWithPolygons() {
	_ = mapkit.NewMKMultiPolygonWithPolygons(
		[]mapkit.MKPolygon{}, // polygons []MKPolygon
	)
	// Output:
}
