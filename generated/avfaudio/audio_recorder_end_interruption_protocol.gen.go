// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// audioRecorderEndInterruptionProtocol is the audioRecorderEndInterruption: protocol.
//
// Availability:
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Use this protocol when registering custom classes that conform to audioRecorderEndInterruption:.
var audioRecorderEndInterruptionProtocol *objc.Protocol

func init() {
	audioRecorderEndInterruptionProtocol = objc.GetProtocol("audioRecorderEndInterruption:")
}

