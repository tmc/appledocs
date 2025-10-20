// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DetectTextRectanglesRequest] class.
var (
	DetectTextRectanglesRequestClass     _DetectTextRectanglesRequestClass
	DetectTextRectanglesRequestClassOnce sync.Once
)

func getDetectTextRectanglesRequestClass() _DetectTextRectanglesRequestClass {
	DetectTextRectanglesRequestClassOnce.Do(func() {
		DetectTextRectanglesRequestClass = _DetectTextRectanglesRequestClass{objc.GetClass("VNDetectTextRectanglesRequest")}
	})
	return DetectTextRectanglesRequestClass
}

type _DetectTextRectanglesRequestClass struct {
	class objc.Class
}

// An interface definition for the [DetectTextRectanglesRequest] class.
type IDetectTextRectanglesRequest interface {
	IImageBasedRequest
}

// An image-analysis request that finds regions of visible text in an image.
//
// This request returns detected text characters as rectangular bounding boxes with origin and size.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTextRectanglesRequest
type DetectTextRectanglesRequest struct {
	ImageBasedRequest
}

// DetectTextRectanglesRequestFrom constructs a [DetectTextRectanglesRequest] from an unsafe.Pointer.
//
// An image-analysis request that finds regions of visible text in an image.
func DetectTextRectanglesRequestFrom(ptr unsafe.Pointer) DetectTextRectanglesRequest {
	return DetectTextRectanglesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DetectTextRectanglesRequestClass) Alloc() DetectTextRectanglesRequest {
	rv := objc.Send[DetectTextRectanglesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DetectTextRectanglesRequestClass) New() DetectTextRectanglesRequest {
	rv := objc.Send[DetectTextRectanglesRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectTextRectanglesRequest) Init() DetectTextRectanglesRequest {
	rv := objc.Send[DetectTextRectanglesRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectTextRectanglesRequest) Autorelease() DetectTextRectanglesRequest {
	rv := objc.Send[DetectTextRectanglesRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectTextRectanglesRequest creates a new DetectTextRectanglesRequest instance.
func NewDetectTextRectanglesRequest() DetectTextRectanglesRequest {
	return getDetectTextRectanglesRequestClass().New()
}




