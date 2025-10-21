// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import "github.com/ebitengine/purego/objc"

// callAudioUnitProtocol is the callAudioUnit: protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to callAudioUnit:.
var callAudioUnitProtocol *objc.Protocol

func init() {
	callAudioUnitProtocol = objc.GetProtocol("callAudioUnit:")
}
