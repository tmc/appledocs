// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PIKSlideshowDataSource is the IKSlideshowDataSource protocol interface.
//
// The   protocol describes the methods that an   object uses to access the contents of its data source object.
//
// Availability:
//   - macOS 10.4+
//
// See: doc://com.apple.quartz/documentation/Quartz/IKSlideshowDataSource
type PIKSlideshowDataSource interface {
	// Required methods
	NumberOfSlideshowItems() uint/* debug [protocol_interface/required_method]: NumberOfSlideshowItems */
	SlideshowItemAtIndex(index uint) objc.ID/* debug [protocol_interface/required_method]: SlideshowItemAtIndex */
	// Optional methods
	CanExportSlideshowItemAtIndexToApplication(index uint, applicationBundleIdentifier objc.IObject /* cross-framework: NSString */) bool
	HasCanExportSlideshowItemAtIndexToApplication() bool
	NameOfSlideshowItemAtIndex(index uint) foundation.String
	HasNameOfSlideshowItemAtIndex() bool
	SlideshowDidChangeCurrentIndex(newIndex uint)
	HasSlideshowDidChangeCurrentIndex() bool
	SlideshowDidStop()
	HasSlideshowDidStop() bool
	SlideshowWillStart()
	HasSlideshowWillStart() bool
}

// IKSlideshowDataSource is a delegate implementation builder for the PIKSlideshowDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type IKSlideshowDataSource struct {
	_CanExportSlideshowItemAtIndexToApplication func(index uint, applicationBundleIdentifier objc.IObject /* cross-framework: NSString */) bool
	_NameOfSlideshowItemAtIndex func(index uint) foundation.String
	_SlideshowDidChangeCurrentIndex func(newIndex uint)
	_SlideshowDidStop func()
	_SlideshowWillStart func()
	_NumberOfSlideshowItems func() uint
	_SlideshowItemAtIndex func(index uint) objc.ID
}

// SetCanExportSlideshowItemAtIndexToApplication sets the handler for the CanExportSlideshowItemAtIndexToApplication delegate method.
//
// Reports whether the export button should be enabled for a slideshow item.
func (d *IKSlideshowDataSource) SetCanExportSlideshowItemAtIndexToApplication(f func(index uint, applicationBundleIdentifier objc.IObject /* cross-framework: NSString */) bool) {
	d._CanExportSlideshowItemAtIndexToApplication = f
}

// SetNameOfSlideshowItemAtIndex sets the handler for the NameOfSlideshowItemAtIndex delegate method.
//
// Returns the display name for item at the specified index.
func (d *IKSlideshowDataSource) SetNameOfSlideshowItemAtIndex(f func(index uint) foundation.String) {
	d._NameOfSlideshowItemAtIndex = f
}

// SetSlideshowDidChangeCurrentIndex sets the handler for the SlideshowDidChangeCurrentIndex delegate method.
//
// Performs custom tasks when the slideshow changes to the item at the specified index.
func (d *IKSlideshowDataSource) SetSlideshowDidChangeCurrentIndex(f func(newIndex uint)) {
	d._SlideshowDidChangeCurrentIndex = f
}

// SetSlideshowDidStop sets the handler for the SlideshowDidStop delegate method.
//
// Performs custom tasks when the slideshow stops.
func (d *IKSlideshowDataSource) SetSlideshowDidStop(f func()) {
	d._SlideshowDidStop = f
}

// SetSlideshowWillStart sets the handler for the SlideshowWillStart delegate method.
//
// Performs custom tasks when the slideshow is about to start.
func (d *IKSlideshowDataSource) SetSlideshowWillStart(f func()) {
	d._SlideshowWillStart = f
}

// SetNumberOfSlideshowItems sets the handler for the NumberOfSlideshowItems delegate method.
//
// Returns the number of items in a slideshow.
func (d *IKSlideshowDataSource) SetNumberOfSlideshowItems(f func() uint) {
	d._NumberOfSlideshowItems = f
}

// SetSlideshowItemAtIndex sets the handler for the SlideshowItemAtIndex delegate method.
//
// Returns the item for a given index
func (d *IKSlideshowDataSource) SetSlideshowItemAtIndex(f func(index uint) objc.ID) {
	d._SlideshowItemAtIndex = f
}

// CanExportSlideshowItemAtIndexToApplication implements the PIKSlideshowDataSource interface.
func (d *IKSlideshowDataSource) CanExportSlideshowItemAtIndexToApplication(index uint, applicationBundleIdentifier objc.IObject /* cross-framework: NSString */) bool {
	if d._CanExportSlideshowItemAtIndexToApplication != nil {
		return d._CanExportSlideshowItemAtIndexToApplication(index, applicationBundleIdentifier)
	}
	var zero bool
	return zero
}

// HasCanExportSlideshowItemAtIndexToApplication returns true if a handler for CanExportSlideshowItemAtIndexToApplication has been set.
func (d *IKSlideshowDataSource) HasCanExportSlideshowItemAtIndexToApplication() bool {
	return d._CanExportSlideshowItemAtIndexToApplication != nil
}

// NameOfSlideshowItemAtIndex implements the PIKSlideshowDataSource interface.
func (d *IKSlideshowDataSource) NameOfSlideshowItemAtIndex(index uint) foundation.String {
	if d._NameOfSlideshowItemAtIndex != nil {
		return d._NameOfSlideshowItemAtIndex(index)
	}
	var zero foundation.String
	return zero
}

// HasNameOfSlideshowItemAtIndex returns true if a handler for NameOfSlideshowItemAtIndex has been set.
func (d *IKSlideshowDataSource) HasNameOfSlideshowItemAtIndex() bool {
	return d._NameOfSlideshowItemAtIndex != nil
}

// SlideshowDidChangeCurrentIndex implements the PIKSlideshowDataSource interface.
func (d *IKSlideshowDataSource) SlideshowDidChangeCurrentIndex(newIndex uint) {
	if d._SlideshowDidChangeCurrentIndex != nil {
		d._SlideshowDidChangeCurrentIndex(newIndex)
	}
}

// HasSlideshowDidChangeCurrentIndex returns true if a handler for SlideshowDidChangeCurrentIndex has been set.
func (d *IKSlideshowDataSource) HasSlideshowDidChangeCurrentIndex() bool {
	return d._SlideshowDidChangeCurrentIndex != nil
}

// SlideshowDidStop implements the PIKSlideshowDataSource interface.
func (d *IKSlideshowDataSource) SlideshowDidStop() {
	if d._SlideshowDidStop != nil {
		d._SlideshowDidStop()
	}
}

// HasSlideshowDidStop returns true if a handler for SlideshowDidStop has been set.
func (d *IKSlideshowDataSource) HasSlideshowDidStop() bool {
	return d._SlideshowDidStop != nil
}

// SlideshowWillStart implements the PIKSlideshowDataSource interface.
func (d *IKSlideshowDataSource) SlideshowWillStart() {
	if d._SlideshowWillStart != nil {
		d._SlideshowWillStart()
	}
}

// HasSlideshowWillStart returns true if a handler for SlideshowWillStart has been set.
func (d *IKSlideshowDataSource) HasSlideshowWillStart() bool {
	return d._SlideshowWillStart != nil
}

// NumberOfSlideshowItems implements the PIKSlideshowDataSource interface.
func (d *IKSlideshowDataSource) NumberOfSlideshowItems() uint {
	if d._NumberOfSlideshowItems != nil {
		return d._NumberOfSlideshowItems()
	}
	var zero uint
	return zero
}

// HasNumberOfSlideshowItems returns true if a handler for NumberOfSlideshowItems has been set.
func (d *IKSlideshowDataSource) HasNumberOfSlideshowItems() bool {
	return d._NumberOfSlideshowItems != nil
}

// SlideshowItemAtIndex implements the PIKSlideshowDataSource interface.
func (d *IKSlideshowDataSource) SlideshowItemAtIndex(index uint) objc.ID {
	if d._SlideshowItemAtIndex != nil {
		return d._SlideshowItemAtIndex(index)
	}
	var zero objc.ID
	return zero
}

// HasSlideshowItemAtIndex returns true if a handler for SlideshowItemAtIndex has been set.
func (d *IKSlideshowDataSource) HasSlideshowItemAtIndex() bool {
	return d._SlideshowItemAtIndex != nil
}
