// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKAddressFilter

// ExampleNewMKAddressFilterExcludingOptions demonstrates how to create a MKAddressFilter instance using NewMKAddressFilterExcludingOptions.
// Creates an address filter with options for excluding results in a search.
func ExampleNewMKAddressFilterExcludingOptions() {
	_ = mapkit.NewMKAddressFilterExcludingOptions(
		mapkit.MKAddressFilterOption{}, // options MKAddressFilterOption
	)
	// Output:
}
// ExampleNewMKAddressFilterIncludingOptions demonstrates how to create a MKAddressFilter instance using NewMKAddressFilterIncludingOptions.
// Creates an address filter with options for including results in a search.
func ExampleNewMKAddressFilterIncludingOptions() {
	_ = mapkit.NewMKAddressFilterIncludingOptions(
		mapkit.MKAddressFilterOption{}, // options MKAddressFilterOption
	)
	// Output:
}
