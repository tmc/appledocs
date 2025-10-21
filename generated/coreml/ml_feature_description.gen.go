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
}

// The name, type, and constraints of an input or output feature.
//
// In Core ML, a is a single input or output of a model. A model can have any number of or . Each feature has a name and a value type, which are defined in the feature’s . Model authors use feature descriptions to help developers integrate their model properly. Each instance has read-only properties that indicate the feature’s name, its type, and whether it’s optional. For examples of features, see . Note the three input features named , , and , and the output feature is named . All four features are of type . An may also include constraints, which specify the limitations of the model’s input and output features. For each input feature, the constraints describe what values the model expects from your app. For each output feature, the constraints describe what values your app should expect from the model. You can also write code to inspect these descriptions before using the model in your app.
//
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


// The constraint for a dictionary feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/dictionaryConstraint
func (f_ FeatureDescription) DictionaryConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("dictionaryConstraint"))
	return rv
}

// The size and format constraints for an image feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/imageConstraint
func (f_ FeatureDescription) ImageConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("imageConstraint"))
	return rv
}

// A Boolean value that indicates whether this feature is optional.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/isOptional
func (f_ FeatureDescription) Optional() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("optional"))
	return rv
}

// The constraints on a multidimensional array feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/multiArrayConstraint
func (f_ FeatureDescription) MultiArrayConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("multiArrayConstraint"))
	return rv
}

// The name of this feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/name
func (f_ FeatureDescription) Name() string {
	rv := objc.Send[string](f_.ID, objc.Sel("name"))
	return rv
}

// The constraints for a sequence feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/sequenceConstraint
func (f_ FeatureDescription) SequenceConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("sequenceConstraint"))
	return rv
}

// The state feature value constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/stateConstraint
func (f_ FeatureDescription) StateConstraint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("stateConstraint"))
	return rv
}

// The type of this feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLFeatureDescription/type
func (f_ FeatureDescription) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("type"))
	return rv
}



