// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Font] class.
var fontClass = _FontClass{objc.GetClass("NSFont")}

type _FontClass struct {
	class objc.Class
}

// The representation of a font in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont

type Font struct {
	objectivec.Object
}

// FontFrom constructs a [Font] from an unsafe.Pointer.
//
// The representation of a font in an app.
func FontFrom(ptr unsafe.Pointer) Font {
	return Font{objectivec.Object{objc.ID(ptr)}}
}



