// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import "github.com/ebitengine/purego/objc"

// StreamOutputProtocol is the SCStreamOutput protocol.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - macOS 12.3+
//
// Use this protocol when registering custom classes that conform to SCStreamOutput.
var StreamOutputProtocol *objc.Protocol

func init() {
	StreamOutputProtocol = objc.GetProtocol("SCStreamOutput")
}
