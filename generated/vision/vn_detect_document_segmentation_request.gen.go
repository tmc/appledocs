// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [DetectDocumentSegmentationRequest] class.
type IDetectDocumentSegmentationRequest interface {
	IImageBasedRequest
}

// An object that detects rectangular regions that contain text in the input image.
//
// Perform this request to detect a document in an image. The result that the request generates contains the four corner points of a document’s quadrilateral and saliency mask.
//
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

// Alloc allocates a new instance without initialization.
func (dc _DetectDocumentSegmentationRequestClass) Alloc() DetectDocumentSegmentationRequest {
	rv := objc.Send[DetectDocumentSegmentationRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The results of a document segmentation request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequest/results
func (d_ DetectDocumentSegmentationRequest) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of a document segmentation request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequest/results
func (d_ DetectDocumentSegmentationRequest) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setResults:"), value)
}

// A constant for specifying revision 1 of the document segmentation request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequestrevision1
func (d_ DetectDocumentSegmentationRequest) VNDetectDocumentSegmentationRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectDocumentSegmentationRequestRevision1"))
	return rv
}



