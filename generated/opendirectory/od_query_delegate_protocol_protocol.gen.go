// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PODQueryDelegate is the ODQueryDelegate protocol interface.
//
// The   protocol defines methods for receiving results returned from an Open Directory query.
//
// Availability:
//   - Mac Catalyst +
//   - macOS +
//
// See: doc://com.apple.opendirectory/documentation/OpenDirectory/ODQueryDelegate
type PODQueryDelegate interface {
	// Required methods
	QueryFoundResultsError(inQuery IODQuery, inResults objc.IObject /* cross-framework: NSArray */, inError objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: QueryFoundResultsError */
}

// ODQueryDelegate is a delegate implementation builder for the PODQueryDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ODQueryDelegate struct {
	_QueryFoundResultsError func(inQuery IODQuery, inResults objc.IObject /* cross-framework: NSArray */, inError objc.IObject /* cross-framework: Error */)
}

// SetQueryFoundResultsError sets the handler for the QueryFoundResultsError delegate method.
//
// The delegate method called as results are returned from a query scheduled in a run loop.
func (d *ODQueryDelegate) SetQueryFoundResultsError(f func(inQuery IODQuery, inResults objc.IObject /* cross-framework: NSArray */, inError objc.IObject /* cross-framework: Error */)) {
	d._QueryFoundResultsError = f
}

// QueryFoundResultsError implements the PODQueryDelegate interface.
func (d *ODQueryDelegate) QueryFoundResultsError(inQuery IODQuery, inResults objc.IObject /* cross-framework: NSArray */, inError objc.IObject /* cross-framework: Error */) {
	if d._QueryFoundResultsError != nil {
		d._QueryFoundResultsError(inQuery, inResults, inError)
	}
}

// HasQueryFoundResultsError returns true if a handler for QueryFoundResultsError has been set.
func (d *ODQueryDelegate) HasQueryFoundResultsError() bool {
	return d._QueryFoundResultsError != nil
}
