// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PPlayerItemMetadataOutputPushDelegate is the AVPlayerItemMetadataOutputPushDelegate protocol interface.
//
// Methods you can implement to provide additional metadata.
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
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerItemMetadataOutputPushDelegate
type PPlayerItemMetadataOutputPushDelegate interface {
	// Optional methods
	MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(output IAVPlayerItemMetadataOutput, groups []TimedMetadataGroup, track IAVPlayerItemTrack)
	HasMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack() bool
}
