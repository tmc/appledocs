// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKLocalPointsOfInterestRequest

// ExampleNewMKLocalPointsOfInterestRequestWithCenterCoordinateRadius demonstrates how to create a MKLocalPointsOfInterestRequest instance using NewMKLocalPointsOfInterestRequestWithCenterCoordinateRadius.
// Creates a points of interest search request centered on the provided coordinate with the provided radius.
func ExampleNewMKLocalPointsOfInterestRequestWithCenterCoordinateRadius() {
	_ = mapkit.NewMKLocalPointsOfInterestRequestWithCenterCoordinateRadius(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coordinate LocationCoordinate2D /* not a class type */
		mapkit.LocationDistance /* not a class type */{}, // radius LocationDistance /* not a class type */
	)
	// Output:
}
