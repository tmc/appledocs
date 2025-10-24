// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor

import "github.com/ebitengine/purego/objc"

// OS_hv_gic_stateProtocol is the OS_hv_gic_state protocol.
//
// Availability:
//   - macOS +
//
// Use this protocol when registering custom classes that conform to OS_hv_gic_state.
var OS_hv_gic_stateProtocol *objc.Protocol

func init() {
	OS_hv_gic_stateProtocol = objc.GetProtocol("OS_hv_gic_state")
}

