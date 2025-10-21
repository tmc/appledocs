// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKLocalSearchRequest

// ExampleNewMKLocalSearchRequest demonstrates how to create a MKLocalSearchRequest instance.
// Creates a local search request.
func ExampleNewMKLocalSearchRequest() {
	_ = mapkit.NewMKLocalSearchRequest()
	// Output:
}
// ExampleNewMKLocalSearchRequestWithNaturalLanguageQuery demonstrates how to create a MKLocalSearchRequest instance using NewMKLocalSearchRequestWithNaturalLanguageQuery.
// Initializes and returns a local search request based on the provided string.
func ExampleNewMKLocalSearchRequestWithNaturalLanguageQuery() {
	_ = mapkit.NewMKLocalSearchRequestWithNaturalLanguageQuery(
		"naturalLanguageQuery", // naturalLanguageQuery string
	)
	// Output:
}
