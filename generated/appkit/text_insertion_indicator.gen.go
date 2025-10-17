// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextInsertionIndicator] class.
var textInsertionIndicatorClass = _TextInsertionIndicatorClass{objc.GetClass("NSTextInsertionIndicator")}

type _TextInsertionIndicatorClass struct {
	class objc.Class
}

// A view that represents the insertion indicator in text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator

type TextInsertionIndicator struct {
	View
}

// TextInsertionIndicatorFrom constructs a [TextInsertionIndicator] from an unsafe.Pointer.
//
// A view that represents the insertion indicator in text.
func TextInsertionIndicatorFrom(ptr unsafe.Pointer) TextInsertionIndicator {
	return TextInsertionIndicator{
		View: ViewFrom(ptr),
	}
}



