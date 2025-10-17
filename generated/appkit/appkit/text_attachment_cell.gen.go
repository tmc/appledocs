// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextAttachmentCell] class.
var TextAttachmentCellClass objc.Class

func init() {
	TextAttachmentCellClass = objc.GetClass("NSTextAttachmentCell")
}

type TextAttachmentCell struct {
	objc.ID
}

func TextAttachmentCellFrom(ptr unsafe.Pointer) TextAttachmentCell {
	return TextAttachmentCell{
		ID: objc.ID(ptr),
	}
}



