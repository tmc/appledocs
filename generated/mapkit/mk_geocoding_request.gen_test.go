// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKGeocodingRequest

// ExampleMKGeocodingRequest_Cancel demonstrates using Cancel on a MKGeocodingRequest instance.
// A function you call to cancel a geocoding request that’s in progress.
func ExampleMKGeocodingRequest_Cancel() {
	obj := mapkit.NewMKGeocodingRequest()
	obj.Cancel()
	// Output:
	}

