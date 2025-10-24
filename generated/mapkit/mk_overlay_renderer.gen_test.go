// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKOverlayRenderer

// ExampleMKOverlayRenderer_SetNeedsDisplay demonstrates using SetNeedsDisplay on a MKOverlayRenderer instance.
// Invalidates the entire contents of the overlay for all zoom scales.
func ExampleMKOverlayRenderer_SetNeedsDisplay() {
	obj := mapkit.NewMKOverlayRenderer()
	obj.SetNeedsDisplay()
	// Output:
	}

