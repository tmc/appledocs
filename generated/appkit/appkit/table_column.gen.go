// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableColumn] class.
var TableColumnClass objc.Class

func init() {
	TableColumnClass = objc.GetClass("NSTableColumn")
}

type TableColumn struct {
	objc.ID
}

func TableColumnFrom(ptr unsafe.Pointer) TableColumn {
	return TableColumn{
		ID: objc.ID(ptr),
	}
}



