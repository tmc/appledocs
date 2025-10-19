// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PopUpButton] class.
var (
	popUpButtonClass     _PopUpButtonClass
	popUpButtonClassOnce sync.Once
)

func getPopUpButtonClass() _PopUpButtonClass {
	popUpButtonClassOnce.Do(func() {
		popUpButtonClass = _PopUpButtonClass{objc.GetClass("NSPopUpButton")}
	})
	return popUpButtonClass
}

type _PopUpButtonClass struct {
	class objc.Class
}

// An interface definition for the [PopUpButton] class.
type IPopUpButton interface {
	IButton
}

// A control for selecting an item from a list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton
type PopUpButton struct {
	Button
}

// PopUpButtonFrom constructs a [PopUpButton] from an unsafe.Pointer.
//
// A control for selecting an item from a list.
func PopUpButtonFrom(ptr unsafe.Pointer) PopUpButton {
	return PopUpButton{
		Button: ButtonFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PopUpButtonClass) Alloc() PopUpButton {
	rv := objc.Send[PopUpButton](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PopUpButtonClass) New() PopUpButton {
	rv := objc.Send[PopUpButton](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PopUpButton) Init() PopUpButton {
	rv := objc.Send[PopUpButton](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PopUpButton) Autorelease() PopUpButton {
	rv := objc.Send[PopUpButton](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPopUpButton creates a new PopUpButton instance.
func NewPopUpButton() PopUpButton {
	return getPopUpButtonClass().New()
}




