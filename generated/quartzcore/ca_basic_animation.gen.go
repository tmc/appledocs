// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BasicAnimation] class.
var (
	BasicAnimationClass     _BasicAnimationClass
	BasicAnimationClassOnce sync.Once
)

func getBasicAnimationClass() _BasicAnimationClass {
	BasicAnimationClassOnce.Do(func() {
		BasicAnimationClass = _BasicAnimationClass{objc.GetClass("CABasicAnimation")}
	})
	return BasicAnimationClass
}

type _BasicAnimationClass struct {
	class objc.Class
}

// An interface definition for the [BasicAnimation] class.
type IBasicAnimation interface {
	IPropertyAnimation
	ByValue() objc.ID
	SetByValue(value objc.ID)
	FromValue() objc.ID
	SetFromValue(value objc.ID)
	ToValue() objc.ID
	SetToValue(value objc.ID)
	Opacity() float32
	SetOpacity(value float32)
	Transform() unsafe.Pointer
	SetTransform(value unsafe.Pointer)
}

// An object that provides basic, single-keyframe animation capabilities for a layer property.
//
// You create an instance of using the inherited method, specifying the key path of the property to be animated in the render tree. For example, you can animate a layer’s scalar (i.e. containing a single value) properties such as its . The following code fades in a layer by animating its opacity from to . Non-scalar properties, such as , can also be animated. Core Animation will interpolate between the color and the color. The animation created in the following code fades a layer’s background color from red to blue. If you want to animate the individual components of a non-scalar property with different values, you pass the values to and as arrays. The following animation moves a layer from to . The can access the individual components of a property. For example, the following animation stretches a layer by animating its object’s from to .
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation
type BasicAnimation struct {
	PropertyAnimation
}

// BasicAnimationFrom constructs a [BasicAnimation] from an unsafe.Pointer.
//
// An object that provides basic, single-keyframe animation capabilities for a layer property.
func BasicAnimationFrom(ptr unsafe.Pointer) BasicAnimation {
	return BasicAnimation{
		PropertyAnimation: PropertyAnimationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := objc.Send[BasicAnimation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BasicAnimationClass) New() BasicAnimation {
	rv := objc.Send[BasicAnimation](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BasicAnimation) Init() BasicAnimation {
	rv := objc.Send[BasicAnimation](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BasicAnimation) Autorelease() BasicAnimation {
	rv := objc.Send[BasicAnimation](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBasicAnimation creates a new BasicAnimation instance.
func NewBasicAnimation() BasicAnimation {
	return getBasicAnimationClass().New()
}


// Defines the value the receiver uses to perform relative interpolation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/byValue
func (b_ BasicAnimation) ByValue() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("byValue"))
	return rv
}


// SetByValue sets the value of the byValue property.
// Defines the value the receiver uses to perform relative interpolation.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/byValue
func (b_ BasicAnimation) SetByValue(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setByValue:"), value)
}

// Defines the value the receiver uses to start interpolation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/fromValue
func (b_ BasicAnimation) FromValue() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("fromValue"))
	return rv
}


// SetFromValue sets the value of the fromValue property.
// Defines the value the receiver uses to start interpolation.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/fromValue
func (b_ BasicAnimation) SetFromValue(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFromValue:"), value)
}

// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/toValue
func (b_ BasicAnimation) ToValue() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("toValue"))
	return rv
}


// SetToValue sets the value of the toValue property.
// Defines the value the receiver uses to end interpolation.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/toValue
func (b_ BasicAnimation) SetToValue(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setToValue:"), value)
}

// The opacity of the receiver. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/opacity
func (b_ BasicAnimation) Opacity() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("opacity"))
	return rv
}


// SetOpacity sets the value of the opacity property.
// The opacity of the receiver. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/opacity
func (b_ BasicAnimation) SetOpacity(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOpacity:"), value)
}

// The transform applied to the layer’s contents. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/transform
func (b_ BasicAnimation) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("transform"))
	return rv
}


// SetTransform sets the value of the transform property.
// The transform applied to the layer’s contents. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/transform
func (b_ BasicAnimation) SetTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransform:"), value)
}



