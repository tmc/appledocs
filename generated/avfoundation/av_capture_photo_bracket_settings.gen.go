// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CapturePhotoBracketSettings] class.
type ICapturePhotoBracketSettings interface {
	ICapturePhotoSettings
	BracketedSettings() []unsafe.Pointer
	IsLensStabilizationEnabled() bool
	SetIsLensStabilizationEnabled(value bool)
	IsHighResolutionPhotoEnabled() bool
	SetIsHighResolutionPhotoEnabled(value bool)
	PreviewPhotoFormat() string
	SetPreviewPhotoFormat(value string)
}

// A specification of the features and settings to use for a photo capture request that captures multiple images with varied settings.
//
// To take a bracketed capture, you create and configure an object, using objects to describe the individual captures in the bracket, and then pass it to the method. To request a bracketed capture, follow these steps: Create an array of objects describing the number of images to capture in the bracket and the variations on capture settings between them. Create a bracketed photo settings object with the initializer, passing the array of bracketed still image settings, along with the processed format (such as JPEG) or RAW format to capture images in. Configure other settings to share across all images in the bracket, such as the property and certain inherited properties. Initiate capture by passing the bracketed photo settings object to your photo output’s method, along with a delegate object to receive messages about the progress and results of the capture. The photo output calls your delegate’s or methods many times corresponding to the number of captures in the bracket. Each call provides the object indicating which capture in the bracket the captured image corresponds to. The following code example illustrates capturing a bracket of three RAW images with varying exposure value settings. Listing 1. Capturing a Multi-Exposure Bracket
//
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

// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoBracketSettingsClass) Alloc() CapturePhotoBracketSettings {
	rv := objc.Send[CapturePhotoBracketSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// An array describing the number of and settings for images to produce in a bracketed capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings/bracketedSettings
func (c_ CapturePhotoBracketSettings) BracketedSettings() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](c_.ID, objc.Sel("bracketedSettings"))
	return rv
}

// A Boolean value that specifies whether to stabilize the lens for the duration of the bracketed capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotobracketsettings/islensstabilizationenabled
func (c_ CapturePhotoBracketSettings) IsLensStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLensStabilizationEnabled"))
	return rv
}


// SetIsLensStabilizationEnabled sets the value of the isLensStabilizationEnabled property.
// A Boolean value that specifies whether to stabilize the lens for the duration of the bracketed capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotobracketsettings/islensstabilizationenabled
func (c_ CapturePhotoBracketSettings) SetIsLensStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLensStabilizationEnabled:"), value)
}

// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled
func (c_ CapturePhotoBracketSettings) IsHighResolutionPhotoEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighResolutionPhotoEnabled"))
	return rv
}


// SetIsHighResolutionPhotoEnabled sets the value of the isHighResolutionPhotoEnabled property.
// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled
func (c_ CapturePhotoBracketSettings) SetIsHighResolutionPhotoEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighResolutionPhotoEnabled:"), value)
}

// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CapturePhotoBracketSettings) PreviewPhotoFormat() string {
	rv := objc.Send[string](c_.ID, objc.Sel("previewPhotoFormat"))
	return rv
}


// SetPreviewPhotoFormat sets the value of the previewPhotoFormat property.
// A dictionary describing the format for delivery of preview-sized images alongside the main photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c_ CapturePhotoBracketSettings) SetPreviewPhotoFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewPhotoFormat:"), objc.String(value))
}



