// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASECardioidDirectivityModelParameters */


/* debug [class_header]: Header for PHASECardioidDirectivityModelParameters */
// The class instance for the [PHASECardioidDirectivityModelParameters] class.
var (
	PHASECardioidDirectivityModelParametersClass     _PHASECardioidDirectivityModelParametersClass
	PHASECardioidDirectivityModelParametersClassOnce sync.Once
)

func getPHASECardioidDirectivityModelParametersClass() _PHASECardioidDirectivityModelParametersClass {
	PHASECardioidDirectivityModelParametersClassOnce.Do(func() {
		PHASECardioidDirectivityModelParametersClass = _PHASECardioidDirectivityModelParametersClass{objc.GetClass("PHASECardioidDirectivityModelParameters")}
	})
	return PHASECardioidDirectivityModelParametersClass
}

type _PHASECardioidDirectivityModelParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASECardioidDirectivityModelParameters */
// An interface definition for the [PHASECardioidDirectivityModelParameters] class.
type IPHASECardioidDirectivityModelParameters interface {
	IPHASEDirectivityModelParameters
	
/* debug [class_interface_properties]: Properties for PHASECardioidDirectivityModelParameters */
	// properties:
	SubbandParameters() []PHASECardioidDirectivityModelSubbandParameters
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASECardioidDirectivityModelParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASECardioidDirectivityModelParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASECardioidDirectivityModelParametersClass) Alloc() PHASECardioidDirectivityModelParameters {
	rv := objc.Send[PHASECardioidDirectivityModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASECardioidDirectivityModelParametersClass) New() PHASECardioidDirectivityModelParameters {
	rv := objc.Send[PHASECardioidDirectivityModelParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASECardioidDirectivityModelParameters) Init() PHASECardioidDirectivityModelParameters {
	rv := objc.Send[PHASECardioidDirectivityModelParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASECardioidDirectivityModelParameters) Autorelease() PHASECardioidDirectivityModelParameters {
	rv := objc.Send[PHASECardioidDirectivityModelParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASECardioidDirectivityModelParameters creates a new PHASECardioidDirectivityModelParameters instance.
func NewPHASECardioidDirectivityModelParameters() PHASECardioidDirectivityModelParameters {
	return getPHASECardioidDirectivityModelParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASECardioidDirectivityModelParameters */
// An object that directs sound in a heart-shaped curve surrounding a sound source.
//
// This class configures a particular frequency range in the audio spectrum that emits sound in an area defined by a mathematical cardioid. PHASE refers to each frequency segment along the audio spectrum as a . This class contains an array of that each can direct sound in a unique cardioid shape. The framework outputs a blend of a frequency’s adjacent subbands for all frequencies that lie outside of those specified in the array.


// An object that directs sound in a heart-shaped curve surrounding a sound source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelParameters
type PHASECardioidDirectivityModelParameters struct {
	PHASEDirectivityModelParameters
}

// PHASECardioidDirectivityModelParametersFrom constructs a [PHASECardioidDirectivityModelParameters] from an unsafe.Pointer.
//
// An object that directs sound in a heart-shaped curve surrounding a sound source.
func PHASECardioidDirectivityModelParametersFrom(ptr unsafe.Pointer) PHASECardioidDirectivityModelParameters {
	return PHASECardioidDirectivityModelParameters{
		PHASEDirectivityModelParameters: PHASEDirectivityModelParametersFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASECardioidDirectivityModelParameters */

// Creates an object that directs sound in a heart-shaped curve surrounding a sound source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelParameters/init(subbandParameters:)
func NewPHASECardioidDirectivityModelParametersWithSubbandParameters(subbandParameters []PHASECardioidDirectivityModelSubbandParameters) PHASECardioidDirectivityModelParameters {
	instance := getPHASECardioidDirectivityModelParametersClass().Alloc()
	rv := objc.Send[PHASECardioidDirectivityModelParameters](instance.ID, objc.Sel("initWithSubbandParameters:"), subbandParameters)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASECardioidDirectivityModelParametersWithSubbandParameters */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASECardioidDirectivityModelParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASECardioidDirectivityModelParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASECardioidDirectivityModelParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASECardioidDirectivityModelParameters */

// An array of frequencies that describe varying sound emission across the spectrum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelParameters/subbandParameters
func (p_ PHASECardioidDirectivityModelParameters) SubbandParameters() []PHASECardioidDirectivityModelSubbandParameters {
	rv := objc.Send[[]PHASECardioidDirectivityModelSubbandParameters](p_.ID, objc.Sel("subbandParameters"))
	return rv
}/* debug [instance_properties/getter]: subbandParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASECardioidDirectivityModelParameters */


