// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MERAWProcessingBooleanParameter */


/* debug [class_header]: Header for MERAWProcessingBooleanParameter */
// The class instance for the [MERAWProcessingBooleanParameter] class.
var (
	MERAWProcessingBooleanParameterClass     _MERAWProcessingBooleanParameterClass
	MERAWProcessingBooleanParameterClassOnce sync.Once
)

func getMERAWProcessingBooleanParameterClass() _MERAWProcessingBooleanParameterClass {
	MERAWProcessingBooleanParameterClassOnce.Do(func() {
		MERAWProcessingBooleanParameterClass = _MERAWProcessingBooleanParameterClass{objc.GetClass("MERAWProcessingBooleanParameter")}
	})
	return MERAWProcessingBooleanParameterClass
}

type _MERAWProcessingBooleanParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MERAWProcessingBooleanParameter */
// An interface definition for the [MERAWProcessingBooleanParameter] class.
type IMERAWProcessingBooleanParameter interface {
	IMERAWProcessingParameter
	
/* debug [class_interface_properties]: Properties for MERAWProcessingBooleanParameter */
	// properties:
	CurrentValue() bool
	SetCurrentValue(value bool)
	InitialValue() bool
	CameraValue() bool
	SetCameraValue(value bool)
	NeutralValue() bool
	SetNeutralValue(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MERAWProcessingBooleanParameter */
	// methods:
	HasCameraValue(outCameraValue bool) bool
	HasNeutralValue(outNeutralValue bool) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MERAWProcessingBooleanParameter */
// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessingBooleanParameterClass) Alloc() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MERAWProcessingBooleanParameterClass) New() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessingBooleanParameter) Init() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessingBooleanParameter) Autorelease() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingBooleanParameter creates a new MERAWProcessingBooleanParameter instance.
func NewMERAWProcessingBooleanParameter() MERAWProcessingBooleanParameter {
	return getMERAWProcessingBooleanParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MERAWProcessingBooleanParameter */
// An object that describes a Boolean parameter of a RAW processor.


// An object that describes a Boolean parameter of a RAW processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Boolean
type MERAWProcessingBooleanParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingBooleanParameterFrom constructs a [MERAWProcessingBooleanParameter] from an unsafe.Pointer.
//
// An object that describes a Boolean parameter of a RAW processor.
func MERAWProcessingBooleanParameterFrom(ptr unsafe.Pointer) MERAWProcessingBooleanParameter {
	return MERAWProcessingBooleanParameter{
		MERAWProcessingParameter: MERAWProcessingParameterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MERAWProcessingBooleanParameter */

// Creates a Boolean parameter object with the initial value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingBooleanParameter/initWithName:key:description:initialValue:
func NewMERAWProcessingBooleanParameterWithNameKeyDescriptionInitialValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue bool) MERAWProcessingBooleanParameter {
	instance := getMERAWProcessingBooleanParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingBooleanParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:"), name, key, description, initialValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingBooleanParameterWithNameKeyDescriptionInitialValue */


// Creates a Boolean parameter object with the initial and camera values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingBooleanParameter/initWithName:key:description:initialValue:cameraValue:
func NewMERAWProcessingBooleanParameterWithNameKeyDescriptionInitialValueCameraValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue bool, cameraValue bool) MERAWProcessingBooleanParameter {
	instance := getMERAWProcessingBooleanParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingBooleanParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:cameraValue:"), name, key, description, initialValue, cameraValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingBooleanParameterWithNameKeyDescriptionInitialValueCameraValue */


// Creates a Boolean parameter object with the initial and neutral values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingBooleanParameter/initWithName:key:description:initialValue:neutralValue:
func NewMERAWProcessingBooleanParameterWithNameKeyDescriptionInitialValueNeutralValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue bool, neutralValue bool) MERAWProcessingBooleanParameter {
	instance := getMERAWProcessingBooleanParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingBooleanParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:neutralValue:"), name, key, description, initialValue, neutralValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingBooleanParameterWithNameKeyDescriptionInitialValueNeutralValue */


// Creates a Boolean parameter object with the initial, neutral, and camera values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingBooleanParameter/initWithName:key:description:initialValue:neutralValue:cameraValue:
func NewMERAWProcessingBooleanParameterWithNameKeyDescriptionInitialValueNeutralValueCameraValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue bool, neutralValue bool, cameraValue bool) MERAWProcessingBooleanParameter {
	instance := getMERAWProcessingBooleanParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingBooleanParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:neutralValue:cameraValue:"), name, key, description, initialValue, neutralValue, cameraValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingBooleanParameterWithNameKeyDescriptionInitialValueNeutralValueCameraValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MERAWProcessingBooleanParameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MERAWProcessingBooleanParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MERAWProcessingBooleanParameter */

// The optional camera value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingBooleanParameter/hasCameraValue:
func (m_ MERAWProcessingBooleanParameter) HasCameraValue(outCameraValue bool) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasCameraValue:"), outCameraValue)
	return rv
}/* debug [instance_methods/method]: HasCameraValue */


// The optional neutral value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingBooleanParameter/hasNeutralValue:
func (m_ MERAWProcessingBooleanParameter) HasNeutralValue(outNeutralValue bool) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasNeutralValue:"), outNeutralValue)
	return rv
}/* debug [instance_methods/method]: HasNeutralValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MERAWProcessingBooleanParameter */

// Get or set the current value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Boolean/currentValue
func (m_ MERAWProcessingBooleanParameter) CurrentValue() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("currentValue"))
	return rv
}/* debug [instance_properties/getter]: currentValue */


// Get or set the current value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Boolean/currentValue
func (m_ MERAWProcessingBooleanParameter) SetCurrentValue(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentValue:"), value)
}/* debug [instance_properties/setter]: currentValue */


// The initial value for this parameter as defined in the sequence metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Boolean/initialValue
func (m_ MERAWProcessingBooleanParameter) InitialValue() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("initialValue"))
	return rv
}/* debug [instance_properties/getter]: initialValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/boolean/cameravalue
func (m_ MERAWProcessingBooleanParameter) CameraValue() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("cameraValue"))
	return rv
}/* debug [instance_properties/getter]: cameraValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/boolean/cameravalue
func (m_ MERAWProcessingBooleanParameter) SetCameraValue(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraValue:"), value)
}/* debug [instance_properties/setter]: cameraValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/boolean/neutralvalue
func (m_ MERAWProcessingBooleanParameter) NeutralValue() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("neutralValue"))
	return rv
}/* debug [instance_properties/getter]: neutralValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/boolean/neutralvalue
func (m_ MERAWProcessingBooleanParameter) SetNeutralValue(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeutralValue:"), value)
}/* debug [instance_properties/setter]: neutralValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MERAWProcessingBooleanParameter */


