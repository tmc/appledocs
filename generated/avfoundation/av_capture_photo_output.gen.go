// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	CapturePhotoWithSettingsDelegate(settings unsafe.Pointer, delegate objc.ID)
}

// A capture output for still image, Live Photos, and other photography workflows.
//
// provides an interface for capture workflows related to still photography. In addition to basic capture of still images, a photo output supports RAW-format capture, bracketed capture of multiple images, Live Photos, and wide-gamut color. You can output captured photos in a variety of formats and codecs, including RAW format DNG files, HEVC format HEIF files, and JPEG files. To capture photos with the class, follow these steps: Create an object. Use its properties to determine supported capture settings and to enable certain features (for example, whether to capture Live Photos). Create and configure an object to choose features and settings for a specific capture (for example, whether to enable image stabilization or flash). Capture an image by passing your photo settings object to the method along with a delegate object implementing the protocol. The photo capture output then calls your delegate to notify you of significant events during the capture process. Some photo capture settings, such as the property, include options for automatic behavior. For such settings, the photo output determines whether to use that feature at the moment of capture—you don’t know when requesting a capture whether the feature will be enabled when the capture completes. When the photo capture output calls your methods with information about the completed or in-progress capture, it also provides an object that details which automatic features are set for that capture. The resolved settings object’s property matches the value of the object you used to request capture. Enabling certain photo features (Live Photo capture and high resolution capture) requires a reconfiguration of the capture render pipeline. To opt into these features, set the , , and properties before calling your object’s method. Changing any of these properties while the session is running disrupts the capture render pipeline: Live Photo captures in progress end immediately, unfulfilled photo requests abort, and video preview temporarily freezes. Using a photo capture output adds other requirements to your object: A capture session can’t support both Live Photo capture and movie file output. If your capture session includes an object, the property becomes . (As an alternative, you can use the class to output video buffers at the same resolution as a simultaneous Live Photo capture). A capture session can’t contain both an object and an object. The class includes all functionality of (and deprecates) the class. The class implicitly supports wide-gamut color photography. If the source object’s value is , the capture output produces photos with wide color information (unless your object specifies an output format that doesn’t support wide color).
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/capturePhoto(with:delegate:)
func (c_ CapturePhotoOutput) CapturePhotoWithSettingsDelegate(settings unsafe.Pointer, delegate objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("capturePhotoWithSettings:delegate:"), settings, delegate)
}

// A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isConstantColorEnabled
func (c_ CapturePhotoOutput) ConstantColorEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("constantColorEnabled"))
	return rv
}


// SetConstantColorEnabled sets the value of the constantColorEnabled property.
// A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isConstantColorEnabled
func (c_ CapturePhotoOutput) SetConstantColorEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstantColorEnabled:"), value)
}

// The currently active color space for capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activecolorspace
func (c_ CapturePhotoOutput) ActiveColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeColorSpace"))
	return rv
}


// SetActiveColorSpace sets the value of the activeColorSpace property.
// The currently active color space for capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activecolorspace
func (c_ CapturePhotoOutput) SetActiveColorSpace(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveColorSpace:"), value)
}

// The portrait effects matte captured with the photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/portraiteffectsmatte
func (c_ CapturePhotoOutput) PortraitEffectsMatte() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("portraitEffectsMatte"))
	return rv
}


// SetPortraitEffectsMatte sets the value of the portraitEffectsMatte property.
// The portrait effects matte captured with the photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/portraiteffectsmatte
func (c_ CapturePhotoOutput) SetPortraitEffectsMatte(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPortraitEffectsMatte:"), value)
}

// An array of video codecs currently available for Live Photo movie captures.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablelivephotovideocodectypes
func (c_ CapturePhotoOutput) AvailableLivePhotoVideoCodecTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableLivePhotoVideoCodecTypes"))
	return rv
}


// SetAvailableLivePhotoVideoCodecTypes sets the value of the availableLivePhotoVideoCodecTypes property.
// An array of video codecs currently available for Live Photo movie captures.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablelivephotovideocodectypes
func (c_ CapturePhotoOutput) SetAvailableLivePhotoVideoCodecTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableLivePhotoVideoCodecTypes:"), value)
}

// The compression codecs this capture output currently supports for photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotocodectypes
func (c_ CapturePhotoOutput) AvailablePhotoCodecTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availablePhotoCodecTypes"))
	return rv
}


// SetAvailablePhotoCodecTypes sets the value of the availablePhotoCodecTypes property.
// The compression codecs this capture output currently supports for photo capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotocodectypes
func (c_ CapturePhotoOutput) SetAvailablePhotoCodecTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailablePhotoCodecTypes:"), value)
}

// The list of file types currently supported for photo capture and output.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotofiletypes
func (c_ CapturePhotoOutput) AvailablePhotoFileTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availablePhotoFileTypes"))
	return rv
}


// SetAvailablePhotoFileTypes sets the value of the availablePhotoFileTypes property.
// The list of file types currently supported for photo capture and output.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotofiletypes
func (c_ CapturePhotoOutput) SetAvailablePhotoFileTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailablePhotoFileTypes:"), value)
}

// The pixel formats the capture output supports for photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotopixelformattypes-3ydgm
func (c_ CapturePhotoOutput) AvailablePhotoPixelFormatTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availablePhotoPixelFormatTypes"))
	return rv
}


// SetAvailablePhotoPixelFormatTypes sets the value of the availablePhotoPixelFormatTypes property.
// The pixel formats the capture output supports for photo capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablephotopixelformattypes-3ydgm
func (c_ CapturePhotoOutput) SetAvailablePhotoPixelFormatTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailablePhotoPixelFormatTypes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotocodectypes
func (c_ CapturePhotoOutput) AvailableRawPhotoCodecTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableRawPhotoCodecTypes"))
	return rv
}


// SetAvailableRawPhotoCodecTypes sets the value of the availableRawPhotoCodecTypes property.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotocodectypes
func (c_ CapturePhotoOutput) SetAvailableRawPhotoCodecTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableRawPhotoCodecTypes:"), value)
}

// The list of file types currently supported for RAW format capture and output.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotofiletypes
func (c_ CapturePhotoOutput) AvailableRawPhotoFileTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableRawPhotoFileTypes"))
	return rv
}


// SetAvailableRawPhotoFileTypes sets the value of the availableRawPhotoFileTypes property.
// The list of file types currently supported for RAW format capture and output.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotofiletypes
func (c_ CapturePhotoOutput) SetAvailableRawPhotoFileTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableRawPhotoFileTypes:"), value)
}

// The pixel formats the capture output supports for RAW photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotopixelformattypes-9t9k5
func (c_ CapturePhotoOutput) AvailableRawPhotoPixelFormatTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableRawPhotoPixelFormatTypes"))
	return rv
}


// SetAvailableRawPhotoPixelFormatTypes sets the value of the availableRawPhotoPixelFormatTypes property.
// The pixel formats the capture output supports for RAW photo capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablerawphotopixelformattypes-9t9k5
func (c_ CapturePhotoOutput) SetAvailableRawPhotoPixelFormatTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableRawPhotoPixelFormatTypes:"), value)
}

// An array of semantic segmentation matte types that may be captured and delivered along with the primary photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablesemanticsegmentationmattetypes
func (c_ CapturePhotoOutput) AvailableSemanticSegmentationMatteTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableSemanticSegmentationMatteTypes"))
	return rv
}


// SetAvailableSemanticSegmentationMatteTypes sets the value of the availableSemanticSegmentationMatteTypes property.
// An array of semantic segmentation matte types that may be captured and delivered along with the primary photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/availablesemanticsegmentationmattetypes
func (c_ CapturePhotoOutput) SetAvailableSemanticSegmentationMatteTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableSemanticSegmentationMatteTypes:"), value)
}

// A value that specifies whether the photo output is ready to respond to new capture requests in a timely manner.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/capturereadiness-swift.property
func (c_ CapturePhotoOutput) CaptureReadiness() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("captureReadiness"))
	return rv
}


// SetCaptureReadiness sets the value of the captureReadiness property.
// A value that specifies whether the photo output is ready to respond to new capture requests in a timely manner.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/capturereadiness-swift.property
func (c_ CapturePhotoOutput) SetCaptureReadiness(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCaptureReadiness:"), value)
}

// The semantic segmentation matte types that the photo render pipeline delivers.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/enabledsemanticsegmentationmattetypes
func (c_ CapturePhotoOutput) EnabledSemanticSegmentationMatteTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("enabledSemanticSegmentationMatteTypes"))
	return rv
}


// SetEnabledSemanticSegmentationMatteTypes sets the value of the enabledSemanticSegmentationMatteTypes property.
// The semantic segmentation matte types that the photo render pipeline delivers.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/enabledsemanticsegmentationmattetypes
func (c_ CapturePhotoOutput) SetEnabledSemanticSegmentationMatteTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabledSemanticSegmentationMatteTypes:"), value)
}

// A Boolean value that indicates whether you’ve configured the photo output to deliver Apple ProRAW formats.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isappleprorawenabled
func (c_ CapturePhotoOutput) IsAppleProRAWEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAppleProRAWEnabled"))
	return rv
}


// SetIsAppleProRAWEnabled sets the value of the isAppleProRAWEnabled property.
// A Boolean value that indicates whether you’ve configured the photo output to deliver Apple ProRAW formats.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isappleprorawenabled
func (c_ CapturePhotoOutput) SetIsAppleProRAWEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAppleProRAWEnabled:"), value)
}

// A Boolean value that indicates whether the current device and configuration supports Apple ProRAW pixel formats.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isappleprorawsupported
func (c_ CapturePhotoOutput) IsAppleProRAWSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAppleProRAWSupported"))
	return rv
}


// SetIsAppleProRAWSupported sets the value of the isAppleProRAWSupported property.
// A Boolean value that indicates whether the current device and configuration supports Apple ProRAW pixel formats.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isappleprorawsupported
func (c_ CapturePhotoOutput) SetIsAppleProRAWSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAppleProRAWSupported:"), value)
}

// A Boolean value that indicates the enabled state of automatic deferred photo delivery.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautodeferredphotodeliveryenabled
func (c_ CapturePhotoOutput) IsAutoDeferredPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoDeferredPhotoDeliveryEnabled"))
	return rv
}


// SetIsAutoDeferredPhotoDeliveryEnabled sets the value of the isAutoDeferredPhotoDeliveryEnabled property.
// A Boolean value that indicates the enabled state of automatic deferred photo delivery.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautodeferredphotodeliveryenabled
func (c_ CapturePhotoOutput) SetIsAutoDeferredPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoDeferredPhotoDeliveryEnabled:"), value)
}

// A Boolean value that indicates whether the photo output supports deferred photo delivery.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautodeferredphotodeliverysupported
func (c_ CapturePhotoOutput) IsAutoDeferredPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoDeferredPhotoDeliverySupported"))
	return rv
}


// SetIsAutoDeferredPhotoDeliverySupported sets the value of the isAutoDeferredPhotoDeliverySupported property.
// A Boolean value that indicates whether the photo output supports deferred photo delivery.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautodeferredphotodeliverysupported
func (c_ CapturePhotoOutput) SetIsAutoDeferredPhotoDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoDeferredPhotoDeliverySupported:"), value)
}

// A Boolean value indicating whether the capture output supports automatic red-eye reduction.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautoredeyereductionsupported
func (c_ CapturePhotoOutput) IsAutoRedEyeReductionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoRedEyeReductionSupported"))
	return rv
}


// SetIsAutoRedEyeReductionSupported sets the value of the isAutoRedEyeReductionSupported property.
// A Boolean value indicating whether the capture output supports automatic red-eye reduction.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isautoredeyereductionsupported
func (c_ CapturePhotoOutput) SetIsAutoRedEyeReductionSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoRedEyeReductionSupported:"), value)
}

// A Boolean value indicating whether the capture output currently supports delivery of camera calibration data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscameracalibrationdatadeliverysupported
func (c_ CapturePhotoOutput) IsCameraCalibrationDataDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraCalibrationDataDeliverySupported"))
	return rv
}


// SetIsCameraCalibrationDataDeliverySupported sets the value of the isCameraCalibrationDataDeliverySupported property.
// A Boolean value indicating whether the capture output currently supports delivery of camera calibration data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscameracalibrationdatadeliverysupported
func (c_ CapturePhotoOutput) SetIsCameraCalibrationDataDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraCalibrationDataDeliverySupported:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscamerasensororientationcompensationenabled
func (c_ CapturePhotoOutput) IsCameraSensorOrientationCompensationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraSensorOrientationCompensationEnabled"))
	return rv
}


// SetIsCameraSensorOrientationCompensationEnabled sets the value of the isCameraSensorOrientationCompensationEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscamerasensororientationcompensationenabled
func (c_ CapturePhotoOutput) SetIsCameraSensorOrientationCompensationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraSensorOrientationCompensationEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscamerasensororientationcompensationsupported
func (c_ CapturePhotoOutput) IsCameraSensorOrientationCompensationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraSensorOrientationCompensationSupported"))
	return rv
}


// SetIsCameraSensorOrientationCompensationSupported sets the value of the isCameraSensorOrientationCompensationSupported property.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscamerasensororientationcompensationsupported
func (c_ CapturePhotoOutput) SetIsCameraSensorOrientationCompensationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraSensorOrientationCompensationSupported:"), value)
}

// A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isconstantcolorenabled
func (c_ CapturePhotoOutput) IsConstantColorEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorEnabled"))
	return rv
}


// SetIsConstantColorEnabled sets the value of the isConstantColorEnabled property.
// A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isconstantcolorenabled
func (c_ CapturePhotoOutput) SetIsConstantColorEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorEnabled:"), value)
}

// A Boolean value that indicates whether a photo output supports constant color capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isconstantcolorsupported
func (c_ CapturePhotoOutput) IsConstantColorSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorSupported"))
	return rv
}


// SetIsConstantColorSupported sets the value of the isConstantColorSupported property.
// A Boolean value that indicates whether a photo output supports constant color capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isconstantcolorsupported
func (c_ CapturePhotoOutput) SetIsConstantColorSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorSupported:"), value)
}

// A Boolean value that indicates whether the photo render pipeline can perform content-aware distortion correction.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscontentawaredistortioncorrectionenabled
func (c_ CapturePhotoOutput) IsContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContentAwareDistortionCorrectionEnabled"))
	return rv
}


// SetIsContentAwareDistortionCorrectionEnabled sets the value of the isContentAwareDistortionCorrectionEnabled property.
// A Boolean value that indicates whether the photo render pipeline can perform content-aware distortion correction.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscontentawaredistortioncorrectionenabled
func (c_ CapturePhotoOutput) SetIsContentAwareDistortionCorrectionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContentAwareDistortionCorrectionEnabled:"), value)
}

// A Boolean value that indicates whether the session’s current configuration supports content-aware distortion correction.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscontentawaredistortioncorrectionsupported
func (c_ CapturePhotoOutput) IsContentAwareDistortionCorrectionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContentAwareDistortionCorrectionSupported"))
	return rv
}


// SetIsContentAwareDistortionCorrectionSupported sets the value of the isContentAwareDistortionCorrectionSupported property.
// A Boolean value that indicates whether the session’s current configuration supports content-aware distortion correction.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iscontentawaredistortioncorrectionsupported
func (c_ CapturePhotoOutput) SetIsContentAwareDistortionCorrectionSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContentAwareDistortionCorrectionSupported:"), value)
}

// A Boolean value that specifies whether to configure the capture pipeline for depth data capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isdepthdatadeliveryenabled
func (c_ CapturePhotoOutput) IsDepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataDeliveryEnabled"))
	return rv
}


// SetIsDepthDataDeliveryEnabled sets the value of the isDepthDataDeliveryEnabled property.
// A Boolean value that specifies whether to configure the capture pipeline for depth data capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isdepthdatadeliveryenabled
func (c_ CapturePhotoOutput) SetIsDepthDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataDeliveryEnabled:"), value)
}

// A Boolean value indicating whether the capture output currently supports depth data capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isdepthdatadeliverysupported
func (c_ CapturePhotoOutput) IsDepthDataDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataDeliverySupported"))
	return rv
}


// SetIsDepthDataDeliverySupported sets the value of the isDepthDataDeliverySupported property.
// A Boolean value indicating whether the capture output currently supports depth data capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isdepthdatadeliverysupported
func (c_ CapturePhotoOutput) SetIsDepthDataDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataDeliverySupported:"), value)
}

// A Boolean value that indicates whether the output enables fast capture prioritization.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isfastcaptureprioritizationenabled
func (c_ CapturePhotoOutput) IsFastCapturePrioritizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFastCapturePrioritizationEnabled"))
	return rv
}


// SetIsFastCapturePrioritizationEnabled sets the value of the isFastCapturePrioritizationEnabled property.
// A Boolean value that indicates whether the output enables fast capture prioritization.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isfastcaptureprioritizationenabled
func (c_ CapturePhotoOutput) SetIsFastCapturePrioritizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFastCapturePrioritizationEnabled:"), value)
}

// A Boolean value that indicates whether the photo output supports fast capture prioritization.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isfastcaptureprioritizationsupported
func (c_ CapturePhotoOutput) IsFastCapturePrioritizationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFastCapturePrioritizationSupported"))
	return rv
}


// SetIsFastCapturePrioritizationSupported sets the value of the isFastCapturePrioritizationSupported property.
// A Boolean value that indicates whether the photo output supports fast capture prioritization.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isfastcaptureprioritizationsupported
func (c_ CapturePhotoOutput) SetIsFastCapturePrioritizationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFastCapturePrioritizationSupported:"), value)
}

// A Boolean value indicating whether the scene currently being previewed by the camera warrants use of the flash.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isflashscene
func (c_ CapturePhotoOutput) IsFlashScene() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFlashScene"))
	return rv
}


// SetIsFlashScene sets the value of the isFlashScene property.
// A Boolean value indicating whether the scene currently being previewed by the camera warrants use of the flash.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isflashscene
func (c_ CapturePhotoOutput) SetIsFlashScene(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFlashScene:"), value)
}

// A Boolean value that specifies whether to configure the capture pipeline for high resolution still image capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/ishighresolutioncaptureenabled
func (c_ CapturePhotoOutput) IsHighResolutionCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighResolutionCaptureEnabled"))
	return rv
}


// SetIsHighResolutionCaptureEnabled sets the value of the isHighResolutionCaptureEnabled property.
// A Boolean value that specifies whether to configure the capture pipeline for high resolution still image capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/ishighresolutioncaptureenabled
func (c_ CapturePhotoOutput) SetIsHighResolutionCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighResolutionCaptureEnabled:"), value)
}

// A Boolean value indicating whether the capture output currently supports lens stabilization during bracketed image capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islensstabilizationduringbracketedcapturesupported
func (c_ CapturePhotoOutput) IsLensStabilizationDuringBracketedCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLensStabilizationDuringBracketedCaptureSupported"))
	return rv
}


// SetIsLensStabilizationDuringBracketedCaptureSupported sets the value of the isLensStabilizationDuringBracketedCaptureSupported property.
// A Boolean value indicating whether the capture output currently supports lens stabilization during bracketed image capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islensstabilizationduringbracketedcapturesupported
func (c_ CapturePhotoOutput) SetIsLensStabilizationDuringBracketedCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLensStabilizationDuringBracketedCaptureSupported:"), value)
}

// A Boolean value that indicates whether to automatically trim Live Photo movie captures to avoid excessive movement.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotoautotrimmingenabled
func (c_ CapturePhotoOutput) IsLivePhotoAutoTrimmingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLivePhotoAutoTrimmingEnabled"))
	return rv
}


// SetIsLivePhotoAutoTrimmingEnabled sets the value of the isLivePhotoAutoTrimmingEnabled property.
// A Boolean value that indicates whether to automatically trim Live Photo movie captures to avoid excessive movement.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotoautotrimmingenabled
func (c_ CapturePhotoOutput) SetIsLivePhotoAutoTrimmingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLivePhotoAutoTrimmingEnabled:"), value)
}

// A Boolean value that indicates whether to configure the capture pipeline for Live Photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocaptureenabled
func (c_ CapturePhotoOutput) IsLivePhotoCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLivePhotoCaptureEnabled"))
	return rv
}


// SetIsLivePhotoCaptureEnabled sets the value of the isLivePhotoCaptureEnabled property.
// A Boolean value that indicates whether to configure the capture pipeline for Live Photo capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocaptureenabled
func (c_ CapturePhotoOutput) SetIsLivePhotoCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLivePhotoCaptureEnabled:"), value)
}

// A Boolean value that indicates whether the capture output currently supports Live Photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocapturesupported
func (c_ CapturePhotoOutput) IsLivePhotoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLivePhotoCaptureSupported"))
	return rv
}


// SetIsLivePhotoCaptureSupported sets the value of the isLivePhotoCaptureSupported property.
// A Boolean value that indicates whether the capture output currently supports Live Photo capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocapturesupported
func (c_ CapturePhotoOutput) SetIsLivePhotoCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLivePhotoCaptureSupported:"), value)
}

// A Boolean value that indicates whether Live Photo capture is currently in a suspended state.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocapturesuspended
func (c_ CapturePhotoOutput) IsLivePhotoCaptureSuspended() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLivePhotoCaptureSuspended"))
	return rv
}


// SetIsLivePhotoCaptureSuspended sets the value of the isLivePhotoCaptureSuspended property.
// A Boolean value that indicates whether Live Photo capture is currently in a suspended state.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/islivephotocapturesuspended
func (c_ CapturePhotoOutput) SetIsLivePhotoCaptureSuspended(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLivePhotoCaptureSuspended:"), value)
}

// A Boolean value indicating whether the capture output generates a portrait effects matte.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isportraiteffectsmattedeliveryenabled
func (c_ CapturePhotoOutput) IsPortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectsMatteDeliveryEnabled"))
	return rv
}


// SetIsPortraitEffectsMatteDeliveryEnabled sets the value of the isPortraitEffectsMatteDeliveryEnabled property.
// A Boolean value indicating whether the capture output generates a portrait effects matte.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isportraiteffectsmattedeliveryenabled
func (c_ CapturePhotoOutput) SetIsPortraitEffectsMatteDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectsMatteDeliveryEnabled:"), value)
}

// A Boolean value indicating whether the capture output currently supports delivery of a portrait effects matte.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isportraiteffectsmattedeliverysupported
func (c_ CapturePhotoOutput) IsPortraitEffectsMatteDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectsMatteDeliverySupported"))
	return rv
}


// SetIsPortraitEffectsMatteDeliverySupported sets the value of the isPortraitEffectsMatteDeliverySupported property.
// A Boolean value indicating whether the capture output currently supports delivery of a portrait effects matte.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isportraiteffectsmattedeliverysupported
func (c_ CapturePhotoOutput) SetIsPortraitEffectsMatteDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectsMatteDeliverySupported:"), value)
}

// A Boolean value that indicates whether the photo output configuration enables responsive capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isresponsivecaptureenabled
func (c_ CapturePhotoOutput) IsResponsiveCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isResponsiveCaptureEnabled"))
	return rv
}


// SetIsResponsiveCaptureEnabled sets the value of the isResponsiveCaptureEnabled property.
// A Boolean value that indicates whether the photo output configuration enables responsive capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isresponsivecaptureenabled
func (c_ CapturePhotoOutput) SetIsResponsiveCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsResponsiveCaptureEnabled:"), value)
}

// A Boolean value that indicates whether the photo output supports responsive capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isresponsivecapturesupported
func (c_ CapturePhotoOutput) IsResponsiveCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isResponsiveCaptureSupported"))
	return rv
}


// SetIsResponsiveCaptureSupported sets the value of the isResponsiveCaptureSupported property.
// A Boolean value that indicates whether the photo output supports responsive capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isresponsivecapturesupported
func (c_ CapturePhotoOutput) SetIsResponsiveCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsResponsiveCaptureSupported:"), value)
}

// A Boolean value that indicates whether the photo output supports suppressing the system shutter sound.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isshuttersoundsuppressionsupported
func (c_ CapturePhotoOutput) IsShutterSoundSuppressionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isShutterSoundSuppressionSupported"))
	return rv
}


// SetIsShutterSoundSuppressionSupported sets the value of the isShutterSoundSuppressionSupported property.
// A Boolean value that indicates whether the photo output supports suppressing the system shutter sound.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isshuttersoundsuppressionsupported
func (c_ CapturePhotoOutput) SetIsShutterSoundSuppressionSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsShutterSoundSuppressionSupported:"), value)
}

// A Boolean value that indicates whether the photo output delivers photos from constituent cameras of a virtual device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdeviceconstituentphotodeliveryenabled
func (c_ CapturePhotoOutput) IsVirtualDeviceConstituentPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDeviceConstituentPhotoDeliveryEnabled"))
	return rv
}


// SetIsVirtualDeviceConstituentPhotoDeliveryEnabled sets the value of the isVirtualDeviceConstituentPhotoDeliveryEnabled property.
// A Boolean value that indicates whether the photo output delivers photos from constituent cameras of a virtual device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdeviceconstituentphotodeliveryenabled
func (c_ CapturePhotoOutput) SetIsVirtualDeviceConstituentPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDeviceConstituentPhotoDeliveryEnabled:"), value)
}

// A Boolean value that indicates whether the photo output configuration supports delivery of photos from constituent cameras of a virtual device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdeviceconstituentphotodeliverysupported
func (c_ CapturePhotoOutput) IsVirtualDeviceConstituentPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDeviceConstituentPhotoDeliverySupported"))
	return rv
}


// SetIsVirtualDeviceConstituentPhotoDeliverySupported sets the value of the isVirtualDeviceConstituentPhotoDeliverySupported property.
// A Boolean value that indicates whether the photo output configuration supports delivery of photos from constituent cameras of a virtual device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdeviceconstituentphotodeliverysupported
func (c_ CapturePhotoOutput) SetIsVirtualDeviceConstituentPhotoDeliverySupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDeviceConstituentPhotoDeliverySupported:"), value)
}

// A Boolean value that indicates whether the device supports virtual device image fusion.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdevicefusionsupported
func (c_ CapturePhotoOutput) IsVirtualDeviceFusionSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDeviceFusionSupported"))
	return rv
}


// SetIsVirtualDeviceFusionSupported sets the value of the isVirtualDeviceFusionSupported property.
// A Boolean value that indicates whether the device supports virtual device image fusion.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/isvirtualdevicefusionsupported
func (c_ CapturePhotoOutput) SetIsVirtualDeviceFusionSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDeviceFusionSupported:"), value)
}

// A Boolean value that indicates whether the photo output configuration enables zero shutter lag.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iszeroshutterlagenabled
func (c_ CapturePhotoOutput) IsZeroShutterLagEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isZeroShutterLagEnabled"))
	return rv
}


// SetIsZeroShutterLagEnabled sets the value of the isZeroShutterLagEnabled property.
// A Boolean value that indicates whether the photo output configuration enables zero shutter lag.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iszeroshutterlagenabled
func (c_ CapturePhotoOutput) SetIsZeroShutterLagEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsZeroShutterLagEnabled:"), value)
}

// A Boolean value that indicates whether the photo output supports zero shutter lag.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iszeroshutterlagsupported
func (c_ CapturePhotoOutput) IsZeroShutterLagSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isZeroShutterLagSupported"))
	return rv
}


// SetIsZeroShutterLagSupported sets the value of the isZeroShutterLagSupported property.
// A Boolean value that indicates whether the photo output supports zero shutter lag.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/iszeroshutterlagsupported
func (c_ CapturePhotoOutput) SetIsZeroShutterLagSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsZeroShutterLagSupported:"), value)
}

// The maximum number of images that the photo capture output can support in a single bracketed capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxbracketedcapturephotocount
func (c_ CapturePhotoOutput) MaxBracketedCapturePhotoCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxBracketedCapturePhotoCount"))
	return rv
}


// SetMaxBracketedCapturePhotoCount sets the value of the maxBracketedCapturePhotoCount property.
// The maximum number of images that the photo capture output can support in a single bracketed capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxbracketedcapturephotocount
func (c_ CapturePhotoOutput) SetMaxBracketedCapturePhotoCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxBracketedCapturePhotoCount:"), value)
}

// The maximum resolution of the requested photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxphotodimensions
func (c_ CapturePhotoOutput) MaxPhotoDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("maxPhotoDimensions"))
	return rv
}


// SetMaxPhotoDimensions sets the value of the maxPhotoDimensions property.
// The maximum resolution of the requested photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxphotodimensions
func (c_ CapturePhotoOutput) SetMaxPhotoDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxPhotoDimensions:"), value)
}

// The highest quality the photo output should prepare to deliver on a capture-by-capture basis.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxphotoqualityprioritization
func (c_ CapturePhotoOutput) MaxPhotoQualityPrioritization() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("maxPhotoQualityPrioritization"))
	return rv
}


// SetMaxPhotoQualityPrioritization sets the value of the maxPhotoQualityPrioritization property.
// The highest quality the photo output should prepare to deliver on a capture-by-capture basis.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/maxphotoqualityprioritization
func (c_ CapturePhotoOutput) SetMaxPhotoQualityPrioritization(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxPhotoQualityPrioritization:"), value)
}

// A photo settings object that controls how the photo output detects and handles automatic flash and stabilization modes.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/photosettingsforscenemonitoring
func (c_ CapturePhotoOutput) PhotoSettingsForSceneMonitoring() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("photoSettingsForSceneMonitoring"))
	return rv
}


// SetPhotoSettingsForSceneMonitoring sets the value of the photoSettingsForSceneMonitoring property.
// A photo settings object that controls how the photo output detects and handles automatic flash and stabilization modes.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/photosettingsforscenemonitoring
func (c_ CapturePhotoOutput) SetPhotoSettingsForSceneMonitoring(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhotoSettingsForSceneMonitoring:"), value)
}

// An array of photo settings for which the photo output has prepared capture resources.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/preparedphotosettingsarray
func (c_ CapturePhotoOutput) PreparedPhotoSettingsArray() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("preparedPhotoSettingsArray"))
	return rv
}


// SetPreparedPhotoSettingsArray sets the value of the preparedPhotoSettingsArray property.
// An array of photo settings for which the photo output has prepared capture resources.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/preparedphotosettingsarray
func (c_ CapturePhotoOutput) SetPreparedPhotoSettingsArray(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreparedPhotoSettingsArray:"), value)
}

// A Boolean value that indicates whether to preserve the suspended state of Live Photo capture when the session stops.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/preserveslivephotocapturesuspendedonsessionstop
func (c_ CapturePhotoOutput) PreservesLivePhotoCaptureSuspendedOnSessionStop() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("preservesLivePhotoCaptureSuspendedOnSessionStop"))
	return rv
}


// SetPreservesLivePhotoCaptureSuspendedOnSessionStop sets the value of the preservesLivePhotoCaptureSuspendedOnSessionStop property.
// A Boolean value that indicates whether to preserve the suspended state of Live Photo capture when the session stops.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/preserveslivephotocapturesuspendedonsessionstop
func (c_ CapturePhotoOutput) SetPreservesLivePhotoCaptureSuspendedOnSessionStop(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreservesLivePhotoCaptureSuspendedOnSessionStop:"), value)
}

// A Swift array of flash settings this capture output currently supports.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/supportedflashmodes-1n6nm
func (c_ CapturePhotoOutput) SupportedFlashModes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supportedFlashModes"))
	return rv
}


// SetSupportedFlashModes sets the value of the supportedFlashModes property.
// A Swift array of flash settings this capture output currently supports.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotooutput/supportedflashmodes-1n6nm
func (c_ CapturePhotoOutput) SetSupportedFlashModes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedFlashModes:"), value)
}

// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c_ CapturePhotoOutput) FlashMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("flashMode"))
	return rv
}


// SetFlashMode sets the value of the flashMode property.
// A setting for whether to fire the flash when capturing photos.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c_ CapturePhotoOutput) SetFlashMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashMode:"), value)
}

// A unique identifier for this photo settings instance.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/uniqueid
func (c_ CapturePhotoOutput) UniqueID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("uniqueID"))
	return rv
}


// SetUniqueID sets the value of the uniqueID property.
// A unique identifier for this photo settings instance.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/uniqueid
func (c_ CapturePhotoOutput) SetUniqueID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUniqueID:"), value)
}



