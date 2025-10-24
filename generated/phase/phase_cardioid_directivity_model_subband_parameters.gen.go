// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Frequency() float64
	SetFrequency(value float64)
	Pattern() float64
	SetPattern(value float64)
	Sharpness() float64
	SetSharpness(value float64)
	// methods:
}

// A data set that projects sound of a certain frequency outward in the shape of a heart.
//
// This class defines one subband in the class’s . Depending on the specific shape you define with and , you can attenuate sound focused at to the sides of the listener, while leaving the sound in front of or behind the listener unchanged.


// A data set that projects sound of a certain frequency outward in the shape of a heart.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasecardioiddirectivitymodelsubbandparameters/frequency
func (p_ PHASECardioidDirectivityModelSubbandParameters) Frequency() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("frequency"))
	return rv
}


// A frequency in the audio spectrum where the pattern and sharpness resonate most.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasecardioiddirectivitymodelsubbandparameters/frequency
func (p_ PHASECardioidDirectivityModelSubbandParameters) SetFrequency(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFrequency:"), value)
}


// A shape that determines the direction of sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasecardioiddirectivitymodelsubbandparameters/pattern
func (p_ PHASECardioidDirectivityModelSubbandParameters) Pattern() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("pattern"))
	return rv
}


// A shape that determines the direction of sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasecardioiddirectivitymodelsubbandparameters/pattern
func (p_ PHASECardioidDirectivityModelSubbandParameters) SetPattern(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPattern:"), value)
}


// The amount that the shape overlaps with bordering subbands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasecardioiddirectivitymodelsubbandparameters/sharpness
func (p_ PHASECardioidDirectivityModelSubbandParameters) Sharpness() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("sharpness"))
	return rv
}


// The amount that the shape overlaps with bordering subbands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasecardioiddirectivitymodelsubbandparameters/sharpness
func (p_ PHASECardioidDirectivityModelSubbandParameters) SetSharpness(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSharpness:"), value)
}



