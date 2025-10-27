// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PPlayerPlaybackCoordinatorDelegate is the AVPlayerPlaybackCoordinatorDelegate protocol interface.
//
// A protocol that defines the methods to implement to participate in playback coordination.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerPlaybackCoordinatorDelegate
type PPlayerPlaybackCoordinatorDelegate interface {
	// Optional methods
	PlaybackCoordinatorIdentifierForPlayerItem(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) foundation.String
	HasPlaybackCoordinatorIdentifierForPlayerItem() bool
	PlaybackCoordinatorInterstitialTimeRangesForPlayerItem(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) []foundation.Value
	HasPlaybackCoordinatorInterstitialTimeRangesForPlayerItem() bool
}
