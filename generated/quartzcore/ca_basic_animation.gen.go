// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CABasicAnimation */


/* debug [class_header]: Header for CABasicAnimation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BasicAnimation */
// An interface definition for the [BasicAnimation] class.
type IBasicAnimation interface {
	IPropertyAnimation
	
/* debug [class_interface_properties]: Properties for BasicAnimation */
	// properties:
	ByValue() objc.ID
	SetByValue(value objc.ID)
	FromValue() objc.ID
	SetFromValue(value objc.ID)
	ToValue() objc.ID
	SetToValue(value objc.ID)
	BackgroundColor() objectivec.IObject
	SetBackgroundColor(value objectivec.IObject)
	Opacity() float32
	SetOpacity(value float32)
	Transform() objc.IObject /* cross-framework: CATransform3D */
	SetTransform(value objc.IObject /* cross-framework: CATransform3D */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BasicAnimation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BasicAnimation */
// Alloc allocates a new instance without initialization.
func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := objc.Send[BasicAnimation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BasicAnimation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BasicAnimation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BasicAnimation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BasicAnimation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BasicAnimation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BasicAnimation */

// Defines the value the receiver uses to perform relative interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/byValue
func (b_ BasicAnimation) ByValue() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("byValue"))
	return rv
}/* debug [instance_properties/getter]: byValue */


// Defines the value the receiver uses to perform relative interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/byValue
func (b_ BasicAnimation) SetByValue(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setByValue:"), value)
}/* debug [instance_properties/setter]: byValue */


// Defines the value the receiver uses to start interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/fromValue
func (b_ BasicAnimation) FromValue() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("fromValue"))
	return rv
}/* debug [instance_properties/getter]: fromValue */


// Defines the value the receiver uses to start interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/fromValue
func (b_ BasicAnimation) SetFromValue(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFromValue:"), value)
}/* debug [instance_properties/setter]: fromValue */


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/toValue
func (b_ BasicAnimation) ToValue() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("toValue"))
	return rv
}/* debug [instance_properties/getter]: toValue */


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/toValue
func (b_ BasicAnimation) SetToValue(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setToValue:"), value)
}/* debug [instance_properties/setter]: toValue */


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (b_ BasicAnimation) BackgroundColor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (b_ BasicAnimation) SetBackgroundColor(value objectivec.IObject) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The opacity of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/opacity
func (b_ BasicAnimation) Opacity() float32 {
	rv := objc.Send[float32](b_.ID, objc.Sel("opacity"))
	return rv
}/* debug [instance_properties/getter]: opacity */


// The opacity of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/opacity
func (b_ BasicAnimation) SetOpacity(value float32) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setOpacity:"), value)
}/* debug [instance_properties/setter]: opacity */


// The transform applied to the layer’s contents. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/transform
func (b_ BasicAnimation) Transform() objc.IObject /* cross-framework: CATransform3D */ {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// The transform applied to the layer’s contents. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/transform
func (b_ BasicAnimation) SetTransform(value objc.IObject /* cross-framework: CATransform3D */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CABasicAnimation */



