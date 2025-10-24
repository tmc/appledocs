// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNRectangleObservation */


/* debug [class_header]: Header for VNRectangleObservation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RectangleObservation */
// An interface definition for the [RectangleObservation] class.
type IRectangleObservation interface {
	IDetectedObjectObservation
	
/* debug [class_interface_properties]: Properties for RectangleObservation */
	// properties:
	BottomLeft() corefoundation.CGPoint
	BottomRight() corefoundation.CGPoint
	TopLeft() corefoundation.CGPoint
	TopRight() corefoundation.CGPoint
	Results() IVNRectangleObservation
	SetResults(value IVNRectangleObservation)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RectangleObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RectangleObservation */
// Alloc allocates a new instance without initialization.
func (rc _RectangleObservationClass) Alloc() RectangleObservation {
	rv := objc.Send[RectangleObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RectangleObservation */
// An object that represents the four vertices of a detected rectangle.


// An object that represents the four vertices of a detected rectangle.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RectangleObservation */

// Creates a rectangle observation from its corner points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:bottomLeft:bottomRight:topRight:)
func NewRectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight(requestRevision uint, topLeft corefoundation.CGPoint, bottomLeft corefoundation.CGPoint, bottomRight corefoundation.CGPoint, topRight corefoundation.CGPoint) RectangleObservation {
	rv := objc.Send[RectangleObservation](objc.ID(getRectangleObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:bottomLeft:bottomRight:topRight:"), requestRevision, topLeft, bottomLeft, bottomRight, topRight)
	return rv
}/* debug [class_init_methods/constructor]: NewRectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight */


// Creates a rectangle observation from its corner points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:topRight:bottomRight:bottomLeft:)
func NewRectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft(requestRevision uint, topLeft corefoundation.CGPoint, topRight corefoundation.CGPoint, bottomRight corefoundation.CGPoint, bottomLeft corefoundation.CGPoint) RectangleObservation {
	rv := objc.Send[RectangleObservation](objc.ID(getRectangleObservationClass().class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:topRight:bottomRight:bottomLeft:"), requestRevision, topLeft, topRight, bottomRight, bottomLeft)
	return rv
}/* debug [class_init_methods/constructor]: NewRectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RectangleObservation */

// Creates a rectangle observation from its corner points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:bottomLeft:bottomRight:topRight:)
func (rc _RectangleObservationClass) RectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight(requestRevision uint, topLeft corefoundation.CGPoint, bottomLeft corefoundation.CGPoint, bottomRight corefoundation.CGPoint, topRight corefoundation.CGPoint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:bottomLeft:bottomRight:topRight:"), requestRevision, topLeft, bottomLeft, bottomRight, topRight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RectangleObservationWithRequestRevisionTopLeftBottomLeftBottomRightTopRight) */


// Creates a rectangle observation from its corner points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/init(requestRevision:topLeft:topRight:bottomRight:bottomLeft:)
func (rc _RectangleObservationClass) RectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft(requestRevision uint, topLeft corefoundation.CGPoint, topRight corefoundation.CGPoint, bottomRight corefoundation.CGPoint, bottomLeft corefoundation.CGPoint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("rectangleObservationWithRequestRevision:topLeft:topRight:bottomRight:bottomLeft:"), requestRevision, topLeft, topRight, bottomRight, bottomLeft)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RectangleObservationWithRequestRevisionTopLeftTopRightBottomRightBottomLeft) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RectangleObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RectangleObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RectangleObservation */

// The coordinates of the lower-left corner of the observation bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/bottomLeft
func (r_ RectangleObservation) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("bottomLeft"))
	return rv
}/* debug [instance_properties/getter]: bottomLeft */


// The coordinates of the lower-right corner of the observation bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/bottomRight
func (r_ RectangleObservation) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("bottomRight"))
	return rv
}/* debug [instance_properties/getter]: bottomRight */


// The coordinates of the upper-left corner of the observation bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/topLeft
func (r_ RectangleObservation) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("topLeft"))
	return rv
}/* debug [instance_properties/getter]: topLeft */


// The coordinates of the upper-right corner of the observation bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRectangleObservation/topRight
func (r_ RectangleObservation) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("topRight"))
	return rv
}/* debug [instance_properties/getter]: topRight */


// The results of a document segmentation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequest/results
func (r_ RectangleObservation) Results() IVNRectangleObservation {
	rv := objc.Send[RectangleObservation](r_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// The results of a document segmentation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectdocumentsegmentationrequest/results
func (r_ RectangleObservation) SetResults(value IVNRectangleObservation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setResults:"), value)
}/* debug [instance_properties/setter]: results */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNRectangleObservation */


