// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCandidateListTouchBarItemDelegate is the NSCandidateListTouchBarItemDelegate protocol interface.
//
// A set of methods that a candidate list item delegate uses to enable selection state and list visibility.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSCandidateListTouchBarItemDelegate
type PCandidateListTouchBarItemDelegate interface {
	// Optional methods
	CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int)
	HasCandidateListTouchBarItemBeginSelectingCandidateAtIndex() bool
	CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(anItem ICandidateListTouchBarItem, previousIndex int, index int)
	HasCandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex() bool
	CandidateListTouchBarItemChangedCandidateListVisibility(anItem ICandidateListTouchBarItem, isVisible bool)
	HasCandidateListTouchBarItemChangedCandidateListVisibility() bool
	CandidateListTouchBarItemEndSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int)
	HasCandidateListTouchBarItemEndSelectingCandidateAtIndex() bool
}

// CandidateListTouchBarItemDelegate is a delegate implementation builder for the PCandidateListTouchBarItemDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CandidateListTouchBarItemDelegate struct {
	_CandidateListTouchBarItemBeginSelectingCandidateAtIndex func(anItem ICandidateListTouchBarItem, index int)
	_CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex func(anItem ICandidateListTouchBarItem, previousIndex int, index int)
	_CandidateListTouchBarItemChangedCandidateListVisibility func(anItem ICandidateListTouchBarItem, isVisible bool)
	_CandidateListTouchBarItemEndSelectingCandidateAtIndex func(anItem ICandidateListTouchBarItem, index int)
}

// SetCandidateListTouchBarItemBeginSelectingCandidateAtIndex sets the handler for the CandidateListTouchBarItemBeginSelectingCandidateAtIndex delegate method.
//
// Tells the delegate that the user has started touching one of the candidates in the candidate list item.
func (d *CandidateListTouchBarItemDelegate) SetCandidateListTouchBarItemBeginSelectingCandidateAtIndex(f func(anItem ICandidateListTouchBarItem, index int)) {
	d._CandidateListTouchBarItemBeginSelectingCandidateAtIndex = f
}

// SetCandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex sets the handler for the CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex delegate method.
//
// Tells the delegate that user has moved from touching one candidate in the candidate list item to another.
func (d *CandidateListTouchBarItemDelegate) SetCandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(f func(anItem ICandidateListTouchBarItem, previousIndex int, index int)) {
	d._CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex = f
}

// SetCandidateListTouchBarItemChangedCandidateListVisibility sets the handler for the CandidateListTouchBarItemChangedCandidateListVisibility delegate method.
//
// Tells the delegate that the visibility of the candidate list has changed.
func (d *CandidateListTouchBarItemDelegate) SetCandidateListTouchBarItemChangedCandidateListVisibility(f func(anItem ICandidateListTouchBarItem, isVisible bool)) {
	d._CandidateListTouchBarItemChangedCandidateListVisibility = f
}

// SetCandidateListTouchBarItemEndSelectingCandidateAtIndex sets the handler for the CandidateListTouchBarItemEndSelectingCandidateAtIndex delegate method.
//
// Tells the delegate that a user has stopped touching candidates in the candidate list item.
func (d *CandidateListTouchBarItemDelegate) SetCandidateListTouchBarItemEndSelectingCandidateAtIndex(f func(anItem ICandidateListTouchBarItem, index int)) {
	d._CandidateListTouchBarItemEndSelectingCandidateAtIndex = f
}

// CandidateListTouchBarItemBeginSelectingCandidateAtIndex implements the PCandidateListTouchBarItemDelegate interface.
func (d *CandidateListTouchBarItemDelegate) CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	if d._CandidateListTouchBarItemBeginSelectingCandidateAtIndex != nil {
		d._CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem, index)
	}
}

// HasCandidateListTouchBarItemBeginSelectingCandidateAtIndex returns true if a handler for CandidateListTouchBarItemBeginSelectingCandidateAtIndex has been set.
func (d *CandidateListTouchBarItemDelegate) HasCandidateListTouchBarItemBeginSelectingCandidateAtIndex() bool {
	return d._CandidateListTouchBarItemBeginSelectingCandidateAtIndex != nil
}

// CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex implements the PCandidateListTouchBarItemDelegate interface.
func (d *CandidateListTouchBarItemDelegate) CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(anItem ICandidateListTouchBarItem, previousIndex int, index int) {
	if d._CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex != nil {
		d._CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(anItem, previousIndex, index)
	}
}

// HasCandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex returns true if a handler for CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex has been set.
func (d *CandidateListTouchBarItemDelegate) HasCandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex() bool {
	return d._CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex != nil
}

// CandidateListTouchBarItemChangedCandidateListVisibility implements the PCandidateListTouchBarItemDelegate interface.
func (d *CandidateListTouchBarItemDelegate) CandidateListTouchBarItemChangedCandidateListVisibility(anItem ICandidateListTouchBarItem, isVisible bool) {
	if d._CandidateListTouchBarItemChangedCandidateListVisibility != nil {
		d._CandidateListTouchBarItemChangedCandidateListVisibility(anItem, isVisible)
	}
}

// HasCandidateListTouchBarItemChangedCandidateListVisibility returns true if a handler for CandidateListTouchBarItemChangedCandidateListVisibility has been set.
func (d *CandidateListTouchBarItemDelegate) HasCandidateListTouchBarItemChangedCandidateListVisibility() bool {
	return d._CandidateListTouchBarItemChangedCandidateListVisibility != nil
}

// CandidateListTouchBarItemEndSelectingCandidateAtIndex implements the PCandidateListTouchBarItemDelegate interface.
func (d *CandidateListTouchBarItemDelegate) CandidateListTouchBarItemEndSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	if d._CandidateListTouchBarItemEndSelectingCandidateAtIndex != nil {
		d._CandidateListTouchBarItemEndSelectingCandidateAtIndex(anItem, index)
	}
}

// HasCandidateListTouchBarItemEndSelectingCandidateAtIndex returns true if a handler for CandidateListTouchBarItemEndSelectingCandidateAtIndex has been set.
func (d *CandidateListTouchBarItemDelegate) HasCandidateListTouchBarItemEndSelectingCandidateAtIndex() bool {
	return d._CandidateListTouchBarItemEndSelectingCandidateAtIndex != nil
}
