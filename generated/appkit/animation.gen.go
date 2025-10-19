// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Animation] class.
var animationClass = _AnimationClass{objc.GetClass("NSAnimation")}

type _AnimationClass struct {
	class objc.Class
}

// An interface definition for the [Animation] class.
type IAnimation interface {
	objectivec.IObject
}

// An object that manages the timing and progress of animations in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation

type Animation struct {
	objectivec.Object
}

// AnimationFrom constructs a [Animation] from an unsafe.Pointer.
//
// An object that manages the timing and progress of animations in the user interface.
func AnimationFrom(ptr unsafe.Pointer) Animation {
	return Animation{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AnimationClass) Alloc() Animation {
	rv := objc.Send[Animation](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AnimationClass) New() Animation {
	rv := objc.Send[Animation](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Animation) Init() Animation {
	rv := objc.Send[Animation](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Animation) Autorelease() Animation {
	rv := objc.Send[Animation](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAnimation creates a new Animation instance.
func NewAnimation() Animation {
	return animationClass.New()
}




