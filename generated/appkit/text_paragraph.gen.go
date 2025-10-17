// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextParagraph] class.
var textParagraphClass = _TextParagraphClass{objc.GetClass("NSTextParagraph")}

type _TextParagraphClass struct {
	class objc.Class
}

// A class that represents a single paragraph backed by an attributed string as the contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextParagraph

type TextParagraph struct {
	TextElement
}

// TextParagraphFrom constructs a [TextParagraph] from an unsafe.Pointer.
//
// A class that represents a single paragraph backed by an attributed string as the contents.
func TextParagraphFrom(ptr unsafe.Pointer) TextParagraph {
	return TextParagraph{
		TextElement: TextElementFrom(ptr),
	}
}



