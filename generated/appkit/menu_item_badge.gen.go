
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MenuItemBadge] class.
var MenuItemBadgeClass _MenuItemBadgeClass

func init() {
	MenuItemBadgeClass = _MenuItemBadgeClass{objc.GetClass("NSMenuItemBadge")}
}

type _MenuItemBadgeClass struct {
	objc.Class
}

// An interface definition for the [MenuItemBadge] class.
type IMenuItemBadge interface {
	ID() objc.ID
}

type MenuItemBadge struct {
	id objc.ID
}

func MenuItemBadgeFrom(ptr unsafe.Pointer) MenuItemBadge {
	return MenuItemBadge{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ MenuItemBadge) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MenuItemBadgeClass) Alloc() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MenuItemBadgeClass) New() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMenuItemBadge creates and returns a new initialized instance.
func NewMenuItemBadge() MenuItemBadge {
	return MenuItemBadgeClass.New()
}

// Init initializes the instance.
func (m_ MenuItemBadge) Init() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](m_.ID(), selInit)
	return rv
}
