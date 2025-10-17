// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontDescriptor] class.
var fontDescriptorClass = _FontDescriptorClass{objc.GetClass("NSFontDescriptor")}

type _FontDescriptorClass struct {
	class objc.Class
}

// A dictionary of attributes that describe a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor

type FontDescriptor struct {
	objectivec.Object
}

// FontDescriptorFrom constructs a [FontDescriptor] from an unsafe.Pointer.
//
// A dictionary of attributes that describe a font.
func FontDescriptorFrom(ptr unsafe.Pointer) FontDescriptor {
	return FontDescriptor{objectivec.Object{objc.ID(ptr)}}
}



