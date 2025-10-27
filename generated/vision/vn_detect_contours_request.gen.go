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
	

	// properties:
	ContrastAdjustment() float32
	SetContrastAdjustment(value float32)
	ContrastPivot() foundation.foundation.INSNumber
	SetContrastPivot(value foundation.foundation.INSNumber)
	DetectDarkOnLight() bool
	SetDetectDarkOnLight(value bool)
	DetectsDarkOnLight() bool
	SetDetectsDarkOnLight(value bool)
	MaximumImageDimension() uint
	SetMaximumImageDimension(value uint)
	Results() []ContoursObservation
	VNDetectContourRequestRevision1() int


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DetectContoursRequestClass) Alloc() DetectContoursRequest {
	rv := objc.Send[DetectContoursRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A request that detects the contours of the edges of an image.


// A request that detects the contours of the edges of an image.
//
// [Full Topic]
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

























// The amount by which to adjust the image contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/contrastAdjustment
func (d_ DetectContoursRequest) ContrastAdjustment() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("contrastAdjustment"))
	return rv
}


// The amount by which to adjust the image contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/contrastAdjustment
func (d_ DetectContoursRequest) SetContrastAdjustment(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContrastAdjustment:"), value)
}


// The pixel value to use as a pivot for the contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/contrastPivot
func (d_ DetectContoursRequest) ContrastPivot() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](d_.ID, objc.Sel("contrastPivot"))
	return rv
}


// The pixel value to use as a pivot for the contrast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/contrastPivot
func (d_ DetectContoursRequest) SetContrastPivot(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContrastPivot:"), value)
}


// A Boolean value that indicates whether the request detects a dark object on a light background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/detectDarkOnLight
func (d_ DetectContoursRequest) DetectDarkOnLight() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("detectDarkOnLight"))
	return rv
}


// A Boolean value that indicates whether the request detects a dark object on a light background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/detectDarkOnLight
func (d_ DetectContoursRequest) SetDetectDarkOnLight(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDetectDarkOnLight:"), value)
}


// A Boolean value that indicates whether the request detects a dark object on a light background to aid in detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/detectsDarkOnLight
func (d_ DetectContoursRequest) DetectsDarkOnLight() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("detectsDarkOnLight"))
	return rv
}


// A Boolean value that indicates whether the request detects a dark object on a light background to aid in detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/detectsDarkOnLight
func (d_ DetectContoursRequest) SetDetectsDarkOnLight(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDetectsDarkOnLight:"), value)
}


// The maximum image dimension to use for contour detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/maximumImageDimension
func (d_ DetectContoursRequest) MaximumImageDimension() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("maximumImageDimension"))
	return rv
}


// The maximum image dimension to use for contour detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/maximumImageDimension
func (d_ DetectContoursRequest) SetMaximumImageDimension(value uint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumImageDimension:"), value)
}


// The results of the request to detect contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/results
func (d_ DetectContoursRequest) Results() []ContoursObservation {
	rv := objc.Send[[]ContoursObservation](d_.ID, objc.Sel("results"))
	return rv
}


// A constant for specifying revision 1 of the contours detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontourrequestrevision1
func (d_ DetectContoursRequest) VNDetectContourRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectContourRequestRevision1"))
	return rv
}








