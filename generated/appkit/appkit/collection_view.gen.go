// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionView] class.
var CollectionViewClass objc.Class

func init() {
	CollectionViewClass = objc.GetClass("NSCollectionView")
}

type CollectionView struct {
	objc.ID
}

func CollectionViewFrom(ptr unsafe.Pointer) CollectionView {
	return CollectionView{
		ID: objc.ID(ptr),
	}
}


// Deletes the items at the specified index paths. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/deleteItems(at:)
func (c_ CollectionView) DeleteItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	sel := objc.RegisterName("deleteItemsAtIndexPaths:")
	c_.ID.Send(sel, indexPaths)
}
// Deletes the specified sections and their contained items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/deleteSections(_:)
func (c_ CollectionView) DeleteSections(sections unsafe.Pointer) {
	sel := objc.RegisterName("deleteSections:")
	c_.ID.Send(sel, sections)
}
// Deselects all items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/deselectAll(_:)
func (c_ CollectionView) DeselectAll(sender objc.ID) {
	sel := objc.RegisterName("deselectAll:")
	c_.ID.Send(sel, sender)
}
// Removes the specified items from the current selection. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/deselectItems(at:)
func (c_ CollectionView) DeselectItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	sel := objc.RegisterName("deselectItemsAtIndexPaths:")
	c_.ID.Send(sel, indexPaths)
}
// Returns an image to use for dragging the specified items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-7rc4k
func (c_ CollectionView) DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("draggingImageForItemsAtIndexPaths:withEvent:offset:")
	ret := c_.ID.Send(sel, indexPaths, event, dragImageOffset)
	return unsafe.Pointer(ret)
}
// This method computes and returns an image to use for dragging. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-951w7
func (c_ CollectionView) DraggingImageForItemsAtIndexesWithEventOffset(indexes unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("draggingImageForItemsAtIndexes:withEvent:offset:")
	ret := c_.ID.Send(sel, indexes, event, dragImageOffset)
	return unsafe.Pointer(ret)
}
// Returns the frame of the collection view item at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/frameForItem(at:)
func (c_ CollectionView) FrameForItemAtIndex(index uint) unsafe.Pointer {
	sel := objc.RegisterName("frameForItemAtIndex:")
	ret := c_.ID.Send(sel, index)
	return unsafe.Pointer(ret)
}
// Returns the frame of an item based on the number of items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/frameForItem(at:withNumberOfItems:)
func (c_ CollectionView) FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) unsafe.Pointer {
	sel := objc.RegisterName("frameForItemAtIndex:withNumberOfItems:")
	ret := c_.ID.Send(sel, index, numberOfItems)
	return unsafe.Pointer(ret)
}
// Returns the index path of the specified item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/indexPath(for:)
func (c_ CollectionView) IndexPathForItem(item unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("indexPathForItem:")
	ret := c_.ID.Send(sel, item)
	return unsafe.Pointer(ret)
}
// Returns the index path of the item at the specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/indexPathForItem(at:)
func (c_ CollectionView) IndexPathForItemAtPoint(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("indexPathForItemAtPoint:")
	ret := c_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Returns the index paths of the currently active items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/indexPathsForVisibleItems()
func (c_ CollectionView) IndexPathsForVisibleItems() unsafe.Pointer {
	sel := objc.RegisterName("indexPathsForVisibleItems")
	ret := c_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns the index paths of the currently active supplementary views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/indexPathsForVisibleSupplementaryElements(ofKind:)
func (c_ CollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("indexPathsForVisibleSupplementaryElementsOfKind:")
	ret := c_.ID.Send(sel, elementKind)
	return unsafe.Pointer(ret)
}
// Inserts new items into the collection view at the specified locations. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/insertItems(at:)
func (c_ CollectionView) InsertItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	sel := objc.RegisterName("insertItemsAtIndexPaths:")
	c_.ID.Send(sel, indexPaths)
}
// Inserts new sections at the specified indexes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/insertSections(_:)
func (c_ CollectionView) InsertSections(sections unsafe.Pointer) {
	sel := objc.RegisterName("insertSections:")
	c_.ID.Send(sel, sections)
}
// Returns the item associated with the specified index path. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/item(at:)-2vx2h
func (c_ CollectionView) ItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("itemAtIndexPath:")
	ret := c_.ID.Send(sel, indexPath)
	return unsafe.Pointer(ret)
}
// Returns the collection view item for the represented object at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/item(at:)-80xze
func (c_ CollectionView) ItemAtIndex(index uint) unsafe.Pointer {
	sel := objc.RegisterName("itemAtIndex:")
	ret := c_.ID.Send(sel, index)
	return unsafe.Pointer(ret)
}
// Returns the layout information for the item at the specified index path. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/layoutAttributesForItem(at:)
func (c_ CollectionView) LayoutAttributesForItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("layoutAttributesForItemAtIndexPath:")
	ret := c_.ID.Send(sel, indexPath)
	return unsafe.Pointer(ret)
}
// Returns the layout information for the supplementary view at the specified index path. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/layoutAttributesForSupplementaryElement(ofKind:at:)
func (c_ CollectionView) LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("layoutAttributesForSupplementaryElementOfKind:atIndexPath:")
	ret := c_.ID.Send(sel, kind, indexPath)
	return unsafe.Pointer(ret)
}
// Creates or returns a reusable item object of the specified type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/makeItem(withIdentifier:for:)
func (c_ CollectionView) MakeItemWithIdentifierForIndexPath(identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("makeItemWithIdentifier:forIndexPath:")
	ret := c_.ID.Send(sel, identifier, indexPath)
	return unsafe.Pointer(ret)
}
// Creates or returns a reusable supplementary view of the specified type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/makeSupplementaryView(ofKind:withIdentifier:for:)
func (c_ CollectionView) MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind unsafe.Pointer, identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("makeSupplementaryViewOfKind:withIdentifier:forIndexPath:")
	ret := c_.ID.Send(sel, elementKind, identifier, indexPath)
	return unsafe.Pointer(ret)
}
// Moves an item from one location to another in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/moveItem(at:to:)
func (c_ CollectionView) MoveItemAtIndexPathToIndexPath(indexPath unsafe.Pointer, newIndexPath unsafe.Pointer) {
	sel := objc.RegisterName("moveItemAtIndexPath:toIndexPath:")
	c_.ID.Send(sel, indexPath, newIndexPath)
}
// Moves a section from its current location to a new location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/moveSection(_:toSection:)
func (c_ CollectionView) MoveSectionToSection(section int, newSection int) {
	sel := objc.RegisterName("moveSection:toSection:")
	c_.ID.Send(sel, section, newSection)
}
// Returns the collection view item that is used for the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/newItem(forRepresentedObject:)
func (c_ CollectionView) NewItemForRepresentedObject(object objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("newItemForRepresentedObject:")
	ret := c_.ID.Send(sel, object)
	return unsafe.Pointer(ret)
}
// Returns the number of items in the specified section. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/numberOfItems(inSection:)
func (c_ CollectionView) NumberOfItemsInSection(section int) int {
	sel := objc.RegisterName("numberOfItemsInSection:")
	ret := c_.ID.Send(sel, section)
	return int(ret)
}
// Encapsulates multiple insert, delete, reload, and move operations into a single animated operation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/performBatchUpdates(_:completionHandler:)
func (c_ CollectionView) PerformBatchUpdatesCompletionHandler(updates unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("performBatchUpdates:completionHandler:")
	c_.ID.Send(sel, updates, completionHandler)
}
// Registers a class to use when creating new items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-6s4i
func (c_ CollectionView) RegisterClassForItemWithIdentifier(itemClass objc.Class, identifier unsafe.Pointer) {
	sel := objc.RegisterName("registerClass:forItemWithIdentifier:")
	c_.ID.Send(sel, itemClass, identifier)
}
// Registers a nib file to use when creating items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-90h1i
func (c_ CollectionView) RegisterNibForItemWithIdentifier(nib unsafe.Pointer, identifier unsafe.Pointer) {
	sel := objc.RegisterName("registerNib:forItemWithIdentifier:")
	c_.ID.Send(sel, nib, identifier)
}
// Registers a class to use when creating new supplementary views in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-3dqa
func (c_ CollectionView) RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind unsafe.Pointer, identifier unsafe.Pointer) {
	sel := objc.RegisterName("registerClass:forSupplementaryViewOfKind:withIdentifier:")
	c_.ID.Send(sel, viewClass, kind, identifier)
}
// Registers a nib file to use when creating supplementary views in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-7gvf2
func (c_ CollectionView) RegisterNibForSupplementaryViewOfKindWithIdentifier(nib unsafe.Pointer, kind unsafe.Pointer, identifier unsafe.Pointer) {
	sel := objc.RegisterName("registerNib:forSupplementaryViewOfKind:withIdentifier:")
	c_.ID.Send(sel, nib, kind, identifier)
}
// Reloads all of the data for the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/reloadData()
func (c_ CollectionView) ReloadData() {
	sel := objc.RegisterName("reloadData")
	c_.ID.Send(sel)
}
// Reloads only the specified items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/reloadItems(at:)
func (c_ CollectionView) ReloadItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	sel := objc.RegisterName("reloadItemsAtIndexPaths:")
	c_.ID.Send(sel, indexPaths)
}
// Reloads the data in the specified sections of the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/reloadSections(_:)
func (c_ CollectionView) ReloadSections(sections unsafe.Pointer) {
	sel := objc.RegisterName("reloadSections:")
	c_.ID.Send(sel, sections)
}
// Scrolls the collection view contents until the specified items are visible. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/scrollToItems(at:scrollPosition:)
func (c_ CollectionView) ScrollToItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer) {
	sel := objc.RegisterName("scrollToItemsAtIndexPaths:scrollPosition:")
	c_.ID.Send(sel, indexPaths, scrollPosition)
}
// Selects all items in the collection view, if doing so is possible. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/selectAll(_:)
func (c_ CollectionView) SelectAll(sender objc.ID) {
	sel := objc.RegisterName("selectAll:")
	c_.ID.Send(sel, sender)
}
// Adds the specified items to the current selection and optionally scrolls the items into position. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/selectItems(at:scrollPosition:)
func (c_ CollectionView) SelectItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer) {
	sel := objc.RegisterName("selectItemsAtIndexPaths:scrollPosition:")
	c_.ID.Send(sel, indexPaths, scrollPosition)
}
// Configures the drag operation mask. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/setDraggingSourceOperationMask(_:forLocal:)
func (c_ CollectionView) SetDraggingSourceOperationMaskForLocal(dragOperationMask unsafe.Pointer, localDestination bool) {
	sel := objc.RegisterName("setDraggingSourceOperationMask:forLocal:")
	c_.ID.Send(sel, dragOperationMask, localDestination)
}
// Returns the supplementary view associated with the specified index path. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/supplementaryView(forElementKind:at:)
func (c_ CollectionView) SupplementaryViewForElementKindAtIndexPath(elementKind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("supplementaryViewForElementKind:atIndexPath:")
	ret := c_.ID.Send(sel, elementKind, indexPath)
	return unsafe.Pointer(ret)
}
// Collapses the section in which the sender resides into a single horizontally scrollable row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/toggleSectionCollapse(_:)
func (c_ CollectionView) ToggleSectionCollapse(sender objc.ID) {
	sel := objc.RegisterName("toggleSectionCollapse:")
	c_.ID.Send(sel, sender)
}
// Returns an array of the actively managed items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/visibleItems()
func (c_ CollectionView) VisibleItems() unsafe.Pointer {
	sel := objc.RegisterName("visibleItems")
	ret := c_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Returns an array of the actively managed supplementary views in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/visibleSupplementaryViews(ofKind:)
func (c_ CollectionView) VisibleSupplementaryViewsOfKind(elementKind unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("visibleSupplementaryViewsOfKind:")
	ret := c_.ID.Send(sel, elementKind)
	return unsafe.Pointer(ret)
}


