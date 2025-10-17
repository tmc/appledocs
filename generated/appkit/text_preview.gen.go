// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextPreview] class.
var textPreviewClass = _TextPreviewClass{objc.GetClass("NSTextPreview")}

type _TextPreviewClass struct {
	class objc.Class
}

// A snapshot of the text in your view, which the system uses to create user-visible effects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextPreview

type TextPreview struct {
	objectivec.Object
}

// TextPreviewFrom constructs a [TextPreview] from an unsafe.Pointer.
//
// A snapshot of the text in your view, which the system uses to create user-visible effects.
func TextPreviewFrom(ptr unsafe.Pointer) TextPreview {
	return TextPreview{objectivec.Object{objc.ID(ptr)}}
}



