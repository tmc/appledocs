// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import "github.com/ebitengine/purego/objc"

// StreamDelegateProtocol is the SCStreamDelegate protocol.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - macOS 12.3+
//
// Use this protocol when registering custom classes that conform to SCStreamDelegate.
var StreamDelegateProtocol *objc.Protocol

func init() {
	StreamDelegateProtocol = objc.GetProtocol("SCStreamDelegate")
}
