// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AttributedStringMarkdownSourcePosition] class.
var attributedStringMarkdownSourcePositionClass = _AttributedStringMarkdownSourcePositionClass{objc.GetClass("NSAttributedStringMarkdownSourcePosition")}

type _AttributedStringMarkdownSourcePositionClass struct {
	class objc.Class
}

// An interface definition for the [AttributedStringMarkdownSourcePosition] class.
type IAttributedStringMarkdownSourcePosition interface {
	objectivec.IObject
}

// The position of attributed string text in its original Markdown source string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition

type AttributedStringMarkdownSourcePosition struct {
	objectivec.Object
}

// AttributedStringMarkdownSourcePositionFrom constructs a [AttributedStringMarkdownSourcePosition] from an unsafe.Pointer.
//
// The position of attributed string text in its original Markdown source string.
func AttributedStringMarkdownSourcePositionFrom(ptr unsafe.Pointer) AttributedStringMarkdownSourcePosition {
	return AttributedStringMarkdownSourcePosition{objectivec.Object{objc.ID(ptr)}}
}



