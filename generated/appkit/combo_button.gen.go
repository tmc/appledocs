// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ComboButton] class.
var (
	comboButtonClass     _ComboButtonClass
	comboButtonClassOnce sync.Once
)

func getComboButtonClass() _ComboButtonClass {
	comboButtonClassOnce.Do(func() {
		comboButtonClass = _ComboButtonClass{objc.GetClass("NSComboButton")}
	})
	return comboButtonClass
}

type _ComboButtonClass struct {
	class objc.Class
}

// An interface definition for the [ComboButton] class.
type IComboButton interface {
	IControl
}

// A button with a pull-down menu and a default action. [Full Topic]
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




