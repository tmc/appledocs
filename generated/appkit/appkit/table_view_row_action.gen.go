// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableViewRowAction] class.
var TableViewRowActionClass objc.Class

func init() {
	TableViewRowActionClass = objc.GetClass("NSTableViewRowAction")
}

type TableViewRowAction struct {
	objc.ID
}

func TableViewRowActionFrom(ptr unsafe.Pointer) TableViewRowAction {
	return TableViewRowAction{
		ID: objc.ID(ptr),
	}
}




