//go:build darwin && ios

// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for SharingServicePickerToolbarItem


// iOS-only properties

// The custom object from an app built with Mac Catalyst that provides the items to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem/activityItemsConfiguration
func (s_ SharingServicePickerToolbarItem) ActivityItemsConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("activityItemsConfiguration"))
	return rv
}
func (s_ SharingServicePickerToolbarItem) SetActivityItemsConfiguration(value unsafe.Pointer) {
	s_.ID.Send(objc.RegisterName("setActivityItemsConfiguration:"), value)
}





