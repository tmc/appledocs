// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// unloadResourceProtocol is the unloadResource: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to unloadResource:.
var unloadResourceProtocol *objc.Protocol

func init() {
	unloadResourceProtocol = objc.GetProtocol("unloadResource:")
}
