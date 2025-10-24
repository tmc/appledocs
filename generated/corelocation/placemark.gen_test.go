// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation_test

import (
	"github.com/tmc/appledocs/generated/corelocation"
)

// Suppress unused import errors
var _ = corelocation.NewPlacemark

// ExampleNewPlacemark demonstrates how to create a Placemark instance.
func ExampleNewPlacemark() {
	_ = corelocation.NewPlacemark()
	// Output:
}

// ExampleNewPlacemarkWithPlacemark demonstrates how to create a Placemark instance using NewPlacemarkWithPlacemark.
// Initializes and returns a placemark object from another placemark object.
func ExampleNewPlacemarkWithPlacemark() {
	_ = corelocation.NewPlacemarkWithPlacemark(
		corelocation.CLPlacemark{}, // placemark CLPlacemark
	)
	// Output:
}
