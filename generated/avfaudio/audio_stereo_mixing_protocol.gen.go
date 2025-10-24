// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// AudioStereoMixingProtocol is the AVAudioStereoMixing protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to AVAudioStereoMixing.
var AudioStereoMixingProtocol *objc.Protocol

func init() {
	AudioStereoMixingProtocol = objc.GetProtocol("AVAudioStereoMixing")
}

