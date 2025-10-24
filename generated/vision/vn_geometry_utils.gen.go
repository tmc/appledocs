// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [GeometryUtils] class.
var (
	GeometryUtilsClass     _GeometryUtilsClass
	GeometryUtilsClassOnce sync.Once
)

func getGeometryUtilsClass() _GeometryUtilsClass {
	GeometryUtilsClassOnce.Do(func() {
		GeometryUtilsClass = _GeometryUtilsClass{objc.GetClass("VNGeometryUtils")}
	})
	return GeometryUtilsClass
}

type _GeometryUtilsClass struct {
	class objc.Class
}





// An interface definition for the [GeometryUtils] class.
type IGeometryUtils interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GeometryUtilsClass) Alloc() GeometryUtils {
	rv := objc.Send[GeometryUtils](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GeometryUtilsClass) New() GeometryUtils {
	rv := objc.Send[GeometryUtils](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GeometryUtils) Init() GeometryUtils {
	rv := objc.Send[GeometryUtils](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GeometryUtils) Autorelease() GeometryUtils {
	rv := objc.Send[GeometryUtils](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGeometryUtils creates a new GeometryUtils instance.
func NewGeometryUtils() GeometryUtils {
	return getGeometryUtilsClass().New()
}





// Utility methods to determine the geometries of various Vision types.


// Utility methods to determine the geometries of various Vision types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeometryUtils
type GeometryUtils struct {
	objectivec.Object
}

// GeometryUtilsFrom constructs a [GeometryUtils] from an unsafe.Pointer.
//
// Utility methods to determine the geometries of various Vision types.
func GeometryUtilsFrom(ptr unsafe.Pointer) GeometryUtils {
	return GeometryUtils{objectivec.Object{objc.ID(ptr)}}
}










// Calculates a bounding circle for the specified contour object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeometryUtils/boundingCircle(for:)-423ll
func (gc _GeometryUtilsClass) BoundingCircleForContourError(contour IVNContour, error_ objectivec.IObject) ICircle {
	rv := objc.Send[Circle](objc.ID(gc.class), objc.Sel("boundingCircleForContour:error:"), contour, error_)
	return rv
}


// Calculates a bounding circle for the specified array of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeometryUtils/boundingCircle(for:)-9dggv
func (gc _GeometryUtilsClass) BoundingCircleForPointsError(points []Point, error_ objectivec.IObject) ICircle {
	rv := objc.Send[Circle](objc.ID(gc.class), objc.Sel("boundingCircleForPoints:error:"), points, error_)
	return rv
}


// Calculates a bounding circle for the specified points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeometryUtils/boundingCircle(forSIMDPoints:pointCount:)
func (gc _GeometryUtilsClass) BoundingCircleForSIMDPointsPointCountError(points objectivec.IObject, pointCount int, error_ objectivec.IObject) ICircle {
	rv := objc.Send[Circle](objc.ID(gc.class), objc.Sel("boundingCircleForSIMDPoints:pointCount:error:"), points, pointCount, error_)
	return rv
}


// Calculates the area for the specified contour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeometryUtils/calculateArea(_:for:orientedArea:)
func (gc _GeometryUtilsClass) CalculateAreaForContourOrientedAreaError(area objectivec.IObject, contour IVNContour, orientedArea bool, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(gc.class), objc.Sel("calculateArea:forContour:orientedArea:error:"), area, contour, orientedArea, error_)
	return rv
}


// Calculates the perimeter of a closed contour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeometryUtils/calculatePerimeter(_:for:)
func (gc _GeometryUtilsClass) CalculatePerimeterForContourError(perimeter objectivec.IObject, contour IVNContour, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(gc.class), objc.Sel("calculatePerimeter:forContour:error:"), perimeter, contour, error_)
	return rv
}























