// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAnimationContext */


/* debug [class_header]: Header for NSAnimationContext */
// The class instance for the [AnimationContext] class.
var (
	AnimationContextClass     _AnimationContextClass
	AnimationContextClassOnce sync.Once
)

func getAnimationContextClass() _AnimationContextClass {
	AnimationContextClassOnce.Do(func() {
		AnimationContextClass = _AnimationContextClass{objc.GetClass("NSAnimationContext")}
	})
	return AnimationContextClass
}

type _AnimationContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AnimationContext */
// An interface definition for the [AnimationContext] class.
type IAnimationContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AnimationContext */
	// properties:
	AllowsImplicitAnimation() bool
	SetAllowsImplicitAnimation(value bool)
	CompletionHandler() func()
	SetCompletionHandler(value func())
	Duration() float64
	SetDuration(value float64)
	TimingFunction() quartzcore.MediaTimingFunction
	SetTimingFunction(value quartzcore.MediaTimingFunction)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AnimationContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AnimationContext */
// Alloc allocates a new instance without initialization.
func (ac _AnimationContextClass) Alloc() AnimationContext {
	rv := objc.Send[AnimationContext](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AnimationContextClass) New() AnimationContext {
	rv := objc.Send[AnimationContext](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AnimationContext) Init() AnimationContext {
	rv := objc.Send[AnimationContext](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AnimationContext) Autorelease() AnimationContext {
	rv := objc.Send[AnimationContext](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAnimationContext creates a new AnimationContext instance.
func NewAnimationContext() AnimationContext {
	return getAnimationContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AnimationContext */
// An animation context, which contains information about environment and state.
//
// is analogous to and is similar in overall concept to . Each thread maintains its own stack of nestable instances, with each new instance initialized as a copy of the instance below (so, inheriting its current properties). Multiple instances can be nested, allowing a given block of code to initiate animations using its own specified duration without affecting animations initiated by surrounding code.


// An animation context, which contains information about environment and state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext
type AnimationContext struct {
	objectivec.Object
}

// AnimationContextFrom constructs a [AnimationContext] from an unsafe.Pointer.
//
// An animation context, which contains information about environment and state.
func AnimationContextFrom(ptr unsafe.Pointer) AnimationContext {
	return AnimationContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AnimationContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AnimationContext */

// Creates a new animation grouping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/beginGrouping()
func (ac _AnimationContextClass) BeginGrouping() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("beginGrouping"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BeginGrouping) */


// Ends the current animation grouping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/endGrouping()
func (ac _AnimationContextClass) EndGrouping() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("endGrouping"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EndGrouping) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/runAnimationGroup(_:)
func (ac _AnimationContextClass) RunAnimationGroup(changes unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("runAnimationGroup:"), changes)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RunAnimationGroup) */


// Allows you to specify a completion block body after the set of animation actions whose completion will trigger the completion block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/runAnimationGroup(_:completionHandler:)
func (ac _AnimationContextClass) RunAnimationGroupCompletionHandler(changes unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("runAnimationGroup:completionHandler:"), changes, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RunAnimationGroupCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AnimationContext */

// Returns the current animation context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/current
func (ac _AnimationContextClass) CurrentContext() AnimationContext {
	rv := objc.Send[AnimationContext](objc.ID(ac.class), objc.Sel("currentContext"))
	return rv
}/* debug [class_properties_class/property]: currentContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AnimationContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AnimationContext */

// Determine if animations are enabled or not for animations that occur as a result of another property change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/allowsImplicitAnimation
func (a_ AnimationContext) AllowsImplicitAnimation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsImplicitAnimation"))
	return rv
}/* debug [instance_properties/getter]: allowsImplicitAnimation */


// Determine if animations are enabled or not for animations that occur as a result of another property change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/allowsImplicitAnimation
func (a_ AnimationContext) SetAllowsImplicitAnimation(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsImplicitAnimation:"), value)
}/* debug [instance_properties/setter]: allowsImplicitAnimation */


// A completion Block that is called when the animations in the grouping are completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/completionHandler
func (a_ AnimationContext) CompletionHandler() func() {
	rv := objc.Send[func()](a_.ID, objc.Sel("completionHandler"))
	return rv
}/* debug [instance_properties/getter]: completionHandler */


// A completion Block that is called when the animations in the grouping are completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/completionHandler
func (a_ AnimationContext) SetCompletionHandler(value func()) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCompletionHandler:"), value)
}/* debug [instance_properties/setter]: completionHandler */


// Returns the current animation context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/current
func (a_ AnimationContext) CurrentContext() IAnimationContext {
	rv := objc.Send[AnimationContext](a_.ID, objc.Sel("currentContext"))
	return rv
}/* debug [instance_properties/getter]: currentContext */


// The duration used by animations created as a result of setting new values for an animatable property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/duration
func (a_ AnimationContext) Duration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The duration used by animations created as a result of setting new values for an animatable property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/duration
func (a_ AnimationContext) SetDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// The timing function used for all animations within this animation proxy group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/timingFunction
func (a_ AnimationContext) TimingFunction() quartzcore.MediaTimingFunction {
	rv := objc.Send[quartzcore.MediaTimingFunction](a_.ID, objc.Sel("timingFunction"))
	return rv
}/* debug [instance_properties/getter]: timingFunction */


// The timing function used for all animations within this animation proxy group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/timingFunction
func (a_ AnimationContext) SetTimingFunction(value quartzcore.MediaTimingFunction) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimingFunction:"), value)
}/* debug [instance_properties/setter]: timingFunction */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAnimationContext */



