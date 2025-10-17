
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutDimension] class.
var LayoutDimensionClass _LayoutDimensionClass

func init() {
	LayoutDimensionClass = _LayoutDimensionClass{objc.GetClass("NSLayoutDimension")}
}

type _LayoutDimensionClass struct {
	objc.Class
}

// An interface definition for the [LayoutDimension] class.
type ILayoutDimension interface {
	ID() objc.ID
}

type LayoutDimension struct {
	id objc.ID
}

func LayoutDimensionFrom(ptr unsafe.Pointer) LayoutDimension {
	return LayoutDimension{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ LayoutDimension) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutDimensionClass) Alloc() LayoutDimension {
	rv := objc.Send[LayoutDimension](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _LayoutDimensionClass) New() LayoutDimension {
	rv := objc.Send[LayoutDimension](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewLayoutDimension creates and returns a new initialized instance.
func NewLayoutDimension() LayoutDimension {
	return LayoutDimensionClass.New()
}

// Init initializes the instance.
func (l_ LayoutDimension) Init() LayoutDimension {
	rv := objc.Send[LayoutDimension](l_.ID(), selInit)
	return rv
}
