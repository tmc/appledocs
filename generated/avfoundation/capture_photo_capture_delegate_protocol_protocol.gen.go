// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	CaptureOutputDidFinishCaptureForResolvedSettingsError(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ foundation.foundation.INSError)
	HasCaptureOutputDidFinishCaptureForResolvedSettingsError() bool
	CaptureOutputDidFinishCapturingDeferredPhotoProxyError(output IAVCapturePhotoOutput, deferredPhotoProxy IAVCaptureDeferredPhotoProxy, error_ foundation.foundation.INSError)
	HasCaptureOutputDidFinishCapturingDeferredPhotoProxyError() bool
	CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError(output IAVCapturePhotoOutput, outputFileURL foundation.foundation.INSURL, duration objectivec.IObject, photoDisplayTime objectivec.IObject, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ foundation.foundation.INSError)
	HasCaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError() bool
	CaptureOutputDidFinishProcessingPhotoError(output IAVCapturePhotoOutput, photo IAVCapturePhoto, error_ foundation.foundation.INSError)
	HasCaptureOutputDidFinishProcessingPhotoError() bool
	CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(output IAVCapturePhotoOutput, photoSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ foundation.foundation.INSError)
	HasCaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError() bool
	CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError(output IAVCapturePhotoOutput, rawSampleBuffer SampleBufferRef /* not a class type */, previewPhotoSampleBuffer SampleBufferRef /* not a class type */, resolvedSettings IAVCaptureResolvedPhotoSettings, bracketSettings IAVCaptureBracketedStillImageSettings, error_ foundation.foundation.INSError)
	HasCaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError() bool
	CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings(output IAVCapturePhotoOutput, outputFileURL foundation.foundation.INSURL, resolvedSettings IAVCaptureResolvedPhotoSettings)
	HasCaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings() bool
	CaptureOutputWillBeginCaptureForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)
	HasCaptureOutputWillBeginCaptureForResolvedSettings() bool
	CaptureOutputWillCapturePhotoForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings)
	HasCaptureOutputWillCapturePhotoForResolvedSettings() bool
}
