// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// isPreallocateInhibitedProtocol is the isPreallocateInhibited protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to isPreallocateInhibited.
var isPreallocateInhibitedProtocol *objc.Protocol

func init() {
	isPreallocateInhibitedProtocol = objc.GetProtocol("isPreallocateInhibited")
}
