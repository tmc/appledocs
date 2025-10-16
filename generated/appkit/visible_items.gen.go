
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [visibleItems] class.
var visibleItemsClass _visibleItemsClass

func init() {
	visibleItemsClass = _visibleItemsClass{objc.GetClass("visibleItems")}
}

type _visibleItemsClass struct {
	objc.Class
}

// An interface definition for the [visibleItems] class.
type IvisibleItems interface {
	ID() objc.ID
}

type visibleItems struct {
	id objc.ID
}

func visibleItemsFrom(ptr unsafe.Pointer) visibleItems {
	return visibleItems{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ visibleItems) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _visibleItemsClass) Alloc() visibleItems {
	rv := objc.Send[visibleItems](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _visibleItemsClass) New() visibleItems {
	rv := objc.Send[visibleItems](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewvisibleItems creates and returns a new initialized instance.
func NewvisibleItems() visibleItems {
	return visibleItemsClass.New()
}

// Init initializes the instance.
func (v_ visibleItems) Init() visibleItems {
	rv := objc.Send[visibleItems](v_.ID(), selInit)
	return rv
}
