// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corevideo"

	"github.com/tmc/appledocs/generated/foundation"
)

// PCapturePhotoCaptureDelegate is the AVCapturePhotoCaptureDelegate protocol interface.
//
// Methods for monitoring progress and receiving results from a photo capture output.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.15+
//   - tvOS 17.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCapturePhotoCaptureDelegate
type PCapturePhotoCaptureDelegate interface {
	// Optional methods
	CaptureOutputDidCapturePhotoForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)
	HasCaptureOutputDidCapturePhotoForResolvedSettings() bool
	CaptureOutputDidFinishCaptureForResolvedSettingsError(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ Error)
	HasCaptureOutputDidFinishCaptureForResolvedSettingsError() bool
	CaptureOutputDidFinishCapturingDeferredPhotoProxyError(output IAVCapturePhotoOutput, deferredPhotoProxy IAVCaptureDeferredPhotoProxy, error_ Error)
	HasCaptureOutputDidFinishCapturingDeferredPhotoProxyError() bool
	CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError(output IAVCapturePhotoOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, duration objc.IObject /* cross-framework: Time */, photoDisplayTime objc.IObject /* cross-framework: Time */, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ Error)
	HasCaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError() bool
	CaptureOutputDidFinishProcessingPhotoError(output IAVCapturePhotoOutput, photo IAVCapturePhoto, error_ Error)
	HasCaptureOutputDidFinishProcessingPhotoError() bool
	CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(output IAVCapturePhotoOutput, photoSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ Error)
	HasCaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError() bool
	CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(output IAVCapturePhotoOutput, rawSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ Error)
	HasCaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError() bool
	CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings(output IAVCapturePhotoOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, resolvedSettings IAVCaptureResolvedPhotoSettings)
	HasCaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings() bool
	CaptureOutputWillBeginCaptureForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)
	HasCaptureOutputWillBeginCaptureForResolvedSettings() bool
	CaptureOutputWillCapturePhotoForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)
	HasCaptureOutputWillCapturePhotoForResolvedSettings() bool
}

// CapturePhotoCaptureDelegate is a delegate implementation builder for the PCapturePhotoCaptureDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CapturePhotoCaptureDelegate struct {
	_CaptureOutputDidCapturePhotoForResolvedSettings func(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)
	_CaptureOutputDidFinishCaptureForResolvedSettingsError func(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ Error)
	_CaptureOutputDidFinishCapturingDeferredPhotoProxyError func(output IAVCapturePhotoOutput, deferredPhotoProxy IAVCaptureDeferredPhotoProxy, error_ Error)
	_CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError func(output IAVCapturePhotoOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, duration objc.IObject /* cross-framework: Time */, photoDisplayTime objc.IObject /* cross-framework: Time */, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ Error)
	_CaptureOutputDidFinishProcessingPhotoError func(output IAVCapturePhotoOutput, photo IAVCapturePhoto, error_ Error)
	_CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError func(output IAVCapturePhotoOutput, photoSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ Error)
	_CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError func(output IAVCapturePhotoOutput, rawSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ Error)
	_CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings func(output IAVCapturePhotoOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, resolvedSettings IAVCaptureResolvedPhotoSettings)
	_CaptureOutputWillBeginCaptureForResolvedSettings func(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)
	_CaptureOutputWillCapturePhotoForResolvedSettings func(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)
}

// SetCaptureOutputDidCapturePhotoForResolvedSettings sets the handler for the CaptureOutputDidCapturePhotoForResolvedSettings delegate method.
//
// Notifies the delegate that the photo has been taken.
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputDidCapturePhotoForResolvedSettings(f func(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)) {
	d._CaptureOutputDidCapturePhotoForResolvedSettings = f
}

// SetCaptureOutputDidFinishCaptureForResolvedSettingsError sets the handler for the CaptureOutputDidFinishCaptureForResolvedSettingsError delegate method.
//
// Notifies the delegate that the capture process is complete.
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputDidFinishCaptureForResolvedSettingsError(f func(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ Error)) {
	d._CaptureOutputDidFinishCaptureForResolvedSettingsError = f
}

// SetCaptureOutputDidFinishCapturingDeferredPhotoProxyError sets the handler for the CaptureOutputDidFinishCapturingDeferredPhotoProxyError delegate method.
//
// Tells the delegate when the system finishes capturing the photo proxy.
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputDidFinishCapturingDeferredPhotoProxyError(f func(output IAVCapturePhotoOutput, deferredPhotoProxy IAVCaptureDeferredPhotoProxy, error_ Error)) {
	d._CaptureOutputDidFinishCapturingDeferredPhotoProxyError = f
}

// SetCaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError sets the handler for the CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError delegate method.
//
// Provides the delegate the movie file URL resulting from a Live Photo capture.
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError(f func(output IAVCapturePhotoOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, duration objc.IObject /* cross-framework: Time */, photoDisplayTime objc.IObject /* cross-framework: Time */, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ Error)) {
	d._CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError = f
}

// SetCaptureOutputDidFinishProcessingPhotoError sets the handler for the CaptureOutputDidFinishProcessingPhotoError delegate method.
//
// Provides the delegate with the captured image and associated metadata resulting from a photo capture.
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputDidFinishProcessingPhotoError(f func(output IAVCapturePhotoOutput, photo IAVCapturePhoto, error_ Error)) {
	d._CaptureOutputDidFinishProcessingPhotoError = f
}

// SetCaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError sets the handler for the CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError delegate method.
//
// Provides the delegate a captured image in a processed format (such as JPEG).
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(f func(output IAVCapturePhotoOutput, photoSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ Error)) {
	d._CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError = f
}

// SetCaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError sets the handler for the CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError delegate method.
//
// Provides the delegate a captured image in RAW format.
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(f func(output IAVCapturePhotoOutput, rawSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ Error)) {
	d._CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError = f
}

// SetCaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings sets the handler for the CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings delegate method.
//
// Notifies the delegate that the movie content of a Live Photo has finished recording.
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings(f func(output IAVCapturePhotoOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, resolvedSettings IAVCaptureResolvedPhotoSettings)) {
	d._CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings = f
}

// SetCaptureOutputWillBeginCaptureForResolvedSettings sets the handler for the CaptureOutputWillBeginCaptureForResolvedSettings delegate method.
//
// Notifies the delegate that the capture output has resolved settings and will soon begin its capture process.
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputWillBeginCaptureForResolvedSettings(f func(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)) {
	d._CaptureOutputWillBeginCaptureForResolvedSettings = f
}

// SetCaptureOutputWillCapturePhotoForResolvedSettings sets the handler for the CaptureOutputWillCapturePhotoForResolvedSettings delegate method.
//
// Notifies the delegate that photo capture is about to occur.
func (d *CapturePhotoCaptureDelegate) SetCaptureOutputWillCapturePhotoForResolvedSettings(f func(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)) {
	d._CaptureOutputWillCapturePhotoForResolvedSettings = f
}

// CaptureOutputDidCapturePhotoForResolvedSettings implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputDidCapturePhotoForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings) {
	if d._CaptureOutputDidCapturePhotoForResolvedSettings != nil {
		d._CaptureOutputDidCapturePhotoForResolvedSettings(output, resolvedSettings)
	}
}

// HasCaptureOutputDidCapturePhotoForResolvedSettings returns true if a handler for CaptureOutputDidCapturePhotoForResolvedSettings has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputDidCapturePhotoForResolvedSettings() bool {
	return d._CaptureOutputDidCapturePhotoForResolvedSettings != nil
}

// CaptureOutputDidFinishCaptureForResolvedSettingsError implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputDidFinishCaptureForResolvedSettingsError(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ Error) {
	if d._CaptureOutputDidFinishCaptureForResolvedSettingsError != nil {
		d._CaptureOutputDidFinishCaptureForResolvedSettingsError(output, resolvedSettings, error_)
	}
}

// HasCaptureOutputDidFinishCaptureForResolvedSettingsError returns true if a handler for CaptureOutputDidFinishCaptureForResolvedSettingsError has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputDidFinishCaptureForResolvedSettingsError() bool {
	return d._CaptureOutputDidFinishCaptureForResolvedSettingsError != nil
}

// CaptureOutputDidFinishCapturingDeferredPhotoProxyError implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputDidFinishCapturingDeferredPhotoProxyError(output IAVCapturePhotoOutput, deferredPhotoProxy IAVCaptureDeferredPhotoProxy, error_ Error) {
	if d._CaptureOutputDidFinishCapturingDeferredPhotoProxyError != nil {
		d._CaptureOutputDidFinishCapturingDeferredPhotoProxyError(output, deferredPhotoProxy, error_)
	}
}

// HasCaptureOutputDidFinishCapturingDeferredPhotoProxyError returns true if a handler for CaptureOutputDidFinishCapturingDeferredPhotoProxyError has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputDidFinishCapturingDeferredPhotoProxyError() bool {
	return d._CaptureOutputDidFinishCapturingDeferredPhotoProxyError != nil
}

// CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError(output IAVCapturePhotoOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, duration objc.IObject /* cross-framework: Time */, photoDisplayTime objc.IObject /* cross-framework: Time */, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ Error) {
	if d._CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError != nil {
		d._CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError(output, outputFileURL, duration, photoDisplayTime, resolvedSettings, error_)
	}
}

// HasCaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError returns true if a handler for CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError() bool {
	return d._CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError != nil
}

// CaptureOutputDidFinishProcessingPhotoError implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputDidFinishProcessingPhotoError(output IAVCapturePhotoOutput, photo IAVCapturePhoto, error_ Error) {
	if d._CaptureOutputDidFinishProcessingPhotoError != nil {
		d._CaptureOutputDidFinishProcessingPhotoError(output, photo, error_)
	}
}

// HasCaptureOutputDidFinishProcessingPhotoError returns true if a handler for CaptureOutputDidFinishProcessingPhotoError has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputDidFinishProcessingPhotoError() bool {
	return d._CaptureOutputDidFinishProcessingPhotoError != nil
}

// CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(output IAVCapturePhotoOutput, photoSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ Error) {
	if d._CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError != nil {
		d._CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(output, photoSampleBuffer, previewPhotoSampleBuffer, resolvedSettings, bracketSettings, error_)
	}
}

// HasCaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError returns true if a handler for CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError() bool {
	return d._CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError != nil
}

// CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(output IAVCapturePhotoOutput, rawSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ Error) {
	if d._CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError != nil {
		d._CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(output, rawSampleBuffer, previewPhotoSampleBuffer, resolvedSettings, bracketSettings, error_)
	}
}

// HasCaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError returns true if a handler for CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError() bool {
	return d._CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError != nil
}

// CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings(output IAVCapturePhotoOutput, outputFileURL objc.IObject /* cross-framework: NSURL */, resolvedSettings IAVCaptureResolvedPhotoSettings) {
	if d._CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings != nil {
		d._CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings(output, outputFileURL, resolvedSettings)
	}
}

// HasCaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings returns true if a handler for CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings() bool {
	return d._CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings != nil
}

// CaptureOutputWillBeginCaptureForResolvedSettings implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputWillBeginCaptureForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings) {
	if d._CaptureOutputWillBeginCaptureForResolvedSettings != nil {
		d._CaptureOutputWillBeginCaptureForResolvedSettings(output, resolvedSettings)
	}
}

// HasCaptureOutputWillBeginCaptureForResolvedSettings returns true if a handler for CaptureOutputWillBeginCaptureForResolvedSettings has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputWillBeginCaptureForResolvedSettings() bool {
	return d._CaptureOutputWillBeginCaptureForResolvedSettings != nil
}

// CaptureOutputWillCapturePhotoForResolvedSettings implements the PCapturePhotoCaptureDelegate interface.
func (d *CapturePhotoCaptureDelegate) CaptureOutputWillCapturePhotoForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings) {
	if d._CaptureOutputWillCapturePhotoForResolvedSettings != nil {
		d._CaptureOutputWillCapturePhotoForResolvedSettings(output, resolvedSettings)
	}
}

// HasCaptureOutputWillCapturePhotoForResolvedSettings returns true if a handler for CaptureOutputWillCapturePhotoForResolvedSettings has been set.
func (d *CapturePhotoCaptureDelegate) HasCaptureOutputWillCapturePhotoForResolvedSettings() bool {
	return d._CaptureOutputWillCapturePhotoForResolvedSettings != nil
}
