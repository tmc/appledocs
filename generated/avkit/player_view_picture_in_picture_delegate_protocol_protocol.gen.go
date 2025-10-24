// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PPlayerViewPictureInPictureDelegate is the AVPlayerViewPictureInPictureDelegate protocol interface.
//
// A protocol that defines the methods to implement to respond to Picture in Picture playback events.
//
// Availability:
//   - macOS 10.15+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVPlayerViewPictureInPictureDelegate
type PPlayerViewPictureInPictureDelegate interface {
	// Optional methods
	PlayerViewFailedToStartPictureInPictureWithError(playerView IAVPlayerView, error_ objc.IObject /* cross-framework: Error */)
	HasPlayerViewFailedToStartPictureInPictureWithError() bool
	PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerView IAVPlayerView, completionHandler unsafe.Pointer)
	HasPlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler() bool
	PlayerViewDidStartPictureInPicture(playerView IAVPlayerView)
	HasPlayerViewDidStartPictureInPicture() bool
	PlayerViewDidStopPictureInPicture(playerView IAVPlayerView)
	HasPlayerViewDidStopPictureInPicture() bool
	PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart(playerView IAVPlayerView) bool
	HasPlayerViewShouldAutomaticallyDismissAtPictureInPictureStart() bool
	PlayerViewWillStartPictureInPicture(playerView IAVPlayerView)
	HasPlayerViewWillStartPictureInPicture() bool
	PlayerViewWillStopPictureInPicture(playerView IAVPlayerView)
	HasPlayerViewWillStopPictureInPicture() bool
}

// PlayerViewPictureInPictureDelegate is a delegate implementation builder for the PPlayerViewPictureInPictureDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayerViewPictureInPictureDelegate struct {
	_PlayerViewFailedToStartPictureInPictureWithError func(playerView IAVPlayerView, error_ objc.IObject /* cross-framework: Error */)
	_PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler func(playerView IAVPlayerView, completionHandler unsafe.Pointer)
	_PlayerViewDidStartPictureInPicture func(playerView IAVPlayerView)
	_PlayerViewDidStopPictureInPicture func(playerView IAVPlayerView)
	_PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart func(playerView IAVPlayerView) bool
	_PlayerViewWillStartPictureInPicture func(playerView IAVPlayerView)
	_PlayerViewWillStopPictureInPicture func(playerView IAVPlayerView)
}

// SetPlayerViewFailedToStartPictureInPictureWithError sets the handler for the PlayerViewFailedToStartPictureInPictureWithError delegate method.
//
// Tells the delegate that Picture in Picture playback failed to start.
func (d *PlayerViewPictureInPictureDelegate) SetPlayerViewFailedToStartPictureInPictureWithError(f func(playerView IAVPlayerView, error_ objc.IObject /* cross-framework: Error */)) {
	d._PlayerViewFailedToStartPictureInPictureWithError = f
}

// SetPlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler sets the handler for the PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler delegate method.
//
// Tells the delegate to restore the user interface before Picture in Picture playback stops.
func (d *PlayerViewPictureInPictureDelegate) SetPlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(f func(playerView IAVPlayerView, completionHandler unsafe.Pointer)) {
	d._PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler = f
}

// SetPlayerViewDidStartPictureInPicture sets the handler for the PlayerViewDidStartPictureInPicture delegate method.
//
// Tells the delegate that Picture in Picture playback started.
func (d *PlayerViewPictureInPictureDelegate) SetPlayerViewDidStartPictureInPicture(f func(playerView IAVPlayerView)) {
	d._PlayerViewDidStartPictureInPicture = f
}

// SetPlayerViewDidStopPictureInPicture sets the handler for the PlayerViewDidStopPictureInPicture delegate method.
//
// Tells the delegate that Picture in Picture playback stopped.
func (d *PlayerViewPictureInPictureDelegate) SetPlayerViewDidStopPictureInPicture(f func(playerView IAVPlayerView)) {
	d._PlayerViewDidStopPictureInPicture = f
}

// SetPlayerViewShouldAutomaticallyDismissAtPictureInPictureStart sets the handler for the PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart delegate method.
//
// Asks the delegate if the player view should miniaturize when Picture in Picture starts.
func (d *PlayerViewPictureInPictureDelegate) SetPlayerViewShouldAutomaticallyDismissAtPictureInPictureStart(f func(playerView IAVPlayerView) bool) {
	d._PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart = f
}

// SetPlayerViewWillStartPictureInPicture sets the handler for the PlayerViewWillStartPictureInPicture delegate method.
//
// Tells the delegate that Picture in Picture playback is about to start.
func (d *PlayerViewPictureInPictureDelegate) SetPlayerViewWillStartPictureInPicture(f func(playerView IAVPlayerView)) {
	d._PlayerViewWillStartPictureInPicture = f
}

