// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [CapturePhoto] class.
var (
	CapturePhotoClass     _CapturePhotoClass
	CapturePhotoClassOnce sync.Once
)

func getCapturePhotoClass() _CapturePhotoClass {
	CapturePhotoClassOnce.Do(func() {
		CapturePhotoClass = _CapturePhotoClass{objc.GetClass("AVCapturePhoto")}
	})
	return CapturePhotoClass
}

type _CapturePhotoClass struct {
	class objc.Class
}

// An interface definition for the [CapturePhoto] class.
type ICapturePhoto interface {
	objectivec.IObject
	CGImageRepresentation() coregraphics.CGImageRef
	FileDataRepresentation() unsafe.Pointer
}

// A container for image data from a photo capture output.
//
// When you capture photos with the class, your delegate object receives each resulting image and related data in the form of an object. This object is an immutable wrapper from which you can retrieve various results of the photo capture. In addition to the photo image pixel buffer, an AVCapturePhoto object can also contain a preview-sized pixel buffer, capture metadata, and, on supported devices, depth data and camera calibration data. From an object, you can generate data appropriate for writing to a file, such as HEVC encoded image data containerized in the HEIC file format and including a preview image, depth data and other attachments. An instance wraps a single image result. For example, if you request a bracketed capture of three images, your callback is called three times, each time delivering a single object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto
type CapturePhoto struct {
	objectivec.Object
}

// CapturePhotoFrom constructs a [CapturePhoto] from an unsafe.Pointer.
//
// A container for image data from a photo capture output.
func CapturePhotoFrom(ptr unsafe.Pointer) CapturePhoto {
	return CapturePhoto{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoClass) Alloc() CapturePhoto {
	rv := objc.Send[CapturePhoto](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CapturePhotoClass) New() CapturePhoto {
	rv := objc.Send[CapturePhoto](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhoto) Init() CapturePhoto {
	rv := objc.Send[CapturePhoto](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhoto) Autorelease() CapturePhoto {
	rv := objc.Send[CapturePhoto](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhoto creates a new CapturePhoto instance.
func NewCapturePhoto() CapturePhoto {
	return getCapturePhotoClass().New()
}


// Extracts and returns the captured photo’s primary image as a Core Graphics image object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/cgImageRepresentation()
func (c_ CapturePhoto) CGImageRepresentation() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c_.ID, objc.Sel("CGImageRepresentation"))
	return rv
}

// Generates and returns a flat data representation of the photo and its attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/fileDataRepresentation()
func (c_ CapturePhoto) FileDataRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fileDataRepresentation"))
	return rv
}

// A Boolean value indicating whether this photo object contains RAW format data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/isRawPhoto
func (c_ CapturePhoto) RawPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rawPhoto"))
	return rv
}

// The uncompressed or RAW image sample buffer for the photo, if requested.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/pixelBuffer
func (c_ CapturePhoto) PixelBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("pixelBuffer"))
	return rv
}



