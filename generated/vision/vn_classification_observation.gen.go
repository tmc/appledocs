// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ClassificationObservation] class.
var (
	ClassificationObservationClass     _ClassificationObservationClass
	ClassificationObservationClassOnce sync.Once
)

func getClassificationObservationClass() _ClassificationObservationClass {
	ClassificationObservationClassOnce.Do(func() {
		ClassificationObservationClass = _ClassificationObservationClass{objc.GetClass("VNClassificationObservation")}
	})
	return ClassificationObservationClass
}

type _ClassificationObservationClass struct {
	class objc.Class
}





// An interface definition for the [ClassificationObservation] class.
type IClassificationObservation interface {
	IObservation
	

	// properties:
	HasPrecisionRecallCurve() bool
	Identifier() objc.IObject /* cross-framework: NSString */
	ModelDescription() coreml.ModelDescription
	SetModelDescription(value coreml.ModelDescription)
	PredictedFeatureName() objc.IObject /* cross-framework: NSString */
	SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */)


	

	// methods:
	HasMinimumPrecisionForRecall(minimumPrecision float32, recall float32) bool
	HasMinimumRecallForPrecision(minimumRecall float32, precision float32) bool


}





// Alloc allocates a new instance without initialization.
func (cc _ClassificationObservationClass) Alloc() ClassificationObservation {
	rv := objc.Send[ClassificationObservation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ClassificationObservationClass) New() ClassificationObservation {
	rv := objc.Send[ClassificationObservation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ClassificationObservation) Init() ClassificationObservation {
	rv := objc.Send[ClassificationObservation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ClassificationObservation) Autorelease() ClassificationObservation {
	rv := objc.Send[ClassificationObservation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewClassificationObservation creates a new ClassificationObservation instance.
func NewClassificationObservation() ClassificationObservation {
	return getClassificationObservationClass().New()
}





// An object that represents classification information that an image-analysis request produces.
//
// This type of observation results from performing a image analysis with a Core ML model whose role is classification (rather than prediction or image-to-image processing). Vision infers that an object is a classifier model if that model predicts a single feature. That is, the model’s object has a non- value for its property.


// An object that represents classification information that an image-analysis request produces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNClassificationObservation
type ClassificationObservation struct {
	Observation
}

// ClassificationObservationFrom constructs a [ClassificationObservation] from an unsafe.Pointer.
//
// An object that represents classification information that an image-analysis request produces.
func ClassificationObservationFrom(ptr unsafe.Pointer) ClassificationObservation {
	return ClassificationObservation{
		Observation: ObservationFrom(ptr),
	}
}




















// Determines whether the observation for a specific recall has a minimum precision value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNClassificationObservation/hasMinimumPrecision(_:forRecall:)
func (c_ ClassificationObservation) HasMinimumPrecisionForRecall(minimumPrecision float32, recall float32) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasMinimumPrecision:forRecall:"), minimumPrecision, recall)
	return rv
}


// Determines whether the observation for a specific precision has a minimum recall value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNClassificationObservation/hasMinimumRecall(_:forPrecision:)
func (c_ ClassificationObservation) HasMinimumRecallForPrecision(minimumRecall float32, precision float32) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasMinimumRecall:forPrecision:"), minimumRecall, precision)
	return rv
}







// A Boolean variable indicating whether the observation contains precision and recall curves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNClassificationObservation/hasPrecisionRecallCurve
func (c_ ClassificationObservation) HasPrecisionRecallCurve() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasPrecisionRecallCurve"))
	return rv
}


// Classification label identifying the type of observation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNClassificationObservation/identifier
func (c_ ClassificationObservation) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ ClassificationObservation) ModelDescription() coreml.ModelDescription {
	rv := objc.Send[coreml.ModelDescription](c_.ID, objc.Sel("modelDescription"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ ClassificationObservation) SetModelDescription(value coreml.ModelDescription) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelDescription:"), value)
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ ClassificationObservation) PredictedFeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("predictedFeatureName"))
	return rv
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ ClassificationObservation) SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredictedFeatureName:"), value)
}








