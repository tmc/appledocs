// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextAttachmentCell] class.
var textAttachmentCellClass = _TextAttachmentCellClass{objc.GetClass("NSTextAttachmentCell")}

type _TextAttachmentCellClass struct {
	class objc.Class
}

// An object that implements the functionality of the text attachment cell protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCell-swift.class

type TextAttachmentCell struct {
	Cell
}

// TextAttachmentCellFrom constructs a [TextAttachmentCell] from an unsafe.Pointer.
//
// An object that implements the functionality of the text attachment cell protocol.
func TextAttachmentCellFrom(ptr unsafe.Pointer) TextAttachmentCell {
	return TextAttachmentCell{
		Cell: CellFrom(ptr),
	}
}



