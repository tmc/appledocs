// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation_test

import (
	"github.com/tmc/appledocs/generated/corelocation"
)

// Suppress unused import errors
var _ = corelocation.NewBeaconRegion

// ExampleNewBeaconRegionWithBeaconIdentityConstraintIdentifier demonstrates how to create a BeaconRegion instance using NewBeaconRegionWithBeaconIdentityConstraintIdentifier.
// Creates and returns a region object that targets beacons that satisfy the specified beacon identity constraints.
func ExampleNewBeaconRegionWithBeaconIdentityConstraintIdentifier() {
	_ = corelocation.NewBeaconRegionWithBeaconIdentityConstraintIdentifier(
		corelocation.CLBeaconIdentityConstraint{}, // beaconIdentityConstraint CLBeaconIdentityConstraint
		"identifier", // identifier string
	)
	// Output:
}
