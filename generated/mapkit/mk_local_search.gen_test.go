// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKLocalSearch

// ExampleMKLocalSearch_Cancel demonstrates using Cancel on a MKLocalSearch instance.
// Cancels an in-progress search operation.
func ExampleMKLocalSearch_Cancel() {
	obj := mapkit.NewMKLocalSearch()
	obj.Cancel()
	// Output:
	}

