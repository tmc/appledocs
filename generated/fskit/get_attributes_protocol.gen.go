// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// getAttributesProtocol is the getAttributes: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to getAttributes:.
var getAttributesProtocol *objc.Protocol

func init() {
	getAttributesProtocol = objc.GetProtocol("getAttributes:")
}
