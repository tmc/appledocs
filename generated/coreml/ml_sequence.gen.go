// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Sequence] class.
var (
	SequenceClass     _SequenceClass
	SequenceClassOnce sync.Once
)

func getSequenceClass() _SequenceClass {
	SequenceClassOnce.Do(func() {
		SequenceClass = _SequenceClass{objc.GetClass("MLSequence")}
	})
	return SequenceClass
}

type _SequenceClass struct {
	class objc.Class
}

// An interface definition for the [Sequence] class.
type ISequence interface {
	objectivec.IObject
}

// A machine learning collection type that stores a series of strings or integers.
//
// A sequence stores a series of integers or strings of any length as the underlying type of an . Some classifier models — typically natural language models, such as an — produce an feature value from their output features.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence
type Sequence struct {
	objectivec.Object
}

// SequenceFrom constructs a [Sequence] from an unsafe.Pointer.
//
// A machine learning collection type that stores a series of strings or integers.
func SequenceFrom(ptr unsafe.Pointer) Sequence {
	return Sequence{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SequenceClass) Alloc() Sequence {
	rv := objc.Send[Sequence](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SequenceClass) New() Sequence {
	rv := objc.Send[Sequence](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Sequence) Init() Sequence {
	rv := objc.Send[Sequence](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Sequence) Autorelease() Sequence {
	rv := objc.Send[Sequence](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSequence creates a new Sequence instance.
func NewSequence() Sequence {
	return getSequenceClass().New()
}




