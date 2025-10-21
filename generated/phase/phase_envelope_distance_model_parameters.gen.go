// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEEnvelopeDistanceModelParameters] class.
var (
	PHASEEnvelopeDistanceModelParametersClass     _PHASEEnvelopeDistanceModelParametersClass
	PHASEEnvelopeDistanceModelParametersClassOnce sync.Once
)

func getPHASEEnvelopeDistanceModelParametersClass() _PHASEEnvelopeDistanceModelParametersClass {
	PHASEEnvelopeDistanceModelParametersClassOnce.Do(func() {
		PHASEEnvelopeDistanceModelParametersClass = _PHASEEnvelopeDistanceModelParametersClass{objc.GetClass("PHASEEnvelopeDistanceModelParameters")}
	})
	return PHASEEnvelopeDistanceModelParametersClass
}

type _PHASEEnvelopeDistanceModelParametersClass struct {
	class objc.Class
}

// An interface definition for the [PHASEEnvelopeDistanceModelParameters] class.
type IPHASEEnvelopeDistanceModelParameters interface {
	IPHASEDistanceModelParameters
}

// A graph of points and curves that shapes the volume of a sound over distance.
//
// This class provides an envelope that the app configures to dissipate the volume of a source’s sound with distance. The envelope describes a graph that the app configures using points and curves, where the input value is the distance between a sound source and the listener, and the output value is the sound’s volume.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeDistanceModelParameters
type PHASEEnvelopeDistanceModelParameters struct {
	PHASEDistanceModelParameters
}

// PHASEEnvelopeDistanceModelParametersFrom constructs a [PHASEEnvelopeDistanceModelParameters] from an unsafe.Pointer.
//
// A graph of points and curves that shapes the volume of a sound over distance.
func PHASEEnvelopeDistanceModelParametersFrom(ptr unsafe.Pointer) PHASEEnvelopeDistanceModelParameters {
	return PHASEEnvelopeDistanceModelParameters{
		PHASEDistanceModelParameters: PHASEDistanceModelParametersFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEEnvelopeDistanceModelParametersClass) Alloc() PHASEEnvelopeDistanceModelParameters {
	rv := objc.Send[PHASEEnvelopeDistanceModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEEnvelopeDistanceModelParametersClass) New() PHASEEnvelopeDistanceModelParameters {
	rv := objc.Send[PHASEEnvelopeDistanceModelParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEEnvelopeDistanceModelParameters) Init() PHASEEnvelopeDistanceModelParameters {
	rv := objc.Send[PHASEEnvelopeDistanceModelParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEEnvelopeDistanceModelParameters) Autorelease() PHASEEnvelopeDistanceModelParameters {
	rv := objc.Send[PHASEEnvelopeDistanceModelParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEEnvelopeDistanceModelParameters creates a new PHASEEnvelopeDistanceModelParameters instance.
func NewPHASEEnvelopeDistanceModelParameters() PHASEEnvelopeDistanceModelParameters {
	return getPHASEEnvelopeDistanceModelParametersClass().New()
}


// Creates the distance model parameters with an envelope.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeDistanceModelParameters/init(envelope:)
func NewPHASEEnvelopeDistanceModelParametersWithEnvelope(envelope unsafe.Pointer) PHASEEnvelopeDistanceModelParameters {
	instance := getPHASEEnvelopeDistanceModelParametersClass().Alloc()
	rv := objc.Send[PHASEEnvelopeDistanceModelParameters](instance.ID, objc.Sel("initWithEnvelope:"), envelope)
	rv.Autorelease()
	return rv
}


// An envelope that shapes sound dissipation over distance.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeDistanceModelParameters/envelope
func (p_ PHASEEnvelopeDistanceModelParameters) Envelope() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("envelope"))
	return rv
}


