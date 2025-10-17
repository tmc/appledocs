// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextAttachmentViewProvider] class.
var textAttachmentViewProviderClass = _TextAttachmentViewProviderClass{objc.GetClass("NSTextAttachmentViewProvider")}

type _TextAttachmentViewProviderClass struct {
	class objc.Class
}

// A container object that associates a text attachment at a particular document location with a view object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAttachmentViewProvider

type TextAttachmentViewProvider struct {
	objectivec.Object
}

// TextAttachmentViewProviderFrom constructs a [TextAttachmentViewProvider] from an unsafe.Pointer.
//
// A container object that associates a text attachment at a particular document location with a view object.
func TextAttachmentViewProviderFrom(ptr unsafe.Pointer) TextAttachmentViewProvider {
	return TextAttachmentViewProvider{objectivec.Object{objc.ID(ptr)}}
}



