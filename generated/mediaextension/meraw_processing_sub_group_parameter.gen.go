// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MERAWProcessingSubGroupParameter */


/* debug [class_header]: Header for MERAWProcessingSubGroupParameter */
// The class instance for the [MERAWProcessingSubGroupParameter] class.
var (
	MERAWProcessingSubGroupParameterClass     _MERAWProcessingSubGroupParameterClass
	MERAWProcessingSubGroupParameterClassOnce sync.Once
)

func getMERAWProcessingSubGroupParameterClass() _MERAWProcessingSubGroupParameterClass {
	MERAWProcessingSubGroupParameterClassOnce.Do(func() {
		MERAWProcessingSubGroupParameterClass = _MERAWProcessingSubGroupParameterClass{objc.GetClass("MERAWProcessingSubGroupParameter")}
	})
	return MERAWProcessingSubGroupParameterClass
}

type _MERAWProcessingSubGroupParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MERAWProcessingSubGroupParameter */
// An interface definition for the [MERAWProcessingSubGroupParameter] class.
type IMERAWProcessingSubGroupParameter interface {
	IMERAWProcessingParameter
	
/* debug [class_interface_properties]: Properties for MERAWProcessingSubGroupParameter */
	// properties:
	SubGroupParameters() []MERAWProcessingParameter
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MERAWProcessingSubGroupParameter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MERAWProcessingSubGroupParameter */
// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessingSubGroupParameterClass) Alloc() MERAWProcessingSubGroupParameter {
	rv := objc.Send[MERAWProcessingSubGroupParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MERAWProcessingSubGroupParameterClass) New() MERAWProcessingSubGroupParameter {
	rv := objc.Send[MERAWProcessingSubGroupParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessingSubGroupParameter) Init() MERAWProcessingSubGroupParameter {
	rv := objc.Send[MERAWProcessingSubGroupParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessingSubGroupParameter) Autorelease() MERAWProcessingSubGroupParameter {
	rv := objc.Send[MERAWProcessingSubGroupParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingSubGroupParameter creates a new MERAWProcessingSubGroupParameter instance.
func NewMERAWProcessingSubGroupParameter() MERAWProcessingSubGroupParameter {
	return getMERAWProcessingSubGroupParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MERAWProcessingSubGroupParameter */
// An object that describes a sub group parameter of a RAW processor.
//
// Sub groups are logical groupings of objects that should be displayed together in an application user interface.


// An object that describes a sub group parameter of a RAW processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/SubGroup
type MERAWProcessingSubGroupParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingSubGroupParameterFrom constructs a [MERAWProcessingSubGroupParameter] from an unsafe.Pointer.
//
// An object that describes a sub group parameter of a RAW processor.
func MERAWProcessingSubGroupParameterFrom(ptr unsafe.Pointer) MERAWProcessingSubGroupParameter {
	return MERAWProcessingSubGroupParameter{
		MERAWProcessingParameter: MERAWProcessingParameterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MERAWProcessingSubGroupParameter */

// Creates a sub group parameter object with the parameters value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/SubGroup/init(name:description:parameters:)
func NewMERAWProcessingSubGroupParameterWithNameDescriptionParameters(name objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, parameters []MERAWProcessingParameter) MERAWProcessingSubGroupParameter {
	instance := getMERAWProcessingSubGroupParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingSubGroupParameter](instance.ID, objc.Sel("initWithName:description:parameters:"), name, description, parameters)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingSubGroupParameterWithNameDescriptionParameters */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MERAWProcessingSubGroupParameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MERAWProcessingSubGroupParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MERAWProcessingSubGroupParameter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MERAWProcessingSubGroupParameter */

// The array of objects in the sub group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/SubGroup/subGroupParameters
func (m_ MERAWProcessingSubGroupParameter) SubGroupParameters() []MERAWProcessingParameter {
	rv := objc.Send[[]MERAWProcessingParameter](m_.ID, objc.Sel("subGroupParameters"))
	return rv
}/* debug [instance_properties/getter]: subGroupParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MERAWProcessingSubGroupParameter */


