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
	SetInnerAngleOuterAngle(innerAngle unsafe.Pointer, outerAngle unsafe.Pointer)
}

// A data set that projects sound of a certain frequency outward in the shape of a cone.
//
// This class defines one subband in the class’s . The inner and outer angles you define with describe a cone that directs sound of a given toward the listener. The cone’s point rests at the 3D position of the sound source. The framework adjusts the volume of the sound according to location of the listener in the 3D scene: If the listener positions in an area outside of the subband’s , the sound emanates from the source at the volume defined by . If the listener positions inside the area defined by , the sound emanates from the source at maximum volume. If the listener positions in between the outer and inner angles, the framework blends the volume to a value between and the maximum.
//
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



// Configures a focus area for cone-based sound directivity.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/setAngles(innerAngle:outerAngle:)
func (p_ PHASEConeDirectivityModelSubbandParameters) SetInnerAngleOuterAngle(innerAngle unsafe.Pointer, outerAngle unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInnerAngle:outerAngle:"), innerAngle, outerAngle)
}

// A frequency in the audio spectrum where the subband resonates most.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/frequency
func (p_ PHASEConeDirectivityModelSubbandParameters) Frequency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("frequency"))
	return rv
}


// SetFrequency sets the value of the frequency property.
// A frequency in the audio spectrum where the subband resonates most.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/frequency
func (p_ PHASEConeDirectivityModelSubbandParameters) SetFrequency(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFrequency:"), value)
}

// An angle, in degrees, that determines the size of the audio emitting area inside the cone.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/innerAngle
func (p_ PHASEConeDirectivityModelSubbandParameters) InnerAngle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("innerAngle"))
	return rv
}

// An angle, in degrees, that determines the size of the audio emitting area outside the cone.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/outerAngle
func (p_ PHASEConeDirectivityModelSubbandParameters) OuterAngle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("outerAngle"))
	return rv
}

// The loudness of the audio the outside area of the cone emits.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/outerGain
func (p_ PHASEConeDirectivityModelSubbandParameters) OuterGain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("outerGain"))
	return rv
}


// SetOuterGain sets the value of the outerGain property.
// The loudness of the audio the outside area of the cone emits.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEConeDirectivityModelSubbandParameters/outerGain
func (p_ PHASEConeDirectivityModelSubbandParameters) SetOuterGain(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOuterGain:"), value)
}


