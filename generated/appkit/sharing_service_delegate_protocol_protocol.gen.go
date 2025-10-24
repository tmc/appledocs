// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PSharingServiceDelegate is the NSSharingServiceDelegate protocol interface.
//
// A set of methods that you use to customize the position and animation of a share sheet, and to be notified whether the item is successfully shared.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSharingServiceDelegate
type PSharingServiceDelegate interface {
	// Optional methods
	AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService ISharingService, positioningRect objc.IObject /* cross-framework: Rect */, preferredEdge RectEdge /* not a class type */) View
	HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool
	SharingServiceDidFailToShareItemsError(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, error_ objc.IObject /* cross-framework: Error */)
	HasSharingServiceDidFailToShareItemsError() bool
	SharingServiceDidShareItems(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */)
	HasSharingServiceDidShareItems() bool
	SharingServiceSourceFrameOnScreenForShareItem(sharingService ISharingService, item objc.IObject) corefoundation.Rect
	HasSharingServiceSourceFrameOnScreenForShareItem() bool
	SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, sharingContentScope SharingContentScope) Window
	HasSharingServiceSourceWindowForShareItemsSharingContentScope() bool
	SharingServiceTransitionImageForShareItemContentRect(sharingService ISharingService, item objc.IObject, contentRect objc.IObject /* cross-framework: Rect */) Image
	HasSharingServiceTransitionImageForShareItemContentRect() bool
	SharingServiceWillShareItems(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */)
	HasSharingServiceWillShareItems() bool
}

// SharingServiceDelegate is a delegate implementation builder for the PSharingServiceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SharingServiceDelegate struct {
	_AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge func(sharingService ISharingService, positioningRect objc.IObject /* cross-framework: Rect */, preferredEdge RectEdge /* not a class type */) View
	_SharingServiceDidFailToShareItemsError func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, error_ objc.IObject /* cross-framework: Error */)
	_SharingServiceDidShareItems func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */)
	_SharingServiceSourceFrameOnScreenForShareItem func(sharingService ISharingService, item objc.IObject) corefoundation.Rect
	_SharingServiceSourceWindowForShareItemsSharingContentScope func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, sharingContentScope SharingContentScope) Window
	_SharingServiceTransitionImageForShareItemContentRect func(sharingService ISharingService, item objc.IObject, contentRect objc.IObject /* cross-framework: Rect */) Image
	_SharingServiceWillShareItems func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */)
}

// SetAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge sets the handler for the AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge delegate method.
func (d *SharingServiceDelegate) SetAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(f func(sharingService ISharingService, positioningRect objc.IObject /* cross-framework: Rect */, preferredEdge RectEdge /* not a class type */) View) {
	d._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge = f
}

// SetSharingServiceDidFailToShareItemsError sets the handler for the SharingServiceDidFailToShareItemsError delegate method.
//
// Invoked when the sharing service encountered an error when sharing items.
func (d *SharingServiceDelegate) SetSharingServiceDidFailToShareItemsError(f func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, error_ objc.IObject /* cross-framework: Error */)) {
	d._SharingServiceDidFailToShareItemsError = f
}

// SetSharingServiceDidShareItems sets the handler for the SharingServiceDidShareItems delegate method.
//
// Invoked when the sharing service has finished sharing the items.
func (d *SharingServiceDelegate) SetSharingServiceDidShareItems(f func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */)) {
	d._SharingServiceDidShareItems = f
}

// SetSharingServiceSourceFrameOnScreenForShareItem sets the handler for the SharingServiceSourceFrameOnScreenForShareItem delegate method.
//
// Invoked when the sharing service is performed and the sharing window is displayed, to present a transition between the original items and the sharing window.
func (d *SharingServiceDelegate) SetSharingServiceSourceFrameOnScreenForShareItem(f func(sharingService ISharingService, item objc.IObject) corefoundation.Rect) {
	d._SharingServiceSourceFrameOnScreenForShareItem = f
}

// SetSharingServiceSourceWindowForShareItemsSharingContentScope sets the handler for the SharingServiceSourceWindowForShareItemsSharingContentScope delegate method.
//
// Returns the window that contained the share items.
func (d *SharingServiceDelegate) SetSharingServiceSourceWindowForShareItemsSharingContentScope(f func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, sharingContentScope SharingContentScope) Window) {
	d._SharingServiceSourceWindowForShareItemsSharingContentScope = f
}

// SetSharingServiceTransitionImageForShareItemContentRect sets the handler for the SharingServiceTransitionImageForShareItemContentRect delegate method.
//
// Invoked to allow returning a custom transition image when sharing an item.
func (d *SharingServiceDelegate) SetSharingServiceTransitionImageForShareItemContentRect(f func(sharingService ISharingService, item objc.IObject, contentRect objc.IObject /* cross-framework: Rect */) Image) {
	d._SharingServiceTransitionImageForShareItemContentRect = f
}

// SetSharingServiceWillShareItems sets the handler for the SharingServiceWillShareItems delegate method.
//
// Invoked when the sharing service will share the specified items.
func (d *SharingServiceDelegate) SetSharingServiceWillShareItems(f func(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */)) {
	d._SharingServiceWillShareItems = f
}

// AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService ISharingService, positioningRect objc.IObject /* cross-framework: Rect */, preferredEdge RectEdge /* not a class type */) View {
	if d._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge != nil {
		return d._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService, positioningRect, preferredEdge)
	}
	var zero View
	return zero
}

// HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge returns true if a handler for AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge has been set.
func (d *SharingServiceDelegate) HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool {
	return d._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge != nil
}

// SharingServiceDidFailToShareItemsError implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceDidFailToShareItemsError(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, error_ objc.IObject /* cross-framework: Error */) {
	if d._SharingServiceDidFailToShareItemsError != nil {
		d._SharingServiceDidFailToShareItemsError(sharingService, items, error_)
	}
}

// HasSharingServiceDidFailToShareItemsError returns true if a handler for SharingServiceDidFailToShareItemsError has been set.
func (d *SharingServiceDelegate) HasSharingServiceDidFailToShareItemsError() bool {
	return d._SharingServiceDidFailToShareItemsError != nil
}

// SharingServiceDidShareItems implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceDidShareItems(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */) {
	if d._SharingServiceDidShareItems != nil {
		d._SharingServiceDidShareItems(sharingService, items)
	}
}

// HasSharingServiceDidShareItems returns true if a handler for SharingServiceDidShareItems has been set.
func (d *SharingServiceDelegate) HasSharingServiceDidShareItems() bool {
	return d._SharingServiceDidShareItems != nil
}

// SharingServiceSourceFrameOnScreenForShareItem implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceSourceFrameOnScreenForShareItem(sharingService ISharingService, item objc.IObject) corefoundation.Rect {
	if d._SharingServiceSourceFrameOnScreenForShareItem != nil {
		return d._SharingServiceSourceFrameOnScreenForShareItem(sharingService, item)
	}
	var zero corefoundation.Rect
	return zero
}

// HasSharingServiceSourceFrameOnScreenForShareItem returns true if a handler for SharingServiceSourceFrameOnScreenForShareItem has been set.
func (d *SharingServiceDelegate) HasSharingServiceSourceFrameOnScreenForShareItem() bool {
	return d._SharingServiceSourceFrameOnScreenForShareItem != nil
}

// SharingServiceSourceWindowForShareItemsSharingContentScope implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */, sharingContentScope SharingContentScope) Window {
	if d._SharingServiceSourceWindowForShareItemsSharingContentScope != nil {
		return d._SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService, items, sharingContentScope)
	}
	var zero Window
	return zero
}

// HasSharingServiceSourceWindowForShareItemsSharingContentScope returns true if a handler for SharingServiceSourceWindowForShareItemsSharingContentScope has been set.
func (d *SharingServiceDelegate) HasSharingServiceSourceWindowForShareItemsSharingContentScope() bool {
	return d._SharingServiceSourceWindowForShareItemsSharingContentScope != nil
}

// SharingServiceTransitionImageForShareItemContentRect implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceTransitionImageForShareItemContentRect(sharingService ISharingService, item objc.IObject, contentRect objc.IObject /* cross-framework: Rect */) Image {
	if d._SharingServiceTransitionImageForShareItemContentRect != nil {
		return d._SharingServiceTransitionImageForShareItemContentRect(sharingService, item, contentRect)
	}
	var zero Image
	return zero
}

// HasSharingServiceTransitionImageForShareItemContentRect returns true if a handler for SharingServiceTransitionImageForShareItemContentRect has been set.
func (d *SharingServiceDelegate) HasSharingServiceTransitionImageForShareItemContentRect() bool {
	return d._SharingServiceTransitionImageForShareItemContentRect != nil
}

// SharingServiceWillShareItems implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceWillShareItems(sharingService ISharingService, items objc.IObject /* cross-framework: NSArray */) {
	if d._SharingServiceWillShareItems != nil {
		d._SharingServiceWillShareItems(sharingService, items)
	}
}

// HasSharingServiceWillShareItems returns true if a handler for SharingServiceWillShareItems has been set.
func (d *SharingServiceDelegate) HasSharingServiceWillShareItems() bool {
	return d._SharingServiceWillShareItems != nil
}
