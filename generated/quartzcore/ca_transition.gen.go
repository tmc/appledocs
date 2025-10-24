// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CATransition */


/* debug [class_header]: Header for CATransition */
// The class instance for the [Transition] class.
var (
	TransitionClass     _TransitionClass
	TransitionClassOnce sync.Once
)

func getTransitionClass() _TransitionClass {
	TransitionClassOnce.Do(func() {
		TransitionClass = _TransitionClass{objc.GetClass("CATransition")}
	})
	return TransitionClass
}

type _TransitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Transition */
// An interface definition for the [Transition] class.
type ITransition interface {
	IAnimation
	
/* debug [class_interface_properties]: Properties for Transition */
	// properties:
	EndProgress() float32
	SetEndProgress(value float32)
	Filter() objc.ID
	SetFilter(value objc.ID)
	StartProgress() float32
	SetStartProgress(value float32)
	Subtype() TransitionSubtype /* typedef */
	SetSubtype(value TransitionSubtype /* typedef */)
	Type() TransitionType /* typedef */
	SetType(value TransitionType /* typedef */)
	BackgroundColor() objectivec.IObject
	SetBackgroundColor(value objectivec.IObject)
	String() objectivec.IObject
	SetString(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Transition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Transition */
// Alloc allocates a new instance without initialization.
func (tc _TransitionClass) Alloc() Transition {
	rv := objc.Send[Transition](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TransitionClass) New() Transition {
	rv := objc.Send[Transition](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Transition) Init() Transition {
	rv := objc.Send[Transition](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Transition) Autorelease() Transition {
	rv := objc.Send[Transition](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTransition creates a new Transition instance.
func NewTransition() Transition {
	return getTransitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Transition */
// An object that provides an animated transition between a layer’s states.
//
// You can transition between a layer’s states by creating and adding a object to it. The default transition is a cross fade, but you can specify different effects from a set of predefined transitions. The following code shows how you can transition between the two states of a named . When the layer is first created, its is set to red and its property is set to . When the function is called, a new object is created and added to , and the state of the layer is changed so that its background color is blue and its rendered text reads . The end result is that the push transition animates the red state from left to right with the blue state entering the scene from the left.


// An object that provides an animated transition between a layer’s states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition
type Transition struct {
	Animation
}

// TransitionFrom constructs a [Transition] from an unsafe.Pointer.
//
// An object that provides an animated transition between a layer’s states.
func TransitionFrom(ptr unsafe.Pointer) Transition {
	return Transition{
		Animation: AnimationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Transition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Transition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Transition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Transition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Transition */

// Indicates the end point of the receiver as a fraction of the entire transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/endProgress
func (t_ Transition) EndProgress() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("endProgress"))
	return rv
}/* debug [instance_properties/getter]: endProgress */


// Indicates the end point of the receiver as a fraction of the entire transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/endProgress
func (t_ Transition) SetEndProgress(value float32) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEndProgress:"), value)
}/* debug [instance_properties/setter]: endProgress */


// An optional Core Image filter object that provides the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/filter
func (t_ Transition) Filter() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("filter"))
	return rv
}/* debug [instance_properties/getter]: filter */


// An optional Core Image filter object that provides the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/filter
func (t_ Transition) SetFilter(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFilter:"), value)
}/* debug [instance_properties/setter]: filter */


// Indicates the start point of the receiver as a fraction of the entire transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/startProgress
func (t_ Transition) StartProgress() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("startProgress"))
	return rv
}/* debug [instance_properties/getter]: startProgress */


// Indicates the start point of the receiver as a fraction of the entire transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/startProgress
func (t_ Transition) SetStartProgress(value float32) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStartProgress:"), value)
}/* debug [instance_properties/setter]: startProgress */


// Specifies an optional subtype that indicates the direction for the predefined motion-based transitions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/subtype
func (t_ Transition) Subtype() TransitionSubtype /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("subtype"))
	return rv
}/* debug [instance_properties/getter]: subtype */


// Specifies an optional subtype that indicates the direction for the predefined motion-based transitions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/subtype
func (t_ Transition) SetSubtype(value TransitionSubtype /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSubtype:"), value)
}/* debug [instance_properties/setter]: subtype */


// Specifies the predefined transition type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/type
func (t_ Transition) Type() TransitionType /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// Specifies the predefined transition type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/type
func (t_ Transition) SetType(value TransitionType /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (t_ Transition) BackgroundColor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (t_ Transition) SetBackgroundColor(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The text to be rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/string
func (t_ Transition) String() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */


// The text to be rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/string
func (t_ Transition) SetString(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}/* debug [instance_properties/setter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CATransition */



