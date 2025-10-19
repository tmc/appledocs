// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Menu] class.
var (
	menuClass     _MenuClass
	menuClassOnce sync.Once
)

func getMenuClass() _MenuClass {
	menuClassOnce.Do(func() {
		menuClass = _MenuClass{objc.GetClass("NSMenu")}
	})
	return menuClass
}

type _MenuClass struct {
	class objc.Class
}

// An interface definition for the [Menu] class.
type IMenu interface {
	objectivec.IObject
}

// An object that manages an app’s menus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu

type Menu struct {
	objectivec.Object
}

// MenuFrom constructs a [Menu] from an unsafe.Pointer.
//
// An object that manages an app’s menus.
func MenuFrom(ptr unsafe.Pointer) Menu {
	return Menu{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MenuClass) Alloc() Menu {
	rv := objc.Send[Menu](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MenuClass) New() Menu {
	rv := objc.Send[Menu](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Menu) Init() Menu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Menu) Autorelease() Menu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenu creates a new Menu instance.
func NewMenu() Menu {
	return getMenuClass().New()
}


// Displays a contextual menu over a view for an event. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (mc _MenuClass) PopUpContextMenuWithEventForView(menu unsafe.Pointer, event unsafe.Pointer, view unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:"), menu, event, view)
}


