// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CollectionView] class.
var collectionViewClass = _CollectionViewClass{objc.GetClass("NSCollectionView")}

type _CollectionViewClass struct {
	class objc.Class
}

// An ordered collection of data items displayed in a customizable layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView

type CollectionView struct {
	View
}

// CollectionViewFrom constructs a [CollectionView] from an unsafe.Pointer.
//
// An ordered collection of data items displayed in a customizable layout.
func CollectionViewFrom(ptr unsafe.Pointer) CollectionView {
	return CollectionView{
		View: ViewFrom(ptr),
	}
}

// Deletes the items at the specified index paths. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deleteItems(at:)
func (c_ CollectionView) DeleteItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteItemsAtIndexPaths:"), indexPaths)
}
// Deletes the specified sections and their contained items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deleteSections(_:)
func (c_ CollectionView) DeleteSections(sections unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteSections:"), sections)
}
// Deselects all items in the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deselectAll(_:)
func (c_ CollectionView) DeselectAll(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deselectAll:"), sender)
}
// Removes the specified items from the current selection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deselectItems(at:)
func (c_ CollectionView) DeselectItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deselectItemsAtIndexPaths:"), indexPaths)
}
// Returns an image to use for dragging the specified items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-7rc4k
func (c_ CollectionView) DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("draggingImageForItemsAtIndexPaths:withEvent:offset:"), indexPaths, event, dragImageOffset)
	return rv
}
// This method computes and returns an image to use for dragging. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-951w7
func (c_ CollectionView) DraggingImageForItemsAtIndexesWithEventOffset(indexes unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("draggingImageForItemsAtIndexes:withEvent:offset:"), indexes, event, dragImageOffset)
	return rv
}
// Returns the frame of the collection view item at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/frameForItem(at:)
func (c_ CollectionView) FrameForItemAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("frameForItemAtIndex:"), index)
	return rv
}
// Returns the frame of an item based on the number of items in the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/frameForItem(at:withNumberOfItems:)
func (c_ CollectionView) FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("frameForItemAtIndex:withNumberOfItems:"), index, numberOfItems)
	return rv
}
// Returns the index path of the specified item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPath(for:)
func (c_ CollectionView) IndexPathForItem(item unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathForItem:"), item)
	return rv
}
// Returns the index path of the item at the specified point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathForItem(at:)
func (c_ CollectionView) IndexPathForItemAtPoint(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathForItemAtPoint:"), point)
	return rv
}
// Returns the index paths of the currently active items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathsForVisibleItems()
func (c_ CollectionView) IndexPathsForVisibleItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsForVisibleItems"))
	return rv
}
// Returns the index paths of the currently active supplementary views. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathsForVisibleSupplementaryElements(ofKind:)
func (c_ CollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsForVisibleSupplementaryElementsOfKind:"), elementKind)
	return rv
}
// Inserts new items into the collection view at the specified locations. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/insertItems(at:)
func (c_ CollectionView) InsertItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertItemsAtIndexPaths:"), indexPaths)
}
// Inserts new sections at the specified indexes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/insertSections(_:)
func (c_ CollectionView) InsertSections(sections unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertSections:"), sections)
}
// Returns the item associated with the specified index path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/item(at:)-2vx2h
func (c_ CollectionView) ItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("itemAtIndexPath:"), indexPath)
	return rv
}
// Returns the collection view item for the represented object at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/item(at:)-80xze
func (c_ CollectionView) ItemAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}
// Returns the layout information for the item at the specified index path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/layoutAttributesForItem(at:)
func (c_ CollectionView) LayoutAttributesForItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("layoutAttributesForItemAtIndexPath:"), indexPath)
	return rv
}
// Returns the layout information for the supplementary view at the specified index path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/layoutAttributesForSupplementaryElement(ofKind:at:)
func (c_ CollectionView) LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("layoutAttributesForSupplementaryElementOfKind:atIndexPath:"), kind, indexPath)
	return rv
}
// Creates or returns a reusable item object of the specified type. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/makeItem(withIdentifier:for:)
func (c_ CollectionView) MakeItemWithIdentifierForIndexPath(identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("makeItemWithIdentifier:forIndexPath:"), identifier, indexPath)
	return rv
}
// Creates or returns a reusable supplementary view of the specified type. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/makeSupplementaryView(ofKind:withIdentifier:for:)
func (c_ CollectionView) MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind unsafe.Pointer, identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("makeSupplementaryViewOfKind:withIdentifier:forIndexPath:"), elementKind, identifier, indexPath)
	return rv
}
// Moves an item from one location to another in the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/moveItem(at:to:)
func (c_ CollectionView) MoveItemAtIndexPathToIndexPath(indexPath unsafe.Pointer, newIndexPath unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("moveItemAtIndexPath:toIndexPath:"), indexPath, newIndexPath)
}
// Moves a section from its current location to a new location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/moveSection(_:toSection:)
func (c_ CollectionView) MoveSectionToSection(section int, newSection int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("moveSection:toSection:"), section, newSection)
}
// Returns the collection view item that is used for the specified object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/newItem(forRepresentedObject:)
func (c_ CollectionView) NewItemForRepresentedObject(object objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("newItemForRepresentedObject:"), object)
	return rv
}
// Returns the number of items in the specified section. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/numberOfItems(inSection:)
func (c_ CollectionView) NumberOfItemsInSection(section int) int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfItemsInSection:"), section)
	return rv
}
// Encapsulates multiple insert, delete, reload, and move operations into a single animated operation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/performBatchUpdates(_:completionHandler:)
func (c_ CollectionView) PerformBatchUpdatesCompletionHandler(updates unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performBatchUpdates:completionHandler:"), updates, completionHandler)
}
// Registers a class to use when creating new items in the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-6s4i
func (c_ CollectionView) RegisterClassForItemWithIdentifier(itemClass objc.Class, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerClass:forItemWithIdentifier:"), itemClass, identifier)
}
// Registers a nib file to use when creating items in the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-90h1i
func (c_ CollectionView) RegisterNibForItemWithIdentifier(nib unsafe.Pointer, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerNib:forItemWithIdentifier:"), nib, identifier)
}
// Registers a class to use when creating new supplementary views in the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-3dqa
func (c_ CollectionView) RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind unsafe.Pointer, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerClass:forSupplementaryViewOfKind:withIdentifier:"), viewClass, kind, identifier)
}
// Registers a nib file to use when creating supplementary views in the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-7gvf2
func (c_ CollectionView) RegisterNibForSupplementaryViewOfKindWithIdentifier(nib unsafe.Pointer, kind unsafe.Pointer, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerNib:forSupplementaryViewOfKind:withIdentifier:"), nib, kind, identifier)
}
// Reloads all of the data for the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadData()
func (c_ CollectionView) ReloadData() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadData"))
}
// Reloads only the specified items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadItems(at:)
func (c_ CollectionView) ReloadItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadItemsAtIndexPaths:"), indexPaths)
}
// Reloads the data in the specified sections of the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadSections(_:)
func (c_ CollectionView) ReloadSections(sections unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadSections:"), sections)
}
// Scrolls the collection view contents until the specified items are visible. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/scrollToItems(at:scrollPosition:)
func (c_ CollectionView) ScrollToItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollToItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}
// Selects all items in the collection view, if doing so is possible. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectAll(_:)
func (c_ CollectionView) SelectAll(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectAll:"), sender)
}
// Adds the specified items to the current selection and optionally scrolls the items into position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectItems(at:scrollPosition:)
func (c_ CollectionView) SelectItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}
// Configures the drag operation mask. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/setDraggingSourceOperationMask(_:forLocal:)
func (c_ CollectionView) SetDraggingSourceOperationMaskForLocal(dragOperationMask unsafe.Pointer, localDestination bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), dragOperationMask, localDestination)
}
// Returns the supplementary view associated with the specified index path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/supplementaryView(forElementKind:at:)
func (c_ CollectionView) SupplementaryViewForElementKindAtIndexPath(elementKind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supplementaryViewForElementKind:atIndexPath:"), elementKind, indexPath)
	return rv
}
// Collapses the section in which the sender resides into a single horizontally scrollable row. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/toggleSectionCollapse(_:)
func (c_ CollectionView) ToggleSectionCollapse(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("toggleSectionCollapse:"), sender)
}
// Returns an array of the actively managed items in the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/visibleItems()
func (c_ CollectionView) VisibleItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("visibleItems"))
	return rv
}
// Returns an array of the actively managed supplementary views in the collection view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/visibleSupplementaryViews(ofKind:)
func (c_ CollectionView) VisibleSupplementaryViewsOfKind(elementKind unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("visibleSupplementaryViewsOfKind:"), elementKind)
	return rv
}


