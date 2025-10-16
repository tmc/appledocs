
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [animationBehavior] class.
var animationBehaviorClass _animationBehaviorClass

func init() {
	animationBehaviorClass = _animationBehaviorClass{objc.GetClass("animationBehavior")}
}

type _animationBehaviorClass struct {
	objc.Class
}

// An interface definition for the [animationBehavior] class.
type IanimationBehavior interface {
	ID() objc.ID
}

type animationBehavior struct {
	id objc.ID
}

func animationBehaviorFrom(ptr unsafe.Pointer) animationBehavior {
	return animationBehavior{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ animationBehavior) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _animationBehaviorClass) Alloc() animationBehavior {
	rv := objc.Send[animationBehavior](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _animationBehaviorClass) New() animationBehavior {
	rv := objc.Send[animationBehavior](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewanimationBehavior creates and returns a new initialized instance.
func NewanimationBehavior() animationBehavior {
	return animationBehaviorClass.New()
}

// Init initializes the instance.
func (a_ animationBehavior) Init() animationBehavior {
	rv := objc.Send[animationBehavior](a_.ID(), selInit)
	return rv
}
