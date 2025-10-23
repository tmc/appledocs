// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXPCConnection

// ExampleNewXPCConnectionWithListenerEndpoint demonstrates how to create a XPCConnection instance using NewXPCConnectionWithListenerEndpoint.
// Initializes an   object to connect to an   object in another process, identified by an   object.
func ExampleNewXPCConnectionWithListenerEndpoint() {
	_ = foundation.NewXPCConnectionWithListenerEndpoint(
		foundation.NSXPCListenerEndpoint{}, // endpoint NSXPCListenerEndpoint
	)
	// Output:
}
// ExampleNewXPCConnectionWithMachServiceNameOptions demonstrates how to create a XPCConnection instance using NewXPCConnectionWithMachServiceNameOptions.
// Initializes an   object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a  .
func ExampleNewXPCConnectionWithMachServiceNameOptions() {
	_ = foundation.NewXPCConnectionWithMachServiceNameOptions(
		"name", // name string
		foundation.XPCConnectionOptions{}, // options XPCConnectionOptions
	)
	// Output:
}
// ExampleNewXPCConnectionWithServiceName demonstrates how to create a XPCConnection instance using NewXPCConnectionWithServiceName.
// Initializes an   object to connect to an   object in an XPC service, identified by a service name.
func ExampleNewXPCConnectionWithServiceName() {
	_ = foundation.NewXPCConnectionWithServiceName(
		"serviceName", // serviceName string
	)
	// Output:
}
