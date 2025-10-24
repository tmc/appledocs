// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAPropertyAnimation */


/* debug [class_header]: Header for CAPropertyAnimation */
// The class instance for the [PropertyAnimation] class.
var (
	PropertyAnimationClass     _PropertyAnimationClass
	PropertyAnimationClassOnce sync.Once
)

func getPropertyAnimationClass() _PropertyAnimationClass {
	PropertyAnimationClassOnce.Do(func() {
		PropertyAnimationClass = _PropertyAnimationClass{objc.GetClass("CAPropertyAnimation")}
	})
	return PropertyAnimationClass
}

type _PropertyAnimationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PropertyAnimation */
// An interface definition for the [PropertyAnimation] class.
type IPropertyAnimation interface {
	IAnimation
	
/* debug [class_interface_properties]: Properties for PropertyAnimation */
	// properties:
	Additive() bool
	SetAdditive(value bool)
	Cumulative() bool
	SetCumulative(value bool)
	KeyPath() objc.IObject /* cross-framework: NSString */
	SetKeyPath(value objc.IObject /* cross-framework: NSString */)
	ValueFunction() IValueFunction
	SetValueFunction(value IValueFunction)
	IsAdditive() bool
	SetIsAdditive(value bool)
	IsCumulative() bool
	SetIsCumulative(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PropertyAnimation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PropertyAnimation */
// Alloc allocates a new instance without initialization.
func (pc _PropertyAnimationClass) Alloc() PropertyAnimation {
	rv := objc.Send[PropertyAnimation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PropertyAnimationClass) New() PropertyAnimation {
	rv := objc.Send[PropertyAnimation](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertyAnimation) Init() PropertyAnimation {
	rv := objc.Send[PropertyAnimation](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertyAnimation) Autorelease() PropertyAnimation {
	rv := objc.Send[PropertyAnimation](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertyAnimation creates a new PropertyAnimation instance.
func NewPropertyAnimation() PropertyAnimation {
	return getPropertyAnimationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PropertyAnimation */
// An abstract subclass for creating animations that manipulate the value of layer properties.
//
// The property to animate is specified using a key path that is relative to the layer using the animation. You do not create instances of : to animate the properties of a Core Animation layer, create instance of the concrete subclasses or .


// An abstract subclass for creating animations that manipulate the value of layer properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation
type PropertyAnimation struct {
	Animation
}

// PropertyAnimationFrom constructs a [PropertyAnimation] from an unsafe.Pointer.
//
// An abstract subclass for creating animations that manipulate the value of layer properties.
func PropertyAnimationFrom(ptr unsafe.Pointer) PropertyAnimation {
	return PropertyAnimation{
		Animation: AnimationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PropertyAnimation */

// Creates and returns an instance for the specified key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/init(keyPath:)
func NewPropertyAnimationWithKeyPath(path objc.IObject /* cross-framework: NSString */) PropertyAnimation {
	rv := objc.Send[PropertyAnimation](objc.ID(getPropertyAnimationClass().class), objc.Sel("animationWithKeyPath:"), path)
	return rv
}/* debug [class_init_methods/constructor]: NewPropertyAnimationWithKeyPath */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PropertyAnimation */

// Creates and returns an instance for the specified key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/init(keyPath:)
func (pc _PropertyAnimationClass) AnimationWithKeyPath(path objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("animationWithKeyPath:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AnimationWithKeyPath) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PropertyAnimation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PropertyAnimation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PropertyAnimation */

// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isAdditive
func (p_ PropertyAnimation) Additive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("additive"))
	return rv
}/* debug [instance_properties/getter]: additive */


// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isAdditive
func (p_ PropertyAnimation) SetAdditive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdditive:"), value)
}/* debug [instance_properties/setter]: additive */


// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isCumulative
func (p_ PropertyAnimation) Cumulative() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("cumulative"))
	return rv
}/* debug [instance_properties/getter]: cumulative */


// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isCumulative
func (p_ PropertyAnimation) SetCumulative(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCumulative:"), value)
}/* debug [instance_properties/setter]: cumulative */


// Specifies the key path the receiver animates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/keyPath
func (p_ PropertyAnimation) KeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("keyPath"))
	return rv
}/* debug [instance_properties/getter]: keyPath */


// Specifies the key path the receiver animates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/keyPath
func (p_ PropertyAnimation) SetKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setKeyPath:"), value)
}/* debug [instance_properties/setter]: keyPath */


// An optional value function that is applied to interpolated values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/valueFunction
func (p_ PropertyAnimation) ValueFunction() IValueFunction {
	rv := objc.Send[ValueFunction](p_.ID, objc.Sel("valueFunction"))
	return rv
}/* debug [instance_properties/getter]: valueFunction */


// An optional value function that is applied to interpolated values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/valueFunction
func (p_ PropertyAnimation) SetValueFunction(value IValueFunction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValueFunction:"), value)
}/* debug [instance_properties/setter]: valueFunction */


// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/isadditive
func (p_ PropertyAnimation) IsAdditive() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isAdditive"))
	return rv
}/* debug [instance_properties/getter]: isAdditive */


// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/isadditive
func (p_ PropertyAnimation) SetIsAdditive(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsAdditive:"), value)
}/* debug [instance_properties/setter]: isAdditive */


// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/iscumulative
func (p_ PropertyAnimation) IsCumulative() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCumulative"))
	return rv
}/* debug [instance_properties/getter]: isCumulative */


// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/capropertyanimation/iscumulative
func (p_ PropertyAnimation) SetIsCumulative(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCumulative:"), value)
}/* debug [instance_properties/setter]: isCumulative */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAPropertyAnimation */



