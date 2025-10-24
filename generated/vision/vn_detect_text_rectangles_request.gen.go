// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNDetectTextRectanglesRequest */


/* debug [class_header]: Header for VNDetectTextRectanglesRequest */
// The class instance for the [DetectTextRectanglesRequest] class.
var (
	DetectTextRectanglesRequestClass     _DetectTextRectanglesRequestClass
	DetectTextRectanglesRequestClassOnce sync.Once
)

func getDetectTextRectanglesRequestClass() _DetectTextRectanglesRequestClass {
	DetectTextRectanglesRequestClassOnce.Do(func() {
		DetectTextRectanglesRequestClass = _DetectTextRectanglesRequestClass{objc.GetClass("VNDetectTextRectanglesRequest")}
	})
	return DetectTextRectanglesRequestClass
}

type _DetectTextRectanglesRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectTextRectanglesRequest */
// An interface definition for the [DetectTextRectanglesRequest] class.
type IDetectTextRectanglesRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for DetectTextRectanglesRequest */
	// properties:
	ReportCharacterBoxes() bool
	SetReportCharacterBoxes(value bool)
	Results() []TextObservation
	VNDetectTextRectanglesRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectTextRectanglesRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectTextRectanglesRequest */
// Alloc allocates a new instance without initialization.
func (dc _DetectTextRectanglesRequestClass) Alloc() DetectTextRectanglesRequest {
	rv := objc.Send[DetectTextRectanglesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectTextRectanglesRequestClass) New() DetectTextRectanglesRequest {
	rv := objc.Send[DetectTextRectanglesRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectTextRectanglesRequest) Init() DetectTextRectanglesRequest {
	rv := objc.Send[DetectTextRectanglesRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectTextRectanglesRequest) Autorelease() DetectTextRectanglesRequest {
	rv := objc.Send[DetectTextRectanglesRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectTextRectanglesRequest creates a new DetectTextRectanglesRequest instance.
func NewDetectTextRectanglesRequest() DetectTextRectanglesRequest {
	return getDetectTextRectanglesRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectTextRectanglesRequest */
// An image-analysis request that finds regions of visible text in an image.
//
// This request returns detected text characters as rectangular bounding boxes with origin and size.


// An image-analysis request that finds regions of visible text in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTextRectanglesRequest
type DetectTextRectanglesRequest struct {
	ImageBasedRequest
}

// DetectTextRectanglesRequestFrom constructs a [DetectTextRectanglesRequest] from an unsafe.Pointer.
//
// An image-analysis request that finds regions of visible text in an image.
func DetectTextRectanglesRequestFrom(ptr unsafe.Pointer) DetectTextRectanglesRequest {
	return DetectTextRectanglesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectTextRectanglesRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectTextRectanglesRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectTextRectanglesRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectTextRectanglesRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectTextRectanglesRequest */

// A Boolean value that indicates whether the request detects character bounding boxes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTextRectanglesRequest/reportCharacterBoxes
func (d_ DetectTextRectanglesRequest) ReportCharacterBoxes() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("reportCharacterBoxes"))
	return rv
}/* debug [instance_properties/getter]: reportCharacterBoxes */


// A Boolean value that indicates whether the request detects character bounding boxes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTextRectanglesRequest/reportCharacterBoxes
func (d_ DetectTextRectanglesRequest) SetReportCharacterBoxes(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setReportCharacterBoxes:"), value)
}/* debug [instance_properties/setter]: reportCharacterBoxes */


// The results of the request to detect text rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTextRectanglesRequest/results
func (d_ DetectTextRectanglesRequest) Results() []TextObservation {
	rv := objc.Send[[]TextObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the text rectangles detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttextrectanglesrequestrevision1
func (d_ DetectTextRectanglesRequest) VNDetectTextRectanglesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectTextRectanglesRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNDetectTextRectanglesRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectTextRectanglesRequest */



