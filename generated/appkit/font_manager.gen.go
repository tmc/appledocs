// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontManager] class.
var fontManagerClass = _FontManagerClass{objc.GetClass("NSFontManager")}

type _FontManagerClass struct {
	class objc.Class
}

// An interface definition for the [FontManager] class.
type IFontManager interface {
	objectivec.IObject
}

// The center of activity for the font-conversion system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontManager

type FontManager struct {
	objectivec.Object
}

// FontManagerFrom constructs a [FontManager] from an unsafe.Pointer.
//
// The center of activity for the font-conversion system.
func FontManagerFrom(ptr unsafe.Pointer) FontManager {
	return FontManager{objectivec.Object{objc.ID(ptr)}}
}



