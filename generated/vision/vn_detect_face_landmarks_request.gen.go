// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNDetectFaceLandmarksRequest */


/* debug [class_header]: Header for VNDetectFaceLandmarksRequest */
// The class instance for the [DetectFaceLandmarksRequest] class.
var (
	DetectFaceLandmarksRequestClass     _DetectFaceLandmarksRequestClass
	DetectFaceLandmarksRequestClassOnce sync.Once
)

func getDetectFaceLandmarksRequestClass() _DetectFaceLandmarksRequestClass {
	DetectFaceLandmarksRequestClassOnce.Do(func() {
		DetectFaceLandmarksRequestClass = _DetectFaceLandmarksRequestClass{objc.GetClass("VNDetectFaceLandmarksRequest")}
	})
	return DetectFaceLandmarksRequestClass
}

type _DetectFaceLandmarksRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectFaceLandmarksRequest */
// An interface definition for the [DetectFaceLandmarksRequest] class.
type IDetectFaceLandmarksRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for DetectFaceLandmarksRequest */
	// properties:
	Constellation() RequestFaceLandmarksConstellation
	SetConstellation(value RequestFaceLandmarksConstellation)
	Results() []FaceObservation
	VNDetectFaceLandmarksRequestRevision1() int
	VNDetectFaceLandmarksRequestRevision2() int
	VNDetectFaceLandmarksRequestRevision3() int
	InputFaceObservations() IVNFaceObservation
	SetInputFaceObservations(value IVNFaceObservation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectFaceLandmarksRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectFaceLandmarksRequest */
// Alloc allocates a new instance without initialization.
func (dc _DetectFaceLandmarksRequestClass) Alloc() DetectFaceLandmarksRequest {
	rv := objc.Send[DetectFaceLandmarksRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectFaceLandmarksRequestClass) New() DetectFaceLandmarksRequest {
	rv := objc.Send[DetectFaceLandmarksRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectFaceLandmarksRequest) Init() DetectFaceLandmarksRequest {
	rv := objc.Send[DetectFaceLandmarksRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectFaceLandmarksRequest) Autorelease() DetectFaceLandmarksRequest {
	rv := objc.Send[DetectFaceLandmarksRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectFaceLandmarksRequest creates a new DetectFaceLandmarksRequest instance.
func NewDetectFaceLandmarksRequest() DetectFaceLandmarksRequest {
	return getDetectFaceLandmarksRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectFaceLandmarksRequest */
// An image-analysis request that finds facial features like eyes and mouth in an image.
//
// By default, a face landmarks request first locates all faces in the input image, then analyzes each to detect facial features. If you’ve already located all the faces in an image, or want to detect landmarks in only a subset of the faces in the image, set the property to an array of objects representing the faces you want to analyze. You can either use face observations output by a or manually create instances with the bounding boxes of the faces you want to analyze.


// An image-analysis request that finds facial features like eyes and mouth in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectFaceLandmarksRequest
type DetectFaceLandmarksRequest struct {
	ImageBasedRequest
}

// DetectFaceLandmarksRequestFrom constructs a [DetectFaceLandmarksRequest] from an unsafe.Pointer.
//
// An image-analysis request that finds facial features like eyes and mouth in an image.
func DetectFaceLandmarksRequestFrom(ptr unsafe.Pointer) DetectFaceLandmarksRequest {
	return DetectFaceLandmarksRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectFaceLandmarksRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectFaceLandmarksRequest */

// Returns a Boolean value that indicates whether a revision supports a constellation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectFaceLandmarksRequest/revision(_:supportsConstellation:)
func (dc _DetectFaceLandmarksRequestClass) RevisionSupportsConstellation(requestRevision uint, constellation RequestFaceLandmarksConstellation) bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("revision:supportsConstellation:"), requestRevision, constellation)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RevisionSupportsConstellation) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectFaceLandmarksRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectFaceLandmarksRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectFaceLandmarksRequest */

// A variable that describes how a face landmarks request orders or enumerates the resulting features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectFaceLandmarksRequest/constellation
func (d_ DetectFaceLandmarksRequest) Constellation() RequestFaceLandmarksConstellation {
	rv := objc.Send[RequestFaceLandmarksConstellation](d_.ID, objc.Sel("constellation"))
	return rv
}/* debug [instance_properties/getter]: constellation */


// A variable that describes how a face landmarks request orders or enumerates the resulting features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectFaceLandmarksRequest/constellation
func (d_ DetectFaceLandmarksRequest) SetConstellation(value RequestFaceLandmarksConstellation) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setConstellation:"), value)
}/* debug [instance_properties/setter]: constellation */


// The results of the face landmarks request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectFaceLandmarksRequest/results
func (d_ DetectFaceLandmarksRequest) Results() []FaceObservation {
	rv := objc.Send[[]FaceObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the face landmarks detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequestrevision1
func (d_ DetectFaceLandmarksRequest) VNDetectFaceLandmarksRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceLandmarksRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNDetectFaceLandmarksRequestRevision1 */


// A constant for specifying revision 2 of the face landmarks detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequestrevision2
func (d_ DetectFaceLandmarksRequest) VNDetectFaceLandmarksRequestRevision2() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceLandmarksRequestRevision2"))
	return rv
}/* debug [instance_properties/getter]: VNDetectFaceLandmarksRequestRevision2 */


// A constant for specifying revision 3 of the face landmarks detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequestrevision3
func (d_ DetectFaceLandmarksRequest) VNDetectFaceLandmarksRequestRevision3() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceLandmarksRequestRevision3"))
	return rv
}/* debug [instance_properties/getter]: VNDetectFaceLandmarksRequestRevision3 */


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservationaccepting/inputfaceobservations
func (d_ DetectFaceLandmarksRequest) InputFaceObservations() IVNFaceObservation {
	rv := objc.Send[FaceObservation](d_.ID, objc.Sel("inputFaceObservations"))
	return rv
}/* debug [instance_properties/getter]: inputFaceObservations */


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservationaccepting/inputfaceobservations
func (d_ DetectFaceLandmarksRequest) SetInputFaceObservations(value IVNFaceObservation) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setInputFaceObservations:"), value)
}/* debug [instance_properties/setter]: inputFaceObservations */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectFaceLandmarksRequest */



