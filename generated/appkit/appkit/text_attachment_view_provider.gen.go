// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextAttachmentViewProvider] class.
var TextAttachmentViewProviderClass objc.Class

func init() {
	TextAttachmentViewProviderClass = objc.GetClass("NSTextAttachmentViewProvider")
}

type TextAttachmentViewProvider struct {
	objc.ID
}

func TextAttachmentViewProviderFrom(ptr unsafe.Pointer) TextAttachmentViewProvider {
	return TextAttachmentViewProvider{
		ID: objc.ID(ptr),
	}
}




