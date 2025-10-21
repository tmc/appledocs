// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import "github.com/ebitengine/purego/objc"

// callHostBlockProtocol is the callHostBlock protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to callHostBlock.
var callHostBlockProtocol *objc.Protocol

func init() {
	callHostBlockProtocol = objc.GetProtocol("callHostBlock")
}

