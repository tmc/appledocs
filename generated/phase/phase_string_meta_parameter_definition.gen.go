// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEStringMetaParameterDefinition] class.
var (
	PHASEStringMetaParameterDefinitionClass     _PHASEStringMetaParameterDefinitionClass
	PHASEStringMetaParameterDefinitionClassOnce sync.Once
)

func getPHASEStringMetaParameterDefinitionClass() _PHASEStringMetaParameterDefinitionClass {
	PHASEStringMetaParameterDefinitionClassOnce.Do(func() {
		PHASEStringMetaParameterDefinitionClass = _PHASEStringMetaParameterDefinitionClass{objc.GetClass("PHASEStringMetaParameterDefinition")}
	})
	return PHASEStringMetaParameterDefinitionClass
}

type _PHASEStringMetaParameterDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASEStringMetaParameterDefinition] class.
type IPHASEStringMetaParameterDefinition interface {
	IPHASEMetaParameterDefinition
	// properties:
	GlobalMetaParameters() IPHASEMetaParameter
	SetGlobalMetaParameters(value IPHASEMetaParameter)
	MetaParameters() IPHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
	// methods:
}

// A specification for a metaparameter defined by text.
//
// Use this class to spawn discrete instances of , for example, a “player speed” metaparameter that the app changes gradually from to . To use a number metaparameter, create an instance of this class and: Register it with the engine by calling , then access the instance of this class in the engine’s dictionary. Pass it to the initializer, , and then access the instance of this class in a sound event’s dictionary.


// A specification for a metaparameter defined by text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStringMetaParameterDefinition
type PHASEStringMetaParameterDefinition struct {
	PHASEMetaParameterDefinition
}

// PHASEStringMetaParameterDefinitionFrom constructs a [PHASEStringMetaParameterDefinition] from an unsafe.Pointer.
//
// A specification for a metaparameter defined by text.
func PHASEStringMetaParameterDefinitionFrom(ptr unsafe.Pointer) PHASEStringMetaParameterDefinition {
	return PHASEStringMetaParameterDefinition{
		PHASEMetaParameterDefinition: PHASEMetaParameterDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEStringMetaParameterDefinitionClass) Alloc() PHASEStringMetaParameterDefinition {
	rv := objc.Send[PHASEStringMetaParameterDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEStringMetaParameterDefinitionClass) New() PHASEStringMetaParameterDefinition {
	rv := objc.Send[PHASEStringMetaParameterDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEStringMetaParameterDefinition) Init() PHASEStringMetaParameterDefinition {
	rv := objc.Send[PHASEStringMetaParameterDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEStringMetaParameterDefinition) Autorelease() PHASEStringMetaParameterDefinition {
	rv := objc.Send[PHASEStringMetaParameterDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEStringMetaParameterDefinition creates a new PHASEStringMetaParameterDefinition instance.
func NewPHASEStringMetaParameterDefinition() PHASEStringMetaParameterDefinition {
	return getPHASEStringMetaParameterDefinitionClass().New()
}



// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameterDefinition) GlobalMetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameterDefinition) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameterDefinition) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameterDefinition) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}



