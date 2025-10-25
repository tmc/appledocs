// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	PlaybackCoordinatorDidIssueSeekCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, seekCommand IAVDelegatingPlaybackCoordinatorSeekCommand, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: PlaybackCoordinatorDidIssueSeekCommandCompletionHandler */
	PlaybackCoordinatorDidIssuePauseCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, pauseCommand IAVDelegatingPlaybackCoordinatorPauseCommand, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: PlaybackCoordinatorDidIssuePauseCommandCompletionHandler */
	PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, playCommand IAVDelegatingPlaybackCoordinatorPlayCommand, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: PlaybackCoordinatorDidIssuePlayCommandCompletionHandler */
	PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, bufferingCommand IAVDelegatingPlaybackCoordinatorBufferingCommand, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler */
}

// PlaybackCoordinatorPlaybackControlDelegate is a delegate implementation builder for the PPlaybackCoordinatorPlaybackControlDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlaybackCoordinatorPlaybackControlDelegate struct {
	_PlaybackCoordinatorDidIssueSeekCommandCompletionHandler func(coordinator IAVDelegatingPlaybackCoordinator, seekCommand IAVDelegatingPlaybackCoordinatorSeekCommand, completionHandler unsafe.Pointer)
	_PlaybackCoordinatorDidIssuePauseCommandCompletionHandler func(coordinator IAVDelegatingPlaybackCoordinator, pauseCommand IAVDelegatingPlaybackCoordinatorPauseCommand, completionHandler unsafe.Pointer)
	_PlaybackCoordinatorDidIssuePlayCommandCompletionHandler func(coordinator IAVDelegatingPlaybackCoordinator, playCommand IAVDelegatingPlaybackCoordinatorPlayCommand, completionHandler unsafe.Pointer)
	_PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler func(coordinator IAVDelegatingPlaybackCoordinator, bufferingCommand IAVDelegatingPlaybackCoordinatorBufferingCommand, completionHandler unsafe.Pointer)
}

// SetPlaybackCoordinatorDidIssueSeekCommandCompletionHandler sets the handler for the PlaybackCoordinatorDidIssueSeekCommandCompletionHandler delegate method.
//
// Tells the delegate to seek to a new time.
func (d *PlaybackCoordinatorPlaybackControlDelegate) SetPlaybackCoordinatorDidIssueSeekCommandCompletionHandler(f func(coordinator IAVDelegatingPlaybackCoordinator, seekCommand IAVDelegatingPlaybackCoordinatorSeekCommand, completionHandler unsafe.Pointer)) {
	d._PlaybackCoordinatorDidIssueSeekCommandCompletionHandler = f
}

// SetPlaybackCoordinatorDidIssuePauseCommandCompletionHandler sets the handler for the PlaybackCoordinatorDidIssuePauseCommandCompletionHandler delegate method.
//
// Tells the delegate to pause playback.
func (d *PlaybackCoordinatorPlaybackControlDelegate) SetPlaybackCoordinatorDidIssuePauseCommandCompletionHandler(f func(coordinator IAVDelegatingPlaybackCoordinator, pauseCommand IAVDelegatingPlaybackCoordinatorPauseCommand, completionHandler unsafe.Pointer)) {
	d._PlaybackCoordinatorDidIssuePauseCommandCompletionHandler = f
}

// SetPlaybackCoordinatorDidIssuePlayCommandCompletionHandler sets the handler for the PlaybackCoordinatorDidIssuePlayCommandCompletionHandler delegate method.
//
// Tells the delegate to match the playback rate to that of the group when the rate is nonzero.
func (d *PlaybackCoordinatorPlaybackControlDelegate) SetPlaybackCoordinatorDidIssuePlayCommandCompletionHandler(f func(coordinator IAVDelegatingPlaybackCoordinator, playCommand IAVDelegatingPlaybackCoordinatorPlayCommand, completionHandler unsafe.Pointer)) {
	d._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler = f
}

// SetPlaybackCoordinatorDidIssueBufferingCommandCompletionHandler sets the handler for the PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler delegate method.
//
// Tells the delegate to expect playback soon and to start buffering media data in preparation.
func (d *PlaybackCoordinatorPlaybackControlDelegate) SetPlaybackCoordinatorDidIssueBufferingCommandCompletionHandler(f func(coordinator IAVDelegatingPlaybackCoordinator, bufferingCommand IAVDelegatingPlaybackCoordinatorBufferingCommand, completionHandler unsafe.Pointer)) {
	d._PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler = f
}

// PlaybackCoordinatorDidIssueSeekCommandCompletionHandler implements the PPlaybackCoordinatorPlaybackControlDelegate interface.
func (d *PlaybackCoordinatorPlaybackControlDelegate) PlaybackCoordinatorDidIssueSeekCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, seekCommand IAVDelegatingPlaybackCoordinatorSeekCommand, completionHandler unsafe.Pointer) {
	if d._PlaybackCoordinatorDidIssueSeekCommandCompletionHandler != nil {
		d._PlaybackCoordinatorDidIssueSeekCommandCompletionHandler(coordinator, seekCommand, completionHandler)
	}
}

// HasPlaybackCoordinatorDidIssueSeekCommandCompletionHandler returns true if a handler for PlaybackCoordinatorDidIssueSeekCommandCompletionHandler has been set.
func (d *PlaybackCoordinatorPlaybackControlDelegate) HasPlaybackCoordinatorDidIssueSeekCommandCompletionHandler() bool {
	return d._PlaybackCoordinatorDidIssueSeekCommandCompletionHandler != nil
}

// PlaybackCoordinatorDidIssuePauseCommandCompletionHandler implements the PPlaybackCoordinatorPlaybackControlDelegate interface.
func (d *PlaybackCoordinatorPlaybackControlDelegate) PlaybackCoordinatorDidIssuePauseCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, pauseCommand IAVDelegatingPlaybackCoordinatorPauseCommand, completionHandler unsafe.Pointer) {
	if d._PlaybackCoordinatorDidIssuePauseCommandCompletionHandler != nil {
		d._PlaybackCoordinatorDidIssuePauseCommandCompletionHandler(coordinator, pauseCommand, completionHandler)
	}
}

// HasPlaybackCoordinatorDidIssuePauseCommandCompletionHandler returns true if a handler for PlaybackCoordinatorDidIssuePauseCommandCompletionHandler has been set.
func (d *PlaybackCoordinatorPlaybackControlDelegate) HasPlaybackCoordinatorDidIssuePauseCommandCompletionHandler() bool {
	return d._PlaybackCoordinatorDidIssuePauseCommandCompletionHandler != nil
}

// PlaybackCoordinatorDidIssuePlayCommandCompletionHandler implements the PPlaybackCoordinatorPlaybackControlDelegate interface.
func (d *PlaybackCoordinatorPlaybackControlDelegate) PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, playCommand IAVDelegatingPlaybackCoordinatorPlayCommand, completionHandler unsafe.Pointer) {
	if d._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler != nil {
		d._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler(coordinator, playCommand, completionHandler)
	}
}

// HasPlaybackCoordinatorDidIssuePlayCommandCompletionHandler returns true if a handler for PlaybackCoordinatorDidIssuePlayCommandCompletionHandler has been set.
func (d *PlaybackCoordinatorPlaybackControlDelegate) HasPlaybackCoordinatorDidIssuePlayCommandCompletionHandler() bool {
	return d._PlaybackCoordinatorDidIssuePlayCommandCompletionHandler != nil
}

// PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler implements the PPlaybackCoordinatorPlaybackControlDelegate interface.
func (d *PlaybackCoordinatorPlaybackControlDelegate) PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler(coordinator IAVDelegatingPlaybackCoordinator, bufferingCommand IAVDelegatingPlaybackCoordinatorBufferingCommand, completionHandler unsafe.Pointer) {
	if d._PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler != nil {
		d._PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler(coordinator, bufferingCommand, completionHandler)
	}
}

// HasPlaybackCoordinatorDidIssueBufferingCommandCompletionHandler returns true if a handler for PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler has been set.
func (d *PlaybackCoordinatorPlaybackControlDelegate) HasPlaybackCoordinatorDidIssueBufferingCommandCompletionHandler() bool {
	return d._PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler != nil
}
