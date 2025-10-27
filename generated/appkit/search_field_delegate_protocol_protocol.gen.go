// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PSearchFieldDelegate is the NSSearchFieldDelegate protocol interface.
//
// A protocol that a search field delegate can use to determine when a search started or ended.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSSearchFieldDelegate
type PSearchFieldDelegate interface {
	// Optional methods
	SearchFieldDidEndSearching(sender ISearchField)
	HasSearchFieldDidEndSearching() bool
	SearchFieldDidStartSearching(sender ISearchField)
	HasSearchFieldDidStartSearching() bool
}

// SearchFieldDelegate is a delegate implementation builder for the PSearchFieldDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SearchFieldDelegate struct {
	_SearchFieldDidEndSearching func(sender ISearchField)
	_SearchFieldDidStartSearching func(sender ISearchField)
}

// SetSearchFieldDidEndSearching sets the handler for the SearchFieldDidEndSearching delegate method.
//
// The method that is called when the search field has ended its search for content.
func (d *SearchFieldDelegate) SetSearchFieldDidEndSearching(f func(sender ISearchField)) {
	d._SearchFieldDidEndSearching = f
}

// SetSearchFieldDidStartSearching sets the handler for the SearchFieldDidStartSearching delegate method.
//
// The method that is called when the search field begins searching for content.
func (d *SearchFieldDelegate) SetSearchFieldDidStartSearching(f func(sender ISearchField)) {
	d._SearchFieldDidStartSearching = f
}

// SearchFieldDidEndSearching implements the PSearchFieldDelegate interface.
func (d *SearchFieldDelegate) SearchFieldDidEndSearching(sender ISearchField) {
	if d._SearchFieldDidEndSearching != nil {
		d._SearchFieldDidEndSearching(sender)
	}
}

// HasSearchFieldDidEndSearching returns true if a handler for SearchFieldDidEndSearching has been set.
func (d *SearchFieldDelegate) HasSearchFieldDidEndSearching() bool {
	return d._SearchFieldDidEndSearching != nil
}

// SearchFieldDidStartSearching implements the PSearchFieldDelegate interface.
func (d *SearchFieldDelegate) SearchFieldDidStartSearching(sender ISearchField) {
	if d._SearchFieldDidStartSearching != nil {
		d._SearchFieldDidStartSearching(sender)
	}
}

// HasSearchFieldDidStartSearching returns true if a handler for SearchFieldDidStartSearching has been set.
func (d *SearchFieldDelegate) HasSearchFieldDidStartSearching() bool {
	return d._SearchFieldDidStartSearching != nil
}

// SearchFieldDelegateObject wraps an existing Objective-C object that conforms to the PSearchFieldDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type SearchFieldDelegateObject struct {
	objectivec.Object
}

// NewSearchFieldDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSSearchFieldDelegate protocol.
func NewSearchFieldDelegateObject(obj objectivec.Object) *SearchFieldDelegateObject {
	return &SearchFieldDelegateObject{obj}
}

// Make sure SearchFieldDelegateObject implements PSearchFieldDelegate.
var _ PSearchFieldDelegate = (*SearchFieldDelegateObject)(nil)

// SearchFieldDidEndSearching implements the PSearchFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SearchFieldDelegateObject) SearchFieldDidEndSearching(sender ISearchField) {
	objc.Send[objc.ID](o.ID, objc.Sel("searchFieldDidEndSearching:"), sender)
}

// HasSearchFieldDidEndSearching returns true; this is a placeholder for optional method checks.
func (o *SearchFieldDelegateObject) HasSearchFieldDidEndSearching() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SearchFieldDidStartSearching implements the PSearchFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *SearchFieldDelegateObject) SearchFieldDidStartSearching(sender ISearchField) {
	objc.Send[objc.ID](o.ID, objc.Sel("searchFieldDidStartSearching:"), sender)
}

// HasSearchFieldDidStartSearching returns true; this is a placeholder for optional method checks.
func (o *SearchFieldDelegateObject) HasSearchFieldDidStartSearching() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
