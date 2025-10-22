// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	DictionaryConstraint() MLDictionaryConstraint
	SetDictionaryConstraint(value IMLDictionaryConstraint)
	ImageConstraint() MLImageConstraint
	SetImageConstraint(value IMLImageConstraint)
	MultiArrayConstraint() MLMultiArrayConstraint
	SetMultiArrayConstraint(value IMLMultiArrayConstraint)
	SequenceConstraint() MLSequenceConstraint
	SetSequenceConstraint(value IMLSequenceConstraint)
	StateConstraint() MLStateConstraint
	SetStateConstraint(value IMLStateConstraint)
	CountRange() foundation.Range
	SetCountRange(value foundation.Range)
	ValueDescription() MLFeatureDescription
	SetValueDescription(value IMLFeatureDescription)
}

// The constraints for a sequence feature.


// The constraints for a sequence feature.
//
// [Full Topic]
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



// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint

func (s_ SequenceConstraint) DictionaryConstraint() MLDictionaryConstraint {
	rv := objc.Send[MLDictionaryConstraint](s_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}


// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint

func (s_ SequenceConstraint) SetDictionaryConstraint(value IMLDictionaryConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDictionaryConstraint:"), value)
}


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint

func (s_ SequenceConstraint) ImageConstraint() MLImageConstraint {
	rv := objc.Send[MLImageConstraint](s_.ID, objc.Sel("imageConstraint"))
	return rv
}


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint

func (s_ SequenceConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageConstraint:"), value)
}


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint

func (s_ SequenceConstraint) MultiArrayConstraint() MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](s_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint

func (s_ SequenceConstraint) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMultiArrayConstraint:"), value)
}


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint

func (s_ SequenceConstraint) SequenceConstraint() MLSequenceConstraint {
	rv := objc.Send[MLSequenceConstraint](s_.ID, objc.Sel("sequenceConstraint"))
	return rv
}


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint

func (s_ SequenceConstraint) SetSequenceConstraint(value IMLSequenceConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSequenceConstraint:"), value)
}


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint

func (s_ SequenceConstraint) StateConstraint() MLStateConstraint {
	rv := objc.Send[MLStateConstraint](s_.ID, objc.Sel("stateConstraint"))
	return rv
}


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint

func (s_ SequenceConstraint) SetStateConstraint(value IMLStateConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStateConstraint:"), value)
}


// The range of values allowed for the sequence’s length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/countrange

func (s_ SequenceConstraint) CountRange() foundation.Range {
	rv := objc.Send[foundation.Range](s_.ID, objc.Sel("countRange"))
	return rv
}


// The range of values allowed for the sequence’s length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/countrange

func (s_ SequenceConstraint) SetCountRange(value foundation.Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCountRange:"), value)
}


// The description that all sequence elements must match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/valuedescription

func (s_ SequenceConstraint) ValueDescription() MLFeatureDescription {
	rv := objc.Send[MLFeatureDescription](s_.ID, objc.Sel("valueDescription"))
	return rv
}


// The description that all sequence elements must match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/valuedescription

func (s_ SequenceConstraint) SetValueDescription(value IMLFeatureDescription) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValueDescription:"), value)
}



