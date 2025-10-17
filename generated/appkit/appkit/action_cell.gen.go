// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ActionCell] class.
var ActionCellClass objc.Class

func init() {
	ActionCellClass = objc.GetClass("NSActionCell")
}

type ActionCell struct {
	objc.ID
}

func ActionCellFrom(ptr unsafe.Pointer) ActionCell {
	return ActionCell{
		ID: objc.ID(ptr),
	}
}




