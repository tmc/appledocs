// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextSelection] class.
var TextSelectionClass objc.Class

func init() {
	TextSelectionClass = objc.GetClass("NSTextSelection")
}

type TextSelection struct {
	objc.ID
}

func TextSelectionFrom(ptr unsafe.Pointer) TextSelection {
	return TextSelection{
		ID: objc.ID(ptr),
	}
}



