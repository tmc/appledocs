// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService ISharingService, positioningRect corefoundation.CGRect, preferredEdge RectEdge /* not a class type */) IView
	HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool
	SharingServiceDidFailToShareItemsError(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError)
	HasSharingServiceDidFailToShareItemsError() bool
	SharingServiceDidShareItems(sharingService ISharingService, items foundation.foundation.INSArray)
	HasSharingServiceDidShareItems() bool
	SharingServiceSourceFrameOnScreenForShareItem(sharingService ISharingService, item objectivec.IObject) corefoundation.CGRect
	HasSharingServiceSourceFrameOnScreenForShareItem() bool
	SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService ISharingService, items foundation.foundation.INSArray, sharingContentScope SharingContentScope) IWindow
	HasSharingServiceSourceWindowForShareItemsSharingContentScope() bool
	SharingServiceTransitionImageForShareItemContentRect(sharingService ISharingService, item objectivec.IObject, contentRect corefoundation.CGRect) IImage
	HasSharingServiceTransitionImageForShareItemContentRect() bool
	SharingServiceWillShareItems(sharingService ISharingService, items foundation.foundation.INSArray)
	HasSharingServiceWillShareItems() bool
}

// SharingServiceDelegate is a delegate implementation builder for the PSharingServiceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SharingServiceDelegate struct {
	_AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge func(sharingService ISharingService, positioningRect corefoundation.CGRect, preferredEdge RectEdge /* not a class type */) IView
	_SharingServiceDidFailToShareItemsError func(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError)
	_SharingServiceDidShareItems func(sharingService ISharingService, items foundation.foundation.INSArray)
	_SharingServiceSourceFrameOnScreenForShareItem func(sharingService ISharingService, item objectivec.IObject) corefoundation.CGRect
	_SharingServiceSourceWindowForShareItemsSharingContentScope func(sharingService ISharingService, items foundation.foundation.INSArray, sharingContentScope SharingContentScope) IWindow
	_SharingServiceTransitionImageForShareItemContentRect func(sharingService ISharingService, item objectivec.IObject, contentRect corefoundation.CGRect) IImage
	_SharingServiceWillShareItems func(sharingService ISharingService, items foundation.foundation.INSArray)
}

// SetAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge sets the handler for the AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge delegate method.
func (d *SharingServiceDelegate) SetAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(f func(sharingService ISharingService, positioningRect corefoundation.CGRect, preferredEdge RectEdge /* not a class type */) IView) {
	d._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge = f
}

// SetSharingServiceDidFailToShareItemsError sets the handler for the SharingServiceDidFailToShareItemsError delegate method.
//
// Invoked when the sharing service encountered an error when sharing items.
func (d *SharingServiceDelegate) SetSharingServiceDidFailToShareItemsError(f func(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError)) {
	d._SharingServiceDidFailToShareItemsError = f
}

// SetSharingServiceDidShareItems sets the handler for the SharingServiceDidShareItems delegate method.
//
// Invoked when the sharing service has finished sharing the items.
func (d *SharingServiceDelegate) SetSharingServiceDidShareItems(f func(sharingService ISharingService, items foundation.foundation.INSArray)) {
	d._SharingServiceDidShareItems = f
}

// SetSharingServiceSourceFrameOnScreenForShareItem sets the handler for the SharingServiceSourceFrameOnScreenForShareItem delegate method.
//
// Invoked when the sharing service is performed and the sharing window is displayed, to present a transition between the original items and the sharing window.
func (d *SharingServiceDelegate) SetSharingServiceSourceFrameOnScreenForShareItem(f func(sharingService ISharingService, item objectivec.IObject) corefoundation.CGRect) {
	d._SharingServiceSourceFrameOnScreenForShareItem = f
}

// SetSharingServiceSourceWindowForShareItemsSharingContentScope sets the handler for the SharingServiceSourceWindowForShareItemsSharingContentScope delegate method.
//
// Returns the window that contained the share items.
func (d *SharingServiceDelegate) SetSharingServiceSourceWindowForShareItemsSharingContentScope(f func(sharingService ISharingService, items foundation.foundation.INSArray, sharingContentScope SharingContentScope) IWindow) {
	d._SharingServiceSourceWindowForShareItemsSharingContentScope = f
}

// SetSharingServiceTransitionImageForShareItemContentRect sets the handler for the SharingServiceTransitionImageForShareItemContentRect delegate method.
//
// Invoked to allow returning a custom transition image when sharing an item.
func (d *SharingServiceDelegate) SetSharingServiceTransitionImageForShareItemContentRect(f func(sharingService ISharingService, item objectivec.IObject, contentRect corefoundation.CGRect) IImage) {
	d._SharingServiceTransitionImageForShareItemContentRect = f
}

// SetSharingServiceWillShareItems sets the handler for the SharingServiceWillShareItems delegate method.
//
// Invoked when the sharing service will share the specified items.
func (d *SharingServiceDelegate) SetSharingServiceWillShareItems(f func(sharingService ISharingService, items foundation.foundation.INSArray)) {
	d._SharingServiceWillShareItems = f
}

// AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService ISharingService, positioningRect corefoundation.CGRect, preferredEdge RectEdge /* not a class type */) IView {
	if d._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge != nil {
		return d._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService, positioningRect, preferredEdge)
	}
	var zero IView
	return zero
}

// HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge returns true if a handler for AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge has been set.
func (d *SharingServiceDelegate) HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool {
	return d._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge != nil
}

// SharingServiceDidFailToShareItemsError implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceDidFailToShareItemsError(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError) {
	if d._SharingServiceDidFailToShareItemsError != nil {
		d._SharingServiceDidFailToShareItemsError(sharingService, items, error_)
	}
}

// HasSharingServiceDidFailToShareItemsError returns true if a handler for SharingServiceDidFailToShareItemsError has been set.
func (d *SharingServiceDelegate) HasSharingServiceDidFailToShareItemsError() bool {
	return d._SharingServiceDidFailToShareItemsError != nil
}

// SharingServiceDidShareItems implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceDidShareItems(sharingService ISharingService, items foundation.foundation.INSArray) {
	if d._SharingServiceDidShareItems != nil {
		d._SharingServiceDidShareItems(sharingService, items)
	}
}

// HasSharingServiceDidShareItems returns true if a handler for SharingServiceDidShareItems has been set.
func (d *SharingServiceDelegate) HasSharingServiceDidShareItems() bool {
	return d._SharingServiceDidShareItems != nil
}

// SharingServiceSourceFrameOnScreenForShareItem implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceSourceFrameOnScreenForShareItem(sharingService ISharingService, item objectivec.IObject) corefoundation.CGRect {
	if d._SharingServiceSourceFrameOnScreenForShareItem != nil {
		return d._SharingServiceSourceFrameOnScreenForShareItem(sharingService, item)
	}
	var zero corefoundation.CGRect
	return zero
}

// HasSharingServiceSourceFrameOnScreenForShareItem returns true if a handler for SharingServiceSourceFrameOnScreenForShareItem has been set.
func (d *SharingServiceDelegate) HasSharingServiceSourceFrameOnScreenForShareItem() bool {
	return d._SharingServiceSourceFrameOnScreenForShareItem != nil
}

// SharingServiceSourceWindowForShareItemsSharingContentScope implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService ISharingService, items foundation.foundation.INSArray, sharingContentScope SharingContentScope) IWindow {
	if d._SharingServiceSourceWindowForShareItemsSharingContentScope != nil {
		return d._SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService, items, sharingContentScope)
	}
	var zero IWindow
	return zero
}

// HasSharingServiceSourceWindowForShareItemsSharingContentScope returns true if a handler for SharingServiceSourceWindowForShareItemsSharingContentScope has been set.
func (d *SharingServiceDelegate) HasSharingServiceSourceWindowForShareItemsSharingContentScope() bool {
	return d._SharingServiceSourceWindowForShareItemsSharingContentScope != nil
}

// SharingServiceTransitionImageForShareItemContentRect implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceTransitionImageForShareItemContentRect(sharingService ISharingService, item objectivec.IObject, contentRect corefoundation.CGRect) IImage {
	if d._SharingServiceTransitionImageForShareItemContentRect != nil {
		return d._SharingServiceTransitionImageForShareItemContentRect(sharingService, item, contentRect)
	}
	var zero IImage
	return zero
}

// HasSharingServiceTransitionImageForShareItemContentRect returns true if a handler for SharingServiceTransitionImageForShareItemContentRect has been set.
func (d *SharingServiceDelegate) HasSharingServiceTransitionImageForShareItemContentRect() bool {
	return d._SharingServiceTransitionImageForShareItemContentRect != nil
}

// SharingServiceWillShareItems implements the PSharingServiceDelegate interface.
func (d *SharingServiceDelegate) SharingServiceWillShareItems(sharingService ISharingService, items foundation.foundation.INSArray) {
	if d._SharingServiceWillShareItems != nil {
		d._SharingServiceWillShareItems(sharingService, items)
	}
}

