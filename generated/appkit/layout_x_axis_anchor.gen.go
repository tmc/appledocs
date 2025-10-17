
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutXAxisAnchor] class.
var LayoutXAxisAnchorClass _LayoutXAxisAnchorClass

func init() {
	LayoutXAxisAnchorClass = _LayoutXAxisAnchorClass{objc.GetClass("NSLayoutXAxisAnchor")}
}

type _LayoutXAxisAnchorClass struct {
	objc.Class
}

// An interface definition for the [LayoutXAxisAnchor] class.
type ILayoutXAxisAnchor interface {
	ID() objc.ID
}

type LayoutXAxisAnchor struct {
	id objc.ID
}

func LayoutXAxisAnchorFrom(ptr unsafe.Pointer) LayoutXAxisAnchor {
	return LayoutXAxisAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ LayoutXAxisAnchor) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutXAxisAnchorClass) Alloc() LayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _LayoutXAxisAnchorClass) New() LayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewLayoutXAxisAnchor creates and returns a new initialized instance.
func NewLayoutXAxisAnchor() LayoutXAxisAnchor {
	return LayoutXAxisAnchorClass.New()
}

// Init initializes the instance.
func (l_ LayoutXAxisAnchor) Init() LayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID(), selInit)
	return rv
}
