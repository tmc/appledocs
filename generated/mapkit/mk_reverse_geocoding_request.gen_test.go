// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKReverseGeocodingRequest

// ExampleMKReverseGeocodingRequest_Cancel demonstrates using Cancel on a MKReverseGeocodingRequest instance.
// A method you call to cancel a reverse geocoding request that’s in progress.
func ExampleMKReverseGeocodingRequest_Cancel() {
	obj := mapkit.NewMKReverseGeocodingRequest()
	obj.Cancel()
	// Output:
	}

