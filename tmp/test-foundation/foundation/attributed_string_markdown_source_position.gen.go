// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var AttributedStringMarkdownSourcePositionClass _AttributedStringMarkdownSourcePositionClass

func init() {
	AttributedStringMarkdownSourcePositionClass = _AttributedStringMarkdownSourcePositionClass{objc.GetClass("NSAttributedStringMarkdownSourcePosition")}
}

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




