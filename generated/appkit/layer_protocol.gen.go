// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// layerProtocol is the layer: protocol.
//
// Availability:
//   - macOS 10.7+
//
// Use this protocol when registering custom classes that conform to layer:.
var layerProtocol *objc.Protocol

func init() {
	layerProtocol = objc.GetProtocol("layer:")
}
