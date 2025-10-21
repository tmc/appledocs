// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit_test

import (
	"github.com/tmc/appledocs/generated/coreaudiokit"
)

// Suppress unused import errors
var _ = coreaudiokit.NewAudioUnitViewConfiguration


// ExampleNewAudioUnitViewConfigurationWithWidthHeightHostHasController demonstrates how to create a AudioUnitViewConfiguration instance using NewAudioUnitViewConfigurationWithWidthHeightHostHasController.
// Creates a new configuration object.
func ExampleNewAudioUnitViewConfigurationWithWidthHeightHostHasController() {
	_ = coreaudiokit.NewAudioUnitViewConfigurationWithWidthHeightHostHasController(
		0.0, // width float64
		0.0, // height float64
		false, // hostHasController bool
	)
	// Output:
}


