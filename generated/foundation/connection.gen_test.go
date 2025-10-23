// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewConnection

// ExampleNewConnectionWithReceivePortSendPort demonstrates how to create a Connection instance using NewConnectionWithReceivePortSendPort.
// Returns an   object initialized with given send and receive ports.
func ExampleNewConnectionWithReceivePortSendPort() {
	_ = foundation.NewConnectionWithReceivePortSendPort(
		foundation.NSPort{}, // receivePort NSPort
		foundation.NSPort{}, // sendPort NSPort
	)
	// Output:
}
