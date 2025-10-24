// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Location() IPoint
	SetLocation(value IPoint)
	X() float64
	SetX(value float64)
	Y() float64
	SetY(value float64)
	// methods:
}

// An immutable object that represents a single 2D point in an image.


// An immutable object that represents a single 2D point in an image.
//
// [Full Topic]
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



// The Core Graphics point for this point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/location
func (p_ Point) Location() IPoint {
	rv := objc.Send[Point](p_.ID, objc.Sel("location"))
	return rv
}


// The Core Graphics point for this point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/location
func (p_ Point) SetLocation(value IPoint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocation:"), value)
}


// The x-coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/x
func (p_ Point) X() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("x"))
	return rv
}


// The x-coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/x
func (p_ Point) SetX(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setX:"), value)
}


// The y-coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/y
func (p_ Point) Y() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("y"))
	return rv
}


// The y-coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnpoint/y
func (p_ Point) SetY(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setY:"), value)
}



