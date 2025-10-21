// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MultiArrayShapeConstraint] class.
var (
	MultiArrayShapeConstraintClass     _MultiArrayShapeConstraintClass
	MultiArrayShapeConstraintClassOnce sync.Once
)

func getMultiArrayShapeConstraintClass() _MultiArrayShapeConstraintClass {
	MultiArrayShapeConstraintClassOnce.Do(func() {
		MultiArrayShapeConstraintClass = _MultiArrayShapeConstraintClass{objc.GetClass("MLMultiArrayShapeConstraint")}
	})
	return MultiArrayShapeConstraintClass
}

type _MultiArrayShapeConstraintClass struct {
	class objc.Class
}

// An interface definition for the [MultiArrayShapeConstraint] class.
type IMultiArrayShapeConstraint interface {
	objectivec.IObject
}

// The lists of shapes or ranges of shapes that constrain a multiarray feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayShapeConstraint
type MultiArrayShapeConstraint struct {
	objectivec.Object
}

// MultiArrayShapeConstraintFrom constructs a [MultiArrayShapeConstraint] from an unsafe.Pointer.
//
// The lists of shapes or ranges of shapes that constrain a multiarray feature.
func MultiArrayShapeConstraintFrom(ptr unsafe.Pointer) MultiArrayShapeConstraint {
	return MultiArrayShapeConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MultiArrayShapeConstraintClass) Alloc() MultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MultiArrayShapeConstraintClass) New() MultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiArrayShapeConstraint) Init() MultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiArrayShapeConstraint) Autorelease() MultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiArrayShapeConstraint creates a new MultiArrayShapeConstraint instance.
func NewMultiArrayShapeConstraint() MultiArrayShapeConstraint {
	return getMultiArrayShapeConstraintClass().New()
}




