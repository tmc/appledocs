
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [layoutSubtreeIfNeeded] class.
var layoutSubtreeIfNeededClass _layoutSubtreeIfNeededClass

func init() {
	layoutSubtreeIfNeededClass = _layoutSubtreeIfNeededClass{objc.GetClass("layoutSubtreeIfNeeded")}
}

type _layoutSubtreeIfNeededClass struct {
	objc.Class
}

// An interface definition for the [layoutSubtreeIfNeeded] class.
type IlayoutSubtreeIfNeeded interface {
	ID() objc.ID
}

type layoutSubtreeIfNeeded struct {
	id objc.ID
}

func layoutSubtreeIfNeededFrom(ptr unsafe.Pointer) layoutSubtreeIfNeeded {
	return layoutSubtreeIfNeeded{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ layoutSubtreeIfNeeded) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _layoutSubtreeIfNeededClass) Alloc() layoutSubtreeIfNeeded {
	rv := objc.Send[layoutSubtreeIfNeeded](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _layoutSubtreeIfNeededClass) New() layoutSubtreeIfNeeded {
	rv := objc.Send[layoutSubtreeIfNeeded](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlayoutSubtreeIfNeeded creates and returns a new initialized instance.
func NewlayoutSubtreeIfNeeded() layoutSubtreeIfNeeded {
	return layoutSubtreeIfNeededClass.New()
}

// Init initializes the instance.
func (l_ layoutSubtreeIfNeeded) Init() layoutSubtreeIfNeeded {
	rv := objc.Send[layoutSubtreeIfNeeded](l_.ID(), selInit)
	return rv
}
