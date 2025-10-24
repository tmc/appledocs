// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNGenerateForegroundInstanceMaskRequest */


/* debug [class_header]: Header for VNGenerateForegroundInstanceMaskRequest */
// The class instance for the [GenerateForegroundInstanceMaskRequest] class.
var (
	GenerateForegroundInstanceMaskRequestClass     _GenerateForegroundInstanceMaskRequestClass
	GenerateForegroundInstanceMaskRequestClassOnce sync.Once
)

func getGenerateForegroundInstanceMaskRequestClass() _GenerateForegroundInstanceMaskRequestClass {
	GenerateForegroundInstanceMaskRequestClassOnce.Do(func() {
		GenerateForegroundInstanceMaskRequestClass = _GenerateForegroundInstanceMaskRequestClass{objc.GetClass("VNGenerateForegroundInstanceMaskRequest")}
	})
	return GenerateForegroundInstanceMaskRequestClass
}

type _GenerateForegroundInstanceMaskRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GenerateForegroundInstanceMaskRequest */
// An interface definition for the [GenerateForegroundInstanceMaskRequest] class.
type IGenerateForegroundInstanceMaskRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for GenerateForegroundInstanceMaskRequest */
	// properties:
	Results() []InstanceMaskObservation
	VNGenerateForegroundInstanceMaskRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GenerateForegroundInstanceMaskRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GenerateForegroundInstanceMaskRequest */
// Alloc allocates a new instance without initialization.
func (gc _GenerateForegroundInstanceMaskRequestClass) Alloc() GenerateForegroundInstanceMaskRequest {
	rv := objc.Send[GenerateForegroundInstanceMaskRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GenerateForegroundInstanceMaskRequestClass) New() GenerateForegroundInstanceMaskRequest {
	rv := objc.Send[GenerateForegroundInstanceMaskRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenerateForegroundInstanceMaskRequest) Init() GenerateForegroundInstanceMaskRequest {
	rv := objc.Send[GenerateForegroundInstanceMaskRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenerateForegroundInstanceMaskRequest) Autorelease() GenerateForegroundInstanceMaskRequest {
	rv := objc.Send[GenerateForegroundInstanceMaskRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenerateForegroundInstanceMaskRequest creates a new GenerateForegroundInstanceMaskRequest instance.
func NewGenerateForegroundInstanceMaskRequest() GenerateForegroundInstanceMaskRequest {
	return getGenerateForegroundInstanceMaskRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GenerateForegroundInstanceMaskRequest */
// A request that generates an instance mask of noticable objects to separate from the background.


// A request that generates an instance mask of noticable objects to separate from the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateForegroundInstanceMaskRequest
type GenerateForegroundInstanceMaskRequest struct {
	ImageBasedRequest
}

// GenerateForegroundInstanceMaskRequestFrom constructs a [GenerateForegroundInstanceMaskRequest] from an unsafe.Pointer.
//
// A request that generates an instance mask of noticable objects to separate from the background.
func GenerateForegroundInstanceMaskRequestFrom(ptr unsafe.Pointer) GenerateForegroundInstanceMaskRequest {
	return GenerateForegroundInstanceMaskRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GenerateForegroundInstanceMaskRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GenerateForegroundInstanceMaskRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GenerateForegroundInstanceMaskRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GenerateForegroundInstanceMaskRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GenerateForegroundInstanceMaskRequest */

// The instance masks the request observes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateForegroundInstanceMaskRequest/results
func (g_ GenerateForegroundInstanceMaskRequest) Results() []InstanceMaskObservation {
	rv := objc.Send[[]InstanceMaskObservation](g_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying the first revision of the foreground instance mask request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateforegroundinstancemaskrequestrevision1
func (g_ GenerateForegroundInstanceMaskRequest) VNGenerateForegroundInstanceMaskRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateForegroundInstanceMaskRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNGenerateForegroundInstanceMaskRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNGenerateForegroundInstanceMaskRequest */



