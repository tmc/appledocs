// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corevideo"
)

// PPictureInPictureSampleBufferPlaybackDelegate is the AVPictureInPictureSampleBufferPlaybackDelegate protocol interface.
//
// A protocol for controlling playback from a sample buffer display layer in Picture in Picture.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVPictureInPictureSampleBufferPlaybackDelegate
type PPictureInPictureSampleBufferPlaybackDelegate interface {
	// Required methods
	PictureInPictureControllerDidTransitionToRenderSize(pictureInPictureController IAVPictureInPictureController, newRenderSize VideoDimensions /* not a class type */)/* debug [protocol_interface/required_method]: PictureInPictureControllerDidTransitionToRenderSize */
	PictureInPictureControllerSetPlaying(pictureInPictureController IAVPictureInPictureController, playing bool)/* debug [protocol_interface/required_method]: PictureInPictureControllerSetPlaying */
	PictureInPictureControllerSkipByIntervalCompletionHandler(pictureInPictureController IAVPictureInPictureController, skipInterval objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: PictureInPictureControllerSkipByIntervalCompletionHandler */
	PictureInPictureControllerIsPlaybackPaused(pictureInPictureController IAVPictureInPictureController) bool/* debug [protocol_interface/required_method]: PictureInPictureControllerIsPlaybackPaused */
	PictureInPictureControllerTimeRangeForPlayback(pictureInPictureController IAVPictureInPictureController) TimeRange/* debug [protocol_interface/required_method]: PictureInPictureControllerTimeRangeForPlayback */
	// Optional methods
	PictureInPictureControllerShouldProhibitBackgroundAudioPlayback(pictureInPictureController IAVPictureInPictureController) bool
	HasPictureInPictureControllerShouldProhibitBackgroundAudioPlayback() bool
}

// PictureInPictureSampleBufferPlaybackDelegate is a delegate implementation builder for the PPictureInPictureSampleBufferPlaybackDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PictureInPictureSampleBufferPlaybackDelegate struct {
	_PictureInPictureControllerShouldProhibitBackgroundAudioPlayback func(pictureInPictureController IAVPictureInPictureController) bool
	_PictureInPictureControllerDidTransitionToRenderSize func(pictureInPictureController IAVPictureInPictureController, newRenderSize VideoDimensions /* not a class type */)
	_PictureInPictureControllerSetPlaying func(pictureInPictureController IAVPictureInPictureController, playing bool)
	_PictureInPictureControllerSkipByIntervalCompletionHandler func(pictureInPictureController IAVPictureInPictureController, skipInterval objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)
	_PictureInPictureControllerIsPlaybackPaused func(pictureInPictureController IAVPictureInPictureController) bool
	_PictureInPictureControllerTimeRangeForPlayback func(pictureInPictureController IAVPictureInPictureController) TimeRange
}

// SetPictureInPictureControllerShouldProhibitBackgroundAudioPlayback sets the handler for the PictureInPictureControllerShouldProhibitBackgroundAudioPlayback delegate method.
//
// Asks the delegate whether to always prohibit background audio playback.
func (d *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerShouldProhibitBackgroundAudioPlayback(f func(pictureInPictureController IAVPictureInPictureController) bool) {
	d._PictureInPictureControllerShouldProhibitBackgroundAudioPlayback = f
}

// SetPictureInPictureControllerDidTransitionToRenderSize sets the handler for the PictureInPictureControllerDidTransitionToRenderSize delegate method.
//
// Tells the delegate when the system Picture in Picture window changes size.
func (d *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerDidTransitionToRenderSize(f func(pictureInPictureController IAVPictureInPictureController, newRenderSize VideoDimensions /* not a class type */)) {
	d._PictureInPictureControllerDidTransitionToRenderSize = f
}

// SetPictureInPictureControllerSetPlaying sets the handler for the PictureInPictureControllerSetPlaying delegate method.
//
// Tells the delegate that the user requested to begin or pause playback.
func (d *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerSetPlaying(f func(pictureInPictureController IAVPictureInPictureController, playing bool)) {
	d._PictureInPictureControllerSetPlaying = f
}

// SetPictureInPictureControllerSkipByIntervalCompletionHandler sets the handler for the PictureInPictureControllerSkipByIntervalCompletionHandler delegate method.
//
// Tells the delegate that the user has requested skipping forward or backward by the indicated time interval.
func (d *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerSkipByIntervalCompletionHandler(f func(pictureInPictureController IAVPictureInPictureController, skipInterval objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)) {
	d._PictureInPictureControllerSkipByIntervalCompletionHandler = f
}

// SetPictureInPictureControllerIsPlaybackPaused sets the handler for the PictureInPictureControllerIsPlaybackPaused delegate method.
//
// Asks delegate to indicate whether the playback UI reflects a playing or paused state, regardless of the current playback rate.
func (d *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerIsPlaybackPaused(f func(pictureInPictureController IAVPictureInPictureController) bool) {
	d._PictureInPictureControllerIsPlaybackPaused = f
}

// SetPictureInPictureControllerTimeRangeForPlayback sets the handler for the PictureInPictureControllerTimeRangeForPlayback delegate method.
//
// Asks the delegate for the current playable time range.
func (d *PictureInPictureSampleBufferPlaybackDelegate) SetPictureInPictureControllerTimeRangeForPlayback(f func(pictureInPictureController IAVPictureInPictureController) TimeRange) {
	d._PictureInPictureControllerTimeRangeForPlayback = f
}

// PictureInPictureControllerShouldProhibitBackgroundAudioPlayback implements the PPictureInPictureSampleBufferPlaybackDelegate interface.
func (d *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerShouldProhibitBackgroundAudioPlayback(pictureInPictureController IAVPictureInPictureController) bool {
	if d._PictureInPictureControllerShouldProhibitBackgroundAudioPlayback != nil {
		return d._PictureInPictureControllerShouldProhibitBackgroundAudioPlayback(pictureInPictureController)
	}
	var zero bool
	return zero
}

// HasPictureInPictureControllerShouldProhibitBackgroundAudioPlayback returns true if a handler for PictureInPictureControllerShouldProhibitBackgroundAudioPlayback has been set.
func (d *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerShouldProhibitBackgroundAudioPlayback() bool {
	return d._PictureInPictureControllerShouldProhibitBackgroundAudioPlayback != nil
}

// PictureInPictureControllerDidTransitionToRenderSize implements the PPictureInPictureSampleBufferPlaybackDelegate interface.
func (d *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerDidTransitionToRenderSize(pictureInPictureController IAVPictureInPictureController, newRenderSize VideoDimensions /* not a class type */) {
	if d._PictureInPictureControllerDidTransitionToRenderSize != nil {
		d._PictureInPictureControllerDidTransitionToRenderSize(pictureInPictureController, newRenderSize)
	}
}

// HasPictureInPictureControllerDidTransitionToRenderSize returns true if a handler for PictureInPictureControllerDidTransitionToRenderSize has been set.
func (d *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerDidTransitionToRenderSize() bool {
	return d._PictureInPictureControllerDidTransitionToRenderSize != nil
}

// PictureInPictureControllerSetPlaying implements the PPictureInPictureSampleBufferPlaybackDelegate interface.
func (d *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerSetPlaying(pictureInPictureController IAVPictureInPictureController, playing bool) {
	if d._PictureInPictureControllerSetPlaying != nil {
		d._PictureInPictureControllerSetPlaying(pictureInPictureController, playing)
	}
}

// HasPictureInPictureControllerSetPlaying returns true if a handler for PictureInPictureControllerSetPlaying has been set.
func (d *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerSetPlaying() bool {
	return d._PictureInPictureControllerSetPlaying != nil
}

// PictureInPictureControllerSkipByIntervalCompletionHandler implements the PPictureInPictureSampleBufferPlaybackDelegate interface.
func (d *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerSkipByIntervalCompletionHandler(pictureInPictureController IAVPictureInPictureController, skipInterval objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	if d._PictureInPictureControllerSkipByIntervalCompletionHandler != nil {
		d._PictureInPictureControllerSkipByIntervalCompletionHandler(pictureInPictureController, skipInterval, completionHandler)
	}
}

// HasPictureInPictureControllerSkipByIntervalCompletionHandler returns true if a handler for PictureInPictureControllerSkipByIntervalCompletionHandler has been set.
func (d *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerSkipByIntervalCompletionHandler() bool {
	return d._PictureInPictureControllerSkipByIntervalCompletionHandler != nil
}

// PictureInPictureControllerIsPlaybackPaused implements the PPictureInPictureSampleBufferPlaybackDelegate interface.
func (d *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerIsPlaybackPaused(pictureInPictureController IAVPictureInPictureController) bool {
	if d._PictureInPictureControllerIsPlaybackPaused != nil {
		return d._PictureInPictureControllerIsPlaybackPaused(pictureInPictureController)
	}
	var zero bool
	return zero
}

// HasPictureInPictureControllerIsPlaybackPaused returns true if a handler for PictureInPictureControllerIsPlaybackPaused has been set.
func (d *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerIsPlaybackPaused() bool {
	return d._PictureInPictureControllerIsPlaybackPaused != nil
}

// PictureInPictureControllerTimeRangeForPlayback implements the PPictureInPictureSampleBufferPlaybackDelegate interface.
func (d *PictureInPictureSampleBufferPlaybackDelegate) PictureInPictureControllerTimeRangeForPlayback(pictureInPictureController IAVPictureInPictureController) TimeRange {
	if d._PictureInPictureControllerTimeRangeForPlayback != nil {
		return d._PictureInPictureControllerTimeRangeForPlayback(pictureInPictureController)
	}
	var zero TimeRange
	return zero
}

// HasPictureInPictureControllerTimeRangeForPlayback returns true if a handler for PictureInPictureControllerTimeRangeForPlayback has been set.
func (d *PictureInPictureSampleBufferPlaybackDelegate) HasPictureInPictureControllerTimeRangeForPlayback() bool {
	return d._PictureInPictureControllerTimeRangeForPlayback != nil
}
