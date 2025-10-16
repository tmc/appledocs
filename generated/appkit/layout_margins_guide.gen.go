
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [layoutMarginsGuide] class.
var layoutMarginsGuideClass _layoutMarginsGuideClass

func init() {
	layoutMarginsGuideClass = _layoutMarginsGuideClass{objc.GetClass("layoutMarginsGuide")}
}

type _layoutMarginsGuideClass struct {
	objc.Class
}

// An interface definition for the [layoutMarginsGuide] class.
type IlayoutMarginsGuide interface {
	ID() objc.ID
}

type layoutMarginsGuide struct {
	id objc.ID
}

func layoutMarginsGuideFrom(ptr unsafe.Pointer) layoutMarginsGuide {
	return layoutMarginsGuide{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ layoutMarginsGuide) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _layoutMarginsGuideClass) Alloc() layoutMarginsGuide {
	rv := objc.Send[layoutMarginsGuide](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _layoutMarginsGuideClass) New() layoutMarginsGuide {
	rv := objc.Send[layoutMarginsGuide](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlayoutMarginsGuide creates and returns a new initialized instance.
func NewlayoutMarginsGuide() layoutMarginsGuide {
	return layoutMarginsGuideClass.New()
}

// Init initializes the instance.
func (l_ layoutMarginsGuide) Init() layoutMarginsGuide {
	rv := objc.Send[layoutMarginsGuide](l_.ID(), selInit)
	return rv
}
