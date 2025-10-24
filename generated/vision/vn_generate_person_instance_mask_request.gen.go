// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNGeneratePersonInstanceMaskRequest */


/* debug [class_header]: Header for VNGeneratePersonInstanceMaskRequest */
// The class instance for the [GeneratePersonInstanceMaskRequest] class.
var (
	GeneratePersonInstanceMaskRequestClass     _GeneratePersonInstanceMaskRequestClass
	GeneratePersonInstanceMaskRequestClassOnce sync.Once
)

func getGeneratePersonInstanceMaskRequestClass() _GeneratePersonInstanceMaskRequestClass {
	GeneratePersonInstanceMaskRequestClassOnce.Do(func() {
		GeneratePersonInstanceMaskRequestClass = _GeneratePersonInstanceMaskRequestClass{objc.GetClass("VNGeneratePersonInstanceMaskRequest")}
	})
	return GeneratePersonInstanceMaskRequestClass
}

type _GeneratePersonInstanceMaskRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GeneratePersonInstanceMaskRequest */
// An interface definition for the [GeneratePersonInstanceMaskRequest] class.
type IGeneratePersonInstanceMaskRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for GeneratePersonInstanceMaskRequest */
	// properties:
	Results() []InstanceMaskObservation
	VNGeneratePersonInstanceMaskRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GeneratePersonInstanceMaskRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GeneratePersonInstanceMaskRequest */
// Alloc allocates a new instance without initialization.
func (gc _GeneratePersonInstanceMaskRequestClass) Alloc() GeneratePersonInstanceMaskRequest {
	rv := objc.Send[GeneratePersonInstanceMaskRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GeneratePersonInstanceMaskRequestClass) New() GeneratePersonInstanceMaskRequest {
	rv := objc.Send[GeneratePersonInstanceMaskRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GeneratePersonInstanceMaskRequest) Init() GeneratePersonInstanceMaskRequest {
	rv := objc.Send[GeneratePersonInstanceMaskRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GeneratePersonInstanceMaskRequest) Autorelease() GeneratePersonInstanceMaskRequest {
	rv := objc.Send[GeneratePersonInstanceMaskRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGeneratePersonInstanceMaskRequest creates a new GeneratePersonInstanceMaskRequest instance.
func NewGeneratePersonInstanceMaskRequest() GeneratePersonInstanceMaskRequest {
	return getGeneratePersonInstanceMaskRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GeneratePersonInstanceMaskRequest */
// An object that produces a mask of individual people it finds in the input image.


// An object that produces a mask of individual people it finds in the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonInstanceMaskRequest
type GeneratePersonInstanceMaskRequest struct {
	ImageBasedRequest
}

// GeneratePersonInstanceMaskRequestFrom constructs a [GeneratePersonInstanceMaskRequest] from an unsafe.Pointer.
//
// An object that produces a mask of individual people it finds in the input image.
func GeneratePersonInstanceMaskRequestFrom(ptr unsafe.Pointer) GeneratePersonInstanceMaskRequest {
	return GeneratePersonInstanceMaskRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GeneratePersonInstanceMaskRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GeneratePersonInstanceMaskRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GeneratePersonInstanceMaskRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GeneratePersonInstanceMaskRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GeneratePersonInstanceMaskRequest */

// The results of the instance mask request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonInstanceMaskRequest/results
func (g_ GeneratePersonInstanceMaskRequest) Results() []InstanceMaskObservation {
	rv := objc.Send[[]InstanceMaskObservation](g_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the person instance mask request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersoninstancemaskrequestrevision1
func (g_ GeneratePersonInstanceMaskRequest) VNGeneratePersonInstanceMaskRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGeneratePersonInstanceMaskRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNGeneratePersonInstanceMaskRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNGeneratePersonInstanceMaskRequest */



