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
	HasMinimumPrecisionForRecall(minimumPrecision unsafe.Pointer, recall unsafe.Pointer) bool
	HasMinimumRecallForPrecision(minimumRecall unsafe.Pointer, precision unsafe.Pointer) bool
}

// An object that represents classification information that an image-analysis request produces.
//
// This type of observation results from performing a image analysis with a Core ML model whose role is classification (rather than prediction or image-to-image processing). Vision infers that an object is a classifier model if that model predicts a single feature. That is, the model’s object has a non- value for its property.
//
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

// Alloc allocates a new instance without initialization.
func (cc _ClassificationObservationClass) Alloc() ClassificationObservation {
	rv := objc.Send[ClassificationObservation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Determines whether the observation for a specific recall has a minimum precision value.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNClassificationObservation/hasMinimumPrecision(_:forRecall:)
func (c_ ClassificationObservation) HasMinimumPrecisionForRecall(minimumPrecision unsafe.Pointer, recall unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasMinimumPrecision:forRecall:"), minimumPrecision, recall)
	return rv
}

// Determines whether the observation for a specific precision has a minimum recall value.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNClassificationObservation/hasMinimumRecall(_:forPrecision:)
func (c_ ClassificationObservation) HasMinimumRecallForPrecision(minimumRecall unsafe.Pointer, precision unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasMinimumRecall:forPrecision:"), minimumRecall, precision)
	return rv
}

// Classification label identifying the type of observation.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassificationobservation/identifier
func (c_ ClassificationObservation) Identifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// Classification label identifying the type of observation.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnclassificationobservation/identifier
func (c_ ClassificationObservation) SetIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// The name of the primary prediction feature output description.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ ClassificationObservation) PredictedFeatureName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("predictedFeatureName"))
	return rv
}


// SetPredictedFeatureName sets the value of the predictedFeatureName property.
// The name of the primary prediction feature output description.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ ClassificationObservation) SetPredictedFeatureName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredictedFeatureName:"), objc.String(value))
}

// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ ClassificationObservation) ModelDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modelDescription"))
	return rv
}


// SetModelDescription sets the value of the modelDescription property.
// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ ClassificationObservation) SetModelDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelDescription:"), value)
}

// A Boolean variable indicating whether the observation contains precision and recall curves.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNClassificationObservation/hasPrecisionRecallCurve
func (c_ ClassificationObservation) HasPrecisionRecallCurve() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasPrecisionRecallCurve"))
	return rv
}



