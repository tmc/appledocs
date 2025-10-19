// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCContentSharingPicker] class.
var (
	sCContentSharingPickerClass     _SCContentSharingPickerClass
	sCContentSharingPickerClassOnce sync.Once
)

func getSCContentSharingPickerClass() _SCContentSharingPickerClass {
	sCContentSharingPickerClassOnce.Do(func() {
		sCContentSharingPickerClass = _SCContentSharingPickerClass{objc.GetClass("SCContentSharingPicker")}
	})
	return sCContentSharingPickerClass
}

type _SCContentSharingPickerClass struct {
	class objc.Class
}

// An interface definition for the [SCContentSharingPicker] class.
type ISCContentSharingPicker interface {
	objectivec.IObject
}

// An instance of a picker presented by the operating system for managing frame-capture streams.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker
type SCContentSharingPicker struct {
	objectivec.Object
}

// SCContentSharingPickerFrom constructs a [SCContentSharingPicker] from an unsafe.Pointer.
//
// An instance of a picker presented by the operating system for managing frame-capture streams.
func SCContentSharingPickerFrom(ptr unsafe.Pointer) SCContentSharingPicker {
	return SCContentSharingPicker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCContentSharingPickerClass) Alloc() SCContentSharingPicker {
	rv := objc.Send[SCContentSharingPicker](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCContentSharingPickerClass) New() SCContentSharingPicker {
	rv := objc.Send[SCContentSharingPicker](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCContentSharingPicker) Init() SCContentSharingPicker {
	rv := objc.Send[SCContentSharingPicker](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCContentSharingPicker) Autorelease() SCContentSharingPicker {
	rv := objc.Send[SCContentSharingPicker](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCContentSharingPicker creates a new SCContentSharingPicker instance.
func NewSCContentSharingPicker() SCContentSharingPicker {
	return getSCContentSharingPickerClass().New()
}




