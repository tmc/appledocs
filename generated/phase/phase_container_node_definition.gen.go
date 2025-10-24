// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEContainerNodeDefinition] class.
var (
	PHASEContainerNodeDefinitionClass     _PHASEContainerNodeDefinitionClass
	PHASEContainerNodeDefinitionClassOnce sync.Once
)

func getPHASEContainerNodeDefinitionClass() _PHASEContainerNodeDefinitionClass {
	PHASEContainerNodeDefinitionClassOnce.Do(func() {
		PHASEContainerNodeDefinitionClass = _PHASEContainerNodeDefinitionClass{objc.GetClass("PHASEContainerNodeDefinition")}
	})
	return PHASEContainerNodeDefinitionClass
}

type _PHASEContainerNodeDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASEContainerNodeDefinition] class.
type IPHASEContainerNodeDefinition interface {
	IPHASESoundEventNodeDefinition
	// properties:
	// methods:
}

// A node that plays all its children at the same time.
//
// This node adds structure to the sound event tree while performing no conditional logic or audio playback of its own. By passing invocation to all its children at once, this class invokes the child nodes’ actions simultaneously.


// A node that plays all its children at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEContainerNodeDefinition
type PHASEContainerNodeDefinition struct {
	PHASESoundEventNodeDefinition
}

// PHASEContainerNodeDefinitionFrom constructs a [PHASEContainerNodeDefinition] from an unsafe.Pointer.
//
// A node that plays all its children at the same time.
func PHASEContainerNodeDefinitionFrom(ptr unsafe.Pointer) PHASEContainerNodeDefinition {
	return PHASEContainerNodeDefinition{
		PHASESoundEventNodeDefinition: PHASESoundEventNodeDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEContainerNodeDefinitionClass) Alloc() PHASEContainerNodeDefinition {
	rv := objc.Send[PHASEContainerNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEContainerNodeDefinitionClass) New() PHASEContainerNodeDefinition {
	rv := objc.Send[PHASEContainerNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEContainerNodeDefinition) Init() PHASEContainerNodeDefinition {
	rv := objc.Send[PHASEContainerNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEContainerNodeDefinition) Autorelease() PHASEContainerNodeDefinition {
	rv := objc.Send[PHASEContainerNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEContainerNodeDefinition creates a new PHASEContainerNodeDefinition instance.
func NewPHASEContainerNodeDefinition() PHASEContainerNodeDefinition {
	return getPHASEContainerNodeDefinitionClass().New()
}




