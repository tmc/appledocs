// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ConstraintConflict] class.
var constraintConflictClass = _ConstraintConflictClass{objc.GetClass("NSConstraintConflict")}

type _ConstraintConflictClass struct {
	class objc.Class
}

// An encapsulation of conflicts that occur during an attempt to save a managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict

type ConstraintConflict struct {
	objectivec.Object
}

// ConstraintConflictFrom constructs a [ConstraintConflict] from an unsafe.Pointer.
//
// An encapsulation of conflicts that occur during an attempt to save a managed object.
func ConstraintConflictFrom(ptr unsafe.Pointer) ConstraintConflict {
	return ConstraintConflict{objectivec.Object{objc.ID(ptr)}}
}



