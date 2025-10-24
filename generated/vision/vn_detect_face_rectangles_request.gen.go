// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNDetectFaceRectanglesRequest */


/* debug [class_header]: Header for VNDetectFaceRectanglesRequest */
// The class instance for the [DetectFaceRectanglesRequest] class.
var (
	DetectFaceRectanglesRequestClass     _DetectFaceRectanglesRequestClass
	DetectFaceRectanglesRequestClassOnce sync.Once
)

func getDetectFaceRectanglesRequestClass() _DetectFaceRectanglesRequestClass {
	DetectFaceRectanglesRequestClassOnce.Do(func() {
		DetectFaceRectanglesRequestClass = _DetectFaceRectanglesRequestClass{objc.GetClass("VNDetectFaceRectanglesRequest")}
	})
	return DetectFaceRectanglesRequestClass
}

type _DetectFaceRectanglesRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectFaceRectanglesRequest */
// An interface definition for the [DetectFaceRectanglesRequest] class.
type IDetectFaceRectanglesRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for DetectFaceRectanglesRequest */
	// properties:
	Results() []FaceObservation
	VNDetectFaceRectanglesRequestRevision1() int
	VNDetectFaceRectanglesRequestRevision2() int
	VNDetectFaceRectanglesRequestRevision3() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectFaceRectanglesRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectFaceRectanglesRequest */
// Alloc allocates a new instance without initialization.
func (dc _DetectFaceRectanglesRequestClass) Alloc() DetectFaceRectanglesRequest {
	rv := objc.Send[DetectFaceRectanglesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectFaceRectanglesRequestClass) New() DetectFaceRectanglesRequest {
	rv := objc.Send[DetectFaceRectanglesRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectFaceRectanglesRequest) Init() DetectFaceRectanglesRequest {
	rv := objc.Send[DetectFaceRectanglesRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectFaceRectanglesRequest) Autorelease() DetectFaceRectanglesRequest {
	rv := objc.Send[DetectFaceRectanglesRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectFaceRectanglesRequest creates a new DetectFaceRectanglesRequest instance.
func NewDetectFaceRectanglesRequest() DetectFaceRectanglesRequest {
	return getDetectFaceRectanglesRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectFaceRectanglesRequest */
// A request that finds faces within an image.
//
// This request returns faces as rectangular bounding boxes with origin and size.


// A request that finds faces within an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectFaceRectanglesRequest
type DetectFaceRectanglesRequest struct {
	ImageBasedRequest
}

// DetectFaceRectanglesRequestFrom constructs a [DetectFaceRectanglesRequest] from an unsafe.Pointer.
//
// A request that finds faces within an image.
func DetectFaceRectanglesRequestFrom(ptr unsafe.Pointer) DetectFaceRectanglesRequest {
	return DetectFaceRectanglesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectFaceRectanglesRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectFaceRectanglesRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectFaceRectanglesRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectFaceRectanglesRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectFaceRectanglesRequest */

// The results of the face detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectFaceRectanglesRequest/results
func (d_ DetectFaceRectanglesRequest) Results() []FaceObservation {
	rv := objc.Send[[]FaceObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the face rectangles detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacerectanglesrequestrevision1
func (d_ DetectFaceRectanglesRequest) VNDetectFaceRectanglesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceRectanglesRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNDetectFaceRectanglesRequestRevision1 */


// A constant for specifying revision 2 of the face rectangles detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacerectanglesrequestrevision2
func (d_ DetectFaceRectanglesRequest) VNDetectFaceRectanglesRequestRevision2() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceRectanglesRequestRevision2"))
	return rv
}/* debug [instance_properties/getter]: VNDetectFaceRectanglesRequestRevision2 */


// A constant for specifying revision 3 of the face rectangles detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacerectanglesrequestrevision3
func (d_ DetectFaceRectanglesRequest) VNDetectFaceRectanglesRequestRevision3() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceRectanglesRequestRevision3"))
	return rv
}/* debug [instance_properties/getter]: VNDetectFaceRectanglesRequestRevision3 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectFaceRectanglesRequest */



