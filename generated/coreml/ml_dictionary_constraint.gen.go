// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLDictionaryConstraint */


/* debug [class_header]: Header for MLDictionaryConstraint */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DictionaryConstraint */
// An interface definition for the [DictionaryConstraint] class.
type IDictionaryConstraint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DictionaryConstraint */
	// properties:
	KeyType() FeatureType
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

	
/* debug [class_interface_methods]: Methods for DictionaryConstraint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DictionaryConstraint */
// Alloc allocates a new instance without initialization.
func (dc _DictionaryConstraintClass) Alloc() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DictionaryConstraint */
// The constraint on the keys for a dictionary feature.


// The constraint on the keys for a dictionary feature.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DictionaryConstraint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DictionaryConstraint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DictionaryConstraint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DictionaryConstraint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DictionaryConstraint */

// The key type for the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint/keyType
func (d_ DictionaryConstraint) KeyType() FeatureType {
	rv := objc.Send[FeatureType](d_.ID, objc.Sel("keyType"))
	return rv
}/* debug [instance_properties/getter]: keyType */


// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (d_ DictionaryConstraint) DictionaryConstraint() IMLDictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](d_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}/* debug [instance_properties/getter]: dictionaryConstraint */


// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (d_ DictionaryConstraint) SetDictionaryConstraint(value IMLDictionaryConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDictionaryConstraint:"), value)
}/* debug [instance_properties/setter]: dictionaryConstraint */


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (d_ DictionaryConstraint) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[ImageConstraint](d_.ID, objc.Sel("imageConstraint"))
	return rv
}/* debug [instance_properties/getter]: imageConstraint */


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (d_ DictionaryConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setImageConstraint:"), value)
}/* debug [instance_properties/setter]: imageConstraint */


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (d_ DictionaryConstraint) MultiArrayConstraint() IMLMultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](d_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}/* debug [instance_properties/getter]: multiArrayConstraint */


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (d_ DictionaryConstraint) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMultiArrayConstraint:"), value)
}/* debug [instance_properties/setter]: multiArrayConstraint */


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (d_ DictionaryConstraint) SequenceConstraint() IMLSequenceConstraint {
	rv := objc.Send[SequenceConstraint](d_.ID, objc.Sel("sequenceConstraint"))
	return rv
}/* debug [instance_properties/getter]: sequenceConstraint */


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (d_ DictionaryConstraint) SetSequenceConstraint(value IMLSequenceConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSequenceConstraint:"), value)
}/* debug [instance_properties/setter]: sequenceConstraint */


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (d_ DictionaryConstraint) StateConstraint() IMLStateConstraint {
	rv := objc.Send[StateConstraint](d_.ID, objc.Sel("stateConstraint"))
	return rv
}/* debug [instance_properties/getter]: stateConstraint */


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (d_ DictionaryConstraint) SetStateConstraint(value IMLStateConstraint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStateConstraint:"), value)
}/* debug [instance_properties/setter]: stateConstraint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLDictionaryConstraint */



