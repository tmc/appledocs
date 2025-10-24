// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEGeometricSpreadingDistanceModelParameters */


/* debug [class_header]: Header for PHASEGeometricSpreadingDistanceModelParameters */
// The class instance for the [PHASEGeometricSpreadingDistanceModelParameters] class.
var (
	PHASEGeometricSpreadingDistanceModelParametersClass     _PHASEGeometricSpreadingDistanceModelParametersClass
	PHASEGeometricSpreadingDistanceModelParametersClassOnce sync.Once
)

func getPHASEGeometricSpreadingDistanceModelParametersClass() _PHASEGeometricSpreadingDistanceModelParametersClass {
	PHASEGeometricSpreadingDistanceModelParametersClassOnce.Do(func() {
		PHASEGeometricSpreadingDistanceModelParametersClass = _PHASEGeometricSpreadingDistanceModelParametersClass{objc.GetClass("PHASEGeometricSpreadingDistanceModelParameters")}
	})
	return PHASEGeometricSpreadingDistanceModelParametersClass
}

type _PHASEGeometricSpreadingDistanceModelParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEGeometricSpreadingDistanceModelParameters */
// An interface definition for the [PHASEGeometricSpreadingDistanceModelParameters] class.
type IPHASEGeometricSpreadingDistanceModelParameters interface {
	IPHASEDistanceModelParameters
	
/* debug [class_interface_properties]: Properties for PHASEGeometricSpreadingDistanceModelParameters */
	// properties:
	RolloffFactor() float64
	SetRolloffFactor(value float64)
	DistanceModelParameters() IPHASEDistanceModelParameters
	SetDistanceModelParameters(value IPHASEDistanceModelParameters)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEGeometricSpreadingDistanceModelParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEGeometricSpreadingDistanceModelParameters */
// Alloc allocates a new instance without initialization.
func (pc _PHASEGeometricSpreadingDistanceModelParametersClass) Alloc() PHASEGeometricSpreadingDistanceModelParameters {
	rv := objc.Send[PHASEGeometricSpreadingDistanceModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEGeometricSpreadingDistanceModelParametersClass) New() PHASEGeometricSpreadingDistanceModelParameters {
	rv := objc.Send[PHASEGeometricSpreadingDistanceModelParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEGeometricSpreadingDistanceModelParameters) Init() PHASEGeometricSpreadingDistanceModelParameters {
	rv := objc.Send[PHASEGeometricSpreadingDistanceModelParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEGeometricSpreadingDistanceModelParameters) Autorelease() PHASEGeometricSpreadingDistanceModelParameters {
	rv := objc.Send[PHASEGeometricSpreadingDistanceModelParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEGeometricSpreadingDistanceModelParameters creates a new PHASEGeometricSpreadingDistanceModelParameters instance.
func NewPHASEGeometricSpreadingDistanceModelParameters() PHASEGeometricSpreadingDistanceModelParameters {
	return getPHASEGeometricSpreadingDistanceModelParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEGeometricSpreadingDistanceModelParameters */
// An object that dissipates sound frequencies over distance.
//
// This class implements a effect — a strategy that aims to model the real-world manner in which sound changes with distance. When the distance between a sound and listener changes, the roll-off effect dissipates certain audio frequencies more than others.


// An object that dissipates sound frequencies over distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeometricSpreadingDistanceModelParameters
type PHASEGeometricSpreadingDistanceModelParameters struct {
	PHASEDistanceModelParameters
}

// PHASEGeometricSpreadingDistanceModelParametersFrom constructs a [PHASEGeometricSpreadingDistanceModelParameters] from an unsafe.Pointer.
//
// An object that dissipates sound frequencies over distance.
func PHASEGeometricSpreadingDistanceModelParametersFrom(ptr unsafe.Pointer) PHASEGeometricSpreadingDistanceModelParameters {
	return PHASEGeometricSpreadingDistanceModelParameters{
		PHASEDistanceModelParameters: PHASEDistanceModelParametersFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEGeometricSpreadingDistanceModelParameters */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEGeometricSpreadingDistanceModelParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEGeometricSpreadingDistanceModelParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEGeometricSpreadingDistanceModelParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEGeometricSpreadingDistanceModelParameters */

// A value that fades specific frequencies over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeometricSpreadingDistanceModelParameters/rolloffFactor
func (p_ PHASEGeometricSpreadingDistanceModelParameters) RolloffFactor() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rolloffFactor"))
	return rv
}/* debug [instance_properties/getter]: rolloffFactor */


// A value that fades specific frequencies over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeometricSpreadingDistanceModelParameters/rolloffFactor
func (p_ PHASEGeometricSpreadingDistanceModelParameters) SetRolloffFactor(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRolloffFactor:"), value)
}/* debug [instance_properties/setter]: rolloffFactor */


// An effect that changes sound as it carries over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASEGeometricSpreadingDistanceModelParameters) DistanceModelParameters() IPHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](p_.ID, objc.Sel("distanceModelParameters"))
	return rv
}/* debug [instance_properties/getter]: distanceModelParameters */


// An effect that changes sound as it carries over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASEGeometricSpreadingDistanceModelParameters) SetDistanceModelParameters(value IPHASEDistanceModelParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDistanceModelParameters:"), value)
}/* debug [instance_properties/setter]: distanceModelParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEGeometricSpreadingDistanceModelParameters */


