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
		corelocation.CLLocationCoordinate2D /* not a class type */ {}, // center CLLocationCoordinate2D /* not a class type */
		corelocation.LocationDistance /* not a class type */ {},       // radius LocationDistance /* not a class type */
	)
	// Output:
}
