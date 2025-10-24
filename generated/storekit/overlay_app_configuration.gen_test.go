// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit_test

import (
	"github.com/tmc/appledocs/generated/storekit"
)

// Suppress unused import errors
var _ = storekit.NewOverlayAppConfiguration

// ExampleNewOverlayAppConfigurationWithAppIdentifierPosition demonstrates how to create a OverlayAppConfiguration instance using NewOverlayAppConfigurationWithAppIdentifierPosition.
// Creates an object that represents the attributes of an overlay you use to recommend another app on the App Store.
func ExampleNewOverlayAppConfigurationWithAppIdentifierPosition() {
	_ = storekit.NewOverlayAppConfigurationWithAppIdentifierPosition(
		"appIdentifier",            // appIdentifier string
		storekit.OverlayPosition{}, // position OverlayPosition
	)
	// Output:
}
