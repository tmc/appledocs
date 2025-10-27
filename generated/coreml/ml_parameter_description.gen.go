// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ParameterDescription] class.
type IParameterDescription interface {
	objectivec.IObject
	

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


	

	// methods:


}





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

























// The default value for the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterDescription/defaultValue
func (p_ ParameterDescription) DefaultValue() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("defaultValue"))
	return rv
}


// The key for this parameter description value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterDescription/key
func (p_ ParameterDescription) Key() IMLParameterKey {
	rv := objc.Send[ParameterKey](p_.ID, objc.Sel("key"))
	return rv
}


// The constraints of this paramter description value, if and only if the value is numerical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLParameterDescription/numericConstraint
func (p_ ParameterDescription) NumericConstraint() IMLNumericConstraint {
	rv := objc.Send[NumericConstraint](p_.ID, objc.Sel("numericConstraint"))
	return rv
}


// A Boolean value that indicates whether you can update the model with additional training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/isupdatable
func (p_ ParameterDescription) IsUpdatable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isUpdatable"))
	return rv
}


// A Boolean value that indicates whether you can update the model with additional training.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/isupdatable
func (p_ ParameterDescription) SetIsUpdatable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsUpdatable:"), value)
}


// A dictionary of the descriptions for the model’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/parameterdescriptionsbykey
func (p_ ParameterDescription) ParameterDescriptionsByKey() IMLParameterDescription {
	rv := objc.Send[ParameterDescription](p_.ID, objc.Sel("parameterDescriptionsByKey"))
	return rv
}


// A dictionary of the descriptions for the model’s parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/parameterdescriptionsbykey
func (p_ ParameterDescription) SetParameterDescriptionsByKey(value IMLParameterDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setParameterDescriptionsByKey:"), value)
}


// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/traininginputdescriptionsbyname
func (p_ ParameterDescription) TrainingInputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](p_.ID, objc.Sel("trainingInputDescriptionsByName"))
	return rv
}


// A dictionary of the training input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/traininginputdescriptionsbyname
func (p_ ParameterDescription) SetTrainingInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTrainingInputDescriptionsByName:"), value)
}








