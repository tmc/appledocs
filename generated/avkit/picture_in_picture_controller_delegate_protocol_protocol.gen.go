// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PPictureInPictureControllerDelegate is the AVPictureInPictureControllerDelegate protocol interface.
//
// A protocol to adopt to respond to Picture in Picture events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.15+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVPictureInPictureControllerDelegate
type PPictureInPictureControllerDelegate interface {
	// Optional methods
	PictureInPictureControllerFailedToStartPictureInPictureWithError(pictureInPictureController IAVPictureInPictureController, error_ objc.IObject /* cross-framework: Error */)
	HasPictureInPictureControllerFailedToStartPictureInPictureWithError() bool
	PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(pictureInPictureController IAVPictureInPictureController, completionHandler unsafe.Pointer)
	HasPictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler() bool
	PictureInPictureControllerDidStartPictureInPicture(pictureInPictureController IAVPictureInPictureController)
	HasPictureInPictureControllerDidStartPictureInPicture() bool
	PictureInPictureControllerDidStopPictureInPicture(pictureInPictureController IAVPictureInPictureController)
	HasPictureInPictureControllerDidStopPictureInPicture() bool
	PictureInPictureControllerWillStartPictureInPicture(pictureInPictureController IAVPictureInPictureController)
	HasPictureInPictureControllerWillStartPictureInPicture() bool
	PictureInPictureControllerWillStopPictureInPicture(pictureInPictureController IAVPictureInPictureController)
	HasPictureInPictureControllerWillStopPictureInPicture() bool
}

// PictureInPictureControllerDelegate is a delegate implementation builder for the PPictureInPictureControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PictureInPictureControllerDelegate struct {
	_PictureInPictureControllerFailedToStartPictureInPictureWithError func(pictureInPictureController IAVPictureInPictureController, error_ objc.IObject /* cross-framework: Error */)
	_PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler func(pictureInPictureController IAVPictureInPictureController, completionHandler unsafe.Pointer)
	_PictureInPictureControllerDidStartPictureInPicture func(pictureInPictureController IAVPictureInPictureController)
	_PictureInPictureControllerDidStopPictureInPicture func(pictureInPictureController IAVPictureInPictureController)
	_PictureInPictureControllerWillStartPictureInPicture func(pictureInPictureController IAVPictureInPictureController)
	_PictureInPictureControllerWillStopPictureInPicture func(pictureInPictureController IAVPictureInPictureController)
}

// SetPictureInPictureControllerFailedToStartPictureInPictureWithError sets the handler for the PictureInPictureControllerFailedToStartPictureInPictureWithError delegate method.
//
// Tells the delegate that Picture in Picture failed to start.
func (d *PictureInPictureControllerDelegate) SetPictureInPictureControllerFailedToStartPictureInPictureWithError(f func(pictureInPictureController IAVPictureInPictureController, error_ objc.IObject /* cross-framework: Error */)) {
	d._PictureInPictureControllerFailedToStartPictureInPictureWithError = f
}

// SetPictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler sets the handler for the PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler delegate method.
//
// Tells the delegate to restore the user interface before Picture in Picture stops.
func (d *PictureInPictureControllerDelegate) SetPictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(f func(pictureInPictureController IAVPictureInPictureController, completionHandler unsafe.Pointer)) {
	d._PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler = f
}

// SetPictureInPictureControllerDidStartPictureInPicture sets the handler for the PictureInPictureControllerDidStartPictureInPicture delegate method.
//
// Tells the delegate that Picture in Picture started.
func (d *PictureInPictureControllerDelegate) SetPictureInPictureControllerDidStartPictureInPicture(f func(pictureInPictureController IAVPictureInPictureController)) {
	d._PictureInPictureControllerDidStartPictureInPicture = f
}

// SetPictureInPictureControllerDidStopPictureInPicture sets the handler for the PictureInPictureControllerDidStopPictureInPicture delegate method.
//
// Tells the delegate that Picture in Picture stopped.
func (d *PictureInPictureControllerDelegate) SetPictureInPictureControllerDidStopPictureInPicture(f func(pictureInPictureController IAVPictureInPictureController)) {
	d._PictureInPictureControllerDidStopPictureInPicture = f
}

// SetPictureInPictureControllerWillStartPictureInPicture sets the handler for the PictureInPictureControllerWillStartPictureInPicture delegate method.
//
// Tells the delegate that Picture in Picture is about to start.
func (d *PictureInPictureControllerDelegate) SetPictureInPictureControllerWillStartPictureInPicture(f func(pictureInPictureController IAVPictureInPictureController)) {
	d._PictureInPictureControllerWillStartPictureInPicture = f
}

// SetPictureInPictureControllerWillStopPictureInPicture sets the handler for the PictureInPictureControllerWillStopPictureInPicture delegate method.
//
// Tells the delegate that Picture in Picture is about to stop.
func (d *PictureInPictureControllerDelegate) SetPictureInPictureControllerWillStopPictureInPicture(f func(pictureInPictureController IAVPictureInPictureController)) {
	d._PictureInPictureControllerWillStopPictureInPicture = f
}

// PictureInPictureControllerFailedToStartPictureInPictureWithError implements the PPictureInPictureControllerDelegate interface.
func (d *PictureInPictureControllerDelegate) PictureInPictureControllerFailedToStartPictureInPictureWithError(pictureInPictureController IAVPictureInPictureController, error_ objc.IObject /* cross-framework: Error */) {
	if d._PictureInPictureControllerFailedToStartPictureInPictureWithError != nil {
		d._PictureInPictureControllerFailedToStartPictureInPictureWithError(pictureInPictureController, error_)
	}
}

// HasPictureInPictureControllerFailedToStartPictureInPictureWithError returns true if a handler for PictureInPictureControllerFailedToStartPictureInPictureWithError has been set.
func (d *PictureInPictureControllerDelegate) HasPictureInPictureControllerFailedToStartPictureInPictureWithError() bool {
	return d._PictureInPictureControllerFailedToStartPictureInPictureWithError != nil
}

// PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler implements the PPictureInPictureControllerDelegate interface.
func (d *PictureInPictureControllerDelegate) PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(pictureInPictureController IAVPictureInPictureController, completionHandler unsafe.Pointer) {
	if d._PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler != nil {
		d._PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler(pictureInPictureController, completionHandler)
	}
}

// HasPictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler returns true if a handler for PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler has been set.
func (d *PictureInPictureControllerDelegate) HasPictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler() bool {
	return d._PictureInPictureControllerRestoreUserInterfaceForPictureInPictureStopWithCompletionHandler != nil
}

// PictureInPictureControllerDidStartPictureInPicture implements the PPictureInPictureControllerDelegate interface.
func (d *PictureInPictureControllerDelegate) PictureInPictureControllerDidStartPictureInPicture(pictureInPictureController IAVPictureInPictureController) {
	if d._PictureInPictureControllerDidStartPictureInPicture != nil {
		d._PictureInPictureControllerDidStartPictureInPicture(pictureInPictureController)
	}
}

// HasPictureInPictureControllerDidStartPictureInPicture returns true if a handler for PictureInPictureControllerDidStartPictureInPicture has been set.
func (d *PictureInPictureControllerDelegate) HasPictureInPictureControllerDidStartPictureInPicture() bool {
	return d._PictureInPictureControllerDidStartPictureInPicture != nil
}

// PictureInPictureControllerDidStopPictureInPicture implements the PPictureInPictureControllerDelegate interface.
func (d *PictureInPictureControllerDelegate) PictureInPictureControllerDidStopPictureInPicture(pictureInPictureController IAVPictureInPictureController) {
	if d._PictureInPictureControllerDidStopPictureInPicture != nil {
		d._PictureInPictureControllerDidStopPictureInPicture(pictureInPictureController)
	}
}

// HasPictureInPictureControllerDidStopPictureInPicture returns true if a handler for PictureInPictureControllerDidStopPictureInPicture has been set.
func (d *PictureInPictureControllerDelegate) HasPictureInPictureControllerDidStopPictureInPicture() bool {
	return d._PictureInPictureControllerDidStopPictureInPicture != nil
}

// PictureInPictureControllerWillStartPictureInPicture implements the PPictureInPictureControllerDelegate interface.
func (d *PictureInPictureControllerDelegate) PictureInPictureControllerWillStartPictureInPicture(pictureInPictureController IAVPictureInPictureController) {
	if d._PictureInPictureControllerWillStartPictureInPicture != nil {
		d._PictureInPictureControllerWillStartPictureInPicture(pictureInPictureController)
	}
}

// HasPictureInPictureControllerWillStartPictureInPicture returns true if a handler for PictureInPictureControllerWillStartPictureInPicture has been set.
func (d *PictureInPictureControllerDelegate) HasPictureInPictureControllerWillStartPictureInPicture() bool {
	return d._PictureInPictureControllerWillStartPictureInPicture != nil
}

// PictureInPictureControllerWillStopPictureInPicture implements the PPictureInPictureControllerDelegate interface.
func (d *PictureInPictureControllerDelegate) PictureInPictureControllerWillStopPictureInPicture(pictureInPictureController IAVPictureInPictureController) {
	if d._PictureInPictureControllerWillStopPictureInPicture != nil {
		d._PictureInPictureControllerWillStopPictureInPicture(pictureInPictureController)
	}
}

// HasPictureInPictureControllerWillStopPictureInPicture returns true if a handler for PictureInPictureControllerWillStopPictureInPicture has been set.
func (d *PictureInPictureControllerDelegate) HasPictureInPictureControllerWillStopPictureInPicture() bool {
	return d._PictureInPictureControllerWillStopPictureInPicture != nil
}
