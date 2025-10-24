// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
