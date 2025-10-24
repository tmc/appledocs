// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNDetectDocumentSegmentationRequest */


/* debug [class_header]: Header for VNDetectDocumentSegmentationRequest */
// The class instance for the [DetectDocumentSegmentationRequest] class.
var (
	DetectDocumentSegmentationRequestClass     _DetectDocumentSegmentationRequestClass
	DetectDocumentSegmentationRequestClassOnce sync.Once
)

func getDetectDocumentSegmentationRequestClass() _DetectDocumentSegmentationRequestClass {
	DetectDocumentSegmentationRequestClassOnce.Do(func() {
		DetectDocumentSegmentationRequestClass = _DetectDocumentSegmentationRequestClass{objc.GetClass("VNDetectDocumentSegmentationRequest")}
	})
	return DetectDocumentSegmentationRequestClass
}

type _DetectDocumentSegmentationRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectDocumentSegmentationRequest */
// An interface definition for the [DetectDocumentSegmentationRequest] class.
type IDetectDocumentSegmentationRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for DetectDocumentSegmentationRequest */
	// properties:
	Results() []RectangleObservation
	VNDetectDocumentSegmentationRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectDocumentSegmentationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectDocumentSegmentationRequest */
// Alloc allocates a new instance without initialization.
func (dc _DetectDocumentSegmentationRequestClass) Alloc() DetectDocumentSegmentationRequest {
	rv := objc.Send[DetectDocumentSegmentationRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectDocumentSegmentationRequestClass) New() DetectDocumentSegmentationRequest {
	rv := objc.Send[DetectDocumentSegmentationRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectDocumentSegmentationRequest) Init() DetectDocumentSegmentationRequest {
	rv := objc.Send[DetectDocumentSegmentationRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectDocumentSegmentationRequest) Autorelease() DetectDocumentSegmentationRequest {
	rv := objc.Send[DetectDocumentSegmentationRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectDocumentSegmentationRequest creates a new DetectDocumentSegmentationRequest instance.
func NewDetectDocumentSegmentationRequest() DetectDocumentSegmentationRequest {
	return getDetectDocumentSegmentationRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectDocumentSegmentationRequest */
// An object that detects rectangular regions that contain text in the input image.
//
// Perform this request to detect a document in an image. The result that the request generates contains the four corner points of a document’s quadrilateral and saliency mask.


// An object that detects rectangular regions that contain text in the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectDocumentSegmentationRequest
type DetectDocumentSegmentationRequest struct {
	ImageBasedRequest
}

// DetectDocumentSegmentationRequestFrom constructs a [DetectDocumentSegmentationRequest] from an unsafe.Pointer.
//
// An object that detects rectangular regions that contain text in the input image.
func DetectDocumentSegmentationRequestFrom(ptr unsafe.Pointer) DetectDocumentSegmentationRequest {
	return DetectDocumentSegmentationRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectDocumentSegmentationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectDocumentSegmentationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectDocumentSegmentationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectDocumentSegmentationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectDocumentSegmentationRequest */

// The results of a document segmentation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectDocumentSegmentationRequest/results
func (d_ DetectDocumentSegmentationRequest) Results() []RectangleObservation {
	rv := objc.Send[[]RectangleObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the document segmentation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequestrevision1
func (d_ DetectDocumentSegmentationRequest) VNDetectDocumentSegmentationRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectDocumentSegmentationRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNDetectDocumentSegmentationRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectDocumentSegmentationRequest */



