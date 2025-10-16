
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MenuItem] class.
var MenuItemClass _MenuItemClass

func init() {
	MenuItemClass = _MenuItemClass{objc.GetClass("NSMenuItem")}
}

type _MenuItemClass struct {
	objc.Class
}

// An interface definition for the [MenuItem] class.
type IMenuItem interface {
	ID() objc.ID
}

type MenuItem struct {
	id objc.ID
}

func MenuItemFrom(ptr unsafe.Pointer) MenuItem {
	return MenuItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ MenuItem) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MenuItemClass) Alloc() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MenuItemClass) New() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMenuItem creates and returns a new initialized instance.
func NewMenuItem() MenuItem {
	return MenuItemClass.New()
}

// Init initializes the instance.
func (m_ MenuItem) Init() MenuItem {
	rv := objc.Send[MenuItem](m_.ID(), selInit)
	return rv
}
