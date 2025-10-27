// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	EndProgress() float32
	SetEndProgress(value float32)
	Filter() objc.ID
	SetFilter(value objc.ID)
	StartProgress() float32
	SetStartProgress(value float32)
	Subtype() TransitionSubtype
	SetSubtype(value TransitionSubtype)
	Type() TransitionType
	SetType(value TransitionType)
	BackgroundColor() objectivec.IObject
	SetBackgroundColor(value objectivec.IObject)
	String() objectivec.IObject
	SetString(value objectivec.IObject)


	

	// methods:


}





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

























// Indicates the end point of the receiver as a fraction of the entire transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/endProgress
func (t_ Transition) EndProgress() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("endProgress"))
	return rv
}


// Indicates the end point of the receiver as a fraction of the entire transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/endProgress
func (t_ Transition) SetEndProgress(value float32) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEndProgress:"), value)
}


// An optional Core Image filter object that provides the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/filter
func (t_ Transition) Filter() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("filter"))
	return rv
}


// An optional Core Image filter object that provides the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/filter
func (t_ Transition) SetFilter(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFilter:"), value)
}


// Indicates the start point of the receiver as a fraction of the entire transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/startProgress
func (t_ Transition) StartProgress() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("startProgress"))
	return rv
}


// Indicates the start point of the receiver as a fraction of the entire transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/startProgress
func (t_ Transition) SetStartProgress(value float32) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStartProgress:"), value)
}


// Specifies an optional subtype that indicates the direction for the predefined motion-based transitions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/subtype
func (t_ Transition) Subtype() TransitionSubtype {
	rv := objc.Send[TransitionSubtype](t_.ID, objc.Sel("subtype"))
	return rv
}


// Specifies an optional subtype that indicates the direction for the predefined motion-based transitions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/subtype
func (t_ Transition) SetSubtype(value TransitionSubtype) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSubtype:"), value)
}


// Specifies the predefined transition type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/type
func (t_ Transition) Type() TransitionType {
	rv := objc.Send[TransitionType](t_.ID, objc.Sel("type"))
	return rv
}


// Specifies the predefined transition type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition/type
func (t_ Transition) SetType(value TransitionType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setType:"), value)
}


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (t_ Transition) BackgroundColor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (t_ Transition) SetBackgroundColor(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The text to be rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/string
func (t_ Transition) String() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("string"))
	return rv
}


// The text to be rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catextlayer/string
func (t_ Transition) SetString(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}








