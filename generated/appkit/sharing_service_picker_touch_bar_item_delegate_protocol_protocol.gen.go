// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem ISharingServicePickerTouchBarItem) foundation.Array
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

// SharingServicePickerTouchBarItemDelegateObject wraps an existing Objective-C object that conforms to the PSharingServicePickerTouchBarItemDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type SharingServicePickerTouchBarItemDelegateObject struct {
	objectivec.Object
}

// NewSharingServicePickerTouchBarItemDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSSharingServicePickerTouchBarItemDelegate protocol.
func NewSharingServicePickerTouchBarItemDelegateObject(obj objectivec.Object) *SharingServicePickerTouchBarItemDelegateObject {
	return &SharingServicePickerTouchBarItemDelegateObject{obj}
}

// Make sure SharingServicePickerTouchBarItemDelegateObject implements PSharingServicePickerTouchBarItemDelegate.
var _ PSharingServicePickerTouchBarItemDelegate = (*SharingServicePickerTouchBarItemDelegateObject)(nil)

// ItemsForSharingServicePickerTouchBarItem implements the PSharingServicePickerTouchBarItemDelegate interface.
// This required method is always available on objects conforming to ItemsForSharingServicePickerTouchBarItem.
func (o *SharingServicePickerTouchBarItemDelegateObject) ItemsForSharingServicePickerTouchBarItem(pickerTouchBarItem ISharingServicePickerTouchBarItem) foundation.Array {
	return objc.Send[foundation.Array](o.ID, objc.Sel("itemsForSharingServicePickerTouchBarItem:"), pickerTouchBarItem)
}
