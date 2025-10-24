// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MERAWProcessingFloatParameter */


/* debug [class_header]: Header for MERAWProcessingFloatParameter */
// The class instance for the [MERAWProcessingFloatParameter] class.
var (
	MERAWProcessingFloatParameterClass     _MERAWProcessingFloatParameterClass
	MERAWProcessingFloatParameterClassOnce sync.Once
)

func getMERAWProcessingFloatParameterClass() _MERAWProcessingFloatParameterClass {
	MERAWProcessingFloatParameterClassOnce.Do(func() {
		MERAWProcessingFloatParameterClass = _MERAWProcessingFloatParameterClass{objc.GetClass("MERAWProcessingFloatParameter")}
	})
	return MERAWProcessingFloatParameterClass
}

type _MERAWProcessingFloatParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MERAWProcessingFloatParameter */
// An interface definition for the [MERAWProcessingFloatParameter] class.
type IMERAWProcessingFloatParameter interface {
	IMERAWProcessingParameter
	
/* debug [class_interface_properties]: Properties for MERAWProcessingFloatParameter */
	// properties:
	CurrentValue() float32
	SetCurrentValue(value float32)
	InitialValue() float32
	MaximumValue() float32
	MinimumValue() float32
	CameraValue() float32
	SetCameraValue(value float32)
	NeutralValue() float32
	SetNeutralValue(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MERAWProcessingFloatParameter */
	// methods:
	HasCameraValue(outCameraValue unsafe.Pointer) bool
	HasNeutralValue(outNeutralValue unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MERAWProcessingFloatParameter */
// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessingFloatParameterClass) Alloc() MERAWProcessingFloatParameter {
	rv := objc.Send[MERAWProcessingFloatParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MERAWProcessingFloatParameterClass) New() MERAWProcessingFloatParameter {
	rv := objc.Send[MERAWProcessingFloatParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessingFloatParameter) Init() MERAWProcessingFloatParameter {
	rv := objc.Send[MERAWProcessingFloatParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessingFloatParameter) Autorelease() MERAWProcessingFloatParameter {
	rv := objc.Send[MERAWProcessingFloatParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingFloatParameter creates a new MERAWProcessingFloatParameter instance.
func NewMERAWProcessingFloatParameter() MERAWProcessingFloatParameter {
	return getMERAWProcessingFloatParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MERAWProcessingFloatParameter */
// An object that describes a floating-point parameter of a RAW processor.


// An object that describes a floating-point parameter of a RAW processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint
type MERAWProcessingFloatParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingFloatParameterFrom constructs a [MERAWProcessingFloatParameter] from an unsafe.Pointer.
//
// An object that describes a floating-point parameter of a RAW processor.
func MERAWProcessingFloatParameterFrom(ptr unsafe.Pointer) MERAWProcessingFloatParameter {
	return MERAWProcessingFloatParameter{
		MERAWProcessingParameter: MERAWProcessingParameterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MERAWProcessingFloatParameter */

// Creates a floating-point parameter object with the initial value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingFloatParameter/initWithName:key:description:initialValue:maximum:minimum:
func NewMERAWProcessingFloatParameterWithNameKeyDescriptionInitialValueMaximumMinimum(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue float32, maximum float32, minimum float32) MERAWProcessingFloatParameter {
	instance := getMERAWProcessingFloatParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingFloatParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:maximum:minimum:"), name, key, description, initialValue, maximum, minimum)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingFloatParameterWithNameKeyDescriptionInitialValueMaximumMinimum */


// Creates a floating-point parameter object with the initial and camera values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingFloatParameter/initWithName:key:description:initialValue:maximum:minimum:cameraValue:
func NewMERAWProcessingFloatParameterWithNameKeyDescriptionInitialValueMaximumMinimumCameraValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue float32, maximum float32, minimum float32, cameraValue float32) MERAWProcessingFloatParameter {
	instance := getMERAWProcessingFloatParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingFloatParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:maximum:minimum:cameraValue:"), name, key, description, initialValue, maximum, minimum, cameraValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingFloatParameterWithNameKeyDescriptionInitialValueMaximumMinimumCameraValue */


// Creates a floating-point parameter object with the initial and neutral values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingFloatParameter/initWithName:key:description:initialValue:maximum:minimum:neutralValue:
func NewMERAWProcessingFloatParameterWithNameKeyDescriptionInitialValueMaximumMinimumNeutralValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue float32, maximum float32, minimum float32, neutralValue float32) MERAWProcessingFloatParameter {
	instance := getMERAWProcessingFloatParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingFloatParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:maximum:minimum:neutralValue:"), name, key, description, initialValue, maximum, minimum, neutralValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingFloatParameterWithNameKeyDescriptionInitialValueMaximumMinimumNeutralValue */


// Creates a floating-point parameter object with the initial, neutral, and camera values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingFloatParameter/initWithName:key:description:initialValue:maximum:minimum:neutralValue:cameraValue:
func NewMERAWProcessingFloatParameterWithNameKeyDescriptionInitialValueMaximumMinimumNeutralValueCameraValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, initialValue float32, maximum float32, minimum float32, neutralValue float32, cameraValue float32) MERAWProcessingFloatParameter {
	instance := getMERAWProcessingFloatParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingFloatParameter](instance.ID, objc.Sel("initWithName:key:description:initialValue:maximum:minimum:neutralValue:cameraValue:"), name, key, description, initialValue, maximum, minimum, neutralValue, cameraValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingFloatParameterWithNameKeyDescriptionInitialValueMaximumMinimumNeutralValueCameraValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MERAWProcessingFloatParameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MERAWProcessingFloatParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MERAWProcessingFloatParameter */

// The optional camera value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingFloatParameter/hasCameraValue:
func (m_ MERAWProcessingFloatParameter) HasCameraValue(outCameraValue unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasCameraValue:"), outCameraValue)
	return rv
}/* debug [instance_methods/method]: HasCameraValue */


// The optional neutral value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingFloatParameter/hasNeutralValue:
func (m_ MERAWProcessingFloatParameter) HasNeutralValue(outNeutralValue unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasNeutralValue:"), outNeutralValue)
	return rv
}/* debug [instance_methods/method]: HasNeutralValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MERAWProcessingFloatParameter */

// Get or set the current value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint/currentValue
func (m_ MERAWProcessingFloatParameter) CurrentValue() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("currentValue"))
	return rv
}/* debug [instance_properties/getter]: currentValue */


// Get or set the current value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint/currentValue
func (m_ MERAWProcessingFloatParameter) SetCurrentValue(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentValue:"), value)
}/* debug [instance_properties/setter]: currentValue */


// The initial value for this parameter as defined in the sequence metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint/initialValue
func (m_ MERAWProcessingFloatParameter) InitialValue() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("initialValue"))
	return rv
}/* debug [instance_properties/getter]: initialValue */


// The maximum value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint/maximumValue
func (m_ MERAWProcessingFloatParameter) MaximumValue() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("maximumValue"))
	return rv
}/* debug [instance_properties/getter]: maximumValue */


// The minimum value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint/minimumValue
func (m_ MERAWProcessingFloatParameter) MinimumValue() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("minimumValue"))
	return rv
}/* debug [instance_properties/getter]: minimumValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/floatingpoint/cameravalue
func (m_ MERAWProcessingFloatParameter) CameraValue() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("cameraValue"))
	return rv
}/* debug [instance_properties/getter]: cameraValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/floatingpoint/cameravalue
func (m_ MERAWProcessingFloatParameter) SetCameraValue(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraValue:"), value)
}/* debug [instance_properties/setter]: cameraValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/floatingpoint/neutralvalue
func (m_ MERAWProcessingFloatParameter) NeutralValue() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("neutralValue"))
	return rv
}/* debug [instance_properties/getter]: neutralValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/floatingpoint/neutralvalue
func (m_ MERAWProcessingFloatParameter) SetNeutralValue(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeutralValue:"), value)
}/* debug [instance_properties/setter]: neutralValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MERAWProcessingFloatParameter */


