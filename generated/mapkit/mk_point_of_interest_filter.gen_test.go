// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKPointOfInterestFilter

// ExampleNewMKPointOfInterestFilterExcludingCategories demonstrates how to create a MKPointOfInterestFilter instance using NewMKPointOfInterestFilterExcludingCategories.
// Initialize the point of interest filter with a list of categories to exclude.
func ExampleNewMKPointOfInterestFilterExcludingCategories() {
	_ = mapkit.NewMKPointOfInterestFilterExcludingCategories(
		[]mapkit.string{}, // categories []string
	)
	// Output:
}
// ExampleNewMKPointOfInterestFilterIncludingCategories demonstrates how to create a MKPointOfInterestFilter instance using NewMKPointOfInterestFilterIncludingCategories.
// Initialize the point of interest filter with a list of categories to include.
func ExampleNewMKPointOfInterestFilterIncludingCategories() {
	_ = mapkit.NewMKPointOfInterestFilterIncludingCategories(
		[]mapkit.string{}, // categories []string
	)
	// Output:
}
