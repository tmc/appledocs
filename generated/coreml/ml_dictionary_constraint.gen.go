// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DictionaryConstraint] class.
var (
	DictionaryConstraintClass     _DictionaryConstraintClass
	DictionaryConstraintClassOnce sync.Once
)

func getDictionaryConstraintClass() _DictionaryConstraintClass {
	DictionaryConstraintClassOnce.Do(func() {
		DictionaryConstraintClass = _DictionaryConstraintClass{objc.GetClass("MLDictionaryConstraint")}
	})
	return DictionaryConstraintClass
}

type _DictionaryConstraintClass struct {
	class objc.Class
}

// An interface definition for the [DictionaryConstraint] class.
type IDictionaryConstraint interface {
	objectivec.IObject
}

// The constraint on the keys for a dictionary feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint
type DictionaryConstraint struct {
	objectivec.Object
}

// DictionaryConstraintFrom constructs a [DictionaryConstraint] from an unsafe.Pointer.
//
// The constraint on the keys for a dictionary feature.
func DictionaryConstraintFrom(ptr unsafe.Pointer) DictionaryConstraint {
	return DictionaryConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DictionaryConstraintClass) Alloc() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DictionaryConstraintClass) New() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DictionaryConstraint) Init() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DictionaryConstraint) Autorelease() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDictionaryConstraint creates a new DictionaryConstraint instance.
func NewDictionaryConstraint() DictionaryConstraint {
	return getDictionaryConstraintClass().New()
}


// The key type for the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mldictionaryconstraint/keytype
func (d_ DictionaryConstraint) KeyType() FeatureType {
	rv := objc.Send[FeatureType](d_.ID, objc.Sel("keyType"))
	return rv
}


// SetKeyType sets the value of the keyType property.
// The key type for the dictionary.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mldictionaryconstraint/keytype
func (d_ DictionaryConstraint) SetKeyType(value FeatureType) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setKeyType:"), value)
}

// The constraint for a dictionary feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (d_ DictionaryConstraint) DictionaryConstraint() MLDictionaryConstraint {
	rv := objc.Send[MLDictionaryConstraint](d_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}


// SetDictionaryConstraint sets the value of the dictionaryConstraint property.
// The constraint for a dictionary feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (d_ DictionaryConstraint) SetDictionaryConstraint(value IMLDictionaryConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDictionaryConstraint:"), value)
}

// The size and format constraints for an image feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (d_ DictionaryConstraint) ImageConstraint() MLImageConstraint {
	rv := objc.Send[MLImageConstraint](d_.ID, objc.Sel("imageConstraint"))
	return rv
}


// SetImageConstraint sets the value of the imageConstraint property.
// The size and format constraints for an image feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (d_ DictionaryConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setImageConstraint:"), value)
}

// The constraints on a multidimensional array feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (d_ DictionaryConstraint) MultiArrayConstraint() MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](d_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}


// SetMultiArrayConstraint sets the value of the multiArrayConstraint property.
// The constraints on a multidimensional array feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (d_ DictionaryConstraint) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMultiArrayConstraint:"), value)
}

// The constraints for a sequence feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (d_ DictionaryConstraint) SequenceConstraint() MLSequenceConstraint {
	rv := objc.Send[MLSequenceConstraint](d_.ID, objc.Sel("sequenceConstraint"))
	return rv
}


// SetSequenceConstraint sets the value of the sequenceConstraint property.
// The constraints for a sequence feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (d_ DictionaryConstraint) SetSequenceConstraint(value IMLSequenceConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSequenceConstraint:"), value)
}

// The state feature value constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (d_ DictionaryConstraint) StateConstraint() MLStateConstraint {
	rv := objc.Send[MLStateConstraint](d_.ID, objc.Sel("stateConstraint"))
	return rv
}


// SetStateConstraint sets the value of the stateConstraint property.
// The state feature value constraint.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (d_ DictionaryConstraint) SetStateConstraint(value IMLStateConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStateConstraint:"), value)
}



