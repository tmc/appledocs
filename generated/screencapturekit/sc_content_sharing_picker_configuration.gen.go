// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCContentSharingPickerConfiguration] class.
var (
	sCContentSharingPickerConfigurationClass     _SCContentSharingPickerConfigurationClass
	sCContentSharingPickerConfigurationClassOnce sync.Once
)

func getSCContentSharingPickerConfigurationClass() _SCContentSharingPickerConfigurationClass {
	sCContentSharingPickerConfigurationClassOnce.Do(func() {
		sCContentSharingPickerConfigurationClass = _SCContentSharingPickerConfigurationClass{objc.GetClass("SCContentSharingPickerConfiguration")}
	})
	return sCContentSharingPickerConfigurationClass
}

type _SCContentSharingPickerConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [SCContentSharingPickerConfiguration] class.
type ISCContentSharingPickerConfiguration interface {
	objectivec.IObject
}

// An instance for configuring the system content-sharing picker.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class
type SCContentSharingPickerConfiguration struct {
	objectivec.Object
}

// SCContentSharingPickerConfigurationFrom constructs a [SCContentSharingPickerConfiguration] from an unsafe.Pointer.
//
// An instance for configuring the system content-sharing picker.
func SCContentSharingPickerConfigurationFrom(ptr unsafe.Pointer) SCContentSharingPickerConfiguration {
	return SCContentSharingPickerConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCContentSharingPickerConfigurationClass) Alloc() SCContentSharingPickerConfiguration {
	rv := objc.Send[SCContentSharingPickerConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCContentSharingPickerConfigurationClass) New() SCContentSharingPickerConfiguration {
	rv := objc.Send[SCContentSharingPickerConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCContentSharingPickerConfiguration) Init() SCContentSharingPickerConfiguration {
	rv := objc.Send[SCContentSharingPickerConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCContentSharingPickerConfiguration) Autorelease() SCContentSharingPickerConfiguration {
	rv := objc.Send[SCContentSharingPickerConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCContentSharingPickerConfiguration creates a new SCContentSharingPickerConfiguration instance.
func NewSCContentSharingPickerConfiguration() SCContentSharingPickerConfiguration {
	return getSCContentSharingPickerConfigurationClass().New()
}




