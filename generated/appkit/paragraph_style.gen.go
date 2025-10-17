// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ParagraphStyle] class.
var paragraphStyleClass = _ParagraphStyleClass{objc.GetClass("NSParagraphStyle")}

type _ParagraphStyleClass struct {
	class objc.Class
}

// The paragraph or ruler attributes for an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle

type ParagraphStyle struct {
	objectivec.Object
}

// ParagraphStyleFrom constructs a [ParagraphStyle] from an unsafe.Pointer.
//
// The paragraph or ruler attributes for an attributed string.
func ParagraphStyleFrom(ptr unsafe.Pointer) ParagraphStyle {
	return ParagraphStyle{objectivec.Object{objc.ID(ptr)}}
}