// SetPlayerViewWillStopPictureInPicture sets the handler for the PlayerViewWillStopPictureInPicture delegate method.
//
// Tells the delegate that Picture in Picture playback is about to stop.
func (d *PlayerViewPictureInPictureDelegate) SetPlayerViewWillStopPictureInPicture(f func(playerView IAVPlayerView)) {
	d._PlayerViewWillStopPictureInPicture = f
}

// PlayerViewFailedToStartPictureInPictureWithError implements the PPlayerViewPictureInPictureDelegate interface.
func (d *PlayerViewPictureInPictureDelegate) PlayerViewFailedToStartPictureInPictureWithError(playerView IAVPlayerView, error_ objc.IObject /* cross-framework: Error */) {
	if d._PlayerViewFailedToStartPictureInPictureWithError != nil {
		d._PlayerViewFailedToStartPictureInPictureWithError(playerView, error_)
	}
}

// HasPlayerViewFailedToStartPictureInPictureWithError returns true if a handler for PlayerViewFailedToStartPictureInPictureWithError has been set.
func (d *PlayerViewPictureInPictureDelegate) HasPlayerViewFailedToStartPictureInPictureWithError() bool {
	return d._PlayerViewFailedToStartPictureInPictureWithError != nil
}

// PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler implements the PPlayerViewPictureInPictureDelegate interface.
func (d *PlayerViewPictureInPictureDelegate) PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerView IAVPlayerView, completionHandler unsafe.Pointer) {
	if d._PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler != nil {
		d._PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(playerView, completionHandler)
	}
}

// HasPlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler returns true if a handler for PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler has been set.
func (d *PlayerViewPictureInPictureDelegate) HasPlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler() bool {
	return d._PlayerViewRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler != nil
}

// PlayerViewDidStartPictureInPicture implements the PPlayerViewPictureInPictureDelegate interface.
func (d *PlayerViewPictureInPictureDelegate) PlayerViewDidStartPictureInPicture(playerView IAVPlayerView) {
	if d._PlayerViewDidStartPictureInPicture != nil {
		d._PlayerViewDidStartPictureInPicture(playerView)
	}
}

// HasPlayerViewDidStartPictureInPicture returns true if a handler for PlayerViewDidStartPictureInPicture has been set.
func (d *PlayerViewPictureInPictureDelegate) HasPlayerViewDidStartPictureInPicture() bool {
	return d._PlayerViewDidStartPictureInPicture != nil
}

// PlayerViewDidStopPictureInPicture implements the PPlayerViewPictureInPictureDelegate interface.
func (d *PlayerViewPictureInPictureDelegate) PlayerViewDidStopPictureInPicture(playerView IAVPlayerView) {
	if d._PlayerViewDidStopPictureInPicture != nil {
		d._PlayerViewDidStopPictureInPicture(playerView)
	}
}

// HasPlayerViewDidStopPictureInPicture returns true if a handler for PlayerViewDidStopPictureInPicture has been set.
func (d *PlayerViewPictureInPictureDelegate) HasPlayerViewDidStopPictureInPicture() bool {
	return d._PlayerViewDidStopPictureInPicture != nil
}

// PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart implements the PPlayerViewPictureInPictureDelegate interface.
func (d *PlayerViewPictureInPictureDelegate) PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart(playerView IAVPlayerView) bool {
	if d._PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart != nil {
		return d._PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart(playerView)
	}
	var zero bool
	return zero
}

// HasPlayerViewShouldAutomaticallyDismissAtPictureInPictureStart returns true if a handler for PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart has been set.
func (d *PlayerViewPictureInPictureDelegate) HasPlayerViewShouldAutomaticallyDismissAtPictureInPictureStart() bool {
	return d._PlayerViewShouldAutomaticallyDismissAtPictureInPictureStart != nil
}

// PlayerViewWillStartPictureInPicture implements the PPlayerViewPictureInPictureDelegate interface.
func (d *PlayerViewPictureInPictureDelegate) PlayerViewWillStartPictureInPicture(playerView IAVPlayerView) {
	if d._PlayerViewWillStartPictureInPicture != nil {
		d._PlayerViewWillStartPictureInPicture(playerView)
	}
}

// HasPlayerViewWillStartPictureInPicture returns true if a handler for PlayerViewWillStartPictureInPicture has been set.
func (d *PlayerViewPictureInPictureDelegate) HasPlayerViewWillStartPictureInPicture() bool {
	return d._PlayerViewWillStartPictureInPicture != nil
}

// PlayerViewWillStopPictureInPicture implements the PPlayerViewPictureInPictureDelegate interface.
func (d *PlayerViewPictureInPictureDelegate) PlayerViewWillStopPictureInPicture(playerView IAVPlayerView) {
	if d._PlayerViewWillStopPictureInPicture != nil {
		d._PlayerViewWillStopPictureInPicture(playerView)
	}
}

// HasPlayerViewWillStopPictureInPicture returns true if a handler for PlayerViewWillStopPictureInPicture has been set.
func (d *PlayerViewPictureInPictureDelegate) HasPlayerViewWillStopPictureInPicture() bool {
	return d._PlayerViewWillStopPictureInPicture != nil
}
