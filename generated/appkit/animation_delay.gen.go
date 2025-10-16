
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [animationDelay] class.
var animationDelayClass _animationDelayClass

func init() {
	animationDelayClass = _animationDelayClass{objc.GetClass("animationDelay")}
}

type _animationDelayClass struct {
	objc.Class
}

// An interface definition for the [animationDelay] class.
type IanimationDelay interface {
	ID() objc.ID
}

type animationDelay struct {
	id objc.ID
}

func animationDelayFrom(ptr unsafe.Pointer) animationDelay {
	return animationDelay{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ animationDelay) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _animationDelayClass) Alloc() animationDelay {
	rv := objc.Send[animationDelay](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _animationDelayClass) New() animationDelay {
	rv := objc.Send[animationDelay](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewanimationDelay creates and returns a new initialized instance.
func NewanimationDelay() animationDelay {
	return animationDelayClass.New()
}

// Init initializes the instance.
func (a_ animationDelay) Init() animationDelay {
	rv := objc.Send[animationDelay](a_.ID(), selInit)
	return rv
}
