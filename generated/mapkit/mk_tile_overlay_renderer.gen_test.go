// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKTileOverlayRenderer

// ExampleMKTileOverlayRenderer_ReloadData demonstrates using ReloadData on a MKTileOverlayRenderer instance.
// Forces the tile overlay renderer to reload and redisplay the tiles.
func ExampleMKTileOverlayRenderer_ReloadData() {
	obj := mapkit.NewMKTileOverlayRenderer()
	obj.ReloadData()
	// Output:
	}

