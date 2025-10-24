// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNGenerateObjectnessBasedSaliencyImageRequest */


/* debug [class_header]: Header for VNGenerateObjectnessBasedSaliencyImageRequest */
// The class instance for the [GenerateObjectnessBasedSaliencyImageRequest] class.
var (
	GenerateObjectnessBasedSaliencyImageRequestClass     _GenerateObjectnessBasedSaliencyImageRequestClass
	GenerateObjectnessBasedSaliencyImageRequestClassOnce sync.Once
)

func getGenerateObjectnessBasedSaliencyImageRequestClass() _GenerateObjectnessBasedSaliencyImageRequestClass {
	GenerateObjectnessBasedSaliencyImageRequestClassOnce.Do(func() {
		GenerateObjectnessBasedSaliencyImageRequestClass = _GenerateObjectnessBasedSaliencyImageRequestClass{objc.GetClass("VNGenerateObjectnessBasedSaliencyImageRequest")}
	})
	return GenerateObjectnessBasedSaliencyImageRequestClass
}

type _GenerateObjectnessBasedSaliencyImageRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GenerateObjectnessBasedSaliencyImageRequest */
// An interface definition for the [GenerateObjectnessBasedSaliencyImageRequest] class.
type IGenerateObjectnessBasedSaliencyImageRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for GenerateObjectnessBasedSaliencyImageRequest */
	// properties:
	Results() []SaliencyImageObservation
	VNGenerateObjectnessBasedSaliencyImageRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GenerateObjectnessBasedSaliencyImageRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GenerateObjectnessBasedSaliencyImageRequest */
// Alloc allocates a new instance without initialization.
func (gc _GenerateObjectnessBasedSaliencyImageRequestClass) Alloc() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[GenerateObjectnessBasedSaliencyImageRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GenerateObjectnessBasedSaliencyImageRequestClass) New() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[GenerateObjectnessBasedSaliencyImageRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenerateObjectnessBasedSaliencyImageRequest) Init() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[GenerateObjectnessBasedSaliencyImageRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenerateObjectnessBasedSaliencyImageRequest) Autorelease() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[GenerateObjectnessBasedSaliencyImageRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenerateObjectnessBasedSaliencyImageRequest creates a new GenerateObjectnessBasedSaliencyImageRequest instance.
func NewGenerateObjectnessBasedSaliencyImageRequest() GenerateObjectnessBasedSaliencyImageRequest {
	return getGenerateObjectnessBasedSaliencyImageRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GenerateObjectnessBasedSaliencyImageRequest */
// A request that generates a heat map that identifies the parts of an image most likely to represent objects.
//
// The resulting observation, , encodes this data as a heat map, which you can use to highlight regions of interest.


// A request that generates a heat map that identifies the parts of an image most likely to represent objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateObjectnessBasedSaliencyImageRequest
type GenerateObjectnessBasedSaliencyImageRequest struct {
	ImageBasedRequest
}

// GenerateObjectnessBasedSaliencyImageRequestFrom constructs a [GenerateObjectnessBasedSaliencyImageRequest] from an unsafe.Pointer.
//
// A request that generates a heat map that identifies the parts of an image most likely to represent objects.
func GenerateObjectnessBasedSaliencyImageRequestFrom(ptr unsafe.Pointer) GenerateObjectnessBasedSaliencyImageRequest {
	return GenerateObjectnessBasedSaliencyImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GenerateObjectnessBasedSaliencyImageRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GenerateObjectnessBasedSaliencyImageRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GenerateObjectnessBasedSaliencyImageRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GenerateObjectnessBasedSaliencyImageRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GenerateObjectnessBasedSaliencyImageRequest */

// The results of the image saliency request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateObjectnessBasedSaliencyImageRequest/results
func (g_ GenerateObjectnessBasedSaliencyImageRequest) Results() []SaliencyImageObservation {
	rv := objc.Send[[]SaliencyImageObservation](g_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the image saliency request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateobjectnessbasedsaliencyimagerequestrevision1
func (g_ GenerateObjectnessBasedSaliencyImageRequest) VNGenerateObjectnessBasedSaliencyImageRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateObjectnessBasedSaliencyImageRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNGenerateObjectnessBasedSaliencyImageRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNGenerateObjectnessBasedSaliencyImageRequest */



