// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PHASEPushStreamNodeDefinition] class.
type IPHASEPushStreamNodeDefinition interface {
	IPHASEGeneratorNodeDefinition
}

// A node that plays a sequence of audio buffers.
//
// Use this node to create sound events for a piecemeal audio source, for example, an audio stream that your app accesses over the network or loads from a memory-mapped file on disk.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEPushStreamNodeDefinitionClass) Alloc() PHASEPushStreamNodeDefinition {
	rv := objc.Send[PHASEPushStreamNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a node definition for audio streams.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/init(mixerDefinition:format:)
func NewPHASEPushStreamNodeDefinitionWithMixerDefinitionFormat(mixerDefinition unsafe.Pointer, format unsafe.Pointer) PHASEPushStreamNodeDefinition {
	instance := getPHASEPushStreamNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEPushStreamNodeDefinition](instance.ID, objc.Sel("initWithMixerDefinition:format:"), mixerDefinition, format)
	rv.Autorelease()
	return rv
}



// Creates a named node definition for audio streams.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/init(mixerDefinition:format:identifier:)
func NewPHASEPushStreamNodeDefinitionWithMixerDefinitionFormatIdentifier(mixerDefinition unsafe.Pointer, format unsafe.Pointer, identifier string) PHASEPushStreamNodeDefinition {
	instance := getPHASEPushStreamNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEPushStreamNodeDefinition](instance.ID, objc.Sel("initWithMixerDefinition:format:identifier:"), mixerDefinition, format, objc.String(identifier))
	rv.Autorelease()
	return rv
}


// The format of the audio stream data.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/format
func (p_ PHASEPushStreamNodeDefinition) Format() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("format"))
	return rv
}

// An option that resizes loudness of the audio stream for consistency.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/normalize
func (p_ PHASEPushStreamNodeDefinition) Normalize() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("normalize"))
	return rv
}


// SetNormalize sets the value of the normalize property.
// An option that resizes loudness of the audio stream for consistency.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNodeDefinition/normalize
func (p_ PHASEPushStreamNodeDefinition) SetNormalize(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNormalize:"), value)
}


