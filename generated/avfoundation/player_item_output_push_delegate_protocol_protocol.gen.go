// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PPlayerItemOutputPushDelegate is the AVPlayerItemOutputPushDelegate protocol interface.
//
// A protocol that defines the methods to implement to respond to changes in the media data sequence.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerItemOutputPushDelegate
type PPlayerItemOutputPushDelegate interface {
	// Optional methods
	OutputSequenceWasFlushed(output IAVPlayerItemOutput)
	HasOutputSequenceWasFlushed() bool
}
