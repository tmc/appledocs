// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// readFromFileProtocol is the readFromFile: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to readFromFile:.
var readFromFileProtocol *objc.Protocol

func init() {
	readFromFileProtocol = objc.GetProtocol("readFromFile:")
}
