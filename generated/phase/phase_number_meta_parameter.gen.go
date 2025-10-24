// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASENumberMetaParameter */


/* debug [class_header]: Header for PHASENumberMetaParameter */
// The class instance for the [PHASENumberMetaParameter] class.
var (
	PHASENumberMetaParameterClass     _PHASENumberMetaParameterClass
	PHASENumberMetaParameterClassOnce sync.Once
)

func getPHASENumberMetaParameterClass() _PHASENumberMetaParameterClass {
	PHASENumberMetaParameterClassOnce.Do(func() {
		PHASENumberMetaParameterClass = _PHASENumberMetaParameterClass{objc.GetClass("PHASENumberMetaParameter")}
	})
	return PHASENumberMetaParameterClass
}

type _PHASENumberMetaParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASENumberMetaParameter */
// An interface definition for the [PHASENumberMetaParameter] class.
type IPHASENumberMetaParameter interface {
	IPHASEMetaParameter
	
/* debug [class_interface_properties]: Properties for PHASENumberMetaParameter */
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

	
/* debug [class_interface_methods]: Methods for PHASENumberMetaParameter */
	// methods:
	FadeToValueDuration(value float64, duration float64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASENumberMetaParameter */
// Alloc allocates a new instance without initialization.
func (pc _PHASENumberMetaParameterClass) Alloc() PHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASENumberMetaParameterClass) New() PHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASENumberMetaParameter) Init() PHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASENumberMetaParameter) Autorelease() PHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASENumberMetaParameter creates a new PHASENumberMetaParameter instance.
func NewPHASENumberMetaParameter() PHASENumberMetaParameter {
	return getPHASENumberMetaParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASENumberMetaParameter */
// A metaparameter defined by a number that can change over time.
//
// This class contains a number that updates, like a “player speed” metaparameter that the app changes gradually from to . To create an instance of this class, first create a , and either: Register it with the engine by calling , then access the instance of this class in the engine’s dictionary. Pass it to the initializer, , and then access the instance of this class in a sound event’s dictionary. Use it as the input value for a by passing it into the initializer. Then, access the instance of this class using the mapped parameter’s property.


// A metaparameter defined by a number that can change over time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameter
type PHASENumberMetaParameter struct {
	PHASEMetaParameter
}

// PHASENumberMetaParameterFrom constructs a [PHASENumberMetaParameter] from an unsafe.Pointer.
//
// A metaparameter defined by a number that can change over time.
func PHASENumberMetaParameterFrom(ptr unsafe.Pointer) PHASENumberMetaParameter {
	return PHASENumberMetaParameter{
		PHASEMetaParameter: PHASEMetaParameterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASENumberMetaParameter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASENumberMetaParameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASENumberMetaParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASENumberMetaParameter */

// Sets the value gradually over the given amount of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameter/fade(value:duration:)
func (p_ PHASENumberMetaParameter) FadeToValueDuration(value float64, duration float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("fadeToValue:duration:"), value, duration)
}/* debug [instance_methods/method]: FadeToValueDuration */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASENumberMetaParameter */

// The highest possible number for the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameter/maximum
func (p_ PHASENumberMetaParameter) Maximum() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("maximum"))
	return rv
}/* debug [instance_properties/getter]: maximum */


// The lowest possible number for the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASENumberMetaParameter/minimum
func (p_ PHASENumberMetaParameter) Minimum() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minimum"))
	return rv
}/* debug [instance_properties/getter]: minimum */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASENumberMetaParameter) GlobalMetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}/* debug [instance_properties/getter]: globalMetaParameters */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASENumberMetaParameter) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}/* debug [instance_properties/setter]: globalMetaParameters */


// A linear input value to plot on a curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/inputmetaparameterdefinition
func (p_ PHASENumberMetaParameter) InputMetaParameterDefinition() IPHASENumberMetaParameterDefinition {
	rv := objc.Send[PHASENumberMetaParameterDefinition](p_.ID, objc.Sel("inputMetaParameterDefinition"))
	return rv
}/* debug [instance_properties/getter]: inputMetaParameterDefinition */


// A linear input value to plot on a curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasemappedmetaparameterdefinition/inputmetaparameterdefinition
func (p_ PHASENumberMetaParameter) SetInputMetaParameterDefinition(value IPHASENumberMetaParameterDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInputMetaParameterDefinition:"), value)
}/* debug [instance_properties/setter]: inputMetaParameterDefinition */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASENumberMetaParameter) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}/* debug [instance_properties/getter]: metaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASENumberMetaParameter) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}/* debug [instance_properties/setter]: metaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASENumberMetaParameter */



