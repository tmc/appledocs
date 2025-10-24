// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit_test

import (
	"github.com/tmc/appledocs/generated/metalkit"
)

// Suppress unused import errors
var _ = metalkit.NewView

// ExampleView_Draw demonstrates using Draw on a View instance.
// Redraws the view’s contents immediately.
func ExampleView_Draw() {
	obj := metalkit.NewView()
	obj.Draw()
	// Output:
	}

// ExampleView_ReleaseDrawables demonstrates using ReleaseDrawables on a View instance.
// Releases the   and   objects.
func ExampleView_ReleaseDrawables() {
	obj := metalkit.NewView()
	obj.ReleaseDrawables()
	// Output:
	}


