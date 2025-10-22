// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CapturePhotoOutput] class.
var (
	CapturePhotoOutputClass     _CapturePhotoOutputClass
	CapturePhotoOutputClassOnce sync.Once
)

func getCapturePhotoOutputClass() _CapturePhotoOutputClass {
	CapturePhotoOutputClassOnce.Do(func() {
		CapturePhotoOutputClass = _CapturePhotoOutputClass{objc.GetClass("AVCapturePhotoOutput")}
	})
	return CapturePhotoOutputClass
}

type _CapturePhotoOutputClass struct {
	class objc.Class
}

// An interface definition for the [CapturePhotoOutput] class.
type ICapturePhotoOutput interface {
	ICaptureOutput
	CapturePhotoWithSettingsDelegate(settings IAVCapturePhotoSettings, delegate objectivec.IObject)
	ConstantColorEnabled() bool
	SetConstantColorEnabled(value bool)
	ActiveColorSpace() CaptureColorSpace
	SetActiveColorSpace(value ICaptureColorSpace)
	PortraitEffectsMatte() AVPortraitEffectsMatte
	SetPortraitEffectsMatte(value IAVPortraitEffectsMatte)
	AvailableLivePhotoVideoCodecTypes() VideoCodecType
	SetAvailableLivePhotoVideoCodecTypes(value IVideoCodecType)
	AvailablePhotoCodecTypes() VideoCodecType
	SetAvailablePhotoCodecTypes(value IVideoCodecType)
	AvailablePhotoFileTypes() FileType
	SetAvailablePhotoFileTypes(value IFileType)
	AvailablePhotoPixelFormatTypes() unsafe.Pointer
	SetAvailablePhotoPixelFormatTypes(value unsafe.Pointer)
	AvailableRawPhotoCodecTypes() VideoCodecType
	SetAvailableRawPhotoCodecTypes(value IVideoCodecType)
	AvailableRawPhotoFileTypes() FileType
	SetAvailableRawPhotoFileTypes(value IFileType)
	AvailableRawPhotoPixelFormatTypes() unsafe.Pointer
	SetAvailableRawPhotoPixelFormatTypes(value unsafe.Pointer)
	AvailableSemanticSegmentationMatteTypes() unsafe.Pointer
	SetAvailableSemanticSegmentationMatteTypes(value unsafe.Pointer)
	CaptureReadiness() unsafe.Pointer
	SetCaptureReadiness(value unsafe.Pointer)
	EnabledSemanticSegmentationMatteTypes() unsafe.Pointer
	SetEnabledSemanticSegmentationMatteTypes(value unsafe.Pointer)
	IsAppleProRAWEnabled() bool
	SetIsAppleProRAWEnabled(value bool)
	IsAppleProRAWSupported() bool
	SetIsAppleProRAWSupported(value bool)
	IsAutoDeferredPhotoDeliveryEnabled() bool
	SetIsAutoDeferredPhotoDeliveryEnabled(value bool)
	IsAutoDeferredPhotoDeliverySupported() bool
	SetIsAutoDeferredPhotoDeliverySupported(value bool)
	IsAutoRedEyeReductionSupported() bool
	SetIsAutoRedEyeReductionSupported(value bool)
	IsCameraCalibrationDataDeliverySupported() bool
	SetIsCameraCalibrationDataDeliverySupported(value bool)
	IsCameraSensorOrientationCompensationEnabled() bool
	SetIsCameraSensorOrientationCompensationEnabled(value bool)
	IsCameraSensorOrientationCompensationSupported() bool
	SetIsCameraSensorOrientationCompensationSupported(value bool)
	IsConstantColorEnabled() bool
	SetIsConstantColorEnabled(value bool)
	IsConstantColorSupported() bool
	SetIsConstantColorSupported(value bool)
	IsContentAwareDistortionCorrectionEnabled() bool
	SetIsContentAwareDistortionCorrectionEnabled(value bool)
	IsContentAwareDistortionCorrectionSupported() bool
	SetIsContentAwareDistortionCorrectionSupported(value bool)
	IsDepthDataDeliveryEnabled() bool
	SetIsDepthDataDeliveryEnabled(value bool)
	IsDepthDataDeliverySupported() bool
	SetIsDepthDataDeliverySupported(value bool)
	IsFastCapturePrioritizationEnabled() bool
	SetIsFastCapturePrioritizationEnabled(value bool)
	IsFastCapturePrioritizationSupported() bool
	SetIsFastCapturePrioritizationSupported(value bool)
	IsFlashScene() bool
	SetIsFlashScene(value bool)
	IsHighResolutionCaptureEnabled() bool
	SetIsHighResolutionCaptureEnabled(value bool)
	IsLensStabilizationDuringBracketedCaptureSupported() bool
	SetIsLensStabilizationDuringBracketedCaptureSupported(value bool)
	IsLivePhotoAutoTrimmingEnabled() bool
	SetIsLivePhotoAutoTrimmingEnabled(value bool)
	IsLivePhotoCaptureEnabled() bool
	SetIsLivePhotoCaptureEnabled(value bool)
	IsLivePhotoCaptureSupported() bool
	SetIsLivePhotoCaptureSupported(value bool)
	IsLivePhotoCaptureSuspended() bool
	SetIsLivePhotoCaptureSuspended(value bool)
	IsPortraitEffectsMatteDeliveryEnabled() bool
	SetIsPortraitEffectsMatteDeliveryEnabled(value bool)
	IsPortraitEffectsMatteDeliverySupported() bool
	SetIsPortraitEffectsMatteDeliverySupported(value bool)
	IsResponsiveCaptureEnabled() bool
	SetIsResponsiveCaptureEnabled(value bool)
	IsResponsiveCaptureSupported() bool
	SetIsResponsiveCaptureSupported(value bool)
	IsShutterSoundSuppressionSupported() bool
	SetIsShutterSoundSuppressionSupported(value bool)
	IsVirtualDeviceConstituentPhotoDeliveryEnabled() bool
	SetIsVirtualDeviceConstituentPhotoDeliveryEnabled(value bool)
	IsVirtualDeviceConstituentPhotoDeliverySupported() bool
	SetIsVirtualDeviceConstituentPhotoDeliverySupported(value bool)
	IsVirtualDeviceFusionSupported() bool
	SetIsVirtualDeviceFusionSupported(value bool)
	IsZeroShutterLagEnabled() bool
	SetIsZeroShutterLagEnabled(value bool)
	IsZeroShutterLagSupported() bool
	SetIsZeroShutterLagSupported(value bool)
	MaxBracketedCapturePhotoCount() int
	SetMaxBracketedCapturePhotoCount(value int)
	MaxPhotoDimensions() unsafe.Pointer
	SetMaxPhotoDimensions(value unsafe.Pointer)
	MaxPhotoQualityPrioritization() unsafe.Pointer
	SetMaxPhotoQualityPrioritization(value unsafe.Pointer)
	PhotoSettingsForSceneMonitoring() AVCapturePhotoSettings
	SetPhotoSettingsForSceneMonitoring(value IAVCapturePhotoSettings)
	PreparedPhotoSettingsArray() AVCapturePhotoSettings
	SetPreparedPhotoSettingsArray(value IAVCapturePhotoSettings)
	PreservesLivePhotoCaptureSuspendedOnSessionStop() bool
	SetPreservesLivePhotoCaptureSuspendedOnSessionStop(value bool)
	SupportedFlashModes() unsafe.Pointer
	SetSupportedFlashModes(value unsafe.Pointer)
	FlashMode() unsafe.Pointer
	SetFlashMode(value unsafe.Pointer)
	UniqueID() unsafe.Pointer
	SetUniqueID(value unsafe.Pointer)
}

// A capture output for still image, Live Photos, and other photography workflows.
//
// provides an interface for capture workflows related to still photography. In addition to basic capture of still images, a photo output supports RAW-format capture, bracketed capture of multiple images, Live Photos, and wide-gamut color. You can output captured photos in a variety of formats and codecs, including RAW format DNG files, HEVC format HEIF files, and JPEG files. To capture photos with the class, follow these steps: Create an object. Use its properties to determine supported capture settings and to enable certain features (for example, whether to capture Live Photos). Create and configure an object to choose features and settings for a specific capture (for example, whether to enable image stabilization or flash). Capture an image by passing your photo settings object to the method along with a delegate object implementing the protocol. The photo capture output then calls your delegate to notify you of significant events during the capture process. Some photo capture settings, such as the property, include options for automatic behavior. For such settings, the photo output determines whether to use that feature at the moment of capture—you don’t know when requesting a capture whether the feature will be enabled when the capture completes. When the photo capture output calls your methods with information about the completed or in-progress capture, it also provides an object that details which automatic features are set for that capture. The resolved settings object’s property matches the value of the object you used to request capture. Enabling certain photo features (Live Photo capture and high resolution capture) requires a reconfiguration of the capture render pipeline. To opt into these features, set the , , and properties before calling your object’s method. Changing any of these properties while the session is running disrupts the capture render pipeline: Live Photo captures in progress end immediately, unfulfilled photo requests abort, and video preview temporarily freezes. Using a photo capture output adds other requirements to your object: A capture session can’t support both Live Photo capture and movie file output. If your capture session includes an object, the property becomes . (As an alternative, you can use the class to output video buffers at the same resolution as a simultaneous Live Photo capture). A capture session can’t contain both an object and an object. The class includes all functionality of (and deprecates) the class. The class implicitly supports wide-gamut color photography. If the source object’s value is , the capture output produces photos with wide color information (unless your object specifies an output format that doesn’t support wide color).


// A capture output for still image, Live Photos, and other photography workflows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput

type CapturePhotoOutput struct {
	CaptureOutput
}

// CapturePhotoOutputFrom constructs a [CapturePhotoOutput] from an unsafe.Pointer.
//
// A capture output for still image, Live Photos, and other photography workflows.
func CapturePhotoOutputFrom(ptr unsafe.Pointer) CapturePhotoOutput {
	return CapturePhotoOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoOutputClass) Alloc() CapturePhotoOutput {
	rv := objc.Send[CapturePhotoOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CapturePhotoOutputClass) New() CapturePhotoOutput {
	rv := objc.Send[CapturePhotoOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhotoOutput) Init() CapturePhotoOutput {
	rv := objc.Send[CapturePhotoOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhotoOutput) Autorelease() CapturePhotoOutput {
	rv := objc.Send[CapturePhotoOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhotoOutput creates a new CapturePhotoOutput instance.
func NewCapturePhotoOutput() CapturePhotoOutput {
	return getCapturePhotoOutputClass().New()
}




// Initiates a photo capture using the specified settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/capturePhoto(with:delegate:)

func (c_ CapturePhotoOutput) CapturePhotoWithSettingsDelegate(settings IAVCapturePhotoSettings, delegate objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("capturePhotoWithSettings:delegate:"), settings, delegate)
}


// A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isConstantColorEnabled

func (c_ CapturePhotoOutput) ConstantColorEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("constantColorEnabled"))
	return rv
}


// A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isConstantColorEnabled

func (c_ CapturePhotoOutput) SetConstantColorEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstantColorEnabled:"), value)
}


// The currently active color space for capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activecolorspace

func (c_ CapturePhotoOutput) ActiveColorSpace() CaptureColorSpace {
	rv := objc.Send[CaptureColorSpace](c_.ID, objc.Sel("activeColorSpace"))
	return rv
}


// The currently active color space for capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activecolorspace

func (c_ CapturePhotoOutput) SetActiveColorSpace(value ICaptureColorSpace) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveColorSpace:"), value)
}


// The portrait effects matte captured with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/portraiteffectsmatte

func (c_ CapturePhotoOutput) PortraitEffectsMatte() AVPortraitEffectsMatte {
	rv := objc.Send[AVPortraitEffectsMatte](c_.ID, objc.Sel("portraitEffectsMatte"))
	return rv
}


// The portrait effects matte captured with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/portraiteffectsmatte

func (c_ CapturePhotoOutput) SetPortraitEffectsMatte(value IAVPortraitEffectsMatte) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPortraitEffectsMatte:"), value)
}


// An array of video codecs currently available for Live Photo movie captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablelivephotovideocodectypes

func (c_ CapturePhotoOutput) AvailableLivePhotoVideoCodecTypes() VideoCodecType {
	rv := objc.Send[VideoCodecType](c_.ID, objc.Sel("availableLivePhotoVideoCodecTypes"))
	return rv
}


// An array of video codecs currently available for Live Photo movie captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablelivephotovideocodectypes

func (c_ CapturePhotoOutput) SetAvailableLivePhotoVideoCodecTypes(value IVideoCodecType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableLivePhotoVideoCodecTypes:"), value)
}


// The compression codecs this capture output currently supports for photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotocodectypes

func (c_ CapturePhotoOutput) AvailablePhotoCodecTypes() VideoCodecType {
	rv := objc.Send[VideoCodecType](c_.ID, objc.Sel("availablePhotoCodecTypes"))
	return rv
}


// The compression codecs this capture output currently supports for photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotocodectypes

func (c_ CapturePhotoOutput) SetAvailablePhotoCodecTypes(value IVideoCodecType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailablePhotoCodecTypes:"), value)
}


// The list of file types currently supported for photo capture and output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotofiletypes

func (c_ CapturePhotoOutput) AvailablePhotoFileTypes() FileType {
	rv := objc.Send[FileType](c_.ID, objc.Sel("availablePhotoFileTypes"))
	return rv
}


// The list of file types currently supported for photo capture and output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotofiletypes

func (c_ CapturePhotoOutput) SetAvailablePhotoFileTypes(value IFileType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailablePhotoFileTypes:"), value)
}


// The pixel formats the capture output supports for photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotopixelformattypes-3ydgm

func (c_ CapturePhotoOutput) AvailablePhotoPixelFormatTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availablePhotoPixelFormatTypes"))
	return rv
}


// The pixel formats the capture output supports for photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotopixelformattypes-3ydgm

func (c_ CapturePhotoOutput) SetAvailablePhotoPixelFormatTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailablePhotoPixelFormatTypes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotocodectypes

func (c_ CapturePhotoOutput) AvailableRawPhotoCodecTypes() VideoCodecType {
	rv := objc.Send[VideoCodecType](c_.ID, objc.Sel("availableRawPhotoCodecTypes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotocodectypes

func (c_ CapturePhotoOutput) SetAvailableRawPhotoCodecTypes(value IVideoCodecType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableRawPhotoCodecTypes:"), value)
}


// The list of file types currently supported for RAW format capture and output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotofiletypes

func (c_ CapturePhotoOutput) AvailableRawPhotoFileTypes() FileType {
	rv := objc.Send[FileType](c_.ID, objc.Sel("availableRawPhotoFileTypes"))
	return rv
}


// The list of file types currently supported for RAW format capture and output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotofiletypes

func (c_ CapturePhotoOutput) SetAvailableRawPhotoFileTypes(value IFileType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableRawPhotoFileTypes:"), value)
}


// The pixel formats the capture output supports for RAW photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotopixelformattypes-9t9k5

func (c_ CapturePhotoOutput) AvailableRawPhotoPixelFormatTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableRawPhotoPixelFormatTypes"))
	return rv
}


// The pixel formats the capture output supports for RAW photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotopixelformattypes-9t9k5

func (c_ CapturePhotoOutput) SetAvailableRawPhotoPixelFormatTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableRawPhotoPixelFormatTypes:"), value)
}


// An array of semantic segmentation matte types that may be captured and delivered along with the primary photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablesemanticsegmentationmattetypes

func (c_ CapturePhotoOutput) AvailableSemanticSegmentationMatteTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableSemanticSegmentationMatteTypes"))
	return rv
}


// An array of semantic segmentation matte types that may be captured and delivered along with the primary photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablesemanticsegmentationmattetypes

func (c_ CapturePhotoOutput) SetAvailableSemanticSegmentationMatteTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableSemanticSegmentationMatteTypes:"), value)
}


// A value that specifies whether the photo output is ready to respond to new capture requests in a timely manner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/capturereadiness-swift.property

func (c_ CapturePhotoOutput) CaptureReadiness() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("captureReadiness"))
	return rv
}


// A value that specifies whether the photo output is ready to respond to new capture requests in a timely manner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/capturereadiness-swift.property

func (c_ CapturePhotoOutput) SetCaptureReadiness(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCaptureReadiness:"), value)
}


// The semantic segmentation matte types that the photo render pipeline delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/enabledsemanticsegmentationmattetypes

func (c_ CapturePhotoOutput) EnabledSemanticSegmentationMatteTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("enabledSemanticSegmentationMatteTypes"))
	return rv
}


// The semantic segmentation matte types that the photo render pipeline delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/enabledsemanticsegmentationmattetypes

func (c_ CapturePhotoOutput) SetEnabledSemanticSegmentationMatteTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabledSemanticSegmentationMatteTypes:"), value)
}


// A Boolean value that indicates whether you’ve configured the photo output to deliver Apple ProRAW formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isappleprorawenabled

func (c_ CapturePhotoOutput) IsAppleProRAWEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAppleProRAWEnabled"))
	return rv
}


// A Boolean value that indicates whether you’ve configured the photo output to deliver Apple ProRAW formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isappleprorawenabled

func (c_ CapturePhotoOutput) SetIsAppleProRAWEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAppleProRAWEnabled:"), value)
}


// A Boolean value that indicates whether the current device and configuration supports Apple ProRAW pixel formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isappleprorawsupported

func (c_ CapturePhotoOutput) IsAppleProRAWSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAppleProRAWSupported"))
	return rv
}


// A Boolean value that indicates whether the current device and configuration supports Apple ProRAW pixel formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isappleprorawsupported

func (c_ CapturePhotoOutput) SetIsAppleProRAWSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAppleProRAWSupported:"), value)
}


// A Boolean value that indicates the enabled state of automatic deferred photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautodeferredphotodeliveryenabled

func (c_ CapturePhotoOutput) IsAutoDeferredPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoDeferredPhotoDeliveryEnabled"))
	return rv
}


// A Boolean value that indicates the enabled state of automatic deferred photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautodeferredphotodeliveryenabled

func (c_ CapturePhotoOutput) SetIsAutoDeferredPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoDeferredPhotoDeliveryEnabled:"), value)
}


// A Boolean value that indicates whether the photo output supports deferred photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautodeferredphotodeliverysupported

func (c_ CapturePhotoOutput) IsAutoDeferredPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoDeferredPhotoDeliverySupported"))
	return rv
}


// A Boolean value that indicates whether the photo output supports deferred photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautodeferredphotodeliverysupported

func (c_ CapturePhotoOutput) SetIsAutoDeferredPhotoDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoDeferredPhotoDeliverySupported:"), value)
}


// A Boolean value indicating whether the capture output supports automatic red-eye reduction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautoredeyereductionsupported

func (c_ CapturePhotoOutput) IsAutoRedEyeReductionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoRedEyeReductionSupported"))
	return rv
}


// A Boolean value indicating whether the capture output supports automatic red-eye reduction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautoredeyereductionsupported

func (c_ CapturePhotoOutput) SetIsAutoRedEyeReductionSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoRedEyeReductionSupported:"), value)
}


// A Boolean value indicating whether the capture output currently supports delivery of camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscameracalibrationdatadeliverysupported

func (c_ CapturePhotoOutput) IsCameraCalibrationDataDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraCalibrationDataDeliverySupported"))
	return rv
}


// A Boolean value indicating whether the capture output currently supports delivery of camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscameracalibrationdatadeliverysupported

func (c_ CapturePhotoOutput) SetIsCameraCalibrationDataDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraCalibrationDataDeliverySupported:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscamerasensororientationcompensationenabled

func (c_ CapturePhotoOutput) IsCameraSensorOrientationCompensationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraSensorOrientationCompensationEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscamerasensororientationcompensationenabled

func (c_ CapturePhotoOutput) SetIsCameraSensorOrientationCompensationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraSensorOrientationCompensationEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscamerasensororientationcompensationsupported

func (c_ CapturePhotoOutput) IsCameraSensorOrientationCompensationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraSensorOrientationCompensationSupported"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscamerasensororientationcompensationsupported

func (c_ CapturePhotoOutput) SetIsCameraSensorOrientationCompensationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraSensorOrientationCompensationSupported:"), value)
}


// A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isconstantcolorenabled

func (c_ CapturePhotoOutput) IsConstantColorEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorEnabled"))
	return rv
}


// A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isconstantcolorenabled

func (c_ CapturePhotoOutput) SetIsConstantColorEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorEnabled:"), value)
}


// A Boolean value that indicates whether a photo output supports constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isconstantcolorsupported

func (c_ CapturePhotoOutput) IsConstantColorSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorSupported"))
	return rv
}


// A Boolean value that indicates whether a photo output supports constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isconstantcolorsupported

func (c_ CapturePhotoOutput) SetIsConstantColorSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorSupported:"), value)
}


// A Boolean value that indicates whether the photo render pipeline can perform content-aware distortion correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscontentawaredistortioncorrectionenabled

func (c_ CapturePhotoOutput) IsContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContentAwareDistortionCorrectionEnabled"))
	return rv
}


// A Boolean value that indicates whether the photo render pipeline can perform content-aware distortion correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscontentawaredistortioncorrectionenabled

func (c_ CapturePhotoOutput) SetIsContentAwareDistortionCorrectionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContentAwareDistortionCorrectionEnabled:"), value)
}


// A Boolean value that indicates whether the session’s current configuration supports content-aware distortion correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscontentawaredistortioncorrectionsupported

func (c_ CapturePhotoOutput) IsContentAwareDistortionCorrectionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContentAwareDistortionCorrectionSupported"))
	return rv
}


// A Boolean value that indicates whether the session’s current configuration supports content-aware distortion correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscontentawaredistortioncorrectionsupported

func (c_ CapturePhotoOutput) SetIsContentAwareDistortionCorrectionSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContentAwareDistortionCorrectionSupported:"), value)
}


// A Boolean value that specifies whether to configure the capture pipeline for depth data capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isdepthdatadeliveryenabled

func (c_ CapturePhotoOutput) IsDepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataDeliveryEnabled"))
	return rv
}


// A Boolean value that specifies whether to configure the capture pipeline for depth data capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isdepthdatadeliveryenabled

func (c_ CapturePhotoOutput) SetIsDepthDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataDeliveryEnabled:"), value)
}


// A Boolean value indicating whether the capture output currently supports depth data capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isdepthdatadeliverysupported

func (c_ CapturePhotoOutput) IsDepthDataDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataDeliverySupported"))
	return rv
}


// A Boolean value indicating whether the capture output currently supports depth data capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isdepthdatadeliverysupported

func (c_ CapturePhotoOutput) SetIsDepthDataDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataDeliverySupported:"), value)
}


// A Boolean value that indicates whether the output enables fast capture prioritization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isfastcaptureprioritizationenabled

func (c_ CapturePhotoOutput) IsFastCapturePrioritizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFastCapturePrioritizationEnabled"))
	return rv
}


// A Boolean value that indicates whether the output enables fast capture prioritization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isfastcaptureprioritizationenabled

func (c_ CapturePhotoOutput) SetIsFastCapturePrioritizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFastCapturePrioritizationEnabled:"), value)
}


// A Boolean value that indicates whether the photo output supports fast capture prioritization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isfastcaptureprioritizationsupported

func (c_ CapturePhotoOutput) IsFastCapturePrioritizationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFastCapturePrioritizationSupported"))
	return rv
}


// A Boolean value that indicates whether the photo output supports fast capture prioritization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isfastcaptureprioritizationsupported

func (c_ CapturePhotoOutput) SetIsFastCapturePrioritizationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFastCapturePrioritizationSupported:"), value)
}


// A Boolean value indicating whether the scene currently being previewed by the camera warrants use of the flash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isflashscene

func (c_ CapturePhotoOutput) IsFlashScene() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFlashScene"))
	return rv
}


// A Boolean value indicating whether the scene currently being previewed by the camera warrants use of the flash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isflashscene

func (c_ CapturePhotoOutput) SetIsFlashScene(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFlashScene:"), value)
}


// A Boolean value that specifies whether to configure the capture pipeline for high resolution still image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/ishighresolutioncaptureenabled

func (c_ CapturePhotoOutput) IsHighResolutionCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighResolutionCaptureEnabled"))
	return rv
}


// A Boolean value that specifies whether to configure the capture pipeline for high resolution still image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/ishighresolutioncaptureenabled

func (c_ CapturePhotoOutput) SetIsHighResolutionCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighResolutionCaptureEnabled:"), value)
}


// A Boolean value indicating whether the capture output currently supports lens stabilization during bracketed image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islensstabilizationduringbracketedcapturesupported

func (c_ CapturePhotoOutput) IsLensStabilizationDuringBracketedCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLensStabilizationDuringBracketedCaptureSupported"))
	return rv
}


// A Boolean value indicating whether the capture output currently supports lens stabilization during bracketed image capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islensstabilizationduringbracketedcapturesupported

func (c_ CapturePhotoOutput) SetIsLensStabilizationDuringBracketedCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLensStabilizationDuringBracketedCaptureSupported:"), value)
}


// A Boolean value that indicates whether to automatically trim Live Photo movie captures to avoid excessive movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotoautotrimmingenabled

func (c_ CapturePhotoOutput) IsLivePhotoAutoTrimmingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLivePhotoAutoTrimmingEnabled"))
	return rv
}


// A Boolean value that indicates whether to automatically trim Live Photo movie captures to avoid excessive movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotoautotrimmingenabled

func (c_ CapturePhotoOutput) SetIsLivePhotoAutoTrimmingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLivePhotoAutoTrimmingEnabled:"), value)
}


// A Boolean value that indicates whether to configure the capture pipeline for Live Photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocaptureenabled

func (c_ CapturePhotoOutput) IsLivePhotoCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLivePhotoCaptureEnabled"))
	return rv
}


// A Boolean value that indicates whether to configure the capture pipeline for Live Photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocaptureenabled

func (c_ CapturePhotoOutput) SetIsLivePhotoCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLivePhotoCaptureEnabled:"), value)
}


// A Boolean value that indicates whether the capture output currently supports Live Photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocapturesupported

func (c_ CapturePhotoOutput) IsLivePhotoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLivePhotoCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether the capture output currently supports Live Photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocapturesupported

func (c_ CapturePhotoOutput) SetIsLivePhotoCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLivePhotoCaptureSupported:"), value)
}


// A Boolean value that indicates whether Live Photo capture is currently in a suspended state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocapturesuspended

func (c_ CapturePhotoOutput) IsLivePhotoCaptureSuspended() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLivePhotoCaptureSuspended"))
	return rv
}


// A Boolean value that indicates whether Live Photo capture is currently in a suspended state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocapturesuspended

func (c_ CapturePhotoOutput) SetIsLivePhotoCaptureSuspended(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLivePhotoCaptureSuspended:"), value)
}


// A Boolean value indicating whether the capture output generates a portrait effects matte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isportraiteffectsmattedeliveryenabled

func (c_ CapturePhotoOutput) IsPortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectsMatteDeliveryEnabled"))
	return rv
}


// A Boolean value indicating whether the capture output generates a portrait effects matte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isportraiteffectsmattedeliveryenabled

func (c_ CapturePhotoOutput) SetIsPortraitEffectsMatteDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectsMatteDeliveryEnabled:"), value)
}


// A Boolean value indicating whether the capture output currently supports delivery of a portrait effects matte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isportraiteffectsmattedeliverysupported

func (c_ CapturePhotoOutput) IsPortraitEffectsMatteDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectsMatteDeliverySupported"))
	return rv
}


// A Boolean value indicating whether the capture output currently supports delivery of a portrait effects matte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isportraiteffectsmattedeliverysupported

func (c_ CapturePhotoOutput) SetIsPortraitEffectsMatteDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectsMatteDeliverySupported:"), value)
}


// A Boolean value that indicates whether the photo output configuration enables responsive capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isresponsivecaptureenabled

func (c_ CapturePhotoOutput) IsResponsiveCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isResponsiveCaptureEnabled"))
	return rv
}


// A Boolean value that indicates whether the photo output configuration enables responsive capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isresponsivecaptureenabled

func (c_ CapturePhotoOutput) SetIsResponsiveCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsResponsiveCaptureEnabled:"), value)
}


// A Boolean value that indicates whether the photo output supports responsive capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isresponsivecapturesupported

func (c_ CapturePhotoOutput) IsResponsiveCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isResponsiveCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether the photo output supports responsive capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isresponsivecapturesupported

func (c_ CapturePhotoOutput) SetIsResponsiveCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsResponsiveCaptureSupported:"), value)
}


// A Boolean value that indicates whether the photo output supports suppressing the system shutter sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isshuttersoundsuppressionsupported

func (c_ CapturePhotoOutput) IsShutterSoundSuppressionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isShutterSoundSuppressionSupported"))
	return rv
}


// A Boolean value that indicates whether the photo output supports suppressing the system shutter sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isshuttersoundsuppressionsupported

func (c_ CapturePhotoOutput) SetIsShutterSoundSuppressionSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsShutterSoundSuppressionSupported:"), value)
}


// A Boolean value that indicates whether the photo output delivers photos from constituent cameras of a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdeviceconstituentphotodeliveryenabled

func (c_ CapturePhotoOutput) IsVirtualDeviceConstituentPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDeviceConstituentPhotoDeliveryEnabled"))
	return rv
}


// A Boolean value that indicates whether the photo output delivers photos from constituent cameras of a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdeviceconstituentphotodeliveryenabled

func (c_ CapturePhotoOutput) SetIsVirtualDeviceConstituentPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDeviceConstituentPhotoDeliveryEnabled:"), value)
}


// A Boolean value that indicates whether the photo output configuration supports delivery of photos from constituent cameras of a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdeviceconstituentphotodeliverysupported

func (c_ CapturePhotoOutput) IsVirtualDeviceConstituentPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDeviceConstituentPhotoDeliverySupported"))
	return rv
}


// A Boolean value that indicates whether the photo output configuration supports delivery of photos from constituent cameras of a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdeviceconstituentphotodeliverysupported

func (c_ CapturePhotoOutput) SetIsVirtualDeviceConstituentPhotoDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDeviceConstituentPhotoDeliverySupported:"), value)
}


// A Boolean value that indicates whether the device supports virtual device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdevicefusionsupported

func (c_ CapturePhotoOutput) IsVirtualDeviceFusionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDeviceFusionSupported"))
	return rv
}


// A Boolean value that indicates whether the device supports virtual device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdevicefusionsupported

func (c_ CapturePhotoOutput) SetIsVirtualDeviceFusionSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDeviceFusionSupported:"), value)
}


// A Boolean value that indicates whether the photo output configuration enables zero shutter lag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iszeroshutterlagenabled

func (c_ CapturePhotoOutput) IsZeroShutterLagEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isZeroShutterLagEnabled"))
	return rv
}


// A Boolean value that indicates whether the photo output configuration enables zero shutter lag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iszeroshutterlagenabled

func (c_ CapturePhotoOutput) SetIsZeroShutterLagEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsZeroShutterLagEnabled:"), value)
}


// A Boolean value that indicates whether the photo output supports zero shutter lag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iszeroshutterlagsupported

func (c_ CapturePhotoOutput) IsZeroShutterLagSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isZeroShutterLagSupported"))
	return rv
}


// A Boolean value that indicates whether the photo output supports zero shutter lag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iszeroshutterlagsupported

func (c_ CapturePhotoOutput) SetIsZeroShutterLagSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsZeroShutterLagSupported:"), value)
}


// The maximum number of images that the photo capture output can support in a single bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxbracketedcapturephotocount

func (c_ CapturePhotoOutput) MaxBracketedCapturePhotoCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxBracketedCapturePhotoCount"))
	return rv
}


// The maximum number of images that the photo capture output can support in a single bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxbracketedcapturephotocount

func (c_ CapturePhotoOutput) SetMaxBracketedCapturePhotoCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxBracketedCapturePhotoCount:"), value)
}


// The maximum resolution of the requested photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxphotodimensions

func (c_ CapturePhotoOutput) MaxPhotoDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("maxPhotoDimensions"))
	return rv
}


// The maximum resolution of the requested photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxphotodimensions

func (c_ CapturePhotoOutput) SetMaxPhotoDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxPhotoDimensions:"), value)
}


// The highest quality the photo output should prepare to deliver on a capture-by-capture basis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxphotoqualityprioritization

func (c_ CapturePhotoOutput) MaxPhotoQualityPrioritization() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("maxPhotoQualityPrioritization"))
	return rv
}


// The highest quality the photo output should prepare to deliver on a capture-by-capture basis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxphotoqualityprioritization

func (c_ CapturePhotoOutput) SetMaxPhotoQualityPrioritization(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxPhotoQualityPrioritization:"), value)
}


// A photo settings object that controls how the photo output detects and handles automatic flash and stabilization modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/photosettingsforscenemonitoring

func (c_ CapturePhotoOutput) PhotoSettingsForSceneMonitoring() AVCapturePhotoSettings {
	rv := objc.Send[AVCapturePhotoSettings](c_.ID, objc.Sel("photoSettingsForSceneMonitoring"))
	return rv
}


// A photo settings object that controls how the photo output detects and handles automatic flash and stabilization modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/photosettingsforscenemonitoring

func (c_ CapturePhotoOutput) SetPhotoSettingsForSceneMonitoring(value IAVCapturePhotoSettings) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhotoSettingsForSceneMonitoring:"), value)
}


// An array of photo settings for which the photo output has prepared capture resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/preparedphotosettingsarray

func (c_ CapturePhotoOutput) PreparedPhotoSettingsArray() AVCapturePhotoSettings {
	rv := objc.Send[AVCapturePhotoSettings](c_.ID, objc.Sel("preparedPhotoSettingsArray"))
	return rv
}


// An array of photo settings for which the photo output has prepared capture resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/preparedphotosettingsarray

func (c_ CapturePhotoOutput) SetPreparedPhotoSettingsArray(value IAVCapturePhotoSettings) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreparedPhotoSettingsArray:"), value)
}


// A Boolean value that indicates whether to preserve the suspended state of Live Photo capture when the session stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/preserveslivephotocapturesuspendedonsessionstop

func (c_ CapturePhotoOutput) PreservesLivePhotoCaptureSuspendedOnSessionStop() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preservesLivePhotoCaptureSuspendedOnSessionStop"))
	return rv
}


// A Boolean value that indicates whether to preserve the suspended state of Live Photo capture when the session stops.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/preserveslivephotocapturesuspendedonsessionstop

func (c_ CapturePhotoOutput) SetPreservesLivePhotoCaptureSuspendedOnSessionStop(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreservesLivePhotoCaptureSuspendedOnSessionStop:"), value)
}


// A Swift array of flash settings this capture output currently supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/supportedflashmodes-1n6nm

func (c_ CapturePhotoOutput) SupportedFlashModes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supportedFlashModes"))
	return rv
}


// A Swift array of flash settings this capture output currently supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/supportedflashmodes-1n6nm

func (c_ CapturePhotoOutput) SetSupportedFlashModes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedFlashModes:"), value)
}


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode

func (c_ CapturePhotoOutput) FlashMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("flashMode"))
	return rv
}


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode

func (c_ CapturePhotoOutput) SetFlashMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashMode:"), value)
}


// A unique identifier for this photo settings instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/uniqueid

func (c_ CapturePhotoOutput) UniqueID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("uniqueID"))
	return rv
}


// A unique identifier for this photo settings instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/uniqueid

func (c_ CapturePhotoOutput) SetUniqueID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUniqueID:"), value)
}



