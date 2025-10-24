// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [DetectHumanRectanglesRequest] class.
var (
	DetectHumanRectanglesRequestClass     _DetectHumanRectanglesRequestClass
	DetectHumanRectanglesRequestClassOnce sync.Once
)

func getDetectHumanRectanglesRequestClass() _DetectHumanRectanglesRequestClass {
	DetectHumanRectanglesRequestClassOnce.Do(func() {
		DetectHumanRectanglesRequestClass = _DetectHumanRectanglesRequestClass{objc.GetClass("VNDetectHumanRectanglesRequest")}
	})
	return DetectHumanRectanglesRequestClass
}

type _DetectHumanRectanglesRequestClass struct {
	class objc.Class
}





// An interface definition for the [DetectHumanRectanglesRequest] class.
type IDetectHumanRectanglesRequest interface {
	IImageBasedRequest
	

	// properties:
	Results() []HumanObservation
	UpperBodyOnly() bool
	SetUpperBodyOnly(value bool)
	VNDetectHumanRectanglesRequestRevision1() int
	VNDetectHumanRectanglesRequestRevision2() int


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DetectHumanRectanglesRequestClass) Alloc() DetectHumanRectanglesRequest {
	rv := objc.Send[DetectHumanRectanglesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectHumanRectanglesRequestClass) New() DetectHumanRectanglesRequest {
	rv := objc.Send[DetectHumanRectanglesRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectHumanRectanglesRequest) Init() DetectHumanRectanglesRequest {
	rv := objc.Send[DetectHumanRectanglesRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectHumanRectanglesRequest) Autorelease() DetectHumanRectanglesRequest {
	rv := objc.Send[DetectHumanRectanglesRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectHumanRectanglesRequest creates a new DetectHumanRectanglesRequest instance.
func NewDetectHumanRectanglesRequest() DetectHumanRectanglesRequest {
	return getDetectHumanRectanglesRequestClass().New()
}





// A request that finds rectangular regions that contain people in an image.


// A request that finds rectangular regions that contain people in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanRectanglesRequest
type DetectHumanRectanglesRequest struct {
	ImageBasedRequest
}

// DetectHumanRectanglesRequestFrom constructs a [DetectHumanRectanglesRequest] from an unsafe.Pointer.
//
// A request that finds rectangular regions that contain people in an image.
func DetectHumanRectanglesRequestFrom(ptr unsafe.Pointer) DetectHumanRectanglesRequest {
	return DetectHumanRectanglesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

























// The results of the request to find rectangular regions that contain people in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanRectanglesRequest/results
func (d_ DetectHumanRectanglesRequest) Results() []HumanObservation {
	rv := objc.Send[[]HumanObservation](d_.ID, objc.Sel("results"))
	return rv
}


// A Boolean value that indicates whether the request requires detecting a full body or upper body only to produce a result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanRectanglesRequest/upperBodyOnly
func (d_ DetectHumanRectanglesRequest) UpperBodyOnly() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("upperBodyOnly"))
	return rv
}


// A Boolean value that indicates whether the request requires detecting a full body or upper body only to produce a result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanRectanglesRequest/upperBodyOnly
func (d_ DetectHumanRectanglesRequest) SetUpperBodyOnly(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUpperBodyOnly:"), value)
}


// A constant for specifying revision 1 of the human rectangles detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanrectanglesrequestrevision1
func (d_ DetectHumanRectanglesRequest) VNDetectHumanRectanglesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHumanRectanglesRequestRevision1"))
	return rv
}


// A constant for specifying revision 2 of the human rectangles detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanrectanglesrequestrevision2
func (d_ DetectHumanRectanglesRequest) VNDetectHumanRectanglesRequestRevision2() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHumanRectanglesRequestRevision2"))
	return rv
}








