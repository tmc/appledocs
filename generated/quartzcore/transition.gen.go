// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Transition] class.
var transitionClass = _TransitionClass{objc.GetClass("CATransition")}

type _TransitionClass struct {
	class objc.Class
}

// An interface definition for the [Transition] class.
type ITransition interface {
	IAnimation
}

// An object that provides an animated transition between a layer’s states. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (tc _TransitionClass) Alloc() Transition {
	rv := objc.Send[Transition](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return transitionClass.New()
}




