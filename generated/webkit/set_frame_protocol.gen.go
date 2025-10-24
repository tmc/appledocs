// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// setFrameProtocol is the setFrame: protocol.
//
// Availability:
//   - Mac Catalyst 18.4+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 15.4+
//   - visionOS 2.4+
//
// Use this protocol when registering custom classes that conform to setFrame:.
var setFrameProtocol *objc.Protocol

func init() {
	setFrameProtocol = objc.GetProtocol("setFrame:")
}
