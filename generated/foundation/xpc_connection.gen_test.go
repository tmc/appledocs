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
