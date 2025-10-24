// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNContour */


/* debug [class_header]: Header for VNContour */
// The class instance for the [Contour] class.
var (
	ContourClass     _ContourClass
	ContourClassOnce sync.Once
)

func getContourClass() _ContourClass {
	ContourClassOnce.Do(func() {
		ContourClass = _ContourClass{objc.GetClass("VNContour")}
	})
	return ContourClass
}

type _ContourClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Contour */
// An interface definition for the [Contour] class.
type IContour interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Contour */
	// properties:
	AspectRatio() float32
	ChildContourCount() int
	ChildContours() []Contour
	IndexPath() foundation.IndexPath
	NormalizedPath() PathRef /* not a class type */
	NormalizedPoints() objectivec.IObject
	PointCount() int
	ContourCount() int
	SetContourCount(value int)
	TopLevelContourCount() int
	SetTopLevelContourCount(value int)
	TopLevelContours() IVNContour
	SetTopLevelContours(value IVNContour)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Contour */
	// methods:
	ChildContourAtIndexError(childContourIndex uint, error_ objectivec.IObject) IContour
	PolygonApproximationWithEpsilonError(epsilon float32, error_ objectivec.IObject) IContour
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Contour */
// Alloc allocates a new instance without initialization.
func (cc _ContourClass) Alloc() Contour {
	rv := objc.Send[Contour](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContourClass) New() Contour {
	rv := objc.Send[Contour](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Contour) Init() Contour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Contour) Autorelease() Contour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContour creates a new Contour instance.
func NewContour() Contour {
	return getContourClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Contour */
// A class that represents a detected contour in an image.


// A class that represents a detected contour in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour
type Contour struct {
	objectivec.Object
}

// ContourFrom constructs a [Contour] from an unsafe.Pointer.
//
// A class that represents a detected contour in an image.
func ContourFrom(ptr unsafe.Pointer) Contour {
	return Contour{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Contour *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Contour */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Contour */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Contour */

// Retrieves the child contour object at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour/childContour(at:)
func (c_ Contour) ChildContourAtIndexError(childContourIndex uint, error_ objectivec.IObject) IContour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("childContourAtIndex:error:"), childContourIndex, error_)
	return rv
}/* debug [instance_methods/method]: ChildContourAtIndexError */


// Simplifies the contour to a polygon using a Ramer-Douglas-Peucker algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour/polygonApproximation(epsilon:)
func (c_ Contour) PolygonApproximationWithEpsilonError(epsilon float32, error_ objectivec.IObject) IContour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("polygonApproximationWithEpsilon:error:"), epsilon, error_)
	return rv
}/* debug [instance_methods/method]: PolygonApproximationWithEpsilonError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Contour */

// The aspect ratio of the contour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour/aspectRatio
func (c_ Contour) AspectRatio() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("aspectRatio"))
	return rv
}/* debug [instance_properties/getter]: aspectRatio */


// The total number of detected child contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour/childContourCount
func (c_ Contour) ChildContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("childContourCount"))
	return rv
}/* debug [instance_properties/getter]: childContourCount */


// An array of contours that this contour encloses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour/childContours
func (c_ Contour) ChildContours() []Contour {
	rv := objc.Send[[]Contour](c_.ID, objc.Sel("childContours"))
	return rv
}/* debug [instance_properties/getter]: childContours */


// The contour object’s index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour/indexPath
func (c_ Contour) IndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("indexPath"))
	return rv
}/* debug [instance_properties/getter]: indexPath */


// The contour object as a path in normalized coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour/normalizedPath
func (c_ Contour) NormalizedPath() PathRef /* not a class type */ {
	rv := objc.Send[PathRef](c_.ID, objc.Sel("normalizedPath"))
	return rv
}/* debug [instance_properties/getter]: normalizedPath */


// The contour’s array of points in normalized coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour/normalizedPoints-2orqj
func (c_ Contour) NormalizedPoints() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("normalizedPoints"))
	return rv
}/* debug [instance_properties/getter]: normalizedPoints */


// The contour’s number of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNContour/pointCount
func (c_ Contour) PointCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("pointCount"))
	return rv
}/* debug [instance_properties/getter]: pointCount */


// The total number of detected contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/contourcount
func (c_ Contour) ContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("contourCount"))
	return rv
}/* debug [instance_properties/getter]: contourCount */


// The total number of detected contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/contourcount
func (c_ Contour) SetContourCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContourCount:"), value)
}/* debug [instance_properties/setter]: contourCount */


// The total number of detected top-level contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontourcount
func (c_ Contour) TopLevelContourCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("topLevelContourCount"))
	return rv
}/* debug [instance_properties/getter]: topLevelContourCount */


// The total number of detected top-level contours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontourcount
func (c_ Contour) SetTopLevelContourCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTopLevelContourCount:"), value)
}/* debug [instance_properties/setter]: topLevelContourCount */


// An array of contours that don’t have another contour enclosing them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontours
func (c_ Contour) TopLevelContours() IVNContour {
	rv := objc.Send[Contour](c_.ID, objc.Sel("topLevelContours"))
	return rv
}/* debug [instance_properties/getter]: topLevelContours */


// An array of contours that don’t have another contour enclosing them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncontoursobservation/toplevelcontours
func (c_ Contour) SetTopLevelContours(value IVNContour) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTopLevelContours:"), value)
}/* debug [instance_properties/setter]: topLevelContours */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNContour */



