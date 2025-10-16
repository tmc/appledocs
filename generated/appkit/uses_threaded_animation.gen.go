
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [usesThreadedAnimation] class.
var usesThreadedAnimationClass _usesThreadedAnimationClass

func init() {
	usesThreadedAnimationClass = _usesThreadedAnimationClass{objc.GetClass("usesThreadedAnimation")}
}

type _usesThreadedAnimationClass struct {
	objc.Class
}

// An interface definition for the [usesThreadedAnimation] class.
type IusesThreadedAnimation interface {
	ID() objc.ID
}

type usesThreadedAnimation struct {
	id objc.ID
}

func usesThreadedAnimationFrom(ptr unsafe.Pointer) usesThreadedAnimation {
	return usesThreadedAnimation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ usesThreadedAnimation) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _usesThreadedAnimationClass) Alloc() usesThreadedAnimation {
	rv := objc.Send[usesThreadedAnimation](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _usesThreadedAnimationClass) New() usesThreadedAnimation {
	rv := objc.Send[usesThreadedAnimation](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewusesThreadedAnimation creates and returns a new initialized instance.
func NewusesThreadedAnimation() usesThreadedAnimation {
	return usesThreadedAnimationClass.New()
}

// Init initializes the instance.
func (u_ usesThreadedAnimation) Init() usesThreadedAnimation {
	rv := objc.Send[usesThreadedAnimation](u_.ID(), selInit)
	return rv
}
