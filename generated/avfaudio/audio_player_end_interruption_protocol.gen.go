// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// audioPlayerEndInterruptionProtocol is the audioPlayerEndInterruption: protocol.
//
// Availability:
//   - tvOS +
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Use this protocol when registering custom classes that conform to audioPlayerEndInterruption:.
var audioPlayerEndInterruptionProtocol *objc.Protocol

func init() {
	audioPlayerEndInterruptionProtocol = objc.GetProtocol("audioPlayerEndInterruption:")
}

