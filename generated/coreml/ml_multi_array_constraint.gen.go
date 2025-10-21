// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MultiArrayConstraint] class.
var (
	MultiArrayConstraintClass     _MultiArrayConstraintClass
	MultiArrayConstraintClassOnce sync.Once
)

func getMultiArrayConstraintClass() _MultiArrayConstraintClass {
	MultiArrayConstraintClassOnce.Do(func() {
		MultiArrayConstraintClass = _MultiArrayConstraintClass{objc.GetClass("MLMultiArrayConstraint")}
	})
	return MultiArrayConstraintClass
}

type _MultiArrayConstraintClass struct {
	class objc.Class
}

// An interface definition for the [MultiArrayConstraint] class.
type IMultiArrayConstraint interface {
	objectivec.IObject
}

// The shape and data type constraints for a multidimensional array feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint
type MultiArrayConstraint struct {
	objectivec.Object
}

// MultiArrayConstraintFrom constructs a [MultiArrayConstraint] from an unsafe.Pointer.
//
// The shape and data type constraints for a multidimensional array feature.
func MultiArrayConstraintFrom(ptr unsafe.Pointer) MultiArrayConstraint {
	return MultiArrayConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MultiArrayConstraintClass) Alloc() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MultiArrayConstraintClass) New() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiArrayConstraint) Init() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiArrayConstraint) Autorelease() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiArrayConstraint creates a new MultiArrayConstraint instance.
func NewMultiArrayConstraint() MultiArrayConstraint {
	return getMultiArrayConstraintClass().New()
}


// The type for the multi array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/dataType
func (m_ MultiArrayConstraint) DataType() MultiArrayDataType {
	rv := objc.Send[MultiArrayDataType](m_.ID, objc.Sel("dataType"))
	return rv
}

// The shape of the multi array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/shape
func (m_ MultiArrayConstraint) Shape() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("shape"))
	return rv
}

// The constraint on the shape of the multiarray.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/shapeConstraint
func (m_ MultiArrayConstraint) ShapeConstraint() MLMultiArrayShapeConstraint {
	rv := objc.Send[MLMultiArrayShapeConstraint](m_.ID, objc.Sel("shapeConstraint"))
	return rv
}

// The constraint for a dictionary feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (m_ MultiArrayConstraint) DictionaryConstraint() MLDictionaryConstraint {
	rv := objc.Send[MLDictionaryConstraint](m_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}


// SetDictionaryConstraint sets the value of the dictionaryConstraint property.
// The constraint for a dictionary feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (m_ MultiArrayConstraint) SetDictionaryConstraint(value IMLDictionaryConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDictionaryConstraint:"), value)
}

// The size and format constraints for an image feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (m_ MultiArrayConstraint) ImageConstraint() MLImageConstraint {
	rv := objc.Send[MLImageConstraint](m_.ID, objc.Sel("imageConstraint"))
	return rv
}


// SetImageConstraint sets the value of the imageConstraint property.
// The size and format constraints for an image feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (m_ MultiArrayConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageConstraint:"), value)
}

// The constraints on a multidimensional array feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (m_ MultiArrayConstraint) MultiArrayConstraint() MLMultiArrayConstraint {
	rv := objc.Send[MLMultiArrayConstraint](m_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}


// SetMultiArrayConstraint sets the value of the multiArrayConstraint property.
// The constraints on a multidimensional array feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (m_ MultiArrayConstraint) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMultiArrayConstraint:"), value)
}

// The constraints for a sequence feature.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (m_ MultiArrayConstraint) SequenceConstraint() MLSequenceConstraint {
	rv := objc.Send[MLSequenceConstraint](m_.ID, objc.Sel("sequenceConstraint"))
	return rv
}


// SetSequenceConstraint sets the value of the sequenceConstraint property.
// The constraints for a sequence feature.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (m_ MultiArrayConstraint) SetSequenceConstraint(value IMLSequenceConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSequenceConstraint:"), value)
}

// The state feature value constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (m_ MultiArrayConstraint) StateConstraint() MLStateConstraint {
	rv := objc.Send[MLStateConstraint](m_.ID, objc.Sel("stateConstraint"))
	return rv
}


// SetStateConstraint sets the value of the stateConstraint property.
// The state feature value constraint.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (m_ MultiArrayConstraint) SetStateConstraint(value IMLStateConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStateConstraint:"), value)
}



