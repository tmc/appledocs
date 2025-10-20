// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Constraint] class.
var (
	constraintClass     _ConstraintClass
	constraintClassOnce sync.Once
)

func getConstraintClass() _ConstraintClass {
	constraintClassOnce.Do(func() {
		constraintClass = _ConstraintClass{objc.GetClass("CAConstraint")}
	})
	return constraintClass
}

type _ConstraintClass struct {
	class objc.Class
}

// An interface definition for the [Constraint] class.
type IConstraint interface {
	objectivec.IObject
}

// A representation of a single layout constraint between two layers.
//
// Each instance encapsulates one geometry relationship between two layers on the same axis. Sibling layers are referenced by name, using the name property of each layer. The special name is used to refer to the layer’s superlayer. For example, to specify that a layer should be horizontally centered in its superview you would use the following: A minimum of two relationships must be specified per axis. If you specify constraints for the left and right edges of a layer, the width will vary. If you specify constraints for the left edge and the width, the right edge of the layer will move relative to the superlayer’s frame. Often you’ll specify only a single edge constraint, the layer’s size in the same axis will be used as the second relationship.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraint
type Constraint struct {
	objectivec.Object
}

// ConstraintFrom constructs a [Constraint] from an unsafe.Pointer.
//
// A representation of a single layout constraint between two layers.
func ConstraintFrom(ptr unsafe.Pointer) Constraint {
	return Constraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ConstraintClass) Alloc() Constraint {
	rv := objc.Send[Constraint](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConstraintClass) New() Constraint {
	rv := objc.Send[Constraint](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Constraint) Init() Constraint {
	rv := objc.Send[Constraint](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Constraint) Autorelease() Constraint {
	rv := objc.Send[Constraint](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConstraint creates a new Constraint instance.
func NewConstraint() Constraint {
	return getConstraintClass().New()
}


// The constraint attribute of the layer the receiver is calculated relative to
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraint/sourceAttribute
func (c_ Constraint) SourceAttribute() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sourceAttribute"))
	return rv
}


