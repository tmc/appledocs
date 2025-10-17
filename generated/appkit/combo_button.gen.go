// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ComboButton] class.
var comboButtonClass = _ComboButtonClass{objc.GetClass("NSComboButton")}

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



