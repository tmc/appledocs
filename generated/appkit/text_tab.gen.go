// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextTab] class.
var textTabClass = _TextTabClass{objc.GetClass("NSTextTab")}

type _TextTabClass struct {
	class objc.Class
}

// A tab in a paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab

type TextTab struct {
	objectivec.Object
}

// TextTabFrom constructs a [TextTab] from an unsafe.Pointer.
//
// A tab in a paragraph.
func TextTabFrom(ptr unsafe.Pointer) TextTab {
	return TextTab{objectivec.Object{objc.ID(ptr)}}
}



