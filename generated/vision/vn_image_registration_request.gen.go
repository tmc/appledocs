// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ImageRegistrationRequest] class.
var (
	ImageRegistrationRequestClass     _ImageRegistrationRequestClass
	ImageRegistrationRequestClassOnce sync.Once
)

func getImageRegistrationRequestClass() _ImageRegistrationRequestClass {
	ImageRegistrationRequestClassOnce.Do(func() {
		ImageRegistrationRequestClass = _ImageRegistrationRequestClass{objc.GetClass("VNImageRegistrationRequest")}
	})
	return ImageRegistrationRequestClass
}

type _ImageRegistrationRequestClass struct {
	class objc.Class
}

// An interface definition for the [ImageRegistrationRequest] class.
type IImageRegistrationRequest interface {
	ITargetedImageRequest
	// properties:
	// methods:
}

// The abstract superclass for image-analysis requests that align images according to their content.
//
// This abstract superclass forms the basis of image alignment or registration requests. Make specific requests through one of its subclasses, or . Don’t create an instance of this superclass yourself.


// The abstract superclass for image-analysis requests that align images according to their content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRegistrationRequest
type ImageRegistrationRequest struct {
	TargetedImageRequest
}

// ImageRegistrationRequestFrom constructs a [ImageRegistrationRequest] from an unsafe.Pointer.
//
// The abstract superclass for image-analysis requests that align images according to their content.
func ImageRegistrationRequestFrom(ptr unsafe.Pointer) ImageRegistrationRequest {
	return ImageRegistrationRequest{
		TargetedImageRequest: TargetedImageRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageRegistrationRequestClass) Alloc() ImageRegistrationRequest {
	rv := objc.Send[ImageRegistrationRequest](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageRegistrationRequestClass) New() ImageRegistrationRequest {
	rv := objc.Send[ImageRegistrationRequest](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageRegistrationRequest) Init() ImageRegistrationRequest {
	rv := objc.Send[ImageRegistrationRequest](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageRegistrationRequest) Autorelease() ImageRegistrationRequest {
	rv := objc.Send[ImageRegistrationRequest](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageRegistrationRequest creates a new ImageRegistrationRequest instance.
func NewImageRegistrationRequest() ImageRegistrationRequest {
	return getImageRegistrationRequestClass().New()
}




