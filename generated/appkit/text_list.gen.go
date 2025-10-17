// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextList] class.
var TextListClass objc.Class

func init() {
	TextListClass = objc.GetClass("NSTextList")
}

type TextList struct {
	objc.ID
}

func TextListFrom(ptr unsafe.Pointer) TextList {
	return TextList{
		ID: objc.ID(ptr),
	}
}



