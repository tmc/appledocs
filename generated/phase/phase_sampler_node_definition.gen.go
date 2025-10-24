// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [PHASESamplerNodeDefinition] class.
type IPHASESamplerNodeDefinition interface {
	IPHASEGeneratorNodeDefinition
	// properties:
	PlaybackMode() PHASEPlaybackMode
	SetPlaybackMode(value PHASEPlaybackMode)
	AssetIdentifier() objc.IObject /* cross-framework: NSString */
	SetAssetIdentifier(value objc.IObject /* cross-framework: NSString */)
	CullOption() unsafe.Pointer
	SetCullOption(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (pc _PHASESamplerNodeDefinitionClass) Alloc() PHASESamplerNodeDefinition {
	rv := objc.Send[PHASESamplerNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An option that determines whether the node’s audio plays in a loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition/playbackMode
func (p_ PHASESamplerNodeDefinition) PlaybackMode() PHASEPlaybackMode {
	rv := objc.Send[PHASEPlaybackMode](p_.ID, objc.Sel("playbackMode"))
	return rv
}


// An option that determines whether the node’s audio plays in a loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESamplerNodeDefinition/playbackMode
func (p_ PHASESamplerNodeDefinition) SetPlaybackMode(value PHASEPlaybackMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackMode:"), value)
}


// The name of the audio this node plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesamplernodedefinition/assetidentifier
func (p_ PHASESamplerNodeDefinition) AssetIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("assetIdentifier"))
	return rv
}


// The name of the audio this node plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesamplernodedefinition/assetidentifier
func (p_ PHASESamplerNodeDefinition) SetAssetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAssetIdentifier:"), value)
}


// The action the engine performs after it temporarily removes the node’s sound from the audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesamplernodedefinition/culloption
func (p_ PHASESamplerNodeDefinition) CullOption() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("cullOption"))
	return rv
}


// The action the engine performs after it temporarily removes the node’s sound from the audio output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesamplernodedefinition/culloption
func (p_ PHASESamplerNodeDefinition) SetCullOption(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCullOption:"), value)
}



