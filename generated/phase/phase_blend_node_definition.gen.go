// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEBlendNodeDefinition] class.
var (
	PHASEBlendNodeDefinitionClass     _PHASEBlendNodeDefinitionClass
	PHASEBlendNodeDefinitionClassOnce sync.Once
)

func getPHASEBlendNodeDefinitionClass() _PHASEBlendNodeDefinitionClass {
	PHASEBlendNodeDefinitionClassOnce.Do(func() {
		PHASEBlendNodeDefinitionClass = _PHASEBlendNodeDefinitionClass{objc.GetClass("PHASEBlendNodeDefinition")}
	})
	return PHASEBlendNodeDefinitionClass
}

type _PHASEBlendNodeDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASEBlendNodeDefinition] class.
type IPHASEBlendNodeDefinition interface {
	IPHASESoundEventNodeDefinition
	AddRangeWithEnvelopeSubtree(envelope unsafe.Pointer, subtree unsafe.Pointer)
}

// A node that smoothly fades between the audio of its child nodes.
//
// This class defines a threshold and a numeric parameter the app increases and decreases to fade between child nodes. Each child node defines a range within the threshold in which the child node plays audio. As the app moves the blend parameter value between and the threshold, the blend node plays the audio of its child nodes whose range and fade curve overlap at the current value.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition
type PHASEBlendNodeDefinition struct {
	PHASESoundEventNodeDefinition
}

// PHASEBlendNodeDefinitionFrom constructs a [PHASEBlendNodeDefinition] from an unsafe.Pointer.
//
// A node that smoothly fades between the audio of its child nodes.
func PHASEBlendNodeDefinitionFrom(ptr unsafe.Pointer) PHASEBlendNodeDefinition {
	return PHASEBlendNodeDefinition{
		PHASESoundEventNodeDefinition: PHASESoundEventNodeDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEBlendNodeDefinitionClass) Alloc() PHASEBlendNodeDefinition {
	rv := objc.Send[PHASEBlendNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEBlendNodeDefinitionClass) New() PHASEBlendNodeDefinition {
	rv := objc.Send[PHASEBlendNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEBlendNodeDefinition) Init() PHASEBlendNodeDefinition {
	rv := objc.Send[PHASEBlendNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEBlendNodeDefinition) Autorelease() PHASEBlendNodeDefinition {
	rv := objc.Send[PHASEBlendNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEBlendNodeDefinition creates a new PHASEBlendNodeDefinition instance.
func NewPHASEBlendNodeDefinition() PHASEBlendNodeDefinition {
	return getPHASEBlendNodeDefinitionClass().New()
}


// Creates a named blend node with a maxiumum blend range value.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/init(blendMetaParameterDefinition:identifier:)
func NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinitionIdentifier(blendMetaParameterDefinition unsafe.Pointer, identifier string) PHASEBlendNodeDefinition {
	instance := getPHASEBlendNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEBlendNodeDefinition](instance.ID, objc.Sel("initWithBlendMetaParameterDefinition:identifier:"), blendMetaParameterDefinition, objc.String(identifier))
	rv.Autorelease()
	return rv
}

// Creates a blend node with a maxiumum blend range value.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/init(blendMetaParameterDefinition:)
func NewPHASEBlendNodeDefinitionWithBlendMetaParameterDefinition(blendMetaParameterDefinition unsafe.Pointer) PHASEBlendNodeDefinition {
	instance := getPHASEBlendNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEBlendNodeDefinition](instance.ID, objc.Sel("initWithBlendMetaParameterDefinition:"), blendMetaParameterDefinition)
	rv.Autorelease()
	return rv
}


// Adds a child node with an envelope.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEBlendNodeDefinition/addRange(envelope:subtree:)
func (p_ PHASEBlendNodeDefinition) AddRangeWithEnvelopeSubtree(envelope unsafe.Pointer, subtree unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addRangeWithEnvelope:subtree:"), envelope, subtree)
}


