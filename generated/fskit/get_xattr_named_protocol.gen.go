// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// getXattrNamedProtocol is the getXattrNamed: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to getXattrNamed:.
var getXattrNamedProtocol *objc.Protocol

func init() {
	getXattrNamedProtocol = objc.GetProtocol("getXattrNamed:")
}

