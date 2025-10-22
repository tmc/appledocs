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
	Center() VNPoint
	SetCenter(value IVNPoint)
	Diameter() float64
	SetDiameter(value float64)
	Radius() float64
	SetRadius(value float64)
}

// An immutable 2D circle represented by its center point and radius.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CircleClass) Alloc() Circle {
	rv := objc.Send[Circle](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The circle’s center point.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/center
func (c_ Circle) Center() VNPoint {
	rv := objc.Send[VNPoint](c_.ID, objc.Sel("center"))
	return rv
}


// SetCenter sets the value of the center property.
// The circle’s center point.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/center
func (c_ Circle) SetCenter(value IVNPoint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCenter:"), value)
}

// The circle’s diameter.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/diameter
func (c_ Circle) Diameter() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("diameter"))
	return rv
}


// SetDiameter sets the value of the diameter property.
// The circle’s diameter.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/diameter
func (c_ Circle) SetDiameter(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDiameter:"), value)
}

// The circle’s radius.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/radius
func (c_ Circle) Radius() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("radius"))
	return rv
}


// SetRadius sets the value of the radius property.
// The circle’s radius.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncircle/radius
func (c_ Circle) SetRadius(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRadius:"), value)
}



