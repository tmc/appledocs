// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DetectHorizonRequest] class.
var (
	DetectHorizonRequestClass     _DetectHorizonRequestClass
	DetectHorizonRequestClassOnce sync.Once
)

func getDetectHorizonRequestClass() _DetectHorizonRequestClass {
	DetectHorizonRequestClassOnce.Do(func() {
		DetectHorizonRequestClass = _DetectHorizonRequestClass{objc.GetClass("VNDetectHorizonRequest")}
	})
	return DetectHorizonRequestClass
}

type _DetectHorizonRequestClass struct {
	class objc.Class
}

// An interface definition for the [DetectHorizonRequest] class.
type IDetectHorizonRequest interface {
	IImageBasedRequest
}

// An image-analysis request that determines the horizon angle in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHorizonRequest
type DetectHorizonRequest struct {
	ImageBasedRequest
}

// DetectHorizonRequestFrom constructs a [DetectHorizonRequest] from an unsafe.Pointer.
//
// An image-analysis request that determines the horizon angle in an image.
func DetectHorizonRequestFrom(ptr unsafe.Pointer) DetectHorizonRequest {
	return DetectHorizonRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DetectHorizonRequestClass) Alloc() DetectHorizonRequest {
	rv := objc.Send[DetectHorizonRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DetectHorizonRequestClass) New() DetectHorizonRequest {
	rv := objc.Send[DetectHorizonRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectHorizonRequest) Init() DetectHorizonRequest {
	rv := objc.Send[DetectHorizonRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectHorizonRequest) Autorelease() DetectHorizonRequest {
	rv := objc.Send[DetectHorizonRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectHorizonRequest creates a new DetectHorizonRequest instance.
func NewDetectHorizonRequest() DetectHorizonRequest {
	return getDetectHorizonRequestClass().New()
}


// The results of the horizon detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthorizonrequest/results
func (d_ DetectHorizonRequest) Results() VNHorizonObservation {
	rv := objc.Send[VNHorizonObservation](d_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of the horizon detection request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthorizonrequest/results
func (d_ DetectHorizonRequest) SetResults(value IVNHorizonObservation) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setResults:"), value)
}

// A constant for specifying revision 1 of the horizon detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthorizonrequestrevision1
func (d_ DetectHorizonRequest) VNDetectHorizonRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHorizonRequestRevision1"))
	return rv
}



