// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextLayoutManager] class.
var TextLayoutManagerClass objc.Class

func init() {
	TextLayoutManagerClass = objc.GetClass("NSTextLayoutManager")
}

type TextLayoutManager struct {
	objc.ID
}

func TextLayoutManagerFrom(ptr unsafe.Pointer) TextLayoutManager {
	return TextLayoutManager{
		ID: objc.ID(ptr),
	}
}




