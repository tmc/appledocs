// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCapturePhotoSettings */


/* debug [class_header]: Header for AVCapturePhotoSettings */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CapturePhotoSettings */
// An interface definition for the [CapturePhotoSettings] class.
type ICapturePhotoSettings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CapturePhotoSettings */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CapturePhotoSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CapturePhotoSettings */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CapturePhotoSettings */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CapturePhotoSettings */

// Creates a unique photo settings object, copying all settings values from the specified photo settings object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(from:)
func NewCapturePhotoSettingsFromPhotoSettings(photoSettings IAVCapturePhotoSettings) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsFromPhotoSettings:"), photoSettings)
	return rv
}/* debug [class_init_methods/constructor]: NewCapturePhotoSettingsFromPhotoSettings */


// Creates a photo settings object with the specified output format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(format:)
func NewCapturePhotoSettingsWithFormat(format foundation.IDictionary) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithFormat:"), format)
	return rv
}/* debug [class_init_methods/constructor]: NewCapturePhotoSettingsWithFormat */


// Creates a photo settings object for RAW-format-only capture with the specified pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:)
func NewCapturePhotoSettingsWithRawPixelFormatType(rawPixelFormatType uint32 /* not a class type */) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithRawPixelFormatType:"), rawPixelFormatType)
	return rv
}/* debug [class_init_methods/constructor]: NewCapturePhotoSettingsWithRawPixelFormatType */


// Creates a photo settings object for capture in both RAW format and a processed format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:processedFormat:)
func NewCapturePhotoSettingsWithRawPixelFormatTypeProcessedFormat(rawPixelFormatType uint32 /* not a class type */, processedFormat foundation.IDictionary) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithRawPixelFormatType:processedFormat:"), rawPixelFormatType, processedFormat)
	return rv
}/* debug [class_init_methods/constructor]: NewCapturePhotoSettingsWithRawPixelFormatTypeProcessedFormat */


// Creates a photo settings object for capture in both RAW format and a processed format with the specified output file types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:rawFileType:processedFormat:processedFileType:)
func NewCapturePhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType(rawPixelFormatType uint32 /* not a class type */, rawFileType FileType /* typedef */, processedFormat foundation.IDictionary, processedFileType FileType /* typedef */) CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(getCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithRawPixelFormatType:rawFileType:processedFormat:processedFileType:"), rawPixelFormatType, rawFileType, processedFormat, processedFileType)
	return rv
}/* debug [class_init_methods/constructor]: NewCapturePhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CapturePhotoSettings */

// Creates a photo settings object with the specified output format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(format:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsWithFormat(format foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsWithFormat:"), format)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhotoSettingsWithFormat) */


// Creates a unique photo settings object, copying all settings values from the specified photo settings object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(from:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsFromPhotoSettings(photoSettings IAVCapturePhotoSettings) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsFromPhotoSettings:"), photoSettings)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhotoSettingsFromPhotoSettings) */


// Creates a photo settings object for RAW-format-only capture with the specified pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsWithRawPixelFormatType(rawPixelFormatType uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsWithRawPixelFormatType:"), rawPixelFormatType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhotoSettingsWithRawPixelFormatType) */


// Creates a photo settings object for capture in both RAW format and a processed format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:processedFormat:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsWithRawPixelFormatTypeProcessedFormat(rawPixelFormatType uint32 /* not a class type */, processedFormat foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsWithRawPixelFormatType:processedFormat:"), rawPixelFormatType, processedFormat)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhotoSettingsWithRawPixelFormatTypeProcessedFormat) */


// Creates a photo settings object for capture in both RAW format and a processed format with the specified output file types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:rawFileType:processedFormat:processedFileType:)
func (cc _CapturePhotoSettingsClass) PhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType(rawPixelFormatType uint32 /* not a class type */, rawFileType FileType /* typedef */, processedFormat foundation.IDictionary, processedFileType FileType /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettingsWithRawPixelFormatType:rawFileType:processedFormat:processedFileType:"), rawPixelFormatType, rawFileType, processedFormat, processedFileType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType) */


// Creates a photo settings object with default settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/photoSettings
func (cc _CapturePhotoSettingsClass) PhotoSettings() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoSettings"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhotoSettings) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CapturePhotoSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CapturePhotoSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CapturePhotoSettings */

// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/flashMode
func (c_ CapturePhotoSettings) FlashMode() CaptureFlashMode {
	rv := objc.Send[CaptureFlashMode](c_.ID, objc.Sel("flashMode"))
	return rv
}/* debug [instance_properties/getter]: flashMode */


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/flashMode
func (c_ CapturePhotoSettings) SetFlashMode(value CaptureFlashMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashMode:"), value)
}/* debug [instance_properties/setter]: flashMode */


// A dictionary describing the processed format (for example, JPEG) to deliver captured photos in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/format
func (c_ CapturePhotoSettings) Format() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorEnabled
func (c_ CapturePhotoSettings) ConstantColorEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("constantColorEnabled"))
	return rv
}/* debug [instance_properties/getter]: constantColorEnabled */


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorEnabled
func (c_ CapturePhotoSettings) SetConstantColorEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstantColorEnabled:"), value)
}/* debug [instance_properties/setter]: constantColorEnabled */


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorFallbackPhotoDeliveryEnabled
func (c_ CapturePhotoSettings) ConstantColorFallbackPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("constantColorFallbackPhotoDeliveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: constantColorFallbackPhotoDeliveryEnabled */


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorFallbackPhotoDeliveryEnabled
func (c_ CapturePhotoSettings) SetConstantColorFallbackPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstantColorFallbackPhotoDeliveryEnabled:"), value)
}/* debug [instance_properties/setter]: constantColorFallbackPhotoDeliveryEnabled */


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isHighResolutionPhotoEnabled
func (c_ CapturePhotoSettings) HighResolutionPhotoEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highResolutionPhotoEnabled"))
	return rv
}/* debug [instance_properties/getter]: highResolutionPhotoEnabled */


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isHighResolutionPhotoEnabled
func (c_ CapturePhotoSettings) SetHighResolutionPhotoEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighResolutionPhotoEnabled:"), value)
}/* debug [instance_properties/setter]: highResolutionPhotoEnabled */


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isShutterSoundSuppressionEnabled
func (c_ CapturePhotoSettings) ShutterSoundSuppressionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shutterSoundSuppressionEnabled"))
	return rv
}/* debug [instance_properties/getter]: shutterSoundSuppressionEnabled */


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isShutterSoundSuppressionEnabled
func (c_ CapturePhotoSettings) SetShutterSoundSuppressionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShutterSoundSuppressionEnabled:"), value)
}/* debug [instance_properties/setter]: shutterSoundSuppressionEnabled */


// The maximum resolution of the photo to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/maxPhotoDimensions
func (c_ CapturePhotoSettings) MaxPhotoDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("maxPhotoDimensions"))
	return rv
}/* debug [instance_properties/getter]: maxPhotoDimensions */


// The maximum resolution of the photo to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/maxPhotoDimensions
func (c_ CapturePhotoSettings) SetMaxPhotoDimensions(value VideoDimensions /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxPhotoDimensions:"), value)
}/* debug [instance_properties/setter]: maxPhotoDimensions */


// A setting that indicates how to prioritize photo quality against speed of photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/photoQualityPrioritization
func (c_ CapturePhotoSettings) PhotoQualityPrioritization() CapturePhotoQualityPrioritization {
	rv := objc.Send[CapturePhotoQualityPrioritization](c_.ID, objc.Sel("photoQualityPrioritization"))
	return rv
}/* debug [instance_properties/getter]: photoQualityPrioritization */


// A setting that indicates how to prioritize photo quality against speed of photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/photoQualityPrioritization
func (c_ CapturePhotoSettings) SetPhotoQualityPrioritization(value CapturePhotoQualityPrioritization) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhotoQualityPrioritization:"), value)
}/* debug [instance_properties/setter]: photoQualityPrioritization */


// The container file format for eventual output of the processed image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/processedFileType
func (c_ CapturePhotoSettings) ProcessedFileType() FileType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("processedFileType"))
	return rv
}/* debug [instance_properties/getter]: processedFileType */


// A unique identifier for this photo settings instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/uniqueID
func (c_ CapturePhotoSettings) UniqueID() int64 {
	rv := objc.Send[int64](c_.ID, objc.Sel("uniqueID"))
	return rv
}/* debug [instance_properties/getter]: uniqueID */


// A Boolean value that specifies whether the photo output, at its discretion, uses content-aware distortion correction on this photo request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautocontentawaredistortioncorrectionenabled
func (c_ CapturePhotoSettings) IsAutoContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoContentAwareDistortionCorrectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutoContentAwareDistortionCorrectionEnabled */


// A Boolean value that specifies whether the photo output, at its discretion, uses content-aware distortion correction on this photo request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautocontentawaredistortioncorrectionenabled
func (c_ CapturePhotoSettings) SetIsAutoContentAwareDistortionCorrectionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoContentAwareDistortionCorrectionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutoContentAwareDistortionCorrectionEnabled */


// A Boolean value that specifies whether captures automatically combine data from a dual camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautodualcamerafusionenabled
func (c_ CapturePhotoSettings) IsAutoDualCameraFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoDualCameraFusionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutoDualCameraFusionEnabled */


// A Boolean value that specifies whether captures automatically combine data from a dual camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautodualcamerafusionenabled
func (c_ CapturePhotoSettings) SetIsAutoDualCameraFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoDualCameraFusionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutoDualCameraFusionEnabled */


// A Boolean value that indicates whether to use auto red-eye reduction on flash captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautoredeyereductionenabled
func (c_ CapturePhotoSettings) IsAutoRedEyeReductionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoRedEyeReductionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutoRedEyeReductionEnabled */


// A Boolean value that indicates whether to use auto red-eye reduction on flash captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautoredeyereductionenabled
func (c_ CapturePhotoSettings) SetIsAutoRedEyeReductionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoRedEyeReductionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutoRedEyeReductionEnabled */


// A Boolean value that specifies whether captures use automatic image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautostillimagestabilizationenabled
func (c_ CapturePhotoSettings) IsAutoStillImageStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoStillImageStabilizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutoStillImageStabilizationEnabled */


// A Boolean value that specifies whether captures use automatic image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautostillimagestabilizationenabled
func (c_ CapturePhotoSettings) SetIsAutoStillImageStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoStillImageStabilizationEnabled:"), value)
}/* debug [instance_properties/setter]: isAutoStillImageStabilizationEnabled */


// A Boolean value that specifies whether to use automatic virtual-device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautovirtualdevicefusionenabled
func (c_ CapturePhotoSettings) IsAutoVirtualDeviceFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoVirtualDeviceFusionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutoVirtualDeviceFusionEnabled */


// A Boolean value that specifies whether to use automatic virtual-device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautovirtualdevicefusionenabled
func (c_ CapturePhotoSettings) SetIsAutoVirtualDeviceFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoVirtualDeviceFusionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutoVirtualDeviceFusionEnabled */


// A Boolean value that determines whether a dual photo capture also delivers camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/iscameracalibrationdatadeliveryenabled
func (c_ CapturePhotoSettings) IsCameraCalibrationDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraCalibrationDataDeliveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: isCameraCalibrationDataDeliveryEnabled */


// A Boolean value that determines whether a dual photo capture also delivers camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/iscameracalibrationdatadeliveryenabled
func (c_ CapturePhotoSettings) SetIsCameraCalibrationDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraCalibrationDataDeliveryEnabled:"), value)
}/* debug [instance_properties/setter]: isCameraCalibrationDataDeliveryEnabled */


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorenabled
func (c_ CapturePhotoSettings) IsConstantColorEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorEnabled"))
	return rv
}/* debug [instance_properties/getter]: isConstantColorEnabled */


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorenabled
func (c_ CapturePhotoSettings) SetIsConstantColorEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorEnabled:"), value)
}/* debug [instance_properties/setter]: isConstantColorEnabled */


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorfallbackphotodeliveryenabled
func (c_ CapturePhotoSettings) IsConstantColorFallbackPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorFallbackPhotoDeliveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: isConstantColorFallbackPhotoDeliveryEnabled */


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorfallbackphotodeliveryenabled
func (c_ CapturePhotoSettings) SetIsConstantColorFallbackPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorFallbackPhotoDeliveryEnabled:"), value)
}/* debug [instance_properties/setter]: isConstantColorFallbackPhotoDeliveryEnabled */


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled
func (c_ CapturePhotoSettings) IsDepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataDeliveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDepthDataDeliveryEnabled */


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled
func (c_ CapturePhotoSettings) SetIsDepthDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataDeliveryEnabled:"), value)
}/* debug [instance_properties/setter]: isDepthDataDeliveryEnabled */


// A Boolean value that determines whether to smooth noise and fill in missing values in depth data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatafiltered
func (c_ CapturePhotoSettings) IsDepthDataFiltered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataFiltered"))
	return rv
}/* debug [instance_properties/getter]: isDepthDataFiltered */


// A Boolean value that determines whether to smooth noise and fill in missing values in depth data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatafiltered
func (c_ CapturePhotoSettings) SetIsDepthDataFiltered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataFiltered:"), value)
}/* debug [instance_properties/setter]: isDepthDataFiltered */


// A Boolean value that determines whether a dual camera device delivers images from both cameras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdualcameradualphotodeliveryenabled
func (c_ CapturePhotoSettings) IsDualCameraDualPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDualCameraDualPhotoDeliveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDualCameraDualPhotoDeliveryEnabled */


// A Boolean value that determines whether a dual camera device delivers images from both cameras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdualcameradualphotodeliveryenabled
func (c_ CapturePhotoSettings) SetIsDualCameraDualPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDualCameraDualPhotoDeliveryEnabled:"), value)
}/* debug [instance_properties/setter]: isDualCameraDualPhotoDeliveryEnabled */


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled
func (c_ CapturePhotoSettings) IsHighResolutionPhotoEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighResolutionPhotoEnabled"))
	return rv
}/* debug [instance_properties/getter]: isHighResolutionPhotoEnabled */


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled
func (c_ CapturePhotoSettings) SetIsHighResolutionPhotoEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighResolutionPhotoEnabled:"), value)
}/* debug [instance_properties/setter]: isHighResolutionPhotoEnabled */


// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isportraiteffectsmattedeliveryenabled
func (c_ CapturePhotoSettings) IsPortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectsMatteDeliveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: isPortraitEffectsMatteDeliveryEnabled */


// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isportraiteffectsmattedeliveryenabled
func (c_ CapturePhotoSettings) SetIsPortraitEffectsMatteDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectsMatteDeliveryEnabled:"), value)
}/* debug [instance_properties/setter]: isPortraitEffectsMatteDeliveryEnabled */


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isshuttersoundsuppressionenabled
func (c_ CapturePhotoSettings) IsShutterSoundSuppressionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isShutterSoundSuppressionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isShutterSoundSuppressionEnabled */


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isshuttersoundsuppressionenabled
func (c_ CapturePhotoSettings) SetIsShutterSoundSuppressionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsShutterSoundSuppressionEnabled:"), value)
}/* debug [instance_properties/setter]: isShutterSoundSuppressionEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCapturePhotoSettings */


