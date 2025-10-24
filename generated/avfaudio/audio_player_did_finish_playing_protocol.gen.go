// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// audioPlayerDidFinishPlayingProtocol is the audioPlayerDidFinishPlaying: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.2+
//   - iPadOS 2.2+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// Use this protocol when registering custom classes that conform to audioPlayerDidFinishPlaying:.
var audioPlayerDidFinishPlayingProtocol *objc.Protocol

func init() {
	audioPlayerDidFinishPlayingProtocol = objc.GetProtocol("audioPlayerDidFinishPlaying:")
}

