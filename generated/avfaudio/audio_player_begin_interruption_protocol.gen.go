// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// audioPlayerBeginInterruptionProtocol is the audioPlayerBeginInterruption: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.2+ (Deprecated in 8.0)
//   - iPadOS 2.2+ (Deprecated in 8.0)
//   - tvOS +
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Use this protocol when registering custom classes that conform to audioPlayerBeginInterruption:.
var audioPlayerBeginInterruptionProtocol *objc.Protocol

func init() {
	audioPlayerBeginInterruptionProtocol = objc.GetProtocol("audioPlayerBeginInterruption:")
}

