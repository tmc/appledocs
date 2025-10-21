// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Menu] class.
var (
	MenuClass     _MenuClass
	MenuClassOnce sync.Once
)

func getMenuClass() _MenuClass {
	MenuClassOnce.Do(func() {
		MenuClass = _MenuClass{objc.GetClass("NSMenu")}
	})
	return MenuClass
}

type _MenuClass struct {
	class objc.Class
}

// An interface definition for the [Menu] class.
type IMenu interface {
	objectivec.IObject
	HelpRequested(eventPtr unsafe.Pointer)
	PopUpMenuPositioningItemAtLocationInView(item unsafe.Pointer, location coregraphics.CGPoint, view unsafe.Pointer) bool
}

// An object that manages an app’s menus.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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

// Creates a palette style menu displaying user-selectable color tags.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/paletteMenuWithColors:titles:selectionHandler:
func (mc _MenuClass) PaletteMenuWithColorsTitlesSelectionHandler(colors unsafe.Pointer, itemTitles unsafe.Pointer, onSelectionChange unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("paletteMenuWithColors:titles:selectionHandler:"), colors, itemTitles, onSelectionChange)
	return rv
}

// Displays a contextual menu over a view for an event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (mc _MenuClass) PopUpContextMenuWithEventForView(menu unsafe.Pointer, event unsafe.Pointer, view unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:"), menu, event, view)
}

// Displays a contextual menu over a view for an event using a specified font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:with:)
func (mc _MenuClass) PopUpContextMenuWithEventForViewWithFont(menu unsafe.Pointer, event unsafe.Pointer, view unsafe.Pointer, font unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:withFont:"), menu, event, view, font)
}

// Overridden by subclasses to implement specialized context-sensitive help behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/helpRequested(with:)
func (m_ Menu) HelpRequested(eventPtr unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("helpRequested:"), eventPtr)
}

// Pops up the menu at the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUp(positioning:at:in:)
func (m_ Menu) PopUpMenuPositioningItemAtLocationInView(item unsafe.Pointer, location coregraphics.CGPoint, view unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("popUpMenuPositioningItem:atLocation:inView:"), item, location, view)
	return rv
}
