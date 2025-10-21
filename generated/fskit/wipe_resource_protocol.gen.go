// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// wipeResourceProtocol is the wipeResource: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to wipeResource:.
var wipeResourceProtocol *objc.Protocol

func init() {
	wipeResourceProtocol = objc.GetProtocol("wipeResource:")
}
