
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ViewAnimation] class.
var ViewAnimationClass _ViewAnimationClass

func init() {
	ViewAnimationClass = _ViewAnimationClass{objc.GetClass("NSViewAnimation")}
}

type _ViewAnimationClass struct {
	objc.Class
}

// An interface definition for the [ViewAnimation] class.
type IViewAnimation interface {
	ID() objc.ID
}

type ViewAnimation struct {
	id objc.ID
}

func ViewAnimationFrom(ptr unsafe.Pointer) ViewAnimation {
	return ViewAnimation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ ViewAnimation) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _ViewAnimationClass) Alloc() ViewAnimation {
	rv := objc.Send[ViewAnimation](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _ViewAnimationClass) New() ViewAnimation {
	rv := objc.Send[ViewAnimation](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewViewAnimation creates and returns a new initialized instance.
func NewViewAnimation() ViewAnimation {
	return ViewAnimationClass.New()
}

// Init initializes the instance.
func (v_ ViewAnimation) Init() ViewAnimation {
	rv := objc.Send[ViewAnimation](v_.ID(), selInit)
	return rv
}
