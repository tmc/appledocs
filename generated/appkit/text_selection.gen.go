// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextSelection] class.
var textSelectionClass = _TextSelectionClass{objc.GetClass("NSTextSelection")}

type _TextSelectionClass struct {
	class objc.Class
}

// A class that represents a single logical selection context that corresponds to an insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection

type TextSelection struct {
	objectivec.Object
}

// TextSelectionFrom constructs a [TextSelection] from an unsafe.Pointer.
//
// A class that represents a single logical selection context that corresponds to an insertion point.
func TextSelectionFrom(ptr unsafe.Pointer) TextSelection {
	return TextSelection{objectivec.Object{objc.ID(ptr)}}
}



