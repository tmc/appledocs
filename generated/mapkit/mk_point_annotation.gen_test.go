// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKPointAnnotation

// ExampleNewMKPointAnnotation demonstrates how to create a MKPointAnnotation instance.
// Creates a map annotation that shows a title string at a point on a map.
func ExampleNewMKPointAnnotation() {
	_ = mapkit.NewMKPointAnnotation()
	// Output:
}
// ExampleNewMKPointAnnotationWithCoordinate demonstrates how to create a MKPointAnnotation instance using NewMKPointAnnotationWithCoordinate.
// Creates a point annotation at the specified coordinate on the map.
func ExampleNewMKPointAnnotationWithCoordinate() {
	_ = mapkit.NewMKPointAnnotationWithCoordinate(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coordinate LocationCoordinate2D /* not a class type */
	)
	// Output:
}
