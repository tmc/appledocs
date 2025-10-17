
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DraggingItem] class.
var DraggingItemClass _DraggingItemClass

func init() {
	DraggingItemClass = _DraggingItemClass{objc.GetClass("NSDraggingItem")}
}

type _DraggingItemClass struct {
	objc.Class
}

// An interface definition for the [DraggingItem] class.
type IDraggingItem interface {
	ID() objc.ID
}

type DraggingItem struct {
	id objc.ID
}

func DraggingItemFrom(ptr unsafe.Pointer) DraggingItem {
	return DraggingItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DraggingItem) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DraggingItemClass) Alloc() DraggingItem {
	rv := objc.Send[DraggingItem](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DraggingItemClass) New() DraggingItem {
	rv := objc.Send[DraggingItem](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDraggingItem creates and returns a new initialized instance.
func NewDraggingItem() DraggingItem {
	return DraggingItemClass.New()
}

// Init initializes the instance.
func (d_ DraggingItem) Init() DraggingItem {
	rv := objc.Send[DraggingItem](d_.ID(), selInit)
	return rv
}
