// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureResolvedPhotoSettings */


/* debug [class_header]: Header for AVCaptureResolvedPhotoSettings */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureResolvedPhotoSettings */
// An interface definition for the [CaptureResolvedPhotoSettings] class.
type ICaptureResolvedPhotoSettings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureResolvedPhotoSettings */
	// properties:
	ExpectedPhotoCount() uint
	FastCapturePrioritizationEnabled() bool
	PhotoDimensions() VideoDimensions /* not a class type */
	UniqueID() int64
	FlashMode() objectivec.IObject
	SetFlashMode(value objectivec.IObject)
	PreviewPhotoFormat() objc.IObject /* cross-framework: NSString */
	SetPreviewPhotoFormat(value objc.IObject /* cross-framework: NSString */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureResolvedPhotoSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureResolvedPhotoSettings */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureResolvedPhotoSettings */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureResolvedPhotoSettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureResolvedPhotoSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureResolvedPhotoSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureResolvedPhotoSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureResolvedPhotoSettings */

// The number of photo capture results in the capture request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/expectedPhotoCount
func (c_ CaptureResolvedPhotoSettings) ExpectedPhotoCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("expectedPhotoCount"))
	return rv
}/* debug [instance_properties/getter]: expectedPhotoCount */


// A Boolean value that indicates whether the system uses fast capture prioritization when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isFastCapturePrioritizationEnabled
func (c_ CaptureResolvedPhotoSettings) FastCapturePrioritizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fastCapturePrioritizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: fastCapturePrioritizationEnabled */


// The size, in pixels, of the photo image (in a processed format, such as JPEG) that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/photoDimensions
func (c_ CaptureResolvedPhotoSettings) PhotoDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("photoDimensions"))
	return rv
}/* debug [instance_properties/getter]: photoDimensions */


// The unique identifier for the photo capture this settings object corresponds to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/uniqueID
func (c_ CaptureResolvedPhotoSettings) UniqueID() int64 {
	rv := objc.Send[int64](c_.ID, objc.Sel("uniqueID"))
	return rv
}/* debug [instance_properties/getter]: uniqueID */


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c_ CaptureResolvedPhotoSettings) FlashMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("flashMode"))
	return rv
}/* debug [instance_properties/getter]: flashMode */


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c_ CaptureResolvedPhotoSettings) SetFlashMode(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashMode:"), value)
}/* debug [instance_properties/setter]: flashMode */


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CaptureResolvedPhotoSettings) PreviewPhotoFormat() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("previewPhotoFormat"))
	return rv
}/* debug [instance_properties/getter]: previewPhotoFormat */


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CaptureResolvedPhotoSettings) SetPreviewPhotoFormat(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewPhotoFormat:"), value)
}/* debug [instance_properties/setter]: previewPhotoFormat */


// A Boolean value that indicates whether the system applies content-aware distortion correction when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/iscontentawaredistortioncorrectionenabled
func (c_ CaptureResolvedPhotoSettings) IsContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContentAwareDistortionCorrectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isContentAwareDistortionCorrectionEnabled */


// A Boolean value that indicates whether the system applies content-aware distortion correction when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/iscontentawaredistortioncorrectionenabled
func (c_ CaptureResolvedPhotoSettings) SetIsContentAwareDistortionCorrectionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContentAwareDistortionCorrectionEnabled:"), value)
}/* debug [instance_properties/setter]: isContentAwareDistortionCorrectionEnabled */


// A Boolean value indicating whether this capture combines image data from a dual camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isdualcamerafusionenabled
func (c_ CaptureResolvedPhotoSettings) IsDualCameraFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDualCameraFusionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDualCameraFusionEnabled */


// A Boolean value indicating whether this capture combines image data from a dual camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isdualcamerafusionenabled
func (c_ CaptureResolvedPhotoSettings) SetIsDualCameraFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDualCameraFusionEnabled:"), value)
}/* debug [instance_properties/setter]: isDualCameraFusionEnabled */


// A Boolean value that indicates whether the system uses fast capture prioritization when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isfastcaptureprioritizationenabled
func (c_ CaptureResolvedPhotoSettings) IsFastCapturePrioritizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFastCapturePrioritizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isFastCapturePrioritizationEnabled */


// A Boolean value that indicates whether the system uses fast capture prioritization when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isfastcaptureprioritizationenabled
func (c_ CaptureResolvedPhotoSettings) SetIsFastCapturePrioritizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFastCapturePrioritizationEnabled:"), value)
}/* debug [instance_properties/setter]: isFastCapturePrioritizationEnabled */


// A Boolean value indicating whether the camera flash fires for this capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isflashenabled
func (c_ CaptureResolvedPhotoSettings) IsFlashEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFlashEnabled"))
	return rv
}/* debug [instance_properties/getter]: isFlashEnabled */


// A Boolean value indicating whether the camera flash fires for this capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isflashenabled
func (c_ CaptureResolvedPhotoSettings) SetIsFlashEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFlashEnabled:"), value)
}/* debug [instance_properties/setter]: isFlashEnabled */


// A Boolean value indicating whether the camera automatically reduces red-eye when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isredeyereductionenabled
func (c_ CaptureResolvedPhotoSettings) IsRedEyeReductionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRedEyeReductionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isRedEyeReductionEnabled */


// A Boolean value indicating whether the camera automatically reduces red-eye when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isredeyereductionenabled
func (c_ CaptureResolvedPhotoSettings) SetIsRedEyeReductionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRedEyeReductionEnabled:"), value)
}/* debug [instance_properties/setter]: isRedEyeReductionEnabled */


// A Boolean value indicating whether this capture uses image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isstillimagestabilizationenabled
func (c_ CaptureResolvedPhotoSettings) IsStillImageStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStillImageStabilizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isStillImageStabilizationEnabled */


// A Boolean value indicating whether this capture uses image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isstillimagestabilizationenabled
func (c_ CaptureResolvedPhotoSettings) SetIsStillImageStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStillImageStabilizationEnabled:"), value)
}/* debug [instance_properties/setter]: isStillImageStabilizationEnabled */


// A Boolean value that specifies whether the system automatically uses virtual device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isvirtualdevicefusionenabled
func (c_ CaptureResolvedPhotoSettings) IsVirtualDeviceFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDeviceFusionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isVirtualDeviceFusionEnabled */


// A Boolean value that specifies whether the system automatically uses virtual device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/isvirtualdevicefusionenabled
func (c_ CaptureResolvedPhotoSettings) SetIsVirtualDeviceFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDeviceFusionEnabled:"), value)
}/* debug [instance_properties/setter]: isVirtualDeviceFusionEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureResolvedPhotoSettings */


