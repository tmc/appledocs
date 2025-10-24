// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ByValue() unsafe.Pointer
	SetByValue(value unsafe.Pointer)
	FromValue() unsafe.Pointer
	SetFromValue(value unsafe.Pointer)
	ToValue() unsafe.Pointer
	SetToValue(value unsafe.Pointer)
	BackgroundColor() objectivec.IObject
	SetBackgroundColor(value objectivec.IObject)
	Opacity() float32
	SetOpacity(value float32)
	Transform() CATransform3D /* not a class type */
	SetTransform(value CATransform3D /* not a class type */)
	// methods:
}

// An object that provides basic, single-keyframe animation capabilities for a layer property.
//
// You create an instance of using the inherited method, specifying the key path of the property to be animated in the render tree. For example, you can animate a layer’s scalar (i.e. containing a single value) properties such as its . The following code fades in a layer by animating its opacity from to . Non-scalar properties, such as , can also be animated. Core Animation will interpolate between the color and the color. The animation created in the following code fades a layer’s background color from red to blue. If you want to animate the individual components of a non-scalar property with different values, you pass the values to and as arrays. The following animation moves a layer from to . The can access the individual components of a property. For example, the following animation stretches a layer by animating its object’s from to .


// An object that provides basic, single-keyframe animation capabilities for a layer property.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/byvalue
func (b_ BasicAnimation) ByValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("byValue"))
	return rv
}


// Defines the value the receiver uses to perform relative interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/byvalue
func (b_ BasicAnimation) SetByValue(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setByValue:"), value)
}


// Defines the value the receiver uses to start interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/fromvalue
func (b_ BasicAnimation) FromValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("fromValue"))
	return rv
}


// Defines the value the receiver uses to start interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/fromvalue
func (b_ BasicAnimation) SetFromValue(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFromValue:"), value)
}


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/tovalue
func (b_ BasicAnimation) ToValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("toValue"))
	return rv
}


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/tovalue
func (b_ BasicAnimation) SetToValue(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setToValue:"), value)
}


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (b_ BasicAnimation) BackgroundColor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (b_ BasicAnimation) SetBackgroundColor(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The opacity of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/opacity
func (b_ BasicAnimation) Opacity() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("opacity"))
	return rv
}


// The opacity of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/opacity
func (b_ BasicAnimation) SetOpacity(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOpacity:"), value)
}


// The transform applied to the layer’s contents. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/transform
func (b_ BasicAnimation) Transform() CATransform3D /* not a class type */ {
	rv := objc.Send[Transform3D](b_.ID, objc.Sel("transform"))
	return rv
}


// The transform applied to the layer’s contents. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/transform
func (b_ BasicAnimation) SetTransform(value CATransform3D /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransform:"), value)
}



