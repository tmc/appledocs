// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASENumberMetaParameterDefinition */


/* debug [class_header]: Header for PHASENumberMetaParameterDefinition */
// The class instance for the [PHASENumberMetaParameterDefinition] class.
var (
	PHASENumberMetaParameterDefinitionClass     _PHASENumberMetaParameterDefinitionClass
	PHASENumberMetaParameterDefinitionClassOnce sync.Once
)

func getPHASENumberMetaParameterDefinitionClass() _PHASENumberMetaParameterDefinitionClass {
	PHASENumberMetaParameterDefinitionClassOnce.Do(func() {
		PHASENumberMetaParameterDefinitionClass = _PHASENumberMetaParameterDefinitionClass{objc.GetClass("PHASENumberMetaParameterDefinition")}
	})
	return PHASENumberMetaParameterDefinitionClass
}

type _PHASENumberMetaParameterDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASENumberMetaParameterDefinition */
// An interface definition for the [PHASENumberMetaParameterDefinition] class.
type IPHASENumberMetaParameterDefinition interface {
	IPHASEMetaParameterDefinition
	
/* debug [class_interface_properties]: Properties for PHASENumberMetaParameterDefinition */
	// properties:
	Maximum() float64
	Minimum() float64
	GlobalMetaParameters() IPHASEMetaParameter
	SetGlobalMetaParameters(value IPHASEMetaParameter)
	InputMetaParameterDefinition() IPHASENumberMetaParameterDefinition
	SetInputMetaParameterDefinition(value IPHASENumberMetaParameterDefinition)
	MetaParameters() IPHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASENumberMetaParameterDefinition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASENumberMetaParameterDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASENumberMetaParameterDefinitionClass) Alloc() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASENumberMetaParameterDefinitionClass) New() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASENumberMetaParameterDefinition) Init() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASENumberMetaParameterDefinition) Autorelease() PHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASENumberMetaParameterDefinition creates a new PHASENumberMetaParameterDefinition instance.
func NewPHASENumberMetaParameterDefinition() PHASENumberMetaParameterDefinition {
	return getPHASENumberMetaParameterDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASENumberMetaParameterDefinition */
// A specification for a metaparameter defined by a number.
//
// Use this class to spawn discrete instances of , for example, a “player speed” metaparameter that the app changes gradually from to . To use a number metaparameter, create an instance of this class and: Register it with the engine by calling , then access the instance of this class in the engine’s dictionary. Pass it to the initializer, , and then access the instance of this class in a sound event’s dictionary. Pass it into the initializer, . Then, access the instance of this class using the mapped parameter’s property.


// A specification for a metaparameter defined by a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameterDefinition
type PHASENumberMetaParameterDefinition struct {
	PHASEMetaParameterDefinition
}

// PHASENumberMetaParameterDefinitionFrom constructs a [PHASENumberMetaParameterDefinition] from an unsafe.Pointer.
//
// A specification for a metaparameter defined by a number.
func PHASENumberMetaParameterDefinitionFrom(ptr unsafe.Pointer) PHASENumberMetaParameterDefinition {
	return PHASENumberMetaParameterDefinition{
		PHASEMetaParameterDefinition: PHASEMetaParameterDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASENumberMetaParameterDefinition */

// Creates a specification for a metaparameter with the given numeric value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameterDefinition/init(value:)
func NewPHASENumberMetaParameterDefinitionWithValue(value float64) PHASENumberMetaParameterDefinition {
	instance := getPHASENumberMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASENumberMetaParameterDefinition](instance.ID, objc.Sel("initWithValue:"), value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASENumberMetaParameterDefinitionWithValue */


// Creates a specification for a named metaparameter with the given numeric value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameterDefinition/init(value:identifier:)
func NewPHASENumberMetaParameterDefinitionWithValueIdentifier(value float64, identifier objc.IObject /* cross-framework: NSString */) PHASENumberMetaParameterDefinition {
	instance := getPHASENumberMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASENumberMetaParameterDefinition](instance.ID, objc.Sel("initWithValue:identifier:"), value, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASENumberMetaParameterDefinitionWithValueIdentifier */


// Creates a specification for a metaparameter with the given numeric value and range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameterDefinition/init(value:minimum:maximum:)
func NewPHASENumberMetaParameterDefinitionWithValueMinimumMaximum(value float64, minimum float64, maximum float64) PHASENumberMetaParameterDefinition {
	instance := getPHASENumberMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASENumberMetaParameterDefinition](instance.ID, objc.Sel("initWithValue:minimum:maximum:"), value, minimum, maximum)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASENumberMetaParameterDefinitionWithValueMinimumMaximum */


// Creates a specification for a named metaparameter with the given numeric value and range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameterDefinition/init(value:minimum:maximum:identifier:)
func NewPHASENumberMetaParameterDefinitionWithValueMinimumMaximumIdentifier(value float64, minimum float64, maximum float64, identifier objc.IObject /* cross-framework: NSString */) PHASENumberMetaParameterDefinition {
	instance := getPHASENumberMetaParameterDefinitionClass().Alloc()
	rv := objc.Send[PHASENumberMetaParameterDefinition](instance.ID, objc.Sel("initWithValue:minimum:maximum:identifier:"), value, minimum, maximum, identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASENumberMetaParameterDefinitionWithValueMinimumMaximumIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASENumberMetaParameterDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASENumberMetaParameterDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASENumberMetaParameterDefinition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASENumberMetaParameterDefinition */

// The highest possible number for the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameterDefinition/maximum
func (p_ PHASENumberMetaParameterDefinition) Maximum() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("maximum"))
	return rv
}/* debug [instance_properties/getter]: maximum */


// The lowest possible number for the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameterDefinition/minimum
func (p_ PHASENumberMetaParameterDefinition) Minimum() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minimum"))
	return rv
}/* debug [instance_properties/getter]: minimum */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASENumberMetaParameterDefinition) GlobalMetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}/* debug [instance_properties/getter]: globalMetaParameters */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASENumberMetaParameterDefinition) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}/* debug [instance_properties/setter]: globalMetaParameters */


// A linear input value to plot on a curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/inputmetaparameterdefinition
func (p_ PHASENumberMetaParameterDefinition) InputMetaParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("inputMetaParameterDefinition"))
	return rv
}/* debug [instance_properties/getter]: inputMetaParameterDefinition */


// A linear input value to plot on a curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/inputmetaparameterdefinition
func (p_ PHASENumberMetaParameterDefinition) SetInputMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInputMetaParameterDefinition:"), value)
}/* debug [instance_properties/setter]: inputMetaParameterDefinition */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASENumberMetaParameterDefinition) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}/* debug [instance_properties/getter]: metaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASENumberMetaParameterDefinition) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}/* debug [instance_properties/setter]: metaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASENumberMetaParameterDefinition */


