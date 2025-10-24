// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

// PSystemMusicPlayerController is the MPSystemMusicPlayerController protocol interface.
//
// A protocol for playing videos in the Music app.
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
// See: doc://com.apple.mediaplayer/documentation/MediaPlayer/MPSystemMusicPlayerController
type PSystemMusicPlayerController interface {
	// Required methods
	OpenToPlayQueueDescriptor(queueDescriptor IMPMusicPlayerQueueDescriptor)/* debug [protocol_interface/required_method]: OpenToPlayQueueDescriptor */
}
