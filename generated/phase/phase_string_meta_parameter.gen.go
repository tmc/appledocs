// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEStringMetaParameter */


/* debug [class_header]: Header for PHASEStringMetaParameter */
// The class instance for the [PHASEStringMetaParameter] class.
var (
	PHASEStringMetaParameterClass     _PHASEStringMetaParameterClass
	PHASEStringMetaParameterClassOnce sync.Once
)

func getPHASEStringMetaParameterClass() _PHASEStringMetaParameterClass {
	PHASEStringMetaParameterClassOnce.Do(func() {
		PHASEStringMetaParameterClass = _PHASEStringMetaParameterClass{objc.GetClass("PHASEStringMetaParameter")}
	})
	return PHASEStringMetaParameterClass
}

type _PHASEStringMetaParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEStringMetaParameter */
// An interface definition for the [PHASEStringMetaParameter] class.
type IPHASEStringMetaParameter interface {
	IPHASEMetaParameter
	
/* debug [class_interface_properties]: Properties for PHASEStringMetaParameter */
	// properties:
	GlobalMetaParameters() IPHASEMetaParameter
	SetGlobalMetaParameters(value IPHASEMetaParameter)
	MetaParameters() IPHASEMetaParameter
	SetMetaParameters(value IPHASEMetaParameter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEStringMetaParameter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEStringMetaParameter */
// Alloc allocates a new instance without initialization.
func (pc _PHASEStringMetaParameterClass) Alloc() PHASEStringMetaParameter {
	rv := objc.Send[PHASEStringMetaParameter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEStringMetaParameterClass) New() PHASEStringMetaParameter {
	rv := objc.Send[PHASEStringMetaParameter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEStringMetaParameter) Init() PHASEStringMetaParameter {
	rv := objc.Send[PHASEStringMetaParameter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEStringMetaParameter) Autorelease() PHASEStringMetaParameter {
	rv := objc.Send[PHASEStringMetaParameter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEStringMetaParameter creates a new PHASEStringMetaParameter instance.
func NewPHASEStringMetaParameter() PHASEStringMetaParameter {
	return getPHASEStringMetaParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEStringMetaParameter */
// A metaparameter with a text definition that can change over time.
//
// This class contains text that updates, like a “weather” metaparameter that the app changes from “rainy” to “sunny.” To create an instance of this class, first create a , and either: Register it with the engine by calling , then access the instance of this class in the engine’s dictionary. Pass it to the initializer, , and then access the instance of this class in a sound event’s dictionary.


// A metaparameter with a text definition that can change over time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStringMetaParameter
type PHASEStringMetaParameter struct {
	PHASEMetaParameter
}

// PHASEStringMetaParameterFrom constructs a [PHASEStringMetaParameter] from an unsafe.Pointer.
//
// A metaparameter with a text definition that can change over time.
func PHASEStringMetaParameterFrom(ptr unsafe.Pointer) PHASEStringMetaParameter {
	return PHASEStringMetaParameter{
		PHASEMetaParameter: PHASEMetaParameterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEStringMetaParameter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEStringMetaParameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEStringMetaParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEStringMetaParameter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEStringMetaParameter */

// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameter) GlobalMetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("globalMetaParameters"))
	return rv
}/* debug [instance_properties/getter]: globalMetaParameters */


// A dictionary of metaparameters that all sound event assets share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseassetregistry/globalmetaparameters
func (p_ PHASEStringMetaParameter) SetGlobalMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGlobalMetaParameters:"), value)
}/* debug [instance_properties/setter]: globalMetaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameter) MetaParameters() IPHASEMetaParameter {
	rv := objc.Send[PHASEMetaParameter](p_.ID, objc.Sel("metaParameters"))
	return rv
}/* debug [instance_properties/getter]: metaParameters */


// The object’s meta parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/metaparameters
func (p_ PHASEStringMetaParameter) SetMetaParameters(value IPHASEMetaParameter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetaParameters:"), value)
}/* debug [instance_properties/setter]: metaParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEStringMetaParameter */



