// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PNCWidgetSearchViewDelegate is the NCWidgetSearchViewDelegate protocol interface.
//
// The interface for enabling user searches in the search view controller of a macOS Today widget.
//
// Availability:
//   - macOS 10.10+ (Deprecated in 11.0)
//
// See: doc://com.apple.notificationcenter/documentation/NotificationCenter/NCWidgetSearchViewDelegate
type PNCWidgetSearchViewDelegate interface {
	// Required methods
	WidgetSearchResultSelected(controller INCWidgetSearchViewController, object objc.IObject)/* debug [protocol_interface/required_method]: WidgetSearchResultSelected */
	WidgetSearchSearchForTermMaxResults(controller INCWidgetSearchViewController, searchTerm objc.IObject /* cross-framework: NSString */, max uint)/* debug [protocol_interface/required_method]: WidgetSearchSearchForTermMaxResults */
	WidgetSearchTermCleared(controller INCWidgetSearchViewController)/* debug [protocol_interface/required_method]: WidgetSearchTermCleared */
}

// NCWidgetSearchViewDelegate is a delegate implementation builder for the PNCWidgetSearchViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type NCWidgetSearchViewDelegate struct {
	_WidgetSearchResultSelected func(controller INCWidgetSearchViewController, object objc.IObject)
	_WidgetSearchSearchForTermMaxResults func(controller INCWidgetSearchViewController, searchTerm objc.IObject /* cross-framework: NSString */, max uint)
	_WidgetSearchTermCleared func(controller INCWidgetSearchViewController)
}

// SetWidgetSearchResultSelected sets the handler for the WidgetSearchResultSelected delegate method.
//
// Tells the delegate that a user chose the specified search result.
func (d *NCWidgetSearchViewDelegate) SetWidgetSearchResultSelected(f func(controller INCWidgetSearchViewController, object objc.IObject)) {
	d._WidgetSearchResultSelected = f
}

// SetWidgetSearchSearchForTermMaxResults sets the handler for the WidgetSearchSearchForTermMaxResults delegate method.
//
// Asks the delegate to search using the specified term.
func (d *NCWidgetSearchViewDelegate) SetWidgetSearchSearchForTermMaxResults(f func(controller INCWidgetSearchViewController, searchTerm objc.IObject /* cross-framework: NSString */, max uint)) {
	d._WidgetSearchSearchForTermMaxResults = f
}

// SetWidgetSearchTermCleared sets the handler for the WidgetSearchTermCleared delegate method.
//
// Tells the delegate that a user cleared the search field.
func (d *NCWidgetSearchViewDelegate) SetWidgetSearchTermCleared(f func(controller INCWidgetSearchViewController)) {
	d._WidgetSearchTermCleared = f
}

// WidgetSearchResultSelected implements the PNCWidgetSearchViewDelegate interface.
func (d *NCWidgetSearchViewDelegate) WidgetSearchResultSelected(controller INCWidgetSearchViewController, object objc.IObject) {
	if d._WidgetSearchResultSelected != nil {
		d._WidgetSearchResultSelected(controller, object)
	}
}

// HasWidgetSearchResultSelected returns true if a handler for WidgetSearchResultSelected has been set.
func (d *NCWidgetSearchViewDelegate) HasWidgetSearchResultSelected() bool {
	return d._WidgetSearchResultSelected != nil
}

// WidgetSearchSearchForTermMaxResults implements the PNCWidgetSearchViewDelegate interface.
func (d *NCWidgetSearchViewDelegate) WidgetSearchSearchForTermMaxResults(controller INCWidgetSearchViewController, searchTerm objc.IObject /* cross-framework: NSString */, max uint) {
	if d._WidgetSearchSearchForTermMaxResults != nil {
		d._WidgetSearchSearchForTermMaxResults(controller, searchTerm, max)
	}
}

// HasWidgetSearchSearchForTermMaxResults returns true if a handler for WidgetSearchSearchForTermMaxResults has been set.
func (d *NCWidgetSearchViewDelegate) HasWidgetSearchSearchForTermMaxResults() bool {
	return d._WidgetSearchSearchForTermMaxResults != nil
}

// WidgetSearchTermCleared implements the PNCWidgetSearchViewDelegate interface.
func (d *NCWidgetSearchViewDelegate) WidgetSearchTermCleared(controller INCWidgetSearchViewController) {
	if d._WidgetSearchTermCleared != nil {
		d._WidgetSearchTermCleared(controller)
	}
}

// HasWidgetSearchTermCleared returns true if a handler for WidgetSearchTermCleared has been set.
func (d *NCWidgetSearchViewDelegate) HasWidgetSearchTermCleared() bool {
	return d._WidgetSearchTermCleared != nil
}
