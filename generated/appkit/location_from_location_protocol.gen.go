// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// locationFromLocationProtocol is the locationFromLocation: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to locationFromLocation:.
var locationFromLocationProtocol *objc.Protocol

func init() {
	locationFromLocationProtocol = objc.GetProtocol("locationFromLocation:")
}
