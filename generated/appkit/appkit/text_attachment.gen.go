// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextAttachment] class.
var TextAttachmentClass objc.Class

func init() {
	TextAttachmentClass = objc.GetClass("NSTextAttachment")
}

type TextAttachment struct {
	objc.ID
}

func TextAttachmentFrom(ptr unsafe.Pointer) TextAttachment {
	return TextAttachment{
		ID: objc.ID(ptr),
	}
}




