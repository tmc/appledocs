// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTableViewDataSource is the NSTableViewDataSource protocol interface.
//
// A set of methods that a table view uses to provide data to a table view and to allow the editing of the table view’s data source object.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTableViewDataSource
type PTableViewDataSource interface {
	// Optional methods
	NumberOfRowsInTableView(tableView ITableView) int
	HasNumberOfRowsInTableView() bool
	TableViewAcceptDropRowDropOperation(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) bool
	HasTableViewAcceptDropRowDropOperation() bool
	TableViewDraggingSessionEndedAtPointOperation(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, operation DragOperation)
	HasTableViewDraggingSessionEndedAtPointOperation() bool
	TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, rowIndexes foundation.IndexSet)
	HasTableViewDraggingSessionWillBeginAtPointForRowIndexes() bool
	TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes(tableView ITableView, dropDestination foundation.foundation.INSURL, indexSet foundation.IndexSet) []string
	HasTableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes() bool
	TableViewPasteboardWriterForRow(tableView ITableView, row int) unsafe.Pointer
	HasTableViewPasteboardWriterForRow() bool
	TableViewSetObjectValueForTableColumnRow(tableView ITableView, object objectivec.IObject, tableColumn ITableColumn, row int)
	HasTableViewSetObjectValueForTableColumnRow() bool
	TableViewSortDescriptorsDidChange(tableView ITableView, oldDescriptors []foundation.SortDescriptor)
	HasTableViewSortDescriptorsDidChange() bool
	TableViewUpdateDraggingItemsForDrag(tableView ITableView, draggingInfo unsafe.Pointer)
	HasTableViewUpdateDraggingItemsForDrag() bool
	TableViewValidateDropProposedRowProposedDropOperation(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) DragOperation
	HasTableViewValidateDropProposedRowProposedDropOperation() bool
	TableViewWriteRowsWithIndexesToPasteboard(tableView ITableView, rowIndexes foundation.IndexSet, pboard IPasteboard) bool
	HasTableViewWriteRowsWithIndexesToPasteboard() bool
	TableViewObjectValueForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) objc.ID
	HasTableViewObjectValueForTableColumnRow() bool
}

// TableViewDataSource is a delegate implementation builder for the PTableViewDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TableViewDataSource struct {
	_NumberOfRowsInTableView func(tableView ITableView) int
	_TableViewAcceptDropRowDropOperation func(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) bool
	_TableViewDraggingSessionEndedAtPointOperation func(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, operation DragOperation)
	_TableViewDraggingSessionWillBeginAtPointForRowIndexes func(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, rowIndexes foundation.IndexSet)
	_TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes func(tableView ITableView, dropDestination foundation.foundation.INSURL, indexSet foundation.IndexSet) []string
	_TableViewPasteboardWriterForRow func(tableView ITableView, row int) unsafe.Pointer
	_TableViewSetObjectValueForTableColumnRow func(tableView ITableView, object objectivec.IObject, tableColumn ITableColumn, row int)
	_TableViewSortDescriptorsDidChange func(tableView ITableView, oldDescriptors []foundation.SortDescriptor)
	_TableViewUpdateDraggingItemsForDrag func(tableView ITableView, draggingInfo unsafe.Pointer)
	_TableViewValidateDropProposedRowProposedDropOperation func(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) DragOperation
	_TableViewWriteRowsWithIndexesToPasteboard func(tableView ITableView, rowIndexes foundation.IndexSet, pboard IPasteboard) bool
	_TableViewObjectValueForTableColumnRow func(tableView ITableView, tableColumn ITableColumn, row int) objc.ID
}

// SetNumberOfRowsInTableView sets the handler for the NumberOfRowsInTableView delegate method.
//
// Returns the number of records managed for   by the data source object.
func (d *TableViewDataSource) SetNumberOfRowsInTableView(f func(tableView ITableView) int) {
	d._NumberOfRowsInTableView = f
}

// SetTableViewAcceptDropRowDropOperation sets the handler for the TableViewAcceptDropRowDropOperation delegate method.
//
// Called by   when the mouse button is released over a table view that previously decided to allow a drop.
func (d *TableViewDataSource) SetTableViewAcceptDropRowDropOperation(f func(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) bool) {
	d._TableViewAcceptDropRowDropOperation = f
}

// SetTableViewDraggingSessionEndedAtPointOperation sets the handler for the TableViewDraggingSessionEndedAtPointOperation delegate method.
//
// Implement this method to determine when a dragging session has ended.
func (d *TableViewDataSource) SetTableViewDraggingSessionEndedAtPointOperation(f func(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, operation DragOperation)) {
	d._TableViewDraggingSessionEndedAtPointOperation = f
}

// SetTableViewDraggingSessionWillBeginAtPointForRowIndexes sets the handler for the TableViewDraggingSessionWillBeginAtPointForRowIndexes delegate method.
//
// Implement this method to determine when a dragging session will begin.
func (d *TableViewDataSource) SetTableViewDraggingSessionWillBeginAtPointForRowIndexes(f func(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, rowIndexes foundation.IndexSet)) {
	d._TableViewDraggingSessionWillBeginAtPointForRowIndexes = f
}

// SetTableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes sets the handler for the TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes delegate method.
//
// Returns an array of filenames that represent the   rows for a drag to  .
func (d *TableViewDataSource) SetTableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes(f func(tableView ITableView, dropDestination foundation.foundation.INSURL, indexSet foundation.IndexSet) []string) {
	d._TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes = f
}

// SetTableViewPasteboardWriterForRow sets the handler for the TableViewPasteboardWriterForRow delegate method.
//
// Called to allow the table to support multiple item dragging.
func (d *TableViewDataSource) SetTableViewPasteboardWriterForRow(f func(tableView ITableView, row int) unsafe.Pointer) {
	d._TableViewPasteboardWriterForRow = f
}

// SetTableViewSetObjectValueForTableColumnRow sets the handler for the TableViewSetObjectValueForTableColumnRow delegate method.
//
// Sets the data object for an item in the specified row and column.
func (d *TableViewDataSource) SetTableViewSetObjectValueForTableColumnRow(f func(tableView ITableView, object objectivec.IObject, tableColumn ITableColumn, row int)) {
	d._TableViewSetObjectValueForTableColumnRow = f
}

// SetTableViewSortDescriptorsDidChange sets the handler for the TableViewSortDescriptorsDidChange delegate method.
//
// Called by   to indicate that sorting may need to be done.
func (d *TableViewDataSource) SetTableViewSortDescriptorsDidChange(f func(tableView ITableView, oldDescriptors []foundation.SortDescriptor)) {
	d._TableViewSortDescriptorsDidChange = f
}

// SetTableViewUpdateDraggingItemsForDrag sets the handler for the TableViewUpdateDraggingItemsForDrag delegate method.
//
// Implement this method to allow the table to update dragging items as they are dragged over a view.
func (d *TableViewDataSource) SetTableViewUpdateDraggingItemsForDrag(f func(tableView ITableView, draggingInfo unsafe.Pointer)) {
	d._TableViewUpdateDraggingItemsForDrag = f
}

// SetTableViewValidateDropProposedRowProposedDropOperation sets the handler for the TableViewValidateDropProposedRowProposedDropOperation delegate method.
//
// Used by   to determine a valid drop target.
func (d *TableViewDataSource) SetTableViewValidateDropProposedRowProposedDropOperation(f func(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) DragOperation) {
	d._TableViewValidateDropProposedRowProposedDropOperation = f
}

// SetTableViewWriteRowsWithIndexesToPasteboard sets the handler for the TableViewWriteRowsWithIndexesToPasteboard delegate method.
//
// Returns a Boolean value that indicates whether a drag operation is allowed.
func (d *TableViewDataSource) SetTableViewWriteRowsWithIndexesToPasteboard(f func(tableView ITableView, rowIndexes foundation.IndexSet, pboard IPasteboard) bool) {
	d._TableViewWriteRowsWithIndexesToPasteboard = f
}

// SetTableViewObjectValueForTableColumnRow sets the handler for the TableViewObjectValueForTableColumnRow delegate method.
//
// Called by the table view to return the data object associated with the specified row and column.
func (d *TableViewDataSource) SetTableViewObjectValueForTableColumnRow(f func(tableView ITableView, tableColumn ITableColumn, row int) objc.ID) {
	d._TableViewObjectValueForTableColumnRow = f
}

// NumberOfRowsInTableView implements the PTableViewDataSource interface.
func (d *TableViewDataSource) NumberOfRowsInTableView(tableView ITableView) int {
	if d._NumberOfRowsInTableView != nil {
		return d._NumberOfRowsInTableView(tableView)
	}
	var zero int
	return zero
}

// HasNumberOfRowsInTableView returns true if a handler for NumberOfRowsInTableView has been set.
func (d *TableViewDataSource) HasNumberOfRowsInTableView() bool {
	return d._NumberOfRowsInTableView != nil
}

// TableViewAcceptDropRowDropOperation implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewAcceptDropRowDropOperation(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) bool {
	if d._TableViewAcceptDropRowDropOperation != nil {
		return d._TableViewAcceptDropRowDropOperation(tableView, info, row, dropOperation)
	}
	var zero bool
	return zero
}

// HasTableViewAcceptDropRowDropOperation returns true if a handler for TableViewAcceptDropRowDropOperation has been set.
func (d *TableViewDataSource) HasTableViewAcceptDropRowDropOperation() bool {
	return d._TableViewAcceptDropRowDropOperation != nil
}

// TableViewDraggingSessionEndedAtPointOperation implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewDraggingSessionEndedAtPointOperation(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, operation DragOperation) {
	if d._TableViewDraggingSessionEndedAtPointOperation != nil {
		d._TableViewDraggingSessionEndedAtPointOperation(tableView, session, screenPoint, operation)
	}
}

// HasTableViewDraggingSessionEndedAtPointOperation returns true if a handler for TableViewDraggingSessionEndedAtPointOperation has been set.
func (d *TableViewDataSource) HasTableViewDraggingSessionEndedAtPointOperation() bool {
	return d._TableViewDraggingSessionEndedAtPointOperation != nil
}

// TableViewDraggingSessionWillBeginAtPointForRowIndexes implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, rowIndexes foundation.IndexSet) {
	if d._TableViewDraggingSessionWillBeginAtPointForRowIndexes != nil {
		d._TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView, session, screenPoint, rowIndexes)
	}
}

// HasTableViewDraggingSessionWillBeginAtPointForRowIndexes returns true if a handler for TableViewDraggingSessionWillBeginAtPointForRowIndexes has been set.
func (d *TableViewDataSource) HasTableViewDraggingSessionWillBeginAtPointForRowIndexes() bool {
	return d._TableViewDraggingSessionWillBeginAtPointForRowIndexes != nil
}

// TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes(tableView ITableView, dropDestination foundation.foundation.INSURL, indexSet foundation.IndexSet) []string {
	if d._TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes != nil {
		return d._TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes(tableView, dropDestination, indexSet)
	}
	var zero []string
	return zero
}

// HasTableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes returns true if a handler for TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes has been set.
func (d *TableViewDataSource) HasTableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes() bool {
	return d._TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes != nil
}

// TableViewPasteboardWriterForRow implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewPasteboardWriterForRow(tableView ITableView, row int) unsafe.Pointer {
	if d._TableViewPasteboardWriterForRow != nil {
		return d._TableViewPasteboardWriterForRow(tableView, row)
	}
	var zero unsafe.Pointer
	return zero
}

// HasTableViewPasteboardWriterForRow returns true if a handler for TableViewPasteboardWriterForRow has been set.
func (d *TableViewDataSource) HasTableViewPasteboardWriterForRow() bool {
	return d._TableViewPasteboardWriterForRow != nil
}

// TableViewSetObjectValueForTableColumnRow implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewSetObjectValueForTableColumnRow(tableView ITableView, object objectivec.IObject, tableColumn ITableColumn, row int) {
	if d._TableViewSetObjectValueForTableColumnRow != nil {
		d._TableViewSetObjectValueForTableColumnRow(tableView, object, tableColumn, row)
	}
}

// HasTableViewSetObjectValueForTableColumnRow returns true if a handler for TableViewSetObjectValueForTableColumnRow has been set.
func (d *TableViewDataSource) HasTableViewSetObjectValueForTableColumnRow() bool {
	return d._TableViewSetObjectValueForTableColumnRow != nil
}

// TableViewSortDescriptorsDidChange implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewSortDescriptorsDidChange(tableView ITableView, oldDescriptors []foundation.SortDescriptor) {
	if d._TableViewSortDescriptorsDidChange != nil {
		d._TableViewSortDescriptorsDidChange(tableView, oldDescriptors)
	}
}

// HasTableViewSortDescriptorsDidChange returns true if a handler for TableViewSortDescriptorsDidChange has been set.
func (d *TableViewDataSource) HasTableViewSortDescriptorsDidChange() bool {
	return d._TableViewSortDescriptorsDidChange != nil
}

// TableViewUpdateDraggingItemsForDrag implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewUpdateDraggingItemsForDrag(tableView ITableView, draggingInfo unsafe.Pointer) {
	if d._TableViewUpdateDraggingItemsForDrag != nil {
		d._TableViewUpdateDraggingItemsForDrag(tableView, draggingInfo)
	}
}

// HasTableViewUpdateDraggingItemsForDrag returns true if a handler for TableViewUpdateDraggingItemsForDrag has been set.
func (d *TableViewDataSource) HasTableViewUpdateDraggingItemsForDrag() bool {
	return d._TableViewUpdateDraggingItemsForDrag != nil
}

// TableViewValidateDropProposedRowProposedDropOperation implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewValidateDropProposedRowProposedDropOperation(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) DragOperation {
	if d._TableViewValidateDropProposedRowProposedDropOperation != nil {
		return d._TableViewValidateDropProposedRowProposedDropOperation(tableView, info, row, dropOperation)
	}
	var zero DragOperation
	return zero
}

// HasTableViewValidateDropProposedRowProposedDropOperation returns true if a handler for TableViewValidateDropProposedRowProposedDropOperation has been set.
func (d *TableViewDataSource) HasTableViewValidateDropProposedRowProposedDropOperation() bool {
	return d._TableViewValidateDropProposedRowProposedDropOperation != nil
}

// TableViewWriteRowsWithIndexesToPasteboard implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewWriteRowsWithIndexesToPasteboard(tableView ITableView, rowIndexes foundation.IndexSet, pboard IPasteboard) bool {
	if d._TableViewWriteRowsWithIndexesToPasteboard != nil {
		return d._TableViewWriteRowsWithIndexesToPasteboard(tableView, rowIndexes, pboard)
	}
	var zero bool
	return zero
}

// HasTableViewWriteRowsWithIndexesToPasteboard returns true if a handler for TableViewWriteRowsWithIndexesToPasteboard has been set.
func (d *TableViewDataSource) HasTableViewWriteRowsWithIndexesToPasteboard() bool {
	return d._TableViewWriteRowsWithIndexesToPasteboard != nil
}

// TableViewObjectValueForTableColumnRow implements the PTableViewDataSource interface.
func (d *TableViewDataSource) TableViewObjectValueForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) objc.ID {
	if d._TableViewObjectValueForTableColumnRow != nil {
		return d._TableViewObjectValueForTableColumnRow(tableView, tableColumn, row)
	}
	var zero objc.ID
	return zero
}

// HasTableViewObjectValueForTableColumnRow returns true if a handler for TableViewObjectValueForTableColumnRow has been set.
func (d *TableViewDataSource) HasTableViewObjectValueForTableColumnRow() bool {
	return d._TableViewObjectValueForTableColumnRow != nil
}

// TableViewDataSourceObject wraps an existing Objective-C object that conforms to the PTableViewDataSource protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TableViewDataSourceObject struct {
	objectivec.Object
}

// NewTableViewDataSourceObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTableViewDataSource protocol.
func NewTableViewDataSourceObject(obj objectivec.Object) *TableViewDataSourceObject {
	return &TableViewDataSourceObject{obj}
}

// Make sure TableViewDataSourceObject implements PTableViewDataSource.
var _ PTableViewDataSource = (*TableViewDataSourceObject)(nil)

// NumberOfRowsInTableView implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) NumberOfRowsInTableView(tableView ITableView) int {
	return objc.Send[int](o.ID, objc.Sel("numberOfRowsInTableView:"), tableView)
}

// HasNumberOfRowsInTableView returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasNumberOfRowsInTableView() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewAcceptDropRowDropOperation implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewAcceptDropRowDropOperation(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) bool {
	return objc.Send[bool](o.ID, objc.Sel("tableView:acceptDrop:row:dropOperation:"), tableView, info, row, dropOperation)
}

