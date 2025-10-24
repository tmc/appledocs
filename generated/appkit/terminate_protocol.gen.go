// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// terminateProtocol is the terminate: protocol.
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Use this protocol when registering custom classes that conform to terminate:.
var terminateProtocol *objc.Protocol

func init() {
	terminateProtocol = objc.GetProtocol("terminate:")
}
