// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASEPushStreamNodeDefinition */


/* debug [class_header]: Header for PHASEPushStreamNodeDefinition */
// The class instance for the [PHASEPushStreamNodeDefinition] class.
var (
	PHASEPushStreamNodeDefinitionClass     _PHASEPushStreamNodeDefinitionClass
	PHASEPushStreamNodeDefinitionClassOnce sync.Once
)

func getPHASEPushStreamNodeDefinitionClass() _PHASEPushStreamNodeDefinitionClass {
	PHASEPushStreamNodeDefinitionClassOnce.Do(func() {
		PHASEPushStreamNodeDefinitionClass = _PHASEPushStreamNodeDefinitionClass{objc.GetClass("PHASEPushStreamNodeDefinition")}
	})
	return PHASEPushStreamNodeDefinitionClass
}

type _PHASEPushStreamNodeDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEPushStreamNodeDefinition */
// An interface definition for the [PHASEPushStreamNodeDefinition] class.
type IPHASEPushStreamNodeDefinition interface {
	IPHASEGeneratorNodeDefinition
	
/* debug [class_interface_properties]: Properties for PHASEPushStreamNodeDefinition */
	// properties:
	Format() avfaudio.AudioFormat
	Normalize() bool
	SetNormalize(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEPushStreamNodeDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEPushStreamNodeDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEPushStreamNodeDefinitionClass) Alloc() PHASEPushStreamNodeDefinition {
	rv := objc.Send[PHASEPushStreamNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEPushStreamNodeDefinitionClass) New() PHASEPushStreamNodeDefinition {
	rv := objc.Send[PHASEPushStreamNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEPushStreamNodeDefinition) Init() PHASEPushStreamNodeDefinition {
	rv := objc.Send[PHASEPushStreamNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEPushStreamNodeDefinition) Autorelease() PHASEPushStreamNodeDefinition {
	rv := objc.Send[PHASEPushStreamNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEPushStreamNodeDefinition creates a new PHASEPushStreamNodeDefinition instance.
func NewPHASEPushStreamNodeDefinition() PHASEPushStreamNodeDefinition {
	return getPHASEPushStreamNodeDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEPushStreamNodeDefinition */
// A node that plays a sequence of audio buffers.
//
// Use this node to create sound events for a piecemeal audio source, for example, an audio stream that your app accesses over the network or loads from a memory-mapped file on disk.


// A node that plays a sequence of audio buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition
type PHASEPushStreamNodeDefinition struct {
	PHASEGeneratorNodeDefinition
}

// PHASEPushStreamNodeDefinitionFrom constructs a [PHASEPushStreamNodeDefinition] from an unsafe.Pointer.
//
// A node that plays a sequence of audio buffers.
func PHASEPushStreamNodeDefinitionFrom(ptr unsafe.Pointer) PHASEPushStreamNodeDefinition {
	return PHASEPushStreamNodeDefinition{
		PHASEGeneratorNodeDefinition: PHASEGeneratorNodeDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEPushStreamNodeDefinition */

// Creates a node definition for audio streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/init(mixerDefinition:format:)
func NewPHASEPushStreamNodeDefinitionWithMixerDefinitionFormat(mixerDefinition IPHASEMixerDefinition, format avfaudio.AudioFormat) PHASEPushStreamNodeDefinition {
	instance := getPHASEPushStreamNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEPushStreamNodeDefinition](instance.ID, objc.Sel("initWithMixerDefinition:format:"), mixerDefinition, format)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEPushStreamNodeDefinitionWithMixerDefinitionFormat */


// Creates a named node definition for audio streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/init(mixerDefinition:format:identifier:)
func NewPHASEPushStreamNodeDefinitionWithMixerDefinitionFormatIdentifier(mixerDefinition IPHASEMixerDefinition, format avfaudio.AudioFormat, identifier objc.IObject /* cross-framework: NSString */) PHASEPushStreamNodeDefinition {
	instance := getPHASEPushStreamNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEPushStreamNodeDefinition](instance.ID, objc.Sel("initWithMixerDefinition:format:identifier:"), mixerDefinition, format, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEPushStreamNodeDefinitionWithMixerDefinitionFormatIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEPushStreamNodeDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEPushStreamNodeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEPushStreamNodeDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEPushStreamNodeDefinition */

// The format of the audio stream data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/format
func (p_ PHASEPushStreamNodeDefinition) Format() avfaudio.AudioFormat {
	rv := objc.Send[avfaudio.AudioFormat](p_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// An option that resizes loudness of the audio stream for consistency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/normalize
func (p_ PHASEPushStreamNodeDefinition) Normalize() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("normalize"))
	return rv
}/* debug [instance_properties/getter]: normalize */


// An option that resizes loudness of the audio stream for consistency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/normalize
func (p_ PHASEPushStreamNodeDefinition) SetNormalize(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNormalize:"), value)
}/* debug [instance_properties/setter]: normalize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEPushStreamNodeDefinition */


