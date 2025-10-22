// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEMixerParameters] class.
var (
	PHASEMixerParametersClass     _PHASEMixerParametersClass
	PHASEMixerParametersClassOnce sync.Once
)

func getPHASEMixerParametersClass() _PHASEMixerParametersClass {
	PHASEMixerParametersClassOnce.Do(func() {
		PHASEMixerParametersClass = _PHASEMixerParametersClass{objc.GetClass("PHASEMixerParameters")}
	})
	return PHASEMixerParametersClass
}

type _PHASEMixerParametersClass struct {
	class objc.Class
}

// An interface definition for the [PHASEMixerParameters] class.
type IPHASEMixerParameters interface {
	objectivec.IObject
	AddAmbientMixerParametersWithIdentifierListener(identifier string, listener IPHASEListener)
	AddSpatialMixerParametersWithIdentifierSourceListener(identifier string, source IPHASESource, listener IPHASEListener)
}

// An object that specifies a mixer for sound events and orients them in 3D space.
//
// This class orients a sound event in 3D space relative to a listener. When you configure an ambient mixer’s orientation and a listener’s orientation, PHASE lowers the volume of the sound event if the two orientations point away from each other, and plays the sound at full volume if they point at each other. To add an instance of this class to a sound event, use the argument of a sound event’s initializer. Alternatively, PHASE can adjust a sound event’s loudness based on its distance from the listener in 3D space. By calling this class’s function, you supply a sound source that defines the location. For more information, see . Ambient sound events define only a listener and play with a consistent loudness, regardless of the listener’s position in the scene. To define a listener and select a particular ambient mixer that outputs the sound, call this class’s function.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerParameters
type PHASEMixerParameters struct {
	objectivec.Object
}

// PHASEMixerParametersFrom constructs a [PHASEMixerParameters] from an unsafe.Pointer.
//
// An object that specifies a mixer for sound events and orients them in 3D space.
func PHASEMixerParametersFrom(ptr unsafe.Pointer) PHASEMixerParameters {
	return PHASEMixerParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEMixerParametersClass) Alloc() PHASEMixerParameters {
	rv := objc.Send[PHASEMixerParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEMixerParametersClass) New() PHASEMixerParameters {
	rv := objc.Send[PHASEMixerParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMixerParameters) Init() PHASEMixerParameters {
	rv := objc.Send[PHASEMixerParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMixerParameters) Autorelease() PHASEMixerParameters {
	rv := objc.Send[PHASEMixerParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMixerParameters creates a new PHASEMixerParameters instance.
func NewPHASEMixerParameters() PHASEMixerParameters {
	return getPHASEMixerParametersClass().New()
}


// Adds runtime parameters for an ambient mixer.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerParameters/addAmbientMixerParameters(identifier:listener:)
func (p_ PHASEMixerParameters) AddAmbientMixerParametersWithIdentifierListener(identifier string, listener IPHASEListener) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addAmbientMixerParametersWithIdentifier:listener:"), objc.String(identifier), listener)
}

// Adds runtime parameters for a spatial mixer.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerParameters/addSpatialMixerParameters(identifier:source:listener:)
func (p_ PHASEMixerParameters) AddSpatialMixerParametersWithIdentifierSourceListener(identifier string, source IPHASESource, listener IPHASEListener) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSpatialMixerParametersWithIdentifier:source:listener:"), objc.String(identifier), source, listener)
}



