// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// loadResourceProtocol is the loadResource: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to loadResource:.
var loadResourceProtocol *objc.Protocol

func init() {
	loadResourceProtocol = objc.GetProtocol("loadResource:")
}

