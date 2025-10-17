// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AttributedStringMarkdownParsingOptions] class.
var attributedStringMarkdownParsingOptionsClass = _AttributedStringMarkdownParsingOptionsClass{objc.GetClass("NSAttributedStringMarkdownParsingOptions")}

type _AttributedStringMarkdownParsingOptionsClass struct {
	class objc.Class
}

// An interface definition for the [AttributedStringMarkdownParsingOptions] class.
type IAttributedStringMarkdownParsingOptions interface {
	objectivec.IObject
}

// Options that affect the parsing of Markdown content into an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions

type AttributedStringMarkdownParsingOptions struct {
	objectivec.Object
}

// AttributedStringMarkdownParsingOptionsFrom constructs a [AttributedStringMarkdownParsingOptions] from an unsafe.Pointer.
//
// Options that affect the parsing of Markdown content into an attributed string.
func AttributedStringMarkdownParsingOptionsFrom(ptr unsafe.Pointer) AttributedStringMarkdownParsingOptions {
	return AttributedStringMarkdownParsingOptions{objectivec.Object{objc.ID(ptr)}}
}



