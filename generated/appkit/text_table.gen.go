// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextTable] class.
var TextTableClass objc.Class

func init() {
	TextTableClass = objc.GetClass("NSTextTable")
}

type TextTable struct {
	objc.ID
}

func TextTableFrom(ptr unsafe.Pointer) TextTable {
	return TextTable{
		ID: objc.ID(ptr),
	}
}



