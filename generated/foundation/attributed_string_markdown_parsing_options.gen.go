// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AttributedStringMarkdownParsingOptions] class.
var AttributedStringMarkdownParsingOptionsClass objc.Class

func init() {
	AttributedStringMarkdownParsingOptionsClass = objc.GetClass("NSAttributedStringMarkdownParsingOptions")
}

type AttributedStringMarkdownParsingOptions struct {
	objc.ID
}

func AttributedStringMarkdownParsingOptionsFrom(ptr unsafe.Pointer) AttributedStringMarkdownParsingOptions {
	return AttributedStringMarkdownParsingOptions{
		ID: objc.ID(ptr),
	}
}



