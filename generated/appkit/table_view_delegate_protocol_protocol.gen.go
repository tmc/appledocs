// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/vision"
)

// PTableViewDelegate is the NSTableViewDelegate protocol interface.
//
// A set of optional methods you implement in a table view delegate to customize the behavior of the table view.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTableViewDelegate
type PTableViewDelegate interface {
	// Optional methods
	SelectionShouldChangeInTableView(tableView ITableView) bool
	HasSelectionShouldChangeInTableView() bool
	TableViewDataCellForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) Cell
	HasTableViewDataCellForTableColumnRow() bool
	TableViewDidAddRowViewForRow(tableView ITableView, rowView ITableRowView, row int)
	HasTableViewDidAddRowViewForRow() bool
	TableViewDidClickTableColumn(tableView ITableView, tableColumn ITableColumn)
	HasTableViewDidClickTableColumn() bool
	TableViewDidDragTableColumn(tableView ITableView, tableColumn ITableColumn)
	HasTableViewDidDragTableColumn() bool
	TableViewDidRemoveRowViewForRow(tableView ITableView, rowView ITableRowView, row int)
	HasTableViewDidRemoveRowViewForRow() bool
	TableViewHeightOfRow(tableView ITableView, row int) float64
	HasTableViewHeightOfRow() bool
	TableViewIsGroupRow(tableView ITableView, row int) bool
	HasTableViewIsGroupRow() bool
	TableViewMouseDownInHeaderOfTableColumn(tableView ITableView, tableColumn ITableColumn)
	HasTableViewMouseDownInHeaderOfTableColumn() bool
	TableViewNextTypeSelectMatchFromRowToRowForString(tableView ITableView, startRow int, endRow int, searchString objc.IObject /* cross-framework: NSString */) int
	HasTableViewNextTypeSelectMatchFromRowToRowForString() bool
	TableViewRowActionsForRowEdge(tableView ITableView, row int, edge TableRowActionEdge) []TableViewRowAction
	HasTableViewRowActionsForRowEdge() bool
	TableViewRowViewForRow(tableView ITableView, row int) TableRowView
	HasTableViewRowViewForRow() bool
	TableViewSelectionIndexesForProposedSelection(tableView ITableView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet
	HasTableViewSelectionIndexesForProposedSelection() bool
	TableViewShouldEditTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) bool
	HasTableViewShouldEditTableColumnRow() bool
	TableViewShouldReorderColumnToColumn(tableView ITableView, columnIndex int, newColumnIndex int) bool
	HasTableViewShouldReorderColumnToColumn() bool
	TableViewShouldSelectTableColumn(tableView ITableView, tableColumn ITableColumn) bool
	HasTableViewShouldSelectTableColumn() bool
	TableViewShouldSelectRow(tableView ITableView, row int) bool
	HasTableViewShouldSelectRow() bool
	TableViewShouldShowCellExpansionForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) bool
	HasTableViewShouldShowCellExpansionForTableColumnRow() bool
	TableViewShouldTrackCellForTableColumnRow(tableView ITableView, cell ICell, tableColumn ITableColumn, row int) bool
	HasTableViewShouldTrackCellForTableColumnRow() bool
	TableViewShouldTypeSelectForEventWithCurrentSearchString(tableView ITableView, event IEvent, searchString objc.IObject /* cross-framework: NSString */) bool
	HasTableViewShouldTypeSelectForEventWithCurrentSearchString() bool
	TableViewSizeToFitWidthOfColumn(tableView ITableView, column int) float64
	HasTableViewSizeToFitWidthOfColumn() bool
	TableViewToolTipForCellRectTableColumnRowMouseLocation(tableView ITableView, cell ICell, rect RectPointer /* not a class type */, tableColumn ITableColumn, row int, mouseLocation vision.Point) foundation.String
	HasTableViewToolTipForCellRectTableColumnRowMouseLocation() bool
	TableViewTypeSelectStringForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) foundation.String
	HasTableViewTypeSelectStringForTableColumnRow() bool
	TableViewUserCanChangeVisibilityOfTableColumn(tableView ITableView, column ITableColumn) bool
	HasTableViewUserCanChangeVisibilityOfTableColumn() bool
	TableViewUserDidChangeVisibilityOfTableColumns(tableView ITableView, columns []TableColumn)
	HasTableViewUserDidChangeVisibilityOfTableColumns() bool
	TableViewViewForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) View
	HasTableViewViewForTableColumnRow() bool
	TableViewWillDisplayCellForTableColumnRow(tableView ITableView, cell objc.IObject, tableColumn ITableColumn, row int)
	HasTableViewWillDisplayCellForTableColumnRow() bool
	TableViewColumnDidMove(notification foundation.Notification)
	HasTableViewColumnDidMove() bool
	TableViewColumnDidResize(notification foundation.Notification)
	HasTableViewColumnDidResize() bool
	TableViewSelectionDidChange(notification foundation.Notification)
	HasTableViewSelectionDidChange() bool
	TableViewSelectionIsChanging(notification foundation.Notification)
	HasTableViewSelectionIsChanging() bool
}

// TableViewDelegate is a delegate implementation builder for the PTableViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TableViewDelegate struct {
	_SelectionShouldChangeInTableView func(tableView ITableView) bool
	_TableViewDataCellForTableColumnRow func(tableView ITableView, tableColumn ITableColumn, row int) Cell
	_TableViewDidAddRowViewForRow func(tableView ITableView, rowView ITableRowView, row int)
	_TableViewDidClickTableColumn func(tableView ITableView, tableColumn ITableColumn)
	_TableViewDidDragTableColumn func(tableView ITableView, tableColumn ITableColumn)
	_TableViewDidRemoveRowViewForRow func(tableView ITableView, rowView ITableRowView, row int)
	_TableViewHeightOfRow func(tableView ITableView, row int) float64
	_TableViewIsGroupRow func(tableView ITableView, row int) bool
	_TableViewMouseDownInHeaderOfTableColumn func(tableView ITableView, tableColumn ITableColumn)
	_TableViewNextTypeSelectMatchFromRowToRowForString func(tableView ITableView, startRow int, endRow int, searchString objc.IObject /* cross-framework: NSString */) int
	_TableViewRowActionsForRowEdge func(tableView ITableView, row int, edge TableRowActionEdge) []TableViewRowAction
	_TableViewRowViewForRow func(tableView ITableView, row int) TableRowView
	_TableViewSelectionIndexesForProposedSelection func(tableView ITableView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet
	_TableViewShouldEditTableColumnRow func(tableView ITableView, tableColumn ITableColumn, row int) bool
	_TableViewShouldReorderColumnToColumn func(tableView ITableView, columnIndex int, newColumnIndex int) bool
	_TableViewShouldSelectTableColumn func(tableView ITableView, tableColumn ITableColumn) bool
	_TableViewShouldSelectRow func(tableView ITableView, row int) bool
	_TableViewShouldShowCellExpansionForTableColumnRow func(tableView ITableView, tableColumn ITableColumn, row int) bool
	_TableViewShouldTrackCellForTableColumnRow func(tableView ITableView, cell ICell, tableColumn ITableColumn, row int) bool
	_TableViewShouldTypeSelectForEventWithCurrentSearchString func(tableView ITableView, event IEvent, searchString objc.IObject /* cross-framework: NSString */) bool
	_TableViewSizeToFitWidthOfColumn func(tableView ITableView, column int) float64
	_TableViewToolTipForCellRectTableColumnRowMouseLocation func(tableView ITableView, cell ICell, rect RectPointer /* not a class type */, tableColumn ITableColumn, row int, mouseLocation vision.Point) foundation.String
	_TableViewTypeSelectStringForTableColumnRow func(tableView ITableView, tableColumn ITableColumn, row int) foundation.String
	_TableViewUserCanChangeVisibilityOfTableColumn func(tableView ITableView, column ITableColumn) bool
	_TableViewUserDidChangeVisibilityOfTableColumns func(tableView ITableView, columns []TableColumn)
	_TableViewViewForTableColumnRow func(tableView ITableView, tableColumn ITableColumn, row int) View
	_TableViewWillDisplayCellForTableColumnRow func(tableView ITableView, cell objc.IObject, tableColumn ITableColumn, row int)
	_TableViewColumnDidMove func(notification foundation.Notification)
	_TableViewColumnDidResize func(notification foundation.Notification)
	_TableViewSelectionDidChange func(notification foundation.Notification)
	_TableViewSelectionIsChanging func(notification foundation.Notification)
}

// SetSelectionShouldChangeInTableView sets the handler for the SelectionShouldChangeInTableView delegate method.
//
// Asks the delegate if the user is allowed to change the selection.
func (d *TableViewDelegate) SetSelectionShouldChangeInTableView(f func(tableView ITableView) bool) {
	d._SelectionShouldChangeInTableView = f
}

// SetTableViewDataCellForTableColumnRow sets the handler for the TableViewDataCellForTableColumnRow delegate method.
//
// Asks the delegate for a custom data cell for the specified row and column.
func (d *TableViewDelegate) SetTableViewDataCellForTableColumnRow(f func(tableView ITableView, tableColumn ITableColumn, row int) Cell) {
	d._TableViewDataCellForTableColumnRow = f
}

// SetTableViewDidAddRowViewForRow sets the handler for the TableViewDidAddRowViewForRow delegate method.
//
// Tells the delegate that a row view was added at the specified row.
func (d *TableViewDelegate) SetTableViewDidAddRowViewForRow(f func(tableView ITableView, rowView ITableRowView, row int)) {
	d._TableViewDidAddRowViewForRow = f
}

// SetTableViewDidClickTableColumn sets the handler for the TableViewDidClickTableColumn delegate method.
//
// Tells the delegate that the mouse button was clicked in the specified table column, but the column was not dragged.
func (d *TableViewDelegate) SetTableViewDidClickTableColumn(f func(tableView ITableView, tableColumn ITableColumn)) {
	d._TableViewDidClickTableColumn = f
}

