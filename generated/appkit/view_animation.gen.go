// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [ViewAnimation] class.
var (
	viewAnimationClass     _ViewAnimationClass
	viewAnimationClassOnce sync.Once
)

func getViewAnimationClass() _ViewAnimationClass {
	viewAnimationClassOnce.Do(func() {
		viewAnimationClass = _ViewAnimationClass{objc.GetClass("NSViewAnimation")}
	})
	return viewAnimationClass
}

type _ViewAnimationClass struct {
	class objc.Class
}

// An interface definition for the [ViewAnimation] class.
type IViewAnimation interface {
	quartzcore.IAnimation
}

// An animation of an app’s views, limited to changes in frame location and size, and to fade-in and fade-out effects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewAnimation
type ViewAnimation struct {
	quartzcore.Animation
}

// ViewAnimationFrom constructs a [ViewAnimation] from an unsafe.Pointer.
//
// An animation of an app’s views, limited to changes in frame location and size, and to fade-in and fade-out effects.
func ViewAnimationFrom(ptr unsafe.Pointer) ViewAnimation {
	return ViewAnimation{
		Animation: quartzcore.AnimationFrom(ptr),
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




