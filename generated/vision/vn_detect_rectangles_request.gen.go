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
	

	// properties:
	MaximumAspectRatio() AspectRatio /* typedef */
	SetMaximumAspectRatio(value AspectRatio /* typedef */)
	MaximumObservations() uint
	SetMaximumObservations(value uint)
	MinimumAspectRatio() AspectRatio /* typedef */
	SetMinimumAspectRatio(value AspectRatio /* typedef */)
	MinimumConfidence() Confidence /* typedef */
	SetMinimumConfidence(value Confidence /* typedef */)
	MinimumSize() float32
	SetMinimumSize(value float32)
	QuadratureTolerance() Degrees /* typedef */
	SetQuadratureTolerance(value Degrees /* typedef */)
	Results() []RectangleObservation
	VNDetectRectanglesRequestRevision1() int


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DetectRectanglesRequestClass) Alloc() DetectRectanglesRequest {
	rv := objc.Send[DetectRectanglesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An image-analysis request that finds projected rectangular regions in an image.
//
// A rectangle detection request locates regions of an image with rectangular shape, like credit cards, business cards, documents, and signs. The request returns its observations in the form of objects, which contain normalized coordinates of bounding boxes containing the rectangle. Use this type of request to find the bounding boxes of rectangles in an image. Vision returns observations for rectangles found in all orientations and sizes, along with a confidence level to indicate how likely it’s that the observation contains an actual rectangle. To further configure or restrict the types of rectangles found, set properties on the request specifying a range of aspect ratios, sizes, and quadrature tolerance.


// An image-analysis request that finds projected rectangular regions in an image.
//
// [Full Topic]
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

























// A specifying the maximum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumAspectRatio
func (d_ DetectRectanglesRequest) MaximumAspectRatio() AspectRatio /* typedef */ {
	rv := objc.Send[float32](d_.ID, objc.Sel("maximumAspectRatio"))
	return rv
}


// A specifying the maximum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumAspectRatio
func (d_ DetectRectanglesRequest) SetMaximumAspectRatio(value AspectRatio /* typedef */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumAspectRatio:"), value)
}


// An integer specifying the maximum number of rectangles Vision returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumObservations
func (d_ DetectRectanglesRequest) MaximumObservations() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("maximumObservations"))
	return rv
}


// An integer specifying the maximum number of rectangles Vision returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/maximumObservations
func (d_ DetectRectanglesRequest) SetMaximumObservations(value uint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumObservations:"), value)
}


// A specifying the minimum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumAspectRatio
func (d_ DetectRectanglesRequest) MinimumAspectRatio() AspectRatio /* typedef */ {
	rv := objc.Send[float32](d_.ID, objc.Sel("minimumAspectRatio"))
	return rv
}


// A specifying the minimum aspect ratio of the rectangle to detect, defined as the shorter dimension over the longer dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumAspectRatio
func (d_ DetectRectanglesRequest) SetMinimumAspectRatio(value AspectRatio /* typedef */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinimumAspectRatio:"), value)
}


// A value specifying the minimum acceptable confidence level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumConfidence
func (d_ DetectRectanglesRequest) MinimumConfidence() Confidence /* typedef */ {
	rv := objc.Send[float32](d_.ID, objc.Sel("minimumConfidence"))
	return rv
}


// A value specifying the minimum acceptable confidence level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumConfidence
func (d_ DetectRectanglesRequest) SetMinimumConfidence(value Confidence /* typedef */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinimumConfidence:"), value)
}


// The minimum size of a rectangle to detect, as a proportion of the smallest dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumSize
func (d_ DetectRectanglesRequest) MinimumSize() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("minimumSize"))
	return rv
}


// The minimum size of a rectangle to detect, as a proportion of the smallest dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/minimumSize
func (d_ DetectRectanglesRequest) SetMinimumSize(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinimumSize:"), value)
}


// A float specifying the number of degrees a rectangle corner angle can deviate from 90°.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/quadratureTolerance
func (d_ DetectRectanglesRequest) QuadratureTolerance() Degrees /* typedef */ {
	rv := objc.Send[float32](d_.ID, objc.Sel("quadratureTolerance"))
	return rv
}


// A float specifying the number of degrees a rectangle corner angle can deviate from 90°.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/quadratureTolerance
func (d_ DetectRectanglesRequest) SetQuadratureTolerance(value Degrees /* typedef */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setQuadratureTolerance:"), value)
}


// The results of the request to detect rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectRectanglesRequest/results
func (d_ DetectRectanglesRequest) Results() []RectangleObservation {
	rv := objc.Send[[]RectangleObservation](d_.ID, objc.Sel("results"))
	return rv
}


// A constant for specifying revision 1 of the rectangle detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectrectanglesrequestrevision1
func (d_ DetectRectanglesRequest) VNDetectRectanglesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectRectanglesRequestRevision1"))
	return rv
}








