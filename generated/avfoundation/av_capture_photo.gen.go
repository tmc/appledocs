// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCapturePhoto */


/* debug [class_header]: Header for AVCapturePhoto */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CapturePhoto */
// An interface definition for the [CapturePhoto] class.
type ICapturePhoto interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CapturePhoto */
	// properties:
	ConstantColorCenterWeightedMeanConfidenceLevel() float32
	ConstantColorConfidenceMap() PixelBufferRef /* not a class type */
	ConstantColorFallbackPhoto() bool
	PhotoCount() int
	PixelBuffer() PixelBufferRef /* not a class type */
	ResolvedSettings() IAVCaptureResolvedPhotoSettings
	Timestamp() objc.IObject /* cross-framework: Time */
	IsConstantColorFallbackPhoto() bool
	SetIsConstantColorFallbackPhoto(value bool)
	IsRawPhoto() bool
	SetIsRawPhoto(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CapturePhoto */
	// methods:
	CGImageRepresentation() ImageRef /* not a class type */
	FileDataRepresentation() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CapturePhoto */
// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoClass) Alloc() CapturePhoto {
	rv := objc.Send[CapturePhoto](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CapturePhoto */
// A container for image data from a photo capture output.
//
// When you capture photos with the class, your delegate object receives each resulting image and related data in the form of an object. This object is an immutable wrapper from which you can retrieve various results of the photo capture. In addition to the photo image pixel buffer, an AVCapturePhoto object can also contain a preview-sized pixel buffer, capture metadata, and, on supported devices, depth data and camera calibration data. From an object, you can generate data appropriate for writing to a file, such as HEVC encoded image data containerized in the HEIC file format and including a preview image, depth data and other attachments. An instance wraps a single image result. For example, if you request a bracketed capture of three images, your callback is called three times, each time delivering a single object.


// A container for image data from a photo capture output.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CapturePhoto *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CapturePhoto */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CapturePhoto */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CapturePhoto */

// Extracts and returns the captured photo’s primary image as a Core Graphics image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/cgImageRepresentation()
func (c_ CapturePhoto) CGImageRepresentation() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](c_.ID, objc.Sel("CGImageRepresentation"))
	return rv
}/* debug [instance_methods/method]: CGImageRepresentation */


// Generates and returns a flat data representation of the photo and its attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/fileDataRepresentation()
func (c_ CapturePhoto) FileDataRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("fileDataRepresentation"))
	return rv
}/* debug [instance_methods/method]: FileDataRepresentation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CapturePhoto */

// A score that summarizes the overall confidence level of a constant color photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/constantColorCenterWeightedMeanConfidenceLevel
func (c_ CapturePhoto) ConstantColorCenterWeightedMeanConfidenceLevel() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("constantColorCenterWeightedMeanConfidenceLevel"))
	return rv
}/* debug [instance_properties/getter]: constantColorCenterWeightedMeanConfidenceLevel */


// A pixel buffer where each pixel value indicates how fully the system achieves the constant color effect in the corresponding region of the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/constantColorConfidenceMap
func (c_ CapturePhoto) ConstantColorConfidenceMap() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](c_.ID, objc.Sel("constantColorConfidenceMap"))
	return rv
}/* debug [instance_properties/getter]: constantColorConfidenceMap */


// A Boolean value that Indicates whether this photo is a fallback photo for a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/isConstantColorFallbackPhoto
func (c_ CapturePhoto) ConstantColorFallbackPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("constantColorFallbackPhoto"))
	return rv
}/* debug [instance_properties/getter]: constantColorFallbackPhoto */


// The 1-based index of this photo capture relative to other results from the same capture request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/photoCount
func (c_ CapturePhoto) PhotoCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("photoCount"))
	return rv
}/* debug [instance_properties/getter]: photoCount */


// The uncompressed or RAW image sample buffer for the photo, if requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/pixelBuffer
func (c_ CapturePhoto) PixelBuffer() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](c_.ID, objc.Sel("pixelBuffer"))
	return rv
}/* debug [instance_properties/getter]: pixelBuffer */


// The settings object that was used to request this photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/resolvedSettings
func (c_ CapturePhoto) ResolvedSettings() IAVCaptureResolvedPhotoSettings {
	rv := objc.Send[CaptureResolvedPhotoSettings](c_.ID, objc.Sel("resolvedSettings"))
	return rv
}/* debug [instance_properties/getter]: resolvedSettings */


// The time at which the image was captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/timestamp
func (c_ CapturePhoto) Timestamp() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("timestamp"))
	return rv
}/* debug [instance_properties/getter]: timestamp */


// A Boolean value that Indicates whether this photo is a fallback photo for a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/isconstantcolorfallbackphoto
func (c_ CapturePhoto) IsConstantColorFallbackPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorFallbackPhoto"))
	return rv
}/* debug [instance_properties/getter]: isConstantColorFallbackPhoto */


// A Boolean value that Indicates whether this photo is a fallback photo for a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/isconstantcolorfallbackphoto
func (c_ CapturePhoto) SetIsConstantColorFallbackPhoto(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorFallbackPhoto:"), value)
}/* debug [instance_properties/setter]: isConstantColorFallbackPhoto */


// A Boolean value indicating whether this photo object contains RAW format data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/israwphoto
func (c_ CapturePhoto) IsRawPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRawPhoto"))
	return rv
}/* debug [instance_properties/getter]: isRawPhoto */


// A Boolean value indicating whether this photo object contains RAW format data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/israwphoto
func (c_ CapturePhoto) SetIsRawPhoto(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRawPhoto:"), value)
}/* debug [instance_properties/setter]: isRawPhoto */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCapturePhoto */


