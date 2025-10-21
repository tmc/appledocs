// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MultiArrayConstraint] class.
var (
	MultiArrayConstraintClass     _MultiArrayConstraintClass
	MultiArrayConstraintClassOnce sync.Once
)

func getMultiArrayConstraintClass() _MultiArrayConstraintClass {
	MultiArrayConstraintClassOnce.Do(func() {
		MultiArrayConstraintClass = _MultiArrayConstraintClass{objc.GetClass("MLMultiArrayConstraint")}
	})
	return MultiArrayConstraintClass
}

type _MultiArrayConstraintClass struct {
	class objc.Class
}

// An interface definition for the [MultiArrayConstraint] class.
type IMultiArrayConstraint interface {
	objectivec.IObject
}

// The shape and data type constraints for a multidimensional array feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint
type MultiArrayConstraint struct {
	objectivec.Object
}

// MultiArrayConstraintFrom constructs a [MultiArrayConstraint] from an unsafe.Pointer.
//
// The shape and data type constraints for a multidimensional array feature.
func MultiArrayConstraintFrom(ptr unsafe.Pointer) MultiArrayConstraint {
	return MultiArrayConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MultiArrayConstraintClass) Alloc() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MultiArrayConstraintClass) New() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiArrayConstraint) Init() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiArrayConstraint) Autorelease() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiArrayConstraint creates a new MultiArrayConstraint instance.
func NewMultiArrayConstraint() MultiArrayConstraint {
	return getMultiArrayConstraintClass().New()
}


// The type for the multi array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/dataType
func (m_ MultiArrayConstraint) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dataType"))
	return rv
}

// The shape of the multi array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/shape
func (m_ MultiArrayConstraint) Shape() []accessibility.NSNumber {
	rv := objc.Send[[]accessibility.NSNumber](m_.ID, objc.Sel("shape"))
	return rv
}

// The constraint on the shape of the multiarray.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/shapeConstraint
func (m_ MultiArrayConstraint) ShapeConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("shapeConstraint"))
	return rv
}



