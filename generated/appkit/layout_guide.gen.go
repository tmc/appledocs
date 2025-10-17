
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutGuide] class.
var LayoutGuideClass _LayoutGuideClass

func init() {
	LayoutGuideClass = _LayoutGuideClass{objc.GetClass("NSLayoutGuide")}
}

type _LayoutGuideClass struct {
	objc.Class
}

// An interface definition for the [LayoutGuide] class.
type ILayoutGuide interface {
	ID() objc.ID
}

type LayoutGuide struct {
	id objc.ID
}

func LayoutGuideFrom(ptr unsafe.Pointer) LayoutGuide {
	return LayoutGuide{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ LayoutGuide) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutGuideClass) Alloc() LayoutGuide {
	rv := objc.Send[LayoutGuide](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _LayoutGuideClass) New() LayoutGuide {
	rv := objc.Send[LayoutGuide](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewLayoutGuide creates and returns a new initialized instance.
func NewLayoutGuide() LayoutGuide {
	return LayoutGuideClass.New()
}

// Init initializes the instance.
func (l_ LayoutGuide) Init() LayoutGuide {
	rv := objc.Send[LayoutGuide](l_.ID(), selInit)
	return rv
}
