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
}

// A specification for a metaparameter defined by text.
//
// Use this class to spawn discrete instances of , for example, a “player speed” metaparameter that the app changes gradually from to . To use a number metaparameter, create an instance of this class and: Register it with the engine by calling , then access the instance of this class in the engine’s dictionary. Pass it to the initializer, , and then access the instance of this class in a sound event’s dictionary.
//
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




// Creates a specification for a textual metaparameter with the given value.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStringMetaParameterDefinition/init(value:)
func NewPHASEStringMetaParameterDefinitionWithValue(value string) PHASEStringMetaParameterDefinition {
	instance := getPHASEStringMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASEStringMetaParameterDefinition](instance.ID, objc.Sel("initWithValue:"), objc.String(value))
	rv.Autorelease()
	return rv
}



// Creates a specification for a named textual metaparameter with the given value.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStringMetaParameterDefinition/init(value:identifier:)
func NewPHASEStringMetaParameterDefinitionWithValueIdentifier(value string, identifier string) PHASEStringMetaParameterDefinition {
	instance := getPHASEStringMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASEStringMetaParameterDefinition](instance.ID, objc.Sel("initWithValue:identifier:"), objc.String(value), objc.String(identifier))
	rv.Autorelease()
	return rv
}


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameterDefinition) GlobalMetaParameters() string {
	rv := objc.Send[string](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}


// SetGlobalMetaParameters sets the value of the globalMetaParameters property.
// A dictionary of metaparameters that all sound event assets share.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameterDefinition) SetGlobalMetaParameters(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), objc.String(value))
}

// The object’s meta parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameterDefinition) MetaParameters() string {
	rv := objc.Send[string](p_.ID, objc.Sel("metaParameters"))
	return rv
}


// SetMetaParameters sets the value of the metaParameters property.
// The object’s meta parameters.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameterDefinition) SetMetaParameters(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), objc.String(value))
}


