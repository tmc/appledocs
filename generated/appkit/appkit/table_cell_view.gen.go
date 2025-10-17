// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableCellView] class.
var TableCellViewClass objc.Class

func init() {
	TableCellViewClass = objc.GetClass("NSTableCellView")
}

type TableCellView struct {
	objc.ID
}

func TableCellViewFrom(ptr unsafe.Pointer) TableCellView {
	return TableCellView{
		ID: objc.ID(ptr),
	}
}



