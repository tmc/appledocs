// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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


// Indicates the end point of the receiver as a fraction of the entire transition.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/endProgress
func (t_ Transition) EndProgress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("endProgress"))
	return rv
}


// SetEndProgress sets the value of the endProgress property.
// Indicates the end point of the receiver as a fraction of the entire transition.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/endProgress
func (t_ Transition) SetEndProgress(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEndProgress:"), value)
}
// An optional Core Image filter object that provides the transition.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/filter
func (t_ Transition) Filter() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("filter"))
	return rv
}


// SetFilter sets the value of the filter property.
// An optional Core Image filter object that provides the transition.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/filter
func (t_ Transition) SetFilter(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFilter:"), value)
}
// Indicates the start point of the receiver as a fraction of the entire transition.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/startProgress
func (t_ Transition) StartProgress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("startProgress"))
	return rv
}


// SetStartProgress sets the value of the startProgress property.
// Indicates the start point of the receiver as a fraction of the entire transition.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/startProgress
func (t_ Transition) SetStartProgress(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStartProgress:"), value)
}
// Specifies an optional subtype that indicates the direction for the predefined motion-based transitions.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/subtype
func (t_ Transition) Subtype() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("subtype"))
	return rv
}


// SetSubtype sets the value of the subtype property.
// Specifies an optional subtype that indicates the direction for the predefined motion-based transitions.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/subtype
func (t_ Transition) SetSubtype(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSubtype:"), value)
}
// Specifies the predefined transition type.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/type
func (t_ Transition) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// Specifies the predefined transition type.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/type
func (t_ Transition) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setType:"), value)
}


