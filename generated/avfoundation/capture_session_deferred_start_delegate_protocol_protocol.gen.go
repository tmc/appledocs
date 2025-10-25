// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCaptureSessionDeferredStartDelegate is the AVCaptureSessionDeferredStartDelegate protocol interface.
//
// A protocol that defines the interface to respond to events about a capture session’s deferred start.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureSessionDeferredStartDelegate
type PCaptureSessionDeferredStartDelegate interface {
	// Required methods
	SessionDidRunDeferredStart(session IAVCaptureSession)/* debug [protocol_interface/required_method]: SessionDidRunDeferredStart */
	SessionWillRunDeferredStart(session IAVCaptureSession)/* debug [protocol_interface/required_method]: SessionWillRunDeferredStart */
}

// CaptureSessionDeferredStartDelegate is a delegate implementation builder for the PCaptureSessionDeferredStartDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureSessionDeferredStartDelegate struct {
	_SessionDidRunDeferredStart func(session IAVCaptureSession)
	_SessionWillRunDeferredStart func(session IAVCaptureSession)
}

// SetSessionDidRunDeferredStart sets the handler for the SessionDidRunDeferredStart delegate method.
//
// This method gets called by the session when deferred start has finished running.
func (d *CaptureSessionDeferredStartDelegate) SetSessionDidRunDeferredStart(f func(session IAVCaptureSession)) {
	d._SessionDidRunDeferredStart = f
}

// SetSessionWillRunDeferredStart sets the handler for the SessionWillRunDeferredStart delegate method.
//
// This method gets called by the session when deferred start is about to run.
func (d *CaptureSessionDeferredStartDelegate) SetSessionWillRunDeferredStart(f func(session IAVCaptureSession)) {
	d._SessionWillRunDeferredStart = f
}

// SessionDidRunDeferredStart implements the PCaptureSessionDeferredStartDelegate interface.
func (d *CaptureSessionDeferredStartDelegate) SessionDidRunDeferredStart(session IAVCaptureSession) {
	if d._SessionDidRunDeferredStart != nil {
		d._SessionDidRunDeferredStart(session)
	}
}

// HasSessionDidRunDeferredStart returns true if a handler for SessionDidRunDeferredStart has been set.
func (d *CaptureSessionDeferredStartDelegate) HasSessionDidRunDeferredStart() bool {
	return d._SessionDidRunDeferredStart != nil
}

// SessionWillRunDeferredStart implements the PCaptureSessionDeferredStartDelegate interface.
func (d *CaptureSessionDeferredStartDelegate) SessionWillRunDeferredStart(session IAVCaptureSession) {
	if d._SessionWillRunDeferredStart != nil {
		d._SessionWillRunDeferredStart(session)
	}
}

// HasSessionWillRunDeferredStart returns true if a handler for SessionWillRunDeferredStart has been set.
func (d *CaptureSessionDeferredStartDelegate) HasSessionWillRunDeferredStart() bool {
	return d._SessionWillRunDeferredStart != nil
}
