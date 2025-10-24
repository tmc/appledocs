// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNCalculateImageAestheticsScoresRequest */


/* debug [class_header]: Header for VNCalculateImageAestheticsScoresRequest */
// The class instance for the [CalculateImageAestheticsScoresRequest] class.
var (
	CalculateImageAestheticsScoresRequestClass     _CalculateImageAestheticsScoresRequestClass
	CalculateImageAestheticsScoresRequestClassOnce sync.Once
)

func getCalculateImageAestheticsScoresRequestClass() _CalculateImageAestheticsScoresRequestClass {
	CalculateImageAestheticsScoresRequestClassOnce.Do(func() {
		CalculateImageAestheticsScoresRequestClass = _CalculateImageAestheticsScoresRequestClass{objc.GetClass("VNCalculateImageAestheticsScoresRequest")}
	})
	return CalculateImageAestheticsScoresRequestClass
}

type _CalculateImageAestheticsScoresRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CalculateImageAestheticsScoresRequest */
// An interface definition for the [CalculateImageAestheticsScoresRequest] class.
type ICalculateImageAestheticsScoresRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for CalculateImageAestheticsScoresRequest */
	// properties:
	Results() []ImageAestheticsScoresObservation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CalculateImageAestheticsScoresRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CalculateImageAestheticsScoresRequest */
// Alloc allocates a new instance without initialization.
func (cc _CalculateImageAestheticsScoresRequestClass) Alloc() CalculateImageAestheticsScoresRequest {
	rv := objc.Send[CalculateImageAestheticsScoresRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CalculateImageAestheticsScoresRequestClass) New() CalculateImageAestheticsScoresRequest {
	rv := objc.Send[CalculateImageAestheticsScoresRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CalculateImageAestheticsScoresRequest) Init() CalculateImageAestheticsScoresRequest {
	rv := objc.Send[CalculateImageAestheticsScoresRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CalculateImageAestheticsScoresRequest) Autorelease() CalculateImageAestheticsScoresRequest {
	rv := objc.Send[CalculateImageAestheticsScoresRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCalculateImageAestheticsScoresRequest creates a new CalculateImageAestheticsScoresRequest instance.
func NewCalculateImageAestheticsScoresRequest() CalculateImageAestheticsScoresRequest {
	return getCalculateImageAestheticsScoresRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CalculateImageAestheticsScoresRequest */
// An object that analyzes an image for aesthetically pleasing attributes.


// An object that analyzes an image for aesthetically pleasing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCalculateImageAestheticsScoresRequest
type CalculateImageAestheticsScoresRequest struct {
	ImageBasedRequest
}

// CalculateImageAestheticsScoresRequestFrom constructs a [CalculateImageAestheticsScoresRequest] from an unsafe.Pointer.
//
// An object that analyzes an image for aesthetically pleasing attributes.
func CalculateImageAestheticsScoresRequestFrom(ptr unsafe.Pointer) CalculateImageAestheticsScoresRequest {
	return CalculateImageAestheticsScoresRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CalculateImageAestheticsScoresRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CalculateImageAestheticsScoresRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CalculateImageAestheticsScoresRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CalculateImageAestheticsScoresRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CalculateImageAestheticsScoresRequest */

// The results of the aesthetics request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCalculateImageAestheticsScoresRequest/results
func (c_ CalculateImageAestheticsScoresRequest) Results() []ImageAestheticsScoresObservation {
	rv := objc.Send[[]ImageAestheticsScoresObservation](c_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNCalculateImageAestheticsScoresRequest */



