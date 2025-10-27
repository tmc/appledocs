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
	

	// properties:
	ExpectedPhotoCount() uint
	FastCapturePrioritizationEnabled() bool
	PhotoDimensions() objectivec.IObject
	UniqueID() int64
	FlashMode() objectivec.IObject
	SetFlashMode(value objectivec.IObject)
	PreviewPhotoFormat() foundation.foundation.INSString
	SetPreviewPhotoFormat(value foundation.foundation.INSString)
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


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureResolvedPhotoSettingsClass) Alloc() CaptureResolvedPhotoSettings {
	rv := objc.Send[CaptureResolvedPhotoSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

























// The number of photo capture results in the capture request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/expectedPhotoCount
func (c_ CaptureResolvedPhotoSettings) ExpectedPhotoCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("expectedPhotoCount"))
	return rv
}


// A Boolean value that indicates whether the system uses fast capture prioritization when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isFastCapturePrioritizationEnabled
func (c_ CaptureResolvedPhotoSettings) FastCapturePrioritizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fastCapturePrioritizationEnabled"))
	return rv
}


// The size, in pixels, of the photo image (in a processed format, such as JPEG) that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/photoDimensions
func (c_ CaptureResolvedPhotoSettings) PhotoDimensions() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("photoDimensions"))
	return rv
}


// The unique identifier for the photo capture this settings object corresponds to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/uniqueID
func (c_ CaptureResolvedPhotoSettings) UniqueID() int64 {
	rv := objc.Send[int64](c_.ID, objc.Sel("uniqueID"))
	return rv
}


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c_ CaptureResolvedPhotoSettings) FlashMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("flashMode"))
	return rv
}


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c_ CaptureResolvedPhotoSettings) SetFlashMode(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashMode:"), value)
}


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CaptureResolvedPhotoSettings) PreviewPhotoFormat() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("previewPhotoFormat"))
	return rv
}


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CaptureResolvedPhotoSettings) SetPreviewPhotoFormat(value foundation.foundation.INSString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewPhotoFormat:"), value)
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







