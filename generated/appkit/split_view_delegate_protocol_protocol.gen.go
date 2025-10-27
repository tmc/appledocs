// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PSplitViewDelegate is the NSSplitViewDelegate protocol interface.
//
// A set of optional methods that a delegate of a split view implements.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSplitViewDelegate
type PSplitViewDelegate interface {
	// Optional methods
	SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView ISplitView, dividerIndex int) corefoundation.CGRect
	HasSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool
	SplitViewCanCollapseSubview(splitView ISplitView, subview IView) bool
	HasSplitViewCanCollapseSubview() bool
	SplitViewConstrainMaxCoordinateOfSubviewAt(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64
	HasSplitViewConstrainMaxCoordinateOfSubviewAt() bool
	SplitViewConstrainMinCoordinateOfSubviewAt(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64
	HasSplitViewConstrainMinCoordinateOfSubviewAt() bool
	SplitViewConstrainSplitPositionOfSubviewAt(splitView ISplitView, proposedPosition float64, dividerIndex int) float64
	HasSplitViewConstrainSplitPositionOfSubviewAt() bool
	SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView ISplitView, proposedEffectiveRect corefoundation.CGRect, drawnRect corefoundation.CGRect, dividerIndex int) corefoundation.CGRect
	HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool
	SplitViewResizeSubviewsWithOldSize(splitView ISplitView, oldSize corefoundation.CGSize)
	HasSplitViewResizeSubviewsWithOldSize() bool
	SplitViewShouldAdjustSizeOfSubview(splitView ISplitView, view IView) bool
	HasSplitViewShouldAdjustSizeOfSubview() bool
	SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(splitView ISplitView, subview IView, dividerIndex int) bool
	HasSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex() bool
	SplitViewShouldHideDividerAtIndex(splitView ISplitView, dividerIndex int) bool
	HasSplitViewShouldHideDividerAtIndex() bool
	SplitViewDidResizeSubviews(notification foundation.foundation.INSNotification)
	HasSplitViewDidResizeSubviews() bool
	SplitViewWillResizeSubviews(notification foundation.foundation.INSNotification)
	HasSplitViewWillResizeSubviews() bool
}

// SplitViewDelegate is a delegate implementation builder for the PSplitViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SplitViewDelegate struct {
	_SplitViewAdditionalEffectiveRectOfDividerAtIndex func(splitView ISplitView, dividerIndex int) corefoundation.CGRect
	_SplitViewCanCollapseSubview func(splitView ISplitView, subview IView) bool
	_SplitViewConstrainMaxCoordinateOfSubviewAt func(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64
	_SplitViewConstrainMinCoordinateOfSubviewAt func(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64
	_SplitViewConstrainSplitPositionOfSubviewAt func(splitView ISplitView, proposedPosition float64, dividerIndex int) float64
	_SplitViewEffectiveRectForDrawnRectOfDividerAtIndex func(splitView ISplitView, proposedEffectiveRect corefoundation.CGRect, drawnRect corefoundation.CGRect, dividerIndex int) corefoundation.CGRect
	_SplitViewResizeSubviewsWithOldSize func(splitView ISplitView, oldSize corefoundation.CGSize)
	_SplitViewShouldAdjustSizeOfSubview func(splitView ISplitView, view IView) bool
	_SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex func(splitView ISplitView, subview IView, dividerIndex int) bool
	_SplitViewShouldHideDividerAtIndex func(splitView ISplitView, dividerIndex int) bool
	_SplitViewDidResizeSubviews func(notification foundation.foundation.INSNotification)
	_SplitViewWillResizeSubviews func(notification foundation.foundation.INSNotification)
}

// SetSplitViewAdditionalEffectiveRectOfDividerAtIndex sets the handler for the SplitViewAdditionalEffectiveRectOfDividerAtIndex delegate method.
//
// Allows the delegate to return an additional rectangle where mouse clicks can initiate divider dragging.
func (d *SplitViewDelegate) SetSplitViewAdditionalEffectiveRectOfDividerAtIndex(f func(splitView ISplitView, dividerIndex int) corefoundation.CGRect) {
	d._SplitViewAdditionalEffectiveRectOfDividerAtIndex = f
}

// SetSplitViewCanCollapseSubview sets the handler for the SplitViewCanCollapseSubview delegate method.
//
// Allows the delegate to determine whether the user can collapse and expand the specified subview.
func (d *SplitViewDelegate) SetSplitViewCanCollapseSubview(f func(splitView ISplitView, subview IView) bool) {
	d._SplitViewCanCollapseSubview = f
}

// SetSplitViewConstrainMaxCoordinateOfSubviewAt sets the handler for the SplitViewConstrainMaxCoordinateOfSubviewAt delegate method.
//
// Allows the delegate to constrain the maximum coordinate limit of a divider when the user drags it.
func (d *SplitViewDelegate) SetSplitViewConstrainMaxCoordinateOfSubviewAt(f func(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64) {
	d._SplitViewConstrainMaxCoordinateOfSubviewAt = f
}

// SetSplitViewConstrainMinCoordinateOfSubviewAt sets the handler for the SplitViewConstrainMinCoordinateOfSubviewAt delegate method.
//
// Allows the delegate to constrain the minimum coordinate limit of a divider when the user drags it.
func (d *SplitViewDelegate) SetSplitViewConstrainMinCoordinateOfSubviewAt(f func(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64) {
	d._SplitViewConstrainMinCoordinateOfSubviewAt = f
}

// SetSplitViewConstrainSplitPositionOfSubviewAt sets the handler for the SplitViewConstrainSplitPositionOfSubviewAt delegate method.
//
// Allows the delegate to constrain the divider to certain positions.
func (d *SplitViewDelegate) SetSplitViewConstrainSplitPositionOfSubviewAt(f func(splitView ISplitView, proposedPosition float64, dividerIndex int) float64) {
	d._SplitViewConstrainSplitPositionOfSubviewAt = f
}

// SetSplitViewEffectiveRectForDrawnRectOfDividerAtIndex sets the handler for the SplitViewEffectiveRectForDrawnRectOfDividerAtIndex delegate method.
//
// Allows the delegate to modify the rectangle where mouse clicks initiate divider dragging.
func (d *SplitViewDelegate) SetSplitViewEffectiveRectForDrawnRectOfDividerAtIndex(f func(splitView ISplitView, proposedEffectiveRect corefoundation.CGRect, drawnRect corefoundation.CGRect, dividerIndex int) corefoundation.CGRect) {
	d._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex = f
}

// SetSplitViewResizeSubviewsWithOldSize sets the handler for the SplitViewResizeSubviewsWithOldSize delegate method.
//
// Allows the delegate to specify custom sizing behavior for the subviews of the split view.
func (d *SplitViewDelegate) SetSplitViewResizeSubviewsWithOldSize(f func(splitView ISplitView, oldSize corefoundation.CGSize)) {
	d._SplitViewResizeSubviewsWithOldSize = f
}

// SetSplitViewShouldAdjustSizeOfSubview sets the handler for the SplitViewShouldAdjustSizeOfSubview delegate method.
//
// Allows the delegate to specify whether to resize the subview.
func (d *SplitViewDelegate) SetSplitViewShouldAdjustSizeOfSubview(f func(splitView ISplitView, view IView) bool) {
	d._SplitViewShouldAdjustSizeOfSubview = f
}

// SetSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex sets the handler for the SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex delegate method.
//
// Allows a delegate to determine if a subview collapses in response to a double click.
func (d *SplitViewDelegate) SetSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(f func(splitView ISplitView, subview IView, dividerIndex int) bool) {
	d._SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex = f
}

// SetSplitViewShouldHideDividerAtIndex sets the handler for the SplitViewShouldHideDividerAtIndex delegate method.
//
// Allows the delegate to determine whether the user can drag a divider or adjust it off the edge of the split view.
func (d *SplitViewDelegate) SetSplitViewShouldHideDividerAtIndex(f func(splitView ISplitView, dividerIndex int) bool) {
	d._SplitViewShouldHideDividerAtIndex = f
}

// SetSplitViewDidResizeSubviews sets the handler for the SplitViewDidResizeSubviews delegate method.
//
// Notifies the delegate when the split view resizes its subviews.
func (d *SplitViewDelegate) SetSplitViewDidResizeSubviews(f func(notification foundation.foundation.INSNotification)) {
	d._SplitViewDidResizeSubviews = f
}

// SetSplitViewWillResizeSubviews sets the handler for the SplitViewWillResizeSubviews delegate method.
//
// Notifies the delegate when the split view is about to resize its subviews.
func (d *SplitViewDelegate) SetSplitViewWillResizeSubviews(f func(notification foundation.foundation.INSNotification)) {
	d._SplitViewWillResizeSubviews = f
}

// SplitViewAdditionalEffectiveRectOfDividerAtIndex implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView ISplitView, dividerIndex int) corefoundation.CGRect {
	if d._SplitViewAdditionalEffectiveRectOfDividerAtIndex != nil {
		return d._SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView, dividerIndex)
	}
	var zero corefoundation.CGRect
	return zero
}

// HasSplitViewAdditionalEffectiveRectOfDividerAtIndex returns true if a handler for SplitViewAdditionalEffectiveRectOfDividerAtIndex has been set.
func (d *SplitViewDelegate) HasSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool {
	return d._SplitViewAdditionalEffectiveRectOfDividerAtIndex != nil
}

// SplitViewCanCollapseSubview implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewCanCollapseSubview(splitView ISplitView, subview IView) bool {
	if d._SplitViewCanCollapseSubview != nil {
		return d._SplitViewCanCollapseSubview(splitView, subview)
	}
	var zero bool
	return zero
}

// HasSplitViewCanCollapseSubview returns true if a handler for SplitViewCanCollapseSubview has been set.
func (d *SplitViewDelegate) HasSplitViewCanCollapseSubview() bool {
	return d._SplitViewCanCollapseSubview != nil
}

// SplitViewConstrainMaxCoordinateOfSubviewAt implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewConstrainMaxCoordinateOfSubviewAt(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	if d._SplitViewConstrainMaxCoordinateOfSubviewAt != nil {
		return d._SplitViewConstrainMaxCoordinateOfSubviewAt(splitView, proposedMaximumPosition, dividerIndex)
	}
	var zero float64
	return zero
}

// HasSplitViewConstrainMaxCoordinateOfSubviewAt returns true if a handler for SplitViewConstrainMaxCoordinateOfSubviewAt has been set.
func (d *SplitViewDelegate) HasSplitViewConstrainMaxCoordinateOfSubviewAt() bool {
	return d._SplitViewConstrainMaxCoordinateOfSubviewAt != nil
}

// SplitViewConstrainMinCoordinateOfSubviewAt implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewConstrainMinCoordinateOfSubviewAt(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	if d._SplitViewConstrainMinCoordinateOfSubviewAt != nil {
		return d._SplitViewConstrainMinCoordinateOfSubviewAt(splitView, proposedMinimumPosition, dividerIndex)
	}
	var zero float64
	return zero
}

// HasSplitViewConstrainMinCoordinateOfSubviewAt returns true if a handler for SplitViewConstrainMinCoordinateOfSubviewAt has been set.
func (d *SplitViewDelegate) HasSplitViewConstrainMinCoordinateOfSubviewAt() bool {
	return d._SplitViewConstrainMinCoordinateOfSubviewAt != nil
}

// SplitViewConstrainSplitPositionOfSubviewAt implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewConstrainSplitPositionOfSubviewAt(splitView ISplitView, proposedPosition float64, dividerIndex int) float64 {
	if d._SplitViewConstrainSplitPositionOfSubviewAt != nil {
		return d._SplitViewConstrainSplitPositionOfSubviewAt(splitView, proposedPosition, dividerIndex)
	}
	var zero float64
	return zero
}

// HasSplitViewConstrainSplitPositionOfSubviewAt returns true if a handler for SplitViewConstrainSplitPositionOfSubviewAt has been set.
func (d *SplitViewDelegate) HasSplitViewConstrainSplitPositionOfSubviewAt() bool {
	return d._SplitViewConstrainSplitPositionOfSubviewAt != nil
}

// SplitViewEffectiveRectForDrawnRectOfDividerAtIndex implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView ISplitView, proposedEffectiveRect corefoundation.CGRect, drawnRect corefoundation.CGRect, dividerIndex int) corefoundation.CGRect {
	if d._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex != nil {
		return d._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView, proposedEffectiveRect, drawnRect, dividerIndex)
	}
	var zero corefoundation.CGRect
	return zero
}

// HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex returns true if a handler for SplitViewEffectiveRectForDrawnRectOfDividerAtIndex has been set.
func (d *SplitViewDelegate) HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool {
	return d._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex != nil
}

// SplitViewResizeSubviewsWithOldSize implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewResizeSubviewsWithOldSize(splitView ISplitView, oldSize corefoundation.CGSize) {
	if d._SplitViewResizeSubviewsWithOldSize != nil {
		d._SplitViewResizeSubviewsWithOldSize(splitView, oldSize)
	}
}

// HasSplitViewResizeSubviewsWithOldSize returns true if a handler for SplitViewResizeSubviewsWithOldSize has been set.
func (d *SplitViewDelegate) HasSplitViewResizeSubviewsWithOldSize() bool {
	return d._SplitViewResizeSubviewsWithOldSize != nil
}

// SplitViewShouldAdjustSizeOfSubview implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewShouldAdjustSizeOfSubview(splitView ISplitView, view IView) bool {
	if d._SplitViewShouldAdjustSizeOfSubview != nil {
		return d._SplitViewShouldAdjustSizeOfSubview(splitView, view)
	}
	var zero bool
	return zero
}

// HasSplitViewShouldAdjustSizeOfSubview returns true if a handler for SplitViewShouldAdjustSizeOfSubview has been set.
func (d *SplitViewDelegate) HasSplitViewShouldAdjustSizeOfSubview() bool {
	return d._SplitViewShouldAdjustSizeOfSubview != nil
}

// SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(splitView ISplitView, subview IView, dividerIndex int) bool {
	if d._SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex != nil {
		return d._SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(splitView, subview, dividerIndex)
	}
	var zero bool
	return zero
}

// HasSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex returns true if a handler for SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex has been set.
func (d *SplitViewDelegate) HasSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex() bool {
	return d._SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex != nil
}

// SplitViewShouldHideDividerAtIndex implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewShouldHideDividerAtIndex(splitView ISplitView, dividerIndex int) bool {
	if d._SplitViewShouldHideDividerAtIndex != nil {
		return d._SplitViewShouldHideDividerAtIndex(splitView, dividerIndex)
	}
	var zero bool
	return zero
}

// HasSplitViewShouldHideDividerAtIndex returns true if a handler for SplitViewShouldHideDividerAtIndex has been set.
func (d *SplitViewDelegate) HasSplitViewShouldHideDividerAtIndex() bool {
	return d._SplitViewShouldHideDividerAtIndex != nil
}

// SplitViewDidResizeSubviews implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewDidResizeSubviews(notification foundation.foundation.INSNotification) {
	if d._SplitViewDidResizeSubviews != nil {
		d._SplitViewDidResizeSubviews(notification)
	}
}

// HasSplitViewDidResizeSubviews returns true if a handler for SplitViewDidResizeSubviews has been set.
func (d *SplitViewDelegate) HasSplitViewDidResizeSubviews() bool {
	return d._SplitViewDidResizeSubviews != nil
}

// SplitViewWillResizeSubviews implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewWillResizeSubviews(notification foundation.foundation.INSNotification) {
	if d._SplitViewWillResizeSubviews != nil {
		d._SplitViewWillResizeSubviews(notification)
	}
}

// HasSplitViewWillResizeSubviews returns true if a handler for SplitViewWillResizeSubviews has been set.
func (d *SplitViewDelegate) HasSplitViewWillResizeSubviews() bool {
	return d._SplitViewWillResizeSubviews != nil
}

// SplitViewDelegateObject wraps an existing Objective-C object that conforms to the PSplitViewDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type SplitViewDelegateObject struct {
	objectivec.Object
}

// NewSplitViewDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSSplitViewDelegate protocol.
func NewSplitViewDelegateObject(obj objectivec.Object) *SplitViewDelegateObject {
	return &SplitViewDelegateObject{obj}
}

// Make sure SplitViewDelegateObject implements PSplitViewDelegate.
var _ PSplitViewDelegate = (*SplitViewDelegateObject)(nil)

// SplitViewAdditionalEffectiveRectOfDividerAtIndex implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView ISplitView, dividerIndex int) corefoundation.CGRect {
	return objc.Send[corefoundation.CGRect](o.ID, objc.Sel("splitView:additionalEffectiveRectOfDividerAtIndex:"), splitView, dividerIndex)
}

// HasSplitViewAdditionalEffectiveRectOfDividerAtIndex returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewCanCollapseSubview implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewCanCollapseSubview(splitView ISplitView, subview IView) bool {
	return objc.Send[bool](o.ID, objc.Sel("splitView:canCollapseSubview:"), splitView, subview)
}

// HasSplitViewCanCollapseSubview returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewCanCollapseSubview() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewConstrainMaxCoordinateOfSubviewAt implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewConstrainMaxCoordinateOfSubviewAt(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	return objc.Send[float64](o.ID, objc.Sel("splitView:constrainMaxCoordinate:ofSubviewAt:"), splitView, proposedMaximumPosition, dividerIndex)
}

// HasSplitViewConstrainMaxCoordinateOfSubviewAt returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewConstrainMaxCoordinateOfSubviewAt() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewConstrainMinCoordinateOfSubviewAt implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewConstrainMinCoordinateOfSubviewAt(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	return objc.Send[float64](o.ID, objc.Sel("splitView:constrainMinCoordinate:ofSubviewAt:"), splitView, proposedMinimumPosition, dividerIndex)
}

// HasSplitViewConstrainMinCoordinateOfSubviewAt returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewConstrainMinCoordinateOfSubviewAt() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewConstrainSplitPositionOfSubviewAt implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewConstrainSplitPositionOfSubviewAt(splitView ISplitView, proposedPosition float64, dividerIndex int) float64 {
	return objc.Send[float64](o.ID, objc.Sel("splitView:constrainSplitPosition:ofSubviewAt:"), splitView, proposedPosition, dividerIndex)
}

// HasSplitViewConstrainSplitPositionOfSubviewAt returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewConstrainSplitPositionOfSubviewAt() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewEffectiveRectForDrawnRectOfDividerAtIndex implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView ISplitView, proposedEffectiveRect corefoundation.CGRect, drawnRect corefoundation.CGRect, dividerIndex int) corefoundation.CGRect {
	return objc.Send[corefoundation.CGRect](o.ID, objc.Sel("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"), splitView, proposedEffectiveRect, drawnRect, dividerIndex)
}

// HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewResizeSubviewsWithOldSize implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewResizeSubviewsWithOldSize(splitView ISplitView, oldSize corefoundation.CGSize) {
	objc.Send[objc.ID](o.ID, objc.Sel("splitView:resizeSubviewsWithOldSize:"), splitView, oldSize)
}

// HasSplitViewResizeSubviewsWithOldSize returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewResizeSubviewsWithOldSize() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewShouldAdjustSizeOfSubview implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewShouldAdjustSizeOfSubview(splitView ISplitView, view IView) bool {
	return objc.Send[bool](o.ID, objc.Sel("splitView:shouldAdjustSizeOfSubview:"), splitView, view)
}

// HasSplitViewShouldAdjustSizeOfSubview returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewShouldAdjustSizeOfSubview() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(splitView ISplitView, subview IView, dividerIndex int) bool {
	return objc.Send[bool](o.ID, objc.Sel("splitView:shouldCollapseSubview:forDoubleClickOnDividerAtIndex:"), splitView, subview, dividerIndex)
}

// HasSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewShouldHideDividerAtIndex implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewShouldHideDividerAtIndex(splitView ISplitView, dividerIndex int) bool {
	return objc.Send[bool](o.ID, objc.Sel("splitView:shouldHideDividerAtIndex:"), splitView, dividerIndex)
}

// HasSplitViewShouldHideDividerAtIndex returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewShouldHideDividerAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewDidResizeSubviews implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewDidResizeSubviews(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("splitViewDidResizeSubviews:"), notification)
}

// HasSplitViewDidResizeSubviews returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewDidResizeSubviews() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SplitViewWillResizeSubviews implements the PSplitViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SplitViewDelegateObject) SplitViewWillResizeSubviews(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("splitViewWillResizeSubviews:"), notification)
}

// HasSplitViewWillResizeSubviews returns true; this is a placeholder for optional method checks.
func (o *SplitViewDelegateObject) HasSplitViewWillResizeSubviews() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
