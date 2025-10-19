// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ConstraintConflict] class.
var (
	constraintConflictClass     _ConstraintConflictClass
	constraintConflictClassOnce sync.Once
)

func getConstraintConflictClass() _ConstraintConflictClass {
	constraintConflictClassOnce.Do(func() {
		constraintConflictClass = _ConstraintConflictClass{objc.GetClass("NSConstraintConflict")}
	})
	return constraintConflictClass
}

type _ConstraintConflictClass struct {
	class objc.Class
}

// An interface definition for the [ConstraintConflict] class.
type IConstraintConflict interface {
	objectivec.IObject
}

// An encapsulation of conflicts that occur during an attempt to save a managed object.
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

// Alloc allocates a new instance without initialization.
func (cc _ConstraintConflictClass) Alloc() ConstraintConflict {
	rv := objc.Send[ConstraintConflict](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConstraintConflictClass) New() ConstraintConflict {
	rv := objc.Send[ConstraintConflict](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConstraintConflict) Init() ConstraintConflict {
	rv := objc.Send[ConstraintConflict](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConstraintConflict) Autorelease() ConstraintConflict {
	rv := objc.Send[ConstraintConflict](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConstraintConflict creates a new ConstraintConflict instance.
func NewConstraintConflict() ConstraintConflict {
	return getConstraintConflictClass().New()
}




