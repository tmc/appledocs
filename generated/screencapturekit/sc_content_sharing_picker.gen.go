// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ContentSharingPicker] class.
var (
	ContentSharingPickerClass     _ContentSharingPickerClass
	ContentSharingPickerClassOnce sync.Once
)

func getContentSharingPickerClass() _ContentSharingPickerClass {
	ContentSharingPickerClassOnce.Do(func() {
		ContentSharingPickerClass = _ContentSharingPickerClass{objc.GetClass("SCContentSharingPicker")}
	})
	return ContentSharingPickerClass
}

type _ContentSharingPickerClass struct {
	class objc.Class
}

// An interface definition for the [ContentSharingPicker] class.
type IContentSharingPicker interface {
	objectivec.IObject
}

// An instance of a picker presented by the operating system for managing frame-capture streams.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker
type ContentSharingPicker struct {
	objectivec.Object
}

// ContentSharingPickerFrom constructs a [ContentSharingPicker] from an unsafe.Pointer.
//
// An instance of a picker presented by the operating system for managing frame-capture streams.
func ContentSharingPickerFrom(ptr unsafe.Pointer) ContentSharingPicker {
	return ContentSharingPicker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContentSharingPickerClass) Alloc() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContentSharingPickerClass) New() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentSharingPicker) Init() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentSharingPicker) Autorelease() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentSharingPicker creates a new ContentSharingPicker instance.
func NewContentSharingPicker() ContentSharingPicker {
	return getContentSharingPickerClass().New()
}


// The system-provided picker UI instance for capturing display and audio content from someone’s Mac.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/shared
func (cc _ContentSharingPickerClass) SharedPicker() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("sharedPicker"))
	return rv
}
// The system-provided picker UI instance for capturing display and audio content from someone’s Mac.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/shared
func (c_ ContentSharingPicker) SharedPicker() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sharedPicker"))
	return rv
}



