// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// POutlineViewDelegate is the NSOutlineViewDelegate protocol interface.
//
// A set of optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSOutlineViewDelegate
type POutlineViewDelegate interface {
	// Optional methods
	OutlineViewDataCellForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) ICell
	HasOutlineViewDataCellForTableColumnItem() bool
	OutlineViewDidAddRowViewForRow(outlineView IOutlineView, rowView ITableRowView, row int)
	HasOutlineViewDidAddRowViewForRow() bool
	OutlineViewDidClickTableColumn(outlineView IOutlineView, tableColumn ITableColumn)
	HasOutlineViewDidClickTableColumn() bool
	OutlineViewDidDragTableColumn(outlineView IOutlineView, tableColumn ITableColumn)
	HasOutlineViewDidDragTableColumn() bool
	OutlineViewDidRemoveRowViewForRow(outlineView IOutlineView, rowView ITableRowView, row int)
	HasOutlineViewDidRemoveRowViewForRow() bool
	OutlineViewHeightOfRowByItem(outlineView IOutlineView, item objectivec.IObject) float64
	HasOutlineViewHeightOfRowByItem() bool
	OutlineViewIsGroupItem(outlineView IOutlineView, item objectivec.IObject) bool
	HasOutlineViewIsGroupItem() bool
	OutlineViewMouseDownInHeaderOfTableColumn(outlineView IOutlineView, tableColumn ITableColumn)
	HasOutlineViewMouseDownInHeaderOfTableColumn() bool
	OutlineViewNextTypeSelectMatchFromItemToItemForString(outlineView IOutlineView, startItem objectivec.IObject, endItem objectivec.IObject, searchString foundation.foundation.INSString) objc.ID
	HasOutlineViewNextTypeSelectMatchFromItemToItemForString() bool
	OutlineViewRowViewForItem(outlineView IOutlineView, item objectivec.IObject) ITableRowView
	HasOutlineViewRowViewForItem() bool
	OutlineViewSelectionIndexesForProposedSelection(outlineView IOutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet
	HasOutlineViewSelectionIndexesForProposedSelection() bool
	OutlineViewShouldCollapseItem(outlineView IOutlineView, item objectivec.IObject) bool
	HasOutlineViewShouldCollapseItem() bool
	OutlineViewShouldEditTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool
	HasOutlineViewShouldEditTableColumnItem() bool
	OutlineViewShouldExpandItem(outlineView IOutlineView, item objectivec.IObject) bool
	HasOutlineViewShouldExpandItem() bool
	OutlineViewShouldReorderColumnToColumn(outlineView IOutlineView, columnIndex int, newColumnIndex int) bool
	HasOutlineViewShouldReorderColumnToColumn() bool
	OutlineViewShouldSelectTableColumn(outlineView IOutlineView, tableColumn ITableColumn) bool
	HasOutlineViewShouldSelectTableColumn() bool
	OutlineViewShouldSelectItem(outlineView IOutlineView, item objectivec.IObject) bool
	HasOutlineViewShouldSelectItem() bool
	OutlineViewShouldShowCellExpansionForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool
	HasOutlineViewShouldShowCellExpansionForTableColumnItem() bool
	OutlineViewShouldShowOutlineCellForItem(outlineView IOutlineView, item objectivec.IObject) bool
	HasOutlineViewShouldShowOutlineCellForItem() bool
	OutlineViewShouldTrackCellForTableColumnItem(outlineView IOutlineView, cell ICell, tableColumn ITableColumn, item objectivec.IObject) bool
	HasOutlineViewShouldTrackCellForTableColumnItem() bool
	OutlineViewShouldTypeSelectForEventWithCurrentSearchString(outlineView IOutlineView, event IEvent, searchString foundation.foundation.INSString) bool
	HasOutlineViewShouldTypeSelectForEventWithCurrentSearchString() bool
	OutlineViewSizeToFitWidthOfColumn(outlineView IOutlineView, column int) float64
	HasOutlineViewSizeToFitWidthOfColumn() bool
	OutlineViewTintConfigurationForItem(outlineView IOutlineView, item objectivec.IObject) ITintConfiguration
	HasOutlineViewTintConfigurationForItem() bool
	OutlineViewToolTipForCellRectTableColumnItemMouseLocation(outlineView IOutlineView, cell ICell, rect RectPointer /* not a class type */, tableColumn ITableColumn, item objectivec.IObject, mouseLocation corefoundation.CGPoint) foundation.String
	HasOutlineViewToolTipForCellRectTableColumnItemMouseLocation() bool
	OutlineViewTypeSelectStringForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) foundation.String
	HasOutlineViewTypeSelectStringForTableColumnItem() bool
	OutlineViewUserCanChangeVisibilityOfTableColumn(outlineView IOutlineView, column ITableColumn) bool
	HasOutlineViewUserCanChangeVisibilityOfTableColumn() bool
	OutlineViewUserDidChangeVisibilityOfTableColumns(outlineView IOutlineView, columns []TableColumn)
	HasOutlineViewUserDidChangeVisibilityOfTableColumns() bool
	OutlineViewViewForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) IView
	HasOutlineViewViewForTableColumnItem() bool
	OutlineViewWillDisplayCellForTableColumnItem(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject)
	HasOutlineViewWillDisplayCellForTableColumnItem() bool
	OutlineViewWillDisplayOutlineCellForTableColumnItem(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject)
	HasOutlineViewWillDisplayOutlineCellForTableColumnItem() bool
	OutlineViewColumnDidMove(notification foundation.foundation.INSNotification)
	HasOutlineViewColumnDidMove() bool
	OutlineViewColumnDidResize(notification foundation.foundation.INSNotification)
	HasOutlineViewColumnDidResize() bool
	OutlineViewItemDidCollapse(notification foundation.foundation.INSNotification)
	HasOutlineViewItemDidCollapse() bool
	OutlineViewItemDidExpand(notification foundation.foundation.INSNotification)
	HasOutlineViewItemDidExpand() bool
	OutlineViewItemWillCollapse(notification foundation.foundation.INSNotification)
	HasOutlineViewItemWillCollapse() bool
	OutlineViewItemWillExpand(notification foundation.foundation.INSNotification)
	HasOutlineViewItemWillExpand() bool
	OutlineViewSelectionDidChange(notification foundation.foundation.INSNotification)
	HasOutlineViewSelectionDidChange() bool
	OutlineViewSelectionIsChanging(notification foundation.foundation.INSNotification)
	HasOutlineViewSelectionIsChanging() bool
	SelectionShouldChangeInOutlineView(outlineView IOutlineView) bool
	HasSelectionShouldChangeInOutlineView() bool
}

