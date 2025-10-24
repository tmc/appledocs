// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import "github.com/ebitengine/purego/objc"

// streamDidBecomeActiveProtocol is the streamDidBecomeActive: protocol.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - macOS 15.2+
//
// Use this protocol when registering custom classes that conform to streamDidBecomeActive:.
var streamDidBecomeActiveProtocol *objc.Protocol

func init() {
	streamDidBecomeActiveProtocol = objc.GetProtocol("streamDidBecomeActive:")
}

