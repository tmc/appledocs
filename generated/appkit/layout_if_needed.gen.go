
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [layoutIfNeeded] class.
var layoutIfNeededClass _layoutIfNeededClass

func init() {
	layoutIfNeededClass = _layoutIfNeededClass{objc.GetClass("layoutIfNeeded")}
}

type _layoutIfNeededClass struct {
	objc.Class
}

// An interface definition for the [layoutIfNeeded] class.
type IlayoutIfNeeded interface {
	ID() objc.ID
}

type layoutIfNeeded struct {
	id objc.ID
}

func layoutIfNeededFrom(ptr unsafe.Pointer) layoutIfNeeded {
	return layoutIfNeeded{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ layoutIfNeeded) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _layoutIfNeededClass) Alloc() layoutIfNeeded {
	rv := objc.Send[layoutIfNeeded](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _layoutIfNeededClass) New() layoutIfNeeded {
	rv := objc.Send[layoutIfNeeded](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlayoutIfNeeded creates and returns a new initialized instance.
func NewlayoutIfNeeded() layoutIfNeeded {
	return layoutIfNeededClass.New()
}

// Init initializes the instance.
func (l_ layoutIfNeeded) Init() layoutIfNeeded {
	rv := objc.Send[layoutIfNeeded](l_.ID(), selInit)
	return rv
}
