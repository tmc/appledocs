// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextRange] class.
var TextRangeClass objc.Class

func init() {
	TextRangeClass = objc.GetClass("NSTextRange")
}

type TextRange struct {
	objc.ID
}

func TextRangeFrom(ptr unsafe.Pointer) TextRange {
	return TextRange{
		ID: objc.ID(ptr),
	}
}




