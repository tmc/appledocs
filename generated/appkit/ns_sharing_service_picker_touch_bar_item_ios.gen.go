//go:build darwin && ios

// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for SharingServicePickerTouchBarItem


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/activityItemsConfiguration
func (s_ SharingServicePickerTouchBarItem) ActivityItemsConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("activityItemsConfiguration"))
	return rv
}
func (s_ SharingServicePickerTouchBarItem) SetActivityItemsConfiguration(value unsafe.Pointer) {
	s_.ID.Send(objc.RegisterName("setActivityItemsConfiguration:"), value)
}





