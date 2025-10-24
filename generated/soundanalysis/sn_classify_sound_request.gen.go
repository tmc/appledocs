// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SNClassifySoundRequest */


/* debug [class_header]: Header for SNClassifySoundRequest */
// The class instance for the [SNClassifySoundRequest] class.
var (
	SNClassifySoundRequestClass     _SNClassifySoundRequestClass
	SNClassifySoundRequestClassOnce sync.Once
)

func getSNClassifySoundRequestClass() _SNClassifySoundRequestClass {
	SNClassifySoundRequestClassOnce.Do(func() {
		SNClassifySoundRequestClass = _SNClassifySoundRequestClass{objc.GetClass("SNClassifySoundRequest")}
	})
	return SNClassifySoundRequestClass
}

type _SNClassifySoundRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SNClassifySoundRequest */
// An interface definition for the [SNClassifySoundRequest] class.
type ISNClassifySoundRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SNClassifySoundRequest */
	// properties:
	KnownClassifications() []string
	OverlapFactor() float64
	SetOverlapFactor(value float64)
	WindowDuration() objc.IObject /* cross-framework: Time */
	SetWindowDuration(value objc.IObject /* cross-framework: Time */)
	WindowDurationConstraint() ISNTimeDurationConstraint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SNClassifySoundRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SNClassifySoundRequest */
// Alloc allocates a new instance without initialization.
func (sc _SNClassifySoundRequestClass) Alloc() SNClassifySoundRequest {
	rv := objc.Send[SNClassifySoundRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SNClassifySoundRequestClass) New() SNClassifySoundRequest {
	rv := objc.Send[SNClassifySoundRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SNClassifySoundRequest) Init() SNClassifySoundRequest {
	rv := objc.Send[SNClassifySoundRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SNClassifySoundRequest) Autorelease() SNClassifySoundRequest {
	rv := objc.Send[SNClassifySoundRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNClassifySoundRequest creates a new SNClassifySoundRequest instance.
func NewSNClassifySoundRequest() SNClassifySoundRequest {
	return getSNClassifySoundRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SNClassifySoundRequest */
// A request that classifies sound using a Core ML model.
//
// An represents a specific sound classification model. Analyze audio data with a sound classification model by: Creating an , either with the Sound Analysis model, or by providing your custom Core ML model. Adding the sound request to an or to process an audio file or stream, respectively. For more information about creating and using classify sound requests, see:


// A request that classifies sound using a Core ML model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest
type SNClassifySoundRequest struct {
	objectivec.Object
}

// SNClassifySoundRequestFrom constructs a [SNClassifySoundRequest] from an unsafe.Pointer.
//
// A request that classifies sound using a Core ML model.
func SNClassifySoundRequestFrom(ptr unsafe.Pointer) SNClassifySoundRequest {
	return SNClassifySoundRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SNClassifySoundRequest */

// Creates a request that uses the framework’s built-in sound classification model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(classifierIdentifier:)
func NewSNClassifySoundRequestWithClassifierIdentifierError(classifierIdentifier SNClassifierIdentifier /* typedef */, error_ unsafe.Pointer) SNClassifySoundRequest {
	instance := getSNClassifySoundRequestClass().Alloc()
	rv := objc.Send[SNClassifySoundRequest](instance.ID, objc.Sel("initWithClassifierIdentifier:error:"), classifierIdentifier, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSNClassifySoundRequestWithClassifierIdentifierError */


// Creates a request that uses a custom sound classification model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(mlModel:)
func NewSNClassifySoundRequestWithMLModelError(mlModel coreml.Model, error_ unsafe.Pointer) SNClassifySoundRequest {
	instance := getSNClassifySoundRequestClass().Alloc()
	rv := objc.Send[SNClassifySoundRequest](instance.ID, objc.Sel("initWithMLModel:error:"), mlModel, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSNClassifySoundRequestWithMLModelError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SNClassifySoundRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SNClassifySoundRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SNClassifySoundRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SNClassifySoundRequest */

// A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/knownClassifications
func (s_ SNClassifySoundRequest) KnownClassifications() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("knownClassifications"))
	return rv
}/* debug [instance_properties/getter]: knownClassifications */


// The amount of overlap between successive analysis windows when the model operates on a fixed-size audio block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/overlapFactor
func (s_ SNClassifySoundRequest) OverlapFactor() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("overlapFactor"))
	return rv
}/* debug [instance_properties/getter]: overlapFactor */


// The amount of overlap between successive analysis windows when the model operates on a fixed-size audio block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/overlapFactor
func (s_ SNClassifySoundRequest) SetOverlapFactor(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOverlapFactor:"), value)
}/* debug [instance_properties/setter]: overlapFactor */


// The duration of the audio buffer the request sends to the underlying sound classifier for each prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/windowDuration
func (s_ SNClassifySoundRequest) WindowDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](s_.ID, objc.Sel("windowDuration"))
	return rv
}/* debug [instance_properties/getter]: windowDuration */


// The duration of the audio buffer the request sends to the underlying sound classifier for each prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/windowDuration
func (s_ SNClassifySoundRequest) SetWindowDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWindowDuration:"), value)
}/* debug [instance_properties/setter]: windowDuration */


// A range or list of sound duration times the request’s underlying sound classifier supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/windowDurationConstraint-5aqvx
func (s_ SNClassifySoundRequest) WindowDurationConstraint() ISNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](s_.ID, objc.Sel("windowDurationConstraint"))
	return rv
}/* debug [instance_properties/getter]: windowDurationConstraint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SNClassifySoundRequest */


