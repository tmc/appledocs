// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHASESoundEventNodeDefinition] class.
var (
	PHASESoundEventNodeDefinitionClass     _PHASESoundEventNodeDefinitionClass
	PHASESoundEventNodeDefinitionClassOnce sync.Once
)

func getPHASESoundEventNodeDefinitionClass() _PHASESoundEventNodeDefinitionClass {
	PHASESoundEventNodeDefinitionClassOnce.Do(func() {
		PHASESoundEventNodeDefinitionClass = _PHASESoundEventNodeDefinitionClass{objc.GetClass("PHASESoundEventNodeDefinition")}
	})
	return PHASESoundEventNodeDefinitionClass
}

type _PHASESoundEventNodeDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASESoundEventNodeDefinition] class.
type IPHASESoundEventNodeDefinition interface {
	IPHASEDefinition
	// properties:
	Children() []IPHASESoundEventNodeDefinition
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	MetaParameters() IPHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
	// methods:
}

// A base class for sound event nodes that connect to form a node hierarchy.
//
// This class defines the base functionality for an object that, depending on the derived class’s type, either plays audio or hands off the invocation to one or more other nodes.


// A base class for sound event nodes that connect to form a node hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventNodeDefinition
type PHASESoundEventNodeDefinition struct {
	PHASEDefinition
}

// PHASESoundEventNodeDefinitionFrom constructs a [PHASESoundEventNodeDefinition] from an unsafe.Pointer.
//
// A base class for sound event nodes that connect to form a node hierarchy.
func PHASESoundEventNodeDefinitionFrom(ptr unsafe.Pointer) PHASESoundEventNodeDefinition {
	return PHASESoundEventNodeDefinition{
		PHASEDefinition: PHASEDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASESoundEventNodeDefinitionClass) Alloc() PHASESoundEventNodeDefinition {
	rv := objc.Send[PHASESoundEventNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASESoundEventNodeDefinitionClass) New() PHASESoundEventNodeDefinition {
	rv := objc.Send[PHASESoundEventNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESoundEventNodeDefinition) Init() PHASESoundEventNodeDefinition {
	rv := objc.Send[PHASESoundEventNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESoundEventNodeDefinition) Autorelease() PHASESoundEventNodeDefinition {
	rv := objc.Send[PHASESoundEventNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESoundEventNodeDefinition creates a new PHASESoundEventNodeDefinition instance.
func NewPHASESoundEventNodeDefinition() PHASESoundEventNodeDefinition {
	return getPHASESoundEventNodeDefinitionClass().New()
}



// An array of child sound event nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESoundEventNodeDefinition/children
func (p_ PHASESoundEventNodeDefinition) Children() []IPHASESoundEventNodeDefinition {
	rv := objc.Send[[]PHASESoundEventNodeDefinition](p_.ID, objc.Sel("children"))
	return rv
}


// A unique name for the definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasedefinition/identifier
func (p_ PHASESoundEventNodeDefinition) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}


// A unique name for the definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasedefinition/identifier
func (p_ PHASESoundEventNodeDefinition) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), value)
}


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASESoundEventNodeDefinition) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASESoundEventNodeDefinition) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}



