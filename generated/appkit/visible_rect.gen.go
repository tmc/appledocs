
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [visibleRect] class.
var visibleRectClass _visibleRectClass

func init() {
	visibleRectClass = _visibleRectClass{objc.GetClass("visibleRect")}
}

type _visibleRectClass struct {
	objc.Class
}

// An interface definition for the [visibleRect] class.
type IvisibleRect interface {
	ID() objc.ID
}

type visibleRect struct {
	id objc.ID
}

func visibleRectFrom(ptr unsafe.Pointer) visibleRect {
	return visibleRect{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ visibleRect) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _visibleRectClass) Alloc() visibleRect {
	rv := objc.Send[visibleRect](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _visibleRectClass) New() visibleRect {
	rv := objc.Send[visibleRect](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewvisibleRect creates and returns a new initialized instance.
func NewvisibleRect() visibleRect {
	return visibleRectClass.New()
}

// Init initializes the instance.
func (v_ visibleRect) Init() visibleRect {
	rv := objc.Send[visibleRect](v_.ID(), selInit)
	return rv
}
