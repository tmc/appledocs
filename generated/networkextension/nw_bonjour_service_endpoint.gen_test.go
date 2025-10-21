// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNWBonjourServiceEndpoint


// ExampleNewNWBonjourServiceEndpointWithNameTypeDomain demonstrates how to create a NWBonjourServiceEndpoint instance using NewNWBonjourServiceEndpointWithNameTypeDomain.
// Create an endpoint with a Bonjour service name, type, and domain. All fields must be specified.
func ExampleNewNWBonjourServiceEndpointWithNameTypeDomain() {
	_ = networkextension.NewNWBonjourServiceEndpointWithNameTypeDomain(
		"name", // name string
		"type", // type string
		"domain", // domain string
	)
	// Output:
}