// OutlineViewDelegate is a delegate implementation builder for the POutlineViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type OutlineViewDelegate struct {
	_OutlineViewDataCellForTableColumnItem func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) ICell
	_OutlineViewDidAddRowViewForRow func(outlineView IOutlineView, rowView ITableRowView, row int)
	_OutlineViewDidClickTableColumn func(outlineView IOutlineView, tableColumn ITableColumn)
	_OutlineViewDidDragTableColumn func(outlineView IOutlineView, tableColumn ITableColumn)
	_OutlineViewDidRemoveRowViewForRow func(outlineView IOutlineView, rowView ITableRowView, row int)
	_OutlineViewHeightOfRowByItem func(outlineView IOutlineView, item objectivec.IObject) float64
	_OutlineViewIsGroupItem func(outlineView IOutlineView, item objectivec.IObject) bool
	_OutlineViewMouseDownInHeaderOfTableColumn func(outlineView IOutlineView, tableColumn ITableColumn)
	_OutlineViewNextTypeSelectMatchFromItemToItemForString func(outlineView IOutlineView, startItem objectivec.IObject, endItem objectivec.IObject, searchString foundation.foundation.INSString) objc.ID
	_OutlineViewRowViewForItem func(outlineView IOutlineView, item objectivec.IObject) ITableRowView
	_OutlineViewSelectionIndexesForProposedSelection func(outlineView IOutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet
	_OutlineViewShouldCollapseItem func(outlineView IOutlineView, item objectivec.IObject) bool
	_OutlineViewShouldEditTableColumnItem func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool
	_OutlineViewShouldExpandItem func(outlineView IOutlineView, item objectivec.IObject) bool
	_OutlineViewShouldReorderColumnToColumn func(outlineView IOutlineView, columnIndex int, newColumnIndex int) bool
	_OutlineViewShouldSelectTableColumn func(outlineView IOutlineView, tableColumn ITableColumn) bool
	_OutlineViewShouldSelectItem func(outlineView IOutlineView, item objectivec.IObject) bool
	_OutlineViewShouldShowCellExpansionForTableColumnItem func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool
	_OutlineViewShouldShowOutlineCellForItem func(outlineView IOutlineView, item objectivec.IObject) bool
	_OutlineViewShouldTrackCellForTableColumnItem func(outlineView IOutlineView, cell ICell, tableColumn ITableColumn, item objectivec.IObject) bool
	_OutlineViewShouldTypeSelectForEventWithCurrentSearchString func(outlineView IOutlineView, event IEvent, searchString foundation.foundation.INSString) bool
	_OutlineViewSizeToFitWidthOfColumn func(outlineView IOutlineView, column int) float64
	_OutlineViewTintConfigurationForItem func(outlineView IOutlineView, item objectivec.IObject) ITintConfiguration
	_OutlineViewToolTipForCellRectTableColumnItemMouseLocation func(outlineView IOutlineView, cell ICell, rect RectPointer /* not a class type */, tableColumn ITableColumn, item objectivec.IObject, mouseLocation corefoundation.CGPoint) foundation.String
	_OutlineViewTypeSelectStringForTableColumnItem func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) foundation.String
	_OutlineViewUserCanChangeVisibilityOfTableColumn func(outlineView IOutlineView, column ITableColumn) bool
	_OutlineViewUserDidChangeVisibilityOfTableColumns func(outlineView IOutlineView, columns []TableColumn)
	_OutlineViewViewForTableColumnItem func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) IView
	_OutlineViewWillDisplayCellForTableColumnItem func(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject)
	_OutlineViewWillDisplayOutlineCellForTableColumnItem func(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject)
	_OutlineViewColumnDidMove func(notification foundation.foundation.INSNotification)
	_OutlineViewColumnDidResize func(notification foundation.foundation.INSNotification)
	_OutlineViewItemDidCollapse func(notification foundation.foundation.INSNotification)
	_OutlineViewItemDidExpand func(notification foundation.foundation.INSNotification)
	_OutlineViewItemWillCollapse func(notification foundation.foundation.INSNotification)
	_OutlineViewItemWillExpand func(notification foundation.foundation.INSNotification)
	_OutlineViewSelectionDidChange func(notification foundation.foundation.INSNotification)
	_OutlineViewSelectionIsChanging func(notification foundation.foundation.INSNotification)
	_SelectionShouldChangeInOutlineView func(outlineView IOutlineView) bool
}

// SetOutlineViewDataCellForTableColumnItem sets the handler for the OutlineViewDataCellForTableColumnItem delegate method.
//
// Returns the cell to use in a given column for a given item.
func (d *OutlineViewDelegate) SetOutlineViewDataCellForTableColumnItem(f func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) ICell) {
	d._OutlineViewDataCellForTableColumnItem = f
}

// SetOutlineViewDidAddRowViewForRow sets the handler for the OutlineViewDidAddRowViewForRow delegate method.
//
// Implemented to know when a new row view is added to the table.
func (d *OutlineViewDelegate) SetOutlineViewDidAddRowViewForRow(f func(outlineView IOutlineView, rowView ITableRowView, row int)) {
	d._OutlineViewDidAddRowViewForRow = f
}

// SetOutlineViewDidClickTableColumn sets the handler for the OutlineViewDidClickTableColumn delegate method.
//
// Sent at the time the mouse button subsequently goes up in   and   has been “clicked” without having been dragged anywhere.
func (d *OutlineViewDelegate) SetOutlineViewDidClickTableColumn(f func(outlineView IOutlineView, tableColumn ITableColumn)) {
	d._OutlineViewDidClickTableColumn = f
}

// SetOutlineViewDidDragTableColumn sets the handler for the OutlineViewDidDragTableColumn delegate method.
//
// Sent at the time the mouse button goes up in   and   has been dragged during the time the mouse button was down.
func (d *OutlineViewDelegate) SetOutlineViewDidDragTableColumn(f func(outlineView IOutlineView, tableColumn ITableColumn)) {
	d._OutlineViewDidDragTableColumn = f
}

// SetOutlineViewDidRemoveRowViewForRow sets the handler for the OutlineViewDidRemoveRowViewForRow delegate method.
//
// Implemented to know when a row view is removed from the table
func (d *OutlineViewDelegate) SetOutlineViewDidRemoveRowViewForRow(f func(outlineView IOutlineView, rowView ITableRowView, row int)) {
	d._OutlineViewDidRemoveRowViewForRow = f
}

// SetOutlineViewHeightOfRowByItem sets the handler for the OutlineViewHeightOfRowByItem delegate method.
//
// Returns the height in points of the row containing  .
func (d *OutlineViewDelegate) SetOutlineViewHeightOfRowByItem(f func(outlineView IOutlineView, item objectivec.IObject) float64) {
	d._OutlineViewHeightOfRowByItem = f
}

// SetOutlineViewIsGroupItem sets the handler for the OutlineViewIsGroupItem delegate method.
//
// Returns a Boolean that indicates whether a given row should be drawn in the “group row” style.
func (d *OutlineViewDelegate) SetOutlineViewIsGroupItem(f func(outlineView IOutlineView, item objectivec.IObject) bool) {
	d._OutlineViewIsGroupItem = f
}

// SetOutlineViewMouseDownInHeaderOfTableColumn sets the handler for the OutlineViewMouseDownInHeaderOfTableColumn delegate method.
//
// Sent to the delegate whenever the mouse button is clicked in   while the cursor is in a column header  .
func (d *OutlineViewDelegate) SetOutlineViewMouseDownInHeaderOfTableColumn(f func(outlineView IOutlineView, tableColumn ITableColumn)) {
	d._OutlineViewMouseDownInHeaderOfTableColumn = f
}

// SetOutlineViewNextTypeSelectMatchFromItemToItemForString sets the handler for the OutlineViewNextTypeSelectMatchFromItemToItemForString delegate method.
//
// Returns the first item that matches the searchString from within the range of startItem to endItem
func (d *OutlineViewDelegate) SetOutlineViewNextTypeSelectMatchFromItemToItemForString(f func(outlineView IOutlineView, startItem objectivec.IObject, endItem objectivec.IObject, searchString foundation.foundation.INSString) objc.ID) {
	d._OutlineViewNextTypeSelectMatchFromItemToItemForString = f
}

// SetOutlineViewRowViewForItem sets the handler for the OutlineViewRowViewForItem delegate method.
//
// implement this method to return a custom   for a particular item.
func (d *OutlineViewDelegate) SetOutlineViewRowViewForItem(f func(outlineView IOutlineView, item objectivec.IObject) ITableRowView) {
	d._OutlineViewRowViewForItem = f
}

// SetOutlineViewSelectionIndexesForProposedSelection sets the handler for the OutlineViewSelectionIndexesForProposedSelection delegate method.
//
// Invoked to allow the delegate to modify the proposed selection.
func (d *OutlineViewDelegate) SetOutlineViewSelectionIndexesForProposedSelection(f func(outlineView IOutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet) {
	d._OutlineViewSelectionIndexesForProposedSelection = f
}

// SetOutlineViewShouldCollapseItem sets the handler for the OutlineViewShouldCollapseItem delegate method.
//
// Returns a Boolean value that indicates whether the outline view should collapse a given item.
func (d *OutlineViewDelegate) SetOutlineViewShouldCollapseItem(f func(outlineView IOutlineView, item objectivec.IObject) bool) {
	d._OutlineViewShouldCollapseItem = f
}

// SetOutlineViewShouldEditTableColumnItem sets the handler for the OutlineViewShouldEditTableColumnItem delegate method.
//
// Returns a Boolean value that indicates whether the outline view should allow editing of a given item in a given table column.
func (d *OutlineViewDelegate) SetOutlineViewShouldEditTableColumnItem(f func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool) {
	d._OutlineViewShouldEditTableColumnItem = f
}

// SetOutlineViewShouldExpandItem sets the handler for the OutlineViewShouldExpandItem delegate method.
//
// Returns a Boolean value that indicates whether the outline view should expand a given item.
func (d *OutlineViewDelegate) SetOutlineViewShouldExpandItem(f func(outlineView IOutlineView, item objectivec.IObject) bool) {
	d._OutlineViewShouldExpandItem = f
}

// SetOutlineViewShouldReorderColumnToColumn sets the handler for the OutlineViewShouldReorderColumnToColumn delegate method.
//
// Sent to the delegate to allow or prohibit the specified column to be dragged to a new location.
func (d *OutlineViewDelegate) SetOutlineViewShouldReorderColumnToColumn(f func(outlineView IOutlineView, columnIndex int, newColumnIndex int) bool) {
	d._OutlineViewShouldReorderColumnToColumn = f
}

// SetOutlineViewShouldSelectTableColumn sets the handler for the OutlineViewShouldSelectTableColumn delegate method.
//
// Returns a Boolean value that indicates whether the outline view should select a given table column.
func (d *OutlineViewDelegate) SetOutlineViewShouldSelectTableColumn(f func(outlineView IOutlineView, tableColumn ITableColumn) bool) {
	d._OutlineViewShouldSelectTableColumn = f
}

// SetOutlineViewShouldSelectItem sets the handler for the OutlineViewShouldSelectItem delegate method.
//
// Returns a Boolean value that indicates whether the outline view should select a given item.
func (d *OutlineViewDelegate) SetOutlineViewShouldSelectItem(f func(outlineView IOutlineView, item objectivec.IObject) bool) {
	d._OutlineViewShouldSelectItem = f
}

// SetOutlineViewShouldShowCellExpansionForTableColumnItem sets the handler for the OutlineViewShouldShowCellExpansionForTableColumnItem delegate method.
//
// Invoked to allow the delegate to control cell expansion for a specific column and item.
func (d *OutlineViewDelegate) SetOutlineViewShouldShowCellExpansionForTableColumnItem(f func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool) {
	d._OutlineViewShouldShowCellExpansionForTableColumnItem = f
}

// SetOutlineViewShouldShowOutlineCellForItem sets the handler for the OutlineViewShouldShowOutlineCellForItem delegate method.
//
// Returns whether the specified item should display the outline cell (the disclosure triangle).
func (d *OutlineViewDelegate) SetOutlineViewShouldShowOutlineCellForItem(f func(outlineView IOutlineView, item objectivec.IObject) bool) {
	d._OutlineViewShouldShowOutlineCellForItem = f
}

// SetOutlineViewShouldTrackCellForTableColumnItem sets the handler for the OutlineViewShouldTrackCellForTableColumnItem delegate method.
//
// Returns a Boolean value that indicates whether a given cell should be tracked.
func (d *OutlineViewDelegate) SetOutlineViewShouldTrackCellForTableColumnItem(f func(outlineView IOutlineView, cell ICell, tableColumn ITableColumn, item objectivec.IObject) bool) {
	d._OutlineViewShouldTrackCellForTableColumnItem = f
}

// SetOutlineViewShouldTypeSelectForEventWithCurrentSearchString sets the handler for the OutlineViewShouldTypeSelectForEventWithCurrentSearchString delegate method.
//
// Returns a Boolean value that indicates whether type select should proceed for a given event and search string.
func (d *OutlineViewDelegate) SetOutlineViewShouldTypeSelectForEventWithCurrentSearchString(f func(outlineView IOutlineView, event IEvent, searchString foundation.foundation.INSString) bool) {
	d._OutlineViewShouldTypeSelectForEventWithCurrentSearchString = f
}

// SetOutlineViewSizeToFitWidthOfColumn sets the handler for the OutlineViewSizeToFitWidthOfColumn delegate method.
//
// Invoked to allow the delegate to provide custom sizing behavior when a column’s resize divider is double clicked.
func (d *OutlineViewDelegate) SetOutlineViewSizeToFitWidthOfColumn(f func(outlineView IOutlineView, column int) float64) {
	d._OutlineViewSizeToFitWidthOfColumn = f
}

// SetOutlineViewTintConfigurationForItem sets the handler for the OutlineViewTintConfigurationForItem delegate method.
//
// Customizes an item’s tinting behavior.
func (d *OutlineViewDelegate) SetOutlineViewTintConfigurationForItem(f func(outlineView IOutlineView, item objectivec.IObject) ITintConfiguration) {
	d._OutlineViewTintConfigurationForItem = f
}

// SetOutlineViewToolTipForCellRectTableColumnItemMouseLocation sets the handler for the OutlineViewToolTipForCellRectTableColumnItemMouseLocation delegate method.
//
// When the cursor pauses over a given cell, the value returned from this method is displayed in a tooltip.
func (d *OutlineViewDelegate) SetOutlineViewToolTipForCellRectTableColumnItemMouseLocation(f func(outlineView IOutlineView, cell ICell, rect RectPointer /* not a class type */, tableColumn ITableColumn, item objectivec.IObject, mouseLocation corefoundation.CGPoint) foundation.String) {
	d._OutlineViewToolTipForCellRectTableColumnItemMouseLocation = f
}

// SetOutlineViewTypeSelectStringForTableColumnItem sets the handler for the OutlineViewTypeSelectStringForTableColumnItem delegate method.
//
// Returns the string that is used for type selection for a given column and item.
func (d *OutlineViewDelegate) SetOutlineViewTypeSelectStringForTableColumnItem(f func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) foundation.String) {
	d._OutlineViewTypeSelectStringForTableColumnItem = f
}

// SetOutlineViewUserCanChangeVisibilityOfTableColumn sets the handler for the OutlineViewUserCanChangeVisibilityOfTableColumn delegate method.
func (d *OutlineViewDelegate) SetOutlineViewUserCanChangeVisibilityOfTableColumn(f func(outlineView IOutlineView, column ITableColumn) bool) {
	d._OutlineViewUserCanChangeVisibilityOfTableColumn = f
}

// SetOutlineViewUserDidChangeVisibilityOfTableColumns sets the handler for the OutlineViewUserDidChangeVisibilityOfTableColumns delegate method.
func (d *OutlineViewDelegate) SetOutlineViewUserDidChangeVisibilityOfTableColumns(f func(outlineView IOutlineView, columns []TableColumn)) {
	d._OutlineViewUserDidChangeVisibilityOfTableColumns = f
}

// SetOutlineViewViewForTableColumnItem sets the handler for the OutlineViewViewForTableColumnItem delegate method.
//
// Implemented to return the view used to display the specified item and column.
func (d *OutlineViewDelegate) SetOutlineViewViewForTableColumnItem(f func(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) IView) {
	d._OutlineViewViewForTableColumnItem = f
}

// SetOutlineViewWillDisplayCellForTableColumnItem sets the handler for the OutlineViewWillDisplayCellForTableColumnItem delegate method.
//
// Informs the delegate that the cell specified by the column and item will be displayed.
func (d *OutlineViewDelegate) SetOutlineViewWillDisplayCellForTableColumnItem(f func(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject)) {
	d._OutlineViewWillDisplayCellForTableColumnItem = f
}

// SetOutlineViewWillDisplayOutlineCellForTableColumnItem sets the handler for the OutlineViewWillDisplayOutlineCellForTableColumnItem delegate method.
//
// Informs the delegate that an outline view is about to display a cell used to draw the expansion symbol.
func (d *OutlineViewDelegate) SetOutlineViewWillDisplayOutlineCellForTableColumnItem(f func(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject)) {
	d._OutlineViewWillDisplayOutlineCellForTableColumnItem = f
}

// SetOutlineViewColumnDidMove sets the handler for the OutlineViewColumnDidMove delegate method.
//
// Invoked whenever the user moves a column in the outline view.
func (d *OutlineViewDelegate) SetOutlineViewColumnDidMove(f func(notification foundation.foundation.INSNotification)) {
	d._OutlineViewColumnDidMove = f
}

// SetOutlineViewColumnDidResize sets the handler for the OutlineViewColumnDidResize delegate method.
//
// Invoked whenever the user resizes a column in the outline view.
func (d *OutlineViewDelegate) SetOutlineViewColumnDidResize(f func(notification foundation.foundation.INSNotification)) {
	d._OutlineViewColumnDidResize = f
}

// SetOutlineViewItemDidCollapse sets the handler for the OutlineViewItemDidCollapse delegate method.
//
// Invoked when the did collapse notification is posted—that is, whenever the user collapses an item in the outline view.
func (d *OutlineViewDelegate) SetOutlineViewItemDidCollapse(f func(notification foundation.foundation.INSNotification)) {
	d._OutlineViewItemDidCollapse = f
}

// SetOutlineViewItemDidExpand sets the handler for the OutlineViewItemDidExpand delegate method.
//
// Invoked when   is posted—that is, whenever the user expands an item in the outline view.
func (d *OutlineViewDelegate) SetOutlineViewItemDidExpand(f func(notification foundation.foundation.INSNotification)) {
	d._OutlineViewItemDidExpand = f
}

// SetOutlineViewItemWillCollapse sets the handler for the OutlineViewItemWillCollapse delegate method.
//
// Invoked when   is posted—that is, whenever the user is about to collapse an item in the outline view.
func (d *OutlineViewDelegate) SetOutlineViewItemWillCollapse(f func(notification foundation.foundation.INSNotification)) {
	d._OutlineViewItemWillCollapse = f
}

// SetOutlineViewItemWillExpand sets the handler for the OutlineViewItemWillExpand delegate method.
//
// Invoked when   is posted—that is, whenever the user is about to expand an item in the outline view.
func (d *OutlineViewDelegate) SetOutlineViewItemWillExpand(f func(notification foundation.foundation.INSNotification)) {
	d._OutlineViewItemWillExpand = f
}

// SetOutlineViewSelectionDidChange sets the handler for the OutlineViewSelectionDidChange delegate method.
//
// Invoked when the selection did change notification is posted—that is, immediately after the outline view’s selection has changed.
func (d *OutlineViewDelegate) SetOutlineViewSelectionDidChange(f func(notification foundation.foundation.INSNotification)) {
	d._OutlineViewSelectionDidChange = f
}

// SetOutlineViewSelectionIsChanging sets the handler for the OutlineViewSelectionIsChanging delegate method.
//
// Invoked when   is posted—that is, whenever the outline view’s selection changes.
func (d *OutlineViewDelegate) SetOutlineViewSelectionIsChanging(f func(notification foundation.foundation.INSNotification)) {
	d._OutlineViewSelectionIsChanging = f
}

// SetSelectionShouldChangeInOutlineView sets the handler for the SelectionShouldChangeInOutlineView delegate method.
//
// Returns a Boolean value that indicates whether the outline view should change its selection.
func (d *OutlineViewDelegate) SetSelectionShouldChangeInOutlineView(f func(outlineView IOutlineView) bool) {
	d._SelectionShouldChangeInOutlineView = f
}

// OutlineViewDataCellForTableColumnItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewDataCellForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) ICell {
	if d._OutlineViewDataCellForTableColumnItem != nil {
		return d._OutlineViewDataCellForTableColumnItem(outlineView, tableColumn, item)
	}
	var zero ICell
	return zero
}

// HasOutlineViewDataCellForTableColumnItem returns true if a handler for OutlineViewDataCellForTableColumnItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewDataCellForTableColumnItem() bool {
	return d._OutlineViewDataCellForTableColumnItem != nil
}

// OutlineViewDidAddRowViewForRow implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewDidAddRowViewForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	if d._OutlineViewDidAddRowViewForRow != nil {
		d._OutlineViewDidAddRowViewForRow(outlineView, rowView, row)
	}
}

// HasOutlineViewDidAddRowViewForRow returns true if a handler for OutlineViewDidAddRowViewForRow has been set.
func (d *OutlineViewDelegate) HasOutlineViewDidAddRowViewForRow() bool {
	return d._OutlineViewDidAddRowViewForRow != nil
}

// OutlineViewDidClickTableColumn implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewDidClickTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	if d._OutlineViewDidClickTableColumn != nil {
		d._OutlineViewDidClickTableColumn(outlineView, tableColumn)
	}
}

// HasOutlineViewDidClickTableColumn returns true if a handler for OutlineViewDidClickTableColumn has been set.
func (d *OutlineViewDelegate) HasOutlineViewDidClickTableColumn() bool {
	return d._OutlineViewDidClickTableColumn != nil
}

// OutlineViewDidDragTableColumn implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewDidDragTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	if d._OutlineViewDidDragTableColumn != nil {
		d._OutlineViewDidDragTableColumn(outlineView, tableColumn)
	}
}

// HasOutlineViewDidDragTableColumn returns true if a handler for OutlineViewDidDragTableColumn has been set.
func (d *OutlineViewDelegate) HasOutlineViewDidDragTableColumn() bool {
	return d._OutlineViewDidDragTableColumn != nil
}

// OutlineViewDidRemoveRowViewForRow implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewDidRemoveRowViewForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	if d._OutlineViewDidRemoveRowViewForRow != nil {
		d._OutlineViewDidRemoveRowViewForRow(outlineView, rowView, row)
	}
}

// HasOutlineViewDidRemoveRowViewForRow returns true if a handler for OutlineViewDidRemoveRowViewForRow has been set.
func (d *OutlineViewDelegate) HasOutlineViewDidRemoveRowViewForRow() bool {
	return d._OutlineViewDidRemoveRowViewForRow != nil
}

// OutlineViewHeightOfRowByItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewHeightOfRowByItem(outlineView IOutlineView, item objectivec.IObject) float64 {
	if d._OutlineViewHeightOfRowByItem != nil {
		return d._OutlineViewHeightOfRowByItem(outlineView, item)
	}
	var zero float64
	return zero
}

// HasOutlineViewHeightOfRowByItem returns true if a handler for OutlineViewHeightOfRowByItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewHeightOfRowByItem() bool {
	return d._OutlineViewHeightOfRowByItem != nil
}

// OutlineViewIsGroupItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewIsGroupItem(outlineView IOutlineView, item objectivec.IObject) bool {
	if d._OutlineViewIsGroupItem != nil {
		return d._OutlineViewIsGroupItem(outlineView, item)
	}
	var zero bool
	return zero
}

// HasOutlineViewIsGroupItem returns true if a handler for OutlineViewIsGroupItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewIsGroupItem() bool {
	return d._OutlineViewIsGroupItem != nil
}

// OutlineViewMouseDownInHeaderOfTableColumn implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewMouseDownInHeaderOfTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	if d._OutlineViewMouseDownInHeaderOfTableColumn != nil {
		d._OutlineViewMouseDownInHeaderOfTableColumn(outlineView, tableColumn)
	}
}

// HasOutlineViewMouseDownInHeaderOfTableColumn returns true if a handler for OutlineViewMouseDownInHeaderOfTableColumn has been set.
func (d *OutlineViewDelegate) HasOutlineViewMouseDownInHeaderOfTableColumn() bool {
	return d._OutlineViewMouseDownInHeaderOfTableColumn != nil
}

// OutlineViewNextTypeSelectMatchFromItemToItemForString implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewNextTypeSelectMatchFromItemToItemForString(outlineView IOutlineView, startItem objectivec.IObject, endItem objectivec.IObject, searchString foundation.foundation.INSString) objc.ID {
	if d._OutlineViewNextTypeSelectMatchFromItemToItemForString != nil {
		return d._OutlineViewNextTypeSelectMatchFromItemToItemForString(outlineView, startItem, endItem, searchString)
	}
	var zero objc.ID
	return zero
}

// HasOutlineViewNextTypeSelectMatchFromItemToItemForString returns true if a handler for OutlineViewNextTypeSelectMatchFromItemToItemForString has been set.
func (d *OutlineViewDelegate) HasOutlineViewNextTypeSelectMatchFromItemToItemForString() bool {
	return d._OutlineViewNextTypeSelectMatchFromItemToItemForString != nil
}

// OutlineViewRowViewForItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewRowViewForItem(outlineView IOutlineView, item objectivec.IObject) ITableRowView {
	if d._OutlineViewRowViewForItem != nil {
		return d._OutlineViewRowViewForItem(outlineView, item)
	}
	var zero ITableRowView
	return zero
}

// HasOutlineViewRowViewForItem returns true if a handler for OutlineViewRowViewForItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewRowViewForItem() bool {
	return d._OutlineViewRowViewForItem != nil
}

// OutlineViewSelectionIndexesForProposedSelection implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewSelectionIndexesForProposedSelection(outlineView IOutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet {
	if d._OutlineViewSelectionIndexesForProposedSelection != nil {
		return d._OutlineViewSelectionIndexesForProposedSelection(outlineView, proposedSelectionIndexes)
	}
	var zero foundation.IndexSet
	return zero
}

// HasOutlineViewSelectionIndexesForProposedSelection returns true if a handler for OutlineViewSelectionIndexesForProposedSelection has been set.
func (d *OutlineViewDelegate) HasOutlineViewSelectionIndexesForProposedSelection() bool {
	return d._OutlineViewSelectionIndexesForProposedSelection != nil
}

// OutlineViewShouldCollapseItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldCollapseItem(outlineView IOutlineView, item objectivec.IObject) bool {
	if d._OutlineViewShouldCollapseItem != nil {
		return d._OutlineViewShouldCollapseItem(outlineView, item)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldCollapseItem returns true if a handler for OutlineViewShouldCollapseItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldCollapseItem() bool {
	return d._OutlineViewShouldCollapseItem != nil
}

// OutlineViewShouldEditTableColumnItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldEditTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool {
	if d._OutlineViewShouldEditTableColumnItem != nil {
		return d._OutlineViewShouldEditTableColumnItem(outlineView, tableColumn, item)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldEditTableColumnItem returns true if a handler for OutlineViewShouldEditTableColumnItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldEditTableColumnItem() bool {
	return d._OutlineViewShouldEditTableColumnItem != nil
}

// OutlineViewShouldExpandItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldExpandItem(outlineView IOutlineView, item objectivec.IObject) bool {
	if d._OutlineViewShouldExpandItem != nil {
		return d._OutlineViewShouldExpandItem(outlineView, item)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldExpandItem returns true if a handler for OutlineViewShouldExpandItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldExpandItem() bool {
	return d._OutlineViewShouldExpandItem != nil
}

// OutlineViewShouldReorderColumnToColumn implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldReorderColumnToColumn(outlineView IOutlineView, columnIndex int, newColumnIndex int) bool {
	if d._OutlineViewShouldReorderColumnToColumn != nil {
		return d._OutlineViewShouldReorderColumnToColumn(outlineView, columnIndex, newColumnIndex)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldReorderColumnToColumn returns true if a handler for OutlineViewShouldReorderColumnToColumn has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldReorderColumnToColumn() bool {
	return d._OutlineViewShouldReorderColumnToColumn != nil
}

// OutlineViewShouldSelectTableColumn implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldSelectTableColumn(outlineView IOutlineView, tableColumn ITableColumn) bool {
	if d._OutlineViewShouldSelectTableColumn != nil {
		return d._OutlineViewShouldSelectTableColumn(outlineView, tableColumn)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldSelectTableColumn returns true if a handler for OutlineViewShouldSelectTableColumn has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldSelectTableColumn() bool {
	return d._OutlineViewShouldSelectTableColumn != nil
}

// OutlineViewShouldSelectItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldSelectItem(outlineView IOutlineView, item objectivec.IObject) bool {
	if d._OutlineViewShouldSelectItem != nil {
		return d._OutlineViewShouldSelectItem(outlineView, item)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldSelectItem returns true if a handler for OutlineViewShouldSelectItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldSelectItem() bool {
	return d._OutlineViewShouldSelectItem != nil
}

// OutlineViewShouldShowCellExpansionForTableColumnItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldShowCellExpansionForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool {
	if d._OutlineViewShouldShowCellExpansionForTableColumnItem != nil {
		return d._OutlineViewShouldShowCellExpansionForTableColumnItem(outlineView, tableColumn, item)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldShowCellExpansionForTableColumnItem returns true if a handler for OutlineViewShouldShowCellExpansionForTableColumnItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldShowCellExpansionForTableColumnItem() bool {
	return d._OutlineViewShouldShowCellExpansionForTableColumnItem != nil
}

// OutlineViewShouldShowOutlineCellForItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldShowOutlineCellForItem(outlineView IOutlineView, item objectivec.IObject) bool {
	if d._OutlineViewShouldShowOutlineCellForItem != nil {
		return d._OutlineViewShouldShowOutlineCellForItem(outlineView, item)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldShowOutlineCellForItem returns true if a handler for OutlineViewShouldShowOutlineCellForItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldShowOutlineCellForItem() bool {
	return d._OutlineViewShouldShowOutlineCellForItem != nil
}

// OutlineViewShouldTrackCellForTableColumnItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldTrackCellForTableColumnItem(outlineView IOutlineView, cell ICell, tableColumn ITableColumn, item objectivec.IObject) bool {
	if d._OutlineViewShouldTrackCellForTableColumnItem != nil {
		return d._OutlineViewShouldTrackCellForTableColumnItem(outlineView, cell, tableColumn, item)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldTrackCellForTableColumnItem returns true if a handler for OutlineViewShouldTrackCellForTableColumnItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldTrackCellForTableColumnItem() bool {
	return d._OutlineViewShouldTrackCellForTableColumnItem != nil
}

// OutlineViewShouldTypeSelectForEventWithCurrentSearchString implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewShouldTypeSelectForEventWithCurrentSearchString(outlineView IOutlineView, event IEvent, searchString foundation.foundation.INSString) bool {
	if d._OutlineViewShouldTypeSelectForEventWithCurrentSearchString != nil {
		return d._OutlineViewShouldTypeSelectForEventWithCurrentSearchString(outlineView, event, searchString)
	}
	var zero bool
	return zero
}

// HasOutlineViewShouldTypeSelectForEventWithCurrentSearchString returns true if a handler for OutlineViewShouldTypeSelectForEventWithCurrentSearchString has been set.
func (d *OutlineViewDelegate) HasOutlineViewShouldTypeSelectForEventWithCurrentSearchString() bool {
	return d._OutlineViewShouldTypeSelectForEventWithCurrentSearchString != nil
}

// OutlineViewSizeToFitWidthOfColumn implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewSizeToFitWidthOfColumn(outlineView IOutlineView, column int) float64 {
	if d._OutlineViewSizeToFitWidthOfColumn != nil {
		return d._OutlineViewSizeToFitWidthOfColumn(outlineView, column)
	}
	var zero float64
	return zero
}

// HasOutlineViewSizeToFitWidthOfColumn returns true if a handler for OutlineViewSizeToFitWidthOfColumn has been set.
func (d *OutlineViewDelegate) HasOutlineViewSizeToFitWidthOfColumn() bool {
	return d._OutlineViewSizeToFitWidthOfColumn != nil
}

// OutlineViewTintConfigurationForItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewTintConfigurationForItem(outlineView IOutlineView, item objectivec.IObject) ITintConfiguration {
	if d._OutlineViewTintConfigurationForItem != nil {
		return d._OutlineViewTintConfigurationForItem(outlineView, item)
	}
	var zero ITintConfiguration
	return zero
}

// HasOutlineViewTintConfigurationForItem returns true if a handler for OutlineViewTintConfigurationForItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewTintConfigurationForItem() bool {
	return d._OutlineViewTintConfigurationForItem != nil
}

// OutlineViewToolTipForCellRectTableColumnItemMouseLocation implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewToolTipForCellRectTableColumnItemMouseLocation(outlineView IOutlineView, cell ICell, rect RectPointer /* not a class type */, tableColumn ITableColumn, item objectivec.IObject, mouseLocation corefoundation.CGPoint) foundation.String {
	if d._OutlineViewToolTipForCellRectTableColumnItemMouseLocation != nil {
		return d._OutlineViewToolTipForCellRectTableColumnItemMouseLocation(outlineView, cell, rect, tableColumn, item, mouseLocation)
	}
	var zero foundation.String
	return zero
}

// HasOutlineViewToolTipForCellRectTableColumnItemMouseLocation returns true if a handler for OutlineViewToolTipForCellRectTableColumnItemMouseLocation has been set.
func (d *OutlineViewDelegate) HasOutlineViewToolTipForCellRectTableColumnItemMouseLocation() bool {
	return d._OutlineViewToolTipForCellRectTableColumnItemMouseLocation != nil
}

// OutlineViewTypeSelectStringForTableColumnItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewTypeSelectStringForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) foundation.String {
	if d._OutlineViewTypeSelectStringForTableColumnItem != nil {
		return d._OutlineViewTypeSelectStringForTableColumnItem(outlineView, tableColumn, item)
	}
	var zero foundation.String
	return zero
}

// HasOutlineViewTypeSelectStringForTableColumnItem returns true if a handler for OutlineViewTypeSelectStringForTableColumnItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewTypeSelectStringForTableColumnItem() bool {
	return d._OutlineViewTypeSelectStringForTableColumnItem != nil
}

// OutlineViewUserCanChangeVisibilityOfTableColumn implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewUserCanChangeVisibilityOfTableColumn(outlineView IOutlineView, column ITableColumn) bool {
	if d._OutlineViewUserCanChangeVisibilityOfTableColumn != nil {
		return d._OutlineViewUserCanChangeVisibilityOfTableColumn(outlineView, column)
	}
	var zero bool
	return zero
}

// HasOutlineViewUserCanChangeVisibilityOfTableColumn returns true if a handler for OutlineViewUserCanChangeVisibilityOfTableColumn has been set.
func (d *OutlineViewDelegate) HasOutlineViewUserCanChangeVisibilityOfTableColumn() bool {
	return d._OutlineViewUserCanChangeVisibilityOfTableColumn != nil
}

// OutlineViewUserDidChangeVisibilityOfTableColumns implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewUserDidChangeVisibilityOfTableColumns(outlineView IOutlineView, columns []TableColumn) {
	if d._OutlineViewUserDidChangeVisibilityOfTableColumns != nil {
		d._OutlineViewUserDidChangeVisibilityOfTableColumns(outlineView, columns)
	}
}

