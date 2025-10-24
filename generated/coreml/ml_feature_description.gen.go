// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLFeatureDescription */


/* debug [class_header]: Header for MLFeatureDescription */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FeatureDescription */
// An interface definition for the [FeatureDescription] class.
type IFeatureDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FeatureDescription */
	// properties:
	DictionaryConstraint() IMLDictionaryConstraint
	ImageConstraint() IMLImageConstraint
	Optional() bool
	MultiArrayConstraint() IMLMultiArrayConstraint
	Name() objc.IObject /* cross-framework: NSString */
	SequenceConstraint() IMLSequenceConstraint
	StateConstraint() IMLStateConstraint
	Type() FeatureType
	IsOptional() bool
	SetIsOptional(value bool)
	InputDescriptionsByName() IMLFeatureDescription
	SetInputDescriptionsByName(value IMLFeatureDescription)
	OutputDescriptionsByName() IMLFeatureDescription
	SetOutputDescriptionsByName(value IMLFeatureDescription)
	StateDescriptionsByName() IMLFeatureDescription
	SetStateDescriptionsByName(value IMLFeatureDescription)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FeatureDescription */
	// methods:
	IsAllowedValue(value IMLFeatureValue) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FeatureDescription */
// Alloc allocates a new instance without initialization.
func (fc _FeatureDescriptionClass) Alloc() FeatureDescription {
	rv := objc.Send[FeatureDescription](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FeatureDescription */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FeatureDescription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FeatureDescription */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FeatureDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FeatureDescription */

// Checks whether the model will accept an input feature value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/isAllowedValue(_:)
func (f_ FeatureDescription) IsAllowedValue(value IMLFeatureValue) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isAllowedValue:"), value)
	return rv
}/* debug [instance_methods/method]: IsAllowedValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FeatureDescription */

// The constraint for a dictionary feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/dictionaryConstraint
func (f_ FeatureDescription) DictionaryConstraint() IMLDictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](f_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}/* debug [instance_properties/getter]: dictionaryConstraint */


// The size and format constraints for an image feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/imageConstraint
func (f_ FeatureDescription) ImageConstraint() IMLImageConstraint {
	rv := objc.Send[ImageConstraint](f_.ID, objc.Sel("imageConstraint"))
	return rv
}/* debug [instance_properties/getter]: imageConstraint */


// A Boolean value that indicates whether this feature is optional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/isOptional
func (f_ FeatureDescription) Optional() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("optional"))
	return rv
}/* debug [instance_properties/getter]: optional */


// The constraints on a multidimensional array feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/multiArrayConstraint
func (f_ FeatureDescription) MultiArrayConstraint() IMLMultiArrayConstraint {
	rv := objc.Send[MultiArrayConstraint](f_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}/* debug [instance_properties/getter]: multiArrayConstraint */


// The name of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/name
func (f_ FeatureDescription) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The constraints for a sequence feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/sequenceConstraint
func (f_ FeatureDescription) SequenceConstraint() IMLSequenceConstraint {
	rv := objc.Send[SequenceConstraint](f_.ID, objc.Sel("sequenceConstraint"))
	return rv
}/* debug [instance_properties/getter]: sequenceConstraint */


// The state feature value constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/stateConstraint
func (f_ FeatureDescription) StateConstraint() IMLStateConstraint {
	rv := objc.Send[StateConstraint](f_.ID, objc.Sel("stateConstraint"))
	return rv
}/* debug [instance_properties/getter]: stateConstraint */


// The type of this feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/type
func (f_ FeatureDescription) Type() FeatureType {
	rv := objc.Send[FeatureType](f_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A Boolean value that indicates whether this feature is optional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/isoptional
func (f_ FeatureDescription) IsOptional() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isOptional"))
	return rv
}/* debug [instance_properties/getter]: isOptional */


// A Boolean value that indicates whether this feature is optional.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlfeaturedescription/isoptional
func (f_ FeatureDescription) SetIsOptional(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsOptional:"), value)
}/* debug [instance_properties/setter]: isOptional */


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (f_ FeatureDescription) InputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](f_.ID, objc.Sel("inputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: inputDescriptionsByName */


// A dictionary of input feature descriptions, which the model keys by the input’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/inputdescriptionsbyname
func (f_ FeatureDescription) SetInputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputDescriptionsByName:"), value)
}/* debug [instance_properties/setter]: inputDescriptionsByName */


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (f_ FeatureDescription) OutputDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](f_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: outputDescriptionsByName */


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/outputdescriptionsbyname
func (f_ FeatureDescription) SetOutputDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}/* debug [instance_properties/setter]: outputDescriptionsByName */


// Description of the state features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/statedescriptionsbyname
func (f_ FeatureDescription) StateDescriptionsByName() IMLFeatureDescription {
	rv := objc.Send[FeatureDescription](f_.ID, objc.Sel("stateDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: stateDescriptionsByName */


// Description of the state features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodeldescription/statedescriptionsbyname
func (f_ FeatureDescription) SetStateDescriptionsByName(value IMLFeatureDescription) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setStateDescriptionsByName:"), value)
}/* debug [instance_properties/setter]: stateDescriptionsByName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLFeatureDescription */



