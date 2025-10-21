// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASESwitchNodeDefinition] class.
var (
	PHASESwitchNodeDefinitionClass     _PHASESwitchNodeDefinitionClass
	PHASESwitchNodeDefinitionClassOnce sync.Once
)

func getPHASESwitchNodeDefinitionClass() _PHASESwitchNodeDefinitionClass {
	PHASESwitchNodeDefinitionClassOnce.Do(func() {
		PHASESwitchNodeDefinitionClass = _PHASESwitchNodeDefinitionClass{objc.GetClass("PHASESwitchNodeDefinition")}
	})
	return PHASESwitchNodeDefinitionClass
}

type _PHASESwitchNodeDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASESwitchNodeDefinition] class.
type IPHASESwitchNodeDefinition interface {
	IPHASESoundEventNodeDefinition
}

// A node that passes invocation to only one of its child nodes.
//
// A switch node takes a different path in a sound-event hierarchy depending on the value that the app supplies for the node’s switch metaparameter. You define the available paths ahead of time by calling at least twice and supplying the subtree’s unique string name as the switch value. When your app invokes a sound event at runtime, PHASE checks the value of to determine which path to take.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESwitchNodeDefinition
type PHASESwitchNodeDefinition struct {
	PHASESoundEventNodeDefinition
}

// PHASESwitchNodeDefinitionFrom constructs a [PHASESwitchNodeDefinition] from an unsafe.Pointer.
//
// A node that passes invocation to only one of its child nodes.
func PHASESwitchNodeDefinitionFrom(ptr unsafe.Pointer) PHASESwitchNodeDefinition {
	return PHASESwitchNodeDefinition{
		PHASESoundEventNodeDefinition: PHASESoundEventNodeDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASESwitchNodeDefinitionClass) Alloc() PHASESwitchNodeDefinition {
	rv := objc.Send[PHASESwitchNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASESwitchNodeDefinitionClass) New() PHASESwitchNodeDefinition {
	rv := objc.Send[PHASESwitchNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASESwitchNodeDefinition) Init() PHASESwitchNodeDefinition {
	rv := objc.Send[PHASESwitchNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASESwitchNodeDefinition) Autorelease() PHASESwitchNodeDefinition {
	rv := objc.Send[PHASESwitchNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASESwitchNodeDefinition creates a new PHASESwitchNodeDefinition instance.
func NewPHASESwitchNodeDefinition() PHASESwitchNodeDefinition {
	return getPHASESwitchNodeDefinitionClass().New()
}




// Creates a node that invokes a child node based on the value of the given parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESwitchNodeDefinition/init(switchMetaParameterDefinition:)
func NewPHASESwitchNodeDefinitionWithSwitchMetaParameterDefinition(switchMetaParameterDefinition unsafe.Pointer) PHASESwitchNodeDefinition {
	instance := getPHASESwitchNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASESwitchNodeDefinition](instance.ID, objc.Sel("initWithSwitchMetaParameterDefinition:"), switchMetaParameterDefinition)
	rv.Autorelease()
	return rv
}


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASESwitchNodeDefinition) GlobalMetaParameters() string {
	rv := objc.Send[string](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}


// SetGlobalMetaParameters sets the value of the globalMetaParameters property.
// A dictionary of metaparameters that all sound event assets share.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASESwitchNodeDefinition) SetGlobalMetaParameters(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), objc.String(value))
}

// The object’s meta parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASESwitchNodeDefinition) MetaParameters() string {
	rv := objc.Send[string](p_.ID, objc.Sel("metaParameters"))
	return rv
}


// SetMetaParameters sets the value of the metaParameters property.
// The object’s meta parameters.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASESwitchNodeDefinition) SetMetaParameters(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), objc.String(value))
}

// The meta parameter that holds the name of the child node to invoke.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseswitchnodedefinition/switchmetaparameterdefinition
func (p_ PHASESwitchNodeDefinition) SwitchMetaParameterDefinition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("switchMetaParameterDefinition"))
	return rv
}


// SetSwitchMetaParameterDefinition sets the value of the switchMetaParameterDefinition property.
// The meta parameter that holds the name of the child node to invoke.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseswitchnodedefinition/switchmetaparameterdefinition
func (p_ PHASESwitchNodeDefinition) SetSwitchMetaParameterDefinition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSwitchMetaParameterDefinition:"), value)
}


