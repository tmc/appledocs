// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MenuItemCell] class.
var (
	MenuItemCellClass     _MenuItemCellClass
	MenuItemCellClassOnce sync.Once
)

func getMenuItemCellClass() _MenuItemCellClass {
	MenuItemCellClassOnce.Do(func() {
		MenuItemCellClass = _MenuItemCellClass{objc.GetClass("NSMenuItemCell")}
	})
	return MenuItemCellClass
}

type _MenuItemCellClass struct {
	class objc.Class
}

// An interface definition for the [MenuItemCell] class.
type IMenuItemCell interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type MenuItemCell struct {
	objectivec.Object
}

// MenuItemCellFrom constructs a [MenuItemCell] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func MenuItemCellFrom(ptr unsafe.Pointer) MenuItemCell {
	return MenuItemCell{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MenuItemCellClass) Alloc() MenuItemCell {
	rv := objc.Send[MenuItemCell](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MenuItemCellClass) New() MenuItemCell {
	rv := objc.Send[MenuItemCell](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuItemCell) Init() MenuItemCell {
	rv := objc.Send[MenuItemCell](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuItemCell) Autorelease() MenuItemCell {
	rv := objc.Send[MenuItemCell](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuItemCell creates a new MenuItemCell instance.
func NewMenuItemCell() MenuItemCell {
	return getMenuItemCellClass().New()
}




