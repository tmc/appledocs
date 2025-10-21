// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SequenceConstraint] class.
var (
	SequenceConstraintClass     _SequenceConstraintClass
	SequenceConstraintClassOnce sync.Once
)

func getSequenceConstraintClass() _SequenceConstraintClass {
	SequenceConstraintClassOnce.Do(func() {
		SequenceConstraintClass = _SequenceConstraintClass{objc.GetClass("MLSequenceConstraint")}
	})
	return SequenceConstraintClass
}

type _SequenceConstraintClass struct {
	class objc.Class
}

// An interface definition for the [SequenceConstraint] class.
type ISequenceConstraint interface {
	objectivec.IObject
}

// The constraints for a sequence feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint
type SequenceConstraint struct {
	objectivec.Object
}

// SequenceConstraintFrom constructs a [SequenceConstraint] from an unsafe.Pointer.
//
// The constraints for a sequence feature.
func SequenceConstraintFrom(ptr unsafe.Pointer) SequenceConstraint {
	return SequenceConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SequenceConstraintClass) Alloc() SequenceConstraint {
	rv := objc.Send[SequenceConstraint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SequenceConstraintClass) New() SequenceConstraint {
	rv := objc.Send[SequenceConstraint](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SequenceConstraint) Init() SequenceConstraint {
	rv := objc.Send[SequenceConstraint](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SequenceConstraint) Autorelease() SequenceConstraint {
	rv := objc.Send[SequenceConstraint](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSequenceConstraint creates a new SequenceConstraint instance.
func NewSequenceConstraint() SequenceConstraint {
	return getSequenceConstraintClass().New()
}




