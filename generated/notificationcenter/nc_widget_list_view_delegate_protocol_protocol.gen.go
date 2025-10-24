// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/appkit"
)

// PNCWidgetListViewDelegate is the NCWidgetListViewDelegate protocol interface.
//
// The interface for handling content display and editing in the list view of a macOS Today widget.
//
// Availability:
//   - macOS 10.10+ (Deprecated in 11.0)
//
// See: doc://com.apple.notificationcenter/documentation/NotificationCenter/NCWidgetListViewDelegate
type PNCWidgetListViewDelegate interface {
	// Required methods
	WidgetListViewControllerForRow(list INCWidgetListViewController, row uint) appkit.ViewController/* debug [protocol_interface/required_method]: WidgetListViewControllerForRow */
	// Optional methods
	WidgetListDidRemoveRow(list INCWidgetListViewController, row uint)
	HasWidgetListDidRemoveRow() bool
	WidgetListDidReorderRowToRow(list INCWidgetListViewController, row uint, newIndex uint)
	HasWidgetListDidReorderRowToRow() bool
	WidgetListShouldRemoveRow(list INCWidgetListViewController, row uint) bool
	HasWidgetListShouldRemoveRow() bool
	WidgetListShouldReorderRow(list INCWidgetListViewController, row uint) bool
	HasWidgetListShouldReorderRow() bool
	WidgetListPerformAddAction(list INCWidgetListViewController)
	HasWidgetListPerformAddAction() bool
}

// NCWidgetListViewDelegate is a delegate implementation builder for the PNCWidgetListViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type NCWidgetListViewDelegate struct {
	_WidgetListDidRemoveRow func(list INCWidgetListViewController, row uint)
	_WidgetListDidReorderRowToRow func(list INCWidgetListViewController, row uint, newIndex uint)
	_WidgetListShouldRemoveRow func(list INCWidgetListViewController, row uint) bool
	_WidgetListShouldReorderRow func(list INCWidgetListViewController, row uint) bool
	_WidgetListPerformAddAction func(list INCWidgetListViewController)
	_WidgetListViewControllerForRow func(list INCWidgetListViewController, row uint) appkit.ViewController
}

// SetWidgetListDidRemoveRow sets the handler for the WidgetListDidRemoveRow delegate method.
//
// Tells the delegate that the specified row was removed from the list.
func (d *NCWidgetListViewDelegate) SetWidgetListDidRemoveRow(f func(list INCWidgetListViewController, row uint)) {
	d._WidgetListDidRemoveRow = f
}

// SetWidgetListDidReorderRowToRow sets the handler for the WidgetListDidReorderRowToRow delegate method.
//
// Tells the delegate that the specified row was moved to a new location.
func (d *NCWidgetListViewDelegate) SetWidgetListDidReorderRowToRow(f func(list INCWidgetListViewController, row uint, newIndex uint)) {
	d._WidgetListDidReorderRowToRow = f
}

// SetWidgetListShouldRemoveRow sets the handler for the WidgetListShouldRemoveRow delegate method.
//
// Asks the delegate to allow or prohibit the specified row to be removed from the list.
func (d *NCWidgetListViewDelegate) SetWidgetListShouldRemoveRow(f func(list INCWidgetListViewController, row uint) bool) {
	d._WidgetListShouldRemoveRow = f
}

// SetWidgetListShouldReorderRow sets the handler for the WidgetListShouldReorderRow delegate method.
//
// Asks the delegate to allow or prohibit the specified row to be moved to a new location in the list.
func (d *NCWidgetListViewDelegate) SetWidgetListShouldReorderRow(f func(list INCWidgetListViewController, row uint) bool) {
	d._WidgetListShouldReorderRow = f
}

// SetWidgetListPerformAddAction sets the handler for the WidgetListPerformAddAction delegate method.
//
// Asks the delegate to perform an action when the Add (+) button is clicked.
func (d *NCWidgetListViewDelegate) SetWidgetListPerformAddAction(f func(list INCWidgetListViewController)) {
	d._WidgetListPerformAddAction = f
}

// SetWidgetListViewControllerForRow sets the handler for the WidgetListViewControllerForRow delegate method.
//
// Asks the delegate for a content view controller for the specified row.
func (d *NCWidgetListViewDelegate) SetWidgetListViewControllerForRow(f func(list INCWidgetListViewController, row uint) appkit.ViewController) {
	d._WidgetListViewControllerForRow = f
}

// WidgetListDidRemoveRow implements the PNCWidgetListViewDelegate interface.
func (d *NCWidgetListViewDelegate) WidgetListDidRemoveRow(list INCWidgetListViewController, row uint) {
	if d._WidgetListDidRemoveRow != nil {
		d._WidgetListDidRemoveRow(list, row)
	}
}

// HasWidgetListDidRemoveRow returns true if a handler for WidgetListDidRemoveRow has been set.
func (d *NCWidgetListViewDelegate) HasWidgetListDidRemoveRow() bool {
	return d._WidgetListDidRemoveRow != nil
}

// WidgetListDidReorderRowToRow implements the PNCWidgetListViewDelegate interface.
func (d *NCWidgetListViewDelegate) WidgetListDidReorderRowToRow(list INCWidgetListViewController, row uint, newIndex uint) {
	if d._WidgetListDidReorderRowToRow != nil {
		d._WidgetListDidReorderRowToRow(list, row, newIndex)
	}
}

// HasWidgetListDidReorderRowToRow returns true if a handler for WidgetListDidReorderRowToRow has been set.
func (d *NCWidgetListViewDelegate) HasWidgetListDidReorderRowToRow() bool {
	return d._WidgetListDidReorderRowToRow != nil
}

// WidgetListShouldRemoveRow implements the PNCWidgetListViewDelegate interface.
func (d *NCWidgetListViewDelegate) WidgetListShouldRemoveRow(list INCWidgetListViewController, row uint) bool {
	if d._WidgetListShouldRemoveRow != nil {
		return d._WidgetListShouldRemoveRow(list, row)
	}
	var zero bool
	return zero
}

// HasWidgetListShouldRemoveRow returns true if a handler for WidgetListShouldRemoveRow has been set.
func (d *NCWidgetListViewDelegate) HasWidgetListShouldRemoveRow() bool {
	return d._WidgetListShouldRemoveRow != nil
}

// WidgetListShouldReorderRow implements the PNCWidgetListViewDelegate interface.
func (d *NCWidgetListViewDelegate) WidgetListShouldReorderRow(list INCWidgetListViewController, row uint) bool {
	if d._WidgetListShouldReorderRow != nil {
		return d._WidgetListShouldReorderRow(list, row)
	}
	var zero bool
	return zero
}

// HasWidgetListShouldReorderRow returns true if a handler for WidgetListShouldReorderRow has been set.
func (d *NCWidgetListViewDelegate) HasWidgetListShouldReorderRow() bool {
	return d._WidgetListShouldReorderRow != nil
}

// WidgetListPerformAddAction implements the PNCWidgetListViewDelegate interface.
func (d *NCWidgetListViewDelegate) WidgetListPerformAddAction(list INCWidgetListViewController) {
	if d._WidgetListPerformAddAction != nil {
		d._WidgetListPerformAddAction(list)
	}
}

// HasWidgetListPerformAddAction returns true if a handler for WidgetListPerformAddAction has been set.
func (d *NCWidgetListViewDelegate) HasWidgetListPerformAddAction() bool {
	return d._WidgetListPerformAddAction != nil
}

// WidgetListViewControllerForRow implements the PNCWidgetListViewDelegate interface.
func (d *NCWidgetListViewDelegate) WidgetListViewControllerForRow(list INCWidgetListViewController, row uint) appkit.ViewController {
	if d._WidgetListViewControllerForRow != nil {
		return d._WidgetListViewControllerForRow(list, row)
	}
	var zero appkit.ViewController
	return zero
}

// HasWidgetListViewControllerForRow returns true if a handler for WidgetListViewControllerForRow has been set.
func (d *NCWidgetListViewDelegate) HasWidgetListViewControllerForRow() bool {
	return d._WidgetListViewControllerForRow != nil
}
