// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEGeneratorNodeDefinition] class.
var (
	PHASEGeneratorNodeDefinitionClass     _PHASEGeneratorNodeDefinitionClass
	PHASEGeneratorNodeDefinitionClassOnce sync.Once
)

func getPHASEGeneratorNodeDefinitionClass() _PHASEGeneratorNodeDefinitionClass {
	PHASEGeneratorNodeDefinitionClassOnce.Do(func() {
		PHASEGeneratorNodeDefinitionClass = _PHASEGeneratorNodeDefinitionClass{objc.GetClass("PHASEGeneratorNodeDefinition")}
	})
	return PHASEGeneratorNodeDefinitionClass
}

type _PHASEGeneratorNodeDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASEGeneratorNodeDefinition] class.
type IPHASEGeneratorNodeDefinition interface {
	IPHASESoundEventNodeDefinition
	// properties:
	GainMetaParameterDefinition() IPHASENumberMetaParameterDefinition
	SetGainMetaParameterDefinition(value IPHASENumberMetaParameterDefinition)
	RateMetaParameterDefinition() IPHASENumberMetaParameterDefinition
	SetRateMetaParameterDefinition(value IPHASENumberMetaParameterDefinition)
	CalibrationMode() PHASECalibrationMode
	SetCalibrationMode(value PHASECalibrationMode)
	Group() IPHASEGroup
	SetGroup(value IPHASEGroup)
	Level() float64
	SetLevel(value float64)
	MixerDefinition() IPHASEMixerDefinition
	SetMixerDefinition(value IPHASEMixerDefinition)
	Rate() float64
	SetRate(value float64)
	// methods:
}

// A base class for nodes that provide audio data to generate sound.
//
// This class encapsulates shared logic for subclasses that provide audio data to a mixer for sound output, namely and .


// A base class for nodes that provide audio data to generate sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition
type PHASEGeneratorNodeDefinition struct {
	PHASESoundEventNodeDefinition
}

// PHASEGeneratorNodeDefinitionFrom constructs a [PHASEGeneratorNodeDefinition] from an unsafe.Pointer.
//
// A base class for nodes that provide audio data to generate sound.
func PHASEGeneratorNodeDefinitionFrom(ptr unsafe.Pointer) PHASEGeneratorNodeDefinition {
	return PHASEGeneratorNodeDefinition{
		PHASESoundEventNodeDefinition: PHASESoundEventNodeDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEGeneratorNodeDefinitionClass) Alloc() PHASEGeneratorNodeDefinition {
	rv := objc.Send[PHASEGeneratorNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEGeneratorNodeDefinitionClass) New() PHASEGeneratorNodeDefinition {
	rv := objc.Send[PHASEGeneratorNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEGeneratorNodeDefinition) Init() PHASEGeneratorNodeDefinition {
	rv := objc.Send[PHASEGeneratorNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEGeneratorNodeDefinition) Autorelease() PHASEGeneratorNodeDefinition {
	rv := objc.Send[PHASEGeneratorNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEGeneratorNodeDefinition creates a new PHASEGeneratorNodeDefinition instance.
func NewPHASEGeneratorNodeDefinition() PHASEGeneratorNodeDefinition {
	return getPHASEGeneratorNodeDefinitionClass().New()
}



// A meta parameter that dynamically changes the audio’s loudness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/gainMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) GainMetaParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("gainMetaParameterDefinition"))
	return rv
}


// A meta parameter that dynamically changes the audio’s loudness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/gainMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) SetGainMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGainMetaParameterDefinition:"), value)
}


// A meta parameter that dynamically changes the audio’s rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rateMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) RateMetaParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("rateMetaParameterDefinition"))
	return rv
}


// A meta parameter that dynamically changes the audio’s rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rateMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) SetRateMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRateMetaParameterDefinition:"), value)
}


// A sound pressure level strategy for loudness correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/calibrationmode
func (p_ PHASEGeneratorNodeDefinition) CalibrationMode() PHASECalibrationMode {
	rv := objc.Send[PHASECalibrationMode](p_.ID, objc.Sel("calibrationMode"))
	return rv
}


// A sound pressure level strategy for loudness correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/calibrationmode
func (p_ PHASEGeneratorNodeDefinition) SetCalibrationMode(value PHASECalibrationMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCalibrationMode:"), value)
}


// A group this node conforms to for gain and rate control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/group
func (p_ PHASEGeneratorNodeDefinition) Group() IPHASEGroup {
	rv := objc.Send[PHASEGroup](p_.ID, objc.Sel("group"))
	return rv
}


// A group this node conforms to for gain and rate control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/group
func (p_ PHASEGeneratorNodeDefinition) SetGroup(value IPHASEGroup) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGroup:"), value)
}


// The node’s loudness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/level
func (p_ PHASEGeneratorNodeDefinition) Level() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("level"))
	return rv
}


// The node’s loudness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/level
func (p_ PHASEGeneratorNodeDefinition) SetLevel(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLevel:"), value)
}


// An object that combines audio layers for the node’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/mixerdefinition
func (p_ PHASEGeneratorNodeDefinition) MixerDefinition() IPHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](p_.ID, objc.Sel("mixerDefinition"))
	return rv
}


// An object that combines audio layers for the node’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/mixerdefinition
func (p_ PHASEGeneratorNodeDefinition) SetMixerDefinition(value IPHASEMixerDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMixerDefinition:"), value)
}


// A playback speed for the node’s audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/rate
func (p_ PHASEGeneratorNodeDefinition) Rate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rate"))
	return rv
}


// A playback speed for the node’s audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasegeneratornodedefinition/rate
func (p_ PHASEGeneratorNodeDefinition) SetRate(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRate:"), value)
}



