// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


// The amount by which to adjust the image contrast.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/contrastadjustment
func (d_ DetectContoursRequest) ContrastAdjustment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("contrastAdjustment"))
	return rv
}


// SetContrastAdjustment sets the value of the contrastAdjustment property.
// The amount by which to adjust the image contrast.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/contrastadjustment
func (d_ DetectContoursRequest) SetContrastAdjustment(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContrastAdjustment:"), value)
}

// A constant for specifying revision 1 of the contours detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontourrequestrevision1
func (d_ DetectContoursRequest) VNDetectContourRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectContourRequestRevision1"))
	return rv
}

// The pixel value to use as a pivot for the contrast.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/contrastpivot
func (d_ DetectContoursRequest) ContrastPivot() foundation.Number {
	rv := objc.Send[foundation.Number](d_.ID, objc.Sel("contrastPivot"))
	return rv
}


// SetContrastPivot sets the value of the contrastPivot property.
// The pixel value to use as a pivot for the contrast.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/contrastpivot
func (d_ DetectContoursRequest) SetContrastPivot(value foundation.Number) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContrastPivot:"), value)
}

// The maximum image dimension to use for contour detection.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/maximumimagedimension
func (d_ DetectContoursRequest) MaximumImageDimension() int {
	rv := objc.Send[int](d_.ID, objc.Sel("maximumImageDimension"))
	return rv
}


// SetMaximumImageDimension sets the value of the maximumImageDimension property.
// The maximum image dimension to use for contour detection.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/maximumimagedimension
func (d_ DetectContoursRequest) SetMaximumImageDimension(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumImageDimension:"), value)
}

// A Boolean value that indicates whether the request detects a dark object on a light background to aid in detection.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/detectsdarkonlight
func (d_ DetectContoursRequest) DetectsDarkOnLight() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("detectsDarkOnLight"))
	return rv
}


// SetDetectsDarkOnLight sets the value of the detectsDarkOnLight property.
// A Boolean value that indicates whether the request detects a dark object on a light background to aid in detection.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/detectsdarkonlight
func (d_ DetectContoursRequest) SetDetectsDarkOnLight(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDetectsDarkOnLight:"), value)
}

// A Boolean value that indicates whether the request detects a dark object on a light background.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/detectdarkonlight
func (d_ DetectContoursRequest) DetectDarkOnLight() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("detectDarkOnLight"))
	return rv
}


// SetDetectDarkOnLight sets the value of the detectDarkOnLight property.
// A Boolean value that indicates whether the request detects a dark object on a light background.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectcontoursrequest/detectdarkonlight
func (d_ DetectContoursRequest) SetDetectDarkOnLight(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDetectDarkOnLight:"), value)
}

// The results of the request to detect contours.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectContoursRequest/results
func (d_ DetectContoursRequest) Results() []ContoursObservation {
	rv := objc.Send[[]ContoursObservation](d_.ID, objc.Sel("results"))
	return rv
}



