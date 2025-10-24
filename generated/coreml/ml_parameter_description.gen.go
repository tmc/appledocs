// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLParameterDescription */


/* debug [class_header]: Header for MLParameterDescription */
// The class instance for the [ParameterDescription] class.
var (
	ParameterDescriptionClass     _ParameterDescriptionClass
	ParameterDescriptionClassOnce sync.Once
)

func getParameterDescriptionClass() _ParameterDescriptionClass {
	ParameterDescriptionClassOnce.Do(func() {
		ParameterDescriptionClass = _ParameterDescriptionClass{objc.GetClass("MLParameterDescription")}
	})
	return ParameterDescriptionClass
}

type _ParameterDescriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ParameterDescription */
// An interface definition for the [ParameterDescription] class.
type IParameterDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ParameterDescription */
	// properties:
	DefaultValue() objc.ID
	Key() IMLParameterKey
	NumericConstraint() IMLNumericConstraint
	IsUpdatable() bool
	SetIsUpdatable(value bool)
	ParameterDescriptionsByKey() IMLParameterDescription
	SetParameterDescriptionsByKey(value IMLParameterDescription)
	TrainingInputDescriptionsByName() IMLFeatureDescription
	SetTrainingInputDescriptionsByName(value IMLFeatureDescription)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ParameterDescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ParameterDescription */
// Alloc allocates a new instance without initialization.
func (pc _ParameterDescriptionClass) Alloc() ParameterDescription {
	rv := objc.Send[ParameterDescription](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ParameterDescriptionClass) New() ParameterDescription {
	rv := objc.Send[ParameterDescription](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParameterDescription) Init() ParameterDescription {
	rv := objc.Send[ParameterDescription](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParameterDescription) Autorelease() ParameterDescription {
	rv := objc.Send[ParameterDescription](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParameterDescription creates a new ParameterDescription instance.
func NewParameterDescription() ParameterDescription {
	return getParameterDescriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ParameterDescription */
// A description of a model parameter that includes a default value and a constraint, if applicable.


// A description of a model parameter that includes a default value and a constraint, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterDescription
type ParameterDescription struct {
	objectivec.Object
}

// ParameterDescriptionFrom constructs a [ParameterDescription] from an unsafe.Pointer.
//
// A description of a model parameter that includes a default value and a constraint, if applicable.
func ParameterDescriptionFrom(ptr unsafe.Pointer) ParameterDescription {
	return ParameterDescription{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ParameterDescription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ParameterDescription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ParameterDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ParameterDescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ParameterDescription */

// The default value for the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterDescription/defaultValue
func (p_ ParameterDescription) DefaultValue() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("defaultValue"))
	return rv
}/* debug [instance_properties/getter]: defaultValue */


// The key for this parameter description value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterDescription/key
func (p_ ParameterDescription) Key() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("key"))
	return rv
}/* debug [instance_properties/getter]: key */


// The constraints of this paramter description value, if and only if the value is numerical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterDescription/numericConstraint
func (p_ ParameterDescription) NumericConstraint() IMLNumericConstraint {
	rv := objc.Send[NumericConstraint](p_.ID, objc.Sel("numericConstraint"))
	return rv
}/* debug [instance_properties/getter]: numericConstraint */


// A Boolean value that indicates whether you can update the model with additional training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/isupdatable
func (p_ ParameterDescription) IsUpdatable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isUpdatable"))
	return rv
}/* debug [instance_properties/getter]: isUpdatable */


// A Boolean value that indicates whether you can update the model with additional training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/isupdatable
func (p_ ParameterDescription) SetIsUpdatable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsUpdatable:"), value)
}/* debug [instance_properties/setter]: isUpdatable */


// A dictionary of the descriptions for the model’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/parameterdescriptionsbykey
func (p_ ParameterDescription) ParameterDescriptionsByKey() IMLParameterDescription {
	rv := objc.Send[ParameterDescription](p_.ID, objc.Sel("parameterDescriptionsByKey"))
	return rv
}/* debug [instance_properties/getter]: parameterDescriptionsByKey */


// A dictionary of the descriptions for the model’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/parameterdescriptionsbykey
func (p_ ParameterDescription) SetParameterDescriptionsByKey(value IMLParameterDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParameterDescriptionsByKey:"), value)
}/* debug [instance_properties/setter]: parameterDescriptionsByKey */


// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/traininginputdescriptionsbyname
func (p_ ParameterDescription) TrainingInputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](p_.ID, objc.Sel("trainingInputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: trainingInputDescriptionsByName */


// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/traininginputdescriptionsbyname
func (p_ ParameterDescription) SetTrainingInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTrainingInputDescriptionsByName:"), value)
}/* debug [instance_properties/setter]: trainingInputDescriptionsByName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLParameterDescription */



