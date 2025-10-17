// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SharingServicePicker] class.
var sharingServicePickerClass = _SharingServicePickerClass{objc.GetClass("NSSharingServicePicker")}

type _SharingServicePickerClass struct {
	class objc.Class
}

// An interface definition for the [SharingServicePicker] class.
type ISharingServicePicker interface {
	objectivec.IObject
	ShowRelativeToRectOfViewPreferredEdge(rect unsafe.Pointer, view unsafe.Pointer, preferredEdge int)
}

// A list of sharing services that the user can choose from. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker

type SharingServicePicker struct {
	objectivec.Object
}

// SharingServicePickerFrom constructs a [SharingServicePicker] from an unsafe.Pointer.
//
// A list of sharing services that the user can choose from.
func SharingServicePickerFrom(ptr unsafe.Pointer) SharingServicePicker {
	return SharingServicePicker{objectivec.Object{objc.ID(ptr)}}
}

// Shows the picker interface and populates it with the relevant sharing services. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/show(relativeTo:of:preferredEdge:)
func (s_ SharingServicePicker) ShowRelativeToRectOfViewPreferredEdge(rect unsafe.Pointer, view unsafe.Pointer, preferredEdge int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), rect, view, preferredEdge)
}


