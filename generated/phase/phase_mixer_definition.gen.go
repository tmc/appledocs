// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEMixerDefinition] class.
var (
	PHASEMixerDefinitionClass     _PHASEMixerDefinitionClass
	PHASEMixerDefinitionClassOnce sync.Once
)

func getPHASEMixerDefinitionClass() _PHASEMixerDefinitionClass {
	PHASEMixerDefinitionClassOnce.Do(func() {
		PHASEMixerDefinitionClass = _PHASEMixerDefinitionClass{objc.GetClass("PHASEMixerDefinition")}
	})
	return PHASEMixerDefinitionClass
}

type _PHASEMixerDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASEMixerDefinition] class.
type IPHASEMixerDefinition interface {
	IPHASEDefinition
}

// An object to initialize a mixer with a given configuration.
//
// A mixer combines multiple layers of audio to a single signal for transmission to the output device. The framework creates a mixer when you provide a mixer definition. Instead of creating an instance of this class, instantiate one of the mixer definition subclasses instead:
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition
type PHASEMixerDefinition struct {
	PHASEDefinition
}

// PHASEMixerDefinitionFrom constructs a [PHASEMixerDefinition] from an unsafe.Pointer.
//
// An object to initialize a mixer with a given configuration.
func PHASEMixerDefinitionFrom(ptr unsafe.Pointer) PHASEMixerDefinition {
	return PHASEMixerDefinition{
		PHASEDefinition: PHASEDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEMixerDefinitionClass) Alloc() PHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEMixerDefinitionClass) New() PHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMixerDefinition) Init() PHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMixerDefinition) Autorelease() PHASEMixerDefinition {
	rv := objc.Send[PHASEMixerDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMixerDefinition creates a new PHASEMixerDefinition instance.
func NewPHASEMixerDefinition() PHASEMixerDefinition {
	return getPHASEMixerDefinitionClass().New()
}


// The mixer’s volume.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition/gain
func (p_ PHASEMixerDefinition) Gain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("gain"))
	return rv
}


// SetGain sets the value of the gain property.
// The mixer’s volume.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition/gain
func (p_ PHASEMixerDefinition) SetGain(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGain:"), value)
}

// A template for a parameter that changes the mixer’s volume gradually over a period of time.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition/gainMetaParameterDefinition
func (p_ PHASEMixerDefinition) GainMetaParameterDefinition() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("gainMetaParameterDefinition"))
	return rv
}


// SetGainMetaParameterDefinition sets the value of the gainMetaParameterDefinition property.
// A template for a parameter that changes the mixer’s volume gradually over a period of time.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixerDefinition/gainMetaParameterDefinition
func (p_ PHASEMixerDefinition) SetGainMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGainMetaParameterDefinition:"), value)
}



