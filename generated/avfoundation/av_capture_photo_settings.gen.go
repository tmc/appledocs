// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CapturePhotoSettings] class.
var (
	CapturePhotoSettingsClass     _CapturePhotoSettingsClass
	CapturePhotoSettingsClassOnce sync.Once
)

func getCapturePhotoSettingsClass() _CapturePhotoSettingsClass {
	CapturePhotoSettingsClassOnce.Do(func() {
		CapturePhotoSettingsClass = _CapturePhotoSettingsClass{objc.GetClass("AVCapturePhotoSettings")}
	})
	return CapturePhotoSettingsClass
}

type _CapturePhotoSettingsClass struct {
	class objc.Class
}





// An interface definition for the [CapturePhotoSettings] class.
type ICapturePhotoSettings interface {
	objectivec.IObject
	

	// properties:
	FlashMode() CaptureFlashMode
	SetFlashMode(value CaptureFlashMode)
	Format() foundation.IDictionary
	ConstantColorEnabled() bool
	SetConstantColorEnabled(value bool)
	ConstantColorFallbackPhotoDeliveryEnabled() bool
	SetConstantColorFallbackPhotoDeliveryEnabled(value bool)
	HighResolutionPhotoEnabled() bool
	SetHighResolutionPhotoEnabled(value bool)
	ShutterSoundSuppressionEnabled() bool
	SetShutterSoundSuppressionEnabled(value bool)
	MaxPhotoDimensions() VideoDimensions /* not a class type */
	SetMaxPhotoDimensions(value VideoDimensions /* not a class type */)
	PhotoQualityPrioritization() CapturePhotoQualityPrioritization
	SetPhotoQualityPrioritization(value CapturePhotoQualityPrioritization)
	ProcessedFileType() FileType /* typedef */
	UniqueID() int64
	IsAutoContentAwareDistortionCorrectionEnabled() bool
	SetIsAutoContentAwareDistortionCorrectionEnabled(value bool)
	IsAutoDualCameraFusionEnabled() bool
	SetIsAutoDualCameraFusionEnabled(value bool)
	IsAutoRedEyeReductionEnabled() bool
	SetIsAutoRedEyeReductionEnabled(value bool)
	IsAutoStillImageStabilizationEnabled() bool
	SetIsAutoStillImageStabilizationEnabled(value bool)
	IsAutoVirtualDeviceFusionEnabled() bool
	SetIsAutoVirtualDeviceFusionEnabled(value bool)
	IsCameraCalibrationDataDeliveryEnabled() bool
	SetIsCameraCalibrationDataDeliveryEnabled(value bool)
	IsConstantColorEnabled() bool
	SetIsConstantColorEnabled(value bool)
	IsConstantColorFallbackPhotoDeliveryEnabled() bool
	SetIsConstantColorFallbackPhotoDeliveryEnabled(value bool)
	IsDepthDataDeliveryEnabled() bool
	SetIsDepthDataDeliveryEnabled(value bool)
	IsDepthDataFiltered() bool
	SetIsDepthDataFiltered(value bool)
	IsDualCameraDualPhotoDeliveryEnabled() bool
	SetIsDualCameraDualPhotoDeliveryEnabled(value bool)
	IsHighResolutionPhotoEnabled() bool
	SetIsHighResolutionPhotoEnabled(value bool)
	IsPortraitEffectsMatteDeliveryEnabled() bool
	SetIsPortraitEffectsMatteDeliveryEnabled(value bool)
	IsShutterSoundSuppressionEnabled() bool
	SetIsShutterSoundSuppressionEnabled(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoSettingsClass) Alloc() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CapturePhotoSettingsClass) New() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhotoSettings) Init() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhotoSettings) Autorelease() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhotoSettings creates a new CapturePhotoSettings instance.
func NewCapturePhotoSettings() CapturePhotoSettings {
	return getCapturePhotoSettingsClass().New()
}





// A specification of the features and settings to use for a single photo capture request.
//
// To take a photo, you create and configure a object, then pass it to the method. A instance can include any combination of settings, regardless of whether that combination is valid for a given capture session. When you initiate a capture by passing a photo settings object to the method, the photo capture output validates your settings to ensure deterministic behavior. For example, the setting must specify a value that’s present in the photo output’s array. For detailed validation rules, see each property description below.


// A specification of the features and settings to use for a single photo capture request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings
type CapturePhotoSettings struct {
	objectivec.Object
}

// CapturePhotoSettingsFrom constructs a [CapturePhotoSettings] from an unsafe.Pointer.
//
// A specification of the features and settings to use for a single photo capture request.
func CapturePhotoSettingsFrom(ptr unsafe.Pointer) CapturePhotoSettings {
	return CapturePhotoSettings{objectivec.Object{objc.ID(ptr)}}
}






// Creates a unique photo settings object, copying all settings values from the specified photo settings object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(from:)
func NewCapturePhotoSettingsFromPhotoSettings(photoSettings IAVCapturePhotoSettings) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsFromPhotoSettings:"), photoSettings)
	return rv
}


// Creates a photo settings object with the specified output format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(format:)
func NewCapturePhotoSettingsWithFormat(format foundation.IDictionary) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithFormat:"), format)
	return rv
}


// Creates a photo settings object for RAW-format-only capture with the specified pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:)
func NewCapturePhotoSettingsWithRawPixelFormatType(rawPixelFormatType uint32 /* not a class type */) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithRawPixelFormatType:"), rawPixelFormatType)
	return rv
}


// Creates a photo settings object for capture in both RAW format and a processed format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:processedFormat:)
func NewCapturePhotoSettingsWithRawPixelFormatTypeProcessedFormat(rawPixelFormatType uint32 /* not a class type */, processedFormat foundation.IDictionary) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithRawPixelFormatType:processedFormat:"), rawPixelFormatType, processedFormat)
	return rv
}


// Creates a photo settings object for capture in both RAW format and a processed format with the specified output file types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:rawFileType:processedFormat:processedFileType:)
func NewCapturePhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType(rawPixelFormatType uint32 /* not a class type */, rawFileType FileType /* typedef */, processedFormat foundation.IDictionary, processedFileType FileType /* typedef */) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithRawPixelFormatType:rawFileType:processedFormat:processedFileType:"), rawPixelFormatType, rawFileType, processedFormat, processedFileType)
	return rv
}







// Creates a photo settings object with the specified output format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(format:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsWithFormat(format foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsWithFormat:"), format)
	return rv
}


// Creates a unique photo settings object, copying all settings values from the specified photo settings object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(from:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsFromPhotoSettings(photoSettings IAVCapturePhotoSettings) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsFromPhotoSettings:"), photoSettings)
	return rv
}


// Creates a photo settings object for RAW-format-only capture with the specified pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsWithRawPixelFormatType(rawPixelFormatType uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsWithRawPixelFormatType:"), rawPixelFormatType)
	return rv
}


// Creates a photo settings object for capture in both RAW format and a processed format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:processedFormat:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsWithRawPixelFormatTypeProcessedFormat(rawPixelFormatType uint32 /* not a class type */, processedFormat foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsWithRawPixelFormatType:processedFormat:"), rawPixelFormatType, processedFormat)
	return rv
}


// Creates a photo settings object for capture in both RAW format and a processed format with the specified output file types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:rawFileType:processedFormat:processedFileType:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType(rawPixelFormatType uint32 /* not a class type */, rawFileType FileType /* typedef */, processedFormat foundation.IDictionary, processedFileType FileType /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsWithRawPixelFormatType:rawFileType:processedFormat:processedFileType:"), rawPixelFormatType, rawFileType, processedFormat, processedFileType)
	return rv
}


// Creates a photo settings object with default settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/photoSettings
func (cc _CapturePhotoSettingsClass) PhotoSettings() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettings"))
	return rv
}

















// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/flashMode
func (c_ CapturePhotoSettings) FlashMode() CaptureFlashMode {
	rv := objc.Send[CaptureFlashMode](c_.ID, objc.Sel("flashMode"))
	return rv
}


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/flashMode
func (c_ CapturePhotoSettings) SetFlashMode(value CaptureFlashMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashMode:"), value)
}


// A dictionary describing the processed format (for example, JPEG) to deliver captured photos in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/format
func (c_ CapturePhotoSettings) Format() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("format"))
	return rv
}


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorEnabled
func (c_ CapturePhotoSettings) ConstantColorEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("constantColorEnabled"))
	return rv
}


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorEnabled
func (c_ CapturePhotoSettings) SetConstantColorEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstantColorEnabled:"), value)
}


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorFallbackPhotoDeliveryEnabled
func (c_ CapturePhotoSettings) ConstantColorFallbackPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("constantColorFallbackPhotoDeliveryEnabled"))
	return rv
}


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorFallbackPhotoDeliveryEnabled
func (c_ CapturePhotoSettings) SetConstantColorFallbackPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstantColorFallbackPhotoDeliveryEnabled:"), value)
}


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isHighResolutionPhotoEnabled
func (c_ CapturePhotoSettings) HighResolutionPhotoEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highResolutionPhotoEnabled"))
	return rv
}


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isHighResolutionPhotoEnabled
func (c_ CapturePhotoSettings) SetHighResolutionPhotoEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighResolutionPhotoEnabled:"), value)
}


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isShutterSoundSuppressionEnabled
func (c_ CapturePhotoSettings) ShutterSoundSuppressionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shutterSoundSuppressionEnabled"))
	return rv
}


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isShutterSoundSuppressionEnabled
func (c_ CapturePhotoSettings) SetShutterSoundSuppressionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShutterSoundSuppressionEnabled:"), value)
}


// The maximum resolution of the photo to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/maxPhotoDimensions
func (c_ CapturePhotoSettings) MaxPhotoDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("maxPhotoDimensions"))
	return rv
}


// The maximum resolution of the photo to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/maxPhotoDimensions
func (c_ CapturePhotoSettings) SetMaxPhotoDimensions(value VideoDimensions /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxPhotoDimensions:"), value)
}


// A setting that indicates how to prioritize photo quality against speed of photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/photoQualityPrioritization
func (c_ CapturePhotoSettings) PhotoQualityPrioritization() CapturePhotoQualityPrioritization {
	rv := objc.Send[CapturePhotoQualityPrioritization](c_.ID, objc.Sel("photoQualityPrioritization"))
	return rv
}


// A setting that indicates how to prioritize photo quality against speed of photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/photoQualityPrioritization
func (c_ CapturePhotoSettings) SetPhotoQualityPrioritization(value CapturePhotoQualityPrioritization) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhotoQualityPrioritization:"), value)
}


// The container file format for eventual output of the processed image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/processedFileType
func (c_ CapturePhotoSettings) ProcessedFileType() FileType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("processedFileType"))
	return rv
}


// A unique identifier for this photo settings instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/uniqueID
func (c_ CapturePhotoSettings) UniqueID() int64 {
	rv := objc.Send[int64](c_.ID, objc.Sel("uniqueID"))
	return rv
}


// A Boolean value that specifies whether the photo output, at its discretion, uses content-aware distortion correction on this photo request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautocontentawaredistortioncorrectionenabled
func (c_ CapturePhotoSettings) IsAutoContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoContentAwareDistortionCorrectionEnabled"))
	return rv
}


// A Boolean value that specifies whether the photo output, at its discretion, uses content-aware distortion correction on this photo request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautocontentawaredistortioncorrectionenabled
func (c_ CapturePhotoSettings) SetIsAutoContentAwareDistortionCorrectionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoContentAwareDistortionCorrectionEnabled:"), value)
}


// A Boolean value that specifies whether captures automatically combine data from a dual camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautodualcamerafusionenabled
func (c_ CapturePhotoSettings) IsAutoDualCameraFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoDualCameraFusionEnabled"))
	return rv
}


// A Boolean value that specifies whether captures automatically combine data from a dual camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautodualcamerafusionenabled
func (c_ CapturePhotoSettings) SetIsAutoDualCameraFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoDualCameraFusionEnabled:"), value)
}


// A Boolean value that indicates whether to use auto red-eye reduction on flash captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautoredeyereductionenabled
func (c_ CapturePhotoSettings) IsAutoRedEyeReductionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoRedEyeReductionEnabled"))
	return rv
}


// A Boolean value that indicates whether to use auto red-eye reduction on flash captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautoredeyereductionenabled
func (c_ CapturePhotoSettings) SetIsAutoRedEyeReductionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoRedEyeReductionEnabled:"), value)
}


// A Boolean value that specifies whether captures use automatic image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautostillimagestabilizationenabled
func (c_ CapturePhotoSettings) IsAutoStillImageStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoStillImageStabilizationEnabled"))
	return rv
}


// A Boolean value that specifies whether captures use automatic image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautostillimagestabilizationenabled
func (c_ CapturePhotoSettings) SetIsAutoStillImageStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoStillImageStabilizationEnabled:"), value)
}


// A Boolean value that specifies whether to use automatic virtual-device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautovirtualdevicefusionenabled
func (c_ CapturePhotoSettings) IsAutoVirtualDeviceFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoVirtualDeviceFusionEnabled"))
	return rv
}


// A Boolean value that specifies whether to use automatic virtual-device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautovirtualdevicefusionenabled
func (c_ CapturePhotoSettings) SetIsAutoVirtualDeviceFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoVirtualDeviceFusionEnabled:"), value)
}


// A Boolean value that determines whether a dual photo capture also delivers camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/iscameracalibrationdatadeliveryenabled
func (c_ CapturePhotoSettings) IsCameraCalibrationDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraCalibrationDataDeliveryEnabled"))
	return rv
}


// A Boolean value that determines whether a dual photo capture also delivers camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/iscameracalibrationdatadeliveryenabled
func (c_ CapturePhotoSettings) SetIsCameraCalibrationDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraCalibrationDataDeliveryEnabled:"), value)
}


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorenabled
func (c_ CapturePhotoSettings) IsConstantColorEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorEnabled"))
	return rv
}


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorenabled
func (c_ CapturePhotoSettings) SetIsConstantColorEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorEnabled:"), value)
}


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorfallbackphotodeliveryenabled
func (c_ CapturePhotoSettings) IsConstantColorFallbackPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorFallbackPhotoDeliveryEnabled"))
	return rv
}


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorfallbackphotodeliveryenabled
func (c_ CapturePhotoSettings) SetIsConstantColorFallbackPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorFallbackPhotoDeliveryEnabled:"), value)
}


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled
func (c_ CapturePhotoSettings) IsDepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataDeliveryEnabled"))
	return rv
}


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled
func (c_ CapturePhotoSettings) SetIsDepthDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataDeliveryEnabled:"), value)
}


// A Boolean value that determines whether to smooth noise and fill in missing values in depth data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatafiltered
func (c_ CapturePhotoSettings) IsDepthDataFiltered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataFiltered"))
	return rv
}


// A Boolean value that determines whether to smooth noise and fill in missing values in depth data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatafiltered
func (c_ CapturePhotoSettings) SetIsDepthDataFiltered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataFiltered:"), value)
}


// A Boolean value that determines whether a dual camera device delivers images from both cameras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdualcameradualphotodeliveryenabled
func (c_ CapturePhotoSettings) IsDualCameraDualPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDualCameraDualPhotoDeliveryEnabled"))
	return rv
}


// A Boolean value that determines whether a dual camera device delivers images from both cameras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdualcameradualphotodeliveryenabled
func (c_ CapturePhotoSettings) SetIsDualCameraDualPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDualCameraDualPhotoDeliveryEnabled:"), value)
}


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled
func (c_ CapturePhotoSettings) IsHighResolutionPhotoEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighResolutionPhotoEnabled"))
	return rv
}


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled
func (c_ CapturePhotoSettings) SetIsHighResolutionPhotoEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighResolutionPhotoEnabled:"), value)
}


// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isportraiteffectsmattedeliveryenabled
func (c_ CapturePhotoSettings) IsPortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectsMatteDeliveryEnabled"))
	return rv
}


// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isportraiteffectsmattedeliveryenabled
func (c_ CapturePhotoSettings) SetIsPortraitEffectsMatteDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectsMatteDeliveryEnabled:"), value)
}


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isshuttersoundsuppressionenabled
func (c_ CapturePhotoSettings) IsShutterSoundSuppressionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isShutterSoundSuppressionEnabled"))
	return rv
}


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isshuttersoundsuppressionenabled
func (c_ CapturePhotoSettings) SetIsShutterSoundSuppressionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsShutterSoundSuppressionEnabled:"), value)
}







