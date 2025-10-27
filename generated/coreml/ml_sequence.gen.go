// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Int64Values() []foundation.Number
	StringValues() []string
	Type() FeatureType


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SequenceClass) Alloc() Sequence {
	rv := objc.Send[Sequence](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates an empty sequence of strings or integers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence/init(empty:)
func NewSequenceEmptySequenceWithType(type_ FeatureType) Sequence {
	rv := objc.Send[Sequence](objc.ID(getSequenceClass().class), objc.Sel("emptySequenceWithType:"), type_)
	return rv
}


// Creates a sequence of integers from an array of numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence/init(int64s:)
func NewSequenceWithInt64Array(int64Values []foundation.Number) Sequence {
	rv := objc.Send[Sequence](objc.ID(getSequenceClass().class), objc.Sel("sequenceWithInt64Array:"), int64Values)
	return rv
}


// Creates a sequence of strings from a string array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence/init(strings:)
func NewSequenceWithStringArray(stringValues []string) Sequence {
	rv := objc.Send[Sequence](objc.ID(getSequenceClass().class), objc.Sel("sequenceWithStringArray:"), stringValues)
	return rv
}







// Creates an empty sequence of strings or integers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence/init(empty:)
func (sc _SequenceClass) EmptySequenceWithType(type_ FeatureType) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("emptySequenceWithType:"), type_)
	return rv
}


// Creates a sequence of integers from an array of numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence/init(int64s:)
func (sc _SequenceClass) SequenceWithInt64Array(int64Values []foundation.Number) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sequenceWithInt64Array:"), int64Values)
	return rv
}


// Creates a sequence of strings from a string array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence/init(strings:)
func (sc _SequenceClass) SequenceWithStringArray(stringValues []string) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("sequenceWithStringArray:"), stringValues)
	return rv
}

















// An array of 64-bit integers in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence/int64Values
func (s_ Sequence) Int64Values() []foundation.Number {
	rv := objc.Send[[]foundation.Number](s_.ID, objc.Sel("int64Values"))
	return rv
}


// An array of strings in the sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence/stringValues
func (s_ Sequence) StringValues() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("stringValues"))
	return rv
}


// The underlying type of the sequence’s elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequence/type
func (s_ Sequence) Type() FeatureType {
	rv := objc.Send[FeatureType](s_.ID, objc.Sel("type"))
	return rv
}







