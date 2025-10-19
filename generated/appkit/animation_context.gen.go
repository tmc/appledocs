// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AnimationContext] class.
var (
	animationContextClass     _AnimationContextClass
	animationContextClassOnce sync.Once
)

func getAnimationContextClass() _AnimationContextClass {
	animationContextClassOnce.Do(func() {
		animationContextClass = _AnimationContextClass{objc.GetClass("NSAnimationContext")}
	})
	return animationContextClass
}

type _AnimationContextClass struct {
	class objc.Class
}

// An interface definition for the [AnimationContext] class.
type IAnimationContext interface {
	objectivec.IObject
}

// An animation context, which contains information about environment and state.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AnimationContextClass) Alloc() AnimationContext {
	rv := objc.Send[AnimationContext](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




