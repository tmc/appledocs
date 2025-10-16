
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutAnchor] class.
var LayoutAnchorClass _LayoutAnchorClass

func init() {
	LayoutAnchorClass = _LayoutAnchorClass{objc.GetClass("NSLayoutAnchor")}
}

type _LayoutAnchorClass struct {
	objc.Class
}

// An interface definition for the [LayoutAnchor] class.
type ILayoutAnchor interface {
	ID() objc.ID
}

type LayoutAnchor struct {
	id objc.ID
}

func LayoutAnchorFrom(ptr unsafe.Pointer) LayoutAnchor {
	return LayoutAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ LayoutAnchor) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutAnchorClass) Alloc() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _LayoutAnchorClass) New() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewLayoutAnchor creates and returns a new initialized instance.
func NewLayoutAnchor() LayoutAnchor {
	return LayoutAnchorClass.New()
}

// Init initializes the instance.
func (l_ LayoutAnchor) Init() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](l_.ID(), selInit)
	return rv
}
