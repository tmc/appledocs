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
	SecondaryNativeResolutionZoomFactors() []foundation.Number
	SupportedColorSpaces() []foundation.Number
	SupportedMaxPhotoDimensions() []foundation.Value
	SupportedVideoZoomRangesForDepthDataDelivery() []ZoomRange
	SystemRecommendedExposureBiasRange() IAVExposureBiasRange
	SystemRecommendedVideoZoomRange() IAVZoomRange
	AutoFocusSystem() CaptureAutoFocusSystem
	DefaultSimulatedAperture() float32
	FormatDescription() FormatDescriptionRef /* not a class type */
	AutoVideoFrameRateSupported() bool
	BackgroundReplacementSupported() bool
	CameraLensSmudgeDetectionSupported() bool
	CenterStageSupported() bool
	CinematicVideoCaptureSupported() bool
	HighPhotoQualitySupported() bool
	PortraitEffectSupported() bool
	SpatialVideoCaptureSupported() bool
	StudioLightSupported() bool
	MaxSimulatedAperture() float32
	MediaType() MediaType /* typedef */
	MinSimulatedAperture() float32
	ReactionEffectsSupported() bool
	VideoFrameRateRangeForBackgroundReplacement() IAVFrameRateRange
	VideoFrameRateRangeForCenterStage() IAVFrameRateRange
	VideoFrameRateRangeForCinematicVideo() IAVFrameRateRange
	VideoFrameRateRangeForPortraitEffect() IAVFrameRateRange
	VideoFrameRateRangeForReactionEffectsInProgress() IAVFrameRateRange
	VideoFrameRateRangeForStudioLight() IAVFrameRateRange
	VideoMaxZoomFactorForCenterStage() float64
	VideoMaxZoomFactorForCinematicVideo() float64
	VideoMinZoomFactorForCenterStage() float64
	VideoMinZoomFactorForCinematicVideo() float64
	VideoSupportedFrameRateRanges() []FrameRateRange
	ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool
	IsAutoVideoFrameRateSupported() bool
	SetIsAutoVideoFrameRateSupported(value bool)
	IsBackgroundReplacementSupported() bool
	SetIsBackgroundReplacementSupported(value bool)
	IsCameraLensSmudgeDetectionSupported() bool
	SetIsCameraLensSmudgeDetectionSupported(value bool)
	IsCenterStageSupported() bool
	SetIsCenterStageSupported(value bool)
	IsCinematicVideoCaptureSupported() bool
	SetIsCinematicVideoCaptureSupported(value bool)
	IsGlobalToneMappingSupported() bool
	SetIsGlobalToneMappingSupported(value bool)
	IsHighPhotoQualitySupported() bool
	SetIsHighPhotoQualitySupported(value bool)
	IsHighestPhotoQualitySupported() bool
	SetIsHighestPhotoQualitySupported(value bool)
	IsMultiCamSupported() bool
	SetIsMultiCamSupported(value bool)
	IsPortraitEffectSupported() bool
	SetIsPortraitEffectSupported(value bool)
	IsPortraitEffectsMatteStillImageDeliverySupported() bool
	SetIsPortraitEffectsMatteStillImageDeliverySupported(value bool)
	IsSmartFramingSupported() bool
	SetIsSmartFramingSupported(value bool)
	IsSpatialVideoCaptureSupported() bool
	SetIsSpatialVideoCaptureSupported(value bool)
	IsStudioLightSupported() bool
	SetIsStudioLightSupported(value bool)
	IsVideoBinned() bool
	SetIsVideoBinned(value bool)
	IsVideoHDRSupported() bool
	SetIsVideoHDRSupported(value bool)
	ActiveDepthDataFormat() IAVCaptureDeviceFormat
	SetActiveDepthDataFormat(value IAVCaptureDeviceFormat)
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	Formats() IAVCaptureDeviceFormat
	SetFormats(value IAVCaptureDeviceFormat)
	VideoZoomFactor() float64
	SetVideoZoomFactor(value float64)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceFormatClass) Alloc() CaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

























// The zoom factors at which this device transitions to secondary native resolution modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/secondaryNativeResolutionZoomFactors
func (c_ CaptureDeviceFormat) SecondaryNativeResolutionZoomFactors() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("secondaryNativeResolutionZoomFactors"))
	return rv
}


// The list of color spaces the format supports for image and video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/supportedColorSpaces
func (c_ CaptureDeviceFormat) SupportedColorSpaces() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("supportedColorSpaces"))
	return rv
}


// The maximum photo dimension this format supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/supportedMaxPhotoDimensions
func (c_ CaptureDeviceFormat) SupportedMaxPhotoDimensions() []foundation.Value {
	rv := objc.Send[[]foundation.Value](c_.ID, objc.Sel("supportedMaxPhotoDimensions"))
	return rv
}


// The zoom ranges that support the delivery of depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/supportedVideoZoomRangesForDepthDataDelivery
func (c_ CaptureDeviceFormat) SupportedVideoZoomRangesForDepthDataDelivery() []ZoomRange {
	rv := objc.Send[[]ZoomRange](c_.ID, objc.Sel("supportedVideoZoomRangesForDepthDataDelivery"))
	return rv
}


// The system’s recommended exposure bias range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/systemRecommendedExposureBiasRange
func (c_ CaptureDeviceFormat) SystemRecommendedExposureBiasRange() IAVExposureBiasRange {
	rv := objc.Send[ExposureBiasRange](c_.ID, objc.Sel("systemRecommendedExposureBiasRange"))
	return rv
}


// The system’s recommended zoom range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/systemRecommendedVideoZoomRange
func (c_ CaptureDeviceFormat) SystemRecommendedVideoZoomRange() IAVZoomRange {
	rv := objc.Send[ZoomRange](c_.ID, objc.Sel("systemRecommendedVideoZoomRange"))
	return rv
}


// The auto focus system the format uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/autoFocusSystem-swift.property
func (c_ CaptureDeviceFormat) AutoFocusSystem() CaptureAutoFocusSystem {
	rv := objc.Send[CaptureAutoFocusSystem](c_.ID, objc.Sel("autoFocusSystem"))
	return rv
}


// Default shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/defaultSimulatedAperture
func (c_ CaptureDeviceFormat) DefaultSimulatedAperture() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("defaultSimulatedAperture"))
	return rv
}


// An object describing the capture format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/formatDescription
func (c_ CaptureDeviceFormat) FormatDescription() FormatDescriptionRef /* not a class type */ {
	rv := objc.Send[FormatDescriptionRef](c_.ID, objc.Sel("formatDescription"))
	return rv
}


// A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isAutoVideoFrameRateSupported
func (c_ CaptureDeviceFormat) AutoVideoFrameRateSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoVideoFrameRateSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports background replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isBackgroundReplacementSupported
func (c_ CaptureDeviceFormat) BackgroundReplacementSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("backgroundReplacementSupported"))
	return rv
}


// Whether camera lens smudge detection is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isCameraLensSmudgeDetectionSupported
func (c_ CaptureDeviceFormat) CameraLensSmudgeDetectionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraLensSmudgeDetectionSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports Center Stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isCenterStageSupported
func (c_ CaptureDeviceFormat) CenterStageSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("centerStageSupported"))
	return rv
}


// Indicates whether the format supports Cinematic Video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isCinematicVideoCaptureSupported
func (c_ CaptureDeviceFormat) CinematicVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cinematicVideoCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether this format supports high-quality capture with the current quality prioritization setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isHighPhotoQualitySupported
func (c_ CaptureDeviceFormat) HighPhotoQualitySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highPhotoQualitySupported"))
	return rv
}


// A Boolean value that indicates whether the format supports the Portrait Effect feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isPortraitEffectSupported
func (c_ CaptureDeviceFormat) PortraitEffectSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports capturing spatial video to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isSpatialVideoCaptureSupported
func (c_ CaptureDeviceFormat) SpatialVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("spatialVideoCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports Studio Light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isStudioLightSupported
func (c_ CaptureDeviceFormat) StudioLightSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("studioLightSupported"))
	return rv
}


// Maximum supported shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/maxSimulatedAperture
func (c_ CaptureDeviceFormat) MaxSimulatedAperture() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("maxSimulatedAperture"))
	return rv
}


// A constant describing the media type of an active or supported format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/mediaType
func (c_ CaptureDeviceFormat) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("mediaType"))
	return rv
}


// Minimum supported shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/minSimulatedAperture
func (c_ CaptureDeviceFormat) MinSimulatedAperture() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("minSimulatedAperture"))
	return rv
}


// A Boolean value that indicates whether the device supports reaction effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/reactionEffectsSupported
func (c_ CaptureDeviceFormat) ReactionEffectsSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("reactionEffectsSupported"))
	return rv
}


// The minimum and maximum frame rates available when Background Replacement is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForBackgroundReplacement
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForBackgroundReplacement() IAVFrameRateRange {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForBackgroundReplacement"))
	return rv
}


// The range of frame rates available when Center Stage is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForCenterStage
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForCenterStage() IAVFrameRateRange {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForCenterStage"))
	return rv
}


// Indicates the minimum / maximum frame rates available when Cinematic Video capture is enabled on the device input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForCinematicVideo
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForCinematicVideo() IAVFrameRateRange {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForCinematicVideo"))
	return rv
}


// The range of frame rates available when Portrait Effect is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForPortraitEffect
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForPortraitEffect() IAVFrameRateRange {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForPortraitEffect"))
	return rv
}


// Indicates the minimum and maximum frame rates available when a reaction effect runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForReactionEffectsInProgress
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForReactionEffectsInProgress() IAVFrameRateRange {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForReactionEffectsInProgress"))
	return rv
}


// A value that indicates the minimum and maximum frame rates available when a user enables Studio Light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForStudioLight
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForStudioLight() IAVFrameRateRange {
	rv := objc.Send[FrameRateRange](c_.ID, objc.Sel("videoFrameRateRangeForStudioLight"))
	return rv
}


// The maximum zoom factor available when Center Stage is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMaxZoomFactorForCenterStage
func (c_ CaptureDeviceFormat) VideoMaxZoomFactorForCenterStage() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxZoomFactorForCenterStage"))
	return rv
}


// Indicates the maximum zoom factor available for the property when Cinematic Video capture is enabled on the device input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMaxZoomFactorForCinematicVideo
func (c_ CaptureDeviceFormat) VideoMaxZoomFactorForCinematicVideo() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMaxZoomFactorForCinematicVideo"))
	return rv
}


// The minimum zoom factor available when Center Stage is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMinZoomFactorForCenterStage
func (c_ CaptureDeviceFormat) VideoMinZoomFactorForCenterStage() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMinZoomFactorForCenterStage"))
	return rv
}


// Indicates the minimum zoom factor available for the property when Cinematic Video capture is enabled on the device input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMinZoomFactorForCinematicVideo
func (c_ CaptureDeviceFormat) VideoMinZoomFactorForCinematicVideo() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoMinZoomFactorForCinematicVideo"))
	return rv
}


// A list of frame rate ranges that a format supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoSupportedFrameRateRanges
func (c_ CaptureDeviceFormat) VideoSupportedFrameRateRanges() []FrameRateRange {
	rv := objc.Send[[]FrameRateRange](c_.ID, objc.Sel("videoSupportedFrameRateRanges"))
	return rv
}


// A Boolean value that indicates whether the format supports zoom factors outside the range supported for depth delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/zoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported
func (c_ CaptureDeviceFormat) ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("zoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported"))
	return rv
}


// A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isautovideoframeratesupported
func (c_ CaptureDeviceFormat) IsAutoVideoFrameRateSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoVideoFrameRateSupported"))
	return rv
}


// A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isautovideoframeratesupported
func (c_ CaptureDeviceFormat) SetIsAutoVideoFrameRateSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoVideoFrameRateSupported:"), value)
}


// A Boolean value that indicates whether the format supports background replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isbackgroundreplacementsupported
func (c_ CaptureDeviceFormat) IsBackgroundReplacementSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isBackgroundReplacementSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports background replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isbackgroundreplacementsupported
func (c_ CaptureDeviceFormat) SetIsBackgroundReplacementSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsBackgroundReplacementSupported:"), value)
}


// Whether camera lens smudge detection is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscameralenssmudgedetectionsupported
func (c_ CaptureDeviceFormat) IsCameraLensSmudgeDetectionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraLensSmudgeDetectionSupported"))
	return rv
}


// Whether camera lens smudge detection is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscameralenssmudgedetectionsupported
func (c_ CaptureDeviceFormat) SetIsCameraLensSmudgeDetectionSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraLensSmudgeDetectionSupported:"), value)
}


// A Boolean value that indicates whether the format supports Center Stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscenterstagesupported
func (c_ CaptureDeviceFormat) IsCenterStageSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCenterStageSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports Center Stage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscenterstagesupported
func (c_ CaptureDeviceFormat) SetIsCenterStageSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCenterStageSupported:"), value)
}


// Indicates whether the format supports Cinematic Video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscinematicvideocapturesupported
func (c_ CaptureDeviceFormat) IsCinematicVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCinematicVideoCaptureSupported"))
	return rv
}


// Indicates whether the format supports Cinematic Video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/iscinematicvideocapturesupported
func (c_ CaptureDeviceFormat) SetIsCinematicVideoCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCinematicVideoCaptureSupported:"), value)
}


