// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASESwitchNodeDefinition */


/* debug [class_header]: Header for PHASESwitchNodeDefinition */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASESwitchNodeDefinition */
// An interface definition for the [PHASESwitchNodeDefinition] class.
type IPHASESwitchNodeDefinition interface {
	IPHASESoundEventNodeDefinition
	
/* debug [class_interface_properties]: Properties for PHASESwitchNodeDefinition */
	// properties:
	SwitchMetaParameterDefinition() IPHASEStringMetaParameterDefinition
	GlobalMetaParameters() IPHASEMetaParameter
	SetGlobalMetaParameters(value IPHASEMetaParameter)
	MetaParameters() IPHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASESwitchNodeDefinition */
	// methods:
	AddSubtreeSwitchValue(subtree IPHASESoundEventNodeDefinition, switchValue objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASESwitchNodeDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASESwitchNodeDefinitionClass) Alloc() PHASESwitchNodeDefinition {
	rv := objc.Send[PHASESwitchNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASESwitchNodeDefinition */
// A node that passes invocation to only one of its child nodes.
//
// A switch node takes a different path in a sound-event hierarchy depending on the value that the app supplies for the node’s switch metaparameter. You define the available paths ahead of time by calling at least twice and supplying the subtree’s unique string name as the switch value. When your app invokes a sound event at runtime, PHASE checks the value of to determine which path to take.


// A node that passes invocation to only one of its child nodes.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASESwitchNodeDefinition */

// Creates a node that invokes a child node based on the value of the given parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESwitchNodeDefinition/init(switchMetaParameterDefinition:)
func NewPHASESwitchNodeDefinitionWithSwitchMetaParameterDefinition(switchMetaParameterDefinition IPHASEStringMetaParameterDefinition) PHASESwitchNodeDefinition {
	instance := getPHASESwitchNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASESwitchNodeDefinition](instance.ID, objc.Sel("initWithSwitchMetaParameterDefinition:"), switchMetaParameterDefinition)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESwitchNodeDefinitionWithSwitchMetaParameterDefinition */


// Creates a named node that invokes a child node based on the value of the given parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESwitchNodeDefinition/init(switchMetaParameterDefinition:identifier:)
func NewPHASESwitchNodeDefinitionWithSwitchMetaParameterDefinitionIdentifier(switchMetaParameterDefinition IPHASEStringMetaParameterDefinition, identifier objc.IObject /* cross-framework: NSString */) PHASESwitchNodeDefinition {
	instance := getPHASESwitchNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASESwitchNodeDefinition](instance.ID, objc.Sel("initWithSwitchMetaParameterDefinition:identifier:"), switchMetaParameterDefinition, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASESwitchNodeDefinitionWithSwitchMetaParameterDefinitionIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASESwitchNodeDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASESwitchNodeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASESwitchNodeDefinition */

// Adds a child node with the given switch value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESwitchNodeDefinition/addSubtree(_:switchValue:)
func (p_ PHASESwitchNodeDefinition) AddSubtreeSwitchValue(subtree IPHASESoundEventNodeDefinition, switchValue objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSubtree:switchValue:"), subtree, switchValue)
}/* debug [instance_methods/method]: AddSubtreeSwitchValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASESwitchNodeDefinition */

// The meta parameter that holds the name of the child node to invoke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASESwitchNodeDefinition/switchMetaParameterDefinition
func (p_ PHASESwitchNodeDefinition) SwitchMetaParameterDefinition() IPHASEStringMetaParameterDefinition {
	rv := objc.Send[PHASEStringMetaParameterDefinition](p_.ID, objc.Sel("switchMetaParameterDefinition"))
	return rv
}/* debug [instance_properties/getter]: switchMetaParameterDefinition */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASESwitchNodeDefinition) GlobalMetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}/* debug [instance_properties/getter]: globalMetaParameters */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASESwitchNodeDefinition) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}/* debug [instance_properties/setter]: globalMetaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASESwitchNodeDefinition) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}/* debug [instance_properties/getter]: metaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASESwitchNodeDefinition) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}/* debug [instance_properties/setter]: metaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASESwitchNodeDefinition */


