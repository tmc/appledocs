// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontDescriptor] class.
var FontDescriptorClass objc.Class

func init() {
	FontDescriptorClass = objc.GetClass("NSFontDescriptor")
}

type FontDescriptor struct {
	objc.ID
}

func FontDescriptorFrom(ptr unsafe.Pointer) FontDescriptor {
	return FontDescriptor{
		ID: objc.ID(ptr),
	}
}




