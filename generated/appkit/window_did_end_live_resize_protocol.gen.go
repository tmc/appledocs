// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// windowDidEndLiveResizeProtocol is the windowDidEndLiveResize: protocol.
//
// Availability:
//   - macOS 10.6+
//
// Use this protocol when registering custom classes that conform to windowDidEndLiveResize:.
var windowDidEndLiveResizeProtocol *objc.Protocol

func init() {
	windowDidEndLiveResizeProtocol = objc.GetProtocol("windowDidEndLiveResize:")
}