// HasTableViewAcceptDropRowDropOperation returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewAcceptDropRowDropOperation() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewDraggingSessionEndedAtPointOperation implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewDraggingSessionEndedAtPointOperation(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, operation DragOperation) {
	objc.Send[objc.ID](o.ID, objc.Sel("tableView:draggingSession:endedAtPoint:operation:"), tableView, session, screenPoint, operation)
}

// HasTableViewDraggingSessionEndedAtPointOperation returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewDraggingSessionEndedAtPointOperation() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewDraggingSessionWillBeginAtPointForRowIndexes implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView ITableView, session IDraggingSession, screenPoint corefoundation.CGPoint, rowIndexes foundation.IndexSet) {
	objc.Send[objc.ID](o.ID, objc.Sel("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"), tableView, session, screenPoint, rowIndexes)
}

// HasTableViewDraggingSessionWillBeginAtPointForRowIndexes returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewDraggingSessionWillBeginAtPointForRowIndexes() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes(tableView ITableView, dropDestination foundation.foundation.INSURL, indexSet foundation.IndexSet) []string {
	return objc.Send[[]string](o.ID, objc.Sel("tableView:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:"), tableView, dropDestination, indexSet)
}

// HasTableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexes() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewPasteboardWriterForRow implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewPasteboardWriterForRow(tableView ITableView, row int) unsafe.Pointer {
	return objc.Send[unsafe.Pointer](o.ID, objc.Sel("tableView:pasteboardWriterForRow:"), tableView, row)
}

// HasTableViewPasteboardWriterForRow returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewPasteboardWriterForRow() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewSetObjectValueForTableColumnRow implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewSetObjectValueForTableColumnRow(tableView ITableView, object objectivec.IObject, tableColumn ITableColumn, row int) {
	objc.Send[objc.ID](o.ID, objc.Sel("tableView:setObjectValue:forTableColumn:row:"), tableView, object, tableColumn, row)
}

// HasTableViewSetObjectValueForTableColumnRow returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewSetObjectValueForTableColumnRow() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewSortDescriptorsDidChange implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewSortDescriptorsDidChange(tableView ITableView, oldDescriptors []foundation.SortDescriptor) {
	objc.Send[objc.ID](o.ID, objc.Sel("tableView:sortDescriptorsDidChange:"), tableView, oldDescriptors)
}

// HasTableViewSortDescriptorsDidChange returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewSortDescriptorsDidChange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewUpdateDraggingItemsForDrag implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewUpdateDraggingItemsForDrag(tableView ITableView, draggingInfo unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("tableView:updateDraggingItemsForDrag:"), tableView, draggingInfo)
}

// HasTableViewUpdateDraggingItemsForDrag returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewUpdateDraggingItemsForDrag() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewValidateDropProposedRowProposedDropOperation implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewValidateDropProposedRowProposedDropOperation(tableView ITableView, info unsafe.Pointer, row int, dropOperation TableViewDropOperation) DragOperation {
	return objc.Send[DragOperation](o.ID, objc.Sel("tableView:validateDrop:proposedRow:proposedDropOperation:"), tableView, info, row, dropOperation)
}

// HasTableViewValidateDropProposedRowProposedDropOperation returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewValidateDropProposedRowProposedDropOperation() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewWriteRowsWithIndexesToPasteboard implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewWriteRowsWithIndexesToPasteboard(tableView ITableView, rowIndexes foundation.IndexSet, pboard IPasteboard) bool {
	return objc.Send[bool](o.ID, objc.Sel("tableView:writeRowsWithIndexes:toPasteboard:"), tableView, rowIndexes, pboard)
}

// HasTableViewWriteRowsWithIndexesToPasteboard returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewWriteRowsWithIndexesToPasteboard() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TableViewObjectValueForTableColumnRow implements the PTableViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TableViewDataSourceObject) TableViewObjectValueForTableColumnRow(tableView ITableView, tableColumn ITableColumn, row int) objc.ID {
	return objc.Send[objc.ID](o.ID, objc.Sel("tableView:objectValueForTableColumn:row:"), tableView, tableColumn, row)
}

// HasTableViewObjectValueForTableColumnRow returns true; this is a placeholder for optional method checks.
func (o *TableViewDataSourceObject) HasTableViewObjectValueForTableColumnRow() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
