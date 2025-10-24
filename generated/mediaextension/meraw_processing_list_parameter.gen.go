// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MERAWProcessingListParameter */


/* debug [class_header]: Header for MERAWProcessingListParameter */
// The class instance for the [MERAWProcessingListParameter] class.
var (
	MERAWProcessingListParameterClass     _MERAWProcessingListParameterClass
	MERAWProcessingListParameterClassOnce sync.Once
)

func getMERAWProcessingListParameterClass() _MERAWProcessingListParameterClass {
	MERAWProcessingListParameterClassOnce.Do(func() {
		MERAWProcessingListParameterClass = _MERAWProcessingListParameterClass{objc.GetClass("MERAWProcessingListParameter")}
	})
	return MERAWProcessingListParameterClass
}

type _MERAWProcessingListParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MERAWProcessingListParameter */
// An interface definition for the [MERAWProcessingListParameter] class.
type IMERAWProcessingListParameter interface {
	IMERAWProcessingParameter
	
/* debug [class_interface_properties]: Properties for MERAWProcessingListParameter */
	// properties:
	CurrentValue() int
	SetCurrentValue(value int)
	InitialValue() int
	ListElements() []MERAWProcessingListElementParameter
	CameraValue() int
	SetCameraValue(value int)
	NeutralValue() int
	SetNeutralValue(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MERAWProcessingListParameter */
	// methods:
	HasCameraValue(outCameraValue int) bool
	HasNeutralValue(outNeutralValue int) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MERAWProcessingListParameter */
// Alloc allocates a new instance without initialization.
func (mc _MERAWProcessingListParameterClass) Alloc() MERAWProcessingListParameter {
	rv := objc.Send[MERAWProcessingListParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MERAWProcessingListParameterClass) New() MERAWProcessingListParameter {
	rv := objc.Send[MERAWProcessingListParameter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MERAWProcessingListParameter) Init() MERAWProcessingListParameter {
	rv := objc.Send[MERAWProcessingListParameter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MERAWProcessingListParameter) Autorelease() MERAWProcessingListParameter {
	rv := objc.Send[MERAWProcessingListParameter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingListParameter creates a new MERAWProcessingListParameter instance.
func NewMERAWProcessingListParameter() MERAWProcessingListParameter {
	return getMERAWProcessingListParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MERAWProcessingListParameter */
// An object that describes a list parameter of a RAW processor.


// An object that describes a list parameter of a RAW processor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List
type MERAWProcessingListParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingListParameterFrom constructs a [MERAWProcessingListParameter] from an unsafe.Pointer.
//
// An object that describes a list parameter of a RAW processor.
func MERAWProcessingListParameterFrom(ptr unsafe.Pointer) MERAWProcessingListParameter {
	return MERAWProcessingListParameter{
		MERAWProcessingParameter: MERAWProcessingParameterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MERAWProcessingListParameter */

// Creates a list parameter object with the initial value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingListParameter/initWithName:key:description:list:initialValue:
func NewMERAWProcessingListParameterWithNameKeyDescriptionListInitialValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, listElements []MERAWProcessingListElementParameter, initialValue int) MERAWProcessingListParameter {
	instance := getMERAWProcessingListParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingListParameter](instance.ID, objc.Sel("initWithName:key:description:list:initialValue:"), name, key, description, listElements, initialValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingListParameterWithNameKeyDescriptionListInitialValue */


// Creates a list parameter object with the initial and camera values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingListParameter/initWithName:key:description:list:initialValue:cameraValue:
func NewMERAWProcessingListParameterWithNameKeyDescriptionListInitialValueCameraValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, listElements []MERAWProcessingListElementParameter, initialValue int, cameraValue int) MERAWProcessingListParameter {
	instance := getMERAWProcessingListParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingListParameter](instance.ID, objc.Sel("initWithName:key:description:list:initialValue:cameraValue:"), name, key, description, listElements, initialValue, cameraValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingListParameterWithNameKeyDescriptionListInitialValueCameraValue */


// Creates a list parameter object with the initial and neutral values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingListParameter/initWithName:key:description:list:initialValue:neutralValue:
func NewMERAWProcessingListParameterWithNameKeyDescriptionListInitialValueNeutralValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, listElements []MERAWProcessingListElementParameter, initialValue int, neutralValue int) MERAWProcessingListParameter {
	instance := getMERAWProcessingListParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingListParameter](instance.ID, objc.Sel("initWithName:key:description:list:initialValue:neutralValue:"), name, key, description, listElements, initialValue, neutralValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingListParameterWithNameKeyDescriptionListInitialValueNeutralValue */


// Creates a list parameter object with the initial, neutral, and camera values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingListParameter/initWithName:key:description:list:initialValue:neutralValue:cameraValue:
func NewMERAWProcessingListParameterWithNameKeyDescriptionListInitialValueNeutralValueCameraValue(name objc.IObject /* cross-framework: NSString */, key objc.IObject /* cross-framework: NSString */, description objc.IObject /* cross-framework: NSString */, listElements []MERAWProcessingListElementParameter, initialValue int, neutralValue int, cameraValue int) MERAWProcessingListParameter {
	instance := getMERAWProcessingListParameterClass().Alloc()
	rv := objc.Send[MERAWProcessingListParameter](instance.ID, objc.Sel("initWithName:key:description:list:initialValue:neutralValue:cameraValue:"), name, key, description, listElements, initialValue, neutralValue, cameraValue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMERAWProcessingListParameterWithNameKeyDescriptionListInitialValueNeutralValueCameraValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MERAWProcessingListParameter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MERAWProcessingListParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MERAWProcessingListParameter */

// The optional camera value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingListParameter/hasCameraValue:
func (m_ MERAWProcessingListParameter) HasCameraValue(outCameraValue int) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasCameraValue:"), outCameraValue)
	return rv
}/* debug [instance_methods/method]: HasCameraValue */


// The optional neutral value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingListParameter/hasNeutralValue:
func (m_ MERAWProcessingListParameter) HasNeutralValue(outNeutralValue int) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasNeutralValue:"), outNeutralValue)
	return rv
}/* debug [instance_methods/method]: HasNeutralValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MERAWProcessingListParameter */

// Get or set the current value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List/currentValue
func (m_ MERAWProcessingListParameter) CurrentValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("currentValue"))
	return rv
}/* debug [instance_properties/getter]: currentValue */


// Get or set the current value for this parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List/currentValue
func (m_ MERAWProcessingListParameter) SetCurrentValue(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentValue:"), value)
}/* debug [instance_properties/setter]: currentValue */


// The initial value for this parameter as defined in the sequence metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List/initialValue
func (m_ MERAWProcessingListParameter) InitialValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("initialValue"))
	return rv
}/* debug [instance_properties/getter]: initialValue */


// The ordered array of which make up this list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List/listElements
func (m_ MERAWProcessingListParameter) ListElements() []MERAWProcessingListElementParameter {
	rv := objc.Send[[]MERAWProcessingListElementParameter](m_.ID, objc.Sel("listElements"))
	return rv
}/* debug [instance_properties/getter]: listElements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/list/cameravalue
func (m_ MERAWProcessingListParameter) CameraValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("cameraValue"))
	return rv
}/* debug [instance_properties/getter]: cameraValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/list/cameravalue
func (m_ MERAWProcessingListParameter) SetCameraValue(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCameraValue:"), value)
}/* debug [instance_properties/setter]: cameraValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/list/neutralvalue
func (m_ MERAWProcessingListParameter) NeutralValue() int {
	rv := objc.Send[int](m_.ID, objc.Sel("neutralValue"))
	return rv
}/* debug [instance_properties/getter]: neutralValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/merawprocessingparameter/list/neutralvalue
func (m_ MERAWProcessingListParameter) SetNeutralValue(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNeutralValue:"), value)
}/* debug [instance_properties/setter]: neutralValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MERAWProcessingListParameter */


