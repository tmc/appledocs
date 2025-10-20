// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Transition] class.
var (
	transitionClass     _TransitionClass
	transitionClassOnce sync.Once
)

func getTransitionClass() _TransitionClass {
	transitionClassOnce.Do(func() {
		transitionClass = _TransitionClass{objc.GetClass("CATransition")}
	})
	return transitionClass
}

type _TransitionClass struct {
	class objc.Class
}

// An interface definition for the [Transition] class.
type ITransition interface {
	IAnimation
}

// An object that provides an animated transition between a layer’s states.
//
// You can transition between a layer’s states by creating and adding a object to it. The default transition is a cross fade, but you can specify different effects from a set of predefined transitions. The following code shows how you can transition between the two states of a named . When the layer is first created, its is set to red and its property is set to . When the function is called, a new object is created and added to , and the state of the layer is changed so that its background color is blue and its rendered text reads . The end result is that the push transition animates the red state from left to right with the blue state entering the scene from the left.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




