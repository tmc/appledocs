// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNHomographicImageRegistrationRequest */


/* debug [class_header]: Header for VNHomographicImageRegistrationRequest */
// The class instance for the [HomographicImageRegistrationRequest] class.
var (
	HomographicImageRegistrationRequestClass     _HomographicImageRegistrationRequestClass
	HomographicImageRegistrationRequestClassOnce sync.Once
)

func getHomographicImageRegistrationRequestClass() _HomographicImageRegistrationRequestClass {
	HomographicImageRegistrationRequestClassOnce.Do(func() {
		HomographicImageRegistrationRequestClass = _HomographicImageRegistrationRequestClass{objc.GetClass("VNHomographicImageRegistrationRequest")}
	})
	return HomographicImageRegistrationRequestClass
}

type _HomographicImageRegistrationRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HomographicImageRegistrationRequest */
// An interface definition for the [HomographicImageRegistrationRequest] class.
type IHomographicImageRegistrationRequest interface {
	IImageRegistrationRequest
	
/* debug [class_interface_properties]: Properties for HomographicImageRegistrationRequest */
	// properties:
	Results() []ImageHomographicAlignmentObservation
	VNHomographicImageRegistrationRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HomographicImageRegistrationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HomographicImageRegistrationRequest */
// Alloc allocates a new instance without initialization.
func (hc _HomographicImageRegistrationRequestClass) Alloc() HomographicImageRegistrationRequest {
	rv := objc.Send[HomographicImageRegistrationRequest](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HomographicImageRegistrationRequestClass) New() HomographicImageRegistrationRequest {
	rv := objc.Send[HomographicImageRegistrationRequest](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HomographicImageRegistrationRequest) Init() HomographicImageRegistrationRequest {
	rv := objc.Send[HomographicImageRegistrationRequest](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HomographicImageRegistrationRequest) Autorelease() HomographicImageRegistrationRequest {
	rv := objc.Send[HomographicImageRegistrationRequest](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHomographicImageRegistrationRequest creates a new HomographicImageRegistrationRequest instance.
func NewHomographicImageRegistrationRequest() HomographicImageRegistrationRequest {
	return getHomographicImageRegistrationRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HomographicImageRegistrationRequest */
// An image-analysis request that determines the perspective warp matrix necessary to align the content of two images.
//
// Create and perform a homographic image registration request to align content in two images through a homography. A is an isomorphism of projected spaces, a bijection that maps lines to lines.


// An image-analysis request that determines the perspective warp matrix necessary to align the content of two images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHomographicImageRegistrationRequest
type HomographicImageRegistrationRequest struct {
	ImageRegistrationRequest
}

// HomographicImageRegistrationRequestFrom constructs a [HomographicImageRegistrationRequest] from an unsafe.Pointer.
//
// An image-analysis request that determines the perspective warp matrix necessary to align the content of two images.
func HomographicImageRegistrationRequestFrom(ptr unsafe.Pointer) HomographicImageRegistrationRequest {
	return HomographicImageRegistrationRequest{
		ImageRegistrationRequest: ImageRegistrationRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HomographicImageRegistrationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HomographicImageRegistrationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HomographicImageRegistrationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HomographicImageRegistrationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HomographicImageRegistrationRequest */

// The results of the image registration request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHomographicImageRegistrationRequest/results
func (h_ HomographicImageRegistrationRequest) Results() []ImageHomographicAlignmentObservation {
	rv := objc.Send[[]ImageHomographicAlignmentObservation](h_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the homographic image registration request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnhomographicimageregistrationrequestrevision1
func (h_ HomographicImageRegistrationRequest) VNHomographicImageRegistrationRequestRevision1() int {
	rv := objc.Send[int](h_.ID, objc.Sel("VNHomographicImageRegistrationRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNHomographicImageRegistrationRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNHomographicImageRegistrationRequest */



