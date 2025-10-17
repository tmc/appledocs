// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextElement] class.
var textElementClass = _TextElementClass{objc.GetClass("NSTextElement")}

type _TextElementClass struct {
	class objc.Class
}

// An interface definition for the [TextElement] class.
type ITextElement interface {
	objectivec.IObject
}

// An abstract base class that represents the smallest units of text layout such as paragraphs or attachments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement

type TextElement struct {
	objectivec.Object
}

// TextElementFrom constructs a [TextElement] from an unsafe.Pointer.
//
// An abstract base class that represents the smallest units of text layout such as paragraphs or attachments.
func TextElementFrom(ptr unsafe.Pointer) TextElement {
	return TextElement{objectivec.Object{objc.ID(ptr)}}
}



