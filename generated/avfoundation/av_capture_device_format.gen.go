// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureDeviceFormat] class.
var (
	CaptureDeviceFormatClass     _CaptureDeviceFormatClass
	CaptureDeviceFormatClassOnce sync.Once
)

func getCaptureDeviceFormatClass() _CaptureDeviceFormatClass {
	CaptureDeviceFormatClassOnce.Do(func() {
		CaptureDeviceFormatClass = _CaptureDeviceFormatClass{objc.GetClass("AVCaptureDeviceFormat")}
	})
	return CaptureDeviceFormatClass
}

type _CaptureDeviceFormatClass struct {
	class objc.Class
}

// An interface definition for the [CaptureDeviceFormat] class.
type ICaptureDeviceFormat interface {
	objectivec.IObject
	// properties:
	AutoFocusSystem() unsafe.Pointer
	SetAutoFocusSystem(value unsafe.Pointer)
	DefaultSimulatedAperture() float32 /* primitive/slice/pointer. */
	SetDefaultSimulatedAperture(value float32 /* primitive/slice/pointer. */)
	FormatDescription() FormatDescription /* not a class type */
	SetFormatDescription(value FormatDescription /* not a class type */)
	GeometricDistortionCorrectedVideoFieldOfView() float32 /* primitive/slice/pointer. */
	SetGeometricDistortionCorrectedVideoFieldOfView(value float32 /* primitive/slice/pointer. */)
	IsAutoVideoFrameRateSupported() bool /* primitive/slice/pointer. */
	SetIsAutoVideoFrameRateSupported(value bool /* primitive/slice/pointer. */)
	IsBackgroundReplacementSupported() bool /* primitive/slice/pointer. */
	SetIsBackgroundReplacementSupported(value bool /* primitive/slice/pointer. */)
	IsCameraLensSmudgeDetectionSupported() bool /* primitive/slice/pointer. */
	SetIsCameraLensSmudgeDetectionSupported(value bool /* primitive/slice/pointer. */)
	IsCenterStageSupported() bool /* primitive/slice/pointer. */
	SetIsCenterStageSupported(value bool /* primitive/slice/pointer. */)
	IsCinematicVideoCaptureSupported() bool /* primitive/slice/pointer. */
	SetIsCinematicVideoCaptureSupported(value bool /* primitive/slice/pointer. */)
	IsGlobalToneMappingSupported() bool /* primitive/slice/pointer. */
	SetIsGlobalToneMappingSupported(value bool /* primitive/slice/pointer. */)
	IsHighPhotoQualitySupported() bool /* primitive/slice/pointer. */
	SetIsHighPhotoQualitySupported(value bool /* primitive/slice/pointer. */)
	IsHighestPhotoQualitySupported() bool /* primitive/slice/pointer. */
	SetIsHighestPhotoQualitySupported(value bool /* primitive/slice/pointer. */)
	IsMultiCamSupported() bool /* primitive/slice/pointer. */
	SetIsMultiCamSupported(value bool /* primitive/slice/pointer. */)
	IsPortraitEffectSupported() bool /* primitive/slice/pointer. */
	SetIsPortraitEffectSupported(value bool /* primitive/slice/pointer. */)
	IsPortraitEffectsMatteStillImageDeliverySupported() bool /* primitive/slice/pointer. */
	SetIsPortraitEffectsMatteStillImageDeliverySupported(value bool /* primitive/slice/pointer. */)
	IsSmartFramingSupported() bool /* primitive/slice/pointer. */
	SetIsSmartFramingSupported(value bool /* primitive/slice/pointer. */)
	IsSpatialVideoCaptureSupported() bool /* primitive/slice/pointer. */
	SetIsSpatialVideoCaptureSupported(value bool /* primitive/slice/pointer. */)
	IsStudioLightSupported() bool /* primitive/slice/pointer. */
	SetIsStudioLightSupported(value bool /* primitive/slice/pointer. */)
	IsVideoBinned() bool /* primitive/slice/pointer. */
	SetIsVideoBinned(value bool /* primitive/slice/pointer. */)
	IsVideoHDRSupported() bool /* primitive/slice/pointer. */
	SetIsVideoHDRSupported(value bool /* primitive/slice/pointer. */)
	MaxExposureDuration() Time /* not a class type */
	SetMaxExposureDuration(value Time /* not a class type */)
	MaxISO() float32 /* primitive/slice/pointer. */
	SetMaxISO(value float32 /* primitive/slice/pointer. */)
	MaxSimulatedAperture() float32 /* primitive/slice/pointer. */
	SetMaxSimulatedAperture(value float32 /* primitive/slice/pointer. */)
	MediaType() MediaType /* not a class type */
	SetMediaType(value MediaType /* not a class type */)
	MinExposureDuration() Time /* not a class type */
	SetMinExposureDuration(value Time /* not a class type */)
	MinISO() float32 /* primitive/slice/pointer. */
	SetMinISO(value float32 /* primitive/slice/pointer. */)
	MinSimulatedAperture() float32 /* primitive/slice/pointer. */
	SetMinSimulatedAperture(value float32 /* primitive/slice/pointer. */)
	ReactionEffectsSupported() bool /* primitive/slice/pointer. */
	SetReactionEffectsSupported(value bool /* primitive/slice/pointer. */)
	SecondaryNativeResolutionZoomFactors() float64 /* primitive/slice/pointer. */
	SetSecondaryNativeResolutionZoomFactors(value float64 /* primitive/slice/pointer. */)
	SupportedColorSpaces() CaptureColorSpace /* not a class type */
	SetSupportedColorSpaces(value CaptureColorSpace /* not a class type */)
	SupportedDepthDataFormats() IAVCaptureDeviceFormat
	SetSupportedDepthDataFormats(value IAVCaptureDeviceFormat)
	SupportedDynamicAspectRatios() unsafe.Pointer
	SetSupportedDynamicAspectRatios(value unsafe.Pointer)
	SupportedMaxPhotoDimensions() VideoDimensions /* not a class type */
	SetSupportedMaxPhotoDimensions(value VideoDimensions /* not a class type */)
	SupportedVideoZoomFactorsForDepthDataDelivery() float64 /* primitive/slice/pointer. */
	SetSupportedVideoZoomFactorsForDepthDataDelivery(value float64 /* primitive/slice/pointer. */)
	SupportedVideoZoomRangesForDepthDataDelivery() float64 /* primitive/slice/pointer. */
	SetSupportedVideoZoomRangesForDepthDataDelivery(value float64 /* primitive/slice/pointer. */)
	SystemRecommendedExposureBiasRange() float32 /* primitive/slice/pointer. */
	SetSystemRecommendedExposureBiasRange(value float32 /* primitive/slice/pointer. */)
	SystemRecommendedVideoZoomRange() float64 /* primitive/slice/pointer. */
	SetSystemRecommendedVideoZoomRange(value float64 /* primitive/slice/pointer. */)
	UnsupportedCaptureOutputClasses() unsafe.Pointer
	SetUnsupportedCaptureOutputClasses(value unsafe.Pointer)
	VideoFieldOfView() float32 /* primitive/slice/pointer. */
	SetVideoFieldOfView(value float32 /* primitive/slice/pointer. */)
	VideoFrameRateRangeForBackgroundReplacement() FrameRateRange /* not a class type */
	SetVideoFrameRateRangeForBackgroundReplacement(value FrameRateRange /* not a class type */)
	VideoFrameRateRangeForCenterStage() FrameRateRange /* not a class type */
	SetVideoFrameRateRangeForCenterStage(value FrameRateRange /* not a class type */)
	VideoFrameRateRangeForCinematicVideo() FrameRateRange /* not a class type */
	SetVideoFrameRateRangeForCinematicVideo(value FrameRateRange /* not a class type */)
	VideoFrameRateRangeForPortraitEffect() FrameRateRange /* not a class type */
	SetVideoFrameRateRangeForPortraitEffect(value FrameRateRange /* not a class type */)
	VideoFrameRateRangeForReactionEffectsInProgress() FrameRateRange /* not a class type */
	SetVideoFrameRateRangeForReactionEffectsInProgress(value FrameRateRange /* not a class type */)
	VideoFrameRateRangeForStudioLight() FrameRateRange /* not a class type */
	SetVideoFrameRateRangeForStudioLight(value FrameRateRange /* not a class type */)
	VideoMaxZoomFactor() float64 /* primitive/slice/pointer. */
	SetVideoMaxZoomFactor(value float64 /* primitive/slice/pointer. */)
	VideoMaxZoomFactorForCenterStage() float64 /* primitive/slice/pointer. */
	SetVideoMaxZoomFactorForCenterStage(value float64 /* primitive/slice/pointer. */)
	VideoMaxZoomFactorForCinematicVideo() float64 /* primitive/slice/pointer. */
	SetVideoMaxZoomFactorForCinematicVideo(value float64 /* primitive/slice/pointer. */)
	VideoMinZoomFactorForCenterStage() float64 /* primitive/slice/pointer. */
	SetVideoMinZoomFactorForCenterStage(value float64 /* primitive/slice/pointer. */)
	VideoMinZoomFactorForCinematicVideo() float64 /* primitive/slice/pointer. */
	SetVideoMinZoomFactorForCinematicVideo(value float64 /* primitive/slice/pointer. */)
	VideoSupportedFrameRateRanges() FrameRateRange /* not a class type */
	SetVideoSupportedFrameRateRanges(value FrameRateRange /* not a class type */)
	VideoZoomFactorUpscaleThreshold() float64 /* primitive/slice/pointer. */
	SetVideoZoomFactorUpscaleThreshold(value float64 /* primitive/slice/pointer. */)
	ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool /* primitive/slice/pointer. */
	SetZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported(value bool /* primitive/slice/pointer. */)
	ActiveDepthDataFormat() IAVCaptureDeviceFormat
	SetActiveDepthDataFormat(value IAVCaptureDeviceFormat)
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	Formats() IAVCaptureDeviceFormat
	SetFormats(value IAVCaptureDeviceFormat)
	VideoZoomFactor() float64 /* primitive/slice/pointer. */
	SetVideoZoomFactor(value float64 /* primitive/slice/pointer. */)
	// methods:
}

// A class that defines media formats and capture settings that capture devices support.
//
// A format object provides information about a media capture format to use with a capture device, such as video frame rates and zoom factors. You can find more information about a capture format using its associated Core Media format description (see ), available using the property.


// A class that defines media formats and capture settings that capture devices support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format
type CaptureDeviceFormat struct {
	objectivec.Object
}

// CaptureDeviceFormatFrom constructs a [CaptureDeviceFormat] from an unsafe.Pointer.
//
// A class that defines media formats and capture settings that capture devices support.
func CaptureDeviceFormatFrom(ptr unsafe.Pointer) CaptureDeviceFormat {
	return CaptureDeviceFormat{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceFormatClass) Alloc() CaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureDeviceFormatClass) New() CaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDeviceFormat) Init() CaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDeviceFormat) Autorelease() CaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDeviceFormat creates a new CaptureDeviceFormat instance.
func NewCaptureDeviceFormat() CaptureDeviceFormat {
	return getCaptureDeviceFormatClass().New()
}



// The auto focus system the format uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/autofocussystem-swift.property
func (c_ CaptureDeviceFormat) AutoFocusSystem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("autoFocusSystem"))
	return rv
}


// The auto focus system the format uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/autofocussystem-swift.property
func (c_ CaptureDeviceFormat) SetAutoFocusSystem(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutoFocusSystem:"), value)
}


// Default shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/defaultsimulatedaperture
func (c_ CaptureDeviceFormat) DefaultSimulatedAperture() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("defaultSimulatedAperture"))
	return rv
}


// Default shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/defaultsimulatedaperture
func (c_ CaptureDeviceFormat) SetDefaultSimulatedAperture(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultSimulatedAperture:"), value)
}


// An object describing the capture format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/formatdescription
func (c_ CaptureDeviceFormat) FormatDescription() FormatDescription /* not a class type */ {
	rv := objc.Send[FormatDescription](c_.ID, objc.Sel("formatDescription"))
	return rv
}


// An object describing the capture format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/formatdescription
func (c_ CaptureDeviceFormat) SetFormatDescription(value FormatDescription /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescription:"), value)
}


// A horizontal field of view for the format after correction for geometric distortion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/geometricdistortioncorrectedvideofieldofview
func (c_ CaptureDeviceFormat) GeometricDistortionCorrectedVideoFieldOfView() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("geometricDistortionCorrectedVideoFieldOfView"))
	return rv
}


// A horizontal field of view for the format after correction for geometric distortion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/geometricdistortioncorrectedvideofieldofview
func (c_ CaptureDeviceFormat) SetGeometricDistortionCorrectedVideoFieldOfView(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGeometricDistortionCorrectedVideoFieldOfView:"), value)
}


// A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isautovideoframeratesupported
func (c_ CaptureDeviceFormat) IsAutoVideoFrameRateSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoVideoFrameRateSupported"))
	return rv
}


// A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isautovideoframeratesupported
func (c_ CaptureDeviceFormat) SetIsAutoVideoFrameRateSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoVideoFrameRateSupported:"), value)
}


// A Boolean value that indicates whether the format supports background replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isbackgroundreplacementsupported
func (c_ CaptureDeviceFormat) IsBackgroundReplacementSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBackgroundReplacementSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports background replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isbackgroundreplacementsupported
func (c_ CaptureDeviceFormat) SetIsBackgroundReplacementSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackgroundReplacementSupported:"), value)
}


// Whether camera lens smudge detection is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscameralenssmudgedetectionsupported
func (c_ CaptureDeviceFormat) IsCameraLensSmudgeDetectionSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraLensSmudgeDetectionSupported"))
	return rv
}


// Whether camera lens smudge detection is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscameralenssmudgedetectionsupported
func (c_ CaptureDeviceFormat) SetIsCameraLensSmudgeDetectionSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraLensSmudgeDetectionSupported:"), value)
}


// A Boolean value that indicates whether the format supports Center Stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscenterstagesupported
func (c_ CaptureDeviceFormat) IsCenterStageSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCenterStageSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports Center Stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscenterstagesupported
func (c_ CaptureDeviceFormat) SetIsCenterStageSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCenterStageSupported:"), value)
}


// Indicates whether the format supports Cinematic Video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscinematicvideocapturesupported
func (c_ CaptureDeviceFormat) IsCinematicVideoCaptureSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCinematicVideoCaptureSupported"))
	return rv
}


// Indicates whether the format supports Cinematic Video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscinematicvideocapturesupported
func (c_ CaptureDeviceFormat) SetIsCinematicVideoCaptureSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCinematicVideoCaptureSupported:"), value)
}


// A Boolean value that indicates whether the format supports global tone mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isglobaltonemappingsupported
func (c_ CaptureDeviceFormat) IsGlobalToneMappingSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isGlobalToneMappingSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports global tone mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isglobaltonemappingsupported
func (c_ CaptureDeviceFormat) SetIsGlobalToneMappingSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsGlobalToneMappingSupported:"), value)
}


// A Boolean value that indicates whether this format supports high-quality capture with the current quality prioritization setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ishighphotoqualitysupported
func (c_ CaptureDeviceFormat) IsHighPhotoQualitySupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighPhotoQualitySupported"))
	return rv
}


// A Boolean value that indicates whether this format supports high-quality capture with the current quality prioritization setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ishighphotoqualitysupported
func (c_ CaptureDeviceFormat) SetIsHighPhotoQualitySupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighPhotoQualitySupported:"), value)
}


// A Boolean value that indicates whether this format supports the highest photo quality that the platform delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ishighestphotoqualitysupported
func (c_ CaptureDeviceFormat) IsHighestPhotoQualitySupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighestPhotoQualitySupported"))
	return rv
}


// A Boolean value that indicates whether this format supports the highest photo quality that the platform delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ishighestphotoqualitysupported
func (c_ CaptureDeviceFormat) SetIsHighestPhotoQualitySupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighestPhotoQualitySupported:"), value)
}


// A Boolean value that indicates whether a multi-camera capture session supports this format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ismulticamsupported
func (c_ CaptureDeviceFormat) IsMultiCamSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultiCamSupported"))
	return rv
}


// A Boolean value that indicates whether a multi-camera capture session supports this format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ismulticamsupported
func (c_ CaptureDeviceFormat) SetIsMultiCamSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMultiCamSupported:"), value)
}


// A Boolean value that indicates whether the format supports the Portrait Effect feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isportraiteffectsupported
func (c_ CaptureDeviceFormat) IsPortraitEffectSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports the Portrait Effect feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isportraiteffectsupported
func (c_ CaptureDeviceFormat) SetIsPortraitEffectSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectSupported:"), value)
}


// A Boolean indicating whether the device supports portrait matte effects in still-image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isportraiteffectsmattestillimagedeliverysupported
func (c_ CaptureDeviceFormat) IsPortraitEffectsMatteStillImageDeliverySupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectsMatteStillImageDeliverySupported"))
	return rv
}


// A Boolean indicating whether the device supports portrait matte effects in still-image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isportraiteffectsmattestillimagedeliverysupported
func (c_ CaptureDeviceFormat) SetIsPortraitEffectsMatteStillImageDeliverySupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectsMatteStillImageDeliverySupported:"), value)
}


// Returns
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/issmartframingsupported
func (c_ CaptureDeviceFormat) IsSmartFramingSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSmartFramingSupported"))
	return rv
}


// Returns
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/issmartframingsupported
func (c_ CaptureDeviceFormat) SetIsSmartFramingSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSmartFramingSupported:"), value)
}


// A Boolean value that indicates whether the format supports capturing spatial video to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isspatialvideocapturesupported
func (c_ CaptureDeviceFormat) IsSpatialVideoCaptureSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSpatialVideoCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports capturing spatial video to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isspatialvideocapturesupported
func (c_ CaptureDeviceFormat) SetIsSpatialVideoCaptureSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSpatialVideoCaptureSupported:"), value)
}


// A Boolean value that indicates whether the format supports Studio Light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isstudiolightsupported
func (c_ CaptureDeviceFormat) IsStudioLightSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStudioLightSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports Studio Light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isstudiolightsupported
func (c_ CaptureDeviceFormat) SetIsStudioLightSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStudioLightSupported:"), value)
}


// A Boolean value that indicates whether the format produces video data in a binned format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideobinned
func (c_ CaptureDeviceFormat) IsVideoBinned() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoBinned"))
	return rv
}


// A Boolean value that indicates whether the format produces video data in a binned format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideobinned
func (c_ CaptureDeviceFormat) SetIsVideoBinned(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoBinned:"), value)
}


// A Boolean value that indicates whether the format supports high dynamic range streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideohdrsupported
func (c_ CaptureDeviceFormat) IsVideoHDRSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoHDRSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports high dynamic range streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideohdrsupported
func (c_ CaptureDeviceFormat) SetIsVideoHDRSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoHDRSupported:"), value)
}


// A time value that indicates the maximum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxexposureduration
func (c_ CaptureDeviceFormat) MaxExposureDuration() Time /* not a class type */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("maxExposureDuration"))
	return rv
}


// A time value that indicates the maximum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxexposureduration
func (c_ CaptureDeviceFormat) SetMaxExposureDuration(value Time /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxExposureDuration:"), value)
}


// A floating-point number that indicates the maximum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxiso
func (c_ CaptureDeviceFormat) MaxISO() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("maxISO"))
	return rv
}


// A floating-point number that indicates the maximum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxiso
func (c_ CaptureDeviceFormat) SetMaxISO(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxISO:"), value)
}


// Maximum supported shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxsimulatedaperture
func (c_ CaptureDeviceFormat) MaxSimulatedAperture() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("maxSimulatedAperture"))
	return rv
}


// Maximum supported shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxsimulatedaperture
func (c_ CaptureDeviceFormat) SetMaxSimulatedAperture(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxSimulatedAperture:"), value)
}


// A constant describing the media type of an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/mediatype
func (c_ CaptureDeviceFormat) MediaType() MediaType /* not a class type */ {
	rv := objc.Send[MediaType](c_.ID, objc.Sel("mediaType"))
	return rv
}


// A constant describing the media type of an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/mediatype
func (c_ CaptureDeviceFormat) SetMediaType(value MediaType /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMediaType:"), value)
}


// A time value that indicates the minimum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/minexposureduration
func (c_ CaptureDeviceFormat) MinExposureDuration() Time /* not a class type */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("minExposureDuration"))
	return rv
}


// A time value that indicates the minimum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/minexposureduration
func (c_ CaptureDeviceFormat) SetMinExposureDuration(value Time /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinExposureDuration:"), value)
}


// A floating-point number that indicates the minimum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/miniso
func (c_ CaptureDeviceFormat) MinISO() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("minISO"))
	return rv
}


// A floating-point number that indicates the minimum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/miniso
func (c_ CaptureDeviceFormat) SetMinISO(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinISO:"), value)
}


// Minimum supported shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/minsimulatedaperture
func (c_ CaptureDeviceFormat) MinSimulatedAperture() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("minSimulatedAperture"))
	return rv
}


// Minimum supported shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/minsimulatedaperture
func (c_ CaptureDeviceFormat) SetMinSimulatedAperture(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinSimulatedAperture:"), value)
}


// A Boolean value that indicates whether the device supports reaction effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/reactioneffectssupported
func (c_ CaptureDeviceFormat) ReactionEffectsSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("reactionEffectsSupported"))
	return rv
}


// A Boolean value that indicates whether the device supports reaction effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/reactioneffectssupported
func (c_ CaptureDeviceFormat) SetReactionEffectsSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReactionEffectsSupported:"), value)
}


// The zoom factors at which this device transitions to secondary native resolution modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/secondarynativeresolutionzoomfactors
func (c_ CaptureDeviceFormat) SecondaryNativeResolutionZoomFactors() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("secondaryNativeResolutionZoomFactors"))
	return rv
}


// The zoom factors at which this device transitions to secondary native resolution modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/secondarynativeresolutionzoomfactors
func (c_ CaptureDeviceFormat) SetSecondaryNativeResolutionZoomFactors(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecondaryNativeResolutionZoomFactors:"), value)
}


// The list of the device’s supported color spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supportedcolorspaces
func (c_ CaptureDeviceFormat) SupportedColorSpaces() CaptureColorSpace /* not a class type */ {
	rv := objc.Send[CaptureColorSpace](c_.ID, objc.Sel("supportedColorSpaces"))
	return rv
}


// The list of the device’s supported color spaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supportedcolorspaces
func (c_ CaptureDeviceFormat) SetSupportedColorSpaces(value CaptureColorSpace /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedColorSpaces:"), value)
}


// The list of data formats compatible with this video format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supporteddepthdataformats
func (c_ CaptureDeviceFormat) SupportedDepthDataFormats() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("supportedDepthDataFormats"))
	return rv
}


// The list of data formats compatible with this video format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supporteddepthdataformats
func (c_ CaptureDeviceFormat) SetSupportedDepthDataFormats(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedDepthDataFormats:"), value)
}


// Indicates the supported aspect ratios for the device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supporteddynamicaspectratios
func (c_ CaptureDeviceFormat) SupportedDynamicAspectRatios() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supportedDynamicAspectRatios"))
	return rv
}


// Indicates the supported aspect ratios for the device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supporteddynamicaspectratios
func (c_ CaptureDeviceFormat) SetSupportedDynamicAspectRatios(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedDynamicAspectRatios:"), value)
}


// The maximum photo dimension this format supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supportedmaxphotodimensions
func (c_ CaptureDeviceFormat) SupportedMaxPhotoDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("supportedMaxPhotoDimensions"))
	return rv
}


// The maximum photo dimension this format supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supportedmaxphotodimensions
func (c_ CaptureDeviceFormat) SetSupportedMaxPhotoDimensions(value VideoDimensions /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedMaxPhotoDimensions:"), value)
}


// The zoom factors that a format supports for depth data delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supportedvideozoomfactorsfordepthdatadelivery
func (c_ CaptureDeviceFormat) SupportedVideoZoomFactorsForDepthDataDelivery() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("supportedVideoZoomFactorsForDepthDataDelivery"))
	return rv
}


// The zoom factors that a format supports for depth data delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supportedvideozoomfactorsfordepthdatadelivery
func (c_ CaptureDeviceFormat) SetSupportedVideoZoomFactorsForDepthDataDelivery(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedVideoZoomFactorsForDepthDataDelivery:"), value)
}


// The zoom ranges that support the delivery of depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supportedvideozoomrangesfordepthdatadelivery
func (c_ CaptureDeviceFormat) SupportedVideoZoomRangesForDepthDataDelivery() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("supportedVideoZoomRangesForDepthDataDelivery"))
	return rv
}


// The zoom ranges that support the delivery of depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supportedvideozoomrangesfordepthdatadelivery
func (c_ CaptureDeviceFormat) SetSupportedVideoZoomRangesForDepthDataDelivery(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedVideoZoomRangesForDepthDataDelivery:"), value)
}


// The system’s recommended exposure bias range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedexposurebiasrange
func (c_ CaptureDeviceFormat) SystemRecommendedExposureBiasRange() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("systemRecommendedExposureBiasRange"))
	return rv
}


// The system’s recommended exposure bias range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedexposurebiasrange
func (c_ CaptureDeviceFormat) SetSystemRecommendedExposureBiasRange(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemRecommendedExposureBiasRange:"), value)
}


// The system’s recommended zoom range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedvideozoomrange
func (c_ CaptureDeviceFormat) SystemRecommendedVideoZoomRange() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("systemRecommendedVideoZoomRange"))
	return rv
}


// The system’s recommended zoom range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedvideozoomrange
func (c_ CaptureDeviceFormat) SetSystemRecommendedVideoZoomRange(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemRecommendedVideoZoomRange:"), value)
}


// The list of capture output subclasses not allowed for capture with this format, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/unsupportedcaptureoutputclasses
func (c_ CaptureDeviceFormat) UnsupportedCaptureOutputClasses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("unsupportedCaptureOutputClasses"))
	return rv
}


// The list of capture output subclasses not allowed for capture with this format, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/unsupportedcaptureoutputclasses
func (c_ CaptureDeviceFormat) SetUnsupportedCaptureOutputClasses(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUnsupportedCaptureOutputClasses:"), value)
}


// Indicates the format’s horizontal field of view in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videofieldofview
func (c_ CaptureDeviceFormat) VideoFieldOfView() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("videoFieldOfView"))
	return rv
}


// Indicates the format’s horizontal field of view in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videofieldofview
func (c_ CaptureDeviceFormat) SetVideoFieldOfView(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFieldOfView:"), value)
}


// The minimum and maximum frame rates available when Background Replacement is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforbackgroundreplacement
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForBackgroundReplacement() FrameRateRange /* not a class type */ {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForBackgroundReplacement"))
	return rv
}


// The minimum and maximum frame rates available when Background Replacement is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforbackgroundreplacement
func (c_ CaptureDeviceFormat) SetVideoFrameRateRangeForBackgroundReplacement(value FrameRateRange /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFrameRateRangeForBackgroundReplacement:"), value)
}


// The range of frame rates available when Center Stage is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforcenterstage
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForCenterStage() FrameRateRange /* not a class type */ {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForCenterStage"))
	return rv
}


// The range of frame rates available when Center Stage is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforcenterstage
func (c_ CaptureDeviceFormat) SetVideoFrameRateRangeForCenterStage(value FrameRateRange /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFrameRateRangeForCenterStage:"), value)
}


// Indicates the minimum / maximum frame rates available when Cinematic Video capture is enabled on the device input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforcinematicvideo
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForCinematicVideo() FrameRateRange /* not a class type */ {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForCinematicVideo"))
	return rv
}


// Indicates the minimum / maximum frame rates available when Cinematic Video capture is enabled on the device input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforcinematicvideo
func (c_ CaptureDeviceFormat) SetVideoFrameRateRangeForCinematicVideo(value FrameRateRange /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFrameRateRangeForCinematicVideo:"), value)
}


// The range of frame rates available when Portrait Effect is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforportraiteffect
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForPortraitEffect() FrameRateRange /* not a class type */ {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForPortraitEffect"))
	return rv
}


// The range of frame rates available when Portrait Effect is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforportraiteffect
func (c_ CaptureDeviceFormat) SetVideoFrameRateRangeForPortraitEffect(value FrameRateRange /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFrameRateRangeForPortraitEffect:"), value)
}


// Indicates the minimum and maximum frame rates available when a reaction effect runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforreactioneffectsinprogress
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForReactionEffectsInProgress() FrameRateRange /* not a class type */ {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForReactionEffectsInProgress"))
	return rv
}


// Indicates the minimum and maximum frame rates available when a reaction effect runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforreactioneffectsinprogress
func (c_ CaptureDeviceFormat) SetVideoFrameRateRangeForReactionEffectsInProgress(value FrameRateRange /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFrameRateRangeForReactionEffectsInProgress:"), value)
}


// A value that indicates the minimum and maximum frame rates available when a user enables Studio Light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforstudiolight
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForStudioLight() FrameRateRange /* not a class type */ {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForStudioLight"))
	return rv
}


// A value that indicates the minimum and maximum frame rates available when a user enables Studio Light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videoframeraterangeforstudiolight
func (c_ CaptureDeviceFormat) SetVideoFrameRateRangeForStudioLight(value FrameRateRange /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoFrameRateRangeForStudioLight:"), value)
}


// A maximum zoom factor the format allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videomaxzoomfactor
func (c_ CaptureDeviceFormat) VideoMaxZoomFactor() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxZoomFactor"))
	return rv
}


// A maximum zoom factor the format allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videomaxzoomfactor
func (c_ CaptureDeviceFormat) SetVideoMaxZoomFactor(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxZoomFactor:"), value)
}


// The maximum zoom factor available when Center Stage is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videomaxzoomfactorforcenterstage
func (c_ CaptureDeviceFormat) VideoMaxZoomFactorForCenterStage() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxZoomFactorForCenterStage"))
	return rv
}


// The maximum zoom factor available when Center Stage is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videomaxzoomfactorforcenterstage
func (c_ CaptureDeviceFormat) SetVideoMaxZoomFactorForCenterStage(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxZoomFactorForCenterStage:"), value)
}


// Indicates the maximum zoom factor available for the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videomaxzoomfactorforcinematicvideo
func (c_ CaptureDeviceFormat) VideoMaxZoomFactorForCinematicVideo() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxZoomFactorForCinematicVideo"))
	return rv
}


// Indicates the maximum zoom factor available for the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videomaxzoomfactorforcinematicvideo
func (c_ CaptureDeviceFormat) SetVideoMaxZoomFactorForCinematicVideo(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMaxZoomFactorForCinematicVideo:"), value)
}


// The minimum zoom factor available when Center Stage is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videominzoomfactorforcenterstage
func (c_ CaptureDeviceFormat) VideoMinZoomFactorForCenterStage() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMinZoomFactorForCenterStage"))
	return rv
}


// The minimum zoom factor available when Center Stage is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videominzoomfactorforcenterstage
func (c_ CaptureDeviceFormat) SetVideoMinZoomFactorForCenterStage(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinZoomFactorForCenterStage:"), value)
}


// Indicates the minimum zoom factor available for the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videominzoomfactorforcinematicvideo
func (c_ CaptureDeviceFormat) VideoMinZoomFactorForCinematicVideo() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMinZoomFactorForCinematicVideo"))
	return rv
}


// Indicates the minimum zoom factor available for the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videominzoomfactorforcinematicvideo
func (c_ CaptureDeviceFormat) SetVideoMinZoomFactorForCinematicVideo(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinZoomFactorForCinematicVideo:"), value)
}


// A list of frame rate ranges that a format supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videosupportedframerateranges
func (c_ CaptureDeviceFormat) VideoSupportedFrameRateRanges() FrameRateRange /* not a class type */ {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoSupportedFrameRateRanges"))
	return rv
}


// A list of frame rate ranges that a format supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videosupportedframerateranges
func (c_ CaptureDeviceFormat) SetVideoSupportedFrameRateRanges(value FrameRateRange /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoSupportedFrameRateRanges:"), value)
}


// A threshold at which the system upscales pixel data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videozoomfactorupscalethreshold
func (c_ CaptureDeviceFormat) VideoZoomFactorUpscaleThreshold() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoZoomFactorUpscaleThreshold"))
	return rv
}


// A threshold at which the system upscales pixel data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videozoomfactorupscalethreshold
func (c_ CaptureDeviceFormat) SetVideoZoomFactorUpscaleThreshold(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoZoomFactorUpscaleThreshold:"), value)
}


// A Boolean value that indicates whether the format supports zoom factors outside the range supported for depth delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/zoomfactorsoutsideofvideozoomrangesfordepthdeliverysupported
func (c_ CaptureDeviceFormat) ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("zoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported"))
	return rv
}


// A Boolean value that indicates whether the format supports zoom factors outside the range supported for depth delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/zoomfactorsoutsideofvideozoomrangesfordepthdeliverysupported
func (c_ CaptureDeviceFormat) SetZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported:"), value)
}


// The currently active depth data format of the capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activedepthdataformat
func (c_ CaptureDeviceFormat) ActiveDepthDataFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeDepthDataFormat"))
	return rv
}


// The currently active depth data format of the capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activedepthdataformat
func (c_ CaptureDeviceFormat) SetActiveDepthDataFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveDepthDataFormat:"), value)
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureDeviceFormat) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureDeviceFormat) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}


// The capture formats a device supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/formats
func (c_ CaptureDeviceFormat) Formats() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("formats"))
	return rv
}


// The capture formats a device supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/formats
func (c_ CaptureDeviceFormat) SetFormats(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormats:"), value)
}


// A value that controls the cropping and enlargement of images captured by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/videozoomfactor
func (c_ CaptureDeviceFormat) VideoZoomFactor() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoZoomFactor"))
	return rv
}


// A value that controls the cropping and enlargement of images captured by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/videozoomfactor
func (c_ CaptureDeviceFormat) SetVideoZoomFactor(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoZoomFactor:"), value)
}



