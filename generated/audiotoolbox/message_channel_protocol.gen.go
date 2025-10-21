// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import "github.com/ebitengine/purego/objc"

// MessageChannelProtocol is the AUMessageChannel protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to AUMessageChannel.
var MessageChannelProtocol *objc.Protocol

func init() {
	MessageChannelProtocol = objc.GetProtocol("AUMessageChannel")
}


