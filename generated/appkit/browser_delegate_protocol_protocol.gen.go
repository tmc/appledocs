// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PBrowserDelegate is the NSBrowserDelegate protocol interface.
//
// A set of methods that a browser delegate implements to manage selection, scrolling, sizing, and other behavior.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSBrowserDelegate
type PBrowserDelegate interface {
	// Optional methods
	BrowserAcceptDropAtRowColumnDropOperation(browser IBrowser, info unsafe.Pointer, row int, column int, dropOperation BrowserDropOperation) bool
	HasBrowserAcceptDropAtRowColumnDropOperation() bool
	BrowserCanDragRowsWithIndexesInColumnWithEvent(browser IBrowser, rowIndexes foundation.IndexSet, column int, event IEvent) bool
	HasBrowserCanDragRowsWithIndexesInColumnWithEvent() bool
	BrowserChildOfItem(browser IBrowser, index int, item objc.IObject) objc.ID
	HasBrowserChildOfItem() bool
	BrowserCreateRowsForColumnInMatrix(sender IBrowser, column int, matrix IMatrix)
	HasBrowserCreateRowsForColumnInMatrix() bool
	BrowserDidChangeLastColumnToColumn(browser IBrowser, oldLastColumn int, column int)
	HasBrowserDidChangeLastColumnToColumn() bool
	BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(browser IBrowser, rowIndexes foundation.IndexSet, column int, event IEvent, dragImageOffset PointPointer /* not a class type */) Image
	HasBrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset() bool
	BrowserHeaderViewControllerForItem(browser IBrowser, item objc.IObject) ViewController
	HasBrowserHeaderViewControllerForItem() bool
	BrowserHeightOfRowInColumn(browser IBrowser, row int, columnIndex int) float64
	HasBrowserHeightOfRowInColumn() bool
	BrowserIsColumnValid(sender IBrowser, column int) bool
	HasBrowserIsColumnValid() bool
	BrowserIsLeafItem(browser IBrowser, item objc.IObject) bool
	HasBrowserIsLeafItem() bool
	BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn(browser IBrowser, dropDestination objc.IObject /* cross-framework: NSURL */, rowIndexes foundation.IndexSet, column int) []string
	HasBrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn() bool
	BrowserNextTypeSelectMatchFromRowToRowInColumnForString(browser IBrowser, startRow int, endRow int, column int, searchString objc.IObject /* cross-framework: NSString */) int
	HasBrowserNextTypeSelectMatchFromRowToRowInColumnForString() bool
	BrowserNumberOfChildrenOfItem(browser IBrowser, item objc.IObject) int
	HasBrowserNumberOfChildrenOfItem() bool
	BrowserNumberOfRowsInColumn(sender IBrowser, column int) int
	HasBrowserNumberOfRowsInColumn() bool
	BrowserObjectValueForItem(browser IBrowser, item objc.IObject) objc.ID
	HasBrowserObjectValueForItem() bool
	BrowserPreviewViewControllerForLeafItem(browser IBrowser, item objc.IObject) ViewController
	HasBrowserPreviewViewControllerForLeafItem() bool
	BrowserSelectCellWithStringInColumn(sender IBrowser, title objc.IObject /* cross-framework: NSString */, column int) bool
	HasBrowserSelectCellWithStringInColumn() bool
	BrowserSelectRowInColumn(sender IBrowser, row int, column int) bool
	HasBrowserSelectRowInColumn() bool
	BrowserSelectionIndexesForProposedSelectionInColumn(browser IBrowser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IndexSet
	HasBrowserSelectionIndexesForProposedSelectionInColumn() bool
	BrowserSetObjectValueForItem(browser IBrowser, object objc.IObject, item objc.IObject)
	HasBrowserSetObjectValueForItem() bool
	BrowserShouldEditItem(browser IBrowser, item objc.IObject) bool
	HasBrowserShouldEditItem() bool
	BrowserShouldShowCellExpansionForRowColumn(browser IBrowser, row int, column int) bool
	HasBrowserShouldShowCellExpansionForRowColumn() bool
	BrowserShouldSizeColumnForUserResizeToWidth(browser IBrowser, columnIndex int, forUserResize bool, suggestedWidth float64) float64
	HasBrowserShouldSizeColumnForUserResizeToWidth() bool
	BrowserShouldTypeSelectForEventWithCurrentSearchString(browser IBrowser, event IEvent, searchString objc.IObject /* cross-framework: NSString */) bool
	HasBrowserShouldTypeSelectForEventWithCurrentSearchString() bool
	BrowserSizeToFitWidthOfColumn(browser IBrowser, columnIndex int) float64
	HasBrowserSizeToFitWidthOfColumn() bool
	BrowserTitleOfColumn(sender IBrowser, column int) foundation.String
	HasBrowserTitleOfColumn() bool
	BrowserTypeSelectStringForRowInColumn(browser IBrowser, row int, column int) foundation.String
	HasBrowserTypeSelectStringForRowInColumn() bool
	BrowserValidateDropProposedRowColumnDropOperation(browser IBrowser, info unsafe.Pointer, row int, column int, dropOperation BrowserDropOperation) DragOperation
	HasBrowserValidateDropProposedRowColumnDropOperation() bool
	BrowserWillDisplayCellAtRowColumn(sender IBrowser, cell objc.IObject, row int, column int)
	HasBrowserWillDisplayCellAtRowColumn() bool
	BrowserWriteRowsWithIndexesInColumnToPasteboard(browser IBrowser, rowIndexes foundation.IndexSet, column int, pasteboard IPasteboard) bool
	HasBrowserWriteRowsWithIndexesInColumnToPasteboard() bool
	BrowserColumnConfigurationDidChange(notification foundation.Notification)
	HasBrowserColumnConfigurationDidChange() bool
	BrowserDidScroll(sender IBrowser)
	HasBrowserDidScroll() bool
	BrowserWillScroll(sender IBrowser)
	HasBrowserWillScroll() bool
	RootItemForBrowser(browser IBrowser) objc.ID
	HasRootItemForBrowser() bool
}

// BrowserDelegate is a delegate implementation builder for the PBrowserDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type BrowserDelegate struct {
	_BrowserAcceptDropAtRowColumnDropOperation func(browser IBrowser, info unsafe.Pointer, row int, column int, dropOperation BrowserDropOperation) bool
	_BrowserCanDragRowsWithIndexesInColumnWithEvent func(browser IBrowser, rowIndexes foundation.IndexSet, column int, event IEvent) bool
	_BrowserChildOfItem func(browser IBrowser, index int, item objc.IObject) objc.ID
	_BrowserCreateRowsForColumnInMatrix func(sender IBrowser, column int, matrix IMatrix)
	_BrowserDidChangeLastColumnToColumn func(browser IBrowser, oldLastColumn int, column int)
	_BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset func(browser IBrowser, rowIndexes foundation.IndexSet, column int, event IEvent, dragImageOffset PointPointer /* not a class type */) Image
	_BrowserHeaderViewControllerForItem func(browser IBrowser, item objc.IObject) ViewController
	_BrowserHeightOfRowInColumn func(browser IBrowser, row int, columnIndex int) float64
	_BrowserIsColumnValid func(sender IBrowser, column int) bool
	_BrowserIsLeafItem func(browser IBrowser, item objc.IObject) bool
	_BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn func(browser IBrowser, dropDestination objc.IObject /* cross-framework: NSURL */, rowIndexes foundation.IndexSet, column int) []string
	_BrowserNextTypeSelectMatchFromRowToRowInColumnForString func(browser IBrowser, startRow int, endRow int, column int, searchString objc.IObject /* cross-framework: NSString */) int
	_BrowserNumberOfChildrenOfItem func(browser IBrowser, item objc.IObject) int
	_BrowserNumberOfRowsInColumn func(sender IBrowser, column int) int
	_BrowserObjectValueForItem func(browser IBrowser, item objc.IObject) objc.ID
	_BrowserPreviewViewControllerForLeafItem func(browser IBrowser, item objc.IObject) ViewController
	_BrowserSelectCellWithStringInColumn func(sender IBrowser, title objc.IObject /* cross-framework: NSString */, column int) bool
	_BrowserSelectRowInColumn func(sender IBrowser, row int, column int) bool
	_BrowserSelectionIndexesForProposedSelectionInColumn func(browser IBrowser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IndexSet
	_BrowserSetObjectValueForItem func(browser IBrowser, object objc.IObject, item objc.IObject)
	_BrowserShouldEditItem func(browser IBrowser, item objc.IObject) bool
	_BrowserShouldShowCellExpansionForRowColumn func(browser IBrowser, row int, column int) bool
	_BrowserShouldSizeColumnForUserResizeToWidth func(browser IBrowser, columnIndex int, forUserResize bool, suggestedWidth float64) float64
	_BrowserShouldTypeSelectForEventWithCurrentSearchString func(browser IBrowser, event IEvent, searchString objc.IObject /* cross-framework: NSString */) bool
	_BrowserSizeToFitWidthOfColumn func(browser IBrowser, columnIndex int) float64
	_BrowserTitleOfColumn func(sender IBrowser, column int) foundation.String
	_BrowserTypeSelectStringForRowInColumn func(browser IBrowser, row int, column int) foundation.String
	_BrowserValidateDropProposedRowColumnDropOperation func(browser IBrowser, info unsafe.Pointer, row int, column int, dropOperation BrowserDropOperation) DragOperation
	_BrowserWillDisplayCellAtRowColumn func(sender IBrowser, cell objc.IObject, row int, column int)
	_BrowserWriteRowsWithIndexesInColumnToPasteboard func(browser IBrowser, rowIndexes foundation.IndexSet, column int, pasteboard IPasteboard) bool
	_BrowserColumnConfigurationDidChange func(notification foundation.Notification)
	_BrowserDidScroll func(sender IBrowser)
	_BrowserWillScroll func(sender IBrowser)
	_RootItemForBrowser func(browser IBrowser) objc.ID
}

// SetBrowserAcceptDropAtRowColumnDropOperation sets the handler for the BrowserAcceptDropAtRowColumnDropOperation delegate method.
//
// Sent to the delegate during a dragging session to determine whether to accept the drop.
func (d *BrowserDelegate) SetBrowserAcceptDropAtRowColumnDropOperation(f func(browser IBrowser, info unsafe.Pointer, row int, column int, dropOperation BrowserDropOperation) bool) {
	d._BrowserAcceptDropAtRowColumnDropOperation = f
}

// SetBrowserCanDragRowsWithIndexesInColumnWithEvent sets the handler for the BrowserCanDragRowsWithIndexesInColumnWithEvent delegate method.
//
// Sent to the delegate to determine whether the browser can attempt to initiate a drag of the specified rows for the specified event.
func (d *BrowserDelegate) SetBrowserCanDragRowsWithIndexesInColumnWithEvent(f func(browser IBrowser, rowIndexes foundation.IndexSet, column int, event IEvent) bool) {
	d._BrowserCanDragRowsWithIndexesInColumnWithEvent = f
}

// SetBrowserChildOfItem sets the handler for the BrowserChildOfItem delegate method.
//
// Asks the delegate to return the child of the specified item at the specified index.
func (d *BrowserDelegate) SetBrowserChildOfItem(f func(browser IBrowser, index int, item objc.IObject) objc.ID) {
	d._BrowserChildOfItem = f
}

// SetBrowserCreateRowsForColumnInMatrix sets the handler for the BrowserCreateRowsForColumnInMatrix delegate method.
//
// Creates a row in the given matrix for each row of data in the specified column of the browser.
func (d *BrowserDelegate) SetBrowserCreateRowsForColumnInMatrix(f func(sender IBrowser, column int, matrix IMatrix)) {
	d._BrowserCreateRowsForColumnInMatrix = f
}

// SetBrowserDidChangeLastColumnToColumn sets the handler for the BrowserDidChangeLastColumnToColumn delegate method.
//
// Tells the delegate that the browser’s last column changed.
func (d *BrowserDelegate) SetBrowserDidChangeLastColumnToColumn(f func(browser IBrowser, oldLastColumn int, column int)) {
	d._BrowserDidChangeLastColumnToColumn = f
}

// SetBrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset sets the handler for the BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset delegate method.
//
// Sent to the delegate to obtain an image to represent dragged rows during a drag operation on a browser.
func (d *BrowserDelegate) SetBrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(f func(browser IBrowser, rowIndexes foundation.IndexSet, column int, event IEvent, dragImageOffset PointPointer /* not a class type */) Image) {
	d._BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset = f
}

// SetBrowserHeaderViewControllerForItem sets the handler for the BrowserHeaderViewControllerForItem delegate method.
//
// Asks the delegate for a controller that provides a header view for the specified column item.
func (d *BrowserDelegate) SetBrowserHeaderViewControllerForItem(f func(browser IBrowser, item objc.IObject) ViewController) {
	d._BrowserHeaderViewControllerForItem = f
}

// SetBrowserHeightOfRowInColumn sets the handler for the BrowserHeightOfRowInColumn delegate method.
//
// Specifies the height of the specified row in the specified column.
func (d *BrowserDelegate) SetBrowserHeightOfRowInColumn(f func(browser IBrowser, row int, columnIndex int) float64) {
	d._BrowserHeightOfRowInColumn = f
}

// SetBrowserIsColumnValid sets the handler for the BrowserIsColumnValid delegate method.
//
// Returns whether the contents of the specified column are valid.
func (d *BrowserDelegate) SetBrowserIsColumnValid(f func(sender IBrowser, column int) bool) {
	d._BrowserIsColumnValid = f
}

// SetBrowserIsLeafItem sets the handler for the BrowserIsLeafItem delegate method.
//
// Asks the delegate whether the specified item is a leaf item (an item that cannot be expanded).
func (d *BrowserDelegate) SetBrowserIsLeafItem(f func(browser IBrowser, item objc.IObject) bool) {
	d._BrowserIsLeafItem = f
}

// SetBrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn sets the handler for the BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn delegate method.
//
// Implements file promise drag operations.
func (d *BrowserDelegate) SetBrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn(f func(browser IBrowser, dropDestination objc.IObject /* cross-framework: NSURL */, rowIndexes foundation.IndexSet, column int) []string) {
	d._BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn = f
}

// SetBrowserNextTypeSelectMatchFromRowToRowInColumnForString sets the handler for the BrowserNextTypeSelectMatchFromRowToRowInColumnForString delegate method.
//
// Sent to the delegate to customize a browser’s keyboard-based selection (type select) behavior.
func (d *BrowserDelegate) SetBrowserNextTypeSelectMatchFromRowToRowInColumnForString(f func(browser IBrowser, startRow int, endRow int, column int, searchString objc.IObject /* cross-framework: NSString */) int) {
	d._BrowserNextTypeSelectMatchFromRowToRowInColumnForString = f
}

// SetBrowserNumberOfChildrenOfItem sets the handler for the BrowserNumberOfChildrenOfItem delegate method.
//
// Asks the delegate for the number of children the given item has.
func (d *BrowserDelegate) SetBrowserNumberOfChildrenOfItem(f func(browser IBrowser, item objc.IObject) int) {
	d._BrowserNumberOfChildrenOfItem = f
}

// SetBrowserNumberOfRowsInColumn sets the handler for the BrowserNumberOfRowsInColumn delegate method.
//
// Returns the number of rows of data in the specified column.
func (d *BrowserDelegate) SetBrowserNumberOfRowsInColumn(f func(sender IBrowser, column int) int) {
	d._BrowserNumberOfRowsInColumn = f
}

// SetBrowserObjectValueForItem sets the handler for the BrowserObjectValueForItem delegate method.
//
// Returns the object that the specified item uses to draw its contents.
func (d *BrowserDelegate) SetBrowserObjectValueForItem(f func(browser IBrowser, item objc.IObject) objc.ID) {
	d._BrowserObjectValueForItem = f
}

// SetBrowserPreviewViewControllerForLeafItem sets the handler for the BrowserPreviewViewControllerForLeafItem delegate method.
//
// Asks the delegate for a controller that provides a preview column for the specified leaf item.
func (d *BrowserDelegate) SetBrowserPreviewViewControllerForLeafItem(f func(browser IBrowser, item objc.IObject) ViewController) {
	d._BrowserPreviewViewControllerForLeafItem = f
}

// SetBrowserSelectCellWithStringInColumn sets the handler for the BrowserSelectCellWithStringInColumn delegate method.
//
// Asks the delegate to select the cell with the given title in the specified column.
func (d *BrowserDelegate) SetBrowserSelectCellWithStringInColumn(f func(sender IBrowser, title objc.IObject /* cross-framework: NSString */, column int) bool) {
	d._BrowserSelectCellWithStringInColumn = f
}

// SetBrowserSelectRowInColumn sets the handler for the BrowserSelectRowInColumn delegate method.
//
// Asks the delegate to select the cell at the specified row and column location.
func (d *BrowserDelegate) SetBrowserSelectRowInColumn(f func(sender IBrowser, row int, column int) bool) {
	d._BrowserSelectRowInColumn = f
}

// SetBrowserSelectionIndexesForProposedSelectionInColumn sets the handler for the BrowserSelectionIndexesForProposedSelectionInColumn delegate method.
//
// Asks the delegate for a set of indexes to select when the user changes the selection in the browser with the keyboard or mouse.
func (d *BrowserDelegate) SetBrowserSelectionIndexesForProposedSelectionInColumn(f func(browser IBrowser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IndexSet) {
	d._BrowserSelectionIndexesForProposedSelectionInColumn = f
}

// SetBrowserSetObjectValueForItem sets the handler for the BrowserSetObjectValueForItem delegate method.
//
// Sets the object that the specified item uses to draw its contents to the specified object.
func (d *BrowserDelegate) SetBrowserSetObjectValueForItem(f func(browser IBrowser, object objc.IObject, item objc.IObject)) {
	d._BrowserSetObjectValueForItem = f
}

// SetBrowserShouldEditItem sets the handler for the BrowserShouldEditItem delegate method.
//
// Asks the delegate whether the browser may start an editing session for the specified item.
func (d *BrowserDelegate) SetBrowserShouldEditItem(f func(browser IBrowser, item objc.IObject) bool) {
	d._BrowserShouldEditItem = f
}

// SetBrowserShouldShowCellExpansionForRowColumn sets the handler for the BrowserShouldShowCellExpansionForRowColumn delegate method.
//
// Invoked to allow the delegate to control cell expansion for a specific row and column.
func (d *BrowserDelegate) SetBrowserShouldShowCellExpansionForRowColumn(f func(browser IBrowser, row int, column int) bool) {
	d._BrowserShouldShowCellExpansionForRowColumn = f
}

// SetBrowserShouldSizeColumnForUserResizeToWidth sets the handler for the BrowserShouldSizeColumnForUserResizeToWidth delegate method.
//
// Used to determine a column’s initial size.
func (d *BrowserDelegate) SetBrowserShouldSizeColumnForUserResizeToWidth(f func(browser IBrowser, columnIndex int, forUserResize bool, suggestedWidth float64) float64) {
	d._BrowserShouldSizeColumnForUserResizeToWidth = f
}

// SetBrowserShouldTypeSelectForEventWithCurrentSearchString sets the handler for the BrowserShouldTypeSelectForEventWithCurrentSearchString delegate method.
//
// Sent to the delegate to determine whether keyboard-based selection (type select) for a given event and search string should proceed.
func (d *BrowserDelegate) SetBrowserShouldTypeSelectForEventWithCurrentSearchString(f func(browser IBrowser, event IEvent, searchString objc.IObject /* cross-framework: NSString */) bool) {
	d._BrowserShouldTypeSelectForEventWithCurrentSearchString = f
}

// SetBrowserSizeToFitWidthOfColumn sets the handler for the BrowserSizeToFitWidthOfColumn delegate method.
//
// Returns the ideal width for a column.
func (d *BrowserDelegate) SetBrowserSizeToFitWidthOfColumn(f func(browser IBrowser, columnIndex int) float64) {
	d._BrowserSizeToFitWidthOfColumn = f
}

// SetBrowserTitleOfColumn sets the handler for the BrowserTitleOfColumn delegate method.
//
// Asks the delegate for the title to display above the specified column.
func (d *BrowserDelegate) SetBrowserTitleOfColumn(f func(sender IBrowser, column int) foundation.String) {
	d._BrowserTitleOfColumn = f
}

// SetBrowserTypeSelectStringForRowInColumn sets the handler for the BrowserTypeSelectStringForRowInColumn delegate method.
//
// Sent to the delegate to get the keyboard-based selection (type select) string for the specified row and column.
func (d *BrowserDelegate) SetBrowserTypeSelectStringForRowInColumn(f func(browser IBrowser, row int, column int) foundation.String) {
	d._BrowserTypeSelectStringForRowInColumn = f
}

// SetBrowserValidateDropProposedRowColumnDropOperation sets the handler for the BrowserValidateDropProposedRowColumnDropOperation delegate method.
//
// Sent to the delegate during a dragging session to determine whether a drop should be accepted and to obtain the drop location. This method is required for a browser to be a drag destination.
func (d *BrowserDelegate) SetBrowserValidateDropProposedRowColumnDropOperation(f func(browser IBrowser, info unsafe.Pointer, row int, column int, dropOperation BrowserDropOperation) DragOperation) {
	d._BrowserValidateDropProposedRowColumnDropOperation = f
}

// SetBrowserWillDisplayCellAtRowColumn sets the handler for the BrowserWillDisplayCellAtRowColumn delegate method.
//
// Gives the delegate the opportunity to modify the specified cell at the given row and column location before the browser displays it.
func (d *BrowserDelegate) SetBrowserWillDisplayCellAtRowColumn(f func(sender IBrowser, cell objc.IObject, row int, column int)) {
	d._BrowserWillDisplayCellAtRowColumn = f
}

// SetBrowserWriteRowsWithIndexesInColumnToPasteboard sets the handler for the BrowserWriteRowsWithIndexesInColumnToPasteboard delegate method.
//
// Determines whether a drag operation can proceed. This method is required for a browser to be a drag source.
func (d *BrowserDelegate) SetBrowserWriteRowsWithIndexesInColumnToPasteboard(f func(browser IBrowser, rowIndexes foundation.IndexSet, column int, pasteboard IPasteboard) bool) {
	d._BrowserWriteRowsWithIndexesInColumnToPasteboard = f
}

// SetBrowserColumnConfigurationDidChange sets the handler for the BrowserColumnConfigurationDidChange delegate method.
//
// Used by clients to implement their own column width persistence.
func (d *BrowserDelegate) SetBrowserColumnConfigurationDidChange(f func(notification foundation.Notification)) {
	d._BrowserColumnConfigurationDidChange = f
}

// SetBrowserDidScroll sets the handler for the BrowserDidScroll delegate method.
//
// Notifies the delegate when the browser has scrolled.
func (d *BrowserDelegate) SetBrowserDidScroll(f func(sender IBrowser)) {
	d._BrowserDidScroll = f
}

// SetBrowserWillScroll sets the handler for the BrowserWillScroll delegate method.
//
// Notifies the delegate when the browser will scroll.
func (d *BrowserDelegate) SetBrowserWillScroll(f func(sender IBrowser)) {
	d._BrowserWillScroll = f
}

// SetRootItemForBrowser sets the handler for the RootItemForBrowser delegate method.
//
// Asks the delegate to return the root item of the browser.
func (d *BrowserDelegate) SetRootItemForBrowser(f func(browser IBrowser) objc.ID) {
	d._RootItemForBrowser = f
}

// BrowserAcceptDropAtRowColumnDropOperation implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserAcceptDropAtRowColumnDropOperation(browser IBrowser, info unsafe.Pointer, row int, column int, dropOperation BrowserDropOperation) bool {
	if d._BrowserAcceptDropAtRowColumnDropOperation != nil {
		return d._BrowserAcceptDropAtRowColumnDropOperation(browser, info, row, column, dropOperation)
	}
	var zero bool
	return zero
}

// HasBrowserAcceptDropAtRowColumnDropOperation returns true if a handler for BrowserAcceptDropAtRowColumnDropOperation has been set.
func (d *BrowserDelegate) HasBrowserAcceptDropAtRowColumnDropOperation() bool {
	return d._BrowserAcceptDropAtRowColumnDropOperation != nil
}

// BrowserCanDragRowsWithIndexesInColumnWithEvent implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserCanDragRowsWithIndexesInColumnWithEvent(browser IBrowser, rowIndexes foundation.IndexSet, column int, event IEvent) bool {
	if d._BrowserCanDragRowsWithIndexesInColumnWithEvent != nil {
		return d._BrowserCanDragRowsWithIndexesInColumnWithEvent(browser, rowIndexes, column, event)
	}
	var zero bool
	return zero
}

// HasBrowserCanDragRowsWithIndexesInColumnWithEvent returns true if a handler for BrowserCanDragRowsWithIndexesInColumnWithEvent has been set.
func (d *BrowserDelegate) HasBrowserCanDragRowsWithIndexesInColumnWithEvent() bool {
	return d._BrowserCanDragRowsWithIndexesInColumnWithEvent != nil
}

// BrowserChildOfItem implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserChildOfItem(browser IBrowser, index int, item objc.IObject) objc.ID {
	if d._BrowserChildOfItem != nil {
		return d._BrowserChildOfItem(browser, index, item)
	}
	var zero objc.ID
	return zero
}

// HasBrowserChildOfItem returns true if a handler for BrowserChildOfItem has been set.
func (d *BrowserDelegate) HasBrowserChildOfItem() bool {
	return d._BrowserChildOfItem != nil
}

// BrowserCreateRowsForColumnInMatrix implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserCreateRowsForColumnInMatrix(sender IBrowser, column int, matrix IMatrix) {
	if d._BrowserCreateRowsForColumnInMatrix != nil {
		d._BrowserCreateRowsForColumnInMatrix(sender, column, matrix)
	}
}

// HasBrowserCreateRowsForColumnInMatrix returns true if a handler for BrowserCreateRowsForColumnInMatrix has been set.
func (d *BrowserDelegate) HasBrowserCreateRowsForColumnInMatrix() bool {
	return d._BrowserCreateRowsForColumnInMatrix != nil
}

// BrowserDidChangeLastColumnToColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserDidChangeLastColumnToColumn(browser IBrowser, oldLastColumn int, column int) {
	if d._BrowserDidChangeLastColumnToColumn != nil {
		d._BrowserDidChangeLastColumnToColumn(browser, oldLastColumn, column)
	}
}

// HasBrowserDidChangeLastColumnToColumn returns true if a handler for BrowserDidChangeLastColumnToColumn has been set.
func (d *BrowserDelegate) HasBrowserDidChangeLastColumnToColumn() bool {
	return d._BrowserDidChangeLastColumnToColumn != nil
}

// BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(browser IBrowser, rowIndexes foundation.IndexSet, column int, event IEvent, dragImageOffset PointPointer /* not a class type */) Image {
	if d._BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset != nil {
		return d._BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset(browser, rowIndexes, column, event, dragImageOffset)
	}
	var zero Image
	return zero
}

// HasBrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset returns true if a handler for BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset has been set.
func (d *BrowserDelegate) HasBrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset() bool {
	return d._BrowserDraggingImageForRowsWithIndexesInColumnWithEventOffset != nil
}

// BrowserHeaderViewControllerForItem implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserHeaderViewControllerForItem(browser IBrowser, item objc.IObject) ViewController {
	if d._BrowserHeaderViewControllerForItem != nil {
		return d._BrowserHeaderViewControllerForItem(browser, item)
	}
	var zero ViewController
	return zero
}

// HasBrowserHeaderViewControllerForItem returns true if a handler for BrowserHeaderViewControllerForItem has been set.
func (d *BrowserDelegate) HasBrowserHeaderViewControllerForItem() bool {
	return d._BrowserHeaderViewControllerForItem != nil
}

// BrowserHeightOfRowInColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserHeightOfRowInColumn(browser IBrowser, row int, columnIndex int) float64 {
	if d._BrowserHeightOfRowInColumn != nil {
		return d._BrowserHeightOfRowInColumn(browser, row, columnIndex)
	}
	var zero float64
	return zero
}

// HasBrowserHeightOfRowInColumn returns true if a handler for BrowserHeightOfRowInColumn has been set.
func (d *BrowserDelegate) HasBrowserHeightOfRowInColumn() bool {
	return d._BrowserHeightOfRowInColumn != nil
}

// BrowserIsColumnValid implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserIsColumnValid(sender IBrowser, column int) bool {
	if d._BrowserIsColumnValid != nil {
		return d._BrowserIsColumnValid(sender, column)
	}
	var zero bool
	return zero
}

// HasBrowserIsColumnValid returns true if a handler for BrowserIsColumnValid has been set.
func (d *BrowserDelegate) HasBrowserIsColumnValid() bool {
	return d._BrowserIsColumnValid != nil
}

// BrowserIsLeafItem implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserIsLeafItem(browser IBrowser, item objc.IObject) bool {
	if d._BrowserIsLeafItem != nil {
		return d._BrowserIsLeafItem(browser, item)
	}
	var zero bool
	return zero
}

// HasBrowserIsLeafItem returns true if a handler for BrowserIsLeafItem has been set.
func (d *BrowserDelegate) HasBrowserIsLeafItem() bool {
	return d._BrowserIsLeafItem != nil
}

// BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn(browser IBrowser, dropDestination objc.IObject /* cross-framework: NSURL */, rowIndexes foundation.IndexSet, column int) []string {
	if d._BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn != nil {
		return d._BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn(browser, dropDestination, rowIndexes, column)
	}
	var zero []string
	return zero
}

// HasBrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn returns true if a handler for BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn has been set.
func (d *BrowserDelegate) HasBrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn() bool {
	return d._BrowserNamesOfPromisedFilesDroppedAtDestinationForDraggedRowsWithIndexesInColumn != nil
}

// BrowserNextTypeSelectMatchFromRowToRowInColumnForString implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserNextTypeSelectMatchFromRowToRowInColumnForString(browser IBrowser, startRow int, endRow int, column int, searchString objc.IObject /* cross-framework: NSString */) int {
	if d._BrowserNextTypeSelectMatchFromRowToRowInColumnForString != nil {
		return d._BrowserNextTypeSelectMatchFromRowToRowInColumnForString(browser, startRow, endRow, column, searchString)
	}
	var zero int
	return zero
}

// HasBrowserNextTypeSelectMatchFromRowToRowInColumnForString returns true if a handler for BrowserNextTypeSelectMatchFromRowToRowInColumnForString has been set.
func (d *BrowserDelegate) HasBrowserNextTypeSelectMatchFromRowToRowInColumnForString() bool {
	return d._BrowserNextTypeSelectMatchFromRowToRowInColumnForString != nil
}

// BrowserNumberOfChildrenOfItem implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserNumberOfChildrenOfItem(browser IBrowser, item objc.IObject) int {
	if d._BrowserNumberOfChildrenOfItem != nil {
		return d._BrowserNumberOfChildrenOfItem(browser, item)
	}
	var zero int
	return zero
}

// HasBrowserNumberOfChildrenOfItem returns true if a handler for BrowserNumberOfChildrenOfItem has been set.
func (d *BrowserDelegate) HasBrowserNumberOfChildrenOfItem() bool {
	return d._BrowserNumberOfChildrenOfItem != nil
}

// BrowserNumberOfRowsInColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserNumberOfRowsInColumn(sender IBrowser, column int) int {
	if d._BrowserNumberOfRowsInColumn != nil {
		return d._BrowserNumberOfRowsInColumn(sender, column)
	}
	var zero int
	return zero
}

// HasBrowserNumberOfRowsInColumn returns true if a handler for BrowserNumberOfRowsInColumn has been set.
func (d *BrowserDelegate) HasBrowserNumberOfRowsInColumn() bool {
	return d._BrowserNumberOfRowsInColumn != nil
}

// BrowserObjectValueForItem implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserObjectValueForItem(browser IBrowser, item objc.IObject) objc.ID {
	if d._BrowserObjectValueForItem != nil {
		return d._BrowserObjectValueForItem(browser, item)
	}
	var zero objc.ID
	return zero
}

// HasBrowserObjectValueForItem returns true if a handler for BrowserObjectValueForItem has been set.
func (d *BrowserDelegate) HasBrowserObjectValueForItem() bool {
	return d._BrowserObjectValueForItem != nil
}

// BrowserPreviewViewControllerForLeafItem implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserPreviewViewControllerForLeafItem(browser IBrowser, item objc.IObject) ViewController {
	if d._BrowserPreviewViewControllerForLeafItem != nil {
		return d._BrowserPreviewViewControllerForLeafItem(browser, item)
	}
	var zero ViewController
	return zero
}

// HasBrowserPreviewViewControllerForLeafItem returns true if a handler for BrowserPreviewViewControllerForLeafItem has been set.
func (d *BrowserDelegate) HasBrowserPreviewViewControllerForLeafItem() bool {
	return d._BrowserPreviewViewControllerForLeafItem != nil
}

// BrowserSelectCellWithStringInColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserSelectCellWithStringInColumn(sender IBrowser, title objc.IObject /* cross-framework: NSString */, column int) bool {
	if d._BrowserSelectCellWithStringInColumn != nil {
		return d._BrowserSelectCellWithStringInColumn(sender, title, column)
	}
	var zero bool
	return zero
}

// HasBrowserSelectCellWithStringInColumn returns true if a handler for BrowserSelectCellWithStringInColumn has been set.
func (d *BrowserDelegate) HasBrowserSelectCellWithStringInColumn() bool {
	return d._BrowserSelectCellWithStringInColumn != nil
}

// BrowserSelectRowInColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserSelectRowInColumn(sender IBrowser, row int, column int) bool {
	if d._BrowserSelectRowInColumn != nil {
		return d._BrowserSelectRowInColumn(sender, row, column)
	}
	var zero bool
	return zero
}

// HasBrowserSelectRowInColumn returns true if a handler for BrowserSelectRowInColumn has been set.
func (d *BrowserDelegate) HasBrowserSelectRowInColumn() bool {
	return d._BrowserSelectRowInColumn != nil
}

// BrowserSelectionIndexesForProposedSelectionInColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserSelectionIndexesForProposedSelectionInColumn(browser IBrowser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IndexSet {
	if d._BrowserSelectionIndexesForProposedSelectionInColumn != nil {
		return d._BrowserSelectionIndexesForProposedSelectionInColumn(browser, proposedSelectionIndexes, column)
	}
	var zero foundation.IndexSet
	return zero
}

// HasBrowserSelectionIndexesForProposedSelectionInColumn returns true if a handler for BrowserSelectionIndexesForProposedSelectionInColumn has been set.
func (d *BrowserDelegate) HasBrowserSelectionIndexesForProposedSelectionInColumn() bool {
	return d._BrowserSelectionIndexesForProposedSelectionInColumn != nil
}

// BrowserSetObjectValueForItem implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserSetObjectValueForItem(browser IBrowser, object objc.IObject, item objc.IObject) {
	if d._BrowserSetObjectValueForItem != nil {
		d._BrowserSetObjectValueForItem(browser, object, item)
	}
}

// HasBrowserSetObjectValueForItem returns true if a handler for BrowserSetObjectValueForItem has been set.
func (d *BrowserDelegate) HasBrowserSetObjectValueForItem() bool {
	return d._BrowserSetObjectValueForItem != nil
}

// BrowserShouldEditItem implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserShouldEditItem(browser IBrowser, item objc.IObject) bool {
	if d._BrowserShouldEditItem != nil {
		return d._BrowserShouldEditItem(browser, item)
	}
	var zero bool
	return zero
}

// HasBrowserShouldEditItem returns true if a handler for BrowserShouldEditItem has been set.
func (d *BrowserDelegate) HasBrowserShouldEditItem() bool {
	return d._BrowserShouldEditItem != nil
}

// BrowserShouldShowCellExpansionForRowColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserShouldShowCellExpansionForRowColumn(browser IBrowser, row int, column int) bool {
	if d._BrowserShouldShowCellExpansionForRowColumn != nil {
		return d._BrowserShouldShowCellExpansionForRowColumn(browser, row, column)
	}
	var zero bool
	return zero
}

// HasBrowserShouldShowCellExpansionForRowColumn returns true if a handler for BrowserShouldShowCellExpansionForRowColumn has been set.
func (d *BrowserDelegate) HasBrowserShouldShowCellExpansionForRowColumn() bool {
	return d._BrowserShouldShowCellExpansionForRowColumn != nil
}

// BrowserShouldSizeColumnForUserResizeToWidth implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserShouldSizeColumnForUserResizeToWidth(browser IBrowser, columnIndex int, forUserResize bool, suggestedWidth float64) float64 {
	if d._BrowserShouldSizeColumnForUserResizeToWidth != nil {
		return d._BrowserShouldSizeColumnForUserResizeToWidth(browser, columnIndex, forUserResize, suggestedWidth)
	}
	var zero float64
	return zero
}

// HasBrowserShouldSizeColumnForUserResizeToWidth returns true if a handler for BrowserShouldSizeColumnForUserResizeToWidth has been set.
func (d *BrowserDelegate) HasBrowserShouldSizeColumnForUserResizeToWidth() bool {
	return d._BrowserShouldSizeColumnForUserResizeToWidth != nil
}

// BrowserShouldTypeSelectForEventWithCurrentSearchString implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserShouldTypeSelectForEventWithCurrentSearchString(browser IBrowser, event IEvent, searchString objc.IObject /* cross-framework: NSString */) bool {
	if d._BrowserShouldTypeSelectForEventWithCurrentSearchString != nil {
		return d._BrowserShouldTypeSelectForEventWithCurrentSearchString(browser, event, searchString)
	}
	var zero bool
	return zero
}

// HasBrowserShouldTypeSelectForEventWithCurrentSearchString returns true if a handler for BrowserShouldTypeSelectForEventWithCurrentSearchString has been set.
func (d *BrowserDelegate) HasBrowserShouldTypeSelectForEventWithCurrentSearchString() bool {
	return d._BrowserShouldTypeSelectForEventWithCurrentSearchString != nil
}

// BrowserSizeToFitWidthOfColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserSizeToFitWidthOfColumn(browser IBrowser, columnIndex int) float64 {
	if d._BrowserSizeToFitWidthOfColumn != nil {
		return d._BrowserSizeToFitWidthOfColumn(browser, columnIndex)
	}
	var zero float64
	return zero
}

// HasBrowserSizeToFitWidthOfColumn returns true if a handler for BrowserSizeToFitWidthOfColumn has been set.
func (d *BrowserDelegate) HasBrowserSizeToFitWidthOfColumn() bool {
	return d._BrowserSizeToFitWidthOfColumn != nil
}

// BrowserTitleOfColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserTitleOfColumn(sender IBrowser, column int) foundation.String {
	if d._BrowserTitleOfColumn != nil {
		return d._BrowserTitleOfColumn(sender, column)
	}
	var zero foundation.String
	return zero
}

// HasBrowserTitleOfColumn returns true if a handler for BrowserTitleOfColumn has been set.
func (d *BrowserDelegate) HasBrowserTitleOfColumn() bool {
	return d._BrowserTitleOfColumn != nil
}

// BrowserTypeSelectStringForRowInColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserTypeSelectStringForRowInColumn(browser IBrowser, row int, column int) foundation.String {
	if d._BrowserTypeSelectStringForRowInColumn != nil {
		return d._BrowserTypeSelectStringForRowInColumn(browser, row, column)
	}
	var zero foundation.String
	return zero
}

// HasBrowserTypeSelectStringForRowInColumn returns true if a handler for BrowserTypeSelectStringForRowInColumn has been set.
func (d *BrowserDelegate) HasBrowserTypeSelectStringForRowInColumn() bool {
	return d._BrowserTypeSelectStringForRowInColumn != nil
}

// BrowserValidateDropProposedRowColumnDropOperation implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserValidateDropProposedRowColumnDropOperation(browser IBrowser, info unsafe.Pointer, row int, column int, dropOperation BrowserDropOperation) DragOperation {
	if d._BrowserValidateDropProposedRowColumnDropOperation != nil {
		return d._BrowserValidateDropProposedRowColumnDropOperation(browser, info, row, column, dropOperation)
	}
	var zero DragOperation
	return zero
}

// HasBrowserValidateDropProposedRowColumnDropOperation returns true if a handler for BrowserValidateDropProposedRowColumnDropOperation has been set.
func (d *BrowserDelegate) HasBrowserValidateDropProposedRowColumnDropOperation() bool {
	return d._BrowserValidateDropProposedRowColumnDropOperation != nil
}

// BrowserWillDisplayCellAtRowColumn implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserWillDisplayCellAtRowColumn(sender IBrowser, cell objc.IObject, row int, column int) {
	if d._BrowserWillDisplayCellAtRowColumn != nil {
		d._BrowserWillDisplayCellAtRowColumn(sender, cell, row, column)
	}
}

// HasBrowserWillDisplayCellAtRowColumn returns true if a handler for BrowserWillDisplayCellAtRowColumn has been set.
func (d *BrowserDelegate) HasBrowserWillDisplayCellAtRowColumn() bool {
	return d._BrowserWillDisplayCellAtRowColumn != nil
}

// BrowserWriteRowsWithIndexesInColumnToPasteboard implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserWriteRowsWithIndexesInColumnToPasteboard(browser IBrowser, rowIndexes foundation.IndexSet, column int, pasteboard IPasteboard) bool {
	if d._BrowserWriteRowsWithIndexesInColumnToPasteboard != nil {
		return d._BrowserWriteRowsWithIndexesInColumnToPasteboard(browser, rowIndexes, column, pasteboard)
	}
	var zero bool
	return zero
}

// HasBrowserWriteRowsWithIndexesInColumnToPasteboard returns true if a handler for BrowserWriteRowsWithIndexesInColumnToPasteboard has been set.
func (d *BrowserDelegate) HasBrowserWriteRowsWithIndexesInColumnToPasteboard() bool {
	return d._BrowserWriteRowsWithIndexesInColumnToPasteboard != nil
}

// BrowserColumnConfigurationDidChange implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserColumnConfigurationDidChange(notification foundation.Notification) {
	if d._BrowserColumnConfigurationDidChange != nil {
		d._BrowserColumnConfigurationDidChange(notification)
	}
}

// HasBrowserColumnConfigurationDidChange returns true if a handler for BrowserColumnConfigurationDidChange has been set.
func (d *BrowserDelegate) HasBrowserColumnConfigurationDidChange() bool {
	return d._BrowserColumnConfigurationDidChange != nil
}

// BrowserDidScroll implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserDidScroll(sender IBrowser) {
	if d._BrowserDidScroll != nil {
		d._BrowserDidScroll(sender)
	}
}

// HasBrowserDidScroll returns true if a handler for BrowserDidScroll has been set.
func (d *BrowserDelegate) HasBrowserDidScroll() bool {
	return d._BrowserDidScroll != nil
}

// BrowserWillScroll implements the PBrowserDelegate interface.
func (d *BrowserDelegate) BrowserWillScroll(sender IBrowser) {
	if d._BrowserWillScroll != nil {
		d._BrowserWillScroll(sender)
	}
}

// HasBrowserWillScroll returns true if a handler for BrowserWillScroll has been set.
func (d *BrowserDelegate) HasBrowserWillScroll() bool {
	return d._BrowserWillScroll != nil
}

// RootItemForBrowser implements the PBrowserDelegate interface.
func (d *BrowserDelegate) RootItemForBrowser(browser IBrowser) objc.ID {
	if d._RootItemForBrowser != nil {
		return d._RootItemForBrowser(browser)
	}
	var zero objc.ID
	return zero
}

// HasRootItemForBrowser returns true if a handler for RootItemForBrowser has been set.
func (d *BrowserDelegate) HasRootItemForBrowser() bool {
	return d._RootItemForBrowser != nil
}
