// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEDefinition] class.
var (
	PHASEDefinitionClass     _PHASEDefinitionClass
	PHASEDefinitionClassOnce sync.Once
)

func getPHASEDefinitionClass() _PHASEDefinitionClass {
	PHASEDefinitionClassOnce.Do(func() {
		PHASEDefinitionClass = _PHASEDefinitionClass{objc.GetClass("PHASEDefinition")}
	})
	return PHASEDefinitionClass
}

type _PHASEDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASEDefinition] class.
type IPHASEDefinition interface {
	objectivec.IObject
}

// A base class that adds a name to framework definitions.
//
// Various PHASE classes derive from this class, for example, , , and . This class represents a template from which PHASE creates concrete subclasses at runtime. For example, when you register a global metaparameter definition using , PHASE returns a subclass, , that identifies a usable metaparameter by name. To access the usable metaparameter, pass the into the dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDefinition
type PHASEDefinition struct {
	objectivec.Object
}

// PHASEDefinitionFrom constructs a [PHASEDefinition] from an unsafe.Pointer.
//
// A base class that adds a name to framework definitions.
func PHASEDefinitionFrom(ptr unsafe.Pointer) PHASEDefinition {
	return PHASEDefinition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEDefinitionClass) Alloc() PHASEDefinition {
	rv := objc.Send[PHASEDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEDefinitionClass) New() PHASEDefinition {
	rv := objc.Send[PHASEDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEDefinition) Init() PHASEDefinition {
	rv := objc.Send[PHASEDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEDefinition) Autorelease() PHASEDefinition {
	rv := objc.Send[PHASEDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEDefinition creates a new PHASEDefinition instance.
func NewPHASEDefinition() PHASEDefinition {
	return getPHASEDefinitionClass().New()
}


// A unique name for the definition.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEDefinition/identifier
func (p_ PHASEDefinition) Identifier() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("identifier"))
	return rv
}

// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEDefinition) GlobalMetaParameters() PHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}


// SetGlobalMetaParameters sets the value of the globalMetaParameters property.
// A dictionary of metaparameters that all sound event assets share.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEDefinition) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}



