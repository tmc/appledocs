// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// destinationForMixerProtocol is the destinationForMixer: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to destinationForMixer:.
var destinationForMixerProtocol *objc.Protocol

func init() {
	destinationForMixerProtocol = objc.GetProtocol("destinationForMixer:")
}

