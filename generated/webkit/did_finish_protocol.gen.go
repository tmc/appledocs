// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// didFinishProtocol is the didFinish protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to didFinish.
var didFinishProtocol *objc.Protocol

func init() {
	didFinishProtocol = objc.GetProtocol("didFinish")
}
