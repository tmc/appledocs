// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FeatureDescription] class.
var (
	FeatureDescriptionClass     _FeatureDescriptionClass
	FeatureDescriptionClassOnce sync.Once
)

func getFeatureDescriptionClass() _FeatureDescriptionClass {
	FeatureDescriptionClassOnce.Do(func() {
		FeatureDescriptionClass = _FeatureDescriptionClass{objc.GetClass("MLFeatureDescription")}
	})
	return FeatureDescriptionClass
}

type _FeatureDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [FeatureDescription] class.
type IFeatureDescription interface {
	objectivec.IObject
	MultiArrayConstraint() MultiArrayConstraint
	DictionaryConstraint() DictionaryConstraint
	SetDictionaryConstraint(value DictionaryConstraint)
	ImageConstraint() IMLImageConstraint
	SetImageConstraint(value IMLImageConstraint)
	IsOptional() bool
	SetIsOptional(value bool)
	Name() string
	SetName(value string)
	SequenceConstraint() SequenceConstraint
	SetSequenceConstraint(value SequenceConstraint)
	StateConstraint() IMLStateConstraint
	SetStateConstraint(value IMLStateConstraint)
	Type() MLFeatureType
	SetType(value MLFeatureType)
	InputDescriptionsByName() IMLFeatureDescription
	SetInputDescriptionsByName(value IMLFeatureDescription)
	OutputDescriptionsByName() IMLFeatureDescription
	SetOutputDescriptionsByName(value IMLFeatureDescription)
	StateDescriptionsByName() IMLFeatureDescription
	SetStateDescriptionsByName(value IMLFeatureDescription)
}

// The name, type, and constraints of an input or output feature.
//
// In Core ML, a is a single input or output of a model. A model can have any number of or . Each feature has a name and a value type, which are defined in the feature’s . Model authors use feature descriptions to help developers integrate their model properly. Each instance has read-only properties that indicate the feature’s name, its type, and whether it’s optional. For examples of features, see . Note the three input features named , , and , and the output feature is named . All four features are of type . An may also include constraints, which specify the limitations of the model’s input and output features. For each input feature, the constraints describe what values the model expects from your app. For each output feature, the constraints describe what values your app should expect from the model. You can also write code to inspect these descriptions before using the model in your app.


// The name, type, and constraints of an input or output feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription
type FeatureDescription struct {
	objectivec.Object
}

// FeatureDescriptionFrom constructs a [FeatureDescription] from an unsafe.Pointer.
//
// The name, type, and constraints of an input or output feature.
func FeatureDescriptionFrom(ptr unsafe.Pointer) FeatureDescription {
	return FeatureDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FeatureDescriptionClass) Alloc() FeatureDescription {
	rv := objc.Send[FeatureDescription](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FeatureDescriptionClass) New() FeatureDescription {
	rv := objc.Send[FeatureDescription](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FeatureDescription) Init() FeatureDescription {
	rv := objc.Send[FeatureDescription](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FeatureDescription) Autorelease() FeatureDescription {
	rv := objc.Send[FeatureDescription](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFeatureDescription creates a new FeatureDescription instance.
func NewFeatureDescription() FeatureDescription {
	return getFeatureDescriptionClass().New()
}



// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/multiArrayConstraint
func (f_ FeatureDescription) MultiArrayConstraint() MultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](f_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}


// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (f_ FeatureDescription) DictionaryConstraint() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](f_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}


// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/dictionaryconstraint
func (f_ FeatureDescription) SetDictionaryConstraint(value DictionaryConstraint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDictionaryConstraint:"), value)
}


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (f_ FeatureDescription) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[ImageConstraint](f_.ID, objc.Sel("imageConstraint"))
	return rv
}


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/imageconstraint
func (f_ FeatureDescription) SetImageConstraint(value IMLImageConstraint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setImageConstraint:"), value)
}


// A Boolean value that indicates whether this feature is optional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/isoptional
func (f_ FeatureDescription) IsOptional() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isOptional"))
	return rv
}


// A Boolean value that indicates whether this feature is optional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/isoptional
func (f_ FeatureDescription) SetIsOptional(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsOptional:"), value)
}


// The name of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/name
func (f_ FeatureDescription) Name() string {
	rv := objc.Send[string](f_.ID, objc.Sel("name"))
	return rv
}


// The name of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/name
func (f_ FeatureDescription) SetName(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setName:"), objc.String(value))
}


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (f_ FeatureDescription) SequenceConstraint() SequenceConstraint {
	rv := objc.Send[SequenceConstraint](f_.ID, objc.Sel("sequenceConstraint"))
	return rv
}


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/sequenceconstraint
func (f_ FeatureDescription) SetSequenceConstraint(value SequenceConstraint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSequenceConstraint:"), value)
}


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (f_ FeatureDescription) StateConstraint() IMLStateConstraint {
	rv := objc.Send[StateConstraint](f_.ID, objc.Sel("stateConstraint"))
	return rv
}


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/stateconstraint
func (f_ FeatureDescription) SetStateConstraint(value IMLStateConstraint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setStateConstraint:"), value)
}


// The type of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/type
func (f_ FeatureDescription) Type() MLFeatureType {
	rv := objc.Send[MLFeatureType](f_.ID, objc.Sel("type"))
	return rv
}


// The type of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/type
func (f_ FeatureDescription) SetType(value MLFeatureType) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setType:"), value)
}


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (f_ FeatureDescription) InputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](f_.ID, objc.Sel("inputDescriptionsByName"))
	return rv
}


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (f_ FeatureDescription) SetInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputDescriptionsByName:"), value)
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (f_ FeatureDescription) OutputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](f_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (f_ FeatureDescription) SetOutputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}


// Description of the state features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/statedescriptionsbyname
func (f_ FeatureDescription) StateDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](f_.ID, objc.Sel("stateDescriptionsByName"))
	return rv
}


// Description of the state features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/statedescriptionsbyname
func (f_ FeatureDescription) SetStateDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setStateDescriptionsByName:"), value)
}



