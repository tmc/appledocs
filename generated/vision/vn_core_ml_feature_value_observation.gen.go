// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coreml"
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
}

// An object that represents a collection of key-value information that a Core ML image-analysis request produces.
//
// This type of observation results from performing a image analysis with a Core ML model whose role is prediction rather than classification or image-to-image processing. Vision infers that an object is a predictor model if that model predicts multiple features. You can tell that a model predicts multiple features when its object has a value for its property, or when it inserts its output in an dictionary.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CoreMLFeatureValueObservationClass) Alloc() CoreMLFeatureValueObservation {
	rv := objc.Send[CoreMLFeatureValueObservation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The name used in the model description of the CoreML model that produced this observation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation/featureName
func (c_ CoreMLFeatureValueObservation) FeatureName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("featureName"))
	return rv
}

// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLFeatureValueObservation) ModelDescription() coreml.ModelDescription {
	rv := objc.Send[coreml.ModelDescription](c_.ID, objc.Sel("modelDescription"))
	return rv
}


// SetModelDescription sets the value of the modelDescription property.
// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLFeatureValueObservation) SetModelDescription(value coreml.IModelDescription) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelDescription:"), value)
}

// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (c_ CoreMLFeatureValueObservation) OutputDescriptionsByName() coreml.FeatureDescription {
	rv := objc.Send[coreml.FeatureDescription](c_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}


// SetOutputDescriptionsByName sets the value of the outputDescriptionsByName property.
// A dictionary of output feature descriptions, which the model keys by the output’s name.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (c_ CoreMLFeatureValueObservation) SetOutputDescriptionsByName(value coreml.IFeatureDescription) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}

// The name of the primary prediction feature output description.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLFeatureValueObservation) PredictedFeatureName() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("predictedFeatureName"))
	return rv
}


// SetPredictedFeatureName sets the value of the predictedFeatureName property.
// The name of the primary prediction feature output description.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLFeatureValueObservation) SetPredictedFeatureName(value appkit.string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredictedFeatureName:"), value)
}

// The feature result of a
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlfeaturevalueobservation/featurevalue
func (c_ CoreMLFeatureValueObservation) FeatureValue() coreml.FeatureValue {
	rv := objc.Send[coreml.FeatureValue](c_.ID, objc.Sel("featureValue"))
	return rv
}


// SetFeatureValue sets the value of the featureValue property.
// The feature result of a

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlfeaturevalueobservation/featurevalue
func (c_ CoreMLFeatureValueObservation) SetFeatureValue(value coreml.IFeatureValue) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFeatureValue:"), value)
}



