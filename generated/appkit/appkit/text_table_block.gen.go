// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextTableBlock] class.
var TextTableBlockClass objc.Class

func init() {
	TextTableBlockClass = objc.GetClass("NSTextTableBlock")
}

type TextTableBlock struct {
	objc.ID
}

func TextTableBlockFrom(ptr unsafe.Pointer) TextTableBlock {
	return TextTableBlock{
		ID: objc.ID(ptr),
	}
}




