// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASESamplerNodeDefinition */


/* debug [class_header]: Header for PHASESamplerNodeDefinition */
// The class instance for the [PHASESamplerNodeDefinition] class.
var (
	PHASESamplerNodeDefinitionClass     _PHASESamplerNodeDefinitionClass
	PHASESamplerNodeDefinitionClassOnce sync.Once
)

func getPHASESamplerNodeDefinitionClass() _PHASESamplerNodeDefinitionClass {
	PHASESamplerNodeDefinitionClassOnce.Do(func() {
		PHASESamplerNodeDefinitionClass = _PHASESamplerNodeDefinitionClass{objc.GetClass("PHASESamplerNodeDefinition")}
	})
	return PHASESamplerNodeDefinitionClass
}

type _PHASESamplerNodeDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASESamplerNodeDefinition */
// An interface definition for the [PHASESamplerNodeDefinition] class.
type IPHASESamplerNodeDefinition interface {
	IPHASEGeneratorNodeDefinition
	
/* debug [class_interface_properties]: Properties for PHASESamplerNodeDefinition */
	// properties:
	AssetIdentifier() objc.IObject /* cross-framework: NSString */
	CullOption() PHASECullOption
	SetCullOption(value PHASECullOption)
	PlaybackMode() PHASEPlaybackMode
	SetPlaybackMode(value PHASEPlaybackMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASESamplerNodeDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASESamplerNodeDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASESamplerNodeDefinitionClass) Alloc() PHASESamplerNodeDefinition {
	rv := objc.Send[PHASESamplerNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASESamplerNodeDefinitionClass) New() PHASESamplerNodeDefinition {
	rv := objc.Send[PHASESamplerNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESamplerNodeDefinition) Init() PHASESamplerNodeDefinition {
	rv := objc.Send[PHASESamplerNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESamplerNodeDefinition) Autorelease() PHASESamplerNodeDefinition {
	rv := objc.Send[PHASESamplerNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESamplerNodeDefinition creates a new PHASESamplerNodeDefinition instance.
func NewPHASESamplerNodeDefinition() PHASESamplerNodeDefinition {
	return getPHASESamplerNodeDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASESamplerNodeDefinition */
// A node that plays complete audio data.
//
// Generate sound events from this node to play audio data that your app loads completely, either from disk or from memory.


// A node that plays complete audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition
type PHASESamplerNodeDefinition struct {
	PHASEGeneratorNodeDefinition
}

// PHASESamplerNodeDefinitionFrom constructs a [PHASESamplerNodeDefinition] from an unsafe.Pointer.
//
// A node that plays complete audio data.
func PHASESamplerNodeDefinitionFrom(ptr unsafe.Pointer) PHASESamplerNodeDefinition {
	return PHASESamplerNodeDefinition{
		PHASEGeneratorNodeDefinition: PHASEGeneratorNodeDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASESamplerNodeDefinition */

// Creates a sampler node with the given sound asset and mixer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition/init(soundAssetIdentifier:mixerDefinition:)
func NewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinition(soundAssetIdentifier objc.IObject /* cross-framework: NSString */, mixerDefinition IPHASEMixerDefinition) PHASESamplerNodeDefinition {
	instance := getPHASESamplerNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASESamplerNodeDefinition](instance.ID, objc.Sel("initWithSoundAssetIdentifier:mixerDefinition:"), soundAssetIdentifier, mixerDefinition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinition */


// Creates a named sampler node with the given sound asset and mixer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition/init(soundAssetIdentifier:mixerDefinition:identifier:)
func NewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinitionIdentifier(soundAssetIdentifier objc.IObject /* cross-framework: NSString */, mixerDefinition IPHASEMixerDefinition, identifier objc.IObject /* cross-framework: NSString */) PHASESamplerNodeDefinition {
	instance := getPHASESamplerNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASESamplerNodeDefinition](instance.ID, objc.Sel("initWithSoundAssetIdentifier:mixerDefinition:identifier:"), soundAssetIdentifier, mixerDefinition, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESamplerNodeDefinitionWithSoundAssetIdentifierMixerDefinitionIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASESamplerNodeDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASESamplerNodeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASESamplerNodeDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASESamplerNodeDefinition */

// The name of the audio this node plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition/assetIdentifier
func (p_ PHASESamplerNodeDefinition) AssetIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("assetIdentifier"))
	return rv
}/* debug [instance_properties/getter]: assetIdentifier */


// The action the engine performs after it temporarily removes the node’s sound from the audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition/cullOption
func (p_ PHASESamplerNodeDefinition) CullOption() PHASECullOption {
	rv := objc.Send[PHASECullOption](p_.ID, objc.Sel("cullOption"))
	return rv
}/* debug [instance_properties/getter]: cullOption */


// The action the engine performs after it temporarily removes the node’s sound from the audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition/cullOption
func (p_ PHASESamplerNodeDefinition) SetCullOption(value PHASECullOption) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCullOption:"), value)
}/* debug [instance_properties/setter]: cullOption */


// An option that determines whether the node’s audio plays in a loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition/playbackMode
func (p_ PHASESamplerNodeDefinition) PlaybackMode() PHASEPlaybackMode {
	rv := objc.Send[PHASEPlaybackMode](p_.ID, objc.Sel("playbackMode"))
	return rv
}/* debug [instance_properties/getter]: playbackMode */


// An option that determines whether the node’s audio plays in a loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition/playbackMode
func (p_ PHASESamplerNodeDefinition) SetPlaybackMode(value PHASEPlaybackMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackMode:"), value)
}/* debug [instance_properties/setter]: playbackMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASESamplerNodeDefinition */


