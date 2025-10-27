// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PScrubberFlowLayoutDelegate is the NSScrubberFlowLayoutDelegate protocol interface.
//
// A protocol that a scrubber delegate can adopt to provide the size of an item.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSScrubberFlowLayoutDelegate
type PScrubberFlowLayoutDelegate interface {
	// Optional methods
	ScrubberLayoutSizeForItemAtIndex(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) corefoundation.CGSize
	HasScrubberLayoutSizeForItemAtIndex() bool
}

// ScrubberFlowLayoutDelegate is a delegate implementation builder for the PScrubberFlowLayoutDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ScrubberFlowLayoutDelegate struct {
	_ScrubberLayoutSizeForItemAtIndex func(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) corefoundation.CGSize
}

// SetScrubberLayoutSizeForItemAtIndex sets the handler for the ScrubberLayoutSizeForItemAtIndex delegate method.
//
// Asks the delegate for the size of each item in a scrubber whose items are arranged in a flow layout.
func (d *ScrubberFlowLayoutDelegate) SetScrubberLayoutSizeForItemAtIndex(f func(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) corefoundation.CGSize) {
	d._ScrubberLayoutSizeForItemAtIndex = f
}

// ScrubberLayoutSizeForItemAtIndex implements the PScrubberFlowLayoutDelegate interface.
func (d *ScrubberFlowLayoutDelegate) ScrubberLayoutSizeForItemAtIndex(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) corefoundation.CGSize {
	if d._ScrubberLayoutSizeForItemAtIndex != nil {
		return d._ScrubberLayoutSizeForItemAtIndex(scrubber, layout, itemIndex)
	}
	var zero corefoundation.CGSize
	return zero
}

// HasScrubberLayoutSizeForItemAtIndex returns true if a handler for ScrubberLayoutSizeForItemAtIndex has been set.
func (d *ScrubberFlowLayoutDelegate) HasScrubberLayoutSizeForItemAtIndex() bool {
	return d._ScrubberLayoutSizeForItemAtIndex != nil
}

// ScrubberFlowLayoutDelegateObject wraps an existing Objective-C object that conforms to the PScrubberFlowLayoutDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type ScrubberFlowLayoutDelegateObject struct {
	objectivec.Object
}

// NewScrubberFlowLayoutDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSScrubberFlowLayoutDelegate protocol.
func NewScrubberFlowLayoutDelegateObject(obj objectivec.Object) *ScrubberFlowLayoutDelegateObject {
	return &ScrubberFlowLayoutDelegateObject{obj}
}

// Make sure ScrubberFlowLayoutDelegateObject implements PScrubberFlowLayoutDelegate.
var _ PScrubberFlowLayoutDelegate = (*ScrubberFlowLayoutDelegateObject)(nil)

// ScrubberLayoutSizeForItemAtIndex implements the PScrubberFlowLayoutDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ScrubberFlowLayoutDelegateObject) ScrubberLayoutSizeForItemAtIndex(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) corefoundation.CGSize {
	return objc.Send[corefoundation.CGSize](o.ID, objc.Sel("scrubber:layout:sizeForItemAtIndex:"), scrubber, layout, itemIndex)
}

// HasScrubberLayoutSizeForItemAtIndex returns true; this is a placeholder for optional method checks.
func (o *ScrubberFlowLayoutDelegateObject) HasScrubberLayoutSizeForItemAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
