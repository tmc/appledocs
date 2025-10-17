// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextBlock] class.
var TextBlockClass objc.Class

func init() {
	TextBlockClass = objc.GetClass("NSTextBlock")
}

type TextBlock struct {
	objc.ID
}

func TextBlockFrom(ptr unsafe.Pointer) TextBlock {
	return TextBlock{
		ID: objc.ID(ptr),
	}
}




