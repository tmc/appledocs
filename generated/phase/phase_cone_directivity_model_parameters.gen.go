// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEConeDirectivityModelParameters */


/* debug [class_header]: Header for PHASEConeDirectivityModelParameters */
// The class instance for the [PHASEConeDirectivityModelParameters] class.
var (
	PHASEConeDirectivityModelParametersClass     _PHASEConeDirectivityModelParametersClass
	PHASEConeDirectivityModelParametersClassOnce sync.Once
)

func getPHASEConeDirectivityModelParametersClass() _PHASEConeDirectivityModelParametersClass {
	PHASEConeDirectivityModelParametersClassOnce.Do(func() {
		PHASEConeDirectivityModelParametersClass = _PHASEConeDirectivityModelParametersClass{objc.GetClass("PHASEConeDirectivityModelParameters")}
	})
	return PHASEConeDirectivityModelParametersClass
}

type _PHASEConeDirectivityModelParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEConeDirectivityModelParameters */
// An interface definition for the [PHASEConeDirectivityModelParameters] class.
type IPHASEConeDirectivityModelParameters interface {
	IPHASEDirectivityModelParameters
	
/* debug [class_interface_properties]: Properties for PHASEConeDirectivityModelParameters */
	// properties:
	SubbandParameters() []PHASEConeDirectivityModelSubbandParameters
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEConeDirectivityModelParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEConeDirectivityModelParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASEConeDirectivityModelParametersClass) Alloc() PHASEConeDirectivityModelParameters {
	rv := objc.Send[PHASEConeDirectivityModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEConeDirectivityModelParametersClass) New() PHASEConeDirectivityModelParameters {
	rv := objc.Send[PHASEConeDirectivityModelParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEConeDirectivityModelParameters) Init() PHASEConeDirectivityModelParameters {
	rv := objc.Send[PHASEConeDirectivityModelParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEConeDirectivityModelParameters) Autorelease() PHASEConeDirectivityModelParameters {
	rv := objc.Send[PHASEConeDirectivityModelParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEConeDirectivityModelParameters creates a new PHASEConeDirectivityModelParameters instance.
func NewPHASEConeDirectivityModelParameters() PHASEConeDirectivityModelParameters {
	return getPHASEConeDirectivityModelParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEConeDirectivityModelParameters */
// An object that directs sound in a cone-shaped curve that extends from a sound source.
//
// This class determines that a particular frequency range in the audio spectrum emits sound in an area defined by a mathematical cone. PHASE refers to each frequency segment along the audio spectrum as a . This class contains an array of that each direct sound in a unique cone shape. The framework outputs a blend of a frequency’s adjacent subbands for all frequencies that lie outside of those specified in the array.


// An object that directs sound in a cone-shaped curve that extends from a sound source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelParameters
type PHASEConeDirectivityModelParameters struct {
	PHASEDirectivityModelParameters
}

// PHASEConeDirectivityModelParametersFrom constructs a [PHASEConeDirectivityModelParameters] from an unsafe.Pointer.
//
// An object that directs sound in a cone-shaped curve that extends from a sound source.
func PHASEConeDirectivityModelParametersFrom(ptr unsafe.Pointer) PHASEConeDirectivityModelParameters {
	return PHASEConeDirectivityModelParameters{
		PHASEDirectivityModelParameters: PHASEDirectivityModelParametersFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEConeDirectivityModelParameters */

// Creates an object that directs sound in a cone-shaped curve that extends from a sound source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelParameters/init(subbandParameters:)
func NewPHASEConeDirectivityModelParametersWithSubbandParameters(subbandParameters []PHASEConeDirectivityModelSubbandParameters) PHASEConeDirectivityModelParameters {
	instance := getPHASEConeDirectivityModelParametersClass().Alloc()
	rv := objc.Send[PHASEConeDirectivityModelParameters](instance.ID, objc.Sel("initWithSubbandParameters:"), subbandParameters)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEConeDirectivityModelParametersWithSubbandParameters */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEConeDirectivityModelParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEConeDirectivityModelParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEConeDirectivityModelParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEConeDirectivityModelParameters */

// An array of frequencies that describe varying sound emission across the spectrum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelParameters/subbandParameters
func (p_ PHASEConeDirectivityModelParameters) SubbandParameters() []PHASEConeDirectivityModelSubbandParameters {
	rv := objc.Send[[]PHASEConeDirectivityModelSubbandParameters](p_.ID, objc.Sel("subbandParameters"))
	return rv
}/* debug [instance_properties/getter]: subbandParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEConeDirectivityModelParameters */


