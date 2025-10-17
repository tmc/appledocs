
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// The class instance for the [SharingServicePicker] class.
var SharingServicePickerClass _SharingServicePickerClass

func init() {
	SharingServicePickerClass = _SharingServicePickerClass{objc.GetClass("NSSharingServicePicker")}
}

type _SharingServicePickerClass struct {
	objc.Class
}

// An interface definition for the [SharingServicePicker] class.
type ISharingServicePicker interface {
	ID() objc.ID
	ShowRelativeToRectOfViewPreferredEdge(rect foundation.Rect, view unsafe.Pointer, preferredEdge foundation.RectEdge)
}

type SharingServicePicker struct {
	id objc.ID
}

func SharingServicePickerFrom(ptr unsafe.Pointer) SharingServicePicker {
	return SharingServicePicker{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SharingServicePicker) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerClass) Alloc() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SharingServicePickerClass) New() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSharingServicePicker creates and returns a new initialized instance.
func NewSharingServicePicker() SharingServicePicker {
	return SharingServicePickerClass.New()
}

// Init initializes the instance.
func (s_ SharingServicePicker) Init() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](s_.ID(), selInit)
	return rv
}
// Shows the picker interface and populates it with the relevant sharing services. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSharingServicePicker/show(relativeTo:of:preferredEdge:)
func (s_ SharingServicePicker) ShowRelativeToRectOfViewPreferredEdge(rect foundation.Rect, view unsafe.Pointer, preferredEdge foundation.RectEdge) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("showRelativeToRect:ofView:preferredEdge:"), rect, view, preferredEdge)
}
// The object for managing the sharing service picker. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSharingServicePicker/delegate
func (s_ SharingServicePicker) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSharingServicePicker/delegate
func (s_ SharingServicePicker) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDelegate:"), value)
}
