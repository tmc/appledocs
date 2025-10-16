
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Animation] class.
var AnimationClass _AnimationClass

func init() {
	AnimationClass = _AnimationClass{objc.GetClass("NSAnimation")}
}

type _AnimationClass struct {
	objc.Class
}

// An interface definition for the [Animation] class.
type IAnimation interface {
	ID() objc.ID
}

type Animation struct {
	id objc.ID
}

func AnimationFrom(ptr unsafe.Pointer) Animation {
	return Animation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ Animation) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _AnimationClass) Alloc() Animation {
	rv := objc.Send[Animation](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _AnimationClass) New() Animation {
	rv := objc.Send[Animation](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewAnimation creates and returns a new initialized instance.
func NewAnimation() Animation {
	return AnimationClass.New()
}

// Init initializes the instance.
func (a_ Animation) Init() Animation {
	rv := objc.Send[Animation](a_.ID(), selInit)
	return rv
}
