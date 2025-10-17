// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextAttachment] class.
var textAttachmentClass = _TextAttachmentClass{objc.GetClass("NSTextAttachment")}

type _TextAttachmentClass struct {
	class objc.Class
}

// The values for the attachment characteristics of attributed strings and related objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachment

type TextAttachment struct {
	objectivec.Object
}

// TextAttachmentFrom constructs a [TextAttachment] from an unsafe.Pointer.
//
// The values for the attachment characteristics of attributed strings and related objects.
func TextAttachmentFrom(ptr unsafe.Pointer) TextAttachment {
	return TextAttachment{objectivec.Object{objc.ID(ptr)}}
}



