// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASEAmbientMixerDefinition */


/* debug [class_header]: Header for PHASEAmbientMixerDefinition */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEAmbientMixerDefinition */
// An interface definition for the [PHASEAmbientMixerDefinition] class.
type IPHASEAmbientMixerDefinition interface {
	IPHASEMixerDefinition
	
/* debug [class_interface_properties]: Properties for PHASEAmbientMixerDefinition */
	// properties:
	InputChannelLayout() avfaudio.AudioChannelLayout
	Orientation() unsafe.Pointer
	Transform() unsafe.Pointer
	SetTransform(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEAmbientMixerDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEAmbientMixerDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEAmbientMixerDefinitionClass) Alloc() PHASEAmbientMixerDefinition {
	rv := objc.Send[PHASEAmbientMixerDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEAmbientMixerDefinition */
// An audio-layering object that outputs sound in a particular direction in 3D space.
//
// As an audio-layering object, this class combines multiple audio signals to a single signal for the output device. Play audio with a 3D orientation using this class when you supply a quaternion for the argument of the initializer. For information on orientation the sound, see . You also supply the intitializer with a channel layout in either mono, stereo, or surround formats. Surround audio files create the best listening experience due to their extra channel data. The framework renders each channel from the direction of its corresponding speaker in the channel layout. This class ignores low-frequency effect channels that may be present in the layout.


// An audio-layering object that outputs sound in a particular direction in 3D space.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEAmbientMixerDefinition */

// Creates an ambient mixer with the given channel layout and orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAmbientMixerDefinition/init(channelLayout:orientation:)
func NewPHASEAmbientMixerDefinitionWithChannelLayoutOrientation(layout avfaudio.AudioChannelLayout, orientation unsafe.Pointer) PHASEAmbientMixerDefinition {
	instance := getPHASEAmbientMixerDefinitionClass().Alloc()
	rv := objc.Send[PHASEAmbientMixerDefinition](instance.ID, objc.Sel("initWithChannelLayout:orientation:"), layout, orientation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEAmbientMixerDefinitionWithChannelLayoutOrientation */


// Creates a named ambient mixer with the given channel layout and orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAmbientMixerDefinition/init(channelLayout:orientation:identifier:)
func NewPHASEAmbientMixerDefinitionWithChannelLayoutOrientationIdentifier(layout avfaudio.AudioChannelLayout, orientation unsafe.Pointer, identifier objc.IObject /* cross-framework: NSString */) PHASEAmbientMixerDefinition {
	instance := getPHASEAmbientMixerDefinitionClass().Alloc()
	rv := objc.Send[PHASEAmbientMixerDefinition](instance.ID, objc.Sel("initWithChannelLayout:orientation:identifier:"), layout, orientation, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEAmbientMixerDefinitionWithChannelLayoutOrientationIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEAmbientMixerDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEAmbientMixerDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEAmbientMixerDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEAmbientMixerDefinition */

// The channel layout of input audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAmbientMixerDefinition/inputChannelLayout
func (p_ PHASEAmbientMixerDefinition) InputChannelLayout() avfaudio.AudioChannelLayout {
	rv := objc.Send[avfaudio.AudioChannelLayout](p_.ID, objc.Sel("inputChannelLayout"))
	return rv
}/* debug [instance_properties/getter]: inputChannelLayout */


// A quaternion that describes the orientation of the speaker layout relative to the scene origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEAmbientMixerDefinition/orientation
func (p_ PHASEAmbientMixerDefinition) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("orientation"))
	return rv
}/* debug [instance_properties/getter]: orientation */


// A matrix, in local coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseobject/transform
func (p_ PHASEAmbientMixerDefinition) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// A matrix, in local coordinates, that determines the object’s pose in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseobject/transform
func (p_ PHASEAmbientMixerDefinition) SetTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEAmbientMixerDefinition */


