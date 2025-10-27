// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CoreMLFeatureValueObservation] class.
var (
	CoreMLFeatureValueObservationClass     _CoreMLFeatureValueObservationClass
	CoreMLFeatureValueObservationClassOnce sync.Once
)

func getCoreMLFeatureValueObservationClass() _CoreMLFeatureValueObservationClass {
	CoreMLFeatureValueObservationClassOnce.Do(func() {
		CoreMLFeatureValueObservationClass = _CoreMLFeatureValueObservationClass{objc.GetClass("VNCoreMLFeatureValueObservation")}
	})
	return CoreMLFeatureValueObservationClass
}

type _CoreMLFeatureValueObservationClass struct {
	class objc.Class
}





// An interface definition for the [CoreMLFeatureValueObservation] class.
type ICoreMLFeatureValueObservation interface {
	IObservation
	

	// properties:
	FeatureName() foundation.foundation.INSString
	FeatureValue() coreml.FeatureValue
	ModelDescription() coreml.ModelDescription
	SetModelDescription(value coreml.ModelDescription)
	OutputDescriptionsByName() coreml.FeatureDescription
	SetOutputDescriptionsByName(value coreml.FeatureDescription)
	PredictedFeatureName() foundation.foundation.INSString
	SetPredictedFeatureName(value foundation.foundation.INSString)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CoreMLFeatureValueObservationClass) Alloc() CoreMLFeatureValueObservation {
	rv := objc.Send[CoreMLFeatureValueObservation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CoreMLFeatureValueObservationClass) New() CoreMLFeatureValueObservation {
	rv := objc.Send[CoreMLFeatureValueObservation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CoreMLFeatureValueObservation) Init() CoreMLFeatureValueObservation {
	rv := objc.Send[CoreMLFeatureValueObservation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CoreMLFeatureValueObservation) Autorelease() CoreMLFeatureValueObservation {
	rv := objc.Send[CoreMLFeatureValueObservation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLFeatureValueObservation creates a new CoreMLFeatureValueObservation instance.
func NewCoreMLFeatureValueObservation() CoreMLFeatureValueObservation {
	return getCoreMLFeatureValueObservationClass().New()
}





// An object that represents a collection of key-value information that a Core ML image-analysis request produces.
//
// This type of observation results from performing a image analysis with a Core ML model whose role is prediction rather than classification or image-to-image processing. Vision infers that an object is a predictor model if that model predicts multiple features. You can tell that a model predicts multiple features when its object has a value for its property, or when it inserts its output in an dictionary.


// An object that represents a collection of key-value information that a Core ML image-analysis request produces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation
type CoreMLFeatureValueObservation struct {
	Observation
}

// CoreMLFeatureValueObservationFrom constructs a [CoreMLFeatureValueObservation] from an unsafe.Pointer.
//
// An object that represents a collection of key-value information that a Core ML image-analysis request produces.
func CoreMLFeatureValueObservationFrom(ptr unsafe.Pointer) CoreMLFeatureValueObservation {
	return CoreMLFeatureValueObservation{
		Observation: ObservationFrom(ptr),
	}
}

























// The name used in the model description of the CoreML model that produced this observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation/featureName
func (c_ CoreMLFeatureValueObservation) FeatureName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("featureName"))
	return rv
}


// The feature result of a that outputs neither a classification nor an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation/featureValue
func (c_ CoreMLFeatureValueObservation) FeatureValue() coreml.FeatureValue {
	rv := objc.Send[coreml.FeatureValue](c_.ID, objc.Sel("featureValue"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLFeatureValueObservation) ModelDescription() coreml.ModelDescription {
	rv := objc.Send[coreml.ModelDescription](c_.ID, objc.Sel("modelDescription"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLFeatureValueObservation) SetModelDescription(value coreml.ModelDescription) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelDescription:"), value)
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (c_ CoreMLFeatureValueObservation) OutputDescriptionsByName() coreml.FeatureDescription {
	rv := objc.Send[coreml.FeatureDescription](c_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (c_ CoreMLFeatureValueObservation) SetOutputDescriptionsByName(value coreml.FeatureDescription) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLFeatureValueObservation) PredictedFeatureName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("predictedFeatureName"))
	return rv
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLFeatureValueObservation) SetPredictedFeatureName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredictedFeatureName:"), value)
}








