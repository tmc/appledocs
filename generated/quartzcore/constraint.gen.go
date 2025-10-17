// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Constraint] class.
var constraintClass = _ConstraintClass{objc.GetClass("CAConstraint")}

type _ConstraintClass struct {
	class objc.Class
}

// A representation of a single layout constraint between two layers. [Full Topic]
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



