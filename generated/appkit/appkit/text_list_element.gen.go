// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextListElement] class.
var TextListElementClass objc.Class

func init() {
	TextListElementClass = objc.GetClass("NSTextListElement")
}

type TextListElement struct {
	objc.ID
}

func TextListElementFrom(ptr unsafe.Pointer) TextListElement {
	return TextListElement{
		ID: objc.ID(ptr),
	}
}




