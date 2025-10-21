// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// A description of the features and settings in use for an in-progress or complete photo capture request.
//
// When you request a photo capture using the method, you describe the settings for that capture request in an object. When the capture begins, the photo output calls your delegate methods and provides an object detailing the settings that are in effect for that capture. Resolved photo settings objects are immutable; they describe a request that has already been made. The property of a resolved photo settings object passed to one of your methods matches the value of the object you passed when requesting capture. Use this value to determine which delegate method calls correspond to which capture requests. Some photo capture settings are automatic, such as the property. For such settings, the photo output determines whether to use that feature at the moment of capture—you don’t know when requesting a capture whether the feature is active when the capture completes. When the photo output calls your delegate methods, the provided object details which automatic features have been set for that capture. Likewise, the dimensions of an output image or movie may not be set until the moment of capture. For example, when you specify a thumbnail size with the setting, the photo output chooses dimensions that best match your requested size while preserving the aspect ratio of the captured photo. When the photo output calls your delegate methods, use the property of the resolved settings to find the actual preview image dimensions. See the methods listed in Examining Output Dimensions for other cases where output dimensions can change at capture time.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/embeddedThumbnailDimensions
func (c_ CaptureResolvedPhotoSettings) EmbeddedThumbnailDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("embeddedThumbnailDimensions"))
	return rv
}



