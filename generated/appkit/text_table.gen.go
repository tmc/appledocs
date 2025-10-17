// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextTable] class.
var textTableClass = _TextTableClass{objc.GetClass("NSTextTable")}

type _TextTableClass struct {
	class objc.Class
}

// An object that represents a text table as a whole. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable

type TextTable struct {
	TextBlock
}

// TextTableFrom constructs a [TextTable] from an unsafe.Pointer.
//
// An object that represents a text table as a whole.
func TextTableFrom(ptr unsafe.Pointer) TextTable {
	return TextTable{
		TextBlock: TextBlockFrom(ptr),
	}
}



