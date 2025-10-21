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
}

// A base class for nodes that provide audio data to generate sound.
//
// This class encapsulates shared logic for subclasses that provide audio data to a mixer for sound output, namely and .
//
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
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/gainMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) GainMetaParameterDefinition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("gainMetaParameterDefinition"))
	return rv
}


// SetGainMetaParameterDefinition sets the value of the gainMetaParameterDefinition property.
// A meta parameter that dynamically changes the audio’s loudness.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/gainMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) SetGainMetaParameterDefinition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGainMetaParameterDefinition:"), value)
}

// A playback speed for the node’s audio.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rate
func (p_ PHASEGeneratorNodeDefinition) Rate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
// A playback speed for the node’s audio.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rate
func (p_ PHASEGeneratorNodeDefinition) SetRate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRate:"), value)
}

// A meta parameter that dynamically changes the audio’s rate.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rateMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) RateMetaParameterDefinition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("rateMetaParameterDefinition"))
	return rv
}


// SetRateMetaParameterDefinition sets the value of the rateMetaParameterDefinition property.
// A meta parameter that dynamically changes the audio’s rate.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGeneratorNodeDefinition/rateMetaParameterDefinition
func (p_ PHASEGeneratorNodeDefinition) SetRateMetaParameterDefinition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRateMetaParameterDefinition:"), value)
}



