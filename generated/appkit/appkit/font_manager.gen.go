// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontManager] class.
var FontManagerClass objc.Class

func init() {
	FontManagerClass = objc.GetClass("NSFontManager")
}

type FontManager struct {
	objc.ID
}

func FontManagerFrom(ptr unsafe.Pointer) FontManager {
	return FontManager{
		ID: objc.ID(ptr),
	}
}



