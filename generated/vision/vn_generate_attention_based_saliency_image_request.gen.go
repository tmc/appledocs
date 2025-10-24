// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNGenerateAttentionBasedSaliencyImageRequest */


/* debug [class_header]: Header for VNGenerateAttentionBasedSaliencyImageRequest */
// The class instance for the [GenerateAttentionBasedSaliencyImageRequest] class.
var (
	GenerateAttentionBasedSaliencyImageRequestClass     _GenerateAttentionBasedSaliencyImageRequestClass
	GenerateAttentionBasedSaliencyImageRequestClassOnce sync.Once
)

func getGenerateAttentionBasedSaliencyImageRequestClass() _GenerateAttentionBasedSaliencyImageRequestClass {
	GenerateAttentionBasedSaliencyImageRequestClassOnce.Do(func() {
		GenerateAttentionBasedSaliencyImageRequestClass = _GenerateAttentionBasedSaliencyImageRequestClass{objc.GetClass("VNGenerateAttentionBasedSaliencyImageRequest")}
	})
	return GenerateAttentionBasedSaliencyImageRequestClass
}

type _GenerateAttentionBasedSaliencyImageRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GenerateAttentionBasedSaliencyImageRequest */
// An interface definition for the [GenerateAttentionBasedSaliencyImageRequest] class.
type IGenerateAttentionBasedSaliencyImageRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for GenerateAttentionBasedSaliencyImageRequest */
	// properties:
	Results() []SaliencyImageObservation
	VNGenerateAttentionBasedSaliencyImageRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GenerateAttentionBasedSaliencyImageRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GenerateAttentionBasedSaliencyImageRequest */
// Alloc allocates a new instance without initialization.
func (gc _GenerateAttentionBasedSaliencyImageRequestClass) Alloc() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[GenerateAttentionBasedSaliencyImageRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GenerateAttentionBasedSaliencyImageRequestClass) New() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[GenerateAttentionBasedSaliencyImageRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenerateAttentionBasedSaliencyImageRequest) Init() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[GenerateAttentionBasedSaliencyImageRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenerateAttentionBasedSaliencyImageRequest) Autorelease() GenerateAttentionBasedSaliencyImageRequest {
	rv := objc.Send[GenerateAttentionBasedSaliencyImageRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenerateAttentionBasedSaliencyImageRequest creates a new GenerateAttentionBasedSaliencyImageRequest instance.
func NewGenerateAttentionBasedSaliencyImageRequest() GenerateAttentionBasedSaliencyImageRequest {
	return getGenerateAttentionBasedSaliencyImageRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GenerateAttentionBasedSaliencyImageRequest */
// An object that produces a heat map that identifies the parts of an image most likely to draw attention.


// An object that produces a heat map that identifies the parts of an image most likely to draw attention.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateAttentionBasedSaliencyImageRequest
type GenerateAttentionBasedSaliencyImageRequest struct {
	ImageBasedRequest
}

// GenerateAttentionBasedSaliencyImageRequestFrom constructs a [GenerateAttentionBasedSaliencyImageRequest] from an unsafe.Pointer.
//
// An object that produces a heat map that identifies the parts of an image most likely to draw attention.
func GenerateAttentionBasedSaliencyImageRequestFrom(ptr unsafe.Pointer) GenerateAttentionBasedSaliencyImageRequest {
	return GenerateAttentionBasedSaliencyImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GenerateAttentionBasedSaliencyImageRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GenerateAttentionBasedSaliencyImageRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GenerateAttentionBasedSaliencyImageRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GenerateAttentionBasedSaliencyImageRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GenerateAttentionBasedSaliencyImageRequest */

// The results of the image saliency request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateAttentionBasedSaliencyImageRequest/results
func (g_ GenerateAttentionBasedSaliencyImageRequest) Results() []SaliencyImageObservation {
	rv := objc.Send[[]SaliencyImageObservation](g_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the image saliency request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateattentionbasedsaliencyimagerequestrevision1
func (g_ GenerateAttentionBasedSaliencyImageRequest) VNGenerateAttentionBasedSaliencyImageRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateAttentionBasedSaliencyImageRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNGenerateAttentionBasedSaliencyImageRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNGenerateAttentionBasedSaliencyImageRequest */



