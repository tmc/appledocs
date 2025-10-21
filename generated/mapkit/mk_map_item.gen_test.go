// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKMapItem

// ExampleNewMKMapItemWithPlacemark demonstrates how to create a MKMapItem instance using NewMKMapItemWithPlacemark.
// Creates and returns a map item object using the specified placemark object.
func ExampleNewMKMapItemWithPlacemark() {
	_ = mapkit.NewMKMapItemWithPlacemark(
		mapkit.MKPlacemark{}, // placemark MKPlacemark
	)
	// Output:
}
