// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCapturePhotoBracketSettings */


/* debug [class_header]: Header for AVCapturePhotoBracketSettings */
// The class instance for the [CapturePhotoBracketSettings] class.
var (
	CapturePhotoBracketSettingsClass     _CapturePhotoBracketSettingsClass
	CapturePhotoBracketSettingsClassOnce sync.Once
)

func getCapturePhotoBracketSettingsClass() _CapturePhotoBracketSettingsClass {
	CapturePhotoBracketSettingsClassOnce.Do(func() {
		CapturePhotoBracketSettingsClass = _CapturePhotoBracketSettingsClass{objc.GetClass("AVCapturePhotoBracketSettings")}
	})
	return CapturePhotoBracketSettingsClass
}

type _CapturePhotoBracketSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CapturePhotoBracketSettings */
// An interface definition for the [CapturePhotoBracketSettings] class.
type ICapturePhotoBracketSettings interface {
	ICapturePhotoSettings
	
/* debug [class_interface_properties]: Properties for CapturePhotoBracketSettings */
	// properties:
	IsLensStabilizationEnabled() bool
	SetIsLensStabilizationEnabled(value bool)
	IsHighResolutionPhotoEnabled() bool
	SetIsHighResolutionPhotoEnabled(value bool)
	PreviewPhotoFormat() objc.IObject /* cross-framework: NSString */
	SetPreviewPhotoFormat(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CapturePhotoBracketSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CapturePhotoBracketSettings */
// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoBracketSettingsClass) Alloc() CapturePhotoBracketSettings {
	rv := objc.Send[CapturePhotoBracketSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CapturePhotoBracketSettingsClass) New() CapturePhotoBracketSettings {
	rv := objc.Send[CapturePhotoBracketSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhotoBracketSettings) Init() CapturePhotoBracketSettings {
	rv := objc.Send[CapturePhotoBracketSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhotoBracketSettings) Autorelease() CapturePhotoBracketSettings {
	rv := objc.Send[CapturePhotoBracketSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhotoBracketSettings creates a new CapturePhotoBracketSettings instance.
func NewCapturePhotoBracketSettings() CapturePhotoBracketSettings {
	return getCapturePhotoBracketSettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CapturePhotoBracketSettings */
// A specification of the features and settings to use for a photo capture request that captures multiple images with varied settings.
//
// To take a bracketed capture, you create and configure an object, using objects to describe the individual captures in the bracket, and then pass it to the method. To request a bracketed capture, follow these steps: Create an array of objects describing the number of images to capture in the bracket and the variations on capture settings between them. Create a bracketed photo settings object with the initializer, passing the array of bracketed still image settings, along with the processed format (such as JPEG) or RAW format to capture images in. Configure other settings to share across all images in the bracket, such as the property and certain inherited properties. Initiate capture by passing the bracketed photo settings object to your photo output’s method, along with a delegate object to receive messages about the progress and results of the capture. The photo output calls your delegate’s or methods many times corresponding to the number of captures in the bracket. Each call provides the object indicating which capture in the bracket the captured image corresponds to. The following code example illustrates capturing a bracket of three RAW images with varying exposure value settings. Listing 1. Capturing a Multi-Exposure Bracket


// A specification of the features and settings to use for a photo capture request that captures multiple images with varied settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings
type CapturePhotoBracketSettings struct {
	CapturePhotoSettings
}

// CapturePhotoBracketSettingsFrom constructs a [CapturePhotoBracketSettings] from an unsafe.Pointer.
//
// A specification of the features and settings to use for a photo capture request that captures multiple images with varied settings.
func CapturePhotoBracketSettingsFrom(ptr unsafe.Pointer) CapturePhotoBracketSettings {
	return CapturePhotoBracketSettings{
		CapturePhotoSettings: CapturePhotoSettingsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CapturePhotoBracketSettings */

// Creates a photo settings object for the specified bracket of captures, in the specified formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings/init(rawPixelFormatType:processedFormat:bracketedSettings:)
func NewCapturePhotoBracketSettingsWithRawPixelFormatTypeProcessedFormatBracketedSettings(rawPixelFormatType uint32 /* not a class type */, processedFormat foundation.IDictionary, bracketedSettings []CaptureBracketedStillImageSettings) CapturePhotoBracketSettings {
	rv := objc.Send[CapturePhotoBracketSettings](objc.ID(getCapturePhotoBracketSettingsClass().class), objc.Sel("photoBracketSettingsWithRawPixelFormatType:processedFormat:bracketedSettings:"), rawPixelFormatType, processedFormat, bracketedSettings)
	return rv
}/* debug [class_init_methods/constructor]: NewCapturePhotoBracketSettingsWithRawPixelFormatTypeProcessedFormatBracketedSettings */


// Creates a photo settings object for capture in both RAW format and a processed format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings/init(rawPixelFormatType:rawFileType:processedFormat:processedFileType:bracketedSettings:)
func NewCapturePhotoBracketSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileTypeBracketedSettings(rawPixelFormatType uint32 /* not a class type */, rawFileType FileType /* typedef */, processedFormat foundation.IDictionary, processedFileType FileType /* typedef */, bracketedSettings []CaptureBracketedStillImageSettings) CapturePhotoBracketSettings {
	rv := objc.Send[CapturePhotoBracketSettings](objc.ID(getCapturePhotoBracketSettingsClass().class), objc.Sel("photoBracketSettingsWithRawPixelFormatType:rawFileType:processedFormat:processedFileType:bracketedSettings:"), rawPixelFormatType, rawFileType, processedFormat, processedFileType, bracketedSettings)
	return rv
}/* debug [class_init_methods/constructor]: NewCapturePhotoBracketSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileTypeBracketedSettings */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CapturePhotoBracketSettings */

// Creates a photo settings object for the specified bracket of captures, in the specified formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings/init(rawPixelFormatType:processedFormat:bracketedSettings:)
func (cc _CapturePhotoBracketSettingsClass) PhotoBracketSettingsWithRawPixelFormatTypeProcessedFormatBracketedSettings(rawPixelFormatType uint32 /* not a class type */, processedFormat foundation.IDictionary, bracketedSettings []CaptureBracketedStillImageSettings) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoBracketSettingsWithRawPixelFormatType:processedFormat:bracketedSettings:"), rawPixelFormatType, processedFormat, bracketedSettings)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhotoBracketSettingsWithRawPixelFormatTypeProcessedFormatBracketedSettings) */


// Creates a photo settings object for capture in both RAW format and a processed format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings/init(rawPixelFormatType:rawFileType:processedFormat:processedFileType:bracketedSettings:)
func (cc _CapturePhotoBracketSettingsClass) PhotoBracketSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileTypeBracketedSettings(rawPixelFormatType uint32 /* not a class type */, rawFileType FileType /* typedef */, processedFormat foundation.IDictionary, processedFileType FileType /* typedef */, bracketedSettings []CaptureBracketedStillImageSettings) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("photoBracketSettingsWithRawPixelFormatType:rawFileType:processedFormat:processedFileType:bracketedSettings:"), rawPixelFormatType, rawFileType, processedFormat, processedFileType, bracketedSettings)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PhotoBracketSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileTypeBracketedSettings) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CapturePhotoBracketSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CapturePhotoBracketSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CapturePhotoBracketSettings */

// A Boolean value that specifies whether to stabilize the lens for the duration of the bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotobracketsettings/islensstabilizationenabled
func (c_ CapturePhotoBracketSettings) IsLensStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLensStabilizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isLensStabilizationEnabled */


// A Boolean value that specifies whether to stabilize the lens for the duration of the bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotobracketsettings/islensstabilizationenabled
func (c_ CapturePhotoBracketSettings) SetIsLensStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLensStabilizationEnabled:"), value)
}/* debug [instance_properties/setter]: isLensStabilizationEnabled */


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled
func (c_ CapturePhotoBracketSettings) IsHighResolutionPhotoEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighResolutionPhotoEnabled"))
	return rv
}/* debug [instance_properties/getter]: isHighResolutionPhotoEnabled */


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled
func (c_ CapturePhotoBracketSettings) SetIsHighResolutionPhotoEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighResolutionPhotoEnabled:"), value)
}/* debug [instance_properties/setter]: isHighResolutionPhotoEnabled */


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CapturePhotoBracketSettings) PreviewPhotoFormat() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("previewPhotoFormat"))
	return rv
}/* debug [instance_properties/getter]: previewPhotoFormat */


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CapturePhotoBracketSettings) SetPreviewPhotoFormat(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewPhotoFormat:"), value)
}/* debug [instance_properties/setter]: previewPhotoFormat */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCapturePhotoBracketSettings */


