// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Popover] class.
var PopoverClass objc.Class

func init() {
	PopoverClass = objc.GetClass("NSPopover")
}

type Popover struct {
	objc.ID
}

func PopoverFrom(ptr unsafe.Pointer) Popover {
	return Popover{
		ID: objc.ID(ptr),
	}
}



