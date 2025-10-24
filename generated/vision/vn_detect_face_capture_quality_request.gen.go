// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [DetectFaceCaptureQualityRequest] class.
var (
	DetectFaceCaptureQualityRequestClass     _DetectFaceCaptureQualityRequestClass
	DetectFaceCaptureQualityRequestClassOnce sync.Once
)

func getDetectFaceCaptureQualityRequestClass() _DetectFaceCaptureQualityRequestClass {
	DetectFaceCaptureQualityRequestClassOnce.Do(func() {
		DetectFaceCaptureQualityRequestClass = _DetectFaceCaptureQualityRequestClass{objc.GetClass("VNDetectFaceCaptureQualityRequest")}
	})
	return DetectFaceCaptureQualityRequestClass
}

type _DetectFaceCaptureQualityRequestClass struct {
	class objc.Class
}





// An interface definition for the [DetectFaceCaptureQualityRequest] class.
type IDetectFaceCaptureQualityRequest interface {
	IImageBasedRequest
	

	// properties:
	Results() []FaceObservation
	VNDetectFaceCaptureQualityRequestRevision1() int
	VNDetectFaceCaptureQualityRequestRevision2() int
	FaceCaptureQuality() float32
	SetFaceCaptureQuality(value float32)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DetectFaceCaptureQualityRequestClass) Alloc() DetectFaceCaptureQualityRequest {
	rv := objc.Send[DetectFaceCaptureQualityRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectFaceCaptureQualityRequestClass) New() DetectFaceCaptureQualityRequest {
	rv := objc.Send[DetectFaceCaptureQualityRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectFaceCaptureQualityRequest) Init() DetectFaceCaptureQualityRequest {
	rv := objc.Send[DetectFaceCaptureQualityRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectFaceCaptureQualityRequest) Autorelease() DetectFaceCaptureQualityRequest {
	rv := objc.Send[DetectFaceCaptureQualityRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectFaceCaptureQualityRequest creates a new DetectFaceCaptureQualityRequest instance.
func NewDetectFaceCaptureQualityRequest() DetectFaceCaptureQualityRequest {
	return getDetectFaceCaptureQualityRequestClass().New()
}





// A request that produces a floating-point number that represents the capture quality of a face in a photo.
//
// This request produces or updates a object’s property with a floating-point value. The value ranges from to . Faces with quality closer to are better lit, sharper, and more centrally positioned than faces with quality closer to . If you don’t execute the request, or the request fails, the property is .


// A request that produces a floating-point number that represents the capture quality of a face in a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectFaceCaptureQualityRequest
type DetectFaceCaptureQualityRequest struct {
	ImageBasedRequest
}

// DetectFaceCaptureQualityRequestFrom constructs a [DetectFaceCaptureQualityRequest] from an unsafe.Pointer.
//
// A request that produces a floating-point number that represents the capture quality of a face in a photo.
func DetectFaceCaptureQualityRequestFrom(ptr unsafe.Pointer) DetectFaceCaptureQualityRequest {
	return DetectFaceCaptureQualityRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

























// The results of the face-capture quality request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectFaceCaptureQualityRequest/results
func (d_ DetectFaceCaptureQualityRequest) Results() []FaceObservation {
	rv := objc.Send[[]FaceObservation](d_.ID, objc.Sel("results"))
	return rv
}


// A constant for specifying revision 1 of the face capture detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacecapturequalityrequestrevision1
func (d_ DetectFaceCaptureQualityRequest) VNDetectFaceCaptureQualityRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceCaptureQualityRequestRevision1"))
	return rv
}


// Revision 2 of the request algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectfacecapturequalityrequestrevision2
func (d_ DetectFaceCaptureQualityRequest) VNDetectFaceCaptureQualityRequestRevision2() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectFaceCaptureQualityRequestRevision2"))
	return rv
}


// A value that indicates the quality of the face capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/facecapturequality-bjg5
func (d_ DetectFaceCaptureQualityRequest) FaceCaptureQuality() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("faceCaptureQuality"))
	return rv
}


// A value that indicates the quality of the face capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnfaceobservation/facecapturequality-bjg5
func (d_ DetectFaceCaptureQualityRequest) SetFaceCaptureQuality(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFaceCaptureQuality:"), value)
}








