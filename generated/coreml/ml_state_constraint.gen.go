// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StateConstraint] class.
var (
	StateConstraintClass     _StateConstraintClass
	StateConstraintClassOnce sync.Once
)

func getStateConstraintClass() _StateConstraintClass {
	StateConstraintClassOnce.Do(func() {
		StateConstraintClass = _StateConstraintClass{objc.GetClass("MLStateConstraint")}
	})
	return StateConstraintClass
}

type _StateConstraintClass struct {
	class objc.Class
}

// An interface definition for the [StateConstraint] class.
type IStateConstraint interface {
	objectivec.IObject
}

// Constraint of a state feature value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLStateConstraint
type StateConstraint struct {
	objectivec.Object
}

// StateConstraintFrom constructs a [StateConstraint] from an unsafe.Pointer.
//
// Constraint of a state feature value.
func StateConstraintFrom(ptr unsafe.Pointer) StateConstraint {
	return StateConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StateConstraintClass) Alloc() StateConstraint {
	rv := objc.Send[StateConstraint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StateConstraintClass) New() StateConstraint {
	rv := objc.Send[StateConstraint](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StateConstraint) Init() StateConstraint {
	rv := objc.Send[StateConstraint](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StateConstraint) Autorelease() StateConstraint {
	rv := objc.Send[StateConstraint](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStateConstraint creates a new StateConstraint instance.
func NewStateConstraint() StateConstraint {
	return getStateConstraintClass().New()
}


// The shape of the state buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLStateConstraint/bufferShape-6o5vn
func (s_ StateConstraint) BufferShape() []foundation.NSNumber {
	rv := objc.Send[[]foundation.NSNumber](s_.ID, objc.Sel("bufferShape"))
	return rv
}

// The data type of scalars in the state buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLStateConstraint/dataType
func (s_ StateConstraint) DataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dataType"))
	return rv
}



