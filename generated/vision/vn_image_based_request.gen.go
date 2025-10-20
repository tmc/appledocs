// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// The abstract superclass for image-analysis requests that focus on a specific part of an image.
//
// Other Vision request handlers that operate on still images inherit from this abstract base class. Don’t use it directly.
//
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

// Alloc allocates a new instance without initialization.
func (ic _ImageBasedRequestClass) Alloc() ImageBasedRequest {
	rv := objc.Send[ImageBasedRequest](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The region of the image in which Vision will perform the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageBasedRequest/regionOfInterest
func (i_ ImageBasedRequest) RegionOfInterest() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](i_.ID, objc.Sel("regionOfInterest"))
	return rv
}


// SetRegionOfInterest sets the value of the regionOfInterest property.
// The region of the image in which Vision will perform the request.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageBasedRequest/regionOfInterest
func (i_ ImageBasedRequest) SetRegionOfInterest(value coregraphics.CGRect) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRegionOfInterest:"), value)
}


