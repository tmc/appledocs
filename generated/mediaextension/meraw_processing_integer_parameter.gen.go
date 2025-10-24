// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MERAWProcessingIntegerParameter */


/* debug [class_header]: Header for MERAWProcessingIntegerParameter */
// The class instance for the [MERAWProcessingIntegerParameter] class.
var (
	MERAWProcessingIntegerParameterClass     _MERAWProcessingIntegerParameterClass
	MERAWProcessingIntegerParameterClassOnce sync.Once
)

func getMERAWProcessingIntegerParameterClass() _MERAWProcessingIntegerParameterClass {
	MERAWProcessingIntegerParameterClassOnce.Do(func() {
		MERAWProcessingIntegerParameterClass = _MERAWProcessingIntegerParameterClass{objc.GetClass("MERAWProcessingIntegerParameter")}
	})
	return MERAWProcessingIntegerParameterClass
}

type _MERAWProcessingIntegerParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MERAWProcessingIntegerParameter */
// An interface definition for the [MERAWProcessingIntegerParameter] class.
type IMERAWProcessingIntegerParameter interface {
	IMERAWProcessingParameter
	
/* debug [class_interface_properties]: Properties for MERAWProcessingIntegerParameter */
	// properties:
	CurrentValue() int
	SetCurrentValue(value int)
	InitialValue() int
	MaximumValue() int
	MinimumValue() int
	CameraValue() int
	SetCameraValue(value int)
	NeutralValue() int
	SetNeutralValue(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MERAWProcessingIntegerParameter */
	// methods:
	HasCameraValue(outCameraValue int) bool
	HasNeutralValue(outNeutralValue int) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MERAWProcessingIntegerParameter */
// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessingIntegerParameterClass) Alloc() MERAWProcessingIntegerParameter {
	rv := objc.Send[MERAWProcessingIntegerParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MERAWProcessingIntegerParameterClass) New() MERAWProcessingIntegerParameter {
	rv := objc.Send[MERAWProcessingIntegerParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessingIntegerParameter) Init() MERAWProcessingIntegerParameter {
	rv := objc.Send[MERAWProcessingIntegerParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessingIntegerParameter) Autorelease() MERAWProcessingIntegerParameter {
	rv := objc.Send[MERAWProcessingIntegerParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingIntegerParameter creates a new MERAWProcessingIntegerParameter instance.
func NewMERAWProcessingIntegerParameter() MERAWProcessingIntegerParameter {
	return getMERAWProcessingIntegerParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MERAWProcessingIntegerParameter */
// An object that describes an integer parameter of a RAW processor.


// An object that describes an integer parameter of a RAW processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer
type MERAWProcessingIntegerParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingIntegerParameterFrom constructs a [MERAWProcessingIntegerParameter] from an unsafe.Pointer.
//
// An object that describes an integer parameter of a RAW processor.
func MERAWProcessingIntegerParameterFrom(ptr unsafe.Pointer) MERAWProcessingIntegerParameter {
	return MERAWProcessingIntegerParameter{
		MERAWProcessingParameter: MERAWProcessingParameterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MERAWProcessingIntegerParameter */

// Creates a integer parameter object with the initial value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingIntegerParameter/initWithName:key:description:initialValue:maximum:minimum:
func NewMERAWProcessingIntegerParameterWithNameKeyDescriptionInitialValueMaximumMinimum(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue int, maximum int, minimum int) MERAWProcessingIntegerParameter {
	instance := getMERAWProcessingIntegerParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingIntegerParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:maximum:minimum:"), name, key, description, initialValue, maximum, minimum)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingIntegerParameterWithNameKeyDescriptionInitialValueMaximumMinimum */


// Creates an integer parameter object with the initial and camera values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingIntegerParameter/initWithName:key:description:initialValue:maximum:minimum:cameraValue:
func NewMERAWProcessingIntegerParameterWithNameKeyDescriptionInitialValueMaximumMinimumCameraValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue int, maximum int, minimum int, cameraValue int) MERAWProcessingIntegerParameter {
	instance := getMERAWProcessingIntegerParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingIntegerParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:maximum:minimum:cameraValue:"), name, key, description, initialValue, maximum, minimum, cameraValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingIntegerParameterWithNameKeyDescriptionInitialValueMaximumMinimumCameraValue */


// Creates an integer parameter object with the initial and neutral values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingIntegerParameter/initWithName:key:description:initialValue:maximum:minimum:neutralValue:
func NewMERAWProcessingIntegerParameterWithNameKeyDescriptionInitialValueMaximumMinimumNeutralValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue int, maximum int, minimum int, neutralValue int) MERAWProcessingIntegerParameter {
	instance := getMERAWProcessingIntegerParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingIntegerParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:maximum:minimum:neutralValue:"), name, key, description, initialValue, maximum, minimum, neutralValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingIntegerParameterWithNameKeyDescriptionInitialValueMaximumMinimumNeutralValue */


// Creates an integer parameter object with the initial, neutral, and camera values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingIntegerParameter/initWithName:key:description:initialValue:maximum:minimum:neutralValue:cameraValue:
func NewMERAWProcessingIntegerParameterWithNameKeyDescriptionInitialValueMaximumMinimumNeutralValueCameraValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue int, maximum int, minimum int, neutralValue int, cameraValue int) MERAWProcessingIntegerParameter {
	instance := getMERAWProcessingIntegerParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingIntegerParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:maximum:minimum:neutralValue:cameraValue:"), name, key, description, initialValue, maximum, minimum, neutralValue, cameraValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingIntegerParameterWithNameKeyDescriptionInitialValueMaximumMinimumNeutralValueCameraValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MERAWProcessingIntegerParameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MERAWProcessingIntegerParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MERAWProcessingIntegerParameter */

// The optional camera value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingIntegerParameter/hasCameraValue:
func (m_ MERAWProcessingIntegerParameter) HasCameraValue(outCameraValue int) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasCameraValue:"), outCameraValue)
	return rv
}/* debug [instance_methods/method]: HasCameraValue */


// The optional neutral value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingIntegerParameter/hasNeutralValue:
func (m_ MERAWProcessingIntegerParameter) HasNeutralValue(outNeutralValue int) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasNeutralValue:"), outNeutralValue)
	return rv
}/* debug [instance_methods/method]: HasNeutralValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MERAWProcessingIntegerParameter */

// Get or set the current value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer/currentValue
func (m_ MERAWProcessingIntegerParameter) CurrentValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("currentValue"))
	return rv
}/* debug [instance_properties/getter]: currentValue */


// Get or set the current value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer/currentValue
func (m_ MERAWProcessingIntegerParameter) SetCurrentValue(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentValue:"), value)
}/* debug [instance_properties/setter]: currentValue */


// The initial value for this parameter as defined in the sequence metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer/initialValue
func (m_ MERAWProcessingIntegerParameter) InitialValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("initialValue"))
	return rv
}/* debug [instance_properties/getter]: initialValue */


// The maximum value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer/maximumValue
func (m_ MERAWProcessingIntegerParameter) MaximumValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maximumValue"))
	return rv
}/* debug [instance_properties/getter]: maximumValue */


// The minimum value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer/minimumValue
func (m_ MERAWProcessingIntegerParameter) MinimumValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("minimumValue"))
	return rv
}/* debug [instance_properties/getter]: minimumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/integer/cameravalue
func (m_ MERAWProcessingIntegerParameter) CameraValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cameraValue"))
	return rv
}/* debug [instance_properties/getter]: cameraValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/integer/cameravalue
func (m_ MERAWProcessingIntegerParameter) SetCameraValue(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraValue:"), value)
}/* debug [instance_properties/setter]: cameraValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/integer/neutralvalue
func (m_ MERAWProcessingIntegerParameter) NeutralValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("neutralValue"))
	return rv
}/* debug [instance_properties/getter]: neutralValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/integer/neutralvalue
func (m_ MERAWProcessingIntegerParameter) SetNeutralValue(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeutralValue:"), value)
}/* debug [instance_properties/setter]: neutralValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MERAWProcessingIntegerParameter */


