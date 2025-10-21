// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEAmbientMixerDefinition] class.
var (
	PHASEAmbientMixerDefinitionClass     _PHASEAmbientMixerDefinitionClass
	PHASEAmbientMixerDefinitionClassOnce sync.Once
)

func getPHASEAmbientMixerDefinitionClass() _PHASEAmbientMixerDefinitionClass {
	PHASEAmbientMixerDefinitionClassOnce.Do(func() {
		PHASEAmbientMixerDefinitionClass = _PHASEAmbientMixerDefinitionClass{objc.GetClass("PHASEAmbientMixerDefinition")}
	})
	return PHASEAmbientMixerDefinitionClass
}

type _PHASEAmbientMixerDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASEAmbientMixerDefinition] class.
type IPHASEAmbientMixerDefinition interface {
	IPHASEMixerDefinition
}

// An audio-layering object that outputs sound in a particular direction in 3D space.
//
// As an audio-layering object, this class combines multiple audio signals to a single signal for the output device. Play audio with a 3D orientation using this class when you supply a quaternion for the argument of the initializer. For information on orientation the sound, see . You also supply the intitializer with a channel layout in either mono, stereo, or surround formats. Surround audio files create the best listening experience due to their extra channel data. The framework renders each channel from the direction of its corresponding speaker in the channel layout. This class ignores low-frequency effect channels that may be present in the layout.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAmbientMixerDefinition
type PHASEAmbientMixerDefinition struct {
	PHASEMixerDefinition
}

// PHASEAmbientMixerDefinitionFrom constructs a [PHASEAmbientMixerDefinition] from an unsafe.Pointer.
//
// An audio-layering object that outputs sound in a particular direction in 3D space.
func PHASEAmbientMixerDefinitionFrom(ptr unsafe.Pointer) PHASEAmbientMixerDefinition {
	return PHASEAmbientMixerDefinition{
		PHASEMixerDefinition: PHASEMixerDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEAmbientMixerDefinitionClass) Alloc() PHASEAmbientMixerDefinition {
	rv := objc.Send[PHASEAmbientMixerDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEAmbientMixerDefinitionClass) New() PHASEAmbientMixerDefinition {
	rv := objc.Send[PHASEAmbientMixerDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEAmbientMixerDefinition) Init() PHASEAmbientMixerDefinition {
	rv := objc.Send[PHASEAmbientMixerDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEAmbientMixerDefinition) Autorelease() PHASEAmbientMixerDefinition {
	rv := objc.Send[PHASEAmbientMixerDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEAmbientMixerDefinition creates a new PHASEAmbientMixerDefinition instance.
func NewPHASEAmbientMixerDefinition() PHASEAmbientMixerDefinition {
	return getPHASEAmbientMixerDefinitionClass().New()
}




// Creates an ambient mixer with the given channel layout and orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAmbientMixerDefinition/init(channelLayout:orientation:)
func NewPHASEAmbientMixerDefinitionWithChannelLayoutOrientation(layout unsafe.Pointer, orientation unsafe.Pointer) PHASEAmbientMixerDefinition {
	instance := getPHASEAmbientMixerDefinitionClass().Alloc()
	rv := objc.Send[PHASEAmbientMixerDefinition](instance.ID, objc.Sel("initWithChannelLayout:orientation:"), layout, orientation)
	rv.Autorelease()
	return rv
}



// Creates a named ambient mixer with the given channel layout and orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAmbientMixerDefinition/init(channelLayout:orientation:identifier:)
func NewPHASEAmbientMixerDefinitionWithChannelLayoutOrientationIdentifier(layout unsafe.Pointer, orientation unsafe.Pointer, identifier string) PHASEAmbientMixerDefinition {
	instance := getPHASEAmbientMixerDefinitionClass().Alloc()
	rv := objc.Send[PHASEAmbientMixerDefinition](instance.ID, objc.Sel("initWithChannelLayout:orientation:identifier:"), layout, orientation, objc.String(identifier))
	rv.Autorelease()
	return rv
}


// The channel layout of input audio.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAmbientMixerDefinition/inputChannelLayout
func (p_ PHASEAmbientMixerDefinition) InputChannelLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("inputChannelLayout"))
	return rv
}

// A quaternion that describes the orientation of the speaker layout relative to the scene origin.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAmbientMixerDefinition/orientation
func (p_ PHASEAmbientMixerDefinition) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("orientation"))
	return rv
}

// A matrix, in local coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseobject/transform
func (p_ PHASEAmbientMixerDefinition) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("transform"))
	return rv
}


// SetTransform sets the value of the transform property.
// A matrix, in local coordinates, that determines the object’s pose in the scene.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseobject/transform
func (p_ PHASEAmbientMixerDefinition) SetTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransform:"), value)
}


