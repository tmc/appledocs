// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNDetectHorizonRequest */


/* debug [class_header]: Header for VNDetectHorizonRequest */
// The class instance for the [DetectHorizonRequest] class.
var (
	DetectHorizonRequestClass     _DetectHorizonRequestClass
	DetectHorizonRequestClassOnce sync.Once
)

func getDetectHorizonRequestClass() _DetectHorizonRequestClass {
	DetectHorizonRequestClassOnce.Do(func() {
		DetectHorizonRequestClass = _DetectHorizonRequestClass{objc.GetClass("VNDetectHorizonRequest")}
	})
	return DetectHorizonRequestClass
}

type _DetectHorizonRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectHorizonRequest */
// An interface definition for the [DetectHorizonRequest] class.
type IDetectHorizonRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for DetectHorizonRequest */
	// properties:
	Results() []HorizonObservation
	VNDetectHorizonRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectHorizonRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectHorizonRequest */
// Alloc allocates a new instance without initialization.
func (dc _DetectHorizonRequestClass) Alloc() DetectHorizonRequest {
	rv := objc.Send[DetectHorizonRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectHorizonRequestClass) New() DetectHorizonRequest {
	rv := objc.Send[DetectHorizonRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectHorizonRequest) Init() DetectHorizonRequest {
	rv := objc.Send[DetectHorizonRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectHorizonRequest) Autorelease() DetectHorizonRequest {
	rv := objc.Send[DetectHorizonRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectHorizonRequest creates a new DetectHorizonRequest instance.
func NewDetectHorizonRequest() DetectHorizonRequest {
	return getDetectHorizonRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectHorizonRequest */
// An image-analysis request that determines the horizon angle in an image.


// An image-analysis request that determines the horizon angle in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHorizonRequest
type DetectHorizonRequest struct {
	ImageBasedRequest
}

// DetectHorizonRequestFrom constructs a [DetectHorizonRequest] from an unsafe.Pointer.
//
// An image-analysis request that determines the horizon angle in an image.
func DetectHorizonRequestFrom(ptr unsafe.Pointer) DetectHorizonRequest {
	return DetectHorizonRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectHorizonRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectHorizonRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectHorizonRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectHorizonRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectHorizonRequest */

// The results of the horizon detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHorizonRequest/results
func (d_ DetectHorizonRequest) Results() []HorizonObservation {
	rv := objc.Send[[]HorizonObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the horizon detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthorizonrequestrevision1
func (d_ DetectHorizonRequest) VNDetectHorizonRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHorizonRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNDetectHorizonRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectHorizonRequest */



