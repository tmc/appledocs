// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MERAWProcessingParameter */


/* debug [class_header]: Header for MERAWProcessingParameter */
// The class instance for the [MERAWProcessingParameter] class.
var (
	MERAWProcessingParameterClass     _MERAWProcessingParameterClass
	MERAWProcessingParameterClassOnce sync.Once
)

func getMERAWProcessingParameterClass() _MERAWProcessingParameterClass {
	MERAWProcessingParameterClassOnce.Do(func() {
		MERAWProcessingParameterClass = _MERAWProcessingParameterClass{objc.GetClass("MERAWProcessingParameter")}
	})
	return MERAWProcessingParameterClass
}

type _MERAWProcessingParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MERAWProcessingParameter */
// An interface definition for the [MERAWProcessingParameter] class.
type IMERAWProcessingParameter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MERAWProcessingParameter */
	// properties:
	Enabled() bool
	SetEnabled(value bool)
	Key() objc.IObject /* cross-framework: NSString */
	LongDescription() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	ProcessingParameters() IMERAWProcessingParameter
	SetProcessingParameters(value IMERAWProcessingParameter)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MERAWProcessingParameter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MERAWProcessingParameter */
// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessingParameterClass) Alloc() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MERAWProcessingParameterClass) New() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessingParameter) Init() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessingParameter) Autorelease() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingParameter creates a new MERAWProcessingParameter instance.
func NewMERAWProcessingParameter() MERAWProcessingParameter {
	return getMERAWProcessingParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MERAWProcessingParameter */
// An object for the RAW processor to describe each processing parameter the processor exposes.
//
// This protocol provides an interface for Video Toolbox to query descriptions of the different parameters that can be used to influence RAW processor operation. A distinct is created for each parameter supported by the RAW processor, and the set of supported parameters is returned by the interface.


// An object for the RAW processor to describe each processing parameter the processor exposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter
type MERAWProcessingParameter struct {
	objectivec.Object
}

// MERAWProcessingParameterFrom constructs a [MERAWProcessingParameter] from an unsafe.Pointer.
//
// An object for the RAW processor to describe each processing parameter the processor exposes.
func MERAWProcessingParameterFrom(ptr unsafe.Pointer) MERAWProcessingParameter {
	return MERAWProcessingParameter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MERAWProcessingParameter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MERAWProcessingParameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MERAWProcessingParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MERAWProcessingParameter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MERAWProcessingParameter */

// A Boolean value that indicates whether the extension enables the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/enabled
func (m_ MERAWProcessingParameter) Enabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether the extension enables the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/enabled
func (m_ MERAWProcessingParameter) SetEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A unique key string identifying the parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/key
func (m_ MERAWProcessingParameter) Key() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("key"))
	return rv
}/* debug [instance_properties/getter]: key */


// A localized description of the parameter, suitable for displaying in a tool tip or similar explanatory UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/longDescription
func (m_ MERAWProcessingParameter) LongDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("longDescription"))
	return rv
}/* debug [instance_properties/getter]: longDescription */


// A localized human-readable name for the parameter, suitable for displaying in application UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/name
func (m_ MERAWProcessingParameter) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// Provides a list of processing parameters that can be changed by the client of Video Toolbox session to influence processing behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessor/processingparameters
func (m_ MERAWProcessingParameter) ProcessingParameters() IMERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](m_.ID, objc.Sel("processingParameters"))
	return rv
}/* debug [instance_properties/getter]: processingParameters */


// Provides a list of processing parameters that can be changed by the client of Video Toolbox session to influence processing behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessor/processingparameters
func (m_ MERAWProcessingParameter) SetProcessingParameters(value IMERAWProcessingParameter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProcessingParameters:"), value)
}/* debug [instance_properties/setter]: processingParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MERAWProcessingParameter */



