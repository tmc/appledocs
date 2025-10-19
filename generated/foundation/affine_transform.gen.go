// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AffineTransform] class.
var affineTransformClass = _AffineTransformClass{objc.GetClass("NSAffineTransform")}

type _AffineTransformClass struct {
	class objc.Class
}

// An interface definition for the [AffineTransform] class.
type IAffineTransform interface {
	objectivec.IObject
}

// A graphics coordinate transformation. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return affineTransformClass.New()
}




