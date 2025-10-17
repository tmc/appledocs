// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Switch_] class.
var switch_Class = _Switch_Class{objc.GetClass("NSSwitch")}

type _Switch_Class struct {
	class objc.Class
}

// An interface definition for the [Switch_] class.
type ISwitch_ interface {
	IControl
}

// A control that offers a binary choice. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSwitch

type Switch_ struct {
	Control
}

// Switch_From constructs a [Switch_] from an unsafe.Pointer.
//
// A control that offers a binary choice.
func Switch_From(ptr unsafe.Pointer) Switch_ {
	return Switch_{
		Control: ControlFrom(ptr),
	}
}



