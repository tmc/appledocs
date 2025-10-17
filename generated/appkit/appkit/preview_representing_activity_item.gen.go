// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PreviewRepresentingActivityItem] class.
var PreviewRepresentingActivityItemClass objc.Class

func init() {
	PreviewRepresentingActivityItemClass = objc.GetClass("NSPreviewRepresentingActivityItem")
}

type PreviewRepresentingActivityItem struct {
	objc.ID
}

func PreviewRepresentingActivityItemFrom(ptr unsafe.Pointer) PreviewRepresentingActivityItem {
	return PreviewRepresentingActivityItem{
		ID: objc.ID(ptr),
	}
}



