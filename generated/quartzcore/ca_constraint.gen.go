// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Constraint] class.
var (
	ConstraintClass     _ConstraintClass
	ConstraintClassOnce sync.Once
)

func getConstraintClass() _ConstraintClass {
	ConstraintClassOnce.Do(func() {
		ConstraintClass = _ConstraintClass{objc.GetClass("CAConstraint")}
	})
	return ConstraintClass
}

type _ConstraintClass struct {
	class objc.Class
}

// An interface definition for the [Constraint] class.
type IConstraint interface {
	objectivec.IObject
	// properties:
	SourceAttribute() ConstraintAttribute /* not a class type */
	Attribute() ConstraintAttribute /* not a class type */
	SetAttribute(value ConstraintAttribute /* not a class type */)
	Offset() float64
	SetOffset(value float64)
	Scale() float64
	SetScale(value float64)
	SourceName() objc.IObject /* cross-framework: NSString */
	SetSourceName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A representation of a single layout constraint between two layers.
//
// Each instance encapsulates one geometry relationship between two layers on the same axis. Sibling layers are referenced by name, using the name property of each layer. The special name is used to refer to the layer’s superlayer. For example, to specify that a layer should be horizontally centered in its superview you would use the following: A minimum of two relationships must be specified per axis. If you specify constraints for the left and right edges of a layer, the width will vary. If you specify constraints for the left edge and the width, the right edge of the layer will move relative to the superlayer’s frame. Often you’ll specify only a single edge constraint, the layer’s size in the same axis will be used as the second relationship.


// A representation of a single layout constraint between two layers.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraint/sourceAttribute
func (c_ Constraint) SourceAttribute() ConstraintAttribute /* not a class type */ {
	rv := objc.Send[ConstraintAttribute](c_.ID, objc.Sel("sourceAttribute"))
	return rv
}


// The attribute the constraint affects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/attribute
func (c_ Constraint) Attribute() ConstraintAttribute /* not a class type */ {
	rv := objc.Send[ConstraintAttribute](c_.ID, objc.Sel("attribute"))
	return rv
}


// The attribute the constraint affects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/attribute
func (c_ Constraint) SetAttribute(value ConstraintAttribute /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAttribute:"), value)
}


// Offset value of the constraint attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/offset
func (c_ Constraint) Offset() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("offset"))
	return rv
}


// Offset value of the constraint attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/offset
func (c_ Constraint) SetOffset(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOffset:"), value)
}


// Scale factor of the constraint attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/scale
func (c_ Constraint) Scale() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("scale"))
	return rv
}


// Scale factor of the constraint attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/scale
func (c_ Constraint) SetScale(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScale:"), value)
}


// Name of the layer that the constraint is calculated relative to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/sourcename
func (c_ Constraint) SourceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("sourceName"))
	return rv
}


// Name of the layer that the constraint is calculated relative to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caconstraint/sourcename
func (c_ Constraint) SetSourceName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceName:"), value)
}



