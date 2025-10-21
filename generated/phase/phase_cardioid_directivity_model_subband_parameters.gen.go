// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHASECardioidDirectivityModelSubbandParameters] class.
var (
	PHASECardioidDirectivityModelSubbandParametersClass     _PHASECardioidDirectivityModelSubbandParametersClass
	PHASECardioidDirectivityModelSubbandParametersClassOnce sync.Once
)

func getPHASECardioidDirectivityModelSubbandParametersClass() _PHASECardioidDirectivityModelSubbandParametersClass {
	PHASECardioidDirectivityModelSubbandParametersClassOnce.Do(func() {
		PHASECardioidDirectivityModelSubbandParametersClass = _PHASECardioidDirectivityModelSubbandParametersClass{objc.GetClass("PHASECardioidDirectivityModelSubbandParameters")}
	})
	return PHASECardioidDirectivityModelSubbandParametersClass
}

type _PHASECardioidDirectivityModelSubbandParametersClass struct {
	class objc.Class
}

// An interface definition for the [PHASECardioidDirectivityModelSubbandParameters] class.
type IPHASECardioidDirectivityModelSubbandParameters interface {
	objectivec.IObject
}

// A data set that projects sound of a certain frequency outward in the shape of a heart.
//
// This class defines one subband in the class’s . Depending on the specific shape you define with and , you can attenuate sound focused at to the sides of the listener, while leaving the sound in front of or behind the listener unchanged.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters
type PHASECardioidDirectivityModelSubbandParameters struct {
	objectivec.Object
}

// PHASECardioidDirectivityModelSubbandParametersFrom constructs a [PHASECardioidDirectivityModelSubbandParameters] from an unsafe.Pointer.
//
// A data set that projects sound of a certain frequency outward in the shape of a heart.
func PHASECardioidDirectivityModelSubbandParametersFrom(ptr unsafe.Pointer) PHASECardioidDirectivityModelSubbandParameters {
	return PHASECardioidDirectivityModelSubbandParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASECardioidDirectivityModelSubbandParametersClass) Alloc() PHASECardioidDirectivityModelSubbandParameters {
	rv := objc.Send[PHASECardioidDirectivityModelSubbandParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASECardioidDirectivityModelSubbandParametersClass) New() PHASECardioidDirectivityModelSubbandParameters {
	rv := objc.Send[PHASECardioidDirectivityModelSubbandParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASECardioidDirectivityModelSubbandParameters) Init() PHASECardioidDirectivityModelSubbandParameters {
	rv := objc.Send[PHASECardioidDirectivityModelSubbandParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASECardioidDirectivityModelSubbandParameters) Autorelease() PHASECardioidDirectivityModelSubbandParameters {
	rv := objc.Send[PHASECardioidDirectivityModelSubbandParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASECardioidDirectivityModelSubbandParameters creates a new PHASECardioidDirectivityModelSubbandParameters instance.
func NewPHASECardioidDirectivityModelSubbandParameters() PHASECardioidDirectivityModelSubbandParameters {
	return getPHASECardioidDirectivityModelSubbandParametersClass().New()
}


// A frequency in the audio spectrum where the pattern and sharpness resonate most.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters/frequency
func (p_ PHASECardioidDirectivityModelSubbandParameters) Frequency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("frequency"))
	return rv
}


// SetFrequency sets the value of the frequency property.
// A frequency in the audio spectrum where the pattern and sharpness resonate most.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASECardioidDirectivityModelSubbandParameters/frequency
func (p_ PHASECardioidDirectivityModelSubbandParameters) SetFrequency(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFrequency:"), value)
}


