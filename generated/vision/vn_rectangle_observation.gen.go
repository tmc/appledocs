// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [RectangleObservation] class.
var (
	RectangleObservationClass     _RectangleObservationClass
	RectangleObservationClassOnce sync.Once
)

func getRectangleObservationClass() _RectangleObservationClass {
	RectangleObservationClassOnce.Do(func() {
		RectangleObservationClass = _RectangleObservationClass{objc.GetClass("VNRectangleObservation")}
	})
	return RectangleObservationClass
}

type _RectangleObservationClass struct {
	class objc.Class
}

// An interface definition for the [RectangleObservation] class.
type IRectangleObservation interface {
	IDetectedObjectObservation
	TopRight() coregraphics.CGPoint
	Results() VNRectangleObservation
	SetResults(value IVNRectangleObservation)
	BottomLeft() coregraphics.CGPoint
	SetBottomLeft(value coregraphics.CGPoint)
	BottomRight() coregraphics.CGPoint
	SetBottomRight(value coregraphics.CGPoint)
	TopLeft() coregraphics.CGPoint
	SetTopLeft(value coregraphics.CGPoint)
}

// An object that represents the four vertices of a detected rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation
type RectangleObservation struct {
	DetectedObjectObservation
}

// RectangleObservationFrom constructs a [RectangleObservation] from an unsafe.Pointer.
//
// An object that represents the four vertices of a detected rectangle.
func RectangleObservationFrom(ptr unsafe.Pointer) RectangleObservation {
	return RectangleObservation{
		DetectedObjectObservation: DetectedObjectObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RectangleObservationClass) Alloc() RectangleObservation {
	rv := objc.Send[RectangleObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RectangleObservationClass) New() RectangleObservation {
	rv := objc.Send[RectangleObservation](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RectangleObservation) Init() RectangleObservation {
	rv := objc.Send[RectangleObservation](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RectangleObservation) Autorelease() RectangleObservation {
	rv := objc.Send[RectangleObservation](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRectangleObservation creates a new RectangleObservation instance.
func NewRectangleObservation() RectangleObservation {
	return getRectangleObservationClass().New()
}




// Creates a rectangle observation from its corner points.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:bottomLeft:bottomRight:topRight:)
func NewRectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight(requestRevision uint, topLeft coregraphics.CGPoint, bottomLeft coregraphics.CGPoint, bottomRight coregraphics.CGPoint, topRight coregraphics.CGPoint) RectangleObservation {
	rv := objc.Send[RectangleObservation](objc.ID(getRectangleObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:bottomLeft:bottomRight:topRight:"), requestRevision, topLeft, bottomLeft, bottomRight, topRight)
	return rv
}



// Creates a rectangle observation from its corner points.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:topRight:bottomRight:bottomLeft:)
func NewRectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft(requestRevision uint, topLeft coregraphics.CGPoint, topRight coregraphics.CGPoint, bottomRight coregraphics.CGPoint, bottomLeft coregraphics.CGPoint) RectangleObservation {
	rv := objc.Send[RectangleObservation](objc.ID(getRectangleObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:topRight:bottomRight:bottomLeft:"), requestRevision, topLeft, topRight, bottomRight, bottomLeft)
	return rv
}


// Creates a rectangle observation from its corner points.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:bottomLeft:bottomRight:topRight:)
func (rc _RectangleObservationClass) RectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight(requestRevision uint, topLeft coregraphics.CGPoint, bottomLeft coregraphics.CGPoint, bottomRight coregraphics.CGPoint, topRight coregraphics.CGPoint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:bottomLeft:bottomRight:topRight:"), requestRevision, topLeft, bottomLeft, bottomRight, topRight)
	return rv
}

// Creates a rectangle observation from its corner points.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:topRight:bottomRight:bottomLeft:)
func (rc _RectangleObservationClass) RectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft(requestRevision uint, topLeft coregraphics.CGPoint, topRight coregraphics.CGPoint, bottomRight coregraphics.CGPoint, bottomLeft coregraphics.CGPoint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:topRight:bottomRight:bottomLeft:"), requestRevision, topLeft, topRight, bottomRight, bottomLeft)
	return rv
}

// The coordinates of the upper-right corner of the observation bounding box.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/topRight
func (r_ RectangleObservation) TopRight() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("topRight"))
	return rv
}

// The results of a document segmentation request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequest/results
func (r_ RectangleObservation) Results() VNRectangleObservation {
	rv := objc.Send[VNRectangleObservation](r_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of a document segmentation request.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequest/results
func (r_ RectangleObservation) SetResults(value IVNRectangleObservation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setResults:"), value)
}

// The coordinates of the lower-left corner of the observation bounding box.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/bottomleft
func (r_ RectangleObservation) BottomLeft() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("bottomLeft"))
	return rv
}


// SetBottomLeft sets the value of the bottomLeft property.
// The coordinates of the lower-left corner of the observation bounding box.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/bottomleft
func (r_ RectangleObservation) SetBottomLeft(value coregraphics.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBottomLeft:"), value)
}

// The coordinates of the lower-right corner of the observation bounding box.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/bottomright
func (r_ RectangleObservation) BottomRight() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("bottomRight"))
	return rv
}


// SetBottomRight sets the value of the bottomRight property.
// The coordinates of the lower-right corner of the observation bounding box.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/bottomright
func (r_ RectangleObservation) SetBottomRight(value coregraphics.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBottomRight:"), value)
}

// The coordinates of the upper-left corner of the observation bounding box.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/topleft
func (r_ RectangleObservation) TopLeft() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("topLeft"))
	return rv
}


// SetTopLeft sets the value of the topLeft property.
// The coordinates of the upper-left corner of the observation bounding box.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrectangleobservation/topleft
func (r_ RectangleObservation) SetTopLeft(value coregraphics.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTopLeft:"), value)
}


