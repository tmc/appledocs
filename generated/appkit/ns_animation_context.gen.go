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

// An interface definition for the [AnimationContext] class.
type IAnimationContext interface {
	objectivec.IObject
}

// An animation context, which contains information about environment and state.
//
// is analogous to and is similar in overall concept to . Each thread maintains its own stack of nestable instances, with each new instance initialized as a copy of the instance below (so, inheriting its current properties). Multiple instances can be nested, allowing a given block of code to initiate animations using its own specified duration without affecting animations initiated by surrounding code.
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

// Ends the current animation grouping.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/endGrouping()
func (ac _AnimationContextClass) EndGrouping() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("endGrouping"))
}

// Allows you to specify a completion block body after the set of animation actions whose completion will trigger the completion block.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/runAnimationGroup(_:completionHandler:)
func (ac _AnimationContextClass) RunAnimationGroupCompletionHandler(changes unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("runAnimationGroup:completionHandler:"), changes, completionHandler)
}
