// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A request that classifies sound using a Core ML model.
//
// An represents a specific sound classification model. Analyze audio data with a sound classification model by: Creating an , either with the Sound Analysis model, or by providing your custom Core ML model. Adding the sound request to an or to process an audio file or stream, respectively. For more information about creating and using classify sound requests, see:
//
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
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(classifierIdentifier:)
func NewSNClassifySoundRequestWithClassifierIdentifierError(classifierIdentifier unsafe.Pointer, error_ unsafe.Pointer) SNClassifySoundRequest {
	instance := getSNClassifySoundRequestClass().Alloc()
	rv := objc.Send[SNClassifySoundRequest](instance.ID, objc.Sel("initWithClassifierIdentifier:error:"), classifierIdentifier, error_)
	rv.Autorelease()
	return rv
}

// Creates a request that uses a custom sound classification model.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/init(mlModel:)
func NewSNClassifySoundRequestWithMLModelError(mlModel unsafe.Pointer, error_ unsafe.Pointer) SNClassifySoundRequest {
	instance := getSNClassifySoundRequestClass().Alloc()
	rv := objc.Send[SNClassifySoundRequest](instance.ID, objc.Sel("initWithMLModel:error:"), mlModel, error_)
	rv.Autorelease()
	return rv
}


// A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/knownClassifications
func (s_ SNClassifySoundRequest) KnownClassifications() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("knownClassifications"))
	return rv
}

// The amount of overlap between successive analysis windows when the model operates on a fixed-size audio block.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/overlapFactor
func (s_ SNClassifySoundRequest) OverlapFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("overlapFactor"))
	return rv
}


// SetOverlapFactor sets the value of the overlapFactor property.
// The amount of overlap between successive analysis windows when the model operates on a fixed-size audio block.

//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassifySoundRequest/overlapFactor
func (s_ SNClassifySoundRequest) SetOverlapFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setOverlapFactor:"), value)
}

