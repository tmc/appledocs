// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GenerateObjectnessBasedSaliencyImageRequest] class.
var (
	GenerateObjectnessBasedSaliencyImageRequestClass     _GenerateObjectnessBasedSaliencyImageRequestClass
	GenerateObjectnessBasedSaliencyImageRequestClassOnce sync.Once
)

func getGenerateObjectnessBasedSaliencyImageRequestClass() _GenerateObjectnessBasedSaliencyImageRequestClass {
	GenerateObjectnessBasedSaliencyImageRequestClassOnce.Do(func() {
		GenerateObjectnessBasedSaliencyImageRequestClass = _GenerateObjectnessBasedSaliencyImageRequestClass{objc.GetClass("VNGenerateObjectnessBasedSaliencyImageRequest")}
	})
	return GenerateObjectnessBasedSaliencyImageRequestClass
}

type _GenerateObjectnessBasedSaliencyImageRequestClass struct {
	class objc.Class
}

// An interface definition for the [GenerateObjectnessBasedSaliencyImageRequest] class.
type IGenerateObjectnessBasedSaliencyImageRequest interface {
	IImageBasedRequest
}

// A request that generates a heat map that identifies the parts of an image most likely to represent objects.
//
// The resulting observation, , encodes this data as a heat map, which you can use to highlight regions of interest.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateObjectnessBasedSaliencyImageRequest
type GenerateObjectnessBasedSaliencyImageRequest struct {
	ImageBasedRequest
}

// GenerateObjectnessBasedSaliencyImageRequestFrom constructs a [GenerateObjectnessBasedSaliencyImageRequest] from an unsafe.Pointer.
//
// A request that generates a heat map that identifies the parts of an image most likely to represent objects.
func GenerateObjectnessBasedSaliencyImageRequestFrom(ptr unsafe.Pointer) GenerateObjectnessBasedSaliencyImageRequest {
	return GenerateObjectnessBasedSaliencyImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GenerateObjectnessBasedSaliencyImageRequestClass) Alloc() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[GenerateObjectnessBasedSaliencyImageRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GenerateObjectnessBasedSaliencyImageRequestClass) New() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[GenerateObjectnessBasedSaliencyImageRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenerateObjectnessBasedSaliencyImageRequest) Init() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[GenerateObjectnessBasedSaliencyImageRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenerateObjectnessBasedSaliencyImageRequest) Autorelease() GenerateObjectnessBasedSaliencyImageRequest {
	rv := objc.Send[GenerateObjectnessBasedSaliencyImageRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenerateObjectnessBasedSaliencyImageRequest creates a new GenerateObjectnessBasedSaliencyImageRequest instance.
func NewGenerateObjectnessBasedSaliencyImageRequest() GenerateObjectnessBasedSaliencyImageRequest {
	return getGenerateObjectnessBasedSaliencyImageRequestClass().New()
}


// The results of the image saliency request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateobjectnessbasedsaliencyimagerequest/results
func (g_ GenerateObjectnessBasedSaliencyImageRequest) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of the image saliency request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateobjectnessbasedsaliencyimagerequest/results
func (g_ GenerateObjectnessBasedSaliencyImageRequest) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResults:"), value)
}

// A constant for specifying revision 1 of the image saliency request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateobjectnessbasedsaliencyimagerequestrevision1
func (g_ GenerateObjectnessBasedSaliencyImageRequest) VNGenerateObjectnessBasedSaliencyImageRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateObjectnessBasedSaliencyImageRequestRevision1"))
	return rv
}



