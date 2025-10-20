// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Vector] class.
var (
	VectorClass     _VectorClass
	VectorClassOnce sync.Once
)

func getVectorClass() _VectorClass {
	VectorClassOnce.Do(func() {
		VectorClass = _VectorClass{objc.GetClass("CIVector")}
	})
	return VectorClass
}

type _VectorClass struct {
	class objc.Class
}

// An interface definition for the [Vector] class.
type IVector interface {
	objectivec.IObject
	ValueAtIndex(index unsafe.Pointer) float64
}

// The Core Image class that defines a vector object.
//
// A can store one or more in one object. They can store a group of float values for a variety of different uses such as coordinate points, direction vectors, geometric rectangles, transform matrices, convolution weights, or just a list a parameter values. You use objects in conjunction with other Core Image classes, such as and . Many of the built-in Core Image filters have one or more inputs that you can set to affect the filter’s behavior.
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


// Initialize a Core Image vector object with four values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:w:)
func NewVectorWithXYZW(x float64, y float64, z float64, w float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:Y:Z:W:"), x, y, z, w)
	rv.Autorelease()
	return rv
}

// Initialize a Core Image vector object with two values provided by a structure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgPoint:)
func NewVectorWithCGPoint(p coregraphics.CGPoint) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithCGPoint:"), p)
	rv.Autorelease()
	return rv
}

// Initialize a Core Image vector object with values provided in a string representation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(string:)
func NewVectorWithString(representation string) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithString:"), objc.String(representation))
	rv.Autorelease()
	return rv
}

// Initialize a Core Image vector object with six values provided by a structure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgAffineTransform:)
func NewVectorWithCGAffineTransform(t coregraphics.CGAffineTransform) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithCGAffineTransform:"), t)
	rv.Autorelease()
	return rv
}

// Initialize a Core Image vector object with four values provided by a structure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgRect:)
func NewVectorWithCGRect(r coregraphics.CGRect) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithCGRect:"), r)
	rv.Autorelease()
	return rv
}

// Initialize a Core Image vector object with the specified the values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(values:count:)
func NewVectorWithValuesCount(values unsafe.Pointer, count unsafe.Pointer) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithValues:count:"), values, count)
	rv.Autorelease()
	return rv
}

// Initialize a Core Image vector object with one value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:)
func NewVectorWithX(x float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:"), x)
	rv.Autorelease()
	return rv
}

// Initialize a Core Image vector object with two values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:)
func NewVectorWithXY(x float64, y float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:Y:"), x, y)
	rv.Autorelease()
	return rv
}

// Initialize a Core Image vector object with three values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:)
func NewVectorWithXYZ(x float64, y float64, z float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:Y:Z:"), x, y, z)
	rv.Autorelease()
	return rv
}


// Create a Core Image vector object that is initialized with six values provided by a structure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGAffineTransform:
func (vc _VectorClass) VectorWithCGAffineTransform(t coregraphics.CGAffineTransform) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithCGAffineTransform:"), t)
	return rv
}

// Create a Core Image vector object that is initialized with two values provided by a structure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGPoint:
func (vc _VectorClass) VectorWithCGPoint(p coregraphics.CGPoint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithCGPoint:"), p)
	return rv
}

// Create a Core Image vector object that is initialized with four values provided by a structure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGRect:
func (vc _VectorClass) VectorWithCGRect(r coregraphics.CGRect) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithCGRect:"), r)
	return rv
}

// Create a Core Image vector object with values provided in a string representation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithString:
func (vc _VectorClass) VectorWithString(representation string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithString:"), objc.String(representation))
	return rv
}

// Create a Core Image vector object that is initialized with the specified values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithValues:count:
func (vc _VectorClass) VectorWithValuesCount(values unsafe.Pointer, count unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithValues:count:"), values, count)
	return rv
}

// Create a Core Image vector object that is initialized with one value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:
func (vc _VectorClass) VectorWithX(x float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithX:"), x)
	return rv
}

// Create a Core Image vector object that is initialized with two values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:
func (vc _VectorClass) VectorWithXY(x float64, y float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithX:Y:"), x, y)
	return rv
}

// Create a Core Image vector object that is initialized with three values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:Z:
func (vc _VectorClass) VectorWithXYZ(x float64, y float64, z float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithX:Y:Z:"), x, y, z)
	return rv
}

// Create a Core Image vector object that is initialized with four values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:Z:W:
func (vc _VectorClass) VectorWithXYZW(x float64, y float64, z float64, w float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("vectorWithX:Y:Z:W:"), x, y, z, w)
	return rv
}

// Returns a value from a specific position in the vector.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/value(at:)
func (v_ Vector) ValueAtIndex(index unsafe.Pointer) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("valueAtIndex:"), index)
	return rv
}

// Returns the values in the vector as a structure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/cgAffineTransformValue
func (v_ Vector) CGAffineTransformValue() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](v_.ID, objc.Sel("CGAffineTransformValue"))
	return rv
}

// Returns the values in the vector as a structure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/cgPointValue
func (v_ Vector) CGPointValue() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("CGPointValue"))
	return rv
}

// Returns the values in the vector as a structure.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/cgRectValue
func (v_ Vector) CGRectValue() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("CGRectValue"))
	return rv
}

// The number of items in the vector.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/count
func (v_ Vector) Count() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("count"))
	return rv
}

// Returns a formatted string with all the values of a .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/stringRepresentation
func (v_ Vector) StringRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("stringRepresentation"))
	return rv
}

// The value located in the forth position in the vector.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/w
func (v_ Vector) W() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("W"))
	return rv
}

// The value located in the first position in the vector.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/x
func (v_ Vector) X() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("X"))
	return rv
}

// The value located in the second position in the vector.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/y
func (v_ Vector) Y() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("Y"))
	return rv
}

// The value located in the third position in the vector.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/z
func (v_ Vector) Z() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("Z"))
	return rv
}


