// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TargetedImageRequest] class.
var (
	TargetedImageRequestClass     _TargetedImageRequestClass
	TargetedImageRequestClassOnce sync.Once
)

func getTargetedImageRequestClass() _TargetedImageRequestClass {
	TargetedImageRequestClassOnce.Do(func() {
		TargetedImageRequestClass = _TargetedImageRequestClass{objc.GetClass("VNTargetedImageRequest")}
	})
	return TargetedImageRequestClass
}

type _TargetedImageRequestClass struct {
	class objc.Class
}

// An interface definition for the [TargetedImageRequest] class.
type ITargetedImageRequest interface {
	IImageBasedRequest
	// properties:
	// methods:
}

// The abstract superclass for image analysis requests that operate on both the processed image and a secondary image.
//
// Other Vision request handlers that operate on both the processed image and a secondary image inherit from this abstract base class. Instantiate one of its subclasses to perform image analysis, and pass in auxiliary image data by filling in the dictionary at initialization.


// The abstract superclass for image analysis requests that operate on both the processed image and a secondary image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest
type TargetedImageRequest struct {
	ImageBasedRequest
}

// TargetedImageRequestFrom constructs a [TargetedImageRequest] from an unsafe.Pointer.
//
// The abstract superclass for image analysis requests that operate on both the processed image and a secondary image.
func TargetedImageRequestFrom(ptr unsafe.Pointer) TargetedImageRequest {
	return TargetedImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TargetedImageRequestClass) Alloc() TargetedImageRequest {
	rv := objc.Send[TargetedImageRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TargetedImageRequestClass) New() TargetedImageRequest {
	rv := objc.Send[TargetedImageRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TargetedImageRequest) Init() TargetedImageRequest {
	rv := objc.Send[TargetedImageRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TargetedImageRequest) Autorelease() TargetedImageRequest {
	rv := objc.Send[TargetedImageRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTargetedImageRequest creates a new TargetedImageRequest instance.
func NewTargetedImageRequest() TargetedImageRequest {
	return getTargetedImageRequestClass().New()
}




