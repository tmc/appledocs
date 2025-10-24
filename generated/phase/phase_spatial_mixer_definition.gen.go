// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASESpatialMixerDefinition] class.
var (
	PHASESpatialMixerDefinitionClass     _PHASESpatialMixerDefinitionClass
	PHASESpatialMixerDefinitionClassOnce sync.Once
)

func getPHASESpatialMixerDefinitionClass() _PHASESpatialMixerDefinitionClass {
	PHASESpatialMixerDefinitionClassOnce.Do(func() {
		PHASESpatialMixerDefinitionClass = _PHASESpatialMixerDefinitionClass{objc.GetClass("PHASESpatialMixerDefinition")}
	})
	return PHASESpatialMixerDefinitionClass
}

type _PHASESpatialMixerDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASESpatialMixerDefinition] class.
type IPHASESpatialMixerDefinition interface {
	IPHASEMixerDefinition
	// properties:
	DistanceModelParameters() IPHASEDistanceModelParameters
	SetDistanceModelParameters(value IPHASEDistanceModelParameters)
	ListenerDirectivityModelParameters() IPHASEDirectivityModelParameters
	SetListenerDirectivityModelParameters(value IPHASEDirectivityModelParameters)
	SourceDirectivityModelParameters() IPHASEDirectivityModelParameters
	SetSourceDirectivityModelParameters(value IPHASEDirectivityModelParameters)
	SpatialPipeline() IPHASESpatialPipeline
	SetSpatialPipeline(value IPHASESpatialPipeline)
	// methods:
}

// An audio-layering object that produces environmental effects and plays sound with a 3D position and orientation.
//
// This class enables the app to define a relationship between a source and listener in six degrees of freedom: orientation (roll, pitch, yaw) and a 3D position (x, y, z). The framework plays back an audio source with (see ), direct path transmission effects and any combination of environmental effects, such as reverb (see ), and directivity (see ).  The result enables an app to implement directive point or omnidirectional sound sources — with or without direction, respectively — and volumetric sources with a defined shape. For a walkthrough of spatial mixing, see .


// An audio-layering object that produces environmental effects and plays sound with a 3D position and orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESpatialMixerDefinition
type PHASESpatialMixerDefinition struct {
	PHASEMixerDefinition
}

// PHASESpatialMixerDefinitionFrom constructs a [PHASESpatialMixerDefinition] from an unsafe.Pointer.
//
// An audio-layering object that produces environmental effects and plays sound with a 3D position and orientation.
func PHASESpatialMixerDefinitionFrom(ptr unsafe.Pointer) PHASESpatialMixerDefinition {
	return PHASESpatialMixerDefinition{
		PHASEMixerDefinition: PHASEMixerDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASESpatialMixerDefinitionClass) Alloc() PHASESpatialMixerDefinition {
	rv := objc.Send[PHASESpatialMixerDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASESpatialMixerDefinitionClass) New() PHASESpatialMixerDefinition {
	rv := objc.Send[PHASESpatialMixerDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESpatialMixerDefinition) Init() PHASESpatialMixerDefinition {
	rv := objc.Send[PHASESpatialMixerDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESpatialMixerDefinition) Autorelease() PHASESpatialMixerDefinition {
	rv := objc.Send[PHASESpatialMixerDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESpatialMixerDefinition creates a new PHASESpatialMixerDefinition instance.
func NewPHASESpatialMixerDefinition() PHASESpatialMixerDefinition {
	return getPHASESpatialMixerDefinitionClass().New()
}



// An effect that changes sound as it carries over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASESpatialMixerDefinition) DistanceModelParameters() IPHASEDistanceModelParameters {
	rv := objc.Send[PHASEDistanceModelParameters](p_.ID, objc.Sel("distanceModelParameters"))
	return rv
}


// An effect that changes sound as it carries over a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/distancemodelparameters
func (p_ PHASESpatialMixerDefinition) SetDistanceModelParameters(value IPHASEDistanceModelParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDistanceModelParameters:"), value)
}


// A data set that determines how well the listener hears depending on its direction relative to a sound source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/listenerdirectivitymodelparameters
func (p_ PHASESpatialMixerDefinition) ListenerDirectivityModelParameters() IPHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](p_.ID, objc.Sel("listenerDirectivityModelParameters"))
	return rv
}


// A data set that determines how well the listener hears depending on its direction relative to a sound source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/listenerdirectivitymodelparameters
func (p_ PHASESpatialMixerDefinition) SetListenerDirectivityModelParameters(value IPHASEDirectivityModelParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setListenerDirectivityModelParameters:"), value)
}


// A data set that directs sound such that it’s louder when directed at the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/sourcedirectivitymodelparameters
func (p_ PHASESpatialMixerDefinition) SourceDirectivityModelParameters() IPHASEDirectivityModelParameters {
	rv := objc.Send[PHASEDirectivityModelParameters](p_.ID, objc.Sel("sourceDirectivityModelParameters"))
	return rv
}


// A data set that directs sound such that it’s louder when directed at the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/sourcedirectivitymodelparameters
func (p_ PHASESpatialMixerDefinition) SetSourceDirectivityModelParameters(value IPHASEDirectivityModelParameters) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSourceDirectivityModelParameters:"), value)
}


// An object that adds sound layers for environmental effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/spatialpipeline
func (p_ PHASESpatialMixerDefinition) SpatialPipeline() IPHASESpatialPipeline {
	rv := objc.Send[PHASESpatialPipeline](p_.ID, objc.Sel("spatialPipeline"))
	return rv
}


// An object that adds sound layers for environmental effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasespatialmixerdefinition/spatialpipeline
func (p_ PHASESpatialMixerDefinition) SetSpatialPipeline(value IPHASESpatialPipeline) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSpatialPipeline:"), value)
}



