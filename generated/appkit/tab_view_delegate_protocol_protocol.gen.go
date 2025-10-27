// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTabViewDelegate is the NSTabViewDelegate protocol interface.
//
// The   protocol defines the optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTabViewDelegate
type PTabViewDelegate interface {
	// Optional methods
	TabViewDidSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem)
	HasTabViewDidSelectTabViewItem() bool
	TabViewShouldSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) bool
	HasTabViewShouldSelectTabViewItem() bool
	TabViewWillSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem)
	HasTabViewWillSelectTabViewItem() bool
	TabViewDidChangeNumberOfTabViewItems(tabView ITabView)
	HasTabViewDidChangeNumberOfTabViewItems() bool
}

// TabViewDelegate is a delegate implementation builder for the PTabViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TabViewDelegate struct {
	_TabViewDidSelectTabViewItem func(tabView ITabView, tabViewItem ITabViewItem)
	_TabViewShouldSelectTabViewItem func(tabView ITabView, tabViewItem ITabViewItem) bool
	_TabViewWillSelectTabViewItem func(tabView ITabView, tabViewItem ITabViewItem)
	_TabViewDidChangeNumberOfTabViewItems func(tabView ITabView)
}

// SetTabViewDidSelectTabViewItem sets the handler for the TabViewDidSelectTabViewItem delegate method.
//
// Informs the delegate that   has selected  .
func (d *TabViewDelegate) SetTabViewDidSelectTabViewItem(f func(tabView ITabView, tabViewItem ITabViewItem)) {
	d._TabViewDidSelectTabViewItem = f
}

// SetTabViewShouldSelectTabViewItem sets the handler for the TabViewShouldSelectTabViewItem delegate method.
//
// Invoked just before   in   is selected.
func (d *TabViewDelegate) SetTabViewShouldSelectTabViewItem(f func(tabView ITabView, tabViewItem ITabViewItem) bool) {
	d._TabViewShouldSelectTabViewItem = f
}

// SetTabViewWillSelectTabViewItem sets the handler for the TabViewWillSelectTabViewItem delegate method.
//
// Informs the delegate that   is about to select  .
func (d *TabViewDelegate) SetTabViewWillSelectTabViewItem(f func(tabView ITabView, tabViewItem ITabViewItem)) {
	d._TabViewWillSelectTabViewItem = f
}

// SetTabViewDidChangeNumberOfTabViewItems sets the handler for the TabViewDidChangeNumberOfTabViewItems delegate method.
//
// Informs the delegate that the number of tab view items in   has changed.
func (d *TabViewDelegate) SetTabViewDidChangeNumberOfTabViewItems(f func(tabView ITabView)) {
	d._TabViewDidChangeNumberOfTabViewItems = f
}

// TabViewDidSelectTabViewItem implements the PTabViewDelegate interface.
func (d *TabViewDelegate) TabViewDidSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	if d._TabViewDidSelectTabViewItem != nil {
		d._TabViewDidSelectTabViewItem(tabView, tabViewItem)
	}
}

// HasTabViewDidSelectTabViewItem returns true if a handler for TabViewDidSelectTabViewItem has been set.
func (d *TabViewDelegate) HasTabViewDidSelectTabViewItem() bool {
	return d._TabViewDidSelectTabViewItem != nil
}

// TabViewShouldSelectTabViewItem implements the PTabViewDelegate interface.
func (d *TabViewDelegate) TabViewShouldSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) bool {
	if d._TabViewShouldSelectTabViewItem != nil {
		return d._TabViewShouldSelectTabViewItem(tabView, tabViewItem)
	}
	var zero bool
	return zero
}

// HasTabViewShouldSelectTabViewItem returns true if a handler for TabViewShouldSelectTabViewItem has been set.
func (d *TabViewDelegate) HasTabViewShouldSelectTabViewItem() bool {
	return d._TabViewShouldSelectTabViewItem != nil
}

// TabViewWillSelectTabViewItem implements the PTabViewDelegate interface.
func (d *TabViewDelegate) TabViewWillSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	if d._TabViewWillSelectTabViewItem != nil {
		d._TabViewWillSelectTabViewItem(tabView, tabViewItem)
	}
}

// HasTabViewWillSelectTabViewItem returns true if a handler for TabViewWillSelectTabViewItem has been set.
func (d *TabViewDelegate) HasTabViewWillSelectTabViewItem() bool {
	return d._TabViewWillSelectTabViewItem != nil
}

// TabViewDidChangeNumberOfTabViewItems implements the PTabViewDelegate interface.
func (d *TabViewDelegate) TabViewDidChangeNumberOfTabViewItems(tabView ITabView) {
	if d._TabViewDidChangeNumberOfTabViewItems != nil {
		d._TabViewDidChangeNumberOfTabViewItems(tabView)
	}
}

// HasTabViewDidChangeNumberOfTabViewItems returns true if a handler for TabViewDidChangeNumberOfTabViewItems has been set.
func (d *TabViewDelegate) HasTabViewDidChangeNumberOfTabViewItems() bool {
	return d._TabViewDidChangeNumberOfTabViewItems != nil
}

// TabViewDelegateObject wraps an existing Objective-C object that conforms to the PTabViewDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TabViewDelegateObject struct {
	objectivec.Object
}

// NewTabViewDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTabViewDelegate protocol.
func NewTabViewDelegateObject(obj objectivec.Object) *TabViewDelegateObject {
	return &TabViewDelegateObject{obj}
}

// Make sure TabViewDelegateObject implements PTabViewDelegate.
var _ PTabViewDelegate = (*TabViewDelegateObject)(nil)

// TabViewDidSelectTabViewItem implements the PTabViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TabViewDelegateObject) TabViewDidSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	objc.Send[objc.ID](o.ID, objc.Sel("tabView:didSelectTabViewItem:"), tabView, tabViewItem)
}

// HasTabViewDidSelectTabViewItem returns true; this is a placeholder for optional method checks.
func (o *TabViewDelegateObject) HasTabViewDidSelectTabViewItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TabViewShouldSelectTabViewItem implements the PTabViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TabViewDelegateObject) TabViewShouldSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) bool {
	return objc.Send[bool](o.ID, objc.Sel("tabView:shouldSelectTabViewItem:"), tabView, tabViewItem)
}

// HasTabViewShouldSelectTabViewItem returns true; this is a placeholder for optional method checks.
func (o *TabViewDelegateObject) HasTabViewShouldSelectTabViewItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TabViewWillSelectTabViewItem implements the PTabViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TabViewDelegateObject) TabViewWillSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	objc.Send[objc.ID](o.ID, objc.Sel("tabView:willSelectTabViewItem:"), tabView, tabViewItem)
}

// HasTabViewWillSelectTabViewItem returns true; this is a placeholder for optional method checks.
func (o *TabViewDelegateObject) HasTabViewWillSelectTabViewItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TabViewDidChangeNumberOfTabViewItems implements the PTabViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TabViewDelegateObject) TabViewDidChangeNumberOfTabViewItems(tabView ITabView) {
	objc.Send[objc.ID](o.ID, objc.Sel("tabViewDidChangeNumberOfTabViewItems:"), tabView)
}

// HasTabViewDidChangeNumberOfTabViewItems returns true; this is a placeholder for optional method checks.
func (o *TabViewDelegateObject) HasTabViewDidChangeNumberOfTabViewItems() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
