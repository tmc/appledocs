// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ComboButton] class.
var (
	ComboButtonClass     _ComboButtonClass
	ComboButtonClassOnce sync.Once
)

func getComboButtonClass() _ComboButtonClass {
	ComboButtonClassOnce.Do(func() {
		ComboButtonClass = _ComboButtonClass{objc.GetClass("NSComboButton")}
	})
	return ComboButtonClass
}

type _ComboButtonClass struct {
	class objc.Class
}

// An interface definition for the [ComboButton] class.
type IComboButton interface {
	IControl
}

// A button with a pull-down menu and a default action.
//
// An object is a button that displays a title string, image, and an optional control for displaying a menu. Use this control in places where you want to offer a button with a default action and one or more alternative actions. Clicking the title or image executes the default action you provide, and clicking the menu control displays a menu for selecting a different action. If you configure the button to hide the menu control, a long-press gesture displays the menu. After you create a combo button programmatically or in Interface Builder, choose the button you want and add a title or image for your content. A combo button has a default action, which you specify at creation time. You can also change that action later using the inherited and properties. To specify one or more alternative actions, configure a menu with those actions and assign it to the button’s property. This control doesn’t use an object for its underlying implementation. It also doesn’t support the addition of a contextual menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboButton
type ComboButton struct {
	Control
}

// ComboButtonFrom constructs a [ComboButton] from an unsafe.Pointer.
//
// A button with a pull-down menu and a default action.
func ComboButtonFrom(ptr unsafe.Pointer) ComboButton {
	return ComboButton{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ComboButtonClass) Alloc() ComboButton {
	rv := objc.Send[ComboButton](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComboButtonClass) New() ComboButton {
	rv := objc.Send[ComboButton](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComboButton) Init() ComboButton {
	rv := objc.Send[ComboButton](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComboButton) Autorelease() ComboButton {
	rv := objc.Send[ComboButton](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComboButton creates a new ComboButton instance.
func NewComboButton() ComboButton {
	return getComboButtonClass().New()
}
