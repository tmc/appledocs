// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PScrubberDelegate is the NSScrubberDelegate protocol interface.
//
// A set of methods that a scrubber delegate implements to respond to user interactions.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSScrubberDelegate
type PScrubberDelegate interface {
	// Optional methods
	DidBeginInteractingWithScrubber(scrubber IScrubber)
	HasDidBeginInteractingWithScrubber() bool
	DidCancelInteractingWithScrubber(scrubber IScrubber)
	HasDidCancelInteractingWithScrubber() bool
	DidFinishInteractingWithScrubber(scrubber IScrubber)
	HasDidFinishInteractingWithScrubber() bool
	ScrubberDidChangeVisibleRange(scrubber IScrubber, visibleRange corefoundation.Range)
	HasScrubberDidChangeVisibleRange() bool
	ScrubberDidHighlightItemAtIndex(scrubber IScrubber, highlightedIndex int)
	HasScrubberDidHighlightItemAtIndex() bool
	ScrubberDidSelectItemAtIndex(scrubber IScrubber, selectedIndex int)
	HasScrubberDidSelectItemAtIndex() bool
}

// ScrubberDelegate is a delegate implementation builder for the PScrubberDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ScrubberDelegate struct {
	_DidBeginInteractingWithScrubber func(scrubber IScrubber)
	_DidCancelInteractingWithScrubber func(scrubber IScrubber)
	_DidFinishInteractingWithScrubber func(scrubber IScrubber)
	_ScrubberDidChangeVisibleRange func(scrubber IScrubber, visibleRange corefoundation.Range)
	_ScrubberDidHighlightItemAtIndex func(scrubber IScrubber, highlightedIndex int)
	_ScrubberDidSelectItemAtIndex func(scrubber IScrubber, selectedIndex int)
}

// SetDidBeginInteractingWithScrubber sets the handler for the DidBeginInteractingWithScrubber delegate method.
//
// Tells the delegate that the user is panning or scrolling the scrubber.
func (d *ScrubberDelegate) SetDidBeginInteractingWithScrubber(f func(scrubber IScrubber)) {
	d._DidBeginInteractingWithScrubber = f
}

// SetDidCancelInteractingWithScrubber sets the handler for the DidCancelInteractingWithScrubber delegate method.
//
// Tells the delegate that a user interaction with the scrubber has been canceled.
func (d *ScrubberDelegate) SetDidCancelInteractingWithScrubber(f func(scrubber IScrubber)) {
	d._DidCancelInteractingWithScrubber = f
}

// SetDidFinishInteractingWithScrubber sets the handler for the DidFinishInteractingWithScrubber delegate method.
//
// Tells the delegate that a pan or scroll interaction with the scrubber has ended.
func (d *ScrubberDelegate) SetDidFinishInteractingWithScrubber(f func(scrubber IScrubber)) {
	d._DidFinishInteractingWithScrubber = f
}

// SetScrubberDidChangeVisibleRange sets the handler for the ScrubberDidChangeVisibleRange delegate method.
//
// Tells the delegate that the range of items currently visible in the scrubber has changed.
func (d *ScrubberDelegate) SetScrubberDidChangeVisibleRange(f func(scrubber IScrubber, visibleRange corefoundation.Range)) {
	d._ScrubberDidChangeVisibleRange = f
}

// SetScrubberDidHighlightItemAtIndex sets the handler for the ScrubberDidHighlightItemAtIndex delegate method.
//
// Tells the delegate that the item at the specified index was highlighted.
func (d *ScrubberDelegate) SetScrubberDidHighlightItemAtIndex(f func(scrubber IScrubber, highlightedIndex int)) {
	d._ScrubberDidHighlightItemAtIndex = f
}

// SetScrubberDidSelectItemAtIndex sets the handler for the ScrubberDidSelectItemAtIndex delegate method.
//
// Tells the delegate that the item at the specified index was selected.
func (d *ScrubberDelegate) SetScrubberDidSelectItemAtIndex(f func(scrubber IScrubber, selectedIndex int)) {
	d._ScrubberDidSelectItemAtIndex = f
}

// DidBeginInteractingWithScrubber implements the PScrubberDelegate interface.
func (d *ScrubberDelegate) DidBeginInteractingWithScrubber(scrubber IScrubber) {
	if d._DidBeginInteractingWithScrubber != nil {
		d._DidBeginInteractingWithScrubber(scrubber)
	}
}

// HasDidBeginInteractingWithScrubber returns true if a handler for DidBeginInteractingWithScrubber has been set.
func (d *ScrubberDelegate) HasDidBeginInteractingWithScrubber() bool {
	return d._DidBeginInteractingWithScrubber != nil
}

// DidCancelInteractingWithScrubber implements the PScrubberDelegate interface.
func (d *ScrubberDelegate) DidCancelInteractingWithScrubber(scrubber IScrubber) {
	if d._DidCancelInteractingWithScrubber != nil {
		d._DidCancelInteractingWithScrubber(scrubber)
	}
}

// HasDidCancelInteractingWithScrubber returns true if a handler for DidCancelInteractingWithScrubber has been set.
func (d *ScrubberDelegate) HasDidCancelInteractingWithScrubber() bool {
	return d._DidCancelInteractingWithScrubber != nil
}

// DidFinishInteractingWithScrubber implements the PScrubberDelegate interface.
func (d *ScrubberDelegate) DidFinishInteractingWithScrubber(scrubber IScrubber) {
	if d._DidFinishInteractingWithScrubber != nil {
		d._DidFinishInteractingWithScrubber(scrubber)
	}
}

// HasDidFinishInteractingWithScrubber returns true if a handler for DidFinishInteractingWithScrubber has been set.
func (d *ScrubberDelegate) HasDidFinishInteractingWithScrubber() bool {
	return d._DidFinishInteractingWithScrubber != nil
}

// ScrubberDidChangeVisibleRange implements the PScrubberDelegate interface.
func (d *ScrubberDelegate) ScrubberDidChangeVisibleRange(scrubber IScrubber, visibleRange corefoundation.Range) {
	if d._ScrubberDidChangeVisibleRange != nil {
		d._ScrubberDidChangeVisibleRange(scrubber, visibleRange)
	}
}

// HasScrubberDidChangeVisibleRange returns true if a handler for ScrubberDidChangeVisibleRange has been set.
func (d *ScrubberDelegate) HasScrubberDidChangeVisibleRange() bool {
	return d._ScrubberDidChangeVisibleRange != nil
}

// ScrubberDidHighlightItemAtIndex implements the PScrubberDelegate interface.
func (d *ScrubberDelegate) ScrubberDidHighlightItemAtIndex(scrubber IScrubber, highlightedIndex int) {
	if d._ScrubberDidHighlightItemAtIndex != nil {
		d._ScrubberDidHighlightItemAtIndex(scrubber, highlightedIndex)
	}
}

// HasScrubberDidHighlightItemAtIndex returns true if a handler for ScrubberDidHighlightItemAtIndex has been set.
func (d *ScrubberDelegate) HasScrubberDidHighlightItemAtIndex() bool {
	return d._ScrubberDidHighlightItemAtIndex != nil
}

// ScrubberDidSelectItemAtIndex implements the PScrubberDelegate interface.
func (d *ScrubberDelegate) ScrubberDidSelectItemAtIndex(scrubber IScrubber, selectedIndex int) {
	if d._ScrubberDidSelectItemAtIndex != nil {
		d._ScrubberDidSelectItemAtIndex(scrubber, selectedIndex)
	}
}

// HasScrubberDidSelectItemAtIndex returns true if a handler for ScrubberDidSelectItemAtIndex has been set.
func (d *ScrubberDelegate) HasScrubberDidSelectItemAtIndex() bool {
	return d._ScrubberDidSelectItemAtIndex != nil
}
