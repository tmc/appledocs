// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SearchFieldCell] class.
var SearchFieldCellClass objc.Class

func init() {
	SearchFieldCellClass = objc.GetClass("NSSearchFieldCell")
}

type SearchFieldCell struct {
	objc.ID
}

func SearchFieldCellFrom(ptr unsafe.Pointer) SearchFieldCell {
	return SearchFieldCell{
		ID: objc.ID(ptr),
	}
}




