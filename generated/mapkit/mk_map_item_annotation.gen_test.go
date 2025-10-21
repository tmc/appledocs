// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKMapItemAnnotation

// ExampleNewMKMapItemAnnotationWithMapItem demonstrates how to create a MKMapItemAnnotation instance using NewMKMapItemAnnotationWithMapItem.
// Creates a map item annotation
func ExampleNewMKMapItemAnnotationWithMapItem() {
	_ = mapkit.NewMKMapItemAnnotationWithMapItem(
		mapkit.MKMapItem{}, // mapItem MKMapItem
	)
	// Output:
}
