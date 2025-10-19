// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MenuItem] class.
var menuItemClass = _MenuItemClass{objc.GetClass("NSMenuItem")}

type _MenuItemClass struct {
	class objc.Class
}

// An interface definition for the [MenuItem] class.
type IMenuItem interface {
	objectivec.IObject
}

// A command item in an app menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem

type MenuItem struct {
	objectivec.Object
}

// MenuItemFrom constructs a [MenuItem] from an unsafe.Pointer.
//
// A command item in an app menu.
func MenuItemFrom(ptr unsafe.Pointer) MenuItem {
	return MenuItem{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MenuItemClass) Alloc() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MenuItemClass) New() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuItem) Init() MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuItem) Autorelease() MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuItem creates a new MenuItem instance.
func NewMenuItem() MenuItem {
	return menuItemClass.New()
}