// HasOutlineViewUserDidChangeVisibilityOfTableColumns returns true if a handler for OutlineViewUserDidChangeVisibilityOfTableColumns has been set.
func (d *OutlineViewDelegate) HasOutlineViewUserDidChangeVisibilityOfTableColumns() bool {
	return d._OutlineViewUserDidChangeVisibilityOfTableColumns != nil
}

// OutlineViewViewForTableColumnItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewViewForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) IView {
	if d._OutlineViewViewForTableColumnItem != nil {
		return d._OutlineViewViewForTableColumnItem(outlineView, tableColumn, item)
	}
	var zero IView
	return zero
}

// HasOutlineViewViewForTableColumnItem returns true if a handler for OutlineViewViewForTableColumnItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewViewForTableColumnItem() bool {
	return d._OutlineViewViewForTableColumnItem != nil
}

// OutlineViewWillDisplayCellForTableColumnItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewWillDisplayCellForTableColumnItem(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject) {
	if d._OutlineViewWillDisplayCellForTableColumnItem != nil {
		d._OutlineViewWillDisplayCellForTableColumnItem(outlineView, cell, tableColumn, item)
	}
}

// HasOutlineViewWillDisplayCellForTableColumnItem returns true if a handler for OutlineViewWillDisplayCellForTableColumnItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewWillDisplayCellForTableColumnItem() bool {
	return d._OutlineViewWillDisplayCellForTableColumnItem != nil
}

// OutlineViewWillDisplayOutlineCellForTableColumnItem implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewWillDisplayOutlineCellForTableColumnItem(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject) {
	if d._OutlineViewWillDisplayOutlineCellForTableColumnItem != nil {
		d._OutlineViewWillDisplayOutlineCellForTableColumnItem(outlineView, cell, tableColumn, item)
	}
}

// HasOutlineViewWillDisplayOutlineCellForTableColumnItem returns true if a handler for OutlineViewWillDisplayOutlineCellForTableColumnItem has been set.
func (d *OutlineViewDelegate) HasOutlineViewWillDisplayOutlineCellForTableColumnItem() bool {
	return d._OutlineViewWillDisplayOutlineCellForTableColumnItem != nil
}

// OutlineViewColumnDidMove implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewColumnDidMove(notification foundation.foundation.INSNotification) {
	if d._OutlineViewColumnDidMove != nil {
		d._OutlineViewColumnDidMove(notification)
	}
}

// HasOutlineViewColumnDidMove returns true if a handler for OutlineViewColumnDidMove has been set.
func (d *OutlineViewDelegate) HasOutlineViewColumnDidMove() bool {
	return d._OutlineViewColumnDidMove != nil
}

// OutlineViewColumnDidResize implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewColumnDidResize(notification foundation.foundation.INSNotification) {
	if d._OutlineViewColumnDidResize != nil {
		d._OutlineViewColumnDidResize(notification)
	}
}

// HasOutlineViewColumnDidResize returns true if a handler for OutlineViewColumnDidResize has been set.
func (d *OutlineViewDelegate) HasOutlineViewColumnDidResize() bool {
	return d._OutlineViewColumnDidResize != nil
}

// OutlineViewItemDidCollapse implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewItemDidCollapse(notification foundation.foundation.INSNotification) {
	if d._OutlineViewItemDidCollapse != nil {
		d._OutlineViewItemDidCollapse(notification)
	}
}

// HasOutlineViewItemDidCollapse returns true if a handler for OutlineViewItemDidCollapse has been set.
func (d *OutlineViewDelegate) HasOutlineViewItemDidCollapse() bool {
	return d._OutlineViewItemDidCollapse != nil
}

// OutlineViewItemDidExpand implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewItemDidExpand(notification foundation.foundation.INSNotification) {
	if d._OutlineViewItemDidExpand != nil {
		d._OutlineViewItemDidExpand(notification)
	}
}

// HasOutlineViewItemDidExpand returns true if a handler for OutlineViewItemDidExpand has been set.
func (d *OutlineViewDelegate) HasOutlineViewItemDidExpand() bool {
	return d._OutlineViewItemDidExpand != nil
}

// OutlineViewItemWillCollapse implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewItemWillCollapse(notification foundation.foundation.INSNotification) {
	if d._OutlineViewItemWillCollapse != nil {
		d._OutlineViewItemWillCollapse(notification)
	}
}

// HasOutlineViewItemWillCollapse returns true if a handler for OutlineViewItemWillCollapse has been set.
func (d *OutlineViewDelegate) HasOutlineViewItemWillCollapse() bool {
	return d._OutlineViewItemWillCollapse != nil
}

// OutlineViewItemWillExpand implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewItemWillExpand(notification foundation.foundation.INSNotification) {
	if d._OutlineViewItemWillExpand != nil {
		d._OutlineViewItemWillExpand(notification)
	}
}

// HasOutlineViewItemWillExpand returns true if a handler for OutlineViewItemWillExpand has been set.
func (d *OutlineViewDelegate) HasOutlineViewItemWillExpand() bool {
	return d._OutlineViewItemWillExpand != nil
}

// OutlineViewSelectionDidChange implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewSelectionDidChange(notification foundation.foundation.INSNotification) {
	if d._OutlineViewSelectionDidChange != nil {
		d._OutlineViewSelectionDidChange(notification)
	}
}

// HasOutlineViewSelectionDidChange returns true if a handler for OutlineViewSelectionDidChange has been set.
func (d *OutlineViewDelegate) HasOutlineViewSelectionDidChange() bool {
	return d._OutlineViewSelectionDidChange != nil
}

// OutlineViewSelectionIsChanging implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) OutlineViewSelectionIsChanging(notification foundation.foundation.INSNotification) {
	if d._OutlineViewSelectionIsChanging != nil {
		d._OutlineViewSelectionIsChanging(notification)
	}
}

// HasOutlineViewSelectionIsChanging returns true if a handler for OutlineViewSelectionIsChanging has been set.
func (d *OutlineViewDelegate) HasOutlineViewSelectionIsChanging() bool {
	return d._OutlineViewSelectionIsChanging != nil
}

// SelectionShouldChangeInOutlineView implements the POutlineViewDelegate interface.
func (d *OutlineViewDelegate) SelectionShouldChangeInOutlineView(outlineView IOutlineView) bool {
	if d._SelectionShouldChangeInOutlineView != nil {
		return d._SelectionShouldChangeInOutlineView(outlineView)
	}
	var zero bool
	return zero
}

// HasSelectionShouldChangeInOutlineView returns true if a handler for SelectionShouldChangeInOutlineView has been set.
func (d *OutlineViewDelegate) HasSelectionShouldChangeInOutlineView() bool {
	return d._SelectionShouldChangeInOutlineView != nil
}

// OutlineViewDelegateObject wraps an existing Objective-C object that conforms to the POutlineViewDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type OutlineViewDelegateObject struct {
	objectivec.Object
}

// NewOutlineViewDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSOutlineViewDelegate protocol.
func NewOutlineViewDelegateObject(obj objectivec.Object) *OutlineViewDelegateObject {
	return &OutlineViewDelegateObject{obj}
}

// Make sure OutlineViewDelegateObject implements POutlineViewDelegate.
var _ POutlineViewDelegate = (*OutlineViewDelegateObject)(nil)

// OutlineViewDataCellForTableColumnItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewDataCellForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) ICell {
	return objc.Send[ICell](o.ID, objc.Sel("outlineView:dataCellForTableColumn:item:"), outlineView, tableColumn, item)
}

// HasOutlineViewDataCellForTableColumnItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewDataCellForTableColumnItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewDidAddRowViewForRow implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewDidAddRowViewForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineView:didAddRowView:forRow:"), outlineView, rowView, row)
}

// HasOutlineViewDidAddRowViewForRow returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewDidAddRowViewForRow() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewDidClickTableColumn implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewDidClickTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineView:didClickTableColumn:"), outlineView, tableColumn)
}

// HasOutlineViewDidClickTableColumn returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewDidClickTableColumn() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewDidDragTableColumn implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewDidDragTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineView:didDragTableColumn:"), outlineView, tableColumn)
}

// HasOutlineViewDidDragTableColumn returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewDidDragTableColumn() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewDidRemoveRowViewForRow implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewDidRemoveRowViewForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineView:didRemoveRowView:forRow:"), outlineView, rowView, row)
}

// HasOutlineViewDidRemoveRowViewForRow returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewDidRemoveRowViewForRow() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewHeightOfRowByItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewHeightOfRowByItem(outlineView IOutlineView, item objectivec.IObject) float64 {
	return objc.Send[float64](o.ID, objc.Sel("outlineView:heightOfRowByItem:"), outlineView, item)
}

// HasOutlineViewHeightOfRowByItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewHeightOfRowByItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewIsGroupItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewIsGroupItem(outlineView IOutlineView, item objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:isGroupItem:"), outlineView, item)
}

// HasOutlineViewIsGroupItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewIsGroupItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewMouseDownInHeaderOfTableColumn implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewMouseDownInHeaderOfTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineView:mouseDownInHeaderOfTableColumn:"), outlineView, tableColumn)
}

// HasOutlineViewMouseDownInHeaderOfTableColumn returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewMouseDownInHeaderOfTableColumn() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewNextTypeSelectMatchFromItemToItemForString implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewNextTypeSelectMatchFromItemToItemForString(outlineView IOutlineView, startItem objectivec.IObject, endItem objectivec.IObject, searchString foundation.foundation.INSString) objc.ID {
	return objc.Send[objc.ID](o.ID, objc.Sel("outlineView:nextTypeSelectMatchFromItem:toItem:forString:"), outlineView, startItem, endItem, searchString)
}

// HasOutlineViewNextTypeSelectMatchFromItemToItemForString returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewNextTypeSelectMatchFromItemToItemForString() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewRowViewForItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewRowViewForItem(outlineView IOutlineView, item objectivec.IObject) ITableRowView {
	return objc.Send[ITableRowView](o.ID, objc.Sel("outlineView:rowViewForItem:"), outlineView, item)
}

// HasOutlineViewRowViewForItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewRowViewForItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewSelectionIndexesForProposedSelection implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewSelectionIndexesForProposedSelection(outlineView IOutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IndexSet {
	return objc.Send[foundation.IndexSet](o.ID, objc.Sel("outlineView:selectionIndexesForProposedSelection:"), outlineView, proposedSelectionIndexes)
}

// HasOutlineViewSelectionIndexesForProposedSelection returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewSelectionIndexesForProposedSelection() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldCollapseItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldCollapseItem(outlineView IOutlineView, item objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldCollapseItem:"), outlineView, item)
}

// HasOutlineViewShouldCollapseItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldCollapseItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldEditTableColumnItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldEditTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldEditTableColumn:item:"), outlineView, tableColumn, item)
}

// HasOutlineViewShouldEditTableColumnItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldEditTableColumnItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldExpandItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldExpandItem(outlineView IOutlineView, item objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldExpandItem:"), outlineView, item)
}

// HasOutlineViewShouldExpandItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldExpandItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldReorderColumnToColumn implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldReorderColumnToColumn(outlineView IOutlineView, columnIndex int, newColumnIndex int) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldReorderColumn:toColumn:"), outlineView, columnIndex, newColumnIndex)
}

// HasOutlineViewShouldReorderColumnToColumn returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldReorderColumnToColumn() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldSelectTableColumn implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldSelectTableColumn(outlineView IOutlineView, tableColumn ITableColumn) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldSelectTableColumn:"), outlineView, tableColumn)
}

// HasOutlineViewShouldSelectTableColumn returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldSelectTableColumn() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldSelectItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldSelectItem(outlineView IOutlineView, item objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldSelectItem:"), outlineView, item)
}

// HasOutlineViewShouldSelectItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldSelectItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldShowCellExpansionForTableColumnItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldShowCellExpansionForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldShowCellExpansionForTableColumn:item:"), outlineView, tableColumn, item)
}

// HasOutlineViewShouldShowCellExpansionForTableColumnItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldShowCellExpansionForTableColumnItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldShowOutlineCellForItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldShowOutlineCellForItem(outlineView IOutlineView, item objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldShowOutlineCellForItem:"), outlineView, item)
}

// HasOutlineViewShouldShowOutlineCellForItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldShowOutlineCellForItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldTrackCellForTableColumnItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldTrackCellForTableColumnItem(outlineView IOutlineView, cell ICell, tableColumn ITableColumn, item objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldTrackCell:forTableColumn:item:"), outlineView, cell, tableColumn, item)
}

// HasOutlineViewShouldTrackCellForTableColumnItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldTrackCellForTableColumnItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewShouldTypeSelectForEventWithCurrentSearchString implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewShouldTypeSelectForEventWithCurrentSearchString(outlineView IOutlineView, event IEvent, searchString foundation.foundation.INSString) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:shouldTypeSelectForEvent:withCurrentSearchString:"), outlineView, event, searchString)
}

// HasOutlineViewShouldTypeSelectForEventWithCurrentSearchString returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewShouldTypeSelectForEventWithCurrentSearchString() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewSizeToFitWidthOfColumn implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewSizeToFitWidthOfColumn(outlineView IOutlineView, column int) float64 {
	return objc.Send[float64](o.ID, objc.Sel("outlineView:sizeToFitWidthOfColumn:"), outlineView, column)
}

// HasOutlineViewSizeToFitWidthOfColumn returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewSizeToFitWidthOfColumn() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewTintConfigurationForItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewTintConfigurationForItem(outlineView IOutlineView, item objectivec.IObject) ITintConfiguration {
	return objc.Send[ITintConfiguration](o.ID, objc.Sel("outlineView:tintConfigurationForItem:"), outlineView, item)
}

// HasOutlineViewTintConfigurationForItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewTintConfigurationForItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewToolTipForCellRectTableColumnItemMouseLocation implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewToolTipForCellRectTableColumnItemMouseLocation(outlineView IOutlineView, cell ICell, rect RectPointer /* not a class type */, tableColumn ITableColumn, item objectivec.IObject, mouseLocation corefoundation.CGPoint) foundation.String {
	return objc.Send[foundation.String](o.ID, objc.Sel("outlineView:toolTipForCell:rect:tableColumn:item:mouseLocation:"), outlineView, cell, rect, tableColumn, item, mouseLocation)
}

// HasOutlineViewToolTipForCellRectTableColumnItemMouseLocation returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewToolTipForCellRectTableColumnItemMouseLocation() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewTypeSelectStringForTableColumnItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewTypeSelectStringForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) foundation.String {
	return objc.Send[foundation.String](o.ID, objc.Sel("outlineView:typeSelectStringForTableColumn:item:"), outlineView, tableColumn, item)
}

// HasOutlineViewTypeSelectStringForTableColumnItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewTypeSelectStringForTableColumnItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewUserCanChangeVisibilityOfTableColumn implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewUserCanChangeVisibilityOfTableColumn(outlineView IOutlineView, column ITableColumn) bool {
	return objc.Send[bool](o.ID, objc.Sel("outlineView:userCanChangeVisibilityOfTableColumn:"), outlineView, column)
}

// HasOutlineViewUserCanChangeVisibilityOfTableColumn returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewUserCanChangeVisibilityOfTableColumn() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewUserDidChangeVisibilityOfTableColumns implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewUserDidChangeVisibilityOfTableColumns(outlineView IOutlineView, columns []TableColumn) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineView:userDidChangeVisibilityOfTableColumns:"), outlineView, columns)
}

// HasOutlineViewUserDidChangeVisibilityOfTableColumns returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewUserDidChangeVisibilityOfTableColumns() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewViewForTableColumnItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewViewForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objectivec.IObject) IView {
	return objc.Send[IView](o.ID, objc.Sel("outlineView:viewForTableColumn:item:"), outlineView, tableColumn, item)
}

// HasOutlineViewViewForTableColumnItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewViewForTableColumnItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewWillDisplayCellForTableColumnItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewWillDisplayCellForTableColumnItem(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineView:willDisplayCell:forTableColumn:item:"), outlineView, cell, tableColumn, item)
}

// HasOutlineViewWillDisplayCellForTableColumnItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewWillDisplayCellForTableColumnItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewWillDisplayOutlineCellForTableColumnItem implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewWillDisplayOutlineCellForTableColumnItem(outlineView IOutlineView, cell objectivec.IObject, tableColumn ITableColumn, item objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineView:willDisplayOutlineCell:forTableColumn:item:"), outlineView, cell, tableColumn, item)
}

// HasOutlineViewWillDisplayOutlineCellForTableColumnItem returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewWillDisplayOutlineCellForTableColumnItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewColumnDidMove implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewColumnDidMove(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineViewColumnDidMove:"), notification)
}

// HasOutlineViewColumnDidMove returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewColumnDidMove() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewColumnDidResize implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewColumnDidResize(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineViewColumnDidResize:"), notification)
}

// HasOutlineViewColumnDidResize returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewColumnDidResize() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewItemDidCollapse implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewItemDidCollapse(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineViewItemDidCollapse:"), notification)
}

// HasOutlineViewItemDidCollapse returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewItemDidCollapse() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewItemDidExpand implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewItemDidExpand(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineViewItemDidExpand:"), notification)
}

// HasOutlineViewItemDidExpand returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewItemDidExpand() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewItemWillCollapse implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewItemWillCollapse(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineViewItemWillCollapse:"), notification)
}

// HasOutlineViewItemWillCollapse returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewItemWillCollapse() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewItemWillExpand implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewItemWillExpand(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineViewItemWillExpand:"), notification)
}

// HasOutlineViewItemWillExpand returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewItemWillExpand() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewSelectionDidChange implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewSelectionDidChange(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineViewSelectionDidChange:"), notification)
}

// HasOutlineViewSelectionDidChange returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewSelectionDidChange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// OutlineViewSelectionIsChanging implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) OutlineViewSelectionIsChanging(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("outlineViewSelectionIsChanging:"), notification)
}

// HasOutlineViewSelectionIsChanging returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasOutlineViewSelectionIsChanging() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// SelectionShouldChangeInOutlineView implements the POutlineViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *OutlineViewDelegateObject) SelectionShouldChangeInOutlineView(outlineView IOutlineView) bool {
	return objc.Send[bool](o.ID, objc.Sel("selectionShouldChangeInOutlineView:"), outlineView)
}

// HasSelectionShouldChangeInOutlineView returns true; this is a placeholder for optional method checks.
func (o *OutlineViewDelegateObject) HasSelectionShouldChangeInOutlineView() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
