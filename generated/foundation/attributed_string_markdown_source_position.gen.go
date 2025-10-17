// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AttributedStringMarkdownSourcePosition] class.
var AttributedStringMarkdownSourcePositionClass = _AttributedStringMarkdownSourcePositionClass{objc.GetClass("NSAttributedStringMarkdownSourcePosition")}

type _AttributedStringMarkdownSourcePositionClass struct {
	class objc.Class
}

type AttributedStringMarkdownSourcePosition struct {
	objc.ID
}

func AttributedStringMarkdownSourcePositionFrom(ptr unsafe.Pointer) AttributedStringMarkdownSourcePosition {
	return AttributedStringMarkdownSourcePosition{
		ID: objc.ID(ptr),
	}
}




