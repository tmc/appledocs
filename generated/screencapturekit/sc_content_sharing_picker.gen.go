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
	// properties:
	Configuration() IContentSharingPickerConfiguration
	SetConfiguration(value IContentSharingPickerConfiguration)
	DefaultConfiguration() IContentSharingPickerConfiguration
	SetDefaultConfiguration(value IContentSharingPickerConfiguration)
	IsActive() bool
	SetIsActive(value bool)
	MaximumStreamCount() int
	SetMaximumStreamCount(value int)
	// methods:
}

// An instance of a picker presented by the operating system for managing frame-capture streams.


// An instance of a picker presented by the operating system for managing frame-capture streams.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/shared
func (cc _ContentSharingPickerClass) SharedPicker() ContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](objc.ID(cc.class), objc.Sel("sharedPicker"))
	return rv
}

// The system-provided picker UI instance for capturing display and audio content from someone’s Mac.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/shared
func (c_ ContentSharingPicker) SharedPicker() ISCContentSharingPicker {
	rv := objc.Send[ContentSharingPicker](c_.ID, objc.Sel("sharedPicker"))
	return rv
}


// Sets the configuration for the content capture picker for all streams, providing allowed selection modes and content excluded from selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/configuration
func (c_ ContentSharingPicker) Configuration() IContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}


// Sets the configuration for the content capture picker for all streams, providing allowed selection modes and content excluded from selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/configuration
func (c_ ContentSharingPicker) SetConfiguration(value IContentSharingPickerConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}


// The default configuration to use for the content capture picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/defaultconfiguration-94q2b
func (c_ ContentSharingPicker) DefaultConfiguration() IContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](c_.ID, objc.Sel("defaultConfiguration"))
	return rv
}


// The default configuration to use for the content capture picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/defaultconfiguration-94q2b
func (c_ ContentSharingPicker) SetDefaultConfiguration(value IContentSharingPickerConfiguration) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDefaultConfiguration:"), value)
}


// A Boolean value that indicates if the picker is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/isactive
func (c_ ContentSharingPicker) IsActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isActive"))
	return rv
}


// A Boolean value that indicates if the picker is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/isactive
func (c_ ContentSharingPicker) SetIsActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsActive:"), value)
}


// The maximum number of streams the content capture picker allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/maximumstreamcount-2kuaa
func (c_ ContentSharingPicker) MaximumStreamCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maximumStreamCount"))
	return rv
}


// The maximum number of streams the content capture picker allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/maximumstreamcount-2kuaa
func (c_ ContentSharingPicker) SetMaximumStreamCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumStreamCount:"), value)
}



