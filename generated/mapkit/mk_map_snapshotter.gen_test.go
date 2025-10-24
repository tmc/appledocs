// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKMapSnapshotter

// ExampleMKMapSnapshotter_Cancel demonstrates using Cancel on a MKMapSnapshotter instance.
// Cancels the request to create a snapshot.
func ExampleMKMapSnapshotter_Cancel() {
	obj := mapkit.NewMKMapSnapshotter()
	obj.Cancel()
	// Output:
	}

