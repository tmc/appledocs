// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEEnvelopeDistanceModelParameters */


/* debug [class_header]: Header for PHASEEnvelopeDistanceModelParameters */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEEnvelopeDistanceModelParameters */
// An interface definition for the [PHASEEnvelopeDistanceModelParameters] class.
type IPHASEEnvelopeDistanceModelParameters interface {
	IPHASEDistanceModelParameters
	
/* debug [class_interface_properties]: Properties for PHASEEnvelopeDistanceModelParameters */
	// properties:
	Envelope() IPHASEEnvelope
	RolloffFactor() float64
	SetRolloffFactor(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEEnvelopeDistanceModelParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEEnvelopeDistanceModelParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASEEnvelopeDistanceModelParametersClass) Alloc() PHASEEnvelopeDistanceModelParameters {
	rv := objc.Send[PHASEEnvelopeDistanceModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEEnvelopeDistanceModelParameters */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEEnvelopeDistanceModelParameters */

// Creates the distance model parameters with an envelope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeDistanceModelParameters/init(envelope:)
func NewPHASEEnvelopeDistanceModelParametersWithEnvelope(envelope IPHASEEnvelope) PHASEEnvelopeDistanceModelParameters {
	instance := getPHASEEnvelopeDistanceModelParametersClass().Alloc()
	rv := objc.Send[PHASEEnvelopeDistanceModelParameters](instance.ID, objc.Sel("initWithEnvelope:"), envelope)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEEnvelopeDistanceModelParametersWithEnvelope */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEEnvelopeDistanceModelParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEEnvelopeDistanceModelParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEEnvelopeDistanceModelParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEEnvelopeDistanceModelParameters */

// An envelope that shapes sound dissipation over distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeDistanceModelParameters/envelope
func (p_ PHASEEnvelopeDistanceModelParameters) Envelope() IPHASEEnvelope {
	rv := objc.Send[PHASEEnvelope](p_.ID, objc.Sel("envelope"))
	return rv
}/* debug [instance_properties/getter]: envelope */


// A value that fades specific frequencies over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeometricspreadingdistancemodelparameters/rollofffactor
func (p_ PHASEEnvelopeDistanceModelParameters) RolloffFactor() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rolloffFactor"))
	return rv
}/* debug [instance_properties/getter]: rolloffFactor */


// A value that fades specific frequencies over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeometricspreadingdistancemodelparameters/rollofffactor
func (p_ PHASEEnvelopeDistanceModelParameters) SetRolloffFactor(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRolloffFactor:"), value)
}/* debug [instance_properties/setter]: rolloffFactor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEEnvelopeDistanceModelParameters */


