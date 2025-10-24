// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
	TouchBarMakeItemForIdentifier(touchBar objc.IObject /* cross-framework: TouchBar */, identifier objc.IObject /* cross-framework: TouchBarItemIdentifier */) TouchBarItem
	HasTouchBarMakeItemForIdentifier() bool
}

// TouchBarDelegate is a delegate implementation builder for the PTouchBarDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TouchBarDelegate struct {
	_TouchBarMakeItemForIdentifier func(touchBar objc.IObject /* cross-framework: TouchBar */, identifier objc.IObject /* cross-framework: TouchBarItemIdentifier */) TouchBarItem
}

// SetTouchBarMakeItemForIdentifier sets the handler for the TouchBarMakeItemForIdentifier delegate method.
//
// Asks the delegate object for the bar item for the specified bar and item identifier.
func (d *TouchBarDelegate) SetTouchBarMakeItemForIdentifier(f func(touchBar objc.IObject /* cross-framework: TouchBar */, identifier objc.IObject /* cross-framework: TouchBarItemIdentifier */) TouchBarItem) {
	d._TouchBarMakeItemForIdentifier = f
}

// TouchBarMakeItemForIdentifier implements the PTouchBarDelegate interface.
func (d *TouchBarDelegate) TouchBarMakeItemForIdentifier(touchBar objc.IObject /* cross-framework: TouchBar */, identifier objc.IObject /* cross-framework: TouchBarItemIdentifier */) TouchBarItem {
	if d._TouchBarMakeItemForIdentifier != nil {
		return d._TouchBarMakeItemForIdentifier(touchBar, identifier)
	}
	var zero TouchBarItem
	return zero
}

// HasTouchBarMakeItemForIdentifier returns true if a handler for TouchBarMakeItemForIdentifier has been set.
func (d *TouchBarDelegate) HasTouchBarMakeItemForIdentifier() bool {
	return d._TouchBarMakeItemForIdentifier != nil
}
