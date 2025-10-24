// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PPlayerViewDelegate is the AVPlayerViewDelegate protocol interface.
//
// A protocol that defines the methods to implement to participate in the player view’s full-screen presentation life cycle.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVPlayerViewDelegate
type PPlayerViewDelegate interface {
	// Optional methods
	PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerView IAVPlayerView, completionHandler unsafe.Pointer)
	HasPlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler() bool
	PlayerViewDidEnterFullScreen(playerView IAVPlayerView)
	HasPlayerViewDidEnterFullScreen() bool
	PlayerViewDidExitFullScreen(playerView IAVPlayerView)
	HasPlayerViewDidExitFullScreen() bool
	PlayerViewWillEnterFullScreen(playerView IAVPlayerView)
	HasPlayerViewWillEnterFullScreen() bool
	PlayerViewWillExitFullScreen(playerView IAVPlayerView)
	HasPlayerViewWillExitFullScreen() bool
}

// PlayerViewDelegate is a delegate implementation builder for the PPlayerViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayerViewDelegate struct {
	_PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler func(playerView IAVPlayerView, completionHandler unsafe.Pointer)
	_PlayerViewDidEnterFullScreen func(playerView IAVPlayerView)
	_PlayerViewDidExitFullScreen func(playerView IAVPlayerView)
	_PlayerViewWillEnterFullScreen func(playerView IAVPlayerView)
	_PlayerViewWillExitFullScreen func(playerView IAVPlayerView)
}

// SetPlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler sets the handler for the PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler delegate method.
//
// Tells the delegate to restore the app’s user interface when exiting full-screen mode.
func (d *PlayerViewDelegate) SetPlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler(f func(playerView IAVPlayerView, completionHandler unsafe.Pointer)) {
	d._PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler = f
}

// SetPlayerViewDidEnterFullScreen sets the handler for the PlayerViewDidEnterFullScreen delegate method.
//
// Tells the delegate that the player view entered full-screen mode.
func (d *PlayerViewDelegate) SetPlayerViewDidEnterFullScreen(f func(playerView IAVPlayerView)) {
	d._PlayerViewDidEnterFullScreen = f
}

// SetPlayerViewDidExitFullScreen sets the handler for the PlayerViewDidExitFullScreen delegate method.
//
// Tells the delegate that the player view exited full-screen mode.
func (d *PlayerViewDelegate) SetPlayerViewDidExitFullScreen(f func(playerView IAVPlayerView)) {
	d._PlayerViewDidExitFullScreen = f
}

// SetPlayerViewWillEnterFullScreen sets the handler for the PlayerViewWillEnterFullScreen delegate method.
//
// Tells the delegate that the player view is about to enter full-screen mode.
func (d *PlayerViewDelegate) SetPlayerViewWillEnterFullScreen(f func(playerView IAVPlayerView)) {
	d._PlayerViewWillEnterFullScreen = f
}

// SetPlayerViewWillExitFullScreen sets the handler for the PlayerViewWillExitFullScreen delegate method.
//
// Tells the delegate that the player view is about to exit full-screen mode.
func (d *PlayerViewDelegate) SetPlayerViewWillExitFullScreen(f func(playerView IAVPlayerView)) {
	d._PlayerViewWillExitFullScreen = f
}

// PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler implements the PPlayerViewDelegate interface.
func (d *PlayerViewDelegate) PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerView IAVPlayerView, completionHandler unsafe.Pointer) {
	if d._PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler != nil {
		d._PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerView, completionHandler)
	}
}

// HasPlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler returns true if a handler for PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler has been set.
func (d *PlayerViewDelegate) HasPlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler() bool {
	return d._PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler != nil
}

// PlayerViewDidEnterFullScreen implements the PPlayerViewDelegate interface.
func (d *PlayerViewDelegate) PlayerViewDidEnterFullScreen(playerView IAVPlayerView) {
	if d._PlayerViewDidEnterFullScreen != nil {
		d._PlayerViewDidEnterFullScreen(playerView)
	}
}

// HasPlayerViewDidEnterFullScreen returns true if a handler for PlayerViewDidEnterFullScreen has been set.
func (d *PlayerViewDelegate) HasPlayerViewDidEnterFullScreen() bool {
	return d._PlayerViewDidEnterFullScreen != nil
}

// PlayerViewDidExitFullScreen implements the PPlayerViewDelegate interface.
func (d *PlayerViewDelegate) PlayerViewDidExitFullScreen(playerView IAVPlayerView) {
	if d._PlayerViewDidExitFullScreen != nil {
		d._PlayerViewDidExitFullScreen(playerView)
	}
}

// HasPlayerViewDidExitFullScreen returns true if a handler for PlayerViewDidExitFullScreen has been set.
func (d *PlayerViewDelegate) HasPlayerViewDidExitFullScreen() bool {
	return d._PlayerViewDidExitFullScreen != nil
}

// PlayerViewWillEnterFullScreen implements the PPlayerViewDelegate interface.
func (d *PlayerViewDelegate) PlayerViewWillEnterFullScreen(playerView IAVPlayerView) {
	if d._PlayerViewWillEnterFullScreen != nil {
		d._PlayerViewWillEnterFullScreen(playerView)
	}
}

// HasPlayerViewWillEnterFullScreen returns true if a handler for PlayerViewWillEnterFullScreen has been set.
func (d *PlayerViewDelegate) HasPlayerViewWillEnterFullScreen() bool {
	return d._PlayerViewWillEnterFullScreen != nil
}

// PlayerViewWillExitFullScreen implements the PPlayerViewDelegate interface.
func (d *PlayerViewDelegate) PlayerViewWillExitFullScreen(playerView IAVPlayerView) {
	if d._PlayerViewWillExitFullScreen != nil {
		d._PlayerViewWillExitFullScreen(playerView)
	}
}

// HasPlayerViewWillExitFullScreen returns true if a handler for PlayerViewWillExitFullScreen has been set.
func (d *PlayerViewDelegate) HasPlayerViewWillExitFullScreen() bool {
	return d._PlayerViewWillExitFullScreen != nil
}
