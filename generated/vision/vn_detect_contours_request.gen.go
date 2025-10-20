// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DetectContoursRequest] class.
var (
	DetectContoursRequestClass     _DetectContoursRequestClass
	DetectContoursRequestClassOnce sync.Once
)

func getDetectContoursRequestClass() _DetectContoursRequestClass {
	DetectContoursRequestClassOnce.Do(func() {
		DetectContoursRequestClass = _DetectContoursRequestClass{objc.GetClass("VNDetectContoursRequest")}
	})
	return DetectContoursRequestClass
}

type _DetectContoursRequestClass struct {
	class objc.Class
}

// An interface definition for the [DetectContoursRequest] class.
type IDetectContoursRequest interface {
	IImageBasedRequest
}

// A request that detects the contours of the edges of an image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest
type DetectContoursRequest struct {
	ImageBasedRequest
}

// DetectContoursRequestFrom constructs a [DetectContoursRequest] from an unsafe.Pointer.
//
// A request that detects the contours of the edges of an image.
func DetectContoursRequestFrom(ptr unsafe.Pointer) DetectContoursRequest {
	return DetectContoursRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DetectContoursRequestClass) Alloc() DetectContoursRequest {
	rv := objc.Send[DetectContoursRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DetectContoursRequestClass) New() DetectContoursRequest {
	rv := objc.Send[DetectContoursRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectContoursRequest) Init() DetectContoursRequest {
	rv := objc.Send[DetectContoursRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectContoursRequest) Autorelease() DetectContoursRequest {
	rv := objc.Send[DetectContoursRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectContoursRequest creates a new DetectContoursRequest instance.
func NewDetectContoursRequest() DetectContoursRequest {
	return getDetectContoursRequestClass().New()
}


// The results of the request to detect contours.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/results
func (d_ DetectContoursRequest) Results() []ContoursObservation {
	rv := objc.Send[[]ContoursObservation](d_.ID, objc.Sel("results"))
	return rv
}



