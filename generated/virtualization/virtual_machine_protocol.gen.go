// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import "github.com/ebitengine/purego/objc"

// virtualMachineProtocol is the virtualMachine: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to virtualMachine:.
var virtualMachineProtocol *objc.Protocol

func init() {
	virtualMachineProtocol = objc.GetProtocol("virtualMachine:")
}
