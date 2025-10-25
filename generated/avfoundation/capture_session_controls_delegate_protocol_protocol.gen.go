// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCaptureSessionControlsDelegate is the AVCaptureSessionControlsDelegate protocol interface.
//
// A protocol that defines the interface to respond to capture control activation and presentation events.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureSessionControlsDelegate
type PCaptureSessionControlsDelegate interface {
	// Required methods
	SessionControlsDidBecomeActive(session IAVCaptureSession)/* debug [protocol_interface/required_method]: SessionControlsDidBecomeActive */
	SessionControlsDidBecomeInactive(session IAVCaptureSession)/* debug [protocol_interface/required_method]: SessionControlsDidBecomeInactive */
	SessionControlsWillEnterFullscreenAppearance(session IAVCaptureSession)/* debug [protocol_interface/required_method]: SessionControlsWillEnterFullscreenAppearance */
	SessionControlsWillExitFullscreenAppearance(session IAVCaptureSession)/* debug [protocol_interface/required_method]: SessionControlsWillExitFullscreenAppearance */
}

// CaptureSessionControlsDelegate is a delegate implementation builder for the PCaptureSessionControlsDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureSessionControlsDelegate struct {
	_SessionControlsDidBecomeActive func(session IAVCaptureSession)
	_SessionControlsDidBecomeInactive func(session IAVCaptureSession)
	_SessionControlsWillEnterFullscreenAppearance func(session IAVCaptureSession)
	_SessionControlsWillExitFullscreenAppearance func(session IAVCaptureSession)
}

// SetSessionControlsDidBecomeActive sets the handler for the SessionControlsDidBecomeActive delegate method.
//
// Tells the delegate when a capture session’s controls become active and available for interaction.
func (d *CaptureSessionControlsDelegate) SetSessionControlsDidBecomeActive(f func(session IAVCaptureSession)) {
	d._SessionControlsDidBecomeActive = f
}

// SetSessionControlsDidBecomeInactive sets the handler for the SessionControlsDidBecomeInactive delegate method.
//
// Tells the delegate when a capture session’s controls become inactive and unavailable for interaction.
func (d *CaptureSessionControlsDelegate) SetSessionControlsDidBecomeInactive(f func(session IAVCaptureSession)) {
	d._SessionControlsDidBecomeInactive = f
}

// SetSessionControlsWillEnterFullscreenAppearance sets the handler for the SessionControlsWillEnterFullscreenAppearance delegate method.
//
// Tells the delegate when a capture session’s controls are about to enter a fullscreen appearance.
func (d *CaptureSessionControlsDelegate) SetSessionControlsWillEnterFullscreenAppearance(f func(session IAVCaptureSession)) {
	d._SessionControlsWillEnterFullscreenAppearance = f
}

// SetSessionControlsWillExitFullscreenAppearance sets the handler for the SessionControlsWillExitFullscreenAppearance delegate method.
//
// Tells the delegate when a capture session’s controls are about to exit a fullscreen appearance.
func (d *CaptureSessionControlsDelegate) SetSessionControlsWillExitFullscreenAppearance(f func(session IAVCaptureSession)) {
	d._SessionControlsWillExitFullscreenAppearance = f
}

// SessionControlsDidBecomeActive implements the PCaptureSessionControlsDelegate interface.
func (d *CaptureSessionControlsDelegate) SessionControlsDidBecomeActive(session IAVCaptureSession) {
	if d._SessionControlsDidBecomeActive != nil {
		d._SessionControlsDidBecomeActive(session)
	}
}

// HasSessionControlsDidBecomeActive returns true if a handler for SessionControlsDidBecomeActive has been set.
func (d *CaptureSessionControlsDelegate) HasSessionControlsDidBecomeActive() bool {
	return d._SessionControlsDidBecomeActive != nil
}

// SessionControlsDidBecomeInactive implements the PCaptureSessionControlsDelegate interface.
func (d *CaptureSessionControlsDelegate) SessionControlsDidBecomeInactive(session IAVCaptureSession) {
	if d._SessionControlsDidBecomeInactive != nil {
		d._SessionControlsDidBecomeInactive(session)
	}
}

// HasSessionControlsDidBecomeInactive returns true if a handler for SessionControlsDidBecomeInactive has been set.
func (d *CaptureSessionControlsDelegate) HasSessionControlsDidBecomeInactive() bool {
	return d._SessionControlsDidBecomeInactive != nil
}

// SessionControlsWillEnterFullscreenAppearance implements the PCaptureSessionControlsDelegate interface.
func (d *CaptureSessionControlsDelegate) SessionControlsWillEnterFullscreenAppearance(session IAVCaptureSession) {
	if d._SessionControlsWillEnterFullscreenAppearance != nil {
		d._SessionControlsWillEnterFullscreenAppearance(session)
	}
}

// HasSessionControlsWillEnterFullscreenAppearance returns true if a handler for SessionControlsWillEnterFullscreenAppearance has been set.
func (d *CaptureSessionControlsDelegate) HasSessionControlsWillEnterFullscreenAppearance() bool {
	return d._SessionControlsWillEnterFullscreenAppearance != nil
}

// SessionControlsWillExitFullscreenAppearance implements the PCaptureSessionControlsDelegate interface.
func (d *CaptureSessionControlsDelegate) SessionControlsWillExitFullscreenAppearance(session IAVCaptureSession) {
	if d._SessionControlsWillExitFullscreenAppearance != nil {
		d._SessionControlsWillExitFullscreenAppearance(session)
	}
}

// HasSessionControlsWillExitFullscreenAppearance returns true if a handler for SessionControlsWillExitFullscreenAppearance has been set.
func (d *CaptureSessionControlsDelegate) HasSessionControlsWillExitFullscreenAppearance() bool {
	return d._SessionControlsWillExitFullscreenAppearance != nil
}
