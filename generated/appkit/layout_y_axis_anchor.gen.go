
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutYAxisAnchor] class.
var LayoutYAxisAnchorClass _LayoutYAxisAnchorClass

func init() {
	LayoutYAxisAnchorClass = _LayoutYAxisAnchorClass{objc.GetClass("NSLayoutYAxisAnchor")}
}

type _LayoutYAxisAnchorClass struct {
	objc.Class
}

// An interface definition for the [LayoutYAxisAnchor] class.
type ILayoutYAxisAnchor interface {
	ID() objc.ID
}

type LayoutYAxisAnchor struct {
	id objc.ID
}

func LayoutYAxisAnchorFrom(ptr unsafe.Pointer) LayoutYAxisAnchor {
	return LayoutYAxisAnchor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ LayoutYAxisAnchor) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutYAxisAnchorClass) Alloc() LayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _LayoutYAxisAnchorClass) New() LayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewLayoutYAxisAnchor creates and returns a new initialized instance.
func NewLayoutYAxisAnchor() LayoutYAxisAnchor {
	return LayoutYAxisAnchorClass.New()
}

// Init initializes the instance.
func (l_ LayoutYAxisAnchor) Init() LayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](l_.ID(), selInit)
	return rv
}
