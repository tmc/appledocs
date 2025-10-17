// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextTableBlock] class.
var textTableBlockClass = _TextTableBlockClass{objc.GetClass("NSTextTableBlock")}

type _TextTableBlockClass struct {
	class objc.Class
}

// An interface definition for the [TextTableBlock] class.
type ITextTableBlock interface {
	ITextBlock
}

// A text block that appears as a cell in a text table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTableBlock

type TextTableBlock struct {
	TextBlock
}

// TextTableBlockFrom constructs a [TextTableBlock] from an unsafe.Pointer.
//
// A text block that appears as a cell in a text table.
func TextTableBlockFrom(ptr unsafe.Pointer) TextTableBlock {
	return TextTableBlock{
		TextBlock: TextBlockFrom(ptr),
	}
}



