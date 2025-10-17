// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextTab] class.
var TextTabClass objc.Class

func init() {
	TextTabClass = objc.GetClass("NSTextTab")
}

type TextTab struct {
	objc.ID
}

func TextTabFrom(ptr unsafe.Pointer) TextTab {
	return TextTab{
		ID: objc.ID(ptr),
	}
}



