// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
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
	SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView ISplitView, dividerIndex int) Rect
	HasSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool
	SplitViewCanCollapseSubview(splitView ISplitView, subview IView) bool
	HasSplitViewCanCollapseSubview() bool
	SplitViewConstrainMaxCoordinateOfSubviewAt(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64
	HasSplitViewConstrainMaxCoordinateOfSubviewAt() bool
	SplitViewConstrainMinCoordinateOfSubviewAt(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64
	HasSplitViewConstrainMinCoordinateOfSubviewAt() bool
	SplitViewConstrainSplitPositionOfSubviewAt(splitView ISplitView, proposedPosition float64, dividerIndex int) float64
	HasSplitViewConstrainSplitPositionOfSubviewAt() bool
	SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView ISplitView, proposedEffectiveRect Rect /* not a class type */, drawnRect Rect /* not a class type */, dividerIndex int) Rect
	HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool
	SplitViewResizeSubviewsWithOldSize(splitView ISplitView, oldSize Size /* not a class type */)
	HasSplitViewResizeSubviewsWithOldSize() bool
	SplitViewShouldAdjustSizeOfSubview(splitView ISplitView, view IView) bool
	HasSplitViewShouldAdjustSizeOfSubview() bool
	SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(splitView ISplitView, subview IView, dividerIndex int) bool
	HasSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex() bool
	SplitViewShouldHideDividerAtIndex(splitView ISplitView, dividerIndex int) bool
	HasSplitViewShouldHideDividerAtIndex() bool
	SplitViewDidResizeSubviews(notification foundation.Notification)
	HasSplitViewDidResizeSubviews() bool
	SplitViewWillResizeSubviews(notification foundation.Notification)
	HasSplitViewWillResizeSubviews() bool
}

// SplitViewDelegate is a delegate implementation builder for the PSplitViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SplitViewDelegate struct {
	_SplitViewAdditionalEffectiveRectOfDividerAtIndex func(splitView ISplitView, dividerIndex int) Rect
	_SplitViewCanCollapseSubview func(splitView ISplitView, subview IView) bool
	_SplitViewConstrainMaxCoordinateOfSubviewAt func(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64
	_SplitViewConstrainMinCoordinateOfSubviewAt func(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64
	_SplitViewConstrainSplitPositionOfSubviewAt func(splitView ISplitView, proposedPosition float64, dividerIndex int) float64
	_SplitViewEffectiveRectForDrawnRectOfDividerAtIndex func(splitView ISplitView, proposedEffectiveRect Rect /* not a class type */, drawnRect Rect /* not a class type */, dividerIndex int) Rect
	_SplitViewResizeSubviewsWithOldSize func(splitView ISplitView, oldSize Size /* not a class type */)
	_SplitViewShouldAdjustSizeOfSubview func(splitView ISplitView, view IView) bool
	_SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex func(splitView ISplitView, subview IView, dividerIndex int) bool
	_SplitViewShouldHideDividerAtIndex func(splitView ISplitView, dividerIndex int) bool
	_SplitViewDidResizeSubviews func(notification foundation.Notification)
	_SplitViewWillResizeSubviews func(notification foundation.Notification)
}

// SetSplitViewAdditionalEffectiveRectOfDividerAtIndex sets the handler for the SplitViewAdditionalEffectiveRectOfDividerAtIndex delegate method.
//
// Allows the delegate to return an additional rectangle where mouse clicks can initiate divider dragging.
func (d *SplitViewDelegate) SetSplitViewAdditionalEffectiveRectOfDividerAtIndex(f func(splitView ISplitView, dividerIndex int) Rect) {
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
func (d *SplitViewDelegate) SetSplitViewEffectiveRectForDrawnRectOfDividerAtIndex(f func(splitView ISplitView, proposedEffectiveRect Rect /* not a class type */, drawnRect Rect /* not a class type */, dividerIndex int) Rect) {
	d._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex = f
}

// SetSplitViewResizeSubviewsWithOldSize sets the handler for the SplitViewResizeSubviewsWithOldSize delegate method.
//
// Allows the delegate to specify custom sizing behavior for the subviews of the split view.
func (d *SplitViewDelegate) SetSplitViewResizeSubviewsWithOldSize(f func(splitView ISplitView, oldSize Size /* not a class type */)) {
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
func (d *SplitViewDelegate) SetSplitViewDidResizeSubviews(f func(notification foundation.Notification)) {
	d._SplitViewDidResizeSubviews = f
}

// SetSplitViewWillResizeSubviews sets the handler for the SplitViewWillResizeSubviews delegate method.
//
// Notifies the delegate when the split view is about to resize its subviews.
func (d *SplitViewDelegate) SetSplitViewWillResizeSubviews(f func(notification foundation.Notification)) {
	d._SplitViewWillResizeSubviews = f
}

// SplitViewAdditionalEffectiveRectOfDividerAtIndex implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView ISplitView, dividerIndex int) Rect {
	if d._SplitViewAdditionalEffectiveRectOfDividerAtIndex != nil {
		return d._SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView, dividerIndex)
	}
	var zero Rect
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
func (d *SplitViewDelegate) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView ISplitView, proposedEffectiveRect Rect /* not a class type */, drawnRect Rect /* not a class type */, dividerIndex int) Rect {
	if d._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex != nil {
		return d._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView, proposedEffectiveRect, drawnRect, dividerIndex)
	}
	var zero Rect
	return zero
}

// HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex returns true if a handler for SplitViewEffectiveRectForDrawnRectOfDividerAtIndex has been set.
func (d *SplitViewDelegate) HasSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool {
	return d._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex != nil
}

// SplitViewResizeSubviewsWithOldSize implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewResizeSubviewsWithOldSize(splitView ISplitView, oldSize Size /* not a class type */) {
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
func (d *SplitViewDelegate) SplitViewDidResizeSubviews(notification foundation.Notification) {
	if d._SplitViewDidResizeSubviews != nil {
		d._SplitViewDidResizeSubviews(notification)
	}
}

// HasSplitViewDidResizeSubviews returns true if a handler for SplitViewDidResizeSubviews has been set.
func (d *SplitViewDelegate) HasSplitViewDidResizeSubviews() bool {
	return d._SplitViewDidResizeSubviews != nil
}

// SplitViewWillResizeSubviews implements the PSplitViewDelegate interface.
func (d *SplitViewDelegate) SplitViewWillResizeSubviews(notification foundation.Notification) {
	if d._SplitViewWillResizeSubviews != nil {
		d._SplitViewWillResizeSubviews(notification)
	}
}

// HasSplitViewWillResizeSubviews returns true if a handler for SplitViewWillResizeSubviews has been set.
func (d *SplitViewDelegate) HasSplitViewWillResizeSubviews() bool {
	return d._SplitViewWillResizeSubviews != nil
}
