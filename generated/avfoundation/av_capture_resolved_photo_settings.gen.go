// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureResolvedPhotoSettings] class.
var (
	CaptureResolvedPhotoSettingsClass     _CaptureResolvedPhotoSettingsClass
	CaptureResolvedPhotoSettingsClassOnce sync.Once
)

func getCaptureResolvedPhotoSettingsClass() _CaptureResolvedPhotoSettingsClass {
	CaptureResolvedPhotoSettingsClassOnce.Do(func() {
		CaptureResolvedPhotoSettingsClass = _CaptureResolvedPhotoSettingsClass{objc.GetClass("AVCaptureResolvedPhotoSettings")}
	})
	return CaptureResolvedPhotoSettingsClass
}

type _CaptureResolvedPhotoSettingsClass struct {
	class objc.Class
}

// An interface definition for the [CaptureResolvedPhotoSettings] class.
type ICaptureResolvedPhotoSettings interface {
	objectivec.IObject
	EmbeddedThumbnailDimensions() unsafe.Pointer
	FlashMode() unsafe.Pointer
	SetFlashMode(value unsafe.Pointer)
	PreviewPhotoFormat() string
	SetPreviewPhotoFormat(value string)
	UniqueID() unsafe.Pointer
	SetUniqueID(value unsafe.Pointer)
	DeferredPhotoProxyDimensions() unsafe.Pointer
	SetDeferredPhotoProxyDimensions(value unsafe.Pointer)
	ExpectedPhotoCount() int
	SetExpectedPhotoCount(value int)
	IsContentAwareDistortionCorrectionEnabled() bool
	SetIsContentAwareDistortionCorrectionEnabled(value bool)
	IsDualCameraFusionEnabled() bool
	SetIsDualCameraFusionEnabled(value bool)
	IsFastCapturePrioritizationEnabled() bool
	SetIsFastCapturePrioritizationEnabled(value bool)
	IsFlashEnabled() bool
	SetIsFlashEnabled(value bool)
	IsRedEyeReductionEnabled() bool
	SetIsRedEyeReductionEnabled(value bool)
	IsStillImageStabilizationEnabled() bool
	SetIsStillImageStabilizationEnabled(value bool)
	IsVirtualDeviceFusionEnabled() bool
	SetIsVirtualDeviceFusionEnabled(value bool)
	LivePhotoMovieDimensions() unsafe.Pointer
	SetLivePhotoMovieDimensions(value unsafe.Pointer)
	PhotoDimensions() unsafe.Pointer
	SetPhotoDimensions(value unsafe.Pointer)
	PhotoProcessingTimeRange() unsafe.Pointer
	SetPhotoProcessingTimeRange(value unsafe.Pointer)
	PortraitEffectsMatteDimensions() unsafe.Pointer
	SetPortraitEffectsMatteDimensions(value unsafe.Pointer)
	PreviewDimensions() unsafe.Pointer
	SetPreviewDimensions(value unsafe.Pointer)
	RawEmbeddedThumbnailDimensions() unsafe.Pointer
	SetRawEmbeddedThumbnailDimensions(value unsafe.Pointer)
	RawPhotoDimensions() unsafe.Pointer
	SetRawPhotoDimensions(value unsafe.Pointer)
}

// A description of the features and settings in use for an in-progress or complete photo capture request.
//
// When you request a photo capture using the method, you describe the settings for that capture request in an object. When the capture begins, the photo output calls your delegate methods and provides an object detailing the settings that are in effect for that capture. Resolved photo settings objects are immutable; they describe a request that has already been made. The property of a resolved photo settings object passed to one of your methods matches the value of the object you passed when requesting capture. Use this value to determine which delegate method calls correspond to which capture requests. Some photo capture settings are automatic, such as the property. For such settings, the photo output determines whether to use that feature at the moment of capture—you don’t know when requesting a capture whether the feature is active when the capture completes. When the photo output calls your delegate methods, the provided object details which automatic features have been set for that capture. Likewise, the dimensions of an output image or movie may not be set until the moment of capture. For example, when you specify a thumbnail size with the setting, the photo output chooses dimensions that best match your requested size while preserving the aspect ratio of the captured photo. When the photo output calls your delegate methods, use the property of the resolved settings to find the actual preview image dimensions. See the methods listed in Examining Output Dimensions for other cases where output dimensions can change at capture time.


// A description of the features and settings in use for an in-progress or complete photo capture request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings
type CaptureResolvedPhotoSettings struct {
	objectivec.Object
}

// CaptureResolvedPhotoSettingsFrom constructs a [CaptureResolvedPhotoSettings] from an unsafe.Pointer.
//
// A description of the features and settings in use for an in-progress or complete photo capture request.
func CaptureResolvedPhotoSettingsFrom(ptr unsafe.Pointer) CaptureResolvedPhotoSettings {
	return CaptureResolvedPhotoSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureResolvedPhotoSettingsClass) Alloc() CaptureResolvedPhotoSettings {
	rv := objc.Send[CaptureResolvedPhotoSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureResolvedPhotoSettingsClass) New() CaptureResolvedPhotoSettings {
	rv := objc.Send[CaptureResolvedPhotoSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureResolvedPhotoSettings) Init() CaptureResolvedPhotoSettings {
	rv := objc.Send[CaptureResolvedPhotoSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureResolvedPhotoSettings) Autorelease() CaptureResolvedPhotoSettings {
	rv := objc.Send[CaptureResolvedPhotoSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureResolvedPhotoSettings creates a new CaptureResolvedPhotoSettings instance.
func NewCaptureResolvedPhotoSettings() CaptureResolvedPhotoSettings {
	return getCaptureResolvedPhotoSettingsClass().New()
}



// The size, in pixels, of the thumbnail image that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/embeddedThumbnailDimensions
func (c_ CaptureResolvedPhotoSettings) EmbeddedThumbnailDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("embeddedThumbnailDimensions"))
	return rv
}


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c_ CaptureResolvedPhotoSettings) FlashMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("flashMode"))
	return rv
}


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c_ CaptureResolvedPhotoSettings) SetFlashMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashMode:"), value)
}


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CaptureResolvedPhotoSettings) PreviewPhotoFormat() string {
	rv := objc.Send[string](c_.ID, objc.Sel("previewPhotoFormat"))
	return rv
}


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CaptureResolvedPhotoSettings) SetPreviewPhotoFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewPhotoFormat:"), objc.String(value))
}


// A unique identifier for this photo settings instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/uniqueid
func (c_ CaptureResolvedPhotoSettings) UniqueID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("uniqueID"))
	return rv
}


// A unique identifier for this photo settings instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/uniqueid
func (c_ CaptureResolvedPhotoSettings) SetUniqueID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUniqueID:"), value)
}


// The resolved dimensions of the photo proxy when using deferred photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/deferredphotoproxydimensions
func (c_ CaptureResolvedPhotoSettings) DeferredPhotoProxyDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("deferredPhotoProxyDimensions"))
	return rv
}


// The resolved dimensions of the photo proxy when using deferred photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/deferredphotoproxydimensions
func (c_ CaptureResolvedPhotoSettings) SetDeferredPhotoProxyDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeferredPhotoProxyDimensions:"), value)
}


// The number of photo capture results in the capture request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/expectedphotocount
func (c_ CaptureResolvedPhotoSettings) ExpectedPhotoCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("expectedPhotoCount"))
	return rv
}


// The number of photo capture results in the capture request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/expectedphotocount
func (c_ CaptureResolvedPhotoSettings) SetExpectedPhotoCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExpectedPhotoCount:"), value)
}


// A Boolean value that indicates whether the system applies content-aware distortion correction when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/iscontentawaredistortioncorrectionenabled
func (c_ CaptureResolvedPhotoSettings) IsContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContentAwareDistortionCorrectionEnabled"))
	return rv
}


// A Boolean value that indicates whether the system applies content-aware distortion correction when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/iscontentawaredistortioncorrectionenabled
func (c_ CaptureResolvedPhotoSettings) SetIsContentAwareDistortionCorrectionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContentAwareDistortionCorrectionEnabled:"), value)
}


// A Boolean value indicating whether this capture combines image data from a dual camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isdualcamerafusionenabled
func (c_ CaptureResolvedPhotoSettings) IsDualCameraFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDualCameraFusionEnabled"))
	return rv
}


// A Boolean value indicating whether this capture combines image data from a dual camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isdualcamerafusionenabled
func (c_ CaptureResolvedPhotoSettings) SetIsDualCameraFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDualCameraFusionEnabled:"), value)
}


// A Boolean value that indicates whether the system uses fast capture prioritization when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isfastcaptureprioritizationenabled
func (c_ CaptureResolvedPhotoSettings) IsFastCapturePrioritizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFastCapturePrioritizationEnabled"))
	return rv
}


// A Boolean value that indicates whether the system uses fast capture prioritization when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isfastcaptureprioritizationenabled
func (c_ CaptureResolvedPhotoSettings) SetIsFastCapturePrioritizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFastCapturePrioritizationEnabled:"), value)
}


// A Boolean value indicating whether the camera flash fires for this capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isflashenabled
func (c_ CaptureResolvedPhotoSettings) IsFlashEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFlashEnabled"))
	return rv
}


// A Boolean value indicating whether the camera flash fires for this capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isflashenabled
func (c_ CaptureResolvedPhotoSettings) SetIsFlashEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFlashEnabled:"), value)
}


// A Boolean value indicating whether the camera automatically reduces red-eye when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isredeyereductionenabled
func (c_ CaptureResolvedPhotoSettings) IsRedEyeReductionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRedEyeReductionEnabled"))
	return rv
}


// A Boolean value indicating whether the camera automatically reduces red-eye when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isredeyereductionenabled
func (c_ CaptureResolvedPhotoSettings) SetIsRedEyeReductionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRedEyeReductionEnabled:"), value)
}


// A Boolean value indicating whether this capture uses image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isstillimagestabilizationenabled
func (c_ CaptureResolvedPhotoSettings) IsStillImageStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStillImageStabilizationEnabled"))
	return rv
}


// A Boolean value indicating whether this capture uses image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isstillimagestabilizationenabled
func (c_ CaptureResolvedPhotoSettings) SetIsStillImageStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStillImageStabilizationEnabled:"), value)
}


// A Boolean value that specifies whether the system automatically uses virtual device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isvirtualdevicefusionenabled
func (c_ CaptureResolvedPhotoSettings) IsVirtualDeviceFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDeviceFusionEnabled"))
	return rv
}


// A Boolean value that specifies whether the system automatically uses virtual device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isvirtualdevicefusionenabled
func (c_ CaptureResolvedPhotoSettings) SetIsVirtualDeviceFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDeviceFusionEnabled:"), value)
}


// The size, in pixels, of the Live Photo movie content that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/livephotomoviedimensions
func (c_ CaptureResolvedPhotoSettings) LivePhotoMovieDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("livePhotoMovieDimensions"))
	return rv
}


// The size, in pixels, of the Live Photo movie content that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/livephotomoviedimensions
func (c_ CaptureResolvedPhotoSettings) SetLivePhotoMovieDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLivePhotoMovieDimensions:"), value)
}


// The size, in pixels, of the photo image (in a processed format, such as JPEG) that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/photodimensions
func (c_ CaptureResolvedPhotoSettings) PhotoDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("photoDimensions"))
	return rv
}


// The size, in pixels, of the photo image (in a processed format, such as JPEG) that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/photodimensions
func (c_ CaptureResolvedPhotoSettings) SetPhotoDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhotoDimensions:"), value)
}


// The time range in which to expect the system to deliver the photo to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/photoprocessingtimerange
func (c_ CaptureResolvedPhotoSettings) PhotoProcessingTimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("photoProcessingTimeRange"))
	return rv
}


// The time range in which to expect the system to deliver the photo to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/photoprocessingtimerange
func (c_ CaptureResolvedPhotoSettings) SetPhotoProcessingTimeRange(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhotoProcessingTimeRange:"), value)
}


// The size, in pixels, of the portrait effects matte that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/portraiteffectsmattedimensions
func (c_ CaptureResolvedPhotoSettings) PortraitEffectsMatteDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("portraitEffectsMatteDimensions"))
	return rv
}


// The size, in pixels, of the portrait effects matte that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/portraiteffectsmattedimensions
func (c_ CaptureResolvedPhotoSettings) SetPortraitEffectsMatteDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPortraitEffectsMatteDimensions:"), value)
}


// The size, in pixels, of the preview image that the system delivers with the capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/previewdimensions
func (c_ CaptureResolvedPhotoSettings) PreviewDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("previewDimensions"))
	return rv
}


// The size, in pixels, of the preview image that the system delivers with the capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/previewdimensions
func (c_ CaptureResolvedPhotoSettings) SetPreviewDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewDimensions:"), value)
}


// The size, in pixels, of the RAW-format embedded thumbnail image that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/rawembeddedthumbnaildimensions
func (c_ CaptureResolvedPhotoSettings) RawEmbeddedThumbnailDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("rawEmbeddedThumbnailDimensions"))
	return rv
}


// The size, in pixels, of the RAW-format embedded thumbnail image that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/rawembeddedthumbnaildimensions
func (c_ CaptureResolvedPhotoSettings) SetRawEmbeddedThumbnailDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRawEmbeddedThumbnailDimensions:"), value)
}


// The size, in pixels, of the RAW-format photo image that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/rawphotodimensions
func (c_ CaptureResolvedPhotoSettings) RawPhotoDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("rawPhotoDimensions"))
	return rv
}


// The size, in pixels, of the RAW-format photo image that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/rawphotodimensions
func (c_ CaptureResolvedPhotoSettings) SetRawPhotoDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRawPhotoDimensions:"), value)
}



