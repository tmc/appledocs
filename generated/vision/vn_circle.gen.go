// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Circle] class.
var (
	CircleClass     _CircleClass
	CircleClassOnce sync.Once
)

func getCircleClass() _CircleClass {
	CircleClassOnce.Do(func() {
		CircleClass = _CircleClass{objc.GetClass("VNCircle")}
	})
	return CircleClass
}

type _CircleClass struct {
	class objc.Class
}





// An interface definition for the [Circle] class.
type ICircle interface {
	objectivec.IObject
	

	// properties:
	Center() IVNPoint
	Diameter() float64
	Radius() float64


	

	// methods:
	ContainsPoint(point IVNPoint) bool
	ContainsPointInCircumferentialRingOfWidth(point IVNPoint, ringWidth float64) bool


}





// Alloc allocates a new instance without initialization.
func (cc _CircleClass) Alloc() Circle {
	rv := objc.Send[Circle](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CircleClass) New() Circle {
	rv := objc.Send[Circle](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Circle) Init() Circle {
	rv := objc.Send[Circle](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Circle) Autorelease() Circle {
	rv := objc.Send[Circle](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCircle creates a new Circle instance.
func NewCircle() Circle {
	return getCircleClass().New()
}





// An immutable 2D circle represented by its center point and radius.


// An immutable 2D circle represented by its center point and radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle
type Circle struct {
	objectivec.Object
}

// CircleFrom constructs a [Circle] from an unsafe.Pointer.
//
// An immutable 2D circle represented by its center point and radius.
func CircleFrom(ptr unsafe.Pointer) Circle {
	return Circle{objectivec.Object{objc.ID(ptr)}}
}






// Creates a circle with the specified center and diameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle/init(center:diameter:)
func NewCircleWithCenterDiameter(center IVNPoint, diameter float64) Circle {
	instance := getCircleClass().Alloc()
	rv := objc.Send[Circle](instance.ID, objc.Sel("initWithCenter:diameter:"), center, diameter)
	rv.Autorelease()
	return rv
}


// Creates a circle with the specified center and radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle/init(center:radius:)
func NewCircleWithCenterRadius(center IVNPoint, radius float64) Circle {
	instance := getCircleClass().Alloc()
	rv := objc.Send[Circle](instance.ID, objc.Sel("initWithCenter:radius:"), center, radius)
	rv.Autorelease()
	return rv
}












// A circle object centered at the origin, with a radius of zero.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle/zero
func (cc _CircleClass) ZeroCircle() Circle {
	rv := objc.Send[Circle](objc.ID(cc.class), objc.Sel("zeroCircle"))
	return rv
}






// Determines if this circle, including its boundary, contains the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle/contains(_:)
func (c_ Circle) ContainsPoint(point IVNPoint) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsPoint:"), point)
	return rv
}


// Determines if a ring around this circle’s circumference contains the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle/contains(_:inCircumferentialRingOfWidth:)
func (c_ Circle) ContainsPointInCircumferentialRingOfWidth(point IVNPoint, ringWidth float64) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsPoint:inCircumferentialRingOfWidth:"), point, ringWidth)
	return rv
}







// The circle’s center point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle/center
func (c_ Circle) Center() IVNPoint {
	rv := objc.Send[Point](c_.ID, objc.Sel("center"))
	return rv
}


// The circle’s diameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle/diameter
func (c_ Circle) Diameter() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("diameter"))
	return rv
}


// The circle’s radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle/radius
func (c_ Circle) Radius() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("radius"))
	return rv
}


// A circle object centered at the origin, with a radius of zero.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCircle/zero
func (c_ Circle) ZeroCircle() IVNCircle {
	rv := objc.Send[Circle](c_.ID, objc.Sel("zeroCircle"))
	return rv
}