// A Boolean value that indicates whether the format supports global tone mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isglobaltonemappingsupported
func (c_ CaptureDeviceFormat) IsGlobalToneMappingSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isGlobalToneMappingSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports global tone mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isglobaltonemappingsupported
func (c_ CaptureDeviceFormat) SetIsGlobalToneMappingSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsGlobalToneMappingSupported:"), value)
}


// A Boolean value that indicates whether this format supports high-quality capture with the current quality prioritization setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ishighphotoqualitysupported
func (c_ CaptureDeviceFormat) IsHighPhotoQualitySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighPhotoQualitySupported"))
	return rv
}


// A Boolean value that indicates whether this format supports high-quality capture with the current quality prioritization setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ishighphotoqualitysupported
func (c_ CaptureDeviceFormat) SetIsHighPhotoQualitySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighPhotoQualitySupported:"), value)
}


// A Boolean value that indicates whether this format supports the highest photo quality that the platform delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ishighestphotoqualitysupported
func (c_ CaptureDeviceFormat) IsHighestPhotoQualitySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighestPhotoQualitySupported"))
	return rv
}


// A Boolean value that indicates whether this format supports the highest photo quality that the platform delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ishighestphotoqualitysupported
func (c_ CaptureDeviceFormat) SetIsHighestPhotoQualitySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighestPhotoQualitySupported:"), value)
}


// A Boolean value that indicates whether a multi-camera capture session supports this format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ismulticamsupported
func (c_ CaptureDeviceFormat) IsMultiCamSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultiCamSupported"))
	return rv
}


// A Boolean value that indicates whether a multi-camera capture session supports this format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/ismulticamsupported
func (c_ CaptureDeviceFormat) SetIsMultiCamSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMultiCamSupported:"), value)
}


// A Boolean value that indicates whether the format supports the Portrait Effect feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isportraiteffectsupported
func (c_ CaptureDeviceFormat) IsPortraitEffectSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports the Portrait Effect feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isportraiteffectsupported
func (c_ CaptureDeviceFormat) SetIsPortraitEffectSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectSupported:"), value)
}


// A Boolean indicating whether the device supports portrait matte effects in still-image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isportraiteffectsmattestillimagedeliverysupported
func (c_ CaptureDeviceFormat) IsPortraitEffectsMatteStillImageDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectsMatteStillImageDeliverySupported"))
	return rv
}


// A Boolean indicating whether the device supports portrait matte effects in still-image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isportraiteffectsmattestillimagedeliverysupported
func (c_ CaptureDeviceFormat) SetIsPortraitEffectsMatteStillImageDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectsMatteStillImageDeliverySupported:"), value)
}


// Returns
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/issmartframingsupported
func (c_ CaptureDeviceFormat) IsSmartFramingSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSmartFramingSupported"))
	return rv
}


// Returns
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/issmartframingsupported
func (c_ CaptureDeviceFormat) SetIsSmartFramingSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSmartFramingSupported:"), value)
}


// A Boolean value that indicates whether the format supports capturing spatial video to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isspatialvideocapturesupported
func (c_ CaptureDeviceFormat) IsSpatialVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSpatialVideoCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports capturing spatial video to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isspatialvideocapturesupported
func (c_ CaptureDeviceFormat) SetIsSpatialVideoCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSpatialVideoCaptureSupported:"), value)
}


// A Boolean value that indicates whether the format supports Studio Light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isstudiolightsupported
func (c_ CaptureDeviceFormat) IsStudioLightSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStudioLightSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports Studio Light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isstudiolightsupported
func (c_ CaptureDeviceFormat) SetIsStudioLightSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStudioLightSupported:"), value)
}


// A Boolean value that indicates whether the format produces video data in a binned format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideobinned
func (c_ CaptureDeviceFormat) IsVideoBinned() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoBinned"))
	return rv
}


// A Boolean value that indicates whether the format produces video data in a binned format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideobinned
func (c_ CaptureDeviceFormat) SetIsVideoBinned(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoBinned:"), value)
}


// A Boolean value that indicates whether the format supports high dynamic range streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideohdrsupported
func (c_ CaptureDeviceFormat) IsVideoHDRSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoHDRSupported"))
	return rv
}


// A Boolean value that indicates whether the format supports high dynamic range streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/isvideohdrsupported
func (c_ CaptureDeviceFormat) SetIsVideoHDRSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoHDRSupported:"), value)
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
func (c_ CaptureDeviceFormat) VideoZoomFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoZoomFactor"))
	return rv
}


// A value that controls the cropping and enlargement of images captured by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/videozoomfactor
func (c_ CaptureDeviceFormat) SetVideoZoomFactor(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoZoomFactor:"), value)
}







