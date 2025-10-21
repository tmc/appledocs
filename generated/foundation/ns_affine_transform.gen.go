// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AffineTransform] class.
type IAffineTransform interface {
	objectivec.IObject
	TransformPoint(aPoint Point) Point
	TranslateXByYBy(deltaX float64, deltaY float64)
}

// A graphics coordinate transformation.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. A transformation specifies how points in one coordinate system are transformed to points in another coordinate system. An affine transformation is a special type of transformation that preserves parallel lines in a path but does not necessarily preserve lengths or angles. Scaling, rotation, and translation are the most commonly used manipulations supported by affine transforms, but shearing is also possible. Methods for applying affine transformations to the current graphics context and a method for applying an affine transformation to an object are described in NSAffineTransform Additions Reference in the Application Kit.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AffineTransformClass) Alloc() AffineTransform {
	rv := objc.Send[AffineTransform](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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

// Applies the receiver’s transform to the specified point and returns the result.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transform(_:)-41p16
func (a_ AffineTransform) TransformPoint(aPoint Point) Point {
	rv := objc.Send[Point](a_.ID, objc.Sel("transformPoint:"), aPoint)
	return rv
}

// Applies the specified translation factors to the receiver’s transformation matrix.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform/translateX(by:yBy:)
func (a_ AffineTransform) TranslateXByYBy(deltaX float64, deltaY float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("translateXBy:yBy:"), deltaX, deltaY)
}
