//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CapturePhotoOutput


// Tells the photo capture output to prepare resources for future capture requests with the specified settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/setPreparedPhotoSettingsArray(_:completionHandler:)
func (c_ CapturePhotoOutput) SetPreparedPhotoSettingsArrayCompletionHandler(preparedPhotoSettingsArray []CapturePhotoSettings, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreparedPhotoSettingsArray:completionHandler:"), preparedPhotoSettingsArray, completionHandler)
}

// Returns the list of Bayer RAW pixel formats supported for photo data in the specified file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/supportedRawPhotoPixelFormatTypesForFileType:
func (c_ CapturePhotoOutput) SupportedRawPhotoPixelFormatTypesForFileType(fileType FileType) []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("supportedRawPhotoPixelFormatTypesForFileType:"), fileType)
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/supportedRawPhotoCodecTypes(forRawPhotoPixelFormatType:fileType:)
func (c_ CapturePhotoOutput) SupportedRawPhotoCodecTypesForRawPhotoPixelFormatTypeFileType(pixelFormatType uint32 /* not a class type */, fileType FileType) []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("supportedRawPhotoCodecTypesForRawPhotoPixelFormatType:fileType:"), pixelFormatType, fileType)
	return rv
}

// iOS-only properties

// The pixel formats the capture output supports for RAW photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableRawPhotoPixelFormatTypes-5fatm
func (c_ CapturePhotoOutput) AvailableRawPhotoPixelFormatTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("availableRawPhotoPixelFormatTypes"))
	return rv
}

// An array of video codecs currently available for Live Photo movie captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableLivePhotoVideoCodecTypes
func (c_ CapturePhotoOutput) AvailableLivePhotoVideoCodecTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableLivePhotoVideoCodecTypes"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableRawPhotoCodecTypes
func (c_ CapturePhotoOutput) AvailableRawPhotoCodecTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableRawPhotoCodecTypes"))
	return rv
}

// The list of file types currently supported for RAW format capture and output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableRawPhotoFileTypes
func (c_ CapturePhotoOutput) AvailableRawPhotoFileTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableRawPhotoFileTypes"))
	return rv
}

// An array of semantic segmentation matte types that may be captured and delivered along with the primary photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableSemanticSegmentationMatteTypes
func (c_ CapturePhotoOutput) AvailableSemanticSegmentationMatteTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableSemanticSegmentationMatteTypes"))
	return rv
}

// The semantic segmentation matte types that the photo render pipeline delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/enabledSemanticSegmentationMatteTypes
func (c_ CapturePhotoOutput) EnabledSemanticSegmentationMatteTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("enabledSemanticSegmentationMatteTypes"))
	return rv
}
func (c_ CapturePhotoOutput) SetEnabledSemanticSegmentationMatteTypes(value []string) {
	c_.ID.Send(objc.RegisterName("setEnabledSemanticSegmentationMatteTypes:"), value)
}

// A Boolean value that indicates whether you’ve configured the photo output to deliver Apple ProRAW formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAppleProRAWEnabled
func (c_ CapturePhotoOutput) AppleProRAWEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("appleProRAWEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetAppleProRAWEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAppleProRAWEnabled:"), value)
}

// A Boolean value that indicates whether the current device and configuration supports Apple ProRAW pixel formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAppleProRAWSupported
func (c_ CapturePhotoOutput) AppleProRAWSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("appleProRAWSupported"))
	return rv
}

// A Boolean value that indicates the enabled state of automatic deferred photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAutoDeferredPhotoDeliveryEnabled
func (c_ CapturePhotoOutput) AutoDeferredPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoDeferredPhotoDeliveryEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetAutoDeferredPhotoDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAutoDeferredPhotoDeliveryEnabled:"), value)
}

// A Boolean value that indicates whether the photo output supports deferred photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAutoDeferredPhotoDeliverySupported
func (c_ CapturePhotoOutput) AutoDeferredPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoDeferredPhotoDeliverySupported"))
	return rv
}

