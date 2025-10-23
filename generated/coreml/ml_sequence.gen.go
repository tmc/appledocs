// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Int64Values() foundation.objc.IObject /* cross-framework: Number */
	SetInt64Values(value foundation.objc.IObject /* cross-framework: Number */)
	StringValues() string /* primitive/slice/pointer. */
	SetStringValues(value string /* primitive/slice/pointer. */)
	Type() FeatureType
	SetType(value FeatureType)
	// methods:
}

// A machine learning collection type that stores a series of strings or integers.
//
// A sequence stores a series of integers or strings of any length as the underlying type of an . Some classifier models — typically natural language models, such as an — produce an feature value from their output features.


// A machine learning collection type that stores a series of strings or integers.
//
// [Full Topic]
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



// An array of 64-bit integers in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/int64values
func (s_ Sequence) Int64Values() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](s_.ID, objc.Sel("int64Values"))
	return rv
}


// An array of 64-bit integers in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/int64values
func (s_ Sequence) SetInt64Values(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInt64Values:"), value)
}


// An array of strings in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/stringvalues
func (s_ Sequence) StringValues() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](s_.ID, objc.Sel("stringValues"))
	return rv
}


// An array of strings in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/stringvalues
func (s_ Sequence) SetStringValues(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStringValues:"), objc.String(value))
}


// The underlying type of the sequence’s elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/type
func (s_ Sequence) Type() FeatureType {
	rv := objc.Send[FeatureType](s_.ID, objc.Sel("type"))
	return rv
}


// The underlying type of the sequence’s elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequence/type
func (s_ Sequence) SetType(value FeatureType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setType:"), value)
}



