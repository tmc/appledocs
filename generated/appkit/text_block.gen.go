// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextBlock] class.
var textBlockClass = _TextBlockClass{objc.GetClass("NSTextBlock")}

type _TextBlockClass struct {
	class objc.Class
}

// An interface definition for the [TextBlock] class.
type ITextBlock interface {
	objectivec.IObject
}

// A block of text laid out in a subregion of the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock

type TextBlock struct {
	objectivec.Object
}

// TextBlockFrom constructs a [TextBlock] from an unsafe.Pointer.
//
// A block of text laid out in a subregion of the text container.
func TextBlockFrom(ptr unsafe.Pointer) TextBlock {
	return TextBlock{objectivec.Object{objc.ID(ptr)}}
}



