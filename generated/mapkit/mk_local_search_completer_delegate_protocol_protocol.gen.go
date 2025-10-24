// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PMKLocalSearchCompleterDelegate is the MKLocalSearchCompleterDelegate protocol interface.
//
// Methods the delegate calls with search completion data.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.3+
//   - iPadOS 9.3+
//   - macOS 10.11.4+
//   - tvOS 9.2+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// See: doc://com.apple.mapkit/documentation/MapKit/MKLocalSearchCompleterDelegate
type PMKLocalSearchCompleterDelegate interface {
	// Optional methods
	CompleterDidFailWithError(completer IMKLocalSearchCompleter, error_ objc.IObject /* cross-framework: Error */)
	HasCompleterDidFailWithError() bool
	CompleterDidUpdateResults(completer IMKLocalSearchCompleter)
	HasCompleterDidUpdateResults() bool
}

// MKLocalSearchCompleterDelegate is a delegate implementation builder for the PMKLocalSearchCompleterDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MKLocalSearchCompleterDelegate struct {
	_CompleterDidFailWithError func(completer IMKLocalSearchCompleter, error_ objc.IObject /* cross-framework: Error */)
	_CompleterDidUpdateResults func(completer IMKLocalSearchCompleter)
}

// SetCompleterDidFailWithError sets the handler for the CompleterDidFailWithError delegate method.
//
// Tells the method when the specified search completer is unable to generate a list of search results.
func (d *MKLocalSearchCompleterDelegate) SetCompleterDidFailWithError(f func(completer IMKLocalSearchCompleter, error_ objc.IObject /* cross-framework: Error */)) {
	d._CompleterDidFailWithError = f
}

// SetCompleterDidUpdateResults sets the handler for the CompleterDidUpdateResults delegate method.
//
// Tells the method when the specified search completer updates its array of search completions.
func (d *MKLocalSearchCompleterDelegate) SetCompleterDidUpdateResults(f func(completer IMKLocalSearchCompleter)) {
	d._CompleterDidUpdateResults = f
}

// CompleterDidFailWithError implements the PMKLocalSearchCompleterDelegate interface.
func (d *MKLocalSearchCompleterDelegate) CompleterDidFailWithError(completer IMKLocalSearchCompleter, error_ objc.IObject /* cross-framework: Error */) {
	if d._CompleterDidFailWithError != nil {
		d._CompleterDidFailWithError(completer, error_)
	}
}

// HasCompleterDidFailWithError returns true if a handler for CompleterDidFailWithError has been set.
func (d *MKLocalSearchCompleterDelegate) HasCompleterDidFailWithError() bool {
	return d._CompleterDidFailWithError != nil
}

// CompleterDidUpdateResults implements the PMKLocalSearchCompleterDelegate interface.
func (d *MKLocalSearchCompleterDelegate) CompleterDidUpdateResults(completer IMKLocalSearchCompleter) {
	if d._CompleterDidUpdateResults != nil {
		d._CompleterDidUpdateResults(completer)
	}
}

// HasCompleterDidUpdateResults returns true if a handler for CompleterDidUpdateResults has been set.
func (d *MKLocalSearchCompleterDelegate) HasCompleterDidUpdateResults() bool {
	return d._CompleterDidUpdateResults != nil
}
