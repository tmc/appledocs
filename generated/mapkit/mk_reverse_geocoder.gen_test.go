// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKReverseGeocoder

// ExampleNewMKReverseGeocoderWithCoordinate demonstrates how to create a MKReverseGeocoder instance using NewMKReverseGeocoderWithCoordinate.
// Initializes the reverse geocoder with the specified coordinate value.
func ExampleNewMKReverseGeocoderWithCoordinate() {
	_ = mapkit.NewMKReverseGeocoderWithCoordinate(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coordinate LocationCoordinate2D /* not a class type */
	)
	// Output:
}
