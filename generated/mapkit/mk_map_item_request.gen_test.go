// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKMapItemRequest

// ExampleMKMapItemRequest_Cancel demonstrates using Cancel on a MKMapItemRequest instance.
// Cancels an in-progress map item request.
func ExampleMKMapItemRequest_Cancel() {
	obj := mapkit.NewMKMapItemRequest()
	obj.Cancel()
	// Output:
	}

