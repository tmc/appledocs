// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PScrubberDataSource is the NSScrubberDataSource protocol interface.
//
// A set of methods that a scrubber data source object implements to provide items to the scrubber from an associated data collection in your app.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSScrubberDataSource
type PScrubberDataSource interface {
	// Required methods
	NumberOfItemsForScrubber(scrubber IScrubber) int
	ScrubberViewForItemAtIndex(scrubber IScrubber, index int) IScrubberItemView
}

// ScrubberDataSource is a delegate implementation builder for the PScrubberDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ScrubberDataSource struct {
	_NumberOfItemsForScrubber func(scrubber IScrubber) int
	_ScrubberViewForItemAtIndex func(scrubber IScrubber, index int) IScrubberItemView
}

// SetNumberOfItemsForScrubber sets the handler for the NumberOfItemsForScrubber delegate method.
//
// Asks the data source for the number of items in the scrubber.
func (d *ScrubberDataSource) SetNumberOfItemsForScrubber(f func(scrubber IScrubber) int) {
	d._NumberOfItemsForScrubber = f
}

// SetScrubberViewForItemAtIndex sets the handler for the ScrubberViewForItemAtIndex delegate method.
//
// Asks the data source object for the view the corresponds to the specified item in the scrubber.
func (d *ScrubberDataSource) SetScrubberViewForItemAtIndex(f func(scrubber IScrubber, index int) IScrubberItemView) {
	d._ScrubberViewForItemAtIndex = f
}

// NumberOfItemsForScrubber implements the PScrubberDataSource interface.
func (d *ScrubberDataSource) NumberOfItemsForScrubber(scrubber IScrubber) int {
	if d._NumberOfItemsForScrubber != nil {
		return d._NumberOfItemsForScrubber(scrubber)
	}
	var zero int
	return zero
}

// HasNumberOfItemsForScrubber returns true if a handler for NumberOfItemsForScrubber has been set.
func (d *ScrubberDataSource) HasNumberOfItemsForScrubber() bool {
	return d._NumberOfItemsForScrubber != nil
}

// ScrubberViewForItemAtIndex implements the PScrubberDataSource interface.
func (d *ScrubberDataSource) ScrubberViewForItemAtIndex(scrubber IScrubber, index int) IScrubberItemView {
	if d._ScrubberViewForItemAtIndex != nil {
		return d._ScrubberViewForItemAtIndex(scrubber, index)
	}
	var zero IScrubberItemView
	return zero
}

// HasScrubberViewForItemAtIndex returns true if a handler for ScrubberViewForItemAtIndex has been set.
func (d *ScrubberDataSource) HasScrubberViewForItemAtIndex() bool {
	return d._ScrubberViewForItemAtIndex != nil
}

// ScrubberDataSourceObject wraps an existing Objective-C object that conforms to the PScrubberDataSource protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type ScrubberDataSourceObject struct {
	objectivec.Object
}

// NewScrubberDataSourceObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSScrubberDataSource protocol.
func NewScrubberDataSourceObject(obj objectivec.Object) *ScrubberDataSourceObject {
	return &ScrubberDataSourceObject{obj}
}

// Make sure ScrubberDataSourceObject implements PScrubberDataSource.
var _ PScrubberDataSource = (*ScrubberDataSourceObject)(nil)

// NumberOfItemsForScrubber implements the PScrubberDataSource interface.
// This required method is always available on objects conforming to NumberOfItemsForScrubber.
func (o *ScrubberDataSourceObject) NumberOfItemsForScrubber(scrubber IScrubber) int {
	return objc.Send[int](o.ID, objc.Sel("numberOfItemsForScrubber:"), scrubber)
}

// ScrubberViewForItemAtIndex implements the PScrubberDataSource interface.
// This required method is always available on objects conforming to ScrubberViewForItemAtIndex.
func (o *ScrubberDataSourceObject) ScrubberViewForItemAtIndex(scrubber IScrubber, index int) IScrubberItemView {
	return objc.Send[IScrubberItemView](o.ID, objc.Sel("scrubber:viewForItemAtIndex:"), scrubber, index)
}
