// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import "github.com/ebitengine/purego/objc"

// requestNeedsUserApprovalProtocol is the requestNeedsUserApproval: protocol.
//
// Availability:
//   - macOS 10.15+
//
// Use this protocol when registering custom classes that conform to requestNeedsUserApproval:.
var requestNeedsUserApprovalProtocol *objc.Protocol

func init() {
	requestNeedsUserApprovalProtocol = objc.GetProtocol("requestNeedsUserApproval:")
}
