// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// AudioRecorderDelegateProtocol is the AVAudioRecorderDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.7+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// Use this protocol when registering custom classes that conform to AVAudioRecorderDelegate.
var AudioRecorderDelegateProtocol *objc.Protocol

func init() {
	AudioRecorderDelegateProtocol = objc.GetProtocol("AVAudioRecorderDelegate")
}


