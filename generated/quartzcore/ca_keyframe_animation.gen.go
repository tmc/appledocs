// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CAKeyframeAnimation */


/* debug [class_header]: Header for CAKeyframeAnimation */
// The class instance for the [KeyframeAnimation] class.
var (
	KeyframeAnimationClass     _KeyframeAnimationClass
	KeyframeAnimationClassOnce sync.Once
)

func getKeyframeAnimationClass() _KeyframeAnimationClass {
	KeyframeAnimationClassOnce.Do(func() {
		KeyframeAnimationClass = _KeyframeAnimationClass{objc.GetClass("CAKeyframeAnimation")}
	})
	return KeyframeAnimationClass
}

type _KeyframeAnimationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for KeyframeAnimation */
// An interface definition for the [KeyframeAnimation] class.
type IKeyframeAnimation interface {
	IPropertyAnimation
	
/* debug [class_interface_properties]: Properties for KeyframeAnimation */
	// properties:
	BiasValues() []foundation.Number
	SetBiasValues(value []foundation.Number)
	CalculationMode() AnimationCalculationMode /* typedef */
	SetCalculationMode(value AnimationCalculationMode /* typedef */)
	ContinuityValues() []foundation.Number
	SetContinuityValues(value []foundation.Number)
	KeyTimes() []foundation.Number
	SetKeyTimes(value []foundation.Number)
	Path() PathRef /* not a class type */
	SetPath(value PathRef /* not a class type */)
	RotationMode() AnimationRotationMode /* typedef */
	SetRotationMode(value AnimationRotationMode /* typedef */)
	TensionValues() []foundation.Number
	SetTensionValues(value []foundation.Number)
	TimingFunctions() []MediaTimingFunction
	SetTimingFunctions(value []MediaTimingFunction)
	Values() objc.IObject /* cross-framework: NSArray */
	SetValues(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for KeyframeAnimation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for KeyframeAnimation */
// Alloc allocates a new instance without initialization.
func (kc _KeyframeAnimationClass) Alloc() KeyframeAnimation {
	rv := objc.Send[KeyframeAnimation](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (kc _KeyframeAnimationClass) New() KeyframeAnimation {
	rv := objc.Send[KeyframeAnimation](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyframeAnimation) Init() KeyframeAnimation {
	rv := objc.Send[KeyframeAnimation](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyframeAnimation) Autorelease() KeyframeAnimation {
	rv := objc.Send[KeyframeAnimation](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyframeAnimation creates a new KeyframeAnimation instance.
func NewKeyframeAnimation() KeyframeAnimation {
	return getKeyframeAnimationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for KeyframeAnimation */
// An object that provides keyframe animation capabilities for a layer object.
//
// You create a object using the inherited method, specifying the key path of the property that you want to animate on the layer. You can then specify the keyframe values to use to control the timing and animation behavior. For most types of animations, you specify the keyframe values using the and properties. During the animation, Core Animation generates intermediate values by interpolating between the values you provide. When animating a value that is a coordinate point, such as the layer’s position, you can specify a for that point to follow instead of individual values. The pacing of the animation is controlled by the timing information you provide. The following code shows how to create a keyframe animation that animates a layer’s background color from red to green to blue over a two second duration.


// An object that provides keyframe animation capabilities for a layer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation
type KeyframeAnimation struct {
	PropertyAnimation
}

// KeyframeAnimationFrom constructs a [KeyframeAnimation] from an unsafe.Pointer.
//
// An object that provides keyframe animation capabilities for a layer object.
func KeyframeAnimationFrom(ptr unsafe.Pointer) KeyframeAnimation {
	return KeyframeAnimation{
		PropertyAnimation: PropertyAnimationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for KeyframeAnimation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for KeyframeAnimation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for KeyframeAnimation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for KeyframeAnimation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for KeyframeAnimation */

// An array of numbers that define the position of the curve relative to a control point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/biasValues
func (k_ KeyframeAnimation) BiasValues() []foundation.Number {
	rv := objc.Send[[]foundation.Number](k_.ID, objc.Sel("biasValues"))
	return rv
}/* debug [instance_properties/getter]: biasValues */


// An array of numbers that define the position of the curve relative to a control point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/biasValues
func (k_ KeyframeAnimation) SetBiasValues(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](k_.ID, objc.Sel("setBiasValues:"), nsArray)
}/* debug [instance_properties/setter]: biasValues */


// Specifies how intermediate keyframe values are calculated by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/calculationMode
func (k_ KeyframeAnimation) CalculationMode() AnimationCalculationMode /* typedef */ {
	rv := objc.Send[foundation.NSString](k_.ID, objc.Sel("calculationMode"))
	return rv
}/* debug [instance_properties/getter]: calculationMode */


// Specifies how intermediate keyframe values are calculated by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/calculationMode
func (k_ KeyframeAnimation) SetCalculationMode(value AnimationCalculationMode /* typedef */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setCalculationMode:"), value)
}/* debug [instance_properties/setter]: calculationMode */


// An array of numbers that define the sharpness of the timing curve’s corners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/continuityValues
func (k_ KeyframeAnimation) ContinuityValues() []foundation.Number {
	rv := objc.Send[[]foundation.Number](k_.ID, objc.Sel("continuityValues"))
	return rv
}/* debug [instance_properties/getter]: continuityValues */


// An array of numbers that define the sharpness of the timing curve’s corners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/continuityValues
func (k_ KeyframeAnimation) SetContinuityValues(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](k_.ID, objc.Sel("setContinuityValues:"), nsArray)
}/* debug [instance_properties/setter]: continuityValues */


// An optional array of objects that define the time at which to apply a given keyframe segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/keyTimes
func (k_ KeyframeAnimation) KeyTimes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](k_.ID, objc.Sel("keyTimes"))
	return rv
}/* debug [instance_properties/getter]: keyTimes */


// An optional array of objects that define the time at which to apply a given keyframe segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/keyTimes
func (k_ KeyframeAnimation) SetKeyTimes(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](k_.ID, objc.Sel("setKeyTimes:"), nsArray)
}/* debug [instance_properties/setter]: keyTimes */


// The path for a point-based property to follow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/path
func (k_ KeyframeAnimation) Path() PathRef /* not a class type */ {
	rv := objc.Send[PathRef](k_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// The path for a point-based property to follow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/path
func (k_ KeyframeAnimation) SetPath(value PathRef /* not a class type */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setPath:"), value)
}/* debug [instance_properties/setter]: path */


// Determines whether objects animating along the path rotate to match the path tangent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/rotationMode
func (k_ KeyframeAnimation) RotationMode() AnimationRotationMode /* typedef */ {
	rv := objc.Send[foundation.NSString](k_.ID, objc.Sel("rotationMode"))
	return rv
}/* debug [instance_properties/getter]: rotationMode */


// Determines whether objects animating along the path rotate to match the path tangent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/rotationMode
func (k_ KeyframeAnimation) SetRotationMode(value AnimationRotationMode /* typedef */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setRotationMode:"), value)
}/* debug [instance_properties/setter]: rotationMode */


// An array of numbers that define the tightness of the curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/tensionValues
func (k_ KeyframeAnimation) TensionValues() []foundation.Number {
	rv := objc.Send[[]foundation.Number](k_.ID, objc.Sel("tensionValues"))
	return rv
}/* debug [instance_properties/getter]: tensionValues */


// An array of numbers that define the tightness of the curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/tensionValues
func (k_ KeyframeAnimation) SetTensionValues(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](k_.ID, objc.Sel("setTensionValues:"), nsArray)
}/* debug [instance_properties/setter]: tensionValues */


// An optional array of objects that define the pacing for each keyframe segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/timingFunctions
func (k_ KeyframeAnimation) TimingFunctions() []MediaTimingFunction {
	rv := objc.Send[[]MediaTimingFunction](k_.ID, objc.Sel("timingFunctions"))
	return rv
}/* debug [instance_properties/getter]: timingFunctions */


// An optional array of objects that define the pacing for each keyframe segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/timingFunctions
func (k_ KeyframeAnimation) SetTimingFunctions(value []MediaTimingFunction) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](k_.ID, objc.Sel("setTimingFunctions:"), nsArray)
}/* debug [instance_properties/setter]: timingFunctions */


// An array of objects that specify the keyframe values to use for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/values
func (k_ KeyframeAnimation) Values() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](k_.ID, objc.Sel("values"))
	return rv
}/* debug [instance_properties/getter]: values */


// An array of objects that specify the keyframe values to use for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/values
func (k_ KeyframeAnimation) SetValues(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setValues:"), value)
}/* debug [instance_properties/setter]: values */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAKeyframeAnimation */



