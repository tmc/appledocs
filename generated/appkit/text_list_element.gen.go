// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextListElement] class.
var textListElementClass = _TextListElementClass{objc.GetClass("NSTextListElement")}

type _TextListElementClass struct {
	class objc.Class
}

// An interface definition for the [TextListElement] class.
type ITextListElement interface {
	ITextParagraph
}

// A class that represents a text list node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement

type TextListElement struct {
	TextParagraph
}

// TextListElementFrom constructs a [TextListElement] from an unsafe.Pointer.
//
// A class that represents a text list node.
func TextListElementFrom(ptr unsafe.Pointer) TextListElement {
	return TextListElement{
		TextParagraph: TextParagraphFrom(ptr),
	}
}



