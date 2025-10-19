// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Vector] class.
var vectorClass = _VectorClass{objc.GetClass("CIVector")}

type _VectorClass struct {
	class objc.Class
}

// An interface definition for the [Vector] class.
type IVector interface {
	objectivec.IObject
	ValueAtIndex(index uintptr) float64
}

// The Core Image class that defines a vector object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector

type Vector struct {
	objectivec.Object
}

// VectorFrom constructs a [Vector] from an unsafe.Pointer.
//
// The Core Image class that defines a vector object.
func VectorFrom(ptr unsafe.Pointer) Vector {
	return Vector{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (vc _VectorClass) Alloc() Vector {
	rv := objc.Send[Vector](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return vectorClass.New()
}


// Initialize a Core Image vector object with two values provided by a structure. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgPoint:)
func NewVectorWithCGPoint(p unsafe.Pointer) Vector {
	instance := vectorClass.Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithCGPoint:"), p)
	rv.Autorelease()
	return rv
}
// Initialize a Core Image vector object with four values provided by a structure. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgRect:)
func NewVectorWithCGRect(r unsafe.Pointer) Vector {
	instance := vectorClass.Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithCGRect:"), r)
	rv.Autorelease()
	return rv
}
// Initialize a Core Image vector object with six values provided by a structure. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgAffineTransform:)
func NewVectorWithCGAffineTransform(t coregraphics.AffineTransform) Vector {
	instance := vectorClass.Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithCGAffineTransform:"), t)
	rv.Autorelease()
	return rv
}
// Initialize a Core Image vector object with values provided in a string representation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(string:)
func NewVectorWithString(representation string) Vector {
	instance := vectorClass.Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithString:"), representation)
	rv.Autorelease()
	return rv
}
// Initialize a Core Image vector object with the specified the values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(values:count:)
func NewVectorWithValuesCount(values unsafe.Pointer, count uintptr) Vector {
	instance := vectorClass.Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithValues:count:"), values, count)
	rv.Autorelease()
	return rv
}
// Initialize a Core Image vector object with one value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:)
func NewVectorWithX(x float64) Vector {
	instance := vectorClass.Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:"), x)
	rv.Autorelease()
	return rv
}
// Initialize a Core Image vector object with two values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:)
func NewVectorWithXY(x float64, y float64) Vector {
	instance := vectorClass.Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:Y:"), x, y)
	rv.Autorelease()
	return rv
}
// Initialize a Core Image vector object with three values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:)
func NewVectorWithXYZ(x float64, y float64, z float64) Vector {
	instance := vectorClass.Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:Y:Z:"), x, y, z)
	rv.Autorelease()
	return rv
}
// Initialize a Core Image vector object with four values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:w:)
func NewVectorWithXYZW(x float64, y float64, z float64, w float64) Vector {
	instance := vectorClass.Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:Y:Z:W:"), x, y, z, w)
	rv.Autorelease()
	return rv
}


// Create a Core Image vector object that is initialized with six values provided by a structure. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGAffineTransform:
func (vc _VectorClass) VectorWithCGAffineTransform(t coregraphics.AffineTransform) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithCGAffineTransform:"), t)
	return rv
}
// Create a Core Image vector object that is initialized with two values provided by a structure. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGPoint:
func (vc _VectorClass) VectorWithCGPoint(p unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithCGPoint:"), p)
	return rv
}
// Create a Core Image vector object that is initialized with four values provided by a structure. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGRect:
func (vc _VectorClass) VectorWithCGRect(r unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithCGRect:"), r)
	return rv
}
// Create a Core Image vector object with values provided in a string representation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithString:
func (vc _VectorClass) VectorWithString(representation string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithString:"), representation)
	return rv
}
// Create a Core Image vector object that is initialized with the specified values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithValues:count:
func (vc _VectorClass) VectorWithValuesCount(values unsafe.Pointer, count uintptr) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithValues:count:"), values, count)
	return rv
}
// Create a Core Image vector object that is initialized with one value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:
func (vc _VectorClass) VectorWithX(x float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithX:"), x)
	return rv
}
// Create a Core Image vector object that is initialized with two values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:
func (vc _VectorClass) VectorWithXY(x float64, y float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithX:Y:"), x, y)
	return rv
}
// Create a Core Image vector object that is initialized with three values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:Z:
func (vc _VectorClass) VectorWithXYZ(x float64, y float64, z float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithX:Y:Z:"), x, y, z)
	return rv
}
// Create a Core Image vector object that is initialized with four values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:Z:W:
func (vc _VectorClass) VectorWithXYZW(x float64, y float64, z float64, w float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithX:Y:Z:W:"), x, y, z, w)
	return rv
}
// Returns a value from a specific position in the vector. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/value(at:)
func (v_ Vector) ValueAtIndex(index uintptr) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("valueAtIndex:"), index)
	return rv
}

