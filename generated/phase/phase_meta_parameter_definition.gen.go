// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEMetaParameterDefinition */


/* debug [class_header]: Header for PHASEMetaParameterDefinition */
// The class instance for the [PHASEMetaParameterDefinition] class.
var (
	PHASEMetaParameterDefinitionClass     _PHASEMetaParameterDefinitionClass
	PHASEMetaParameterDefinitionClassOnce sync.Once
)

func getPHASEMetaParameterDefinitionClass() _PHASEMetaParameterDefinitionClass {
	PHASEMetaParameterDefinitionClassOnce.Do(func() {
		PHASEMetaParameterDefinitionClass = _PHASEMetaParameterDefinitionClass{objc.GetClass("PHASEMetaParameterDefinition")}
	})
	return PHASEMetaParameterDefinitionClass
}

type _PHASEMetaParameterDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEMetaParameterDefinition */
// An interface definition for the [PHASEMetaParameterDefinition] class.
type IPHASEMetaParameterDefinition interface {
	IPHASEDefinition
	
/* debug [class_interface_properties]: Properties for PHASEMetaParameterDefinition */
	// properties:
	Value() objc.ID
	GlobalMetaParameters() IPHASEMetaParameter
	SetGlobalMetaParameters(value IPHASEMetaParameter)
	MetaParameters() IPHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEMetaParameterDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEMetaParameterDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEMetaParameterDefinitionClass) Alloc() PHASEMetaParameterDefinition {
	rv := objc.Send[PHASEMetaParameterDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEMetaParameterDefinitionClass) New() PHASEMetaParameterDefinition {
	rv := objc.Send[PHASEMetaParameterDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMetaParameterDefinition) Init() PHASEMetaParameterDefinition {
	rv := objc.Send[PHASEMetaParameterDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMetaParameterDefinition) Autorelease() PHASEMetaParameterDefinition {
	rv := objc.Send[PHASEMetaParameterDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMetaParameterDefinition creates a new PHASEMetaParameterDefinition instance.
func NewPHASEMetaParameterDefinition() PHASEMetaParameterDefinition {
	return getPHASEMetaParameterDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEMetaParameterDefinition */
// A specification for a named parameter with a constant value.
//
// Instances of this class provide an app with dynamic control of various properies of live audio playback. This is a base class for the and subclasses. You create a single instance of one of the subclasses for a specific property you want to adjust. By passing the definition subclass to the framework, you spawn one or more objects for discrete use across your app. For example, when you initialize a with a number metaparameter definition, PHASE registers a in the corresponding sound event’s dictionary. Put metaparameter definitions that you wish to share across different sounds in the asset registery’s dictionary. To add global metaparameter definitions to the dictionary, call . Then pass the definition into several sound event node defintions, such as .


// A specification for a named parameter with a constant value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMetaParameterDefinition
type PHASEMetaParameterDefinition struct {
	PHASEDefinition
}

// PHASEMetaParameterDefinitionFrom constructs a [PHASEMetaParameterDefinition] from an unsafe.Pointer.
//
// A specification for a named parameter with a constant value.
func PHASEMetaParameterDefinitionFrom(ptr unsafe.Pointer) PHASEMetaParameterDefinition {
	return PHASEMetaParameterDefinition{
		PHASEDefinition: PHASEDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEMetaParameterDefinition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEMetaParameterDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEMetaParameterDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEMetaParameterDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEMetaParameterDefinition */

// A constant value for the parameter definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMetaParameterDefinition/value
func (p_ PHASEMetaParameterDefinition) Value() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEMetaParameterDefinition) GlobalMetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}/* debug [instance_properties/getter]: globalMetaParameters */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEMetaParameterDefinition) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}/* debug [instance_properties/setter]: globalMetaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEMetaParameterDefinition) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}/* debug [instance_properties/getter]: metaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEMetaParameterDefinition) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}/* debug [instance_properties/setter]: metaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEMetaParameterDefinition */



