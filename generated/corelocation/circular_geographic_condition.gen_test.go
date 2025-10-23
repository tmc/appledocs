// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation_test

import (
	"github.com/tmc/appledocs/generated/corelocation"
)

// Suppress unused import errors
var _ = corelocation.NewCircularGeographicCondition

// ExampleNewCircularGeographicConditionWithCenterRadius demonstrates how to create a CircularGeographicCondition instance using NewCircularGeographicConditionWithCenterRadius.
// Creates a new circular geographic condition with the center point and radius you provide.
func ExampleNewCircularGeographicConditionWithCenterRadius() {
	_ = corelocation.NewCircularGeographicConditionWithCenterRadius(
		corelocation.LocationCoordinate2D{}, // center LocationCoordinate2D
		corelocation.LocationDistance{}, // radius LocationDistance
	)
	// Output:
}
