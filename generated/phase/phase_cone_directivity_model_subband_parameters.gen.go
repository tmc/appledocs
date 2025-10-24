// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEConeDirectivityModelSubbandParameters] class.
var (
	PHASEConeDirectivityModelSubbandParametersClass     _PHASEConeDirectivityModelSubbandParametersClass
	PHASEConeDirectivityModelSubbandParametersClassOnce sync.Once
)

func getPHASEConeDirectivityModelSubbandParametersClass() _PHASEConeDirectivityModelSubbandParametersClass {
	PHASEConeDirectivityModelSubbandParametersClassOnce.Do(func() {
		PHASEConeDirectivityModelSubbandParametersClass = _PHASEConeDirectivityModelSubbandParametersClass{objc.GetClass("PHASEConeDirectivityModelSubbandParameters")}
	})
	return PHASEConeDirectivityModelSubbandParametersClass
}

type _PHASEConeDirectivityModelSubbandParametersClass struct {
	class objc.Class
}

// An interface definition for the [PHASEConeDirectivityModelSubbandParameters] class.
type IPHASEConeDirectivityModelSubbandParameters interface {
	objectivec.IObject
	// properties:
	Frequency() float64
	SetFrequency(value float64)
	InnerAngle() float64
	SetInnerAngle(value float64)
	OuterAngle() float64
	SetOuterAngle(value float64)
	OuterGain() float64
	SetOuterGain(value float64)
	// methods:
}

// A data set that projects sound of a certain frequency outward in the shape of a cone.
//
// This class defines one subband in the class’s . The inner and outer angles you define with describe a cone that directs sound of a given toward the listener. The cone’s point rests at the 3D position of the sound source. The framework adjusts the volume of the sound according to location of the listener in the 3D scene: If the listener positions in an area outside of the subband’s , the sound emanates from the source at the volume defined by . If the listener positions inside the area defined by , the sound emanates from the source at maximum volume. If the listener positions in between the outer and inner angles, the framework blends the volume to a value between and the maximum.


// A data set that projects sound of a certain frequency outward in the shape of a cone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters
type PHASEConeDirectivityModelSubbandParameters struct {
	objectivec.Object
}

// PHASEConeDirectivityModelSubbandParametersFrom constructs a [PHASEConeDirectivityModelSubbandParameters] from an unsafe.Pointer.
//
// A data set that projects sound of a certain frequency outward in the shape of a cone.
func PHASEConeDirectivityModelSubbandParametersFrom(ptr unsafe.Pointer) PHASEConeDirectivityModelSubbandParameters {
	return PHASEConeDirectivityModelSubbandParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEConeDirectivityModelSubbandParametersClass) Alloc() PHASEConeDirectivityModelSubbandParameters {
	rv := objc.Send[PHASEConeDirectivityModelSubbandParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEConeDirectivityModelSubbandParametersClass) New() PHASEConeDirectivityModelSubbandParameters {
	rv := objc.Send[PHASEConeDirectivityModelSubbandParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEConeDirectivityModelSubbandParameters) Init() PHASEConeDirectivityModelSubbandParameters {
	rv := objc.Send[PHASEConeDirectivityModelSubbandParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEConeDirectivityModelSubbandParameters) Autorelease() PHASEConeDirectivityModelSubbandParameters {
	rv := objc.Send[PHASEConeDirectivityModelSubbandParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEConeDirectivityModelSubbandParameters creates a new PHASEConeDirectivityModelSubbandParameters instance.
func NewPHASEConeDirectivityModelSubbandParameters() PHASEConeDirectivityModelSubbandParameters {
	return getPHASEConeDirectivityModelSubbandParametersClass().New()
}



// A frequency in the audio spectrum where the pattern and sharpness resonate most.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasecardioiddirectivitymodelsubbandparameters/frequency
func (p_ PHASEConeDirectivityModelSubbandParameters) Frequency() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("frequency"))
	return rv
}


// A frequency in the audio spectrum where the pattern and sharpness resonate most.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasecardioiddirectivitymodelsubbandparameters/frequency
func (p_ PHASEConeDirectivityModelSubbandParameters) SetFrequency(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFrequency:"), value)
}


// An angle, in degrees, that determines the size of the audio emitting area inside the cone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseconedirectivitymodelsubbandparameters/innerangle
func (p_ PHASEConeDirectivityModelSubbandParameters) InnerAngle() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("innerAngle"))
	return rv
}


// An angle, in degrees, that determines the size of the audio emitting area inside the cone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseconedirectivitymodelsubbandparameters/innerangle
func (p_ PHASEConeDirectivityModelSubbandParameters) SetInnerAngle(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInnerAngle:"), value)
}


// An angle, in degrees, that determines the size of the audio emitting area outside the cone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseconedirectivitymodelsubbandparameters/outerangle
func (p_ PHASEConeDirectivityModelSubbandParameters) OuterAngle() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("outerAngle"))
	return rv
}


// An angle, in degrees, that determines the size of the audio emitting area outside the cone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseconedirectivitymodelsubbandparameters/outerangle
func (p_ PHASEConeDirectivityModelSubbandParameters) SetOuterAngle(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOuterAngle:"), value)
}


// The loudness of the audio the outside area of the cone emits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseconedirectivitymodelsubbandparameters/outergain
func (p_ PHASEConeDirectivityModelSubbandParameters) OuterGain() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("outerGain"))
	return rv
}


// The loudness of the audio the outside area of the cone emits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseconedirectivitymodelsubbandparameters/outergain
func (p_ PHASEConeDirectivityModelSubbandParameters) SetOuterGain(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOuterGain:"), value)
}



