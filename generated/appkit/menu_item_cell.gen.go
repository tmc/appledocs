
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MenuItemCell] class.
var MenuItemCellClass _MenuItemCellClass

func init() {
	MenuItemCellClass = _MenuItemCellClass{objc.GetClass("NSMenuItemCell")}
}

type _MenuItemCellClass struct {
	objc.Class
}

// An interface definition for the [MenuItemCell] class.
type IMenuItemCell interface {
	ID() objc.ID
}

type MenuItemCell struct {
	id objc.ID
}

func MenuItemCellFrom(ptr unsafe.Pointer) MenuItemCell {
	return MenuItemCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ MenuItemCell) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MenuItemCellClass) Alloc() MenuItemCell {
	rv := objc.Send[MenuItemCell](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MenuItemCellClass) New() MenuItemCell {
	rv := objc.Send[MenuItemCell](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMenuItemCell creates and returns a new initialized instance.
func NewMenuItemCell() MenuItemCell {
	return MenuItemCellClass.New()
}

// Init initializes the instance.
func (m_ MenuItemCell) Init() MenuItemCell {
	rv := objc.Send[MenuItemCell](m_.ID(), selInit)
	return rv
}
