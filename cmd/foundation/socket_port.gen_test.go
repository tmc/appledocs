// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewSocketPort

// ExampleNewSocketPort demonstrates how to create a SocketPort instance.
// Initializes the receiver as a local TCP/IP socket of type  .
func ExampleNewSocketPort() {
	_ = foundation.NewSocketPort()
	// Output:
}
// ExampleNewSocketPortRemoteWithProtocolFamilySocketTypeProtocolAddress demonstrates how to create a SocketPort instance using NewSocketPortRemoteWithProtocolFamilySocketTypeProtocolAddress.
// Initializes the receiver as a remote socket with the provided arguments.
func ExampleNewSocketPortRemoteWithProtocolFamilySocketTypeProtocolAddress() {
	_ = foundation.NewSocketPortRemoteWithProtocolFamilySocketTypeProtocolAddress(
		0, // family int
		0, // type int
		0, // protocol int
		foundation.NSData{}, // address NSData
	)
	// Output:
}
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
// ExampleNewSocketPortWithProtocolFamilySocketTypeProtocolSocket demonstrates how to create a SocketPort instance using NewSocketPortWithProtocolFamilySocketTypeProtocolSocket.
// Initializes the receiver with a previously created local socket.
func ExampleNewSocketPortWithProtocolFamilySocketTypeProtocolSocket() {
	_ = foundation.NewSocketPortWithProtocolFamilySocketTypeProtocolSocket(
		0, // family int
		0, // type int
		0, // protocol int
		foundation.SocketNativeHandle{}, // sock SocketNativeHandle
	)
	// Output:
}
