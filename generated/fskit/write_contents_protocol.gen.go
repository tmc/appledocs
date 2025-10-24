// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// writeContentsProtocol is the writeContents: protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to writeContents:.
var writeContentsProtocol *objc.Protocol

func init() {
	writeContentsProtocol = objc.GetProtocol("writeContents:")
}

