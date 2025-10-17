// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AttributedStringMarkdownParsingOptions] class.
var AttributedStringMarkdownParsingOptionsClass = _AttributedStringMarkdownParsingOptionsClass{objc.GetClass("NSAttributedStringMarkdownParsingOptions")}

type _AttributedStringMarkdownParsingOptionsClass struct {
	class objc.Class
}

type AttributedStringMarkdownParsingOptions struct {
	objc.ID
}

func AttributedStringMarkdownParsingOptionsFrom(ptr unsafe.Pointer) AttributedStringMarkdownParsingOptions {
	return AttributedStringMarkdownParsingOptions{
		ID: objc.ID(ptr),
	}
}




