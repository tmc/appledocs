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
// A Boolean value that indicates if the picker is active.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/isactive
func (c_ ContentSharingPicker) IsActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isActive"))
	return rv
}


// SetIsActive sets the value of the isActive property.
// A Boolean value that indicates if the picker is active.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/isactive
func (c_ ContentSharingPicker) SetIsActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsActive:"), value)
}

// Sets the configuration for the content capture picker for all streams, providing allowed selection modes and content excluded from selection.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/configuration
func (c_ ContentSharingPicker) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("configuration"))
	return rv
}


// SetConfiguration sets the value of the configuration property.
// Sets the configuration for the content capture picker for all streams, providing allowed selection modes and content excluded from selection.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/configuration
func (c_ ContentSharingPicker) SetConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}

// The default configuration to use for the content capture picker.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/defaultconfiguration-94q2b
func (c_ ContentSharingPicker) DefaultConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("defaultConfiguration"))
	return rv
}


// SetDefaultConfiguration sets the value of the defaultConfiguration property.
// The default configuration to use for the content capture picker.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/defaultconfiguration-94q2b
func (c_ ContentSharingPicker) SetDefaultConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultConfiguration:"), value)
}

// The maximum number of streams the content capture picker allows.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/maximumstreamcount-2kuaa
func (c_ ContentSharingPicker) MaximumStreamCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maximumStreamCount"))
	return rv
}


// SetMaximumStreamCount sets the value of the maximumStreamCount property.
// The maximum number of streams the content capture picker allows.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/maximumstreamcount-2kuaa
func (c_ ContentSharingPicker) SetMaximumStreamCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumStreamCount:"), value)
}

// The system-provided picker UI instance for capturing display and audio content from someone’s Mac.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/shared
func (c_ ContentSharingPicker) SharedPicker() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sharedPicker"))
	return rv
}



