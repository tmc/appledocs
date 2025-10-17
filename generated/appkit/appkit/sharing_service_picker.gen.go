// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SharingServicePicker] class.
var SharingServicePickerClass objc.Class

func init() {
	SharingServicePickerClass = objc.GetClass("NSSharingServicePicker")
}

type SharingServicePicker struct {
	objc.ID
}

func SharingServicePickerFrom(ptr unsafe.Pointer) SharingServicePicker {
	return SharingServicePicker{
		ID: objc.ID(ptr),
	}
}


// Shows the picker interface and populates it with the relevant sharing services. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSharingServicePicker/show(relativeTo:of:preferredEdge:)
func (s_ SharingServicePicker) ShowRelativeToRectOfViewPreferredEdge(rect foundation.Rect, view unsafe.Pointer, preferredEdge foundation.RectEdge) {
	sel := objc.RegisterName("showRelativeToRect:ofView:preferredEdge:")
	s_.ID.Send(sel, rect, view, preferredEdge)
}

