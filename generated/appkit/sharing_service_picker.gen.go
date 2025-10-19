// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SharingServicePicker] class.
var (
	sharingServicePickerClass     _SharingServicePickerClass
	sharingServicePickerClassOnce sync.Once
)

func getSharingServicePickerClass() _SharingServicePickerClass {
	sharingServicePickerClassOnce.Do(func() {
		sharingServicePickerClass = _SharingServicePickerClass{objc.GetClass("NSSharingServicePicker")}
	})
	return sharingServicePickerClass
}

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
// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerClass) Alloc() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SharingServicePickerClass) New() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingServicePicker) Init() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingServicePicker) Autorelease() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingServicePicker creates a new SharingServicePicker instance.
func NewSharingServicePicker() SharingServicePicker {
	return getSharingServicePickerClass().New()
}


// Shows the picker interface and populates it with the relevant sharing services. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/show(relativeTo:of:preferredEdge:)
func (s_ SharingServicePicker) ShowRelativeToRectOfViewPreferredEdge(rect unsafe.Pointer, view unsafe.Pointer, preferredEdge int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), rect, view, preferredEdge)
}


