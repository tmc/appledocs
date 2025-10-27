// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTouchBarDelegate is the NSTouchBarDelegate protocol interface.
//
// A protocol that allows you to provide the items for a bar dynamically.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTouchBarDelegate
type PTouchBarDelegate interface {
	// Optional methods
	TouchBarMakeItemForIdentifier(touchBar TouchBar /* not a class type */, identifier TouchBarItemIdentifier) ITouchBarItem
	HasTouchBarMakeItemForIdentifier() bool
}

// TouchBarDelegate is a delegate implementation builder for the PTouchBarDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TouchBarDelegate struct {
	_TouchBarMakeItemForIdentifier func(touchBar TouchBar /* not a class type */, identifier TouchBarItemIdentifier) ITouchBarItem
}

// SetTouchBarMakeItemForIdentifier sets the handler for the TouchBarMakeItemForIdentifier delegate method.
//
// Asks the delegate object for the bar item for the specified bar and item identifier.
func (d *TouchBarDelegate) SetTouchBarMakeItemForIdentifier(f func(touchBar TouchBar /* not a class type */, identifier TouchBarItemIdentifier) ITouchBarItem) {
	d._TouchBarMakeItemForIdentifier = f
}

// TouchBarMakeItemForIdentifier implements the PTouchBarDelegate interface.
func (d *TouchBarDelegate) TouchBarMakeItemForIdentifier(touchBar TouchBar /* not a class type */, identifier TouchBarItemIdentifier) ITouchBarItem {
	if d._TouchBarMakeItemForIdentifier != nil {
		return d._TouchBarMakeItemForIdentifier(touchBar, identifier)
	}
	var zero ITouchBarItem
	return zero
}

// HasTouchBarMakeItemForIdentifier returns true if a handler for TouchBarMakeItemForIdentifier has been set.
func (d *TouchBarDelegate) HasTouchBarMakeItemForIdentifier() bool {
	return d._TouchBarMakeItemForIdentifier != nil
}

// TouchBarDelegateObject wraps an existing Objective-C object that conforms to the PTouchBarDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TouchBarDelegateObject struct {
	objectivec.Object
}

// NewTouchBarDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTouchBarDelegate protocol.
func NewTouchBarDelegateObject(obj objectivec.Object) *TouchBarDelegateObject {
	return &TouchBarDelegateObject{obj}
}

// Make sure TouchBarDelegateObject implements PTouchBarDelegate.
var _ PTouchBarDelegate = (*TouchBarDelegateObject)(nil)

// TouchBarMakeItemForIdentifier implements the PTouchBarDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TouchBarDelegateObject) TouchBarMakeItemForIdentifier(touchBar TouchBar /* not a class type */, identifier TouchBarItemIdentifier) ITouchBarItem {
	return objc.Send[ITouchBarItem](o.ID, objc.Sel("touchBar:makeItemForIdentifier:"), touchBar, identifier)
}

// HasTouchBarMakeItemForIdentifier returns true; this is a placeholder for optional method checks.
func (o *TouchBarDelegateObject) HasTouchBarMakeItemForIdentifier() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