// SetTableViewDidDragTableColumn sets the handler for the TableViewDidDragTableColumn delegate method.
//
// Tells the delegate that the specified table column was dragged.
func (d *TableViewDelegate) SetTableViewDidDragTableColumn(f func(tableView ITableView, tableColumn ITableColumn)) {
	d._TableViewDidDragTableColumn = f
}

// SetTableViewDidRemoveRowViewForRow sets the handler for the TableViewDidRemoveRowViewForRow delegate method.
//
// Tells the delegate that a row view was removed from the table at the specified row.
func (d *TableViewDelegate) SetTableViewDidRemoveRowViewForRow(f func(tableView ITableView, rowView ITableRowView, row int)) {
	d._TableViewDidRemoveRowViewForRow = f
}

// SetTableViewHeightOfRow sets the handler for the TableViewHeightOfRow delegate method.
//
// Asks the delegate for the height of the specified row.
func (d *TableViewDelegate) SetTableViewHeightOfRow(f func(tableView ITableView, row int) float64) {
	d._TableViewHeightOfRow = f
}

// SetTableViewIsGroupRow sets the handler for the TableViewIsGroupRow delegate method.
//
// Returns whether the specified row is a group row.
func (d *TableViewDelegate) SetTableViewIsGroupRow(f func(tableView ITableView, row int) bool) {
	d._TableViewIsGroupRow = f
}

// SetTableViewMouseDownInHeaderOfTableColumn sets the handler for the TableViewMouseDownInHeaderOfTableColumn delegate method.
//
// Tells the delegate that the mouse button was clicked in the specified table column’s header.
func (d *TableViewDelegate) SetTableViewMouseDownInHeaderOfTableColumn(f func(tableView ITableView, tableColumn ITableColumn)) {
	d._TableViewMouseDownInHeaderOfTableColumn = f
}

// SetTableViewNextTypeSelectMatchFromRowToRowForString sets the handler for the TableViewNextTypeSelectMatchFromRowToRowForString delegate method.
//
// Asks the delegate for the row within the specified search range that matches the specified string.
func (d *TableViewDelegate) SetTableViewNextTypeSelectMatchFromRowToRowForString(f func(tableView ITableView, startRow int, endRow int, searchString objc.IObject /* cross-framework: NSString */) int) {
	d._TableViewNextTypeSelectMatchFromRowToRowForString = f
}

// SetTableViewRowActionsForRowEdge sets the handler for the TableViewRowActionsForRowEdge delegate method.
//
// Asks the delegate to provide an array of row actions to be attached to the specified edge of a table row and displayed when the user swipes horizontally across the row.
func (d *TableViewDelegate) SetTableViewRowActionsForRowEdge(f func(tableView ITableView, row int, edge TableRowActionEdge) []TableViewRowAction) {
	d._TableViewRowActionsForRowEdge = f
}

// SetTableViewRowViewForRow sets the handler for the TableViewRowViewForRow delegate method.
//
// Asks the delegate for a view to display the specified row.
func (d *TableViewDelegate) SetTableViewRowViewForRow(f func(tableView ITableView, row int) TableRowView) {
	d._TableViewRowViewForRow = f
}

// SetTableViewSelectionIndexesForProposedSelection sets the handler for the TableViewSelectionIndexesForProposedSelection delegate method.
//
// Asks the delegate to accept or reject the proposed selection.
func (d *TableViewDelegate) SetTableViewSelectionIndexesForProposedSelection(f func(tableView ITableView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet) {
	d._TableViewSelectionIndexesForProposedSelection = f
}

// SetTableViewShouldEditTableColumnRow sets the handler for the TableViewShouldEditTableColumnRow delegate method.
//
// Asks the delegate if the cell at the specified row and column can be edited.
func (d *TableViewDelegate) SetTableViewShouldEditTableColumnRow(f func(tableView ITableView, tableColumn ITableColumn, row int) bool) {
	d._TableViewShouldEditTableColumnRow = f
}

// SetTableViewShouldReorderColumnToColumn sets the handler for the TableViewShouldReorderColumnToColumn delegate method.
//
// Asks the delegate to allow or prohibit the specified column to be dragged to a new location.
func (d *TableViewDelegate) SetTableViewShouldReorderColumnToColumn(f func(tableView ITableView, columnIndex int, newColumnIndex int) bool) {
	d._TableViewShouldReorderColumnToColumn = f
}

// SetTableViewShouldSelectTableColumn sets the handler for the TableViewShouldSelectTableColumn delegate method.
//
// Asks the delegate whether the specified table column can be selected.
func (d *TableViewDelegate) SetTableViewShouldSelectTableColumn(f func(tableView ITableView, tableColumn ITableColumn) bool) {
	d._TableViewShouldSelectTableColumn = f
}

// SetTableViewShouldSelectRow sets the handler for the TableViewShouldSelectRow delegate method.
//
// Asks the delegate if the table view should allow selection of the specified row.
func (d *TableViewDelegate) SetTableViewShouldSelectRow(f func(tableView ITableView, row int) bool) {
	d._TableViewShouldSelectRow = f
}

// SetTableViewShouldShowCellExpansionForTableColumnRow sets the handler for the TableViewShouldShowCellExpansionForTableColumnRow delegate method.
//
// Asks the delegate if an expansion tooltip should be displayed for a specific row and column.
func (d *TableViewDelegate) SetTableViewShouldShowCellExpansionForTableColumnRow(f func(tableView ITableView, tableColumn ITableColumn, row int) bool) {
	d._TableViewShouldShowCellExpansionForTableColumnRow = f
}

// SetTableViewShouldTrackCellForTableColumnRow sets the handler for the TableViewShouldTrackCellForTableColumnRow delegate method.
//
// Asks the delegate whether the specified cell should be tracked.
func (d *TableViewDelegate) SetTableViewShouldTrackCellForTableColumnRow(f func(tableView ITableView, cell ICell, tableColumn ITableColumn, row int) bool) {
	d._TableViewShouldTrackCellForTableColumnRow = f
}

// SetTableViewShouldTypeSelectForEventWithCurrentSearchString sets the handler for the TableViewShouldTypeSelectForEventWithCurrentSearchString delegate method.
//
// Asks the delegate to allow or deny type select for the specified event and current search string.
func (d *TableViewDelegate) SetTableViewShouldTypeSelectForEventWithCurrentSearchString(f func(tableView ITableView, event IEvent, searchString objc.IObject /* cross-framework: NSString */) bool) {
	d._TableViewShouldTypeSelectForEventWithCurrentSearchString = f
}

// SetTableViewSizeToFitWidthOfColumn sets the handler for the TableViewSizeToFitWidthOfColumn delegate method.
//
// Asks the delegate to provide custom sizing behavior when a column’s resize divider is double clicked.
func (d *TableViewDelegate) SetTableViewSizeToFitWidthOfColumn(f func(tableView ITableView, column int) float64) {
	d._TableViewSizeToFitWidthOfColumn = f
}

// SetTableViewToolTipForCellRectTableColumnRowMouseLocation sets the handler for the TableViewToolTipForCellRectTableColumnRowMouseLocation delegate method.
//
// Asks the delegate for a string to display in a tooltip for the specified cell in the column and row.
func (d *TableViewDelegate) SetTableViewToolTipForCellRectTableColumnRowMouseLocation(f func(tableView ITableView, cell ICell, rect RectPointer /* not a class type */, tableColumn ITableColumn, row int, mouseLocation vision.Point) foundation.String) {
	d._TableViewToolTipForCellRectTableColumnRowMouseLocation = f
}

// SetTableViewTypeSelectStringForTableColumnRow sets the handler for the TableViewTypeSelectStringForTableColumnRow delegate method.
//
// Asks the delegate to provide an alternative text value used for type selection for the specified row and column.
func (d *TableViewDelegate) SetTableViewTypeSelectStringForTableColumnRow(f func(tableView ITableView, tableColumn ITableColumn, row int) foundation.String) {
	d._TableViewTypeSelectStringForTableColumnRow = f
}

// SetTableViewUserCanChangeVisibilityOfTableColumn sets the handler for the TableViewUserCanChangeVisibilityOfTableColumn delegate method.
//
// Asks the delegate to verify that the user can change the given column’s visibility.
func (d *TableViewDelegate) SetTableViewUserCanChangeVisibilityOfTableColumn(f func(tableView ITableView, column ITableColumn) bool) {
	d._TableViewUserCanChangeVisibilityOfTableColumn = f
}

// SetTableViewUserDidChangeVisibilityOfTableColumns sets the handler for the TableViewUserDidChangeVisibilityOfTableColumns delegate method.
//
// Tells the delegate that the user changed the visibility of one or more table columns.
func (d *TableViewDelegate) SetTableViewUserDidChangeVisibilityOfTableColumns(f func(tableView ITableView, columns []TableColumn)) {
	d._TableViewUserDidChangeVisibilityOfTableColumns = f
}

// SetTableViewViewForTableColumnRow sets the handler for the TableViewViewForTableColumnRow delegate method.
//
// Asks the delegate for a view to display the specified row and column.
func (d *TableViewDelegate) SetTableViewViewForTableColumnRow(f func(tableView ITableView, tableColumn ITableColumn, row int) View) {
	d._TableViewViewForTableColumnRow = f
}

// SetTableViewWillDisplayCellForTableColumnRow sets the handler for the TableViewWillDisplayCellForTableColumnRow delegate method.
//
// Tells the delegate that the table view will display the specified cell at the specified row and column.
func (d *TableViewDelegate) SetTableViewWillDisplayCellForTableColumnRow(f func(tableView ITableView, cell objc.IObject, tableColumn ITableColumn, row int)) {
	d._TableViewWillDisplayCellForTableColumnRow = f
}

// SetTableViewColumnDidMove sets the handler for the TableViewColumnDidMove delegate method.
//
// Tells the delegate that a table column was moved by user action.
func (d *TableViewDelegate) SetTableViewColumnDidMove(f func(notification foundation.Notification)) {
	d._TableViewColumnDidMove = f
}

// SetTableViewColumnDidResize sets the handler for the TableViewColumnDidResize delegate method.
//
// Tells the delegate that a table column was resized.
func (d *TableViewDelegate) SetTableViewColumnDidResize(f func(notification foundation.Notification)) {
	d._TableViewColumnDidResize = f
}

// SetTableViewSelectionDidChange sets the handler for the TableViewSelectionDidChange delegate method.
//
// Tells the delegate that the table view’s selection has changed.
func (d *TableViewDelegate) SetTableViewSelectionDidChange(f func(notification foundation.Notification)) {
	d._TableViewSelectionDidChange = f
}

// SetTableViewSelectionIsChanging sets the handler for the TableViewSelectionIsChanging delegate method.
//
// Tells the delegate that the table view’s selection is in the process of changing.
func (d *TableViewDelegate) SetTableViewSelectionIsChanging(f func(notification foundation.Notification)) {
	d._TableViewSelectionIsChanging = f
}

// SelectionShouldChangeInTableView implements the PTableViewDelegate interface.
func (d *TableViewDelegate) SelectionShouldChangeInTableView(tableView ITableView) bool {
	if d._SelectionShouldChangeInTableView != nil {
		return d._SelectionShouldChangeInTableView(tableView)
	}
	var zero bool
	return zero
}

// HasSelectionShouldChangeInTableView returns true if a handler for SelectionShouldChangeInTableView has been set.
func (d *TableViewDelegate) HasSelectionShouldChangeInTableView() bool {
	return d._SelectionShouldChangeInTableView != nil
}

// TableViewDataCellForTableColumnRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewDataCellForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) Cell {
	if d._TableViewDataCellForTableColumnRow != nil {
		return d._TableViewDataCellForTableColumnRow(tableView, tableColumn, row)
	}
	var zero Cell
	return zero
}

// HasTableViewDataCellForTableColumnRow returns true if a handler for TableViewDataCellForTableColumnRow has been set.
func (d *TableViewDelegate) HasTableViewDataCellForTableColumnRow() bool {
	return d._TableViewDataCellForTableColumnRow != nil
}

// TableViewDidAddRowViewForRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewDidAddRowViewForRow(tableView ITableView, rowView ITableRowView, row int) {
	if d._TableViewDidAddRowViewForRow != nil {
		d._TableViewDidAddRowViewForRow(tableView, rowView, row)
	}
}

// HasTableViewDidAddRowViewForRow returns true if a handler for TableViewDidAddRowViewForRow has been set.
func (d *TableViewDelegate) HasTableViewDidAddRowViewForRow() bool {
	return d._TableViewDidAddRowViewForRow != nil
}

// TableViewDidClickTableColumn implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewDidClickTableColumn(tableView ITableView, tableColumn ITableColumn) {
	if d._TableViewDidClickTableColumn != nil {
		d._TableViewDidClickTableColumn(tableView, tableColumn)
	}
}

// HasTableViewDidClickTableColumn returns true if a handler for TableViewDidClickTableColumn has been set.
func (d *TableViewDelegate) HasTableViewDidClickTableColumn() bool {
	return d._TableViewDidClickTableColumn != nil
}

// TableViewDidDragTableColumn implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewDidDragTableColumn(tableView ITableView, tableColumn ITableColumn) {
	if d._TableViewDidDragTableColumn != nil {
		d._TableViewDidDragTableColumn(tableView, tableColumn)
	}
}

// HasTableViewDidDragTableColumn returns true if a handler for TableViewDidDragTableColumn has been set.
func (d *TableViewDelegate) HasTableViewDidDragTableColumn() bool {
	return d._TableViewDidDragTableColumn != nil
}

// TableViewDidRemoveRowViewForRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewDidRemoveRowViewForRow(tableView ITableView, rowView ITableRowView, row int) {
	if d._TableViewDidRemoveRowViewForRow != nil {
		d._TableViewDidRemoveRowViewForRow(tableView, rowView, row)
	}
}

// HasTableViewDidRemoveRowViewForRow returns true if a handler for TableViewDidRemoveRowViewForRow has been set.
func (d *TableViewDelegate) HasTableViewDidRemoveRowViewForRow() bool {
	return d._TableViewDidRemoveRowViewForRow != nil
}

// TableViewHeightOfRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewHeightOfRow(tableView ITableView, row int) float64 {
	if d._TableViewHeightOfRow != nil {
		return d._TableViewHeightOfRow(tableView, row)
	}
	var zero float64
	return zero
}

// HasTableViewHeightOfRow returns true if a handler for TableViewHeightOfRow has been set.
func (d *TableViewDelegate) HasTableViewHeightOfRow() bool {
	return d._TableViewHeightOfRow != nil
}

// TableViewIsGroupRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewIsGroupRow(tableView ITableView, row int) bool {
	if d._TableViewIsGroupRow != nil {
		return d._TableViewIsGroupRow(tableView, row)
	}
	var zero bool
	return zero
}

// HasTableViewIsGroupRow returns true if a handler for TableViewIsGroupRow has been set.
func (d *TableViewDelegate) HasTableViewIsGroupRow() bool {
	return d._TableViewIsGroupRow != nil
}

// TableViewMouseDownInHeaderOfTableColumn implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewMouseDownInHeaderOfTableColumn(tableView ITableView, tableColumn ITableColumn) {
	if d._TableViewMouseDownInHeaderOfTableColumn != nil {
		d._TableViewMouseDownInHeaderOfTableColumn(tableView, tableColumn)
	}
}

// HasTableViewMouseDownInHeaderOfTableColumn returns true if a handler for TableViewMouseDownInHeaderOfTableColumn has been set.
func (d *TableViewDelegate) HasTableViewMouseDownInHeaderOfTableColumn() bool {
	return d._TableViewMouseDownInHeaderOfTableColumn != nil
}

// TableViewNextTypeSelectMatchFromRowToRowForString implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewNextTypeSelectMatchFromRowToRowForString(tableView ITableView, startRow int, endRow int, searchString objc.IObject /* cross-framework: NSString */) int {
	if d._TableViewNextTypeSelectMatchFromRowToRowForString != nil {
		return d._TableViewNextTypeSelectMatchFromRowToRowForString(tableView, startRow, endRow, searchString)
	}
	var zero int
	return zero
}

// HasTableViewNextTypeSelectMatchFromRowToRowForString returns true if a handler for TableViewNextTypeSelectMatchFromRowToRowForString has been set.
func (d *TableViewDelegate) HasTableViewNextTypeSelectMatchFromRowToRowForString() bool {
	return d._TableViewNextTypeSelectMatchFromRowToRowForString != nil
}

// TableViewRowActionsForRowEdge implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewRowActionsForRowEdge(tableView ITableView, row int, edge TableRowActionEdge) []TableViewRowAction {
	if d._TableViewRowActionsForRowEdge != nil {
		return d._TableViewRowActionsForRowEdge(tableView, row, edge)
	}
	var zero []TableViewRowAction
	return zero
}

// HasTableViewRowActionsForRowEdge returns true if a handler for TableViewRowActionsForRowEdge has been set.
func (d *TableViewDelegate) HasTableViewRowActionsForRowEdge() bool {
	return d._TableViewRowActionsForRowEdge != nil
}

// TableViewRowViewForRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewRowViewForRow(tableView ITableView, row int) TableRowView {
	if d._TableViewRowViewForRow != nil {
		return d._TableViewRowViewForRow(tableView, row)
	}
	var zero TableRowView
	return zero
}

// HasTableViewRowViewForRow returns true if a handler for TableViewRowViewForRow has been set.
func (d *TableViewDelegate) HasTableViewRowViewForRow() bool {
	return d._TableViewRowViewForRow != nil
}

// TableViewSelectionIndexesForProposedSelection implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewSelectionIndexesForProposedSelection(tableView ITableView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet {
	if d._TableViewSelectionIndexesForProposedSelection != nil {
		return d._TableViewSelectionIndexesForProposedSelection(tableView, proposedSelectionIndexes)
	}
	var zero foundation.IndexSet
	return zero
}

// HasTableViewSelectionIndexesForProposedSelection returns true if a handler for TableViewSelectionIndexesForProposedSelection has been set.
func (d *TableViewDelegate) HasTableViewSelectionIndexesForProposedSelection() bool {
	return d._TableViewSelectionIndexesForProposedSelection != nil
}

// TableViewShouldEditTableColumnRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewShouldEditTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) bool {
	if d._TableViewShouldEditTableColumnRow != nil {
		return d._TableViewShouldEditTableColumnRow(tableView, tableColumn, row)
	}
	var zero bool
	return zero
}

// HasTableViewShouldEditTableColumnRow returns true if a handler for TableViewShouldEditTableColumnRow has been set.
func (d *TableViewDelegate) HasTableViewShouldEditTableColumnRow() bool {
	return d._TableViewShouldEditTableColumnRow != nil
}

// TableViewShouldReorderColumnToColumn implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewShouldReorderColumnToColumn(tableView ITableView, columnIndex int, newColumnIndex int) bool {
	if d._TableViewShouldReorderColumnToColumn != nil {
		return d._TableViewShouldReorderColumnToColumn(tableView, columnIndex, newColumnIndex)
	}
	var zero bool
	return zero
}

// HasTableViewShouldReorderColumnToColumn returns true if a handler for TableViewShouldReorderColumnToColumn has been set.
func (d *TableViewDelegate) HasTableViewShouldReorderColumnToColumn() bool {
	return d._TableViewShouldReorderColumnToColumn != nil
}

// TableViewShouldSelectTableColumn implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewShouldSelectTableColumn(tableView ITableView, tableColumn ITableColumn) bool {
	if d._TableViewShouldSelectTableColumn != nil {
		return d._TableViewShouldSelectTableColumn(tableView, tableColumn)
	}
	var zero bool
	return zero
}

// HasTableViewShouldSelectTableColumn returns true if a handler for TableViewShouldSelectTableColumn has been set.
func (d *TableViewDelegate) HasTableViewShouldSelectTableColumn() bool {
	return d._TableViewShouldSelectTableColumn != nil
}

// TableViewShouldSelectRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewShouldSelectRow(tableView ITableView, row int) bool {
	if d._TableViewShouldSelectRow != nil {
		return d._TableViewShouldSelectRow(tableView, row)
	}
	var zero bool
	return zero
}

// HasTableViewShouldSelectRow returns true if a handler for TableViewShouldSelectRow has been set.
func (d *TableViewDelegate) HasTableViewShouldSelectRow() bool {
	return d._TableViewShouldSelectRow != nil
}

// TableViewShouldShowCellExpansionForTableColumnRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewShouldShowCellExpansionForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) bool {
	if d._TableViewShouldShowCellExpansionForTableColumnRow != nil {
		return d._TableViewShouldShowCellExpansionForTableColumnRow(tableView, tableColumn, row)
	}
	var zero bool
	return zero
}

// HasTableViewShouldShowCellExpansionForTableColumnRow returns true if a handler for TableViewShouldShowCellExpansionForTableColumnRow has been set.
func (d *TableViewDelegate) HasTableViewShouldShowCellExpansionForTableColumnRow() bool {
	return d._TableViewShouldShowCellExpansionForTableColumnRow != nil
}

// TableViewShouldTrackCellForTableColumnRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewShouldTrackCellForTableColumnRow(tableView ITableView, cell ICell, tableColumn ITableColumn, row int) bool {
	if d._TableViewShouldTrackCellForTableColumnRow != nil {
		return d._TableViewShouldTrackCellForTableColumnRow(tableView, cell, tableColumn, row)
	}
	var zero bool
	return zero
}

// HasTableViewShouldTrackCellForTableColumnRow returns true if a handler for TableViewShouldTrackCellForTableColumnRow has been set.
func (d *TableViewDelegate) HasTableViewShouldTrackCellForTableColumnRow() bool {
	return d._TableViewShouldTrackCellForTableColumnRow != nil
}

// TableViewShouldTypeSelectForEventWithCurrentSearchString implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewShouldTypeSelectForEventWithCurrentSearchString(tableView ITableView, event IEvent, searchString objc.IObject /* cross-framework: NSString */) bool {
	if d._TableViewShouldTypeSelectForEventWithCurrentSearchString != nil {
		return d._TableViewShouldTypeSelectForEventWithCurrentSearchString(tableView, event, searchString)
	}
	var zero bool
	return zero
}

// HasTableViewShouldTypeSelectForEventWithCurrentSearchString returns true if a handler for TableViewShouldTypeSelectForEventWithCurrentSearchString has been set.
func (d *TableViewDelegate) HasTableViewShouldTypeSelectForEventWithCurrentSearchString() bool {
	return d._TableViewShouldTypeSelectForEventWithCurrentSearchString != nil
}

// TableViewSizeToFitWidthOfColumn implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewSizeToFitWidthOfColumn(tableView ITableView, column int) float64 {
	if d._TableViewSizeToFitWidthOfColumn != nil {
		return d._TableViewSizeToFitWidthOfColumn(tableView, column)
	}
	var zero float64
	return zero
}

// HasTableViewSizeToFitWidthOfColumn returns true if a handler for TableViewSizeToFitWidthOfColumn has been set.
func (d *TableViewDelegate) HasTableViewSizeToFitWidthOfColumn() bool {
	return d._TableViewSizeToFitWidthOfColumn != nil
}

// TableViewToolTipForCellRectTableColumnRowMouseLocation implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewToolTipForCellRectTableColumnRowMouseLocation(tableView ITableView, cell ICell, rect RectPointer /* not a class type */, tableColumn ITableColumn, row int, mouseLocation vision.Point) foundation.String {
	if d._TableViewToolTipForCellRectTableColumnRowMouseLocation != nil {
		return d._TableViewToolTipForCellRectTableColumnRowMouseLocation(tableView, cell, rect, tableColumn, row, mouseLocation)
	}
	var zero foundation.String
	return zero
}

// HasTableViewToolTipForCellRectTableColumnRowMouseLocation returns true if a handler for TableViewToolTipForCellRectTableColumnRowMouseLocation has been set.
func (d *TableViewDelegate) HasTableViewToolTipForCellRectTableColumnRowMouseLocation() bool {
	return d._TableViewToolTipForCellRectTableColumnRowMouseLocation != nil
}

// TableViewTypeSelectStringForTableColumnRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewTypeSelectStringForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) foundation.String {
	if d._TableViewTypeSelectStringForTableColumnRow != nil {
		return d._TableViewTypeSelectStringForTableColumnRow(tableView, tableColumn, row)
	}
	var zero foundation.String
	return zero
}

// HasTableViewTypeSelectStringForTableColumnRow returns true if a handler for TableViewTypeSelectStringForTableColumnRow has been set.
func (d *TableViewDelegate) HasTableViewTypeSelectStringForTableColumnRow() bool {
	return d._TableViewTypeSelectStringForTableColumnRow != nil
}

// TableViewUserCanChangeVisibilityOfTableColumn implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewUserCanChangeVisibilityOfTableColumn(tableView ITableView, column ITableColumn) bool {
	if d._TableViewUserCanChangeVisibilityOfTableColumn != nil {
		return d._TableViewUserCanChangeVisibilityOfTableColumn(tableView, column)
	}
	var zero bool
	return zero
}

// HasTableViewUserCanChangeVisibilityOfTableColumn returns true if a handler for TableViewUserCanChangeVisibilityOfTableColumn has been set.
func (d *TableViewDelegate) HasTableViewUserCanChangeVisibilityOfTableColumn() bool {
	return d._TableViewUserCanChangeVisibilityOfTableColumn != nil
}

// TableViewUserDidChangeVisibilityOfTableColumns implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewUserDidChangeVisibilityOfTableColumns(tableView ITableView, columns []TableColumn) {
	if d._TableViewUserDidChangeVisibilityOfTableColumns != nil {
		d._TableViewUserDidChangeVisibilityOfTableColumns(tableView, columns)
	}
}

// HasTableViewUserDidChangeVisibilityOfTableColumns returns true if a handler for TableViewUserDidChangeVisibilityOfTableColumns has been set.
func (d *TableViewDelegate) HasTableViewUserDidChangeVisibilityOfTableColumns() bool {
	return d._TableViewUserDidChangeVisibilityOfTableColumns != nil
}

// TableViewViewForTableColumnRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewViewForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) View {
	if d._TableViewViewForTableColumnRow != nil {
		return d._TableViewViewForTableColumnRow(tableView, tableColumn, row)
	}
	var zero View
	return zero
}

// HasTableViewViewForTableColumnRow returns true if a handler for TableViewViewForTableColumnRow has been set.
func (d *TableViewDelegate) HasTableViewViewForTableColumnRow() bool {
	return d._TableViewViewForTableColumnRow != nil
}

// TableViewWillDisplayCellForTableColumnRow implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewWillDisplayCellForTableColumnRow(tableView ITableView, cell objc.IObject, tableColumn ITableColumn, row int) {
	if d._TableViewWillDisplayCellForTableColumnRow != nil {
		d._TableViewWillDisplayCellForTableColumnRow(tableView, cell, tableColumn, row)
	}
}

// HasTableViewWillDisplayCellForTableColumnRow returns true if a handler for TableViewWillDisplayCellForTableColumnRow has been set.
func (d *TableViewDelegate) HasTableViewWillDisplayCellForTableColumnRow() bool {
	return d._TableViewWillDisplayCellForTableColumnRow != nil
}

// TableViewColumnDidMove implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewColumnDidMove(notification foundation.Notification) {
	if d._TableViewColumnDidMove != nil {
		d._TableViewColumnDidMove(notification)
	}
}

// HasTableViewColumnDidMove returns true if a handler for TableViewColumnDidMove has been set.
func (d *TableViewDelegate) HasTableViewColumnDidMove() bool {
	return d._TableViewColumnDidMove != nil
}

// TableViewColumnDidResize implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewColumnDidResize(notification foundation.Notification) {
	if d._TableViewColumnDidResize != nil {
		d._TableViewColumnDidResize(notification)
	}
}

// HasTableViewColumnDidResize returns true if a handler for TableViewColumnDidResize has been set.
func (d *TableViewDelegate) HasTableViewColumnDidResize() bool {
	return d._TableViewColumnDidResize != nil
}

// TableViewSelectionDidChange implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewSelectionDidChange(notification foundation.Notification) {
	if d._TableViewSelectionDidChange != nil {
		d._TableViewSelectionDidChange(notification)
	}
}

// HasTableViewSelectionDidChange returns true if a handler for TableViewSelectionDidChange has been set.
func (d *TableViewDelegate) HasTableViewSelectionDidChange() bool {
	return d._TableViewSelectionDidChange != nil
}

// TableViewSelectionIsChanging implements the PTableViewDelegate interface.
func (d *TableViewDelegate) TableViewSelectionIsChanging(notification foundation.Notification) {
	if d._TableViewSelectionIsChanging != nil {
		d._TableViewSelectionIsChanging(notification)
	}
}

// HasTableViewSelectionIsChanging returns true if a handler for TableViewSelectionIsChanging has been set.
func (d *TableViewDelegate) HasTableViewSelectionIsChanging() bool {
	return d._TableViewSelectionIsChanging != nil
}