// A Boolean value indicating whether the capture output supports automatic red-eye reduction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAutoRedEyeReductionSupported
func (c_ CapturePhotoOutput) AutoRedEyeReductionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoRedEyeReductionSupported"))
	return rv
}

// A Boolean value indicating whether the capture output currently supports delivery of camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isCameraCalibrationDataDeliverySupported
func (c_ CapturePhotoOutput) CameraCalibrationDataDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraCalibrationDataDeliverySupported"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isCameraSensorOrientationCompensationEnabled
func (c_ CapturePhotoOutput) CameraSensorOrientationCompensationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraSensorOrientationCompensationEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetCameraSensorOrientationCompensationEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setCameraSensorOrientationCompensationEnabled:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isCameraSensorOrientationCompensationSupported
func (c_ CapturePhotoOutput) CameraSensorOrientationCompensationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraSensorOrientationCompensationSupported"))
	return rv
}

// A Boolean value that indicates whether the photo render pipeline can perform content-aware distortion correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isContentAwareDistortionCorrectionEnabled
func (c_ CapturePhotoOutput) ContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("contentAwareDistortionCorrectionEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetContentAwareDistortionCorrectionEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setContentAwareDistortionCorrectionEnabled:"), value)
}

// A Boolean value that indicates whether the session’s current configuration supports content-aware distortion correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isContentAwareDistortionCorrectionSupported
func (c_ CapturePhotoOutput) ContentAwareDistortionCorrectionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("contentAwareDistortionCorrectionSupported"))
	return rv
}

// A Boolean value that specifies whether to configure the capture pipeline for depth data capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDepthDataDeliveryEnabled
func (c_ CapturePhotoOutput) DepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("depthDataDeliveryEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetDepthDataDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setDepthDataDeliveryEnabled:"), value)
}

// A Boolean value indicating whether the capture output currently supports depth data capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDepthDataDeliverySupported
func (c_ CapturePhotoOutput) DepthDataDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("depthDataDeliverySupported"))
	return rv
}

// A Boolean value that specifies whether to configure the capture pipeline for simultaneous photo capture with both cameras on a dual-camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDualCameraDualPhotoDeliveryEnabled
func (c_ CapturePhotoOutput) DualCameraDualPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("dualCameraDualPhotoDeliveryEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetDualCameraDualPhotoDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setDualCameraDualPhotoDeliveryEnabled:"), value)
}

// A Boolean value indicating whether the capture output currently supports simultaneous photo capture with both cameras on a dual-camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDualCameraDualPhotoDeliverySupported
func (c_ CapturePhotoOutput) DualCameraDualPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("dualCameraDualPhotoDeliverySupported"))
	return rv
}

// A Boolean value indicating whether the capture output currently supports automatically combining image data on a dual camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDualCameraFusionSupported
func (c_ CapturePhotoOutput) DualCameraFusionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("dualCameraFusionSupported"))
	return rv
}

// A Boolean value indicating whether the scene currently being previewed by the camera warrants use of the flash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isFlashScene
func (c_ CapturePhotoOutput) IsFlashScene() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFlashScene"))
	return rv
}

// A Boolean value indicating whether the capture output currently supports lens stabilization during bracketed image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLensStabilizationDuringBracketedCaptureSupported
func (c_ CapturePhotoOutput) LensStabilizationDuringBracketedCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lensStabilizationDuringBracketedCaptureSupported"))
	return rv
}

// A Boolean value that indicates whether to automatically trim Live Photo movie captures to avoid excessive movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLivePhotoAutoTrimmingEnabled
func (c_ CapturePhotoOutput) LivePhotoAutoTrimmingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("livePhotoAutoTrimmingEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetLivePhotoAutoTrimmingEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setLivePhotoAutoTrimmingEnabled:"), value)
}

// A Boolean value that indicates whether to configure the capture pipeline for Live Photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLivePhotoCaptureEnabled
func (c_ CapturePhotoOutput) LivePhotoCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("livePhotoCaptureEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetLivePhotoCaptureEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setLivePhotoCaptureEnabled:"), value)
}

// A Boolean value that indicates whether the capture output currently supports Live Photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLivePhotoCaptureSupported
func (c_ CapturePhotoOutput) LivePhotoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("livePhotoCaptureSupported"))
	return rv
}

// A Boolean value that indicates whether Live Photo capture is currently in a suspended state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLivePhotoCaptureSuspended
func (c_ CapturePhotoOutput) LivePhotoCaptureSuspended() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("livePhotoCaptureSuspended"))
	return rv
}
func (c_ CapturePhotoOutput) SetLivePhotoCaptureSuspended(value bool) {
	c_.ID.Send(objc.RegisterName("setLivePhotoCaptureSuspended:"), value)
}

// A Boolean value indicating whether the capture output generates a portrait effects matte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isPortraitEffectsMatteDeliveryEnabled
func (c_ CapturePhotoOutput) PortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectsMatteDeliveryEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetPortraitEffectsMatteDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setPortraitEffectsMatteDeliveryEnabled:"), value)
}

// A Boolean value indicating whether the capture output currently supports delivery of a portrait effects matte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isPortraitEffectsMatteDeliverySupported
func (c_ CapturePhotoOutput) PortraitEffectsMatteDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectsMatteDeliverySupported"))
	return rv
}

// A Boolean value indicating whether the scene currently being previewed by the camera warrants image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isStillImageStabilizationScene
func (c_ CapturePhotoOutput) IsStillImageStabilizationScene() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStillImageStabilizationScene"))
	return rv
}

// A Boolean value indicating whether the capture output currently supports automatic stabilization for still image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isStillImageStabilizationSupported
func (c_ CapturePhotoOutput) StillImageStabilizationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("stillImageStabilizationSupported"))
	return rv
}

// A Boolean value that indicates whether the photo output delivers photos from constituent cameras of a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isVirtualDeviceConstituentPhotoDeliveryEnabled
func (c_ CapturePhotoOutput) VirtualDeviceConstituentPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("virtualDeviceConstituentPhotoDeliveryEnabled"))
	return rv
}
func (c_ CapturePhotoOutput) SetVirtualDeviceConstituentPhotoDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setVirtualDeviceConstituentPhotoDeliveryEnabled:"), value)
}

// A Boolean value that indicates whether the photo output configuration supports delivery of photos from constituent cameras of a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isVirtualDeviceConstituentPhotoDeliverySupported
func (c_ CapturePhotoOutput) VirtualDeviceConstituentPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("virtualDeviceConstituentPhotoDeliverySupported"))
	return rv
}

// A Boolean value that indicates whether the device supports virtual device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isVirtualDeviceFusionSupported
func (c_ CapturePhotoOutput) VirtualDeviceFusionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("virtualDeviceFusionSupported"))
	return rv
}

// The maximum number of images that the photo capture output can support in a single bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/maxBracketedCapturePhotoCount
func (c_ CapturePhotoOutput) MaxBracketedCapturePhotoCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxBracketedCapturePhotoCount"))
	return rv
}

// A photo settings object that controls how the photo output detects and handles automatic flash and stabilization modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/photoSettingsForSceneMonitoring
func (c_ CapturePhotoOutput) PhotoSettingsForSceneMonitoring() IAVCapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](c_.ID, objc.Sel("photoSettingsForSceneMonitoring"))
	return rv
}
func (c_ CapturePhotoOutput) SetPhotoSettingsForSceneMonitoring(value IAVCapturePhotoSettings) {
	c_.ID.Send(objc.RegisterName("setPhotoSettingsForSceneMonitoring:"), value)
}

// An array of photo settings for which the photo output has prepared capture resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/preparedPhotoSettingsArray
func (c_ CapturePhotoOutput) PreparedPhotoSettingsArray() []CapturePhotoSettings {
	rv := objc.Send[[]CapturePhotoSettings](c_.ID, objc.Sel("preparedPhotoSettingsArray"))
	return rv
}




