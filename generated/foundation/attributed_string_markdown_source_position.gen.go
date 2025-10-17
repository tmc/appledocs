// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AttributedStringMarkdownSourcePosition] class.
var AttributedStringMarkdownSourcePositionClass objc.Class

func init() {
	AttributedStringMarkdownSourcePositionClass = objc.GetClass("NSAttributedStringMarkdownSourcePosition")
}

type AttributedStringMarkdownSourcePosition struct {
	objc.ID
}

func AttributedStringMarkdownSourcePositionFrom(ptr unsafe.Pointer) AttributedStringMarkdownSourcePosition {
	return AttributedStringMarkdownSourcePosition{
		ID: objc.ID(ptr),
	}
}



