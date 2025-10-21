// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// lookupItemNamedProtocol is the lookupItemNamed: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to lookupItemNamed:.
var lookupItemNamedProtocol *objc.Protocol

func init() {
	lookupItemNamedProtocol = objc.GetProtocol("lookupItemNamed:")
}
