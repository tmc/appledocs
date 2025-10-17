// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextParagraph] class.
var TextParagraphClass objc.Class

func init() {
	TextParagraphClass = objc.GetClass("NSTextParagraph")
}

type TextParagraph struct {
	objc.ID
}

func TextParagraphFrom(ptr unsafe.Pointer) TextParagraph {
	return TextParagraph{
		ID: objc.ID(ptr),
	}
}



