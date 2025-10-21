// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit_test

import (
	"github.com/tmc/appledocs/generated/storekit"
)

// Suppress unused import errors
var _ = storekit.NewOverlay

// ExampleNewOverlayWithConfiguration demonstrates how to create a Overlay instance using NewOverlayWithConfiguration.
// Creates an overlay you use to recommend another app on the App Store.
func ExampleNewOverlayWithConfiguration() {
	_ = storekit.NewOverlayWithConfiguration(
		storekit.SKOverlayConfiguration{}, // configuration SKOverlayConfiguration
	)
	// Output:
}
