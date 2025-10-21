// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNWHostEndpoint


// ExampleNewNWHostEndpointWithHostnamePort demonstrates how to create a NWHostEndpoint instance using NewNWHostEndpointWithHostnamePort.
// Create a host endpoint with a hostname and port.
func ExampleNewNWHostEndpointWithHostnamePort() {
	_ = networkextension.NewNWHostEndpointWithHostnamePort(
		"hostname", // hostname string
		"port", // port string
	)
	// Output:
}


