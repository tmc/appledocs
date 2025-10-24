// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit_test

import (
	"github.com/tmc/appledocs/generated/storekit"
)

// Suppress unused import errors
var _ = storekit.NewOverlayAppClipConfiguration

// ExampleNewOverlayAppClipConfigurationWithPosition demonstrates how to create a OverlayAppClipConfiguration instance using NewOverlayAppClipConfigurationWithPosition.
// Creates an object that represents the attributes of an overlay you use to recommend an App Clip’s corresponding app.
func ExampleNewOverlayAppClipConfigurationWithPosition() {
	_ = storekit.NewOverlayAppClipConfigurationWithPosition(
		storekit.OverlayPosition{}, // position OverlayPosition
	)
	// Output:
}
