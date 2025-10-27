//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureDeviceFormat


// A Boolean value that indicates whether the format supports a given video stabilization mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isVideoStabilizationModeSupported(_:)
func (c_ CaptureDeviceFormat) IsVideoStabilizationModeSupported(videoStabilizationMode CaptureVideoStabilizationMode) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoStabilizationModeSupported:"), videoStabilizationMode)
	return rv
}

// Indicates the horizontal field of view for an aspect ratio, either uncorrected or corrected for geometric distortion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFieldOfView(for:geometricDistortionCorrected:)
func (c_ CaptureDeviceFormat) VideoFieldOfViewForAspectRatioGeometricDistortionCorrected(aspectRatio CaptureAspectRatio, geometricDistortionCorrected bool) float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("videoFieldOfViewForAspectRatio:geometricDistortionCorrected:"), aspectRatio, geometricDistortionCorrected)
	return rv
}

// iOS-only properties

// The zoom factors that a format supports for depth data delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/supportedVideoZoomFactorsForDepthDataDelivery
func (c_ CaptureDeviceFormat) SupportedVideoZoomFactorsForDepthDataDelivery() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("supportedVideoZoomFactorsForDepthDataDelivery"))
	return rv
}

// A horizontal field of view for the format after correction for geometric distortion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/geometricDistortionCorrectedVideoFieldOfView
func (c_ CaptureDeviceFormat) GeometricDistortionCorrectedVideoFieldOfView() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("geometricDistortionCorrectedVideoFieldOfView"))
	return rv
}

// The highest resolution still image the system can produce for this format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/highResolutionStillImageDimensions
func (c_ CaptureDeviceFormat) HighResolutionStillImageDimensions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("highResolutionStillImageDimensions"))
	return rv
}

// A Boolean value that indicates whether the format supports global tone mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isGlobalToneMappingSupported
func (c_ CaptureDeviceFormat) GlobalToneMappingSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("globalToneMappingSupported"))
	return rv
}

// A Boolean value that indicates whether this format supports the highest photo quality that the platform delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isHighestPhotoQualitySupported
func (c_ CaptureDeviceFormat) HighestPhotoQualitySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highestPhotoQualitySupported"))
	return rv
}

// A Boolean value that indicates whether a multi-camera capture session supports this format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isMultiCamSupported
func (c_ CaptureDeviceFormat) MultiCamSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("multiCamSupported"))
	return rv
}

// A Boolean indicating whether the device supports portrait matte effects in still-image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isPortraitEffectsMatteStillImageDeliverySupported
func (c_ CaptureDeviceFormat) PortraitEffectsMatteStillImageDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectsMatteStillImageDeliverySupported"))
	return rv
}

// Returns if smart framing is supported by the current format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isSmartFramingSupported
func (c_ CaptureDeviceFormat) SmartFramingSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("smartFramingSupported"))
	return rv
}

// A Boolean value that indicates whether the format produces video data in a binned format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isVideoBinned
func (c_ CaptureDeviceFormat) VideoBinned() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoBinned"))
	return rv
}

// A Boolean value that indicates whether the format supports high dynamic range streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isVideoHDRSupported
func (c_ CaptureDeviceFormat) VideoHDRSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoHDRSupported"))
	return rv
}

// A Boolean value that indicates whether the format supports video stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isVideoStabilizationSupported
func (c_ CaptureDeviceFormat) VideoStabilizationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoStabilizationSupported"))
	return rv
}

// A time value that indicates the maximum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/maxExposureDuration
func (c_ CaptureDeviceFormat) MaxExposureDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("maxExposureDuration"))
	return rv
}

// A floating-point number that indicates the maximum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/maxISO
func (c_ CaptureDeviceFormat) MaxISO() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("maxISO"))
	return rv
}

// A time value that indicates the minimum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/minExposureDuration
func (c_ CaptureDeviceFormat) MinExposureDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("minExposureDuration"))
	return rv
}

// A floating-point number that indicates the minimum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/minISO
func (c_ CaptureDeviceFormat) MinISO() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("minISO"))
	return rv
}

// The list of data formats compatible with this video format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/supportedDepthDataFormats
func (c_ CaptureDeviceFormat) SupportedDepthDataFormats() []CaptureDeviceFormat {
	rv := objc.Send[[]CaptureDeviceFormat](c_.ID, objc.Sel("supportedDepthDataFormats"))
	return rv
}

// Indicates the supported aspect ratios for the device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/supportedDynamicAspectRatios
func (c_ CaptureDeviceFormat) SupportedDynamicAspectRatios() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("supportedDynamicAspectRatios"))
	return rv
}

// The list of capture output subclasses not allowed for capture with this format, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/unsupportedCaptureOutputClasses
func (c_ CaptureDeviceFormat) UnsupportedCaptureOutputClasses() []objc.Class {
	rv := objc.Send[[]objc.Class](c_.ID, objc.Sel("unsupportedCaptureOutputClasses"))
	return rv
}

// Indicates the format’s horizontal field of view in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFieldOfView
func (c_ CaptureDeviceFormat) VideoFieldOfView() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("videoFieldOfView"))
	return rv
}

// A maximum zoom factor the format allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMaxZoomFactor
func (c_ CaptureDeviceFormat) VideoMaxZoomFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxZoomFactor"))
	return rv
}

// A maximum zoom factor the device supports when configured for depth data delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMaxZoomFactorForDepthDataDelivery
func (c_ CaptureDeviceFormat) VideoMaxZoomFactorForDepthDataDelivery() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxZoomFactorForDepthDataDelivery"))
	return rv
}

// A minimum zoom factor the device supports when configured for depth data delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMinZoomFactorForDepthDataDelivery
func (c_ CaptureDeviceFormat) VideoMinZoomFactorForDepthDataDelivery() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMinZoomFactorForDepthDataDelivery"))
	return rv
}

// A threshold at which the system upscales pixel data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoZoomFactorUpscaleThreshold
func (c_ CaptureDeviceFormat) VideoZoomFactorUpscaleThreshold() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoZoomFactorUpscaleThreshold"))
	return rv
}





