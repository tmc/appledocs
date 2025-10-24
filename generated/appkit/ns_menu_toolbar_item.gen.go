// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMenuToolbarItem */


/* debug [class_header]: Header for NSMenuToolbarItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MenuToolbarItem */
// An interface definition for the [MenuToolbarItem] class.
type IMenuToolbarItem interface {
	IToolbarItem
	
/* debug [class_interface_properties]: Properties for MenuToolbarItem */
	// properties:
	Menu() IMenu
	SetMenu(value IMenu)
	ShowsIndicator() bool
	SetShowsIndicator(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MenuToolbarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MenuToolbarItem */
// Alloc allocates a new instance without initialization.
func (mc _MenuToolbarItemClass) Alloc() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MenuToolbarItem */
// A control that presents a menu in a window’s toolbar.
//
// If you set an action on an control item, the user invokes the action when clicking on the item through pressing and holding to display the menu. If you set an action on the item and to , the system displays the indicator as a separate segment so the user can invoke the menu with a click on that segment. If you don’t set an action on the , a simple click invokes the menu, and the indicator is purely decorative.


// A control that presents a menu in a window’s toolbar.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MenuToolbarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MenuToolbarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MenuToolbarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MenuToolbarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MenuToolbarItem */

// The menu presented from the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/menu
func (m_ MenuToolbarItem) Menu() IMenu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("menu"))
	return rv
}/* debug [instance_properties/getter]: menu */


// The menu presented from the toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/menu
func (m_ MenuToolbarItem) SetMenu(value IMenu) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMenu:"), value)
}/* debug [instance_properties/setter]: menu */


// A Boolean value that determines whether the toolbar item displays an indicator of additional functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/showsIndicator
func (m_ MenuToolbarItem) ShowsIndicator() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsIndicator"))
	return rv
}/* debug [instance_properties/getter]: showsIndicator */


// A Boolean value that determines whether the toolbar item displays an indicator of additional functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/showsIndicator
func (m_ MenuToolbarItem) SetShowsIndicator(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsIndicator:"), value)
}/* debug [instance_properties/setter]: showsIndicator */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMenuToolbarItem */


