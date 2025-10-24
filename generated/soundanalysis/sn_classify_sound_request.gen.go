// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SNClassifySoundRequest] class.
type ISNClassifySoundRequest interface {
	objectivec.IObject
	// properties:
	KnownClassifications() objc.IObject /* cross-framework: NSString */
	SetKnownClassifications(value objc.IObject /* cross-framework: NSString */)
	OverlapFactor() float64
	SetOverlapFactor(value float64)
	WindowDuration() objc.IObject /* cross-framework: Time */
	SetWindowDuration(value objc.IObject /* cross-framework: Time */)
	WindowDurationConstraint() ISNTimeDurationConstraint
	SetWindowDurationConstraint(value ISNTimeDurationConstraint)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (sc _SNClassifySoundRequestClass) Alloc() SNClassifySoundRequest {
	rv := objc.Send[SNClassifySoundRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a request that uses the framework’s built-in sound classification model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(classifierIdentifier:)
func NewSNClassifySoundRequestWithClassifierIdentifierError(classifierIdentifier unsafe.Pointer, error_ unsafe.Pointer) SNClassifySoundRequest {
	instance := getSNClassifySoundRequestClass().Alloc()
	rv := objc.Send[SNClassifySoundRequest](instance.ID, objc.Sel("initWithClassifierIdentifier:error:"), classifierIdentifier, error_)
	rv.Autorelease()
	return rv
}


// Creates a request that uses a custom sound classification model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(mlModel:)
func NewSNClassifySoundRequestWithMLModelError(mlModel objc.IObject /* cross-framework: Model */, error_ unsafe.Pointer) SNClassifySoundRequest {
	instance := getSNClassifySoundRequestClass().Alloc()
	rv := objc.Send[SNClassifySoundRequest](instance.ID, objc.Sel("initWithMLModel:error:"), mlModel, error_)
	rv.Autorelease()
	return rv
}



// A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/knownclassifications
func (s_ SNClassifySoundRequest) KnownClassifications() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("knownClassifications"))
	return rv
}


// A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/knownclassifications
func (s_ SNClassifySoundRequest) SetKnownClassifications(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKnownClassifications:"), value)
}


// The amount of overlap between successive analysis windows when the model operates on a fixed-size audio block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/overlapfactor
func (s_ SNClassifySoundRequest) OverlapFactor() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("overlapFactor"))
	return rv
}


// The amount of overlap between successive analysis windows when the model operates on a fixed-size audio block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/overlapfactor
func (s_ SNClassifySoundRequest) SetOverlapFactor(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOverlapFactor:"), value)
}


// The duration of the audio buffer the request sends to the underlying sound classifier for each prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/windowduration
func (s_ SNClassifySoundRequest) WindowDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](s_.ID, objc.Sel("windowDuration"))
	return rv
}


// The duration of the audio buffer the request sends to the underlying sound classifier for each prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/windowduration
func (s_ SNClassifySoundRequest) SetWindowDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWindowDuration:"), value)
}


// A range or list of sound duration times the request’s underlying sound classifier supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/windowdurationconstraint-5no60
func (s_ SNClassifySoundRequest) WindowDurationConstraint() ISNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](s_.ID, objc.Sel("windowDurationConstraint"))
	return rv
}


// A range or list of sound duration times the request’s underlying sound classifier supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/windowdurationconstraint-5no60
func (s_ SNClassifySoundRequest) SetWindowDurationConstraint(value ISNTimeDurationConstraint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setWindowDurationConstraint:"), value)
}


