// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLSequenceConstraint */


/* debug [class_header]: Header for MLSequenceConstraint */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SequenceConstraint */
// An interface definition for the [SequenceConstraint] class.
type ISequenceConstraint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SequenceConstraint */
	// properties:
	CountRange() corefoundation.Range
	ValueDescription() IMLFeatureDescription
	DictionaryConstraint() IMLDictionaryConstraint
	SetDictionaryConstraint(value IMLDictionaryConstraint)
	ImageConstraint() IMLImageConstraint
	SetImageConstraint(value IMLImageConstraint)
	MultiArrayConstraint() IMLMultiArrayConstraint
	SetMultiArrayConstraint(value IMLMultiArrayConstraint)
	SequenceConstraint() IMLSequenceConstraint
	SetSequenceConstraint(value IMLSequenceConstraint)
	StateConstraint() IMLStateConstraint
	SetStateConstraint(value IMLStateConstraint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SequenceConstraint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SequenceConstraint */
// Alloc allocates a new instance without initialization.
func (sc _SequenceConstraintClass) Alloc() SequenceConstraint {
	rv := objc.Send[SequenceConstraint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SequenceConstraint */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SequenceConstraint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SequenceConstraint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SequenceConstraint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SequenceConstraint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SequenceConstraint */

// The range of values allowed for the sequence’s length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/countRange
func (s_ SequenceConstraint) CountRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](s_.ID, objc.Sel("countRange"))
	return rv
}/* debug [instance_properties/getter]: countRange */


// The description that all sequence elements must match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLSequenceConstraint/valueDescription
func (s_ SequenceConstraint) ValueDescription() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](s_.ID, objc.Sel("valueDescription"))
	return rv
}/* debug [instance_properties/getter]: valueDescription */


// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (s_ SequenceConstraint) DictionaryConstraint() IMLDictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](s_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}/* debug [instance_properties/getter]: dictionaryConstraint */


// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (s_ SequenceConstraint) SetDictionaryConstraint(value IMLDictionaryConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDictionaryConstraint:"), value)
}/* debug [instance_properties/setter]: dictionaryConstraint */


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (s_ SequenceConstraint) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[ImageConstraint](s_.ID, objc.Sel("imageConstraint"))
	return rv
}/* debug [instance_properties/getter]: imageConstraint */


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (s_ SequenceConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setImageConstraint:"), value)
}/* debug [instance_properties/setter]: imageConstraint */


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (s_ SequenceConstraint) MultiArrayConstraint() IMLMultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](s_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}/* debug [instance_properties/getter]: multiArrayConstraint */


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (s_ SequenceConstraint) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMultiArrayConstraint:"), value)
}/* debug [instance_properties/setter]: multiArrayConstraint */


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (s_ SequenceConstraint) SequenceConstraint() IMLSequenceConstraint {
	rv := objc.Send[SequenceConstraint](s_.ID, objc.Sel("sequenceConstraint"))
	return rv
}/* debug [instance_properties/getter]: sequenceConstraint */


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (s_ SequenceConstraint) SetSequenceConstraint(value IMLSequenceConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSequenceConstraint:"), value)
}/* debug [instance_properties/setter]: sequenceConstraint */


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (s_ SequenceConstraint) StateConstraint() IMLStateConstraint {
	rv := objc.Send[StateConstraint](s_.ID, objc.Sel("stateConstraint"))
	return rv
}/* debug [instance_properties/getter]: stateConstraint */


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (s_ SequenceConstraint) SetStateConstraint(value IMLStateConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStateConstraint:"), value)
}/* debug [instance_properties/setter]: stateConstraint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLSequenceConstraint */



