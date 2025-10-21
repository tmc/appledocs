// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKAddress


// ExampleNewMKAddressWithFullAddressShortAddress demonstrates how to create a MKAddress instance using NewMKAddressWithFullAddressShortAddress.
// Initializes a new address with a location’s full address using a string and a short address that provides an abbreviated form of the address such as a street address.
func ExampleNewMKAddressWithFullAddressShortAddress() {
	_ = mapkit.NewMKAddressWithFullAddressShortAddress(
		"fullAddress", // fullAddress string
		"shortAddress", // shortAddress string
	)
	// Output:
}


