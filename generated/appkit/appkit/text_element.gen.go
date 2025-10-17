// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextElement] class.
var TextElementClass objc.Class

func init() {
	TextElementClass = objc.GetClass("NSTextElement")
}

type TextElement struct {
	objc.ID
}

func TextElementFrom(ptr unsafe.Pointer) TextElement {
	return TextElement{
		ID: objc.ID(ptr),
	}
}



