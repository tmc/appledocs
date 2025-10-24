// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASEChannelMixerDefinition */


/* debug [class_header]: Header for PHASEChannelMixerDefinition */
// The class instance for the [PHASEChannelMixerDefinition] class.
var (
	PHASEChannelMixerDefinitionClass     _PHASEChannelMixerDefinitionClass
	PHASEChannelMixerDefinitionClassOnce sync.Once
)

func getPHASEChannelMixerDefinitionClass() _PHASEChannelMixerDefinitionClass {
	PHASEChannelMixerDefinitionClassOnce.Do(func() {
		PHASEChannelMixerDefinitionClass = _PHASEChannelMixerDefinitionClass{objc.GetClass("PHASEChannelMixerDefinition")}
	})
	return PHASEChannelMixerDefinitionClass
}

type _PHASEChannelMixerDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEChannelMixerDefinition */
// An interface definition for the [PHASEChannelMixerDefinition] class.
type IPHASEChannelMixerDefinition interface {
	IPHASEMixerDefinition
	
/* debug [class_interface_properties]: Properties for PHASEChannelMixerDefinition */
	// properties:
	InputChannelLayout() avfaudio.AudioChannelLayout
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEChannelMixerDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEChannelMixerDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEChannelMixerDefinitionClass) Alloc() PHASEChannelMixerDefinition {
	rv := objc.Send[PHASEChannelMixerDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEChannelMixerDefinitionClass) New() PHASEChannelMixerDefinition {
	rv := objc.Send[PHASEChannelMixerDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEChannelMixerDefinition) Init() PHASEChannelMixerDefinition {
	rv := objc.Send[PHASEChannelMixerDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEChannelMixerDefinition) Autorelease() PHASEChannelMixerDefinition {
	rv := objc.Send[PHASEChannelMixerDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEChannelMixerDefinition creates a new PHASEChannelMixerDefinition instance.
func NewPHASEChannelMixerDefinition() PHASEChannelMixerDefinition {
	return getPHASEChannelMixerDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEChannelMixerDefinition */
// An audio-layering object that routes sound directly to the device’s output.
//
// Use this class to play one-time sounds such as menu clicks. This class defines the , which is the strategy the framework uses to send source mono or multichannel assets to the output for playback. The asset’s audio channels route to the output for playback according to the channel layout and runtime output conditions the app designates on an instance of this class. This class minimizes and — that is, source audio channel conversion to a higher or lower number of channels. For example, although a spatial mixer overrides the use of output channels by panning to convey listener position and orientation, the channel mixer maintains source audio channel layout to preserve the listening experience of the source audio.


// An audio-layering object that routes sound directly to the device’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEChannelMixerDefinition
type PHASEChannelMixerDefinition struct {
	PHASEMixerDefinition
}

// PHASEChannelMixerDefinitionFrom constructs a [PHASEChannelMixerDefinition] from an unsafe.Pointer.
//
// An audio-layering object that routes sound directly to the device’s output.
func PHASEChannelMixerDefinitionFrom(ptr unsafe.Pointer) PHASEChannelMixerDefinition {
	return PHASEChannelMixerDefinition{
		PHASEMixerDefinition: PHASEMixerDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEChannelMixerDefinition */

// Creates a channel mixer with the given channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEChannelMixerDefinition/init(channelLayout:)
func NewPHASEChannelMixerDefinitionWithChannelLayout(layout avfaudio.AudioChannelLayout) PHASEChannelMixerDefinition {
	instance := getPHASEChannelMixerDefinitionClass().Alloc()
	rv := objc.Send[PHASEChannelMixerDefinition](instance.ID, objc.Sel("initWithChannelLayout:"), layout)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEChannelMixerDefinitionWithChannelLayout */


// Creates a named channel mixer with the given channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEChannelMixerDefinition/init(channelLayout:identifier:)
func NewPHASEChannelMixerDefinitionWithChannelLayoutIdentifier(layout avfaudio.AudioChannelLayout, identifier objc.IObject /* cross-framework: NSString */) PHASEChannelMixerDefinition {
	instance := getPHASEChannelMixerDefinitionClass().Alloc()
	rv := objc.Send[PHASEChannelMixerDefinition](instance.ID, objc.Sel("initWithChannelLayout:identifier:"), layout, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEChannelMixerDefinitionWithChannelLayoutIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEChannelMixerDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEChannelMixerDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEChannelMixerDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEChannelMixerDefinition */

// The channel layout of the mixer’s input audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEChannelMixerDefinition/inputChannelLayout
func (p_ PHASEChannelMixerDefinition) InputChannelLayout() avfaudio.AudioChannelLayout {
	rv := objc.Send[avfaudio.AudioChannelLayout](p_.ID, objc.Sel("inputChannelLayout"))
	return rv
}/* debug [instance_properties/getter]: inputChannelLayout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEChannelMixerDefinition */


