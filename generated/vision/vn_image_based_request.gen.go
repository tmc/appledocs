// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

/* debug [class.gen.go]: Generating class VNImageBasedRequest */


/* debug [class_header]: Header for VNImageBasedRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageBasedRequest */
// An interface definition for the [ImageBasedRequest] class.
type IImageBasedRequest interface {
	IRequest
	
/* debug [class_interface_properties]: Properties for ImageBasedRequest */
	// properties:
	RegionOfInterest() corefoundation.CGRect
	SetRegionOfInterest(value corefoundation.CGRect)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageBasedRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageBasedRequest */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageBasedRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageBasedRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageBasedRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageBasedRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageBasedRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageBasedRequest */

// The region of the image in which Vision will perform the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageBasedRequest/regionOfInterest
func (i_ ImageBasedRequest) RegionOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i_.ID, objc.Sel("regionOfInterest"))
	return rv
}/* debug [instance_properties/getter]: regionOfInterest */


// The region of the image in which Vision will perform the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageBasedRequest/regionOfInterest
func (i_ ImageBasedRequest) SetRegionOfInterest(value corefoundation.CGRect) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRegionOfInterest:"), value)
}/* debug [instance_properties/setter]: regionOfInterest */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNImageBasedRequest */



