// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [DetectFaceLandmarksRequest] class.
type IDetectFaceLandmarksRequest interface {
	IImageBasedRequest
	// properties:
	Constellation() RequestFaceLandmarksConstellation /* not a class type */
	SetConstellation(value RequestFaceLandmarksConstellation /* not a class type */)
	Results() objc.IObject /* cross-framework: FaceObservation */
	SetResults(value objc.IObject /* cross-framework: FaceObservation */)
	VNDetectFaceLandmarksRequestRevision1() int
	VNDetectFaceLandmarksRequestRevision2() int
	VNDetectFaceLandmarksRequestRevision3() int
	InputFaceObservations() objc.IObject /* cross-framework: FaceObservation */
	SetInputFaceObservations(value objc.IObject /* cross-framework: FaceObservation */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (dc _DetectFaceLandmarksRequestClass) Alloc() DetectFaceLandmarksRequest {
	rv := objc.Send[DetectFaceLandmarksRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A variable that describes how a face landmarks request orders or enumerates the resulting features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequest/constellation
func (d_ DetectFaceLandmarksRequest) Constellation() RequestFaceLandmarksConstellation /* not a class type */ {
	rv := objc.Send[RequestFaceLandmarksConstellation](d_.ID, objc.Sel("constellation"))
	return rv
}


// A variable that describes how a face landmarks request orders or enumerates the resulting features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequest/constellation
func (d_ DetectFaceLandmarksRequest) SetConstellation(value RequestFaceLandmarksConstellation /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setConstellation:"), value)
}


// The results of the face landmarks request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequest/results
func (d_ DetectFaceLandmarksRequest) Results() objc.IObject /* cross-framework: FaceObservation */ {
	rv := objc.Send[FaceObservation](d_.ID, objc.Sel("results"))
	return rv
}


// The results of the face landmarks request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequest/results
func (d_ DetectFaceLandmarksRequest) SetResults(value objc.IObject /* cross-framework: FaceObservation */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setResults:"), value)
}


// A constant for specifying revision 1 of the face landmarks detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequestrevision1
func (d_ DetectFaceLandmarksRequest) VNDetectFaceLandmarksRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceLandmarksRequestRevision1"))
	return rv
}


// A constant for specifying revision 2 of the face landmarks detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequestrevision2
func (d_ DetectFaceLandmarksRequest) VNDetectFaceLandmarksRequestRevision2() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceLandmarksRequestRevision2"))
	return rv
}


// A constant for specifying revision 3 of the face landmarks detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacelandmarksrequestrevision3
func (d_ DetectFaceLandmarksRequest) VNDetectFaceLandmarksRequestRevision3() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceLandmarksRequestRevision3"))
	return rv
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservationaccepting/inputfaceobservations
func (d_ DetectFaceLandmarksRequest) InputFaceObservations() objc.IObject /* cross-framework: FaceObservation */ {
	rv := objc.Send[FaceObservation](d_.ID, objc.Sel("inputFaceObservations"))
	return rv
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservationaccepting/inputfaceobservations
func (d_ DetectFaceLandmarksRequest) SetInputFaceObservations(value objc.IObject /* cross-framework: FaceObservation */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setInputFaceObservations:"), value)
}



