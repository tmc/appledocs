// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Point] class.
var (
	PointClass     _PointClass
	PointClassOnce sync.Once
)

func getPointClass() _PointClass {
	PointClassOnce.Do(func() {
		PointClass = _PointClass{objc.GetClass("VNPoint")}
	})
	return PointClass
}

type _PointClass struct {
	class objc.Class
}

// An interface definition for the [Point] class.
type IPoint interface {
	objectivec.IObject
	DistanceToPoint(point IVNPoint) float64
	Location() coregraphics.CGPoint
	X() float64
	Y() float64
}

// An immutable object that represents a single 2D point in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint
type Point struct {
	objectivec.Object
}

// PointFrom constructs a [Point] from an unsafe.Pointer.
//
// An immutable object that represents a single 2D point in an image.
func PointFrom(ptr unsafe.Pointer) Point {
	return Point{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PointClass) Alloc() Point {
	rv := objc.Send[Point](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PointClass) New() Point {
	rv := objc.Send[Point](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Point) Init() Point {
	rv := objc.Send[Point](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Point) Autorelease() Point {
	rv := objc.Send[Point](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPoint creates a new Point instance.
func NewPoint() Point {
	return getPointClass().New()
}




// Creates a point object from the specified Core Graphics point.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/init(location:)
func NewPointWithLocation(location coregraphics.CGPoint) Point {
	instance := getPointClass().Alloc()
	rv := objc.Send[Point](instance.ID, objc.Sel("initWithLocation:"), location)
	rv.Autorelease()
	return rv
}



// Creates a point object with the specified coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/init(x:y:)
func NewPointWithXY(x float64, y float64) Point {
	instance := getPointClass().Alloc()
	rv := objc.Send[Point](instance.ID, objc.Sel("initWithX:y:"), x, y)
	rv.Autorelease()
	return rv
}


// Creates a point object that’s shifted by the X and Y offsets of the specified vector.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/apply(_:to:)
func (pc _PointClass) PointByApplyingVectorToPoint(vector IVNVector, point IVNPoint) Point {
	rv := objc.Send[Point](objc.ID(pc.class), objc.Sel("pointByApplyingVector:toPoint:"), vector, point)
	return rv
}

// Calculates the distance between two points.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/distance(_:_:)
func (pc _PointClass) DistanceBetweenPointPoint(point1 IVNPoint, point2 IVNPoint) float64 {
	rv := objc.Send[float64](objc.ID(pc.class), objc.Sel("distanceBetweenPoint:point:"), point1, point2)
	return rv
}

// A point object that represents the origin.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/zero
func (pc _PointClass) ZeroPoint() Point {
	rv := objc.Send[VNPoint](objc.ID(pc.class), objc.Sel("zeroPoint"))
	return rv
}
// Returns the distance to another point.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/distance(_:)
func (p_ Point) DistanceToPoint(point IVNPoint) float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("distanceToPoint:"), point)
	return rv
}

// The Core Graphics point for this point.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/location
func (p_ Point) Location() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](p_.ID, objc.Sel("location"))
	return rv
}

// The x-coordinate.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/x
func (p_ Point) X() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("x"))
	return rv
}

// The y-coordinate.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/y
func (p_ Point) Y() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("y"))
	return rv
}

// A point object that represents the origin.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint/zero
func (p_ Point) ZeroPoint() VNPoint {
	rv := objc.Send[VNPoint](p_.ID, objc.Sel("zeroPoint"))
	return rv
}


