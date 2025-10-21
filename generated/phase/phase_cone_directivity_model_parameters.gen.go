// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PHASEConeDirectivityModelParameters] class.
type IPHASEConeDirectivityModelParameters interface {
	IPHASEDirectivityModelParameters
}

// An object that directs sound in a cone-shaped curve that extends from a sound source.
//
// This class determines that a particular frequency range in the audio spectrum emits sound in an area defined by a mathematical cone. PHASE refers to each frequency segment along the audio spectrum as a . This class contains an array of that each direct sound in a unique cone shape. The framework outputs a blend of a frequency’s adjacent subbands for all frequencies that lie outside of those specified in the array.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEConeDirectivityModelParametersClass) Alloc() PHASEConeDirectivityModelParameters {
	rv := objc.Send[PHASEConeDirectivityModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




