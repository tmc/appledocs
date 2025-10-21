// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PHASEGeometricSpreadingDistanceModelParameters] class.
type IPHASEGeometricSpreadingDistanceModelParameters interface {
	IPHASEDistanceModelParameters
}

// An object that dissipates sound frequencies over distance.
//
// This class implements a effect — a strategy that aims to model the real-world manner in which sound changes with distance. When the distance between a sound and listener changes, the roll-off effect dissipates certain audio frequencies more than others.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEGeometricSpreadingDistanceModelParametersClass) Alloc() PHASEGeometricSpreadingDistanceModelParameters {
	rv := objc.Send[PHASEGeometricSpreadingDistanceModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A value that fades specific frequencies over a distance.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeometricSpreadingDistanceModelParameters/rolloffFactor
func (p_ PHASEGeometricSpreadingDistanceModelParameters) RolloffFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("rolloffFactor"))
	return rv
}


// SetRolloffFactor sets the value of the rolloffFactor property.
// A value that fades specific frequencies over a distance.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeometricSpreadingDistanceModelParameters/rolloffFactor
func (p_ PHASEGeometricSpreadingDistanceModelParameters) SetRolloffFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRolloffFactor:"), value)
}

// An effect that changes sound as it carries over a distance.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASEGeometricSpreadingDistanceModelParameters) DistanceModelParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("distanceModelParameters"))
	return rv
}


// SetDistanceModelParameters sets the value of the distanceModelParameters property.
// An effect that changes sound as it carries over a distance.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASEGeometricSpreadingDistanceModelParameters) SetDistanceModelParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDistanceModelParameters:"), value)
}



