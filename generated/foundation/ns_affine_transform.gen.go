// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAffineTransform */


/* debug [class_header]: Header for NSAffineTransform */
// The class instance for the [AffineTransform] class.
var (
	AffineTransformClass     _AffineTransformClass
	AffineTransformClassOnce sync.Once
)

func getAffineTransformClass() _AffineTransformClass {
	AffineTransformClassOnce.Do(func() {
		AffineTransformClass = _AffineTransformClass{objc.GetClass("NSAffineTransform")}
	})
	return AffineTransformClass
}

type _AffineTransformClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AffineTransform */
// An interface definition for the [AffineTransform] class.
type IAffineTransform interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AffineTransform */
	// properties:
	TransformStruct() objc.IObject /* cross-framework: AffineTransformStruct */
	SetTransformStruct(value objc.IObject /* cross-framework: AffineTransformStruct */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AffineTransform */
	// methods:
	AppendTransform(transform IAffineTransform)
	Concat()
	Invert()
	PrependTransform(transform IAffineTransform)
	RotateByDegrees(angle float64)
	RotateByRadians(angle float64)
	ScaleBy(scale float64)
	ScaleXByYBy(scaleX float64, scaleY float64)
	Set()
	TransformPoint(aPoint objectivec.IObject) objectivec.IObject
	TransformSize(aSize Size /* typedef */) Size /* typedef */
	TransformBezierPath(path objectivec.IObject) objectivec.IObject
	TranslateXByYBy(deltaX float64, deltaY float64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AffineTransform */
// Alloc allocates a new instance without initialization.
func (ac _AffineTransformClass) Alloc() AffineTransform {
	rv := objc.Send[AffineTransform](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AffineTransformClass) New() AffineTransform {
	rv := objc.Send[AffineTransform](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AffineTransform) Init() AffineTransform {
	rv := objc.Send[AffineTransform](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AffineTransform) Autorelease() AffineTransform {
	rv := objc.Send[AffineTransform](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAffineTransform creates a new AffineTransform instance.
func NewAffineTransform() AffineTransform {
	return getAffineTransformClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AffineTransform */
// A graphics coordinate transformation.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. A transformation specifies how points in one coordinate system are transformed to points in another coordinate system. An affine transformation is a special type of transformation that preserves parallel lines in a path but does not necessarily preserve lengths or angles. Scaling, rotation, and translation are the most commonly used manipulations supported by affine transforms, but shearing is also possible. Methods for applying affine transformations to the current graphics context and a method for applying an affine transformation to an object are described in NSAffineTransform Additions Reference in the Application Kit.


// A graphics coordinate transformation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform
type AffineTransform struct {
	objectivec.Object
}

// AffineTransformFrom constructs a [AffineTransform] from an unsafe.Pointer.
//
// A graphics coordinate transformation.
func AffineTransformFrom(ptr unsafe.Pointer) AffineTransform {
	return AffineTransform{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AffineTransform */

// Initializes the receiver’s matrix using another transform object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/init(transform:)
func NewAffineTransformWithTransform(transform IAffineTransform) AffineTransform {
	instance := getAffineTransformClass().Alloc()
	rv := objc.Send[AffineTransform](instance.ID, objc.Sel("initWithTransform:"), transform)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAffineTransformWithTransform */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AffineTransform */

// Creates a new affine transform initialized to the identity matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transform
func (ac _AffineTransformClass) Transform() IAffineTransform {
	rv := objc.Send[AffineTransform](objc.ID(ac.class), objc.Sel("transform"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Transform) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AffineTransform */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AffineTransform */

// Appends the specified matrix to the receiver’s matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/append(_:)
func (a_ AffineTransform) AppendTransform(transform IAffineTransform) {
	objc.Send[objc.ID](a_.ID, objc.Sel("appendTransform:"), transform)
}/* debug [instance_methods/method]: AppendTransform */


// Appends the receiver’s matrix to the current transformation matrix stored in the current graphics context, replacing the current transformation matrix with the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/concat()
func (a_ AffineTransform) Concat() {
	objc.Send[objc.ID](a_.ID, objc.Sel("concat"))
}/* debug [instance_methods/method]: Concat */


// Replaces the receiver’s matrix with its inverse matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/invert()
func (a_ AffineTransform) Invert() {
	objc.Send[objc.ID](a_.ID, objc.Sel("invert"))
}/* debug [instance_methods/method]: Invert */


// Prepends the specified matrix to the receiver’s matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/prepend(_:)
func (a_ AffineTransform) PrependTransform(transform IAffineTransform) {
	objc.Send[objc.ID](a_.ID, objc.Sel("prependTransform:"), transform)
}/* debug [instance_methods/method]: PrependTransform */


// Applies a rotation factor (measured in degrees) to the receiver’s transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/rotate(byDegrees:)
func (a_ AffineTransform) RotateByDegrees(angle float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("rotateByDegrees:"), angle)
}/* debug [instance_methods/method]: RotateByDegrees */


// Applies a rotation factor (measured in radians) to the receiver’s transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/rotate(byRadians:)
func (a_ AffineTransform) RotateByRadians(angle float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("rotateByRadians:"), angle)
}/* debug [instance_methods/method]: RotateByRadians */


// Applies the specified scaling factor along both x and y axes to the receiver’s transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/scale(by:)
func (a_ AffineTransform) ScaleBy(scale float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scaleBy:"), scale)
}/* debug [instance_methods/method]: ScaleBy */


// Applies scaling factors to each axis of the receiver’s transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/scaleX(by:yBy:)
func (a_ AffineTransform) ScaleXByYBy(scaleX float64, scaleY float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scaleXBy:yBy:"), scaleX, scaleY)
}/* debug [instance_methods/method]: ScaleXByYBy */


// Sets the current transformation matrix to the receiver’s transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/set()
func (a_ AffineTransform) Set() {
	objc.Send[objc.ID](a_.ID, objc.Sel("set"))
}/* debug [instance_methods/method]: Set */


// Applies the receiver’s transform to the specified point and returns the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transform(_:)-41p16
func (a_ AffineTransform) TransformPoint(aPoint objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("transformPoint:"), aPoint)
	return rv
}/* debug [instance_methods/method]: TransformPoint */


// Applies the receiver’s transform to the specified size and returns the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transform(_:)-5r6ol
func (a_ AffineTransform) TransformSize(aSize Size /* typedef */) Size /* typedef */ {
	rv := objc.Send[CGSize](a_.ID, objc.Sel("transformSize:"), aSize)
	return rv
}/* debug [instance_methods/method]: TransformSize */


// Creates and returns a new Bézier path object with each point in the given path transformed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transform(_:)-6z1xo
func (a_ AffineTransform) TransformBezierPath(path objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("transformBezierPath:"), path)
	return rv
}/* debug [instance_methods/method]: TransformBezierPath */


// Applies the specified translation factors to the receiver’s transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/translateX(by:yBy:)
func (a_ AffineTransform) TranslateXByYBy(deltaX float64, deltaY float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("translateXBy:yBy:"), deltaX, deltaY)
}/* debug [instance_methods/method]: TranslateXByYBy */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AffineTransform */

// The matrix coefficients stored as the transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transformStruct
func (a_ AffineTransform) TransformStruct() objc.IObject /* cross-framework: AffineTransformStruct */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("transformStruct"))
	return rv
}/* debug [instance_properties/getter]: transformStruct */


// The matrix coefficients stored as the transformation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transformStruct
func (a_ AffineTransform) SetTransformStruct(value objc.IObject /* cross-framework: AffineTransformStruct */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransformStruct:"), value)
}/* debug [instance_properties/setter]: transformStruct */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAffineTransform */


