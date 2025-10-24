// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNImageRegistrationRequest */


/* debug [class_header]: Header for VNImageRegistrationRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageRegistrationRequest */
// An interface definition for the [ImageRegistrationRequest] class.
type IImageRegistrationRequest interface {
	ITargetedImageRequest
	
/* debug [class_interface_properties]: Properties for ImageRegistrationRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageRegistrationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageRegistrationRequest */
// Alloc allocates a new instance without initialization.
func (ic _ImageRegistrationRequestClass) Alloc() ImageRegistrationRequest {
	rv := objc.Send[ImageRegistrationRequest](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageRegistrationRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageRegistrationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageRegistrationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageRegistrationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageRegistrationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageRegistrationRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNImageRegistrationRequest */



