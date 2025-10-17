// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PreviewRepresentingActivityItem] class.
var previewRepresentingActivityItemClass = _PreviewRepresentingActivityItemClass{objc.GetClass("NSPreviewRepresentingActivityItem")}

type _PreviewRepresentingActivityItemClass struct {
	class objc.Class
}

// An interface definition for the [PreviewRepresentingActivityItem] class.
type IPreviewRepresentingActivityItem interface {
	objectivec.IObject
}

// A type that adds metadata to an item you share using the macOS share sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPreviewRepresentingActivityItem

type PreviewRepresentingActivityItem struct {
	objectivec.Object
}

// PreviewRepresentingActivityItemFrom constructs a [PreviewRepresentingActivityItem] from an unsafe.Pointer.
//
// A type that adds metadata to an item you share using the macOS share sheet.
func PreviewRepresentingActivityItemFrom(ptr unsafe.Pointer) PreviewRepresentingActivityItem {
	return PreviewRepresentingActivityItem{objectivec.Object{objc.ID(ptr)}}
}



