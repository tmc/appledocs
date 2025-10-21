// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PHASECardioidDirectivityModelParameters] class.
type IPHASECardioidDirectivityModelParameters interface {
	IPHASEDirectivityModelParameters
}

// An object that directs sound in a heart-shaped curve surrounding a sound source.
//
// This class configures a particular frequency range in the audio spectrum that emits sound in an area defined by a mathematical cardioid. PHASE refers to each frequency segment along the audio spectrum as a . This class contains an array of that each can direct sound in a unique cardioid shape. The framework outputs a blend of a frequency’s adjacent subbands for all frequencies that lie outside of those specified in the array.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASECardioidDirectivityModelParametersClass) Alloc() PHASECardioidDirectivityModelParameters {
	rv := objc.Send[PHASECardioidDirectivityModelParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




