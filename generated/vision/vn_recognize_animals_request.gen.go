// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RecognizeAnimalsRequest] class.
var (
	RecognizeAnimalsRequestClass     _RecognizeAnimalsRequestClass
	RecognizeAnimalsRequestClassOnce sync.Once
)

func getRecognizeAnimalsRequestClass() _RecognizeAnimalsRequestClass {
	RecognizeAnimalsRequestClassOnce.Do(func() {
		RecognizeAnimalsRequestClass = _RecognizeAnimalsRequestClass{objc.GetClass("VNRecognizeAnimalsRequest")}
	})
	return RecognizeAnimalsRequestClass
}

type _RecognizeAnimalsRequestClass struct {
	class objc.Class
}

// An interface definition for the [RecognizeAnimalsRequest] class.
type IRecognizeAnimalsRequest interface {
	IImageBasedRequest
}

// A request that recognizes animals in an image.
//
// Use the method to determine which animals the request supports.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRecognizeAnimalsRequest
type RecognizeAnimalsRequest struct {
	ImageBasedRequest
}

// RecognizeAnimalsRequestFrom constructs a [RecognizeAnimalsRequest] from an unsafe.Pointer.
//
// A request that recognizes animals in an image.
func RecognizeAnimalsRequestFrom(ptr unsafe.Pointer) RecognizeAnimalsRequest {
	return RecognizeAnimalsRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RecognizeAnimalsRequestClass) Alloc() RecognizeAnimalsRequest {
	rv := objc.Send[RecognizeAnimalsRequest](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RecognizeAnimalsRequestClass) New() RecognizeAnimalsRequest {
	rv := objc.Send[RecognizeAnimalsRequest](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecognizeAnimalsRequest) Init() RecognizeAnimalsRequest {
	rv := objc.Send[RecognizeAnimalsRequest](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecognizeAnimalsRequest) Autorelease() RecognizeAnimalsRequest {
	rv := objc.Send[RecognizeAnimalsRequest](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecognizeAnimalsRequest creates a new RecognizeAnimalsRequest instance.
func NewRecognizeAnimalsRequest() RecognizeAnimalsRequest {
	return getRecognizeAnimalsRequestClass().New()
}


// A constant for specifying revision 1 of the animal recognition request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequestrevision1
func (r_ RecognizeAnimalsRequest) VNRecognizeAnimalsRequestRevision1() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeAnimalsRequestRevision1"))
	return rv
}

// A constant for specifying revision 2 of the animal recognition request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequestrevision2
func (r_ RecognizeAnimalsRequest) VNRecognizeAnimalsRequestRevision2() int {
	rv := objc.Send[int](r_.ID, objc.Sel("VNRecognizeAnimalsRequestRevision2"))
	return rv
}

// The results of the request to recognize animals.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequest/results
func (r_ RecognizeAnimalsRequest) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of the request to recognize animals.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrecognizeanimalsrequest/results
func (r_ RecognizeAnimalsRequest) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setResults:"), value)
}



