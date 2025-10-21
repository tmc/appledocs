// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// createSymbolicLinkNamedProtocol is the createSymbolicLinkNamed: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to createSymbolicLinkNamed:.
var createSymbolicLinkNamedProtocol *objc.Protocol

func init() {
	createSymbolicLinkNamedProtocol = objc.GetProtocol("createSymbolicLinkNamed:")
}
