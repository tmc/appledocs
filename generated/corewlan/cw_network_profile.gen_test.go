// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan_test

import (
	"github.com/tmc/appledocs/generated/corewlan"
)

// Suppress unused import errors
var _ = corewlan.NewCWNetworkProfile

// ExampleNewCWNetworkProfile demonstrates how to create a CWNetworkProfile instance.
// Creates and returns a CWNetworkProfile object.
func ExampleNewCWNetworkProfile() {
	_ = corewlan.NewCWNetworkProfile()
	// Output:
}
// ExampleNewCWNetworkProfileWithNetworkProfile demonstrates how to create a CWNetworkProfile instance using NewCWNetworkProfileWithNetworkProfile.
// Creates and returns a CWNetworkProfile object initialized with the given CWNetworkProfile object.
func ExampleNewCWNetworkProfileWithNetworkProfile() {
	_ = corewlan.NewCWNetworkProfileWithNetworkProfile(
		corewlan.CWNetworkProfile{}, // networkProfile CWNetworkProfile
	)
	// Output:
}
