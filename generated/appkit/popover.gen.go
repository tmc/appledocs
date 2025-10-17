// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Popover] class.
var popoverClass = _PopoverClass{objc.GetClass("NSPopover")}

type _PopoverClass struct {
	class objc.Class
}

// An interface definition for the [Popover] class.
type IPopover interface {
	IResponder
}

// A means to display additional content related to existing content on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover

type Popover struct {
	Responder
}

// PopoverFrom constructs a [Popover] from an unsafe.Pointer.
//
// A means to display additional content related to existing content on the screen.
func PopoverFrom(ptr unsafe.Pointer) Popover {
	return Popover{
		Responder: ResponderFrom(ptr),
	}
}



