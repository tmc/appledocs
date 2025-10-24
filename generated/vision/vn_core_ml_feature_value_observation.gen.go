// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class VNCoreMLFeatureValueObservation */


/* debug [class_header]: Header for VNCoreMLFeatureValueObservation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CoreMLFeatureValueObservation */
// An interface definition for the [CoreMLFeatureValueObservation] class.
type ICoreMLFeatureValueObservation interface {
	IObservation
	
/* debug [class_interface_properties]: Properties for CoreMLFeatureValueObservation */
	// properties:
	FeatureName() objc.IObject /* cross-framework: NSString */
	FeatureValue() coreml.FeatureValue
	ModelDescription() coreml.ModelDescription
	SetModelDescription(value coreml.ModelDescription)
	OutputDescriptionsByName() coreml.FeatureDescription
	SetOutputDescriptionsByName(value coreml.FeatureDescription)
	PredictedFeatureName() objc.IObject /* cross-framework: NSString */
	SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CoreMLFeatureValueObservation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CoreMLFeatureValueObservation */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CoreMLFeatureValueObservation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CoreMLFeatureValueObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CoreMLFeatureValueObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CoreMLFeatureValueObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CoreMLFeatureValueObservation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CoreMLFeatureValueObservation */

// The name used in the model description of the CoreML model that produced this observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation/featureName
func (c_ CoreMLFeatureValueObservation) FeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("featureName"))
	return rv
}/* debug [instance_properties/getter]: featureName */


// The feature result of a that outputs neither a classification nor an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation/featureValue
func (c_ CoreMLFeatureValueObservation) FeatureValue() coreml.FeatureValue {
	rv := objc.Send[coreml.FeatureValue](c_.ID, objc.Sel("featureValue"))
	return rv
}/* debug [instance_properties/getter]: featureValue */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLFeatureValueObservation) ModelDescription() coreml.ModelDescription {
	rv := objc.Send[coreml.ModelDescription](c_.ID, objc.Sel("modelDescription"))
	return rv
}/* debug [instance_properties/getter]: modelDescription */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLFeatureValueObservation) SetModelDescription(value coreml.ModelDescription) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelDescription:"), value)
}/* debug [instance_properties/setter]: modelDescription */


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (c_ CoreMLFeatureValueObservation) OutputDescriptionsByName() coreml.FeatureDescription {
	rv := objc.Send[coreml.FeatureDescription](c_.ID, objc.Sel("outputDescriptionsByName"))
	return rv
}/* debug [instance_properties/getter]: outputDescriptionsByName */


// A dictionary of output feature descriptions, which the model keys by the output’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (c_ CoreMLFeatureValueObservation) SetOutputDescriptionsByName(value coreml.FeatureDescription) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}/* debug [instance_properties/setter]: outputDescriptionsByName */


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLFeatureValueObservation) PredictedFeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("predictedFeatureName"))
	return rv
}/* debug [instance_properties/getter]: predictedFeatureName */


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLFeatureValueObservation) SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredictedFeatureName:"), value)
}/* debug [instance_properties/setter]: predictedFeatureName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNCoreMLFeatureValueObservation */



