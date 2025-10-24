// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ClassifyImageRequest] class.
var (
	ClassifyImageRequestClass     _ClassifyImageRequestClass
	ClassifyImageRequestClassOnce sync.Once
)

func getClassifyImageRequestClass() _ClassifyImageRequestClass {
	ClassifyImageRequestClassOnce.Do(func() {
		ClassifyImageRequestClass = _ClassifyImageRequestClass{objc.GetClass("VNClassifyImageRequest")}
	})
	return ClassifyImageRequestClass
}

type _ClassifyImageRequestClass struct {
	class objc.Class
}

// An interface definition for the [ClassifyImageRequest] class.
type IClassifyImageRequest interface {
	IImageBasedRequest
	// properties:
	Results() IVNClassificationObservation
	SetResults(value IVNClassificationObservation)
	VNClassifyImageRequestRevision1() int
	// methods:
}

// A request to classify an image.
//
// This type of request produces a collection of objects that describe an image. Access the classifications through .


// A request to classify an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNClassifyImageRequest
type ClassifyImageRequest struct {
	ImageBasedRequest
}

// ClassifyImageRequestFrom constructs a [ClassifyImageRequest] from an unsafe.Pointer.
//
// A request to classify an image.
func ClassifyImageRequestFrom(ptr unsafe.Pointer) ClassifyImageRequest {
	return ClassifyImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ClassifyImageRequestClass) Alloc() ClassifyImageRequest {
	rv := objc.Send[ClassifyImageRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ClassifyImageRequestClass) New() ClassifyImageRequest {
	rv := objc.Send[ClassifyImageRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ClassifyImageRequest) Init() ClassifyImageRequest {
	rv := objc.Send[ClassifyImageRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ClassifyImageRequest) Autorelease() ClassifyImageRequest {
	rv := objc.Send[ClassifyImageRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewClassifyImageRequest creates a new ClassifyImageRequest instance.
func NewClassifyImageRequest() ClassifyImageRequest {
	return getClassifyImageRequestClass().New()
}



// The results of the image classification request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassifyimagerequest/results
func (c_ ClassifyImageRequest) Results() IVNClassificationObservation {
	rv := objc.Send[ClassificationObservation](c_.ID, objc.Sel("results"))
	return rv
}


// The results of the image classification request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassifyimagerequest/results
func (c_ ClassifyImageRequest) SetResults(value IVNClassificationObservation) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResults:"), value)
}


// A constant for specifying the first revision of the image-classification request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassifyimagerequestrevision1
func (c_ ClassifyImageRequest) VNClassifyImageRequestRevision1() int {
	rv := objc.Send[int](c_.ID, objc.Sel("VNClassifyImageRequestRevision1"))
	return rv
}



