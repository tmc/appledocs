// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/vision"
)

// POutlineViewDataSource is the NSOutlineViewDataSource protocol interface.
//
// A set of methods that an outline view calls to retrieve data and information about it from the data source delegate, and—optionally—to update data values.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSOutlineViewDataSource
type POutlineViewDataSource interface {
	// Optional methods
	OutlineViewAcceptDropItemChildIndex(outlineView IOutlineView, info unsafe.Pointer, item objc.IObject, index int) bool
	HasOutlineViewAcceptDropItemChildIndex() bool
	OutlineViewChildOfItem(outlineView IOutlineView, index int, item objc.IObject) objc.ID
	HasOutlineViewChildOfItem() bool
	OutlineViewDraggingSessionEndedAtPointOperation(outlineView IOutlineView, session IDraggingSession, screenPoint vision.Point, operation DragOperation)
	HasOutlineViewDraggingSessionEndedAtPointOperation() bool
	OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView IOutlineView, session IDraggingSession, screenPoint vision.Point, draggedItems objc.IObject /* cross-framework: NSArray */)
	HasOutlineViewDraggingSessionWillBeginAtPointForItems() bool
	OutlineViewIsItemExpandable(outlineView IOutlineView, item objc.IObject) bool
	HasOutlineViewIsItemExpandable() bool
	OutlineViewItemForPersistentObject(outlineView IOutlineView, object objc.IObject) objc.ID
	HasOutlineViewItemForPersistentObject() bool
	OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems(outlineView IOutlineView, dropDestination objc.IObject /* cross-framework: NSURL */, items objc.IObject /* cross-framework: NSArray */) []string
	HasOutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems() bool
	OutlineViewNumberOfChildrenOfItem(outlineView IOutlineView, item objc.IObject) int
	HasOutlineViewNumberOfChildrenOfItem() bool
	OutlineViewObjectValueForTableColumnByItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) objc.ID
	HasOutlineViewObjectValueForTableColumnByItem() bool
	OutlineViewPasteboardWriterForItem(outlineView IOutlineView, item objc.IObject) unsafe.Pointer
	HasOutlineViewPasteboardWriterForItem() bool
	OutlineViewPersistentObjectForItem(outlineView IOutlineView, item objc.IObject) objc.ID
	HasOutlineViewPersistentObjectForItem() bool
	OutlineViewSetObjectValueForTableColumnByItem(outlineView IOutlineView, object objc.IObject, tableColumn ITableColumn, item objc.IObject)
	HasOutlineViewSetObjectValueForTableColumnByItem() bool
	OutlineViewSortDescriptorsDidChange(outlineView IOutlineView, oldDescriptors []objc.IObject)
	HasOutlineViewSortDescriptorsDidChange() bool
	OutlineViewUpdateDraggingItemsForDrag(outlineView IOutlineView, draggingInfo unsafe.Pointer)
	HasOutlineViewUpdateDraggingItemsForDrag() bool
	OutlineViewValidateDropProposedItemProposedChildIndex(outlineView IOutlineView, info unsafe.Pointer, item objc.IObject, index int) DragOperation
	HasOutlineViewValidateDropProposedItemProposedChildIndex() bool
	OutlineViewWriteItemsToPasteboard(outlineView IOutlineView, items objc.IObject /* cross-framework: NSArray */, pasteboard IPasteboard) bool
	HasOutlineViewWriteItemsToPasteboard() bool
}

// OutlineViewDataSource is a delegate implementation builder for the POutlineViewDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type OutlineViewDataSource struct {
	_OutlineViewAcceptDropItemChildIndex func(outlineView IOutlineView, info unsafe.Pointer, item objc.IObject, index int) bool
	_OutlineViewChildOfItem func(outlineView IOutlineView, index int, item objc.IObject) objc.ID
	_OutlineViewDraggingSessionEndedAtPointOperation func(outlineView IOutlineView, session IDraggingSession, screenPoint vision.Point, operation DragOperation)
	_OutlineViewDraggingSessionWillBeginAtPointForItems func(outlineView IOutlineView, session IDraggingSession, screenPoint vision.Point, draggedItems objc.IObject /* cross-framework: NSArray */)
	_OutlineViewIsItemExpandable func(outlineView IOutlineView, item objc.IObject) bool
	_OutlineViewItemForPersistentObject func(outlineView IOutlineView, object objc.IObject) objc.ID
	_OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems func(outlineView IOutlineView, dropDestination objc.IObject /* cross-framework: NSURL */, items objc.IObject /* cross-framework: NSArray */) []string
	_OutlineViewNumberOfChildrenOfItem func(outlineView IOutlineView, item objc.IObject) int
	_OutlineViewObjectValueForTableColumnByItem func(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) objc.ID
	_OutlineViewPasteboardWriterForItem func(outlineView IOutlineView, item objc.IObject) unsafe.Pointer
	_OutlineViewPersistentObjectForItem func(outlineView IOutlineView, item objc.IObject) objc.ID
	_OutlineViewSetObjectValueForTableColumnByItem func(outlineView IOutlineView, object objc.IObject, tableColumn ITableColumn, item objc.IObject)
	_OutlineViewSortDescriptorsDidChange func(outlineView IOutlineView, oldDescriptors []objc.IObject)
	_OutlineViewUpdateDraggingItemsForDrag func(outlineView IOutlineView, draggingInfo unsafe.Pointer)
	_OutlineViewValidateDropProposedItemProposedChildIndex func(outlineView IOutlineView, info unsafe.Pointer, item objc.IObject, index int) DragOperation
	_OutlineViewWriteItemsToPasteboard func(outlineView IOutlineView, items objc.IObject /* cross-framework: NSArray */, pasteboard IPasteboard) bool
}

// SetOutlineViewAcceptDropItemChildIndex sets the handler for the OutlineViewAcceptDropItemChildIndex delegate method.
//
// Returns a Boolean value that indicates whether a drop operation was successful.
func (d *OutlineViewDataSource) SetOutlineViewAcceptDropItemChildIndex(f func(outlineView IOutlineView, info unsafe.Pointer, item objc.IObject, index int) bool) {
	d._OutlineViewAcceptDropItemChildIndex = f
}

// SetOutlineViewChildOfItem sets the handler for the OutlineViewChildOfItem delegate method.
//
// Returns the child item at the specified index of a given item.
func (d *OutlineViewDataSource) SetOutlineViewChildOfItem(f func(outlineView IOutlineView, index int, item objc.IObject) objc.ID) {
	d._OutlineViewChildOfItem = f
}

// SetOutlineViewDraggingSessionEndedAtPointOperation sets the handler for the OutlineViewDraggingSessionEndedAtPointOperation delegate method.
//
// Implement this method to know when the given dragging session has ended.
func (d *OutlineViewDataSource) SetOutlineViewDraggingSessionEndedAtPointOperation(f func(outlineView IOutlineView, session IDraggingSession, screenPoint vision.Point, operation DragOperation)) {
	d._OutlineViewDraggingSessionEndedAtPointOperation = f
}

// SetOutlineViewDraggingSessionWillBeginAtPointForItems sets the handler for the OutlineViewDraggingSessionWillBeginAtPointForItems delegate method.
//
// Implement this method know when the given dragging session is about to begin and potentially modify the dragging session.
func (d *OutlineViewDataSource) SetOutlineViewDraggingSessionWillBeginAtPointForItems(f func(outlineView IOutlineView, session IDraggingSession, screenPoint vision.Point, draggedItems objc.IObject /* cross-framework: NSArray */)) {
	d._OutlineViewDraggingSessionWillBeginAtPointForItems = f
}

// SetOutlineViewIsItemExpandable sets the handler for the OutlineViewIsItemExpandable delegate method.
//
// Returns a Boolean value that indicates whether the a given item is expandable.
func (d *OutlineViewDataSource) SetOutlineViewIsItemExpandable(f func(outlineView IOutlineView, item objc.IObject) bool) {
	d._OutlineViewIsItemExpandable = f
}

// SetOutlineViewItemForPersistentObject sets the handler for the OutlineViewItemForPersistentObject delegate method.
//
// Invoked by   to return the item for the archived  .
func (d *OutlineViewDataSource) SetOutlineViewItemForPersistentObject(f func(outlineView IOutlineView, object objc.IObject) objc.ID) {
	d._OutlineViewItemForPersistentObject = f
}

// SetOutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems sets the handler for the OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems delegate method.
//
// Returns an array of filenames for the created files that the receiver promises to create.
func (d *OutlineViewDataSource) SetOutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems(f func(outlineView IOutlineView, dropDestination objc.IObject /* cross-framework: NSURL */, items objc.IObject /* cross-framework: NSArray */) []string) {
	d._OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems = f
}

// SetOutlineViewNumberOfChildrenOfItem sets the handler for the OutlineViewNumberOfChildrenOfItem delegate method.
//
// Returns the number of child items encompassed by a given item.
func (d *OutlineViewDataSource) SetOutlineViewNumberOfChildrenOfItem(f func(outlineView IOutlineView, item objc.IObject) int) {
	d._OutlineViewNumberOfChildrenOfItem = f
}

// SetOutlineViewObjectValueForTableColumnByItem sets the handler for the OutlineViewObjectValueForTableColumnByItem delegate method.
//
// Invoked by   to return the data object associated with the specified  .
func (d *OutlineViewDataSource) SetOutlineViewObjectValueForTableColumnByItem(f func(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) objc.ID) {
	d._OutlineViewObjectValueForTableColumnByItem = f
}

// SetOutlineViewPasteboardWriterForItem sets the handler for the OutlineViewPasteboardWriterForItem delegate method.
//
// Implement this method to enable the table to be an   that supports dragging multiple items.
func (d *OutlineViewDataSource) SetOutlineViewPasteboardWriterForItem(f func(outlineView IOutlineView, item objc.IObject) unsafe.Pointer) {
	d._OutlineViewPasteboardWriterForItem = f
}

// SetOutlineViewPersistentObjectForItem sets the handler for the OutlineViewPersistentObjectForItem delegate method.
//
// Invoked by   to return an archived object for  .
func (d *OutlineViewDataSource) SetOutlineViewPersistentObjectForItem(f func(outlineView IOutlineView, item objc.IObject) objc.ID) {
	d._OutlineViewPersistentObjectForItem = f
}

// SetOutlineViewSetObjectValueForTableColumnByItem sets the handler for the OutlineViewSetObjectValueForTableColumnByItem delegate method.
//
// Set the data object for a given item in a given column.
func (d *OutlineViewDataSource) SetOutlineViewSetObjectValueForTableColumnByItem(f func(outlineView IOutlineView, object objc.IObject, tableColumn ITableColumn, item objc.IObject)) {
	d._OutlineViewSetObjectValueForTableColumnByItem = f
}

// SetOutlineViewSortDescriptorsDidChange sets the handler for the OutlineViewSortDescriptorsDidChange delegate method.
//
// Invoked by an outline view to notify the data source that the descriptors changed and the data may need to be resorted.
func (d *OutlineViewDataSource) SetOutlineViewSortDescriptorsDidChange(f func(outlineView IOutlineView, oldDescriptors []objc.IObject)) {
	d._OutlineViewSortDescriptorsDidChange = f
}

// SetOutlineViewUpdateDraggingItemsForDrag sets the handler for the OutlineViewUpdateDraggingItemsForDrag delegate method.
//
// Implement this method to enable the table to update dragging items as they are dragged over the view.
func (d *OutlineViewDataSource) SetOutlineViewUpdateDraggingItemsForDrag(f func(outlineView IOutlineView, draggingInfo unsafe.Pointer)) {
	d._OutlineViewUpdateDraggingItemsForDrag = f
}

// SetOutlineViewValidateDropProposedItemProposedChildIndex sets the handler for the OutlineViewValidateDropProposedItemProposedChildIndex delegate method.
//
// Used by an outline view to determine a valid drop target.
func (d *OutlineViewDataSource) SetOutlineViewValidateDropProposedItemProposedChildIndex(f func(outlineView IOutlineView, info unsafe.Pointer, item objc.IObject, index int) DragOperation) {
	d._OutlineViewValidateDropProposedItemProposedChildIndex = f
}

// SetOutlineViewWriteItemsToPasteboard sets the handler for the OutlineViewWriteItemsToPasteboard delegate method.
//
// Returns a Boolean value that indicates whether a drag operation is allowed.
func (d *OutlineViewDataSource) SetOutlineViewWriteItemsToPasteboard(f func(outlineView IOutlineView, items objc.IObject /* cross-framework: NSArray */, pasteboard IPasteboard) bool) {
	d._OutlineViewWriteItemsToPasteboard = f
}

// OutlineViewAcceptDropItemChildIndex implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewAcceptDropItemChildIndex(outlineView IOutlineView, info unsafe.Pointer, item objc.IObject, index int) bool {
	if d._OutlineViewAcceptDropItemChildIndex != nil {
		return d._OutlineViewAcceptDropItemChildIndex(outlineView, info, item, index)
	}
	var zero bool
	return zero
}

// HasOutlineViewAcceptDropItemChildIndex returns true if a handler for OutlineViewAcceptDropItemChildIndex has been set.
func (d *OutlineViewDataSource) HasOutlineViewAcceptDropItemChildIndex() bool {
	return d._OutlineViewAcceptDropItemChildIndex != nil
}

// OutlineViewChildOfItem implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewChildOfItem(outlineView IOutlineView, index int, item objc.IObject) objc.ID {
	if d._OutlineViewChildOfItem != nil {
		return d._OutlineViewChildOfItem(outlineView, index, item)
	}
	var zero objc.ID
	return zero
}

// HasOutlineViewChildOfItem returns true if a handler for OutlineViewChildOfItem has been set.
func (d *OutlineViewDataSource) HasOutlineViewChildOfItem() bool {
	return d._OutlineViewChildOfItem != nil
}

// OutlineViewDraggingSessionEndedAtPointOperation implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewDraggingSessionEndedAtPointOperation(outlineView IOutlineView, session IDraggingSession, screenPoint vision.Point, operation DragOperation) {
	if d._OutlineViewDraggingSessionEndedAtPointOperation != nil {
		d._OutlineViewDraggingSessionEndedAtPointOperation(outlineView, session, screenPoint, operation)
	}
}

// HasOutlineViewDraggingSessionEndedAtPointOperation returns true if a handler for OutlineViewDraggingSessionEndedAtPointOperation has been set.
func (d *OutlineViewDataSource) HasOutlineViewDraggingSessionEndedAtPointOperation() bool {
	return d._OutlineViewDraggingSessionEndedAtPointOperation != nil
}

// OutlineViewDraggingSessionWillBeginAtPointForItems implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView IOutlineView, session IDraggingSession, screenPoint vision.Point, draggedItems objc.IObject /* cross-framework: NSArray */) {
	if d._OutlineViewDraggingSessionWillBeginAtPointForItems != nil {
		d._OutlineViewDraggingSessionWillBeginAtPointForItems(outlineView, session, screenPoint, draggedItems)
	}
}

// HasOutlineViewDraggingSessionWillBeginAtPointForItems returns true if a handler for OutlineViewDraggingSessionWillBeginAtPointForItems has been set.
func (d *OutlineViewDataSource) HasOutlineViewDraggingSessionWillBeginAtPointForItems() bool {
	return d._OutlineViewDraggingSessionWillBeginAtPointForItems != nil
}

// OutlineViewIsItemExpandable implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewIsItemExpandable(outlineView IOutlineView, item objc.IObject) bool {
	if d._OutlineViewIsItemExpandable != nil {
		return d._OutlineViewIsItemExpandable(outlineView, item)
	}
	var zero bool
	return zero
}

// HasOutlineViewIsItemExpandable returns true if a handler for OutlineViewIsItemExpandable has been set.
func (d *OutlineViewDataSource) HasOutlineViewIsItemExpandable() bool {
	return d._OutlineViewIsItemExpandable != nil
}

// OutlineViewItemForPersistentObject implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewItemForPersistentObject(outlineView IOutlineView, object objc.IObject) objc.ID {
	if d._OutlineViewItemForPersistentObject != nil {
		return d._OutlineViewItemForPersistentObject(outlineView, object)
	}
	var zero objc.ID
	return zero
}

// HasOutlineViewItemForPersistentObject returns true if a handler for OutlineViewItemForPersistentObject has been set.
func (d *OutlineViewDataSource) HasOutlineViewItemForPersistentObject() bool {
	return d._OutlineViewItemForPersistentObject != nil
}

// OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems(outlineView IOutlineView, dropDestination objc.IObject /* cross-framework: NSURL */, items objc.IObject /* cross-framework: NSArray */) []string {
	if d._OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems != nil {
		return d._OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems(outlineView, dropDestination, items)
	}
	var zero []string
	return zero
}

// HasOutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems returns true if a handler for OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems has been set.
func (d *OutlineViewDataSource) HasOutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems() bool {
	return d._OutlineViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItems != nil
}

// OutlineViewNumberOfChildrenOfItem implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewNumberOfChildrenOfItem(outlineView IOutlineView, item objc.IObject) int {
	if d._OutlineViewNumberOfChildrenOfItem != nil {
		return d._OutlineViewNumberOfChildrenOfItem(outlineView, item)
	}
	var zero int
	return zero
}

// HasOutlineViewNumberOfChildrenOfItem returns true if a handler for OutlineViewNumberOfChildrenOfItem has been set.
func (d *OutlineViewDataSource) HasOutlineViewNumberOfChildrenOfItem() bool {
	return d._OutlineViewNumberOfChildrenOfItem != nil
}

// OutlineViewObjectValueForTableColumnByItem implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewObjectValueForTableColumnByItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) objc.ID {
	if d._OutlineViewObjectValueForTableColumnByItem != nil {
		return d._OutlineViewObjectValueForTableColumnByItem(outlineView, tableColumn, item)
	}
	var zero objc.ID
	return zero
}

// HasOutlineViewObjectValueForTableColumnByItem returns true if a handler for OutlineViewObjectValueForTableColumnByItem has been set.
func (d *OutlineViewDataSource) HasOutlineViewObjectValueForTableColumnByItem() bool {
	return d._OutlineViewObjectValueForTableColumnByItem != nil
}

// OutlineViewPasteboardWriterForItem implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewPasteboardWriterForItem(outlineView IOutlineView, item objc.IObject) unsafe.Pointer {
	if d._OutlineViewPasteboardWriterForItem != nil {
		return d._OutlineViewPasteboardWriterForItem(outlineView, item)
	}
	var zero unsafe.Pointer
	return zero
}

// HasOutlineViewPasteboardWriterForItem returns true if a handler for OutlineViewPasteboardWriterForItem has been set.
func (d *OutlineViewDataSource) HasOutlineViewPasteboardWriterForItem() bool {
	return d._OutlineViewPasteboardWriterForItem != nil
}

// OutlineViewPersistentObjectForItem implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewPersistentObjectForItem(outlineView IOutlineView, item objc.IObject) objc.ID {
	if d._OutlineViewPersistentObjectForItem != nil {
		return d._OutlineViewPersistentObjectForItem(outlineView, item)
	}
	var zero objc.ID
	return zero
}

// HasOutlineViewPersistentObjectForItem returns true if a handler for OutlineViewPersistentObjectForItem has been set.
func (d *OutlineViewDataSource) HasOutlineViewPersistentObjectForItem() bool {
	return d._OutlineViewPersistentObjectForItem != nil
}

// OutlineViewSetObjectValueForTableColumnByItem implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewSetObjectValueForTableColumnByItem(outlineView IOutlineView, object objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	if d._OutlineViewSetObjectValueForTableColumnByItem != nil {
		d._OutlineViewSetObjectValueForTableColumnByItem(outlineView, object, tableColumn, item)
	}
}

// HasOutlineViewSetObjectValueForTableColumnByItem returns true if a handler for OutlineViewSetObjectValueForTableColumnByItem has been set.
func (d *OutlineViewDataSource) HasOutlineViewSetObjectValueForTableColumnByItem() bool {
	return d._OutlineViewSetObjectValueForTableColumnByItem != nil
}

// OutlineViewSortDescriptorsDidChange implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewSortDescriptorsDidChange(outlineView IOutlineView, oldDescriptors []objc.IObject) {
	if d._OutlineViewSortDescriptorsDidChange != nil {
		d._OutlineViewSortDescriptorsDidChange(outlineView, oldDescriptors)
	}
}

// HasOutlineViewSortDescriptorsDidChange returns true if a handler for OutlineViewSortDescriptorsDidChange has been set.
func (d *OutlineViewDataSource) HasOutlineViewSortDescriptorsDidChange() bool {
	return d._OutlineViewSortDescriptorsDidChange != nil
}

// OutlineViewUpdateDraggingItemsForDrag implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewUpdateDraggingItemsForDrag(outlineView IOutlineView, draggingInfo unsafe.Pointer) {
	if d._OutlineViewUpdateDraggingItemsForDrag != nil {
		d._OutlineViewUpdateDraggingItemsForDrag(outlineView, draggingInfo)
	}
}

// HasOutlineViewUpdateDraggingItemsForDrag returns true if a handler for OutlineViewUpdateDraggingItemsForDrag has been set.
func (d *OutlineViewDataSource) HasOutlineViewUpdateDraggingItemsForDrag() bool {
	return d._OutlineViewUpdateDraggingItemsForDrag != nil
}

// OutlineViewValidateDropProposedItemProposedChildIndex implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewValidateDropProposedItemProposedChildIndex(outlineView IOutlineView, info unsafe.Pointer, item objc.IObject, index int) DragOperation {
	if d._OutlineViewValidateDropProposedItemProposedChildIndex != nil {
		return d._OutlineViewValidateDropProposedItemProposedChildIndex(outlineView, info, item, index)
	}
	var zero DragOperation
	return zero
}

// HasOutlineViewValidateDropProposedItemProposedChildIndex returns true if a handler for OutlineViewValidateDropProposedItemProposedChildIndex has been set.
func (d *OutlineViewDataSource) HasOutlineViewValidateDropProposedItemProposedChildIndex() bool {
	return d._OutlineViewValidateDropProposedItemProposedChildIndex != nil
}

// OutlineViewWriteItemsToPasteboard implements the POutlineViewDataSource interface.
func (d *OutlineViewDataSource) OutlineViewWriteItemsToPasteboard(outlineView IOutlineView, items objc.IObject /* cross-framework: NSArray */, pasteboard IPasteboard) bool {
	if d._OutlineViewWriteItemsToPasteboard != nil {
		return d._OutlineViewWriteItemsToPasteboard(outlineView, items, pasteboard)
	}
	var zero bool
	return zero
}

// HasOutlineViewWriteItemsToPasteboard returns true if a handler for OutlineViewWriteItemsToPasteboard has been set.
func (d *OutlineViewDataSource) HasOutlineViewWriteItemsToPasteboard() bool {
	return d._OutlineViewWriteItemsToPasteboard != nil
}
