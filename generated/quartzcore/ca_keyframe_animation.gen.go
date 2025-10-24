// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [KeyframeAnimation] class.
type IKeyframeAnimation interface {
	IPropertyAnimation
	// properties:
	TimingFunctions() []IMediaTimingFunction
	SetTimingFunctions(value []IMediaTimingFunction)
	BiasValues() objc.IObject /* cross-framework: NSNumber */
	SetBiasValues(value objc.IObject /* cross-framework: NSNumber */)
	CalculationMode() AnimationCalculationMode /* not a class type */
	SetCalculationMode(value AnimationCalculationMode /* not a class type */)
	ContinuityValues() objc.IObject /* cross-framework: NSNumber */
	SetContinuityValues(value objc.IObject /* cross-framework: NSNumber */)
	KeyTimes() objc.IObject /* cross-framework: NSNumber */
	SetKeyTimes(value objc.IObject /* cross-framework: NSNumber */)
	Path() objectivec.IObject
	SetPath(value objectivec.IObject)
	RotationMode() AnimationRotationMode /* not a class type */
	SetRotationMode(value AnimationRotationMode /* not a class type */)
	TensionValues() objc.IObject /* cross-framework: NSNumber */
	SetTensionValues(value objc.IObject /* cross-framework: NSNumber */)
	Values() unsafe.Pointer
	SetValues(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (kc _KeyframeAnimationClass) Alloc() KeyframeAnimation {
	rv := objc.Send[KeyframeAnimation](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An optional array of objects that define the pacing for each keyframe segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/timingFunctions
func (k_ KeyframeAnimation) TimingFunctions() []IMediaTimingFunction {
	rv := objc.Send[[]MediaTimingFunction](k_.ID, objc.Sel("timingFunctions"))
	return rv
}


// An optional array of objects that define the pacing for each keyframe segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/timingFunctions
func (k_ KeyframeAnimation) SetTimingFunctions(value []IMediaTimingFunction) {
	// Convert Go slice to NSArray
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
}


// An array of numbers that define the position of the curve relative to a control point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/biasvalues
func (k_ KeyframeAnimation) BiasValues() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](k_.ID, objc.Sel("biasValues"))
	return rv
}


// An array of numbers that define the position of the curve relative to a control point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/biasvalues
func (k_ KeyframeAnimation) SetBiasValues(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setBiasValues:"), value)
}


// Specifies how intermediate keyframe values are calculated by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/calculationmode
func (k_ KeyframeAnimation) CalculationMode() AnimationCalculationMode /* not a class type */ {
	rv := objc.Send[AnimationCalculationMode](k_.ID, objc.Sel("calculationMode"))
	return rv
}


// Specifies how intermediate keyframe values are calculated by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/calculationmode
func (k_ KeyframeAnimation) SetCalculationMode(value AnimationCalculationMode /* not a class type */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setCalculationMode:"), value)
}


// An array of numbers that define the sharpness of the timing curve’s corners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/continuityvalues
func (k_ KeyframeAnimation) ContinuityValues() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](k_.ID, objc.Sel("continuityValues"))
	return rv
}


// An array of numbers that define the sharpness of the timing curve’s corners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/continuityvalues
func (k_ KeyframeAnimation) SetContinuityValues(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setContinuityValues:"), value)
}


// An optional array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/keytimes
func (k_ KeyframeAnimation) KeyTimes() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](k_.ID, objc.Sel("keyTimes"))
	return rv
}


// An optional array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/keytimes
func (k_ KeyframeAnimation) SetKeyTimes(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setKeyTimes:"), value)
}


// The path for a point-based property to follow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/path
func (k_ KeyframeAnimation) Path() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](k_.ID, objc.Sel("path"))
	return rv
}


// The path for a point-based property to follow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/path
func (k_ KeyframeAnimation) SetPath(value objectivec.IObject) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setPath:"), value)
}


// Determines whether objects animating along the path rotate to match the path tangent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/rotationmode
func (k_ KeyframeAnimation) RotationMode() AnimationRotationMode /* not a class type */ {
	rv := objc.Send[AnimationRotationMode](k_.ID, objc.Sel("rotationMode"))
	return rv
}


// Determines whether objects animating along the path rotate to match the path tangent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/rotationmode
func (k_ KeyframeAnimation) SetRotationMode(value AnimationRotationMode /* not a class type */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setRotationMode:"), value)
}


// An array of numbers that define the tightness of the curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/tensionvalues
func (k_ KeyframeAnimation) TensionValues() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](k_.ID, objc.Sel("tensionValues"))
	return rv
}


// An array of numbers that define the tightness of the curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/tensionvalues
func (k_ KeyframeAnimation) SetTensionValues(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setTensionValues:"), value)
}


// An array of objects that specify the keyframe values to use for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/values
func (k_ KeyframeAnimation) Values() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](k_.ID, objc.Sel("values"))
	return rv
}


// An array of objects that specify the keyframe values to use for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cakeyframeanimation/values
func (k_ KeyframeAnimation) SetValues(value unsafe.Pointer) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setValues:"), value)
}



