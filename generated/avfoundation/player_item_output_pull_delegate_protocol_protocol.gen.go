// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PPlayerItemOutputPullDelegate is the AVPlayerItemOutputPullDelegate protocol interface.
//
// Methods you can implement to respond to pixel buffer changes.
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
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerItemOutputPullDelegate
type PPlayerItemOutputPullDelegate interface {
	// Optional methods
	OutputMediaDataWillChange(sender IAVPlayerItemOutput)
	HasOutputMediaDataWillChange() bool
	OutputSequenceWasFlushed(output IAVPlayerItemOutput)
	HasOutputSequenceWasFlushed() bool
}
