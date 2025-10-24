// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// boundsProtocol is the bounds protocol.
//
// Availability:
//   - macOS 10.15+
//
// Use this protocol when registering custom classes that conform to bounds.
var boundsProtocol *objc.Protocol

func init() {
	boundsProtocol = objc.GetProtocol("bounds")
}
