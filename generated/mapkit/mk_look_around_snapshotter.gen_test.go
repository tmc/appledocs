// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKLookAroundSnapshotter

// ExampleMKLookAroundSnapshotter_Cancel demonstrates using Cancel on a MKLookAroundSnapshotter instance.
// Cancels an in-progress snapshot request.
func ExampleMKLookAroundSnapshotter_Cancel() {
	obj := mapkit.NewMKLookAroundSnapshotter()
	obj.Cancel()
	// Output:
	}

