// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKMapSnapshotter

// ExampleNewMKMapSnapshotterWithOptions demonstrates how to create a MKMapSnapshotter instance using NewMKMapSnapshotterWithOptions.
// Creates and returns a snapshotter object based on the specified options.
func ExampleNewMKMapSnapshotterWithOptions() {
	_ = mapkit.NewMKMapSnapshotterWithOptions(
		mapkit.MKMapSnapshotOptions{}, // options MKMapSnapshotOptions
	)
	// Output:
}
