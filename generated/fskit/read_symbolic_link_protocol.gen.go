// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// readSymbolicLinkProtocol is the readSymbolicLink: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to readSymbolicLink:.
var readSymbolicLinkProtocol *objc.Protocol

func init() {
	readSymbolicLinkProtocol = objc.GetProtocol("readSymbolicLink:")
}
