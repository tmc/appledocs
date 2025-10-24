// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import "github.com/ebitengine/purego/objc"

// recordingOutputProtocol is the recordingOutput: protocol.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - macOS 15.0+
//
// Use this protocol when registering custom classes that conform to recordingOutput:.
var recordingOutputProtocol *objc.Protocol

func init() {
	recordingOutputProtocol = objc.GetProtocol("recordingOutput:")
}

