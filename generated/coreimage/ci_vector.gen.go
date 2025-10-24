// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIVector */


/* debug [class_header]: Header for CIVector */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Vector */
// An interface definition for the [Vector] class.
type IVector interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Vector */
	// properties:
	CGAffineTransformValue() corefoundation.CGAffineTransform
	CGPointValue() corefoundation.CGPoint
	CGRectValue() corefoundation.CGRect
	Count() uintptr /* not a class type */
	StringRepresentation() objc.IObject /* cross-framework: NSString */
	W() float64
	X() float64
	Y() float64
	Z() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Vector */
	// methods:
	ValueAtIndex(index uintptr /* not a class type */) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Vector */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Vector */
// The Core Image class that defines a vector object.
//
// A can store one or more in one object. They can store a group of float values for a variety of different uses such as coordinate points, direction vectors, geometric rectangles, transform matrices, convolution weights, or just a list a parameter values. You use objects in conjunction with other Core Image classes, such as and . Many of the built-in Core Image filters have one or more inputs that you can set to affect the filter’s behavior.


// The Core Image class that defines a vector object.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Vector */

// Initialize a Core Image vector object with six values provided by a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgAffineTransform:)
func NewVectorWithCGAffineTransform(t corefoundation.CGAffineTransform) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithCGAffineTransform:"), t)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVectorWithCGAffineTransform */


// Initialize a Core Image vector object with two values provided by a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgPoint:)
func NewVectorWithCGPoint(p corefoundation.CGPoint) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithCGPoint:"), p)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVectorWithCGPoint */


// Initialize a Core Image vector object with four values provided by a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(cgRect:)
func NewVectorWithCGRect(r corefoundation.CGRect) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithCGRect:"), r)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVectorWithCGRect */


// Initialize a Core Image vector object with values provided in a string representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(string:)
func NewVectorWithString(representation objc.IObject /* cross-framework: NSString */) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithString:"), representation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVectorWithString */


// Initialize a Core Image vector object with the specified the values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(values:count:)
func NewVectorWithValuesCount(values corefoundation.CGFloat, count uintptr /* not a class type */) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithValues:count:"), values, count)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVectorWithValuesCount */


// Initialize a Core Image vector object with one value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:)
func NewVectorWithX(x float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:"), x)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVectorWithX */


// Initialize a Core Image vector object with two values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:)
func NewVectorWithXY(x float64, y float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:Y:"), x, y)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVectorWithXY */


// Initialize a Core Image vector object with three values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:)
func NewVectorWithXYZ(x float64, y float64, z float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:Y:Z:"), x, y, z)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVectorWithXYZ */


// Initialize a Core Image vector object with four values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/init(x:y:z:w:)
func NewVectorWithXYZW(x float64, y float64, z float64, w float64) Vector {
	instance := getVectorClass().Alloc()
	rv := objc.Send[Vector](instance.ID, objc.Sel("initWithX:Y:Z:W:"), x, y, z, w)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVectorWithXYZW */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Vector */

// Create a Core Image vector object that is initialized with six values provided by a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGAffineTransform:
func (vc _VectorClass) VectorWithCGAffineTransform(t corefoundation.CGAffineTransform) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorWithCGAffineTransform:"), t)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorWithCGAffineTransform) */


// Create a Core Image vector object that is initialized with two values provided by a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGPoint:
func (vc _VectorClass) VectorWithCGPoint(p corefoundation.CGPoint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorWithCGPoint:"), p)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorWithCGPoint) */


// Create a Core Image vector object that is initialized with four values provided by a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithCGRect:
func (vc _VectorClass) VectorWithCGRect(r corefoundation.CGRect) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorWithCGRect:"), r)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorWithCGRect) */


// Create a Core Image vector object with values provided in a string representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithString:
func (vc _VectorClass) VectorWithString(representation objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorWithString:"), representation)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorWithString) */


// Create a Core Image vector object that is initialized with the specified values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithValues:count:
func (vc _VectorClass) VectorWithValuesCount(values float64, count uintptr /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorWithValues:count:"), values, count)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorWithValuesCount) */


// Create a Core Image vector object that is initialized with one value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:
func (vc _VectorClass) VectorWithX(x float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorWithX:"), x)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorWithX) */


// Create a Core Image vector object that is initialized with two values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:
func (vc _VectorClass) VectorWithXY(x float64, y float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorWithX:Y:"), x, y)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorWithXY) */


// Create a Core Image vector object that is initialized with three values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:Z:
func (vc _VectorClass) VectorWithXYZ(x float64, y float64, z float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorWithX:Y:Z:"), x, y, z)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorWithXYZ) */


// Create a Core Image vector object that is initialized with four values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/vectorWithX:Y:Z:W:
func (vc _VectorClass) VectorWithXYZW(x float64, y float64, z float64, w float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("vectorWithX:Y:Z:W:"), x, y, z, w)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VectorWithXYZW) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Vector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Vector */

// Returns a value from a specific position in the vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/value(at:)
func (v_ Vector) ValueAtIndex(index uintptr /* not a class type */) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("valueAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ValueAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Vector */

// Returns the values in the vector as a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/cgAffineTransformValue
func (v_ Vector) CGAffineTransformValue() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](v_.ID, objc.Sel("CGAffineTransformValue"))
	return rv
}/* debug [instance_properties/getter]: CGAffineTransformValue */


// Returns the values in the vector as a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/cgPointValue
func (v_ Vector) CGPointValue() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v_.ID, objc.Sel("CGPointValue"))
	return rv
}/* debug [instance_properties/getter]: CGPointValue */


// Returns the values in the vector as a structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/cgRectValue
func (v_ Vector) CGRectValue() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v_.ID, objc.Sel("CGRectValue"))
	return rv
}/* debug [instance_properties/getter]: CGRectValue */


// The number of items in the vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/count
func (v_ Vector) Count() uintptr /* not a class type */ {
	rv := objc.Send[uintptr](v_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// Returns a formatted string with all the values of a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/stringRepresentation
func (v_ Vector) StringRepresentation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("stringRepresentation"))
	return rv
}/* debug [instance_properties/getter]: stringRepresentation */


// The value located in the forth position in the vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/w
func (v_ Vector) W() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("W"))
	return rv
}/* debug [instance_properties/getter]: W */


// The value located in the first position in the vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/x
func (v_ Vector) X() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("X"))
	return rv
}/* debug [instance_properties/getter]: X */


// The value located in the second position in the vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/y
func (v_ Vector) Y() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("Y"))
	return rv
}/* debug [instance_properties/getter]: Y */


// The value located in the third position in the vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIVector/z
func (v_ Vector) Z() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("Z"))
	return rv
}/* debug [instance_properties/getter]: Z */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIVector */


