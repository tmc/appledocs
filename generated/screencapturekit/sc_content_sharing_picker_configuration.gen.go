// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ContentSharingPickerConfiguration] class.
var (
	ContentSharingPickerConfigurationClass     _ContentSharingPickerConfigurationClass
	ContentSharingPickerConfigurationClassOnce sync.Once
)

func getContentSharingPickerConfigurationClass() _ContentSharingPickerConfigurationClass {
	ContentSharingPickerConfigurationClassOnce.Do(func() {
		ContentSharingPickerConfigurationClass = _ContentSharingPickerConfigurationClass{objc.GetClass("SCContentSharingPickerConfiguration")}
	})
	return ContentSharingPickerConfigurationClass
}

type _ContentSharingPickerConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [ContentSharingPickerConfiguration] class.
type IContentSharingPickerConfiguration interface {
	objectivec.IObject
}

// An instance for configuring the system content-sharing picker.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class
type ContentSharingPickerConfiguration struct {
	objectivec.Object
}

// ContentSharingPickerConfigurationFrom constructs a [ContentSharingPickerConfiguration] from an unsafe.Pointer.
//
// An instance for configuring the system content-sharing picker.
func ContentSharingPickerConfigurationFrom(ptr unsafe.Pointer) ContentSharingPickerConfiguration {
	return ContentSharingPickerConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContentSharingPickerConfigurationClass) Alloc() ContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContentSharingPickerConfigurationClass) New() ContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentSharingPickerConfiguration) Init() ContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentSharingPickerConfiguration) Autorelease() ContentSharingPickerConfiguration {
	rv := objc.Send[ContentSharingPickerConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentSharingPickerConfiguration creates a new ContentSharingPickerConfiguration instance.
func NewContentSharingPickerConfiguration() ContentSharingPickerConfiguration {
	return getContentSharingPickerConfigurationClass().New()
}


// The content-selection modes supported by the picker.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/allowedPickerModes
func (c_ ContentSharingPickerConfiguration) AllowedPickerModes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("allowedPickerModes"))
	return rv
}


// SetAllowedPickerModes sets the value of the allowedPickerModes property.
// The content-selection modes supported by the picker.

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/allowedPickerModes
func (c_ ContentSharingPickerConfiguration) SetAllowedPickerModes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowedPickerModes:"), value)
}

// A list of bundle IDs to exclude from the sharing picker.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/excludedBundleIDs
func (c_ ContentSharingPickerConfiguration) ExcludedBundleIDs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("excludedBundleIDs"))
	return rv
}


// SetExcludedBundleIDs sets the value of the excludedBundleIDs property.
// A list of bundle IDs to exclude from the sharing picker.

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/excludedBundleIDs
func (c_ ContentSharingPickerConfiguration) SetExcludedBundleIDs(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setExcludedBundleIDs:"), nsArray)
}



