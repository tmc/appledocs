// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXPCConnection

// ExampleNewXPCConnectionWithServiceName demonstrates how to create a XPCConnection instance using NewXPCConnectionWithServiceName.
// Initializes an   object to connect to an   object in an XPC service, identified by a service name.
func ExampleNewXPCConnectionWithServiceName() {
	_ = foundation.NewXPCConnectionWithServiceName(
		"serviceName", // serviceName string
	)
	// Output:
}

