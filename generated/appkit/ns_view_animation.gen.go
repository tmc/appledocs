// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
)

// The class instance for the [ViewAnimation] class.
var (
	ViewAnimationClass     _ViewAnimationClass
	ViewAnimationClassOnce sync.Once
)

func getViewAnimationClass() _ViewAnimationClass {
	ViewAnimationClassOnce.Do(func() {
		ViewAnimationClass = _ViewAnimationClass{objc.GetClass("NSViewAnimation")}
	})
	return ViewAnimationClass
}

type _ViewAnimationClass struct {
	class objc.Class
}

// An interface definition for the [ViewAnimation] class.
type IViewAnimation interface {
	IAnimation
}

// An animation of an app’s views, limited to changes in frame location and size, and to fade-in and fade-out effects.
//
// An object takes an array of dictionaries from which it determines the objects to animate and the effects to apply to them. Each dictionary must have a target object and, optionally, properties that specify beginning and ending frame and whether to fade in or fade out. (See for further information.) Animations with are, by default, in non-blocking mode over a duration of 0.5 seconds using the ease in-out animation curve. But you can configure the animation to have any duration, curve, frame rate, and blocking mode. You may also set progress marks, assign a delegate, and implement delegation methods in order to animate view and windows concurrent with the ones specified as targets in the view-animation dictionary. Invoking the method on a running object moves the animation to the end frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewAnimation
type ViewAnimation struct {
	Animation
}

// ViewAnimationFrom constructs a [ViewAnimation] from an unsafe.Pointer.
//
// An animation of an app’s views, limited to changes in frame location and size, and to fade-in and fade-out effects.
func ViewAnimationFrom(ptr unsafe.Pointer) ViewAnimation {
	return ViewAnimation{
		Animation: AnimationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _ViewAnimationClass) Alloc() ViewAnimation {
	rv := objc.Send[ViewAnimation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _ViewAnimationClass) New() ViewAnimation {
	rv := objc.Send[ViewAnimation](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ ViewAnimation) Init() ViewAnimation {
	rv := objc.Send[ViewAnimation](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ ViewAnimation) Autorelease() ViewAnimation {
	rv := objc.Send[ViewAnimation](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewViewAnimation creates a new ViewAnimation instance.
func NewViewAnimation() ViewAnimation {
	return getViewAnimationClass().New()
}


// The dictionaries defining the objects to animate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewanimation/viewanimations
func (v_ ViewAnimation) ViewAnimations() coreml.Key {
	rv := objc.Send[coreml.Key](v_.ID, objc.Sel("viewAnimations"))
	return rv
}


// SetViewAnimations sets the value of the viewAnimations property.
// The dictionaries defining the objects to animate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewanimation/viewanimations
func (v_ ViewAnimation) SetViewAnimations(value coreml.IKey) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setViewAnimations:"), value)
}



