// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PopUpButtonCell] class.
var (
	popUpButtonCellClass     _PopUpButtonCellClass
	popUpButtonCellClassOnce sync.Once
)

func getPopUpButtonCellClass() _PopUpButtonCellClass {
	popUpButtonCellClassOnce.Do(func() {
		popUpButtonCellClass = _PopUpButtonCellClass{objc.GetClass("NSPopUpButtonCell")}
	})
	return popUpButtonCellClass
}

type _PopUpButtonCellClass struct {
	class objc.Class
}

// An interface definition for the [PopUpButtonCell] class.
type IPopUpButtonCell interface {
	IMenuItemCell
}

// The class defines the visual appearance of pop-up buttons that display pop-up or pull-down menus. Pop-up menus present the user with a set of choices, much the way radio buttons do, but using much less space. Pull-down menus also provide a set of choices but present the information in a slightly different way, usually to provide a set of commands from which the user can choose. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell

type PopUpButtonCell struct {
	MenuItemCell
}

// PopUpButtonCellFrom constructs a [PopUpButtonCell] from an unsafe.Pointer.
//
// The class defines the visual appearance of pop-up buttons that display pop-up or pull-down menus. Pop-up menus present the user with a set of choices, much the way radio buttons do, but using much less space. Pull-down menus also provide a set of choices but present the information in a slightly different way, usually to provide a set of commands from which the user can choose.
func PopUpButtonCellFrom(ptr unsafe.Pointer) PopUpButtonCell {
	return PopUpButtonCell{
		MenuItemCell: MenuItemCellFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PopUpButtonCellClass) Alloc() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PopUpButtonCellClass) New() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PopUpButtonCell) Init() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PopUpButtonCell) Autorelease() PopUpButtonCell {
	rv := objc.Send[PopUpButtonCell](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPopUpButtonCell creates a new PopUpButtonCell instance.
func NewPopUpButtonCell() PopUpButtonCell {
	return getPopUpButtonCellClass().New()
}




