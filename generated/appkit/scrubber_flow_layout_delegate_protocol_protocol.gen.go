// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
	ScrubberLayoutSizeForItemAtIndex(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) Size
	HasScrubberLayoutSizeForItemAtIndex() bool
}

// ScrubberFlowLayoutDelegate is a delegate implementation builder for the PScrubberFlowLayoutDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ScrubberFlowLayoutDelegate struct {
	_ScrubberLayoutSizeForItemAtIndex func(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) Size
}

// SetScrubberLayoutSizeForItemAtIndex sets the handler for the ScrubberLayoutSizeForItemAtIndex delegate method.
//
// Asks the delegate for the size of each item in a scrubber whose items are arranged in a flow layout.
func (d *ScrubberFlowLayoutDelegate) SetScrubberLayoutSizeForItemAtIndex(f func(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) Size) {
	d._ScrubberLayoutSizeForItemAtIndex = f
}

// ScrubberLayoutSizeForItemAtIndex implements the PScrubberFlowLayoutDelegate interface.
func (d *ScrubberFlowLayoutDelegate) ScrubberLayoutSizeForItemAtIndex(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) Size {
	if d._ScrubberLayoutSizeForItemAtIndex != nil {
		return d._ScrubberLayoutSizeForItemAtIndex(scrubber, layout, itemIndex)
	}
	var zero Size
	return zero
}

// HasScrubberLayoutSizeForItemAtIndex returns true if a handler for ScrubberLayoutSizeForItemAtIndex has been set.
func (d *ScrubberFlowLayoutDelegate) HasScrubberLayoutSizeForItemAtIndex() bool {
	return d._ScrubberLayoutSizeForItemAtIndex != nil
}
