// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MenuToolbarItem] class.
var (
	MenuToolbarItemClass     _MenuToolbarItemClass
	MenuToolbarItemClassOnce sync.Once
)

func getMenuToolbarItemClass() _MenuToolbarItemClass {
	MenuToolbarItemClassOnce.Do(func() {
		MenuToolbarItemClass = _MenuToolbarItemClass{objc.GetClass("NSMenuToolbarItem")}
	})
	return MenuToolbarItemClass
}

type _MenuToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [MenuToolbarItem] class.
type IMenuToolbarItem interface {
	IToolbarItem
}

// A control that presents a menu in a window’s toolbar.
//
// If you set an action on an control item, the user invokes the action when clicking on the item through pressing and holding to display the menu. If you set an action on the item and to , the system displays the indicator as a separate segment so the user can invoke the menu with a click on that segment. If you don’t set an action on the , a simple click invokes the menu, and the indicator is purely decorative.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem
type MenuToolbarItem struct {
	ToolbarItem
}

// MenuToolbarItemFrom constructs a [MenuToolbarItem] from an unsafe.Pointer.
//
// A control that presents a menu in a window’s toolbar.
func MenuToolbarItemFrom(ptr unsafe.Pointer) MenuToolbarItem {
	return MenuToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MenuToolbarItemClass) Alloc() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MenuToolbarItemClass) New() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuToolbarItem) Init() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuToolbarItem) Autorelease() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuToolbarItem creates a new MenuToolbarItem instance.
func NewMenuToolbarItem() MenuToolbarItem {
	return getMenuToolbarItemClass().New()
}


// The menu presented from the toolbar item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/menu
func (m_ MenuToolbarItem) Menu() NSMenu {
	rv := objc.Send[NSMenu](m_.ID, objc.Sel("menu"))
	return rv
}


// SetMenu sets the value of the menu property.
// The menu presented from the toolbar item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/menu
func (m_ MenuToolbarItem) SetMenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenu:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenutoolbaritem/itemmenu
func (m_ MenuToolbarItem) ItemMenu() Menu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("itemMenu"))
	return rv
}


// SetItemMenu sets the value of the itemMenu property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenutoolbaritem/itemmenu
func (m_ MenuToolbarItem) SetItemMenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setItemMenu:"), value)
}

// A Boolean value that determines whether the toolbar item displays an indicator of additional functionality.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenutoolbaritem/showsindicator
func (m_ MenuToolbarItem) ShowsIndicator() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsIndicator"))
	return rv
}


// SetShowsIndicator sets the value of the showsIndicator property.
// A Boolean value that determines whether the toolbar item displays an indicator of additional functionality.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenutoolbaritem/showsindicator
func (m_ MenuToolbarItem) SetShowsIndicator(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsIndicator:"), value)
}



