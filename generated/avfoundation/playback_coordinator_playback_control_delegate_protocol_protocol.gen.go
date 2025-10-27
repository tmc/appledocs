// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
)

// PPlaybackCoordinatorPlaybackControlDelegate is the AVPlaybackCoordinatorPlaybackControlDelegate protocol interface.
//
// A protocol that defines the method to implement to respond to playback commands from the playback coordinator.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlaybackCoordinatorPlaybackControlDelegate
type PPlaybackCoordinatorPlaybackControlDelegate interface {
	// Required methods
	PlaybackCoordinatorDidIssueSeekCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, seekCommand IAVDelegatingPlaybackCoordinatorSeekCommand, completionHandler unsafe.Pointer)
	PlaybackCoordinatorDidIssuePauseCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, pauseCommand IAVDelegatingPlaybackCoordinatorPauseCommand, completionHandler unsafe.Pointer)
	PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, playCommand IAVDelegatingPlaybackCoordinatorPlayCommand, completionHandler unsafe.Pointer)
	PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, bufferingCommand IAVDelegatingPlaybackCoordinatorBufferingCommand, completionHandler unsafe.Pointer)
}
