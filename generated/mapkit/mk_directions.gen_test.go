// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKDirections

// ExampleMKDirections_Cancel demonstrates using Cancel on a MKDirections instance.
// Cancels a pending request.
func ExampleMKDirections_Cancel() {
	obj := mapkit.NewMKDirections()
	obj.Cancel()
	// Output:
	}

