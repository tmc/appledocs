// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLMultiArrayConstraint */


/* debug [class_header]: Header for MLMultiArrayConstraint */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MultiArrayConstraint */
// An interface definition for the [MultiArrayConstraint] class.
type IMultiArrayConstraint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MultiArrayConstraint */
	// properties:
	DataType() MultiArrayDataType
	Shape() []foundation.Number
	ShapeConstraint() IMLMultiArrayShapeConstraint
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

	
/* debug [class_interface_methods]: Methods for MultiArrayConstraint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MultiArrayConstraint */
// Alloc allocates a new instance without initialization.
func (mc _MultiArrayConstraintClass) Alloc() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MultiArrayConstraint */
// The shape and data type constraints for a multidimensional array feature.


// The shape and data type constraints for a multidimensional array feature.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MultiArrayConstraint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MultiArrayConstraint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MultiArrayConstraint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MultiArrayConstraint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MultiArrayConstraint */

// The type for the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/dataType
func (m_ MultiArrayConstraint) DataType() MultiArrayDataType {
	rv := objc.Send[MultiArrayDataType](m_.ID, objc.Sel("dataType"))
	return rv
}/* debug [instance_properties/getter]: dataType */


// The shape of the multi array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/shape
func (m_ MultiArrayConstraint) Shape() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */


// The constraint on the shape of the multiarray.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMultiArrayConstraint/shapeConstraint
func (m_ MultiArrayConstraint) ShapeConstraint() IMLMultiArrayShapeConstraint {
	rv := objc.Send[MultiArrayShapeConstraint](m_.ID, objc.Sel("shapeConstraint"))
	return rv
}/* debug [instance_properties/getter]: shapeConstraint */


// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (m_ MultiArrayConstraint) DictionaryConstraint() IMLDictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](m_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}/* debug [instance_properties/getter]: dictionaryConstraint */


// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (m_ MultiArrayConstraint) SetDictionaryConstraint(value IMLDictionaryConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDictionaryConstraint:"), value)
}/* debug [instance_properties/setter]: dictionaryConstraint */


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (m_ MultiArrayConstraint) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[ImageConstraint](m_.ID, objc.Sel("imageConstraint"))
	return rv
}/* debug [instance_properties/getter]: imageConstraint */


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (m_ MultiArrayConstraint) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageConstraint:"), value)
}/* debug [instance_properties/setter]: imageConstraint */


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (m_ MultiArrayConstraint) MultiArrayConstraint() IMLMultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](m_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}/* debug [instance_properties/getter]: multiArrayConstraint */


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/multiarrayconstraint
func (m_ MultiArrayConstraint) SetMultiArrayConstraint(value IMLMultiArrayConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMultiArrayConstraint:"), value)
}/* debug [instance_properties/setter]: multiArrayConstraint */


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (m_ MultiArrayConstraint) SequenceConstraint() IMLSequenceConstraint {
	rv := objc.Send[SequenceConstraint](m_.ID, objc.Sel("sequenceConstraint"))
	return rv
}/* debug [instance_properties/getter]: sequenceConstraint */


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (m_ MultiArrayConstraint) SetSequenceConstraint(value IMLSequenceConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSequenceConstraint:"), value)
}/* debug [instance_properties/setter]: sequenceConstraint */


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (m_ MultiArrayConstraint) StateConstraint() IMLStateConstraint {
	rv := objc.Send[StateConstraint](m_.ID, objc.Sel("stateConstraint"))
	return rv
}/* debug [instance_properties/getter]: stateConstraint */


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (m_ MultiArrayConstraint) SetStateConstraint(value IMLStateConstraint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStateConstraint:"), value)
}/* debug [instance_properties/setter]: stateConstraint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLMultiArrayConstraint */



