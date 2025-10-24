// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PSharingServicePickerTouchBarItemDelegate is the NSSharingServicePickerTouchBarItemDelegate protocol interface.
//
// A protocol that a sharing service picker item delegate uses to provide a list of items eligible for sharing.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSharingServicePickerTouchBarItemDelegate
type PSharingServicePickerTouchBarItemDelegate interface {
	// Required methods
	ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem ISharingServicePickerTouchBarItem) foundation.Array/* debug [protocol_interface/required_method]: ItemsForSharingServicePickerTouchBarItem */
}

// SharingServicePickerTouchBarItemDelegate is a delegate implementation builder for the PSharingServicePickerTouchBarItemDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SharingServicePickerTouchBarItemDelegate struct {
	_ItemsForSharingServicePickerTouchBarItem func(pickerTouchBarItem ISharingServicePickerTouchBarItem) foundation.Array
}

// SetItemsForSharingServicePickerTouchBarItem sets the handler for the ItemsForSharingServicePickerTouchBarItem delegate method.
//
// Asks the delegate for items that represent the objects to be shared.
func (d *SharingServicePickerTouchBarItemDelegate) SetItemsForSharingServicePickerTouchBarItem(f func(pickerTouchBarItem ISharingServicePickerTouchBarItem) foundation.Array) {
	d._ItemsForSharingServicePickerTouchBarItem = f
}

// ItemsForSharingServicePickerTouchBarItem implements the PSharingServicePickerTouchBarItemDelegate interface.
func (d *SharingServicePickerTouchBarItemDelegate) ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem ISharingServicePickerTouchBarItem) foundation.Array {
	if d._ItemsForSharingServicePickerTouchBarItem != nil {
		return d._ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem)
	}
	var zero foundation.Array
	return zero
}

// HasItemsForSharingServicePickerTouchBarItem returns true if a handler for ItemsForSharingServicePickerTouchBarItem has been set.
func (d *SharingServicePickerTouchBarItemDelegate) HasItemsForSharingServicePickerTouchBarItem() bool {
	return d._ItemsForSharingServicePickerTouchBarItem != nil
}
