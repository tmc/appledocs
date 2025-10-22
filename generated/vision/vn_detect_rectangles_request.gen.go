// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DetectRectanglesRequest] class.
var (
	DetectRectanglesRequestClass     _DetectRectanglesRequestClass
	DetectRectanglesRequestClassOnce sync.Once
)

func getDetectRectanglesRequestClass() _DetectRectanglesRequestClass {
	DetectRectanglesRequestClassOnce.Do(func() {
		DetectRectanglesRequestClass = _DetectRectanglesRequestClass{objc.GetClass("VNDetectRectanglesRequest")}
	})
	return DetectRectanglesRequestClass
}

type _DetectRectanglesRequestClass struct {
	class objc.Class
}

// An interface definition for the [DetectRectanglesRequest] class.
type IDetectRectanglesRequest interface {
	IImageBasedRequest
	MaximumAspectRatio() AspectRatio
	SetMaximumAspectRatio(value IAspectRatio)
	MaximumObservations() uint
	SetMaximumObservations(value uint)
	MinimumAspectRatio() AspectRatio
	SetMinimumAspectRatio(value IAspectRatio)
	MinimumConfidence() Confidence
	SetMinimumConfidence(value IConfidence)
	MinimumSize() float32
	SetMinimumSize(value float32)
	QuadratureTolerance() Degrees
	SetQuadratureTolerance(value IDegrees)
	Results() VNRectangleObservation
	SetResults(value IVNRectangleObservation)
	VNDetectRectanglesRequestRevision1() int
}

// An image-analysis request that finds projected rectangular regions in an image.
//
// A rectangle detection request locates regions of an image with rectangular shape, like credit cards, business cards, documents, and signs. The request returns its observations in the form of objects, which contain normalized coordinates of bounding boxes containing the rectangle. Use this type of request to find the bounding boxes of rectangles in an image. Vision returns observations for rectangles found in all orientations and sizes, along with a confidence level to indicate how likely it’s that the observation contains an actual rectangle. To further configure or restrict the types of rectangles found, set properties on the request specifying a range of aspect ratios, sizes, and quadrature tolerance.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest
type DetectRectanglesRequest struct {
	ImageBasedRequest
}

// DetectRectanglesRequestFrom constructs a [DetectRectanglesRequest] from an unsafe.Pointer.
//
// An image-analysis request that finds projected rectangular regions in an image.
func DetectRectanglesRequestFrom(ptr unsafe.Pointer) DetectRectanglesRequest {
	return DetectRectanglesRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DetectRectanglesRequestClass) Alloc() DetectRectanglesRequest {
	rv := objc.Send[DetectRectanglesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DetectRectanglesRequestClass) New() DetectRectanglesRequest {
	rv := objc.Send[DetectRectanglesRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectRectanglesRequest) Init() DetectRectanglesRequest {
	rv := objc.Send[DetectRectanglesRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectRectanglesRequest) Autorelease() DetectRectanglesRequest {
	rv := objc.Send[DetectRectanglesRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectRectanglesRequest creates a new DetectRectanglesRequest instance.
func NewDetectRectanglesRequest() DetectRectanglesRequest {
	return getDetectRectanglesRequestClass().New()
}


// A specifying the maximum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumAspectRatio
func (d_ DetectRectanglesRequest) MaximumAspectRatio() AspectRatio {
	rv := objc.Send[AspectRatio](d_.ID, objc.Sel("maximumAspectRatio"))
	return rv
}


// SetMaximumAspectRatio sets the value of the maximumAspectRatio property.
// A specifying the maximum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumAspectRatio
func (d_ DetectRectanglesRequest) SetMaximumAspectRatio(value IAspectRatio) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumAspectRatio:"), value)
}

// An integer specifying the maximum number of rectangles Vision returns.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumObservations
func (d_ DetectRectanglesRequest) MaximumObservations() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("maximumObservations"))
	return rv
}


// SetMaximumObservations sets the value of the maximumObservations property.
// An integer specifying the maximum number of rectangles Vision returns.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumObservations
func (d_ DetectRectanglesRequest) SetMaximumObservations(value uint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumObservations:"), value)
}

// A specifying the minimum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumAspectRatio
func (d_ DetectRectanglesRequest) MinimumAspectRatio() AspectRatio {
	rv := objc.Send[AspectRatio](d_.ID, objc.Sel("minimumAspectRatio"))
	return rv
}


// SetMinimumAspectRatio sets the value of the minimumAspectRatio property.
// A specifying the minimum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumAspectRatio
func (d_ DetectRectanglesRequest) SetMinimumAspectRatio(value IAspectRatio) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinimumAspectRatio:"), value)
}

// A value specifying the minimum acceptable confidence level.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumConfidence
func (d_ DetectRectanglesRequest) MinimumConfidence() Confidence {
	rv := objc.Send[Confidence](d_.ID, objc.Sel("minimumConfidence"))
	return rv
}


// SetMinimumConfidence sets the value of the minimumConfidence property.
// A value specifying the minimum acceptable confidence level.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumConfidence
func (d_ DetectRectanglesRequest) SetMinimumConfidence(value IConfidence) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinimumConfidence:"), value)
}

// The minimum size of a rectangle to detect, as a proportion of the smallest dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumSize
func (d_ DetectRectanglesRequest) MinimumSize() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("minimumSize"))
	return rv
}


// SetMinimumSize sets the value of the minimumSize property.
// The minimum size of a rectangle to detect, as a proportion of the smallest dimension.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumSize
func (d_ DetectRectanglesRequest) SetMinimumSize(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinimumSize:"), value)
}

// A float specifying the number of degrees a rectangle corner angle can deviate from 90°.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/quadraturetolerance
func (d_ DetectRectanglesRequest) QuadratureTolerance() Degrees {
	rv := objc.Send[Degrees](d_.ID, objc.Sel("quadratureTolerance"))
	return rv
}


// SetQuadratureTolerance sets the value of the quadratureTolerance property.
// A float specifying the number of degrees a rectangle corner angle can deviate from 90°.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/quadraturetolerance
func (d_ DetectRectanglesRequest) SetQuadratureTolerance(value IDegrees) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setQuadratureTolerance:"), value)
}

// The results of the request to detect rectangles.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/results
func (d_ DetectRectanglesRequest) Results() VNRectangleObservation {
	rv := objc.Send[VNRectangleObservation](d_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of the request to detect rectangles.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequest/results
func (d_ DetectRectanglesRequest) SetResults(value IVNRectangleObservation) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setResults:"), value)
}

// A constant for specifying revision 1 of the rectangle detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequestrevision1
func (d_ DetectRectanglesRequest) VNDetectRectanglesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectRectanglesRequestRevision1"))
	return rv
}



