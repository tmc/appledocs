// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEMixer] class.
var (
	PHASEMixerClass     _PHASEMixerClass
	PHASEMixerClassOnce sync.Once
)

func getPHASEMixerClass() _PHASEMixerClass {
	PHASEMixerClassOnce.Do(func() {
		PHASEMixerClass = _PHASEMixerClass{objc.GetClass("PHASEMixer")}
	})
	return PHASEMixerClass
}

type _PHASEMixerClass struct {
	class objc.Class
}

// An interface definition for the [PHASEMixer] class.
type IPHASEMixer interface {
	objectivec.IObject
}

// An object that combines multiple audio signals into a single signal.
//
// Mixers provide a single point of control over the multiple audio signals they combine. To create a mixer, you provide the framework with a mixer definition; see . Subclasses of this class define unique properties the app sets to control specific features. For example, the spatial mixer ( ) adds environmental effects into the output audio signal.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixer
type PHASEMixer struct {
	objectivec.Object
}

// PHASEMixerFrom constructs a [PHASEMixer] from an unsafe.Pointer.
//
// An object that combines multiple audio signals into a single signal.
func PHASEMixerFrom(ptr unsafe.Pointer) PHASEMixer {
	return PHASEMixer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEMixerClass) Alloc() PHASEMixer {
	rv := objc.Send[PHASEMixer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEMixerClass) New() PHASEMixer {
	rv := objc.Send[PHASEMixer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMixer) Init() PHASEMixer {
	rv := objc.Send[PHASEMixer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMixer) Autorelease() PHASEMixer {
	rv := objc.Send[PHASEMixer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMixer creates a new PHASEMixer instance.
func NewPHASEMixer() PHASEMixer {
	return getPHASEMixerClass().New()
}


// The mixer’s volume.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixer/gain
func (p_ PHASEMixer) Gain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("gain"))
	return rv
}

// A parameter that changes the mixer’s volume gradually over a period of time.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixer/gainMetaParameter
func (p_ PHASEMixer) GainMetaParameter() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("gainMetaParameter"))
	return rv
}

// A unique name for the mixer.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMixer/identifier
func (p_ PHASEMixer) Identifier() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("identifier"))
	return rv
}



