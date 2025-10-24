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
	// properties:
	Envelope() IPHASEEnvelope
	SetEnvelope(value IPHASEEnvelope)
	RolloffFactor() float64
	SetRolloffFactor(value float64)
	// methods:
}

// A graph of points and curves that shapes the volume of a sound over distance.
//
// This class provides an envelope that the app configures to dissipate the volume of a source’s sound with distance. The envelope describes a graph that the app configures using points and curves, where the input value is the distance between a sound source and the listener, and the output value is the sound’s volume.


// A graph of points and curves that shapes the volume of a sound over distance.
//
// [Full Topic]
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



// An envelope that shapes sound dissipation over distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseenvelopedistancemodelparameters/envelope
func (p_ PHASEEnvelopeDistanceModelParameters) Envelope() IPHASEEnvelope {
	rv := objc.Send[PHASEEnvelope](p_.ID, objc.Sel("envelope"))
	return rv
}


// An envelope that shapes sound dissipation over distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseenvelopedistancemodelparameters/envelope
func (p_ PHASEEnvelopeDistanceModelParameters) SetEnvelope(value IPHASEEnvelope) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnvelope:"), value)
}


// A value that fades specific frequencies over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeometricspreadingdistancemodelparameters/rollofffactor
func (p_ PHASEEnvelopeDistanceModelParameters) RolloffFactor() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rolloffFactor"))
	return rv
}


// A value that fades specific frequencies over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeometricspreadingdistancemodelparameters/rollofffactor
func (p_ PHASEEnvelopeDistanceModelParameters) SetRolloffFactor(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRolloffFactor:"), value)
}



