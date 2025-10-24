// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"

	"github.com/tmc/appledocs/generated/vision"
)

// PCollectionViewDelegate is the NSCollectionViewDelegate protocol interface.
//
// A set of methods that you use to manage the behavior of a collection view.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSCollectionViewDelegate
type PCollectionViewDelegate interface {
	// Optional methods
	CollectionViewAcceptDropIndexDropOperation(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, index int, dropOperation CollectionViewDropOperation) bool
	HasCollectionViewAcceptDropIndexDropOperation() bool
	CollectionViewAcceptDropIndexPathDropOperation(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	HasCollectionViewAcceptDropIndexPathDropOperation() bool
	CollectionViewCanDragItemsAtIndexesWithEvent(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, event IEvent) bool
	HasCollectionViewCanDragItemsAtIndexesWithEvent() bool
	CollectionViewCanDragItemsAtIndexPathsWithEvent(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, event IEvent) bool
	HasCollectionViewCanDragItemsAtIndexPathsWithEvent() bool
	CollectionViewDidChangeItemsAtIndexPathsToHighlightState(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, highlightState CollectionViewItemHighlightState)
	HasCollectionViewDidChangeItemsAtIndexPathsToHighlightState() bool
	CollectionViewDidDeselectItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer)
	HasCollectionViewDidDeselectItemsAtIndexPaths() bool
	CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, item ICollectionViewItem, indexPath foundation.IndexPath)
	HasCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath() bool
	CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, view IView, elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath)
	HasCollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath() bool
	CollectionViewDidSelectItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer)
	HasCollectionViewDidSelectItemsAtIndexPaths() bool
	CollectionViewDraggingImageForItemsAtIndexesWithEventOffset(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, event IEvent, dragImageOffset PointPointer /* not a class type */) Image
	HasCollectionViewDraggingImageForItemsAtIndexesWithEventOffset() bool
	CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, event IEvent, dragImageOffset PointPointer /* not a class type */) Image
	HasCollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset() bool
	CollectionViewDraggingSessionEndedAtPointDragOperation(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, operation DragOperation)
	HasCollectionViewDraggingSessionEndedAtPointDragOperation() bool
	CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, indexPaths unsafe.Pointer)
	HasCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths() bool
	CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, indexes foundation.IndexSet)
	HasCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes() bool
	CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, dropURL objc.IObject /* cross-framework: NSURL */, indexPaths unsafe.Pointer) []string
	HasCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths() bool
	CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes(collectionView objc.IObject /* cross-framework: CollectionView */, dropURL objc.IObject /* cross-framework: NSURL */, indexes foundation.IndexSet) []string
	HasCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes() bool
	CollectionViewPasteboardWriterForItemAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, indexPath foundation.IndexPath) unsafe.Pointer
	HasCollectionViewPasteboardWriterForItemAtIndexPath() bool
	CollectionViewPasteboardWriterForItemAtIndex(collectionView objc.IObject /* cross-framework: CollectionView */, index uint) unsafe.Pointer
	HasCollectionViewPasteboardWriterForItemAtIndex() bool
	CollectionViewShouldChangeItemsAtIndexPathsToHighlightState(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, highlightState CollectionViewItemHighlightState) unsafe.Pointer
	HasCollectionViewShouldChangeItemsAtIndexPathsToHighlightState() bool
	CollectionViewShouldDeselectItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) unsafe.Pointer
	HasCollectionViewShouldDeselectItemsAtIndexPaths() bool
	CollectionViewShouldSelectItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) unsafe.Pointer
	HasCollectionViewShouldSelectItemsAtIndexPaths() bool
	CollectionViewTransitionLayoutForOldLayoutNewLayout(collectionView objc.IObject /* cross-framework: CollectionView */, fromLayout ICollectionViewLayout, toLayout ICollectionViewLayout) CollectionViewTransitionLayout
	HasCollectionViewTransitionLayoutForOldLayoutNewLayout() bool
	CollectionViewUpdateDraggingItemsForDrag(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer)
	HasCollectionViewUpdateDraggingItemsForDrag() bool
	CollectionViewValidateDropProposedIndexDropOperation(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, proposedDropIndex int, proposedDropOperation CollectionViewDropOperation) DragOperation
	HasCollectionViewValidateDropProposedIndexDropOperation() bool
	CollectionViewValidateDropProposedIndexPathDropOperation(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, proposedDropIndexPath objectivec.IObject, proposedDropOperation CollectionViewDropOperation) DragOperation
	HasCollectionViewValidateDropProposedIndexPathDropOperation() bool
	CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, item ICollectionViewItem, indexPath foundation.IndexPath)
	HasCollectionViewWillDisplayItemForRepresentedObjectAtIndexPath() bool
	CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, view IView, elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath)
	HasCollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath() bool
	CollectionViewWriteItemsAtIndexPathsToPasteboard(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, pasteboard IPasteboard) bool
	HasCollectionViewWriteItemsAtIndexPathsToPasteboard() bool
	CollectionViewWriteItemsAtIndexesToPasteboard(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, pasteboard IPasteboard) bool
	HasCollectionViewWriteItemsAtIndexesToPasteboard() bool
}

// CollectionViewDelegate is a delegate implementation builder for the PCollectionViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CollectionViewDelegate struct {
	_CollectionViewAcceptDropIndexDropOperation func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, index int, dropOperation CollectionViewDropOperation) bool
	_CollectionViewAcceptDropIndexPathDropOperation func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	_CollectionViewCanDragItemsAtIndexesWithEvent func(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, event IEvent) bool
	_CollectionViewCanDragItemsAtIndexPathsWithEvent func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, event IEvent) bool
	_CollectionViewDidChangeItemsAtIndexPathsToHighlightState func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, highlightState CollectionViewItemHighlightState)
	_CollectionViewDidDeselectItemsAtIndexPaths func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer)
	_CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath func(collectionView objc.IObject /* cross-framework: CollectionView */, item ICollectionViewItem, indexPath foundation.IndexPath)
	_CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath func(collectionView objc.IObject /* cross-framework: CollectionView */, view IView, elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath)
	_CollectionViewDidSelectItemsAtIndexPaths func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer)
	_CollectionViewDraggingImageForItemsAtIndexesWithEventOffset func(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, event IEvent, dragImageOffset PointPointer /* not a class type */) Image
	_CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, event IEvent, dragImageOffset PointPointer /* not a class type */) Image
	_CollectionViewDraggingSessionEndedAtPointDragOperation func(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, operation DragOperation)
	_CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths func(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, indexPaths unsafe.Pointer)
	_CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes func(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, indexes foundation.IndexSet)
	_CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths func(collectionView objc.IObject /* cross-framework: CollectionView */, dropURL objc.IObject /* cross-framework: NSURL */, indexPaths unsafe.Pointer) []string
	_CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes func(collectionView objc.IObject /* cross-framework: CollectionView */, dropURL objc.IObject /* cross-framework: NSURL */, indexes foundation.IndexSet) []string
	_CollectionViewPasteboardWriterForItemAtIndexPath func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPath foundation.IndexPath) unsafe.Pointer
	_CollectionViewPasteboardWriterForItemAtIndex func(collectionView objc.IObject /* cross-framework: CollectionView */, index uint) unsafe.Pointer
	_CollectionViewShouldChangeItemsAtIndexPathsToHighlightState func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, highlightState CollectionViewItemHighlightState) unsafe.Pointer
	_CollectionViewShouldDeselectItemsAtIndexPaths func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) unsafe.Pointer
	_CollectionViewShouldSelectItemsAtIndexPaths func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) unsafe.Pointer
	_CollectionViewTransitionLayoutForOldLayoutNewLayout func(collectionView objc.IObject /* cross-framework: CollectionView */, fromLayout ICollectionViewLayout, toLayout ICollectionViewLayout) CollectionViewTransitionLayout
	_CollectionViewUpdateDraggingItemsForDrag func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer)
	_CollectionViewValidateDropProposedIndexDropOperation func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, proposedDropIndex int, proposedDropOperation CollectionViewDropOperation) DragOperation
	_CollectionViewValidateDropProposedIndexPathDropOperation func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, proposedDropIndexPath objectivec.IObject, proposedDropOperation CollectionViewDropOperation) DragOperation
	_CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath func(collectionView objc.IObject /* cross-framework: CollectionView */, item ICollectionViewItem, indexPath foundation.IndexPath)
	_CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath func(collectionView objc.IObject /* cross-framework: CollectionView */, view IView, elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath)
	_CollectionViewWriteItemsAtIndexPathsToPasteboard func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, pasteboard IPasteboard) bool
	_CollectionViewWriteItemsAtIndexesToPasteboard func(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, pasteboard IPasteboard) bool
}

// SetCollectionViewAcceptDropIndexDropOperation sets the handler for the CollectionViewAcceptDropIndexDropOperation delegate method.
//
// Invoked when the mouse is released over a collection view that previously allowed a drop.
func (d *CollectionViewDelegate) SetCollectionViewAcceptDropIndexDropOperation(f func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, index int, dropOperation CollectionViewDropOperation) bool) {
	d._CollectionViewAcceptDropIndexDropOperation = f
}

// SetCollectionViewAcceptDropIndexPathDropOperation sets the handler for the CollectionViewAcceptDropIndexPathDropOperation delegate method.
//
// Incorporates the dropped content into the collection view.
func (d *CollectionViewDelegate) SetCollectionViewAcceptDropIndexPathDropOperation(f func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool) {
	d._CollectionViewAcceptDropIndexPathDropOperation = f
}

// SetCollectionViewCanDragItemsAtIndexesWithEvent sets the handler for the CollectionViewCanDragItemsAtIndexesWithEvent delegate method.
//
// Returns a Boolean indicating whether the collection view can begin dragging the specified items.
func (d *CollectionViewDelegate) SetCollectionViewCanDragItemsAtIndexesWithEvent(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, event IEvent) bool) {
	d._CollectionViewCanDragItemsAtIndexesWithEvent = f
}

// SetCollectionViewCanDragItemsAtIndexPathsWithEvent sets the handler for the CollectionViewCanDragItemsAtIndexPathsWithEvent delegate method.
//
// Returns a Boolean indicating whether a drag operation involving the specified items can begin.
func (d *CollectionViewDelegate) SetCollectionViewCanDragItemsAtIndexPathsWithEvent(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, event IEvent) bool) {
	d._CollectionViewCanDragItemsAtIndexPathsWithEvent = f
}

// SetCollectionViewDidChangeItemsAtIndexPathsToHighlightState sets the handler for the CollectionViewDidChangeItemsAtIndexPathsToHighlightState delegate method.
//
// Notifies the delegate that the highlight state of the specified items changed.
func (d *CollectionViewDelegate) SetCollectionViewDidChangeItemsAtIndexPathsToHighlightState(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, highlightState CollectionViewItemHighlightState)) {
	d._CollectionViewDidChangeItemsAtIndexPathsToHighlightState = f
}

// SetCollectionViewDidDeselectItemsAtIndexPaths sets the handler for the CollectionViewDidDeselectItemsAtIndexPaths delegate method.
//
// Notifies the delegate object that one or more items were deselected.
func (d *CollectionViewDelegate) SetCollectionViewDidDeselectItemsAtIndexPaths(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer)) {
	d._CollectionViewDidDeselectItemsAtIndexPaths = f
}

// SetCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath sets the handler for the CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath delegate method.
//
// Notifies the delegate that the specified item was removed from the collection view.
func (d *CollectionViewDelegate) SetCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(f func(collectionView objc.IObject /* cross-framework: CollectionView */, item ICollectionViewItem, indexPath foundation.IndexPath)) {
	d._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath = f
}

// SetCollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath sets the handler for the CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath delegate method.
//
// Notifies the delegate that the specified supplementary view was removed from the collection view.
func (d *CollectionViewDelegate) SetCollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(f func(collectionView objc.IObject /* cross-framework: CollectionView */, view IView, elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath)) {
	d._CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath = f
}

// SetCollectionViewDidSelectItemsAtIndexPaths sets the handler for the CollectionViewDidSelectItemsAtIndexPaths delegate method.
//
// Notifies the delegate object that one or more items were selected.
func (d *CollectionViewDelegate) SetCollectionViewDidSelectItemsAtIndexPaths(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer)) {
	d._CollectionViewDidSelectItemsAtIndexPaths = f
}

// SetCollectionViewDraggingImageForItemsAtIndexesWithEventOffset sets the handler for the CollectionViewDraggingImageForItemsAtIndexesWithEventOffset delegate method.
//
// Creates and returns a drag image to represent the specified items during a drag.
func (d *CollectionViewDelegate) SetCollectionViewDraggingImageForItemsAtIndexesWithEventOffset(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, event IEvent, dragImageOffset PointPointer /* not a class type */) Image) {
	d._CollectionViewDraggingImageForItemsAtIndexesWithEventOffset = f
}

// SetCollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset sets the handler for the CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset delegate method.
//
// Creates and returns a drag image to represent the specified items during a drag.
func (d *CollectionViewDelegate) SetCollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, event IEvent, dragImageOffset PointPointer /* not a class type */) Image) {
	d._CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset = f
}

// SetCollectionViewDraggingSessionEndedAtPointDragOperation sets the handler for the CollectionViewDraggingSessionEndedAtPointDragOperation delegate method.
//
// Notifies your delegate that a drag session ended.
func (d *CollectionViewDelegate) SetCollectionViewDraggingSessionEndedAtPointDragOperation(f func(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, operation DragOperation)) {
	d._CollectionViewDraggingSessionEndedAtPointDragOperation = f
}

// SetCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths sets the handler for the CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths delegate method.
//
// Notifies your delegate that a drag session is about to begin.
func (d *CollectionViewDelegate) SetCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(f func(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, indexPaths unsafe.Pointer)) {
	d._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths = f
}

// SetCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes sets the handler for the CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes delegate method.
//
// Notifies your delegate that a drag session is about to begin.
func (d *CollectionViewDelegate) SetCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(f func(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, indexes foundation.IndexSet)) {
	d._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes = f
}

// SetCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths sets the handler for the CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths delegate method.
//
// Returns the names of the promised files that you created for a drag operation.
func (d *CollectionViewDelegate) SetCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths(f func(collectionView objc.IObject /* cross-framework: CollectionView */, dropURL objc.IObject /* cross-framework: NSURL */, indexPaths unsafe.Pointer) []string) {
	d._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths = f
}

// SetCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes sets the handler for the CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes delegate method.
//
// Invoked to return an array of filenames that the receiver promises to create.
func (d *CollectionViewDelegate) SetCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes(f func(collectionView objc.IObject /* cross-framework: CollectionView */, dropURL objc.IObject /* cross-framework: NSURL */, indexes foundation.IndexSet) []string) {
	d._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes = f
}

// SetCollectionViewPasteboardWriterForItemAtIndexPath sets the handler for the CollectionViewPasteboardWriterForItemAtIndexPath delegate method.
//
// Provides the pasteboard writer for the item at the specified index path.
func (d *CollectionViewDelegate) SetCollectionViewPasteboardWriterForItemAtIndexPath(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPath foundation.IndexPath) unsafe.Pointer) {
	d._CollectionViewPasteboardWriterForItemAtIndexPath = f
}

// SetCollectionViewPasteboardWriterForItemAtIndex sets the handler for the CollectionViewPasteboardWriterForItemAtIndex delegate method.
//
// Provides the pasteboard writer for the item at the specified index
func (d *CollectionViewDelegate) SetCollectionViewPasteboardWriterForItemAtIndex(f func(collectionView objc.IObject /* cross-framework: CollectionView */, index uint) unsafe.Pointer) {
	d._CollectionViewPasteboardWriterForItemAtIndex = f
}

// SetCollectionViewShouldChangeItemsAtIndexPathsToHighlightState sets the handler for the CollectionViewShouldChangeItemsAtIndexPathsToHighlightState delegate method.
//
// Asks the delegate to approve the pending highlighting of the specified items.
func (d *CollectionViewDelegate) SetCollectionViewShouldChangeItemsAtIndexPathsToHighlightState(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, highlightState CollectionViewItemHighlightState) unsafe.Pointer) {
	d._CollectionViewShouldChangeItemsAtIndexPathsToHighlightState = f
}

// SetCollectionViewShouldDeselectItemsAtIndexPaths sets the handler for the CollectionViewShouldDeselectItemsAtIndexPaths delegate method.
//
// Asks the delegate object to approve the pending deselection of items.
func (d *CollectionViewDelegate) SetCollectionViewShouldDeselectItemsAtIndexPaths(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) unsafe.Pointer) {
	d._CollectionViewShouldDeselectItemsAtIndexPaths = f
}

// SetCollectionViewShouldSelectItemsAtIndexPaths sets the handler for the CollectionViewShouldSelectItemsAtIndexPaths delegate method.
//
// Asks the delegate to approve the pending selection of items.
func (d *CollectionViewDelegate) SetCollectionViewShouldSelectItemsAtIndexPaths(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) unsafe.Pointer) {
	d._CollectionViewShouldSelectItemsAtIndexPaths = f
}

// SetCollectionViewTransitionLayoutForOldLayoutNewLayout sets the handler for the CollectionViewTransitionLayoutForOldLayoutNewLayout delegate method.
//
// Returns the transition layout object to use when performing an animated change between different layouts.
func (d *CollectionViewDelegate) SetCollectionViewTransitionLayoutForOldLayoutNewLayout(f func(collectionView objc.IObject /* cross-framework: CollectionView */, fromLayout ICollectionViewLayout, toLayout ICollectionViewLayout) CollectionViewTransitionLayout) {
	d._CollectionViewTransitionLayoutForOldLayoutNewLayout = f
}

// SetCollectionViewUpdateDraggingItemsForDrag sets the handler for the CollectionViewUpdateDraggingItemsForDrag delegate method.
//
// Asks your delegate to update the dragging items during a drag operation.
func (d *CollectionViewDelegate) SetCollectionViewUpdateDraggingItemsForDrag(f func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer)) {
	d._CollectionViewUpdateDraggingItemsForDrag = f
}

// SetCollectionViewValidateDropProposedIndexDropOperation sets the handler for the CollectionViewValidateDropProposedIndexDropOperation delegate method.
//
// Validates the specified location to see if it is a valid drop target.
func (d *CollectionViewDelegate) SetCollectionViewValidateDropProposedIndexDropOperation(f func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, proposedDropIndex int, proposedDropOperation CollectionViewDropOperation) DragOperation) {
	d._CollectionViewValidateDropProposedIndexDropOperation = f
}

// SetCollectionViewValidateDropProposedIndexPathDropOperation sets the handler for the CollectionViewValidateDropProposedIndexPathDropOperation delegate method.
//
// Validates whether a drop operation is possible at the specified location.
func (d *CollectionViewDelegate) SetCollectionViewValidateDropProposedIndexPathDropOperation(f func(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, proposedDropIndexPath objectivec.IObject, proposedDropOperation CollectionViewDropOperation) DragOperation) {
	d._CollectionViewValidateDropProposedIndexPathDropOperation = f
}

// SetCollectionViewWillDisplayItemForRepresentedObjectAtIndexPath sets the handler for the CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath delegate method.
//
// Notifies the delegate that the specified item is about to be displayed by the collection view.
func (d *CollectionViewDelegate) SetCollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(f func(collectionView objc.IObject /* cross-framework: CollectionView */, item ICollectionViewItem, indexPath foundation.IndexPath)) {
	d._CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath = f
}

// SetCollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath sets the handler for the CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath delegate method.
//
// Notifies the delegate that the specified supplementary view is about to be displayed by the collection view.
func (d *CollectionViewDelegate) SetCollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(f func(collectionView objc.IObject /* cross-framework: CollectionView */, view IView, elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath)) {
	d._CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath = f
}

// SetCollectionViewWriteItemsAtIndexPathsToPasteboard sets the handler for the CollectionViewWriteItemsAtIndexPathsToPasteboard delegate method.
//
// Places the data for the drag operation on the pasteboard.
func (d *CollectionViewDelegate) SetCollectionViewWriteItemsAtIndexPathsToPasteboard(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, pasteboard IPasteboard) bool) {
	d._CollectionViewWriteItemsAtIndexPathsToPasteboard = f
}

// SetCollectionViewWriteItemsAtIndexesToPasteboard sets the handler for the CollectionViewWriteItemsAtIndexesToPasteboard delegate method.
//
// Invoked after it has been determined that a drag should begin, but before the drag has been started.
func (d *CollectionViewDelegate) SetCollectionViewWriteItemsAtIndexesToPasteboard(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, pasteboard IPasteboard) bool) {
	d._CollectionViewWriteItemsAtIndexesToPasteboard = f
}

// CollectionViewAcceptDropIndexDropOperation implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewAcceptDropIndexDropOperation(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, index int, dropOperation CollectionViewDropOperation) bool {
	if d._CollectionViewAcceptDropIndexDropOperation != nil {
		return d._CollectionViewAcceptDropIndexDropOperation(collectionView, draggingInfo, index, dropOperation)
	}
	var zero bool
	return zero
}

// HasCollectionViewAcceptDropIndexDropOperation returns true if a handler for CollectionViewAcceptDropIndexDropOperation has been set.
func (d *CollectionViewDelegate) HasCollectionViewAcceptDropIndexDropOperation() bool {
	return d._CollectionViewAcceptDropIndexDropOperation != nil
}

// CollectionViewAcceptDropIndexPathDropOperation implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewAcceptDropIndexPathDropOperation(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool {
	if d._CollectionViewAcceptDropIndexPathDropOperation != nil {
		return d._CollectionViewAcceptDropIndexPathDropOperation(collectionView, draggingInfo, indexPath, dropOperation)
	}
	var zero bool
	return zero
}

// HasCollectionViewAcceptDropIndexPathDropOperation returns true if a handler for CollectionViewAcceptDropIndexPathDropOperation has been set.
func (d *CollectionViewDelegate) HasCollectionViewAcceptDropIndexPathDropOperation() bool {
	return d._CollectionViewAcceptDropIndexPathDropOperation != nil
}

// CollectionViewCanDragItemsAtIndexesWithEvent implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewCanDragItemsAtIndexesWithEvent(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, event IEvent) bool {
	if d._CollectionViewCanDragItemsAtIndexesWithEvent != nil {
		return d._CollectionViewCanDragItemsAtIndexesWithEvent(collectionView, indexes, event)
	}
	var zero bool
	return zero
}

// HasCollectionViewCanDragItemsAtIndexesWithEvent returns true if a handler for CollectionViewCanDragItemsAtIndexesWithEvent has been set.
func (d *CollectionViewDelegate) HasCollectionViewCanDragItemsAtIndexesWithEvent() bool {
	return d._CollectionViewCanDragItemsAtIndexesWithEvent != nil
}

// CollectionViewCanDragItemsAtIndexPathsWithEvent implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewCanDragItemsAtIndexPathsWithEvent(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, event IEvent) bool {
	if d._CollectionViewCanDragItemsAtIndexPathsWithEvent != nil {
		return d._CollectionViewCanDragItemsAtIndexPathsWithEvent(collectionView, indexPaths, event)
	}
	var zero bool
	return zero
}

// HasCollectionViewCanDragItemsAtIndexPathsWithEvent returns true if a handler for CollectionViewCanDragItemsAtIndexPathsWithEvent has been set.
func (d *CollectionViewDelegate) HasCollectionViewCanDragItemsAtIndexPathsWithEvent() bool {
	return d._CollectionViewCanDragItemsAtIndexPathsWithEvent != nil
}

// CollectionViewDidChangeItemsAtIndexPathsToHighlightState implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDidChangeItemsAtIndexPathsToHighlightState(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, highlightState CollectionViewItemHighlightState) {
	if d._CollectionViewDidChangeItemsAtIndexPathsToHighlightState != nil {
		d._CollectionViewDidChangeItemsAtIndexPathsToHighlightState(collectionView, indexPaths, highlightState)
	}
}

// HasCollectionViewDidChangeItemsAtIndexPathsToHighlightState returns true if a handler for CollectionViewDidChangeItemsAtIndexPathsToHighlightState has been set.
func (d *CollectionViewDelegate) HasCollectionViewDidChangeItemsAtIndexPathsToHighlightState() bool {
	return d._CollectionViewDidChangeItemsAtIndexPathsToHighlightState != nil
}

// CollectionViewDidDeselectItemsAtIndexPaths implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDidDeselectItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) {
	if d._CollectionViewDidDeselectItemsAtIndexPaths != nil {
		d._CollectionViewDidDeselectItemsAtIndexPaths(collectionView, indexPaths)
	}
}

// HasCollectionViewDidDeselectItemsAtIndexPaths returns true if a handler for CollectionViewDidDeselectItemsAtIndexPaths has been set.
func (d *CollectionViewDelegate) HasCollectionViewDidDeselectItemsAtIndexPaths() bool {
	return d._CollectionViewDidDeselectItemsAtIndexPaths != nil
}

// CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, item ICollectionViewItem, indexPath foundation.IndexPath) {
	if d._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath != nil {
		d._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath(collectionView, item, indexPath)
	}
}

// HasCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath returns true if a handler for CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath has been set.
func (d *CollectionViewDelegate) HasCollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath() bool {
	return d._CollectionViewDidEndDisplayingItemForRepresentedObjectAtIndexPath != nil
}

// CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, view IView, elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath) {
	if d._CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath != nil {
		d._CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath(collectionView, view, elementKind, indexPath)
	}
}

// HasCollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath returns true if a handler for CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath has been set.
func (d *CollectionViewDelegate) HasCollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath() bool {
	return d._CollectionViewDidEndDisplayingSupplementaryViewForElementOfKindAtIndexPath != nil
}

// CollectionViewDidSelectItemsAtIndexPaths implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDidSelectItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) {
	if d._CollectionViewDidSelectItemsAtIndexPaths != nil {
		d._CollectionViewDidSelectItemsAtIndexPaths(collectionView, indexPaths)
	}
}

// HasCollectionViewDidSelectItemsAtIndexPaths returns true if a handler for CollectionViewDidSelectItemsAtIndexPaths has been set.
func (d *CollectionViewDelegate) HasCollectionViewDidSelectItemsAtIndexPaths() bool {
	return d._CollectionViewDidSelectItemsAtIndexPaths != nil
}

// CollectionViewDraggingImageForItemsAtIndexesWithEventOffset implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDraggingImageForItemsAtIndexesWithEventOffset(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, event IEvent, dragImageOffset PointPointer /* not a class type */) Image {
	if d._CollectionViewDraggingImageForItemsAtIndexesWithEventOffset != nil {
		return d._CollectionViewDraggingImageForItemsAtIndexesWithEventOffset(collectionView, indexes, event, dragImageOffset)
	}
	var zero Image
	return zero
}

// HasCollectionViewDraggingImageForItemsAtIndexesWithEventOffset returns true if a handler for CollectionViewDraggingImageForItemsAtIndexesWithEventOffset has been set.
func (d *CollectionViewDelegate) HasCollectionViewDraggingImageForItemsAtIndexesWithEventOffset() bool {
	return d._CollectionViewDraggingImageForItemsAtIndexesWithEventOffset != nil
}

// CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, event IEvent, dragImageOffset PointPointer /* not a class type */) Image {
	if d._CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset != nil {
		return d._CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset(collectionView, indexPaths, event, dragImageOffset)
	}
	var zero Image
	return zero
}

// HasCollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset returns true if a handler for CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset has been set.
func (d *CollectionViewDelegate) HasCollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset() bool {
	return d._CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset != nil
}

// CollectionViewDraggingSessionEndedAtPointDragOperation implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDraggingSessionEndedAtPointDragOperation(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, operation DragOperation) {
	if d._CollectionViewDraggingSessionEndedAtPointDragOperation != nil {
		d._CollectionViewDraggingSessionEndedAtPointDragOperation(collectionView, session, screenPoint, operation)
	}
}

// HasCollectionViewDraggingSessionEndedAtPointDragOperation returns true if a handler for CollectionViewDraggingSessionEndedAtPointDragOperation has been set.
func (d *CollectionViewDelegate) HasCollectionViewDraggingSessionEndedAtPointDragOperation() bool {
	return d._CollectionViewDraggingSessionEndedAtPointDragOperation != nil
}

// CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, indexPaths unsafe.Pointer) {
	if d._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths != nil {
		d._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths(collectionView, session, screenPoint, indexPaths)
	}
}

// HasCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths returns true if a handler for CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths has been set.
func (d *CollectionViewDelegate) HasCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths() bool {
	return d._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexPaths != nil
}

// CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(collectionView objc.IObject /* cross-framework: CollectionView */, session IDraggingSession, screenPoint vision.Point, indexes foundation.IndexSet) {
	if d._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes != nil {
		d._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes(collectionView, session, screenPoint, indexes)
	}
}

// HasCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes returns true if a handler for CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes has been set.
func (d *CollectionViewDelegate) HasCollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes() bool {
	return d._CollectionViewDraggingSessionWillBeginAtPointForItemsAtIndexes != nil
}

// CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, dropURL objc.IObject /* cross-framework: NSURL */, indexPaths unsafe.Pointer) []string {
	if d._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths != nil {
		return d._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths(collectionView, dropURL, indexPaths)
	}
	var zero []string
	return zero
}

// HasCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths returns true if a handler for CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths has been set.
func (d *CollectionViewDelegate) HasCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths() bool {
	return d._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexPaths != nil
}

// CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes(collectionView objc.IObject /* cross-framework: CollectionView */, dropURL objc.IObject /* cross-framework: NSURL */, indexes foundation.IndexSet) []string {
	if d._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes != nil {
		return d._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes(collectionView, dropURL, indexes)
	}
	var zero []string
	return zero
}

// HasCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes returns true if a handler for CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes has been set.
func (d *CollectionViewDelegate) HasCollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes() bool {
	return d._CollectionViewNamesOfPromisedFilesDroppedAtDestinationForDraggedItemsAtIndexes != nil
}

// CollectionViewPasteboardWriterForItemAtIndexPath implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewPasteboardWriterForItemAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, indexPath foundation.IndexPath) unsafe.Pointer {
	if d._CollectionViewPasteboardWriterForItemAtIndexPath != nil {
		return d._CollectionViewPasteboardWriterForItemAtIndexPath(collectionView, indexPath)
	}
	var zero unsafe.Pointer
	return zero
}

// HasCollectionViewPasteboardWriterForItemAtIndexPath returns true if a handler for CollectionViewPasteboardWriterForItemAtIndexPath has been set.
func (d *CollectionViewDelegate) HasCollectionViewPasteboardWriterForItemAtIndexPath() bool {
	return d._CollectionViewPasteboardWriterForItemAtIndexPath != nil
}

// CollectionViewPasteboardWriterForItemAtIndex implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewPasteboardWriterForItemAtIndex(collectionView objc.IObject /* cross-framework: CollectionView */, index uint) unsafe.Pointer {
	if d._CollectionViewPasteboardWriterForItemAtIndex != nil {
		return d._CollectionViewPasteboardWriterForItemAtIndex(collectionView, index)
	}
	var zero unsafe.Pointer
	return zero
}

// HasCollectionViewPasteboardWriterForItemAtIndex returns true if a handler for CollectionViewPasteboardWriterForItemAtIndex has been set.
func (d *CollectionViewDelegate) HasCollectionViewPasteboardWriterForItemAtIndex() bool {
	return d._CollectionViewPasteboardWriterForItemAtIndex != nil
}

// CollectionViewShouldChangeItemsAtIndexPathsToHighlightState implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewShouldChangeItemsAtIndexPathsToHighlightState(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, highlightState CollectionViewItemHighlightState) unsafe.Pointer {
	if d._CollectionViewShouldChangeItemsAtIndexPathsToHighlightState != nil {
		return d._CollectionViewShouldChangeItemsAtIndexPathsToHighlightState(collectionView, indexPaths, highlightState)
	}
	var zero unsafe.Pointer
	return zero
}

// HasCollectionViewShouldChangeItemsAtIndexPathsToHighlightState returns true if a handler for CollectionViewShouldChangeItemsAtIndexPathsToHighlightState has been set.
func (d *CollectionViewDelegate) HasCollectionViewShouldChangeItemsAtIndexPathsToHighlightState() bool {
	return d._CollectionViewShouldChangeItemsAtIndexPathsToHighlightState != nil
}

// CollectionViewShouldDeselectItemsAtIndexPaths implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewShouldDeselectItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) unsafe.Pointer {
	if d._CollectionViewShouldDeselectItemsAtIndexPaths != nil {
		return d._CollectionViewShouldDeselectItemsAtIndexPaths(collectionView, indexPaths)
	}
	var zero unsafe.Pointer
	return zero
}

// HasCollectionViewShouldDeselectItemsAtIndexPaths returns true if a handler for CollectionViewShouldDeselectItemsAtIndexPaths has been set.
func (d *CollectionViewDelegate) HasCollectionViewShouldDeselectItemsAtIndexPaths() bool {
	return d._CollectionViewShouldDeselectItemsAtIndexPaths != nil
}

// CollectionViewShouldSelectItemsAtIndexPaths implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewShouldSelectItemsAtIndexPaths(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer) unsafe.Pointer {
	if d._CollectionViewShouldSelectItemsAtIndexPaths != nil {
		return d._CollectionViewShouldSelectItemsAtIndexPaths(collectionView, indexPaths)
	}
	var zero unsafe.Pointer
	return zero
}

// HasCollectionViewShouldSelectItemsAtIndexPaths returns true if a handler for CollectionViewShouldSelectItemsAtIndexPaths has been set.
func (d *CollectionViewDelegate) HasCollectionViewShouldSelectItemsAtIndexPaths() bool {
	return d._CollectionViewShouldSelectItemsAtIndexPaths != nil
}

// CollectionViewTransitionLayoutForOldLayoutNewLayout implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewTransitionLayoutForOldLayoutNewLayout(collectionView objc.IObject /* cross-framework: CollectionView */, fromLayout ICollectionViewLayout, toLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	if d._CollectionViewTransitionLayoutForOldLayoutNewLayout != nil {
		return d._CollectionViewTransitionLayoutForOldLayoutNewLayout(collectionView, fromLayout, toLayout)
	}
	var zero CollectionViewTransitionLayout
	return zero
}

// HasCollectionViewTransitionLayoutForOldLayoutNewLayout returns true if a handler for CollectionViewTransitionLayoutForOldLayoutNewLayout has been set.
func (d *CollectionViewDelegate) HasCollectionViewTransitionLayoutForOldLayoutNewLayout() bool {
	return d._CollectionViewTransitionLayoutForOldLayoutNewLayout != nil
}

// CollectionViewUpdateDraggingItemsForDrag implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewUpdateDraggingItemsForDrag(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer) {
	if d._CollectionViewUpdateDraggingItemsForDrag != nil {
		d._CollectionViewUpdateDraggingItemsForDrag(collectionView, draggingInfo)
	}
}

// HasCollectionViewUpdateDraggingItemsForDrag returns true if a handler for CollectionViewUpdateDraggingItemsForDrag has been set.
func (d *CollectionViewDelegate) HasCollectionViewUpdateDraggingItemsForDrag() bool {
	return d._CollectionViewUpdateDraggingItemsForDrag != nil
}

// CollectionViewValidateDropProposedIndexDropOperation implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewValidateDropProposedIndexDropOperation(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, proposedDropIndex int, proposedDropOperation CollectionViewDropOperation) DragOperation {
	if d._CollectionViewValidateDropProposedIndexDropOperation != nil {
		return d._CollectionViewValidateDropProposedIndexDropOperation(collectionView, draggingInfo, proposedDropIndex, proposedDropOperation)
	}
	var zero DragOperation
	return zero
}

// HasCollectionViewValidateDropProposedIndexDropOperation returns true if a handler for CollectionViewValidateDropProposedIndexDropOperation has been set.
func (d *CollectionViewDelegate) HasCollectionViewValidateDropProposedIndexDropOperation() bool {
	return d._CollectionViewValidateDropProposedIndexDropOperation != nil
}

// CollectionViewValidateDropProposedIndexPathDropOperation implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewValidateDropProposedIndexPathDropOperation(collectionView objc.IObject /* cross-framework: CollectionView */, draggingInfo unsafe.Pointer, proposedDropIndexPath objectivec.IObject, proposedDropOperation CollectionViewDropOperation) DragOperation {
	if d._CollectionViewValidateDropProposedIndexPathDropOperation != nil {
		return d._CollectionViewValidateDropProposedIndexPathDropOperation(collectionView, draggingInfo, proposedDropIndexPath, proposedDropOperation)
	}
	var zero DragOperation
	return zero
}

// HasCollectionViewValidateDropProposedIndexPathDropOperation returns true if a handler for CollectionViewValidateDropProposedIndexPathDropOperation has been set.
func (d *CollectionViewDelegate) HasCollectionViewValidateDropProposedIndexPathDropOperation() bool {
	return d._CollectionViewValidateDropProposedIndexPathDropOperation != nil
}

// CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, item ICollectionViewItem, indexPath foundation.IndexPath) {
	if d._CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath != nil {
		d._CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath(collectionView, item, indexPath)
	}
}

// HasCollectionViewWillDisplayItemForRepresentedObjectAtIndexPath returns true if a handler for CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath has been set.
func (d *CollectionViewDelegate) HasCollectionViewWillDisplayItemForRepresentedObjectAtIndexPath() bool {
	return d._CollectionViewWillDisplayItemForRepresentedObjectAtIndexPath != nil
}

// CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, view IView, elementKind CollectionViewSupplementaryElementKind /* typedef */, indexPath foundation.IndexPath) {
	if d._CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath != nil {
		d._CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath(collectionView, view, elementKind, indexPath)
	}
}

// HasCollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath returns true if a handler for CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath has been set.
func (d *CollectionViewDelegate) HasCollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath() bool {
	return d._CollectionViewWillDisplaySupplementaryViewForElementKindAtIndexPath != nil
}

// CollectionViewWriteItemsAtIndexPathsToPasteboard implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewWriteItemsAtIndexPathsToPasteboard(collectionView objc.IObject /* cross-framework: CollectionView */, indexPaths unsafe.Pointer, pasteboard IPasteboard) bool {
	if d._CollectionViewWriteItemsAtIndexPathsToPasteboard != nil {
		return d._CollectionViewWriteItemsAtIndexPathsToPasteboard(collectionView, indexPaths, pasteboard)
	}
	var zero bool
	return zero
}

// HasCollectionViewWriteItemsAtIndexPathsToPasteboard returns true if a handler for CollectionViewWriteItemsAtIndexPathsToPasteboard has been set.
func (d *CollectionViewDelegate) HasCollectionViewWriteItemsAtIndexPathsToPasteboard() bool {
	return d._CollectionViewWriteItemsAtIndexPathsToPasteboard != nil
}

// CollectionViewWriteItemsAtIndexesToPasteboard implements the PCollectionViewDelegate interface.
func (d *CollectionViewDelegate) CollectionViewWriteItemsAtIndexesToPasteboard(collectionView objc.IObject /* cross-framework: CollectionView */, indexes foundation.IndexSet, pasteboard IPasteboard) bool {
	if d._CollectionViewWriteItemsAtIndexesToPasteboard != nil {
		return d._CollectionViewWriteItemsAtIndexesToPasteboard(collectionView, indexes, pasteboard)
	}
	var zero bool
	return zero
}

// HasCollectionViewWriteItemsAtIndexesToPasteboard returns true if a handler for CollectionViewWriteItemsAtIndexesToPasteboard has been set.
func (d *CollectionViewDelegate) HasCollectionViewWriteItemsAtIndexesToPasteboard() bool {
	return d._CollectionViewWriteItemsAtIndexesToPasteboard != nil
}
