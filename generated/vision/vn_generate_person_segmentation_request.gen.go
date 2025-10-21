// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GeneratePersonSegmentationRequest] class.
var (
	GeneratePersonSegmentationRequestClass     _GeneratePersonSegmentationRequestClass
	GeneratePersonSegmentationRequestClassOnce sync.Once
)

func getGeneratePersonSegmentationRequestClass() _GeneratePersonSegmentationRequestClass {
	GeneratePersonSegmentationRequestClassOnce.Do(func() {
		GeneratePersonSegmentationRequestClass = _GeneratePersonSegmentationRequestClass{objc.GetClass("VNGeneratePersonSegmentationRequest")}
	})
	return GeneratePersonSegmentationRequestClass
}

type _GeneratePersonSegmentationRequestClass struct {
	class objc.Class
}

// An interface definition for the [GeneratePersonSegmentationRequest] class.
type IGeneratePersonSegmentationRequest interface {
	IStatefulRequest
}

// An object that produces a matte image for a person it finds in the input image.
//
// Perform this request to detect and generate an image mask for a person in an image. The request returns the resulting image mask in an instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest
type GeneratePersonSegmentationRequest struct {
	StatefulRequest
}

// GeneratePersonSegmentationRequestFrom constructs a [GeneratePersonSegmentationRequest] from an unsafe.Pointer.
//
// An object that produces a matte image for a person it finds in the input image.
func GeneratePersonSegmentationRequestFrom(ptr unsafe.Pointer) GeneratePersonSegmentationRequest {
	return GeneratePersonSegmentationRequest{
		StatefulRequest: StatefulRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GeneratePersonSegmentationRequestClass) Alloc() GeneratePersonSegmentationRequest {
	rv := objc.Send[GeneratePersonSegmentationRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GeneratePersonSegmentationRequestClass) New() GeneratePersonSegmentationRequest {
	rv := objc.Send[GeneratePersonSegmentationRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GeneratePersonSegmentationRequest) Init() GeneratePersonSegmentationRequest {
	rv := objc.Send[GeneratePersonSegmentationRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GeneratePersonSegmentationRequest) Autorelease() GeneratePersonSegmentationRequest {
	rv := objc.Send[GeneratePersonSegmentationRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGeneratePersonSegmentationRequest creates a new GeneratePersonSegmentationRequest instance.
func NewGeneratePersonSegmentationRequest() GeneratePersonSegmentationRequest {
	return getGeneratePersonSegmentationRequestClass().New()
}


// The pixel format of the output image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/outputPixelFormat
func (g_ GeneratePersonSegmentationRequest) OutputPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("outputPixelFormat"))
	return rv
}


// SetOutputPixelFormat sets the value of the outputPixelFormat property.
// The pixel format of the output image.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/outputPixelFormat
func (g_ GeneratePersonSegmentationRequest) SetOutputPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputPixelFormat:"), value)
}

// A value that indicates how the request balances accuracy and performance.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest/qualitylevel-swift.property
func (g_ GeneratePersonSegmentationRequest) QualityLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("qualityLevel"))
	return rv
}


// SetQualityLevel sets the value of the qualityLevel property.
// A value that indicates how the request balances accuracy and performance.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest/qualitylevel-swift.property
func (g_ GeneratePersonSegmentationRequest) SetQualityLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setQualityLevel:"), value)
}

// The results of the segmentation request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest/results
func (g_ GeneratePersonSegmentationRequest) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of the segmentation request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequest/results
func (g_ GeneratePersonSegmentationRequest) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResults:"), value)
}

// A constant for specifying revision 1 of the person segmentation generation request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequestrevision1
func (g_ GeneratePersonSegmentationRequest) VNGeneratePersonSegmentationRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGeneratePersonSegmentationRequestRevision1"))
	return rv
}



