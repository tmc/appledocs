
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AnimationContext] class.
var AnimationContextClass _AnimationContextClass

func init() {
	AnimationContextClass = _AnimationContextClass{objc.GetClass("NSAnimationContext")}
}

type _AnimationContextClass struct {
	objc.Class
}

// An interface definition for the [AnimationContext] class.
type IAnimationContext interface {
	ID() objc.ID
}

type AnimationContext struct {
	id objc.ID
}

func AnimationContextFrom(ptr unsafe.Pointer) AnimationContext {
	return AnimationContext{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ AnimationContext) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _AnimationContextClass) Alloc() AnimationContext {
	rv := objc.Send[AnimationContext](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _AnimationContextClass) New() AnimationContext {
	rv := objc.Send[AnimationContext](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewAnimationContext creates and returns a new initialized instance.
func NewAnimationContext() AnimationContext {
	return AnimationContextClass.New()
}

// Init initializes the instance.
func (a_ AnimationContext) Init() AnimationContext {
	rv := objc.Send[AnimationContext](a_.ID(), selInit)
	return rv
}
