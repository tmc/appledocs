// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextRange] class.
var textRangeClass = _TextRangeClass{objc.GetClass("NSTextRange")}

type _TextRangeClass struct {
	class objc.Class
}

// An interface definition for the [TextRange] class.
type ITextRange interface {
	objectivec.IObject
}

// A class that represents a contiguous range between two locations inside document contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextRange

type TextRange struct {
	objectivec.Object
}

// TextRangeFrom constructs a [TextRange] from an unsafe.Pointer.
//
// A class that represents a contiguous range between two locations inside document contents.
func TextRangeFrom(ptr unsafe.Pointer) TextRange {
	return TextRange{objectivec.Object{objc.ID(ptr)}}
}



