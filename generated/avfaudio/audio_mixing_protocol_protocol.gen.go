// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

// PAudioMixing is the AVAudioMixing protocol interface.
//
// A collection of properties that are applicable to the input bus of a mixer node.
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
// See: doc://com.apple.avfaudio/documentation/AVFAudio/AVAudioMixing
type PAudioMixing interface {
	// Required methods
	DestinationForMixerBus(mixer IAVAudioNode, bus AudioNodeBus /* typedef */) AudioMixingDestination/* debug [protocol_interface/required_method]: DestinationForMixerBus */
}