// HasSharingServiceWillShareItems returns true if a handler for SharingServiceWillShareItems has been set.
func (d *SharingServiceDelegate) HasSharingServiceWillShareItems() bool {
	return d._SharingServiceWillShareItems != nil
}

// SharingServiceDelegateObject wraps an existing Objective-C object that conforms to the PSharingServiceDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type SharingServiceDelegateObject struct {
	objectivec.Object
}

// NewSharingServiceDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSSharingServiceDelegate protocol.
func NewSharingServiceDelegateObject(obj objectivec.Object) *SharingServiceDelegateObject {
	return &SharingServiceDelegateObject{obj}
}

// Make sure SharingServiceDelegateObject implements PSharingServiceDelegate.
var _ PSharingServiceDelegate = (*SharingServiceDelegateObject)(nil)

// AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge implements the PSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServiceDelegateObject) AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService ISharingService, positioningRect corefoundation.CGRect, preferredEdge RectEdge /* not a class type */) IView {
	return objc.Send[IView](o.ID, objc.Sel("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"), sharingService, positioningRect, preferredEdge)
}

// HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge returns true; this is a placeholder for optional method checks.
func (o *SharingServiceDelegateObject) HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServiceDidFailToShareItemsError implements the PSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServiceDelegateObject) SharingServiceDidFailToShareItemsError(sharingService ISharingService, items foundation.foundation.INSArray, error_ foundation.foundation.INSError) {
	objc.Send[objc.ID](o.ID, objc.Sel("sharingService:didFailToShareItems:error:"), sharingService, items, error_)
}

// HasSharingServiceDidFailToShareItemsError returns true; this is a placeholder for optional method checks.
func (o *SharingServiceDelegateObject) HasSharingServiceDidFailToShareItemsError() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServiceDidShareItems implements the PSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServiceDelegateObject) SharingServiceDidShareItems(sharingService ISharingService, items foundation.foundation.INSArray) {
	objc.Send[objc.ID](o.ID, objc.Sel("sharingService:didShareItems:"), sharingService, items)
}

// HasSharingServiceDidShareItems returns true; this is a placeholder for optional method checks.
func (o *SharingServiceDelegateObject) HasSharingServiceDidShareItems() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServiceSourceFrameOnScreenForShareItem implements the PSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServiceDelegateObject) SharingServiceSourceFrameOnScreenForShareItem(sharingService ISharingService, item objectivec.IObject) corefoundation.CGRect {
	return objc.Send[corefoundation.CGRect](o.ID, objc.Sel("sharingService:sourceFrameOnScreenForShareItem:"), sharingService, item)
}

// HasSharingServiceSourceFrameOnScreenForShareItem returns true; this is a placeholder for optional method checks.
func (o *SharingServiceDelegateObject) HasSharingServiceSourceFrameOnScreenForShareItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServiceSourceWindowForShareItemsSharingContentScope implements the PSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServiceDelegateObject) SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService ISharingService, items foundation.foundation.INSArray, sharingContentScope SharingContentScope) IWindow {
	return objc.Send[IWindow](o.ID, objc.Sel("sharingService:sourceWindowForShareItems:sharingContentScope:"), sharingService, items, sharingContentScope)
}

// HasSharingServiceSourceWindowForShareItemsSharingContentScope returns true; this is a placeholder for optional method checks.
func (o *SharingServiceDelegateObject) HasSharingServiceSourceWindowForShareItemsSharingContentScope() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServiceTransitionImageForShareItemContentRect implements the PSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServiceDelegateObject) SharingServiceTransitionImageForShareItemContentRect(sharingService ISharingService, item objectivec.IObject, contentRect corefoundation.CGRect) IImage {
	return objc.Send[IImage](o.ID, objc.Sel("sharingService:transitionImageForShareItem:contentRect:"), sharingService, item, contentRect)
}

// HasSharingServiceTransitionImageForShareItemContentRect returns true; this is a placeholder for optional method checks.
func (o *SharingServiceDelegateObject) HasSharingServiceTransitionImageForShareItemContentRect() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SharingServiceWillShareItems implements the PSharingServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SharingServiceDelegateObject) SharingServiceWillShareItems(sharingService ISharingService, items foundation.foundation.INSArray) {
	objc.Send[objc.ID](o.ID, objc.Sel("sharingService:willShareItems:"), sharingService, items)
}

// HasSharingServiceWillShareItems returns true; this is a placeholder for optional method checks.
func (o *SharingServiceDelegateObject) HasSharingServiceWillShareItems() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
