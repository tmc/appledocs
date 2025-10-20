// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import "github.com/ebitengine/purego/objc"

// speechSynthesizerProtocol is the speechSynthesizer: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.14+
//   - tvOS 7.0+
//   - visionOS 1.0+
//   - watchOS 1.0+
//
// Use this protocol when registering custom classes that conform to speechSynthesizer:.
var speechSynthesizerProtocol *objc.Protocol

func init() {
	speechSynthesizerProtocol = objc.GetProtocol("speechSynthesizer:")
}

