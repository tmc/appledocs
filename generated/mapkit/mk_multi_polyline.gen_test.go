// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKMultiPolyline

// ExampleNewMKMultiPolylineWithPolylines demonstrates how to create a MKMultiPolyline instance using NewMKMultiPolylineWithPolylines.
// Creates a multipolyline object using the provided polylines.
func ExampleNewMKMultiPolylineWithPolylines() {
	_ = mapkit.NewMKMultiPolylineWithPolylines(
		[]mapkit.MKPolyline{}, // polylines []MKPolyline
	)
	// Output:
}
