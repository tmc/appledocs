// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextPreview] class.
var TextPreviewClass objc.Class

func init() {
	TextPreviewClass = objc.GetClass("NSTextPreview")
}

type TextPreview struct {
	objc.ID
}

func TextPreviewFrom(ptr unsafe.Pointer) TextPreview {
	return TextPreview{
		ID: objc.ID(ptr),
	}
}




