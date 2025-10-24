// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

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

// PlayerPlaybackCoordinatorDelegate is a delegate implementation builder for the PPlayerPlaybackCoordinatorDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayerPlaybackCoordinatorDelegate struct {
	_PlaybackCoordinatorIdentifierForPlayerItem func(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) foundation.String
	_PlaybackCoordinatorInterstitialTimeRangesForPlayerItem func(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) []foundation.Value
}

// SetPlaybackCoordinatorIdentifierForPlayerItem sets the handler for the PlaybackCoordinatorIdentifierForPlayerItem delegate method.
//
// Returns an identifier for a player item.
func (d *PlayerPlaybackCoordinatorDelegate) SetPlaybackCoordinatorIdentifierForPlayerItem(f func(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) foundation.String) {
	d._PlaybackCoordinatorIdentifierForPlayerItem = f
}

// SetPlaybackCoordinatorInterstitialTimeRangesForPlayerItem sets the handler for the PlaybackCoordinatorInterstitialTimeRangesForPlayerItem delegate method.
//
// Asks the delegate for time ranges in a player item that don’t correspond to the primary content.
func (d *PlayerPlaybackCoordinatorDelegate) SetPlaybackCoordinatorInterstitialTimeRangesForPlayerItem(f func(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) []foundation.Value) {
	d._PlaybackCoordinatorInterstitialTimeRangesForPlayerItem = f
}

// PlaybackCoordinatorIdentifierForPlayerItem implements the PPlayerPlaybackCoordinatorDelegate interface.
func (d *PlayerPlaybackCoordinatorDelegate) PlaybackCoordinatorIdentifierForPlayerItem(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) foundation.String {
	if d._PlaybackCoordinatorIdentifierForPlayerItem != nil {
		return d._PlaybackCoordinatorIdentifierForPlayerItem(coordinator, playerItem)
	}
	var zero foundation.String
	return zero
}

// HasPlaybackCoordinatorIdentifierForPlayerItem returns true if a handler for PlaybackCoordinatorIdentifierForPlayerItem has been set.
func (d *PlayerPlaybackCoordinatorDelegate) HasPlaybackCoordinatorIdentifierForPlayerItem() bool {
	return d._PlaybackCoordinatorIdentifierForPlayerItem != nil
}

// PlaybackCoordinatorInterstitialTimeRangesForPlayerItem implements the PPlayerPlaybackCoordinatorDelegate interface.
func (d *PlayerPlaybackCoordinatorDelegate) PlaybackCoordinatorInterstitialTimeRangesForPlayerItem(coordinator IAVPlayerPlaybackCoordinator, playerItem IAVPlayerItem) []foundation.Value {
	if d._PlaybackCoordinatorInterstitialTimeRangesForPlayerItem != nil {
		return d._PlaybackCoordinatorInterstitialTimeRangesForPlayerItem(coordinator, playerItem)
	}
	var zero []foundation.Value
	return zero
}

// HasPlaybackCoordinatorInterstitialTimeRangesForPlayerItem returns true if a handler for PlaybackCoordinatorInterstitialTimeRangesForPlayerItem has been set.
func (d *PlayerPlaybackCoordinatorDelegate) HasPlaybackCoordinatorInterstitialTimeRangesForPlayerItem() bool {
	return d._PlaybackCoordinatorInterstitialTimeRangesForPlayerItem != nil
}
