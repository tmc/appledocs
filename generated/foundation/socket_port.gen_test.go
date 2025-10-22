// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewSocketPort

// ExampleNewSocketPortWithProtocolFamilySocketTypeProtocolAddress demonstrates how to create a SocketPort instance using NewSocketPortWithProtocolFamilySocketTypeProtocolAddress.
// Initializes the receiver as a local socket with the provided arguments.
func ExampleNewSocketPortWithProtocolFamilySocketTypeProtocolAddress() {
	_ = foundation.NewSocketPortWithProtocolFamilySocketTypeProtocolAddress(
		0, // family int
		0, // type int
		0, // protocol int
		foundation.NSData{}, // address NSData
	)
	// Output:
}
