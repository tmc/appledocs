// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
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
	FeatureName() objc.IObject /* cross-framework: NSString */
	SetFeatureName(value objc.IObject /* cross-framework: NSString */)
	FeatureValue() objc.IObject /* cross-framework: FeatureValue */
	SetFeatureValue(value objc.IObject /* cross-framework: FeatureValue */)
	ModelDescription() objc.IObject /* cross-framework: ModelDescription */
	SetModelDescription(value objc.IObject /* cross-framework: ModelDescription */)
	OutputDescriptionsByName() objc.IObject /* cross-framework: FeatureDescription */
	SetOutputDescriptionsByName(value objc.IObject /* cross-framework: FeatureDescription */)
	PredictedFeatureName() objc.IObject /* cross-framework: NSString */
	SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */)
	// methods:
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlfeaturevalueobservation/featurename
func (c_ CoreMLFeatureValueObservation) FeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("featureName"))
	return rv
}


// The name used in the model description of the CoreML model that produced this observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlfeaturevalueobservation/featurename
func (c_ CoreMLFeatureValueObservation) SetFeatureName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFeatureName:"), value)
}


// The feature result of a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlfeaturevalueobservation/featurevalue
func (c_ CoreMLFeatureValueObservation) FeatureValue() objc.IObject /* cross-framework: FeatureValue */ {
	rv := objc.Send[coreml.FeatureValue](c_.ID, objc.Sel("featureValue"))
	return rv
}


// The feature result of a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlfeaturevalueobservation/featurevalue
func (c_ CoreMLFeatureValueObservation) SetFeatureValue(value objc.IObject /* cross-framework: FeatureValue */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFeatureValue:"), value)
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLFeatureValueObservation) ModelDescription() objc.IObject /* cross-framework: ModelDescription */ {
	rv := objc.Send[coreml.ModelDescription](c_.ID, objc.Sel("modelDescription"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLFeatureValueObservation) SetModelDescription(value objc.IObject /* cross-framework: ModelDescription */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelDescription:"), value)
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (c_ CoreMLFeatureValueObservation) OutputDescriptionsByName() objc.IObject /* cross-framework: FeatureDescription */ {
	rv := objc.Send[coreml.FeatureDescription](c_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (c_ CoreMLFeatureValueObservation) SetOutputDescriptionsByName(value objc.IObject /* cross-framework: FeatureDescription */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLFeatureValueObservation) PredictedFeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("predictedFeatureName"))
	return rv
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLFeatureValueObservation) SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredictedFeatureName:"), value)
}



