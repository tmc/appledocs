// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PRPScreenRecorderDelegate is the RPScreenRecorderDelegate protocol interface.
//
// The protocol you implement to receive notifications from the screen recorder.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 11.0+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.replaykit/documentation/ReplayKit/RPScreenRecorderDelegate
type PRPScreenRecorderDelegate interface {
	// Optional methods
	ScreenRecorderDidStopRecordingWithPreviewViewControllerError(screenRecorder IRPScreenRecorder, previewViewController IRPPreviewViewController, error_ objc.IObject /* cross-framework: Error */)
	HasScreenRecorderDidStopRecordingWithPreviewViewControllerError() bool
	ScreenRecorderDidStopRecordingWithErrorPreviewViewController(screenRecorder IRPScreenRecorder, error_ objc.IObject /* cross-framework: Error */, previewViewController IRPPreviewViewController)
	HasScreenRecorderDidStopRecordingWithErrorPreviewViewController() bool
	ScreenRecorderDidChangeAvailability(screenRecorder IRPScreenRecorder)
	HasScreenRecorderDidChangeAvailability() bool
}

// RPScreenRecorderDelegate is a delegate implementation builder for the PRPScreenRecorderDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type RPScreenRecorderDelegate struct {
	_ScreenRecorderDidStopRecordingWithPreviewViewControllerError func(screenRecorder IRPScreenRecorder, previewViewController IRPPreviewViewController, error_ objc.IObject /* cross-framework: Error */)
	_ScreenRecorderDidStopRecordingWithErrorPreviewViewController func(screenRecorder IRPScreenRecorder, error_ objc.IObject /* cross-framework: Error */, previewViewController IRPPreviewViewController)
	_ScreenRecorderDidChangeAvailability func(screenRecorder IRPScreenRecorder)
}

// SetScreenRecorderDidStopRecordingWithPreviewViewControllerError sets the handler for the ScreenRecorderDidStopRecordingWithPreviewViewControllerError delegate method.
//
// Indicates that the screen recording has stopped.
func (d *RPScreenRecorderDelegate) SetScreenRecorderDidStopRecordingWithPreviewViewControllerError(f func(screenRecorder IRPScreenRecorder, previewViewController IRPPreviewViewController, error_ objc.IObject /* cross-framework: Error */)) {
	d._ScreenRecorderDidStopRecordingWithPreviewViewControllerError = f
}

// SetScreenRecorderDidStopRecordingWithErrorPreviewViewController sets the handler for the ScreenRecorderDidStopRecordingWithErrorPreviewViewController delegate method.
//
// Indicates that the screen recording has stopped.
func (d *RPScreenRecorderDelegate) SetScreenRecorderDidStopRecordingWithErrorPreviewViewController(f func(screenRecorder IRPScreenRecorder, error_ objc.IObject /* cross-framework: Error */, previewViewController IRPPreviewViewController)) {
	d._ScreenRecorderDidStopRecordingWithErrorPreviewViewController = f
}

// SetScreenRecorderDidChangeAvailability sets the handler for the ScreenRecorderDidChangeAvailability delegate method.
//
// Indicates that the recorder has changed states between disabled and enabled.
func (d *RPScreenRecorderDelegate) SetScreenRecorderDidChangeAvailability(f func(screenRecorder IRPScreenRecorder)) {
	d._ScreenRecorderDidChangeAvailability = f
}

// ScreenRecorderDidStopRecordingWithPreviewViewControllerError implements the PRPScreenRecorderDelegate interface.
func (d *RPScreenRecorderDelegate) ScreenRecorderDidStopRecordingWithPreviewViewControllerError(screenRecorder IRPScreenRecorder, previewViewController IRPPreviewViewController, error_ objc.IObject /* cross-framework: Error */) {
	if d._ScreenRecorderDidStopRecordingWithPreviewViewControllerError != nil {
		d._ScreenRecorderDidStopRecordingWithPreviewViewControllerError(screenRecorder, previewViewController, error_)
	}
}

// HasScreenRecorderDidStopRecordingWithPreviewViewControllerError returns true if a handler for ScreenRecorderDidStopRecordingWithPreviewViewControllerError has been set.
func (d *RPScreenRecorderDelegate) HasScreenRecorderDidStopRecordingWithPreviewViewControllerError() bool {
	return d._ScreenRecorderDidStopRecordingWithPreviewViewControllerError != nil
}

// ScreenRecorderDidStopRecordingWithErrorPreviewViewController implements the PRPScreenRecorderDelegate interface.
func (d *RPScreenRecorderDelegate) ScreenRecorderDidStopRecordingWithErrorPreviewViewController(screenRecorder IRPScreenRecorder, error_ objc.IObject /* cross-framework: Error */, previewViewController IRPPreviewViewController) {
	if d._ScreenRecorderDidStopRecordingWithErrorPreviewViewController != nil {
		d._ScreenRecorderDidStopRecordingWithErrorPreviewViewController(screenRecorder, error_, previewViewController)
	}
}

// HasScreenRecorderDidStopRecordingWithErrorPreviewViewController returns true if a handler for ScreenRecorderDidStopRecordingWithErrorPreviewViewController has been set.
func (d *RPScreenRecorderDelegate) HasScreenRecorderDidStopRecordingWithErrorPreviewViewController() bool {
	return d._ScreenRecorderDidStopRecordingWithErrorPreviewViewController != nil
}

// ScreenRecorderDidChangeAvailability implements the PRPScreenRecorderDelegate interface.
func (d *RPScreenRecorderDelegate) ScreenRecorderDidChangeAvailability(screenRecorder IRPScreenRecorder) {
	if d._ScreenRecorderDidChangeAvailability != nil {
		d._ScreenRecorderDidChangeAvailability(screenRecorder)
	}
}

// HasScreenRecorderDidChangeAvailability returns true if a handler for ScreenRecorderDidChangeAvailability has been set.
func (d *RPScreenRecorderDelegate) HasScreenRecorderDidChangeAvailability() bool {
	return d._ScreenRecorderDidChangeAvailability != nil
}
