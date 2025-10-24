// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import "github.com/ebitengine/purego/objc"

// guestDidStopVirtualMachineProtocol is the guestDidStopVirtualMachine: protocol.
//
// Availability:
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to guestDidStopVirtualMachine:.
var guestDidStopVirtualMachineProtocol *objc.Protocol

func init() {
	guestDidStopVirtualMachineProtocol = objc.GetProtocol("guestDidStopVirtualMachine:")
}

