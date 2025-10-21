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



