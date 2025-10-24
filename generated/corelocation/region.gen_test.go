// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation_test

import (
	"github.com/tmc/appledocs/generated/corelocation"
)

// Suppress unused import errors
var _ = corelocation.NewRegion

// ExampleNewRegionCircularRegionWithCenterRadiusIdentifier demonstrates how to create a Region instance using NewRegionCircularRegionWithCenterRadiusIdentifier.
// Initializes and returns a region object defining a circular area.
func ExampleNewRegionCircularRegionWithCenterRadiusIdentifier() {
	_ = corelocation.NewRegionCircularRegionWithCenterRadiusIdentifier(
		corelocation.LocationCoordinate2D{}, // center LocationCoordinate2D
		corelocation.LocationDistance{},     // radius LocationDistance
		"identifier",                        // identifier string
	)
	// Output:
}
