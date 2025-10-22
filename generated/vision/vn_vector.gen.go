// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Vector] class.
var (
	VectorClass     _VectorClass
	VectorClassOnce sync.Once
)

func getVectorClass() _VectorClass {
	VectorClassOnce.Do(func() {
		VectorClass = _VectorClass{objc.GetClass("VNVector")}
	})
	return VectorClass
}

type _VectorClass struct {
	class objc.Class
}

// An interface definition for the [Vector] class.
type IVector interface {
	objectivec.IObject
	Length() float64
	SetLength(value float64)
	R() float64
	SetR(value float64)
	SquaredLength() float64
	SetSquaredLength(value float64)
	Theta() float64
	SetTheta(value float64)
	X() float64
	SetX(value float64)
	Y() float64
	SetY(value float64)
}

// An immutable 2D vector represented by its x-axis and y-axis projections.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector
type Vector struct {
	objectivec.Object
}

// VectorFrom constructs a [Vector] from an unsafe.Pointer.
//
// An immutable 2D vector represented by its x-axis and y-axis projections.
func VectorFrom(ptr unsafe.Pointer) Vector {
	return Vector{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VectorClass) Alloc() Vector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VectorClass) New() Vector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ Vector) Init() Vector {
	rv := objc.Send[Vector](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ Vector) Autorelease() Vector {
	rv := objc.Send[Vector](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVector creates a new Vector instance.
func NewVector() Vector {
	return getVectorClass().New()
}


// The length, or absolute value, of the vector.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/length
func (v_ Vector) Length() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("length"))
	return rv
}


// SetLength sets the value of the length property.
// The length, or absolute value, of the vector.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/length
func (v_ Vector) SetLength(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLength:"), value)
}

// The radius, absolute value, or length of the vector.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/r
func (v_ Vector) R() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("r"))
	return rv
}


// SetR sets the value of the r property.
// The radius, absolute value, or length of the vector.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/r
func (v_ Vector) SetR(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setR:"), value)
}

// The squared length of the vector.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/squaredlength
func (v_ Vector) SquaredLength() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("squaredLength"))
	return rv
}


// SetSquaredLength sets the value of the squaredLength property.
// The squared length of the vector.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/squaredlength
func (v_ Vector) SetSquaredLength(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSquaredLength:"), value)
}

// The angle between the vector direction and the positive direction of the x-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/theta
func (v_ Vector) Theta() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("theta"))
	return rv
}


// SetTheta sets the value of the theta property.
// The angle between the vector direction and the positive direction of the x-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/theta
func (v_ Vector) SetTheta(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTheta:"), value)
}

// A signed projection that indicates the vector’s direction on the x-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/x
func (v_ Vector) X() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("x"))
	return rv
}


// SetX sets the value of the x property.
// A signed projection that indicates the vector’s direction on the x-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/x
func (v_ Vector) SetX(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setX:"), value)
}

// A signed projection that indicates the vector’s direction on the y-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/y
func (v_ Vector) Y() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("y"))
	return rv
}


// SetY sets the value of the y property.
// A signed projection that indicates the vector’s direction on the y-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnvector/y
func (v_ Vector) SetY(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setY:"), value)
}



