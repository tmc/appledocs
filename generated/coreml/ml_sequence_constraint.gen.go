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


// The description that all sequence elements must match.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/valuedescription
func (s_ SequenceConstraint) ValueDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("valueDescription"))
	return rv
}


// SetValueDescription sets the value of the valueDescription property.
// The description that all sequence elements must match.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/valuedescription
func (s_ SequenceConstraint) SetValueDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValueDescription:"), value)
}

// The range of values allowed for the sequence’s length.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/countrange
func (s_ SequenceConstraint) CountRange() foundation.Range {
	rv := objc.Send[foundation.Range](s_.ID, objc.Sel("countRange"))
	return rv
}


// SetCountRange sets the value of the countRange property.
// The range of values allowed for the sequence’s length.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlsequenceconstraint/countrange
func (s_ SequenceConstraint) SetCountRange(value foundation.Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCountRange:"), value)
}

// The size and format constraints for an image feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (s_ SequenceConstraint) ImageConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("imageConstraint"))
	return rv
}


// SetImageConstraint sets the value of the imageConstraint property.
// The size and format constraints for an image feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (s_ SequenceConstraint) SetImageConstraint(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageConstraint:"), value)
}

// The constraints on a multidimensional array feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (s_ SequenceConstraint) MultiArrayConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}


// SetMultiArrayConstraint sets the value of the multiArrayConstraint property.
// The constraints on a multidimensional array feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (s_ SequenceConstraint) SetMultiArrayConstraint(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMultiArrayConstraint:"), value)
}

// The constraint for a dictionary feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (s_ SequenceConstraint) DictionaryConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}


// SetDictionaryConstraint sets the value of the dictionaryConstraint property.
// The constraint for a dictionary feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (s_ SequenceConstraint) SetDictionaryConstraint(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDictionaryConstraint:"), value)
}

// The state feature value constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (s_ SequenceConstraint) StateConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stateConstraint"))
	return rv
}


// SetStateConstraint sets the value of the stateConstraint property.
// The state feature value constraint.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (s_ SequenceConstraint) SetStateConstraint(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStateConstraint:"), value)
}

// The constraints for a sequence feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (s_ SequenceConstraint) SequenceConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sequenceConstraint"))
	return rv
}


// SetSequenceConstraint sets the value of the sequenceConstraint property.
// The constraints for a sequence feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (s_ SequenceConstraint) SetSequenceConstraint(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSequenceConstraint:"), value)
}



