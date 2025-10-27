// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PSharingServicePickerToolbarItemDelegate is the NSSharingServicePickerToolbarItemDelegate protocol interface.
//
// An interface that provides the content to share from the macOS share sheet.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSharingServicePickerToolbarItemDelegate
type PSharingServicePickerToolbarItemDelegate interface {
	// Required methods
	ItemsForSharingServicePickerToolbarItem(pickerToolbarItem ISharingServicePickerToolbarItem) foundation.Array
}

// SharingServicePickerToolbarItemDelegate is a delegate implementation builder for the PSharingServicePickerToolbarItemDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SharingServicePickerToolbarItemDelegate struct {
	_ItemsForSharingServicePickerToolbarItem func(pickerToolbarItem ISharingServicePickerToolbarItem) foundation.Array
}

// SetItemsForSharingServicePickerToolbarItem sets the handler for the ItemsForSharingServicePickerToolbarItem delegate method.
//
// Asks the delegate for the items to share.
func (d *SharingServicePickerToolbarItemDelegate) SetItemsForSharingServicePickerToolbarItem(f func(pickerToolbarItem ISharingServicePickerToolbarItem) foundation.Array) {
	d._ItemsForSharingServicePickerToolbarItem = f
}

// ItemsForSharingServicePickerToolbarItem implements the PSharingServicePickerToolbarItemDelegate interface.
func (d *SharingServicePickerToolbarItemDelegate) ItemsForSharingServicePickerToolbarItem(pickerToolbarItem ISharingServicePickerToolbarItem) foundation.Array {
	if d._ItemsForSharingServicePickerToolbarItem != nil {
		return d._ItemsForSharingServicePickerToolbarItem(pickerToolbarItem)
	}
	var zero foundation.Array
	return zero
}

// HasItemsForSharingServicePickerToolbarItem returns true if a handler for ItemsForSharingServicePickerToolbarItem has been set.
func (d *SharingServicePickerToolbarItemDelegate) HasItemsForSharingServicePickerToolbarItem() bool {
	return d._ItemsForSharingServicePickerToolbarItem != nil
}

// SharingServicePickerToolbarItemDelegateObject wraps an existing Objective-C object that conforms to the PSharingServicePickerToolbarItemDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type SharingServicePickerToolbarItemDelegateObject struct {
	objectivec.Object
}

// NewSharingServicePickerToolbarItemDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSSharingServicePickerToolbarItemDelegate protocol.
func NewSharingServicePickerToolbarItemDelegateObject(obj objectivec.Object) *SharingServicePickerToolbarItemDelegateObject {
	return &SharingServicePickerToolbarItemDelegateObject{obj}
}

// Make sure SharingServicePickerToolbarItemDelegateObject implements PSharingServicePickerToolbarItemDelegate.
var _ PSharingServicePickerToolbarItemDelegate = (*SharingServicePickerToolbarItemDelegateObject)(nil)

// ItemsForSharingServicePickerToolbarItem implements the PSharingServicePickerToolbarItemDelegate interface.
// This required method is always available on objects conforming to ItemsForSharingServicePickerToolbarItem.
func (o *SharingServicePickerToolbarItemDelegateObject) ItemsForSharingServicePickerToolbarItem(pickerToolbarItem ISharingServicePickerToolbarItem) foundation.Array {
	return objc.Send[foundation.Array](o.ID, objc.Sel("itemsForSharingServicePickerToolbarItem:"), pickerToolbarItem)
}
