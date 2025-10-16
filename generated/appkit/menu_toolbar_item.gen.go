
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MenuToolbarItem] class.
var MenuToolbarItemClass _MenuToolbarItemClass

func init() {
	MenuToolbarItemClass = _MenuToolbarItemClass{objc.GetClass("NSMenuToolbarItem")}
}

type _MenuToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [MenuToolbarItem] class.
type IMenuToolbarItem interface {
	ID() objc.ID
}

type MenuToolbarItem struct {
	id objc.ID
}

func MenuToolbarItemFrom(ptr unsafe.Pointer) MenuToolbarItem {
	return MenuToolbarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ MenuToolbarItem) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MenuToolbarItemClass) Alloc() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MenuToolbarItemClass) New() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMenuToolbarItem creates and returns a new initialized instance.
func NewMenuToolbarItem() MenuToolbarItem {
	return MenuToolbarItemClass.New()
}

// Init initializes the instance.
func (m_ MenuToolbarItem) Init() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](m_.ID(), selInit)
	return rv
}
