// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import "github.com/ebitengine/purego/objc"

// maximumFileSizeProtocol is the maximumFileSize protocol.
//
// Availability:
//   - macOS 15.4+
//
// Use this protocol when registering custom classes that conform to maximumFileSize.
var maximumFileSizeProtocol *objc.Protocol

func init() {
	maximumFileSizeProtocol = objc.GetProtocol("maximumFileSize")
}
