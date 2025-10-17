// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableHeaderCell] class.
var TableHeaderCellClass objc.Class

func init() {
	TableHeaderCellClass = objc.GetClass("NSTableHeaderCell")
}

type TableHeaderCell struct {
	objc.ID
}

func TableHeaderCellFrom(ptr unsafe.Pointer) TableHeaderCell {
	return TableHeaderCell{
		ID: objc.ID(ptr),
	}
}



