// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ImageBasedRequest] class.
var (
	ImageBasedRequestClass     _ImageBasedRequestClass
	ImageBasedRequestClassOnce sync.Once
)

func getImageBasedRequestClass() _ImageBasedRequestClass {
	ImageBasedRequestClassOnce.Do(func() {
		ImageBasedRequestClass = _ImageBasedRequestClass{objc.GetClass("VNImageBasedRequest")}
	})
	return ImageBasedRequestClass
}

type _ImageBasedRequestClass struct {
	class objc.Class
}





// An interface definition for the [ImageBasedRequest] class.
type IImageBasedRequest interface {
	IRequest
	

	// properties:
	RegionOfInterest() corefoundation.CGRect
	SetRegionOfInterest(value corefoundation.CGRect)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageBasedRequestClass) Alloc() ImageBasedRequest {
	rv := objc.Send[ImageBasedRequest](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageBasedRequestClass) New() ImageBasedRequest {
	rv := objc.Send[ImageBasedRequest](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageBasedRequest) Init() ImageBasedRequest {
	rv := objc.Send[ImageBasedRequest](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageBasedRequest) Autorelease() ImageBasedRequest {
	rv := objc.Send[ImageBasedRequest](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageBasedRequest creates a new ImageBasedRequest instance.
func NewImageBasedRequest() ImageBasedRequest {
	return getImageBasedRequestClass().New()
}





// The abstract superclass for image-analysis requests that focus on a specific part of an image.
//
// Other Vision request handlers that operate on still images inherit from this abstract base class. Don’t use it directly.


// The abstract superclass for image-analysis requests that focus on a specific part of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageBasedRequest
type ImageBasedRequest struct {
	Request
}

// ImageBasedRequestFrom constructs a [ImageBasedRequest] from an unsafe.Pointer.
//
// The abstract superclass for image-analysis requests that focus on a specific part of an image.
func ImageBasedRequestFrom(ptr unsafe.Pointer) ImageBasedRequest {
	return ImageBasedRequest{
		Request: RequestFrom(ptr),
	}
}

























// The region of the image in which Vision will perform the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageBasedRequest/regionOfInterest
func (i_ ImageBasedRequest) RegionOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i_.ID, objc.Sel("regionOfInterest"))
	return rv
}


// The region of the image in which Vision will perform the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageBasedRequest/regionOfInterest
func (i_ ImageBasedRequest) SetRegionOfInterest(value corefoundation.CGRect) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRegionOfInterest:"), value)
}








