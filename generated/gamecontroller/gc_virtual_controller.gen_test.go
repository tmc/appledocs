// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller_test

import (
	"github.com/tmc/appledocs/generated/gamecontroller"
)

// Suppress unused import errors
var _ = gamecontroller.NewGCVirtualController

// ExampleNewGCVirtualControllerWithConfiguration demonstrates how to create a GCVirtualController instance using NewGCVirtualControllerWithConfiguration.
// Creates a new virtual controller using the configuration you specify.
func ExampleNewGCVirtualControllerWithConfiguration() {
	_ = gamecontroller.NewGCVirtualControllerWithConfiguration(
		gamecontroller.GCVirtualControllerConfiguration{}, // configuration GCVirtualControllerConfiguration
	)
	// Output:
}

