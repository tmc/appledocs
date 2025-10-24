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
	

	// properties:
	Length() float64
	R() float64
	SquaredLength() float64
	Theta() float64
	X() float64
	Y() float64


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VectorClass) Alloc() Vector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An immutable 2D vector represented by its x-axis and y-axis projections.


// An immutable 2D vector represented by its x-axis and y-axis projections.
//
// [Full Topic]
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






// Creates a new vector by adding the specified vectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/init(byAdding:to:)
func NewVectorByAddingVectorToVector(v1 IVNVector, v2 IVNVector) Vector {
	rv := objc.Send[Vector](objc.ID(getVectorClass().class), objc.Sel("vectorByAddingVector:toVector:"), v1, v2)
	return rv
}


// Creates a new vector by multiplying the specified vector’s x-axis and y-axis projections by the scalar value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/init(byMultiplying:byScalar:)
func NewVectorByMultiplyingVectorByScalar(vector IVNVector, scalar float64) Vector {
	rv := objc.Send[Vector](objc.ID(getVectorClass().class), objc.Sel("vectorByMultiplyingVector:byScalar:"), vector, scalar)
	return rv
}


// Creates a new vector by subtracting the first vector from the second vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/init(bySubtracting:from:)
func NewVectorBySubtractingVectorFromVector(v1 IVNVector, v2 IVNVector) Vector {
	rv := objc.Send[Vector](objc.ID(getVectorClass().class), objc.Sel("vectorBySubtractingVector:fromVector:"), v1, v2)
	return rv
}


// Creates a new vector in polar coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/init(r:theta:)
func NewVectorWithRTheta(r float64, theta float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithR:theta:"), r, theta)
	rv.Autorelease()
	return rv
}


// Creates a new vector in Cartesian coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/init(vectorHead:tail:)
func NewVectorWithVectorHeadTail(head IVNPoint, tail IVNPoint) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithVectorHead:tail:"), head, tail)
	rv.Autorelease()
	return rv
}


// Creates a new vector in Cartesian coordinate space, based on its x-axis and y-axis projections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/init(xComponent:yComponent:)
func NewVectorWithXComponentYComponent(x float64, y float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithXComponent:yComponent:"), x, y)
	rv.Autorelease()
	return rv
}







// Caclulates the dot product of two vectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/dotProduct(of:vector:)
func (vc _VectorClass) DotProductOfVectorVector(v1 IVNVector, v2 IVNVector) float64 {
	rv := objc.Send[float64](objc.ID(vc.class), objc.Sel("dotProductOfVector:vector:"), v1, v2)
	return rv
}


// Creates a new vector by adding the specified vectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/init(byAdding:to:)
func (vc _VectorClass) VectorByAddingVectorToVector(v1 IVNVector, v2 IVNVector) IVector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("vectorByAddingVector:toVector:"), v1, v2)
	return rv
}


// Creates a new vector by multiplying the specified vector’s x-axis and y-axis projections by the scalar value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/init(byMultiplying:byScalar:)
func (vc _VectorClass) VectorByMultiplyingVectorByScalar(vector IVNVector, scalar float64) IVector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("vectorByMultiplyingVector:byScalar:"), vector, scalar)
	return rv
}


// Creates a new vector by subtracting the first vector from the second vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/init(bySubtracting:from:)
func (vc _VectorClass) VectorBySubtractingVectorFromVector(v1 IVNVector, v2 IVNVector) IVector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("vectorBySubtractingVector:fromVector:"), v1, v2)
	return rv
}


// Calculates a vector that’s normalized by preserving its direction, so that the vector length equals 1.0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/unitVector(for:)
func (vc _VectorClass) UnitVectorForVector(vector IVNVector) IVector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("unitVectorForVector:"), vector)
	return rv
}







// A vector object with zero length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/zero
func (vc _VectorClass) ZeroVector() Vector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("zeroVector"))
	return rv
}











// The length, or absolute value, of the vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/length
func (v_ Vector) Length() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("length"))
	return rv
}


// The radius, absolute value, or length of the vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/r
func (v_ Vector) R() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("r"))
	return rv
}


// The squared length of the vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/squaredLength
func (v_ Vector) SquaredLength() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("squaredLength"))
	return rv
}


// The angle between the vector direction and the positive direction of the x-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/theta
func (v_ Vector) Theta() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("theta"))
	return rv
}


// A signed projection that indicates the vector’s direction on the x-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/x
func (v_ Vector) X() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("x"))
	return rv
}


// A signed projection that indicates the vector’s direction on the y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/y
func (v_ Vector) Y() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("y"))
	return rv
}


// A vector object with zero length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNVector/zero
func (v_ Vector) ZeroVector() IVNVector {
	rv := objc.Send[Vector](v_.ID, objc.Sel("zeroVector"))
	return rv
}







