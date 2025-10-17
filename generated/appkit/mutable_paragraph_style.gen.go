// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableParagraphStyle] class.
var mutableParagraphStyleClass = _MutableParagraphStyleClass{objc.GetClass("NSMutableParagraphStyle")}

type _MutableParagraphStyleClass struct {
	class objc.Class
}

// An interface definition for the [MutableParagraphStyle] class.
type IMutableParagraphStyle interface {
	IParagraphStyle
}

// An object for changing the values of the subattributes in a paragraph style attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMutableParagraphStyle

type MutableParagraphStyle struct {
	ParagraphStyle
}

// MutableParagraphStyleFrom constructs a [MutableParagraphStyle] from an unsafe.Pointer.
//
// An object for changing the values of the subattributes in a paragraph style attribute.
func MutableParagraphStyleFrom(ptr unsafe.Pointer) MutableParagraphStyle {
	return MutableParagraphStyle{
		ParagraphStyle: ParagraphStyleFrom(ptr),
	}
}



