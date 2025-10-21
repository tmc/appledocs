// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKCircleRenderer

// ExampleNewMKCircleRendererWithCircle demonstrates how to create a MKCircleRenderer instance using NewMKCircleRendererWithCircle.
// Creates a new overlay view using the specified circle overlay object.
func ExampleNewMKCircleRendererWithCircle() {
	_ = mapkit.NewMKCircleRendererWithCircle(
		mapkit.MKCircle{}, // circle MKCircle
	)
	// Output:
}
