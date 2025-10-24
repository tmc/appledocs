// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction_test

import (
	"github.com/tmc/appledocs/generated/nearbyinteraction"
)

// Suppress unused import errors
var _ = nearbyinteraction.NewNIDLTDOAConfiguration

// ExampleNewNIDLTDOAConfigurationWithNetworkIdentifier demonstrates how to create a NIDLTDOAConfiguration instance using NewNIDLTDOAConfigurationWithNetworkIdentifier.
// Initializes a Downlink Time-Difference-of-Arrival (DL-TDoA) configuration for a specific tracked area.
func ExampleNewNIDLTDOAConfigurationWithNetworkIdentifier() {
	_ = nearbyinteraction.NewNIDLTDOAConfigurationWithNetworkIdentifier(
		1, // networkIdentifier int
	)
	// Output:
}
