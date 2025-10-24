// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASEStringMetaParameterDefinition */


/* debug [class_header]: Header for PHASEStringMetaParameterDefinition */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEStringMetaParameterDefinition */
// An interface definition for the [PHASEStringMetaParameterDefinition] class.
type IPHASEStringMetaParameterDefinition interface {
	IPHASEMetaParameterDefinition
	
/* debug [class_interface_properties]: Properties for PHASEStringMetaParameterDefinition */
	// properties:
	GlobalMetaParameters() IPHASEMetaParameter
	SetGlobalMetaParameters(value IPHASEMetaParameter)
	MetaParameters() IPHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEStringMetaParameterDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEStringMetaParameterDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEStringMetaParameterDefinitionClass) Alloc() PHASEStringMetaParameterDefinition {
	rv := objc.Send[PHASEStringMetaParameterDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEStringMetaParameterDefinition */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEStringMetaParameterDefinition */

// Creates a specification for a textual metaparameter with the given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStringMetaParameterDefinition/init(value:)
func NewPHASEStringMetaParameterDefinitionWithValue(value objc.IObject /* cross-framework: NSString */) PHASEStringMetaParameterDefinition {
	instance := getPHASEStringMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASEStringMetaParameterDefinition](instance.ID, objc.Sel("initWithValue:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEStringMetaParameterDefinitionWithValue */


// Creates a specification for a named textual metaparameter with the given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStringMetaParameterDefinition/init(value:identifier:)
func NewPHASEStringMetaParameterDefinitionWithValueIdentifier(value objc.IObject /* cross-framework: NSString */, identifier objc.IObject /* cross-framework: NSString */) PHASEStringMetaParameterDefinition {
	instance := getPHASEStringMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASEStringMetaParameterDefinition](instance.ID, objc.Sel("initWithValue:identifier:"), value, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEStringMetaParameterDefinitionWithValueIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEStringMetaParameterDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEStringMetaParameterDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEStringMetaParameterDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEStringMetaParameterDefinition */

// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameterDefinition) GlobalMetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}/* debug [instance_properties/getter]: globalMetaParameters */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameterDefinition) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}/* debug [instance_properties/setter]: globalMetaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameterDefinition) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}/* debug [instance_properties/getter]: metaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameterDefinition) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}/* debug [instance_properties/setter]: metaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEStringMetaParameterDefinition */


