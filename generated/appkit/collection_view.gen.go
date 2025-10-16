
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CollectionView] class.
var CollectionViewClass _CollectionViewClass

func init() {
	CollectionViewClass = _CollectionViewClass{objc.GetClass("NSCollectionView")}
}

type _CollectionViewClass struct {
	objc.Class
}

// An interface definition for the [CollectionView] class.
type ICollectionView interface {
	ID() objc.ID
	DeleteItemsAtIndexPaths(indexPaths unsafe.Pointer)
	DeleteSections(sections unsafe.Pointer)
	DeselectAll(sender objc.ID)
	DeselectItemsAtIndexPaths(indexPaths unsafe.Pointer)
	DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer
	DraggingImageForItemsAtIndexesWithEventOffset(indexes unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer
	FrameForItemAtIndex(index uint) unsafe.Pointer
	FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) unsafe.Pointer
	IndexPathForItem(item unsafe.Pointer) unsafe.Pointer
	IndexPathForItemAtPoint(point unsafe.Pointer) unsafe.Pointer
	IndexPathsForVisibleItems() unsafe.Pointer
	IndexPathsForVisibleSupplementaryElementsOfKind(elementKind unsafe.Pointer) unsafe.Pointer
	InsertItemsAtIndexPaths(indexPaths unsafe.Pointer)
	InsertSections(sections unsafe.Pointer)
	ItemAtIndex(index uint) unsafe.Pointer
	ItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer
	LayoutAttributesForItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer
	LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer
	MakeItemWithIdentifierForIndexPath(identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer
	MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind unsafe.Pointer, identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer
	MoveItemAtIndexPathToIndexPath(indexPath unsafe.Pointer, newIndexPath unsafe.Pointer)
	MoveSectionToSection(section int, newSection int)
	NewItemForRepresentedObject(object objc.ID) unsafe.Pointer
	NumberOfItemsInSection(section int) int
	PerformBatchUpdatesCompletionHandler(updates unsafe.Pointer, completionHandler unsafe.Pointer)
	RegisterClassForItemWithIdentifier(itemClass objc.Class, identifier unsafe.Pointer)
	RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind unsafe.Pointer, identifier unsafe.Pointer)
	RegisterNibForItemWithIdentifier(nib unsafe.Pointer, identifier unsafe.Pointer)
	RegisterNibForSupplementaryViewOfKindWithIdentifier(nib unsafe.Pointer, kind unsafe.Pointer, identifier unsafe.Pointer)
	ReloadData()
	ReloadItemsAtIndexPaths(indexPaths unsafe.Pointer)
	ReloadSections(sections unsafe.Pointer)
	ScrollToItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer)
	SelectAll(sender objc.ID)
	SelectItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer)
	SetDraggingSourceOperationMaskForLocal(dragOperationMask unsafe.Pointer, localDestination bool)
	SupplementaryViewForElementKindAtIndexPath(elementKind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer
	ToggleSectionCollapse(sender objc.ID)
	VisibleItems() unsafe.Pointer
	VisibleSupplementaryViewsOfKind(elementKind unsafe.Pointer) unsafe.Pointer
}

type CollectionView struct {
	id objc.ID
}

func CollectionViewFrom(ptr unsafe.Pointer) CollectionView {
	return CollectionView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ CollectionView) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewClass) Alloc() CollectionView {
	rv := objc.Send[CollectionView](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _CollectionViewClass) New() CollectionView {
	rv := objc.Send[CollectionView](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCollectionView creates and returns a new initialized instance.
func NewCollectionView() CollectionView {
	return CollectionViewClass.New()
}

// Init initializes the instance.
func (c_ CollectionView) Init() CollectionView {
	rv := objc.Send[CollectionView](c_.ID(), selInit)
	return rv
}
// Deletes the items at the specified index paths. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/deleteItems(at:)
func (c_ CollectionView) DeleteItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("deleteItemsAtIndexPaths:"), indexPaths)
}
// Deletes the specified sections and their contained items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/deleteSections(_:)
func (c_ CollectionView) DeleteSections(sections unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("deleteSections:"), sections)
}
// Deselects all items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/deselectAll(_:)
func (c_ CollectionView) DeselectAll(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("deselectAll:"), sender)
}
// Removes the specified items from the current selection. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/deselectItems(at:)
func (c_ CollectionView) DeselectItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("deselectItemsAtIndexPaths:"), indexPaths)
}
// Returns an image to use for dragging the specified items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-7rc4k
func (c_ CollectionView) DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("draggingImageForItemsAtIndexPaths:withEvent:offset:"), indexPaths, event, dragImageOffset)
	return rv
}
// This method computes and returns an image to use for dragging. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-951w7
func (c_ CollectionView) DraggingImageForItemsAtIndexesWithEventOffset(indexes unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("draggingImageForItemsAtIndexes:withEvent:offset:"), indexes, event, dragImageOffset)
	return rv
}
// Returns the frame of the collection view item at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/frameForItem(at:)
func (c_ CollectionView) FrameForItemAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("frameForItemAtIndex:"), index)
	return rv
}
// Returns the frame of an item based on the number of items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/frameForItem(at:withNumberOfItems:)
func (c_ CollectionView) FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("frameForItemAtIndex:withNumberOfItems:"), index, numberOfItems)
	return rv
}
// Returns the index path of the specified item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/indexPath(for:)
func (c_ CollectionView) IndexPathForItem(item unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("indexPathForItem:"), item)
	return rv
}
// Returns the index path of the item at the specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/indexPathForItem(at:)
func (c_ CollectionView) IndexPathForItemAtPoint(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("indexPathForItemAtPoint:"), point)
	return rv
}
// Returns the index paths of the currently active items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/indexPathsForVisibleItems()
func (c_ CollectionView) IndexPathsForVisibleItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("indexPathsForVisibleItems"))
	return rv
}
// Returns the index paths of the currently active supplementary views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/indexPathsForVisibleSupplementaryElements(ofKind:)
func (c_ CollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("indexPathsForVisibleSupplementaryElementsOfKind:"), elementKind)
	return rv
}
// Inserts new items into the collection view at the specified locations. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/insertItems(at:)
func (c_ CollectionView) InsertItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("insertItemsAtIndexPaths:"), indexPaths)
}
// Inserts new sections at the specified indexes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/insertSections(_:)
func (c_ CollectionView) InsertSections(sections unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("insertSections:"), sections)
}
// Returns the item associated with the specified index path. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/item(at:)-2vx2h
func (c_ CollectionView) ItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("itemAtIndexPath:"), indexPath)
	return rv
}
// Returns the collection view item for the represented object at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/item(at:)-80xze
func (c_ CollectionView) ItemAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("itemAtIndex:"), index)
	return rv
}
// Returns the layout information for the item at the specified index path. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/layoutAttributesForItem(at:)
func (c_ CollectionView) LayoutAttributesForItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("layoutAttributesForItemAtIndexPath:"), indexPath)
	return rv
}
// Returns the layout information for the supplementary view at the specified index path. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/layoutAttributesForSupplementaryElement(ofKind:at:)
func (c_ CollectionView) LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("layoutAttributesForSupplementaryElementOfKind:atIndexPath:"), kind, indexPath)
	return rv
}
// Creates or returns a reusable item object of the specified type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/makeItem(withIdentifier:for:)
func (c_ CollectionView) MakeItemWithIdentifierForIndexPath(identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("makeItemWithIdentifier:forIndexPath:"), identifier, indexPath)
	return rv
}
// Creates or returns a reusable supplementary view of the specified type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/makeSupplementaryView(ofKind:withIdentifier:for:)
func (c_ CollectionView) MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind unsafe.Pointer, identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("makeSupplementaryViewOfKind:withIdentifier:forIndexPath:"), elementKind, identifier, indexPath)
	return rv
}
// Moves an item from one location to another in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/moveItem(at:to:)
func (c_ CollectionView) MoveItemAtIndexPathToIndexPath(indexPath unsafe.Pointer, newIndexPath unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("moveItemAtIndexPath:toIndexPath:"), indexPath, newIndexPath)
}
// Moves a section from its current location to a new location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/moveSection(_:toSection:)
func (c_ CollectionView) MoveSectionToSection(section int, newSection int) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("moveSection:toSection:"), section, newSection)
}
// Returns the collection view item that is used for the specified object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/newItem(forRepresentedObject:)
func (c_ CollectionView) NewItemForRepresentedObject(object objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("newItemForRepresentedObject:"), object)
	return rv
}
// Returns the number of items in the specified section. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/numberOfItems(inSection:)
func (c_ CollectionView) NumberOfItemsInSection(section int) int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("numberOfItemsInSection:"), section)
	return rv
}
// Encapsulates multiple insert, delete, reload, and move operations into a single animated operation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/performBatchUpdates(_:completionHandler:)
func (c_ CollectionView) PerformBatchUpdatesCompletionHandler(updates unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("performBatchUpdates:completionHandler:"), updates, completionHandler)
}
// Registers a class to use when creating new items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-6s4i
func (c_ CollectionView) RegisterClassForItemWithIdentifier(itemClass objc.Class, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("registerClass:forItemWithIdentifier:"), itemClass, identifier)
}
// Registers a nib file to use when creating items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-90h1i
func (c_ CollectionView) RegisterNibForItemWithIdentifier(nib unsafe.Pointer, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("registerNib:forItemWithIdentifier:"), nib, identifier)
}
// Registers a class to use when creating new supplementary views in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-3dqa
func (c_ CollectionView) RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind unsafe.Pointer, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("registerClass:forSupplementaryViewOfKind:withIdentifier:"), viewClass, kind, identifier)
}
// Registers a nib file to use when creating supplementary views in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-7gvf2
func (c_ CollectionView) RegisterNibForSupplementaryViewOfKindWithIdentifier(nib unsafe.Pointer, kind unsafe.Pointer, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("registerNib:forSupplementaryViewOfKind:withIdentifier:"), nib, kind, identifier)
}
// Reloads all of the data for the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/reloadData()
func (c_ CollectionView) ReloadData() {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("reloadData"))
}
// Reloads only the specified items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/reloadItems(at:)
func (c_ CollectionView) ReloadItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("reloadItemsAtIndexPaths:"), indexPaths)
}
// Reloads the data in the specified sections of the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/reloadSections(_:)
func (c_ CollectionView) ReloadSections(sections unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("reloadSections:"), sections)
}
// Scrolls the collection view contents until the specified items are visible. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/scrollToItems(at:scrollPosition:)
func (c_ CollectionView) ScrollToItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("scrollToItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}
// Selects all items in the collection view, if doing so is possible. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/selectAll(_:)
func (c_ CollectionView) SelectAll(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("selectAll:"), sender)
}
// Adds the specified items to the current selection and optionally scrolls the items into position. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/selectItems(at:scrollPosition:)
func (c_ CollectionView) SelectItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("selectItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}
// Configures the drag operation mask. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/setDraggingSourceOperationMask(_:forLocal:)
func (c_ CollectionView) SetDraggingSourceOperationMaskForLocal(dragOperationMask unsafe.Pointer, localDestination bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setDraggingSourceOperationMask:forLocal:"), dragOperationMask, localDestination)
}
// Returns the supplementary view associated with the specified index path. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/supplementaryView(forElementKind:at:)
func (c_ CollectionView) SupplementaryViewForElementKindAtIndexPath(elementKind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("supplementaryViewForElementKind:atIndexPath:"), elementKind, indexPath)
	return rv
}
// Collapses the section in which the sender resides into a single horizontally scrollable row. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/toggleSectionCollapse(_:)
func (c_ CollectionView) ToggleSectionCollapse(sender objc.ID) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("toggleSectionCollapse:"), sender)
}
// Returns an array of the actively managed items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/visibleItems()
func (c_ CollectionView) VisibleItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("visibleItems"))
	return rv
}
// Returns an array of the actively managed supplementary views in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/visibleSupplementaryViews(ofKind:)
func (c_ CollectionView) VisibleSupplementaryViewsOfKind(elementKind unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("visibleSupplementaryViewsOfKind:"), elementKind)
	return rv
}
// A Boolean value indicating whether the collection view may have no selected items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/allowsEmptySelection
func (c_ CollectionView) AllowsEmptySelection() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("allowsEmptySelection"))
	return rv
}
// SetAllowsEmptySelection sets the value of the allowsEmptySelection property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/allowsEmptySelection
func (c_ CollectionView) SetAllowsEmptySelection(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAllowsEmptySelection:"), value)
}
// A Boolean value that indicates whether the user may select more than one item in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/allowsMultipleSelection
func (c_ CollectionView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("allowsMultipleSelection"))
	return rv
}
// SetAllowsMultipleSelection sets the value of the allowsMultipleSelection property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/allowsMultipleSelection
func (c_ CollectionView) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setAllowsMultipleSelection:"), value)
}
// An array containing the collection view’s background colors. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/backgroundColors
func (c_ CollectionView) BackgroundColors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("backgroundColors"))
	return rv
}
// SetBackgroundColors sets the value of the backgroundColors property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/backgroundColors
func (c_ CollectionView) SetBackgroundColors(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBackgroundColors:"), value)
}
// The background view placed behind all items and supplementary views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/backgroundView
func (c_ CollectionView) BackgroundView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("backgroundView"))
	return rv
}
// SetBackgroundView sets the value of the backgroundView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/backgroundView
func (c_ CollectionView) SetBackgroundView(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBackgroundView:"), value)
}
// A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/backgroundViewScrollsWithContent
func (c_ CollectionView) BackgroundViewScrollsWithContent() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("backgroundViewScrollsWithContent"))
	return rv
}
// SetBackgroundViewScrollsWithContent sets the value of the backgroundViewScrollsWithContent property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/backgroundViewScrollsWithContent
func (c_ CollectionView) SetBackgroundViewScrollsWithContent(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setBackgroundViewScrollsWithContent:"), value)
}
// The layout object used to organize the collection view’s content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/collectionViewLayout
func (c_ CollectionView) CollectionViewLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("collectionViewLayout"))
	return rv
}
// SetCollectionViewLayout sets the value of the collectionViewLayout property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/collectionViewLayout
func (c_ CollectionView) SetCollectionViewLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setCollectionViewLayout:"), value)
}
// An array that provides data for the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/content
func (c_ CollectionView) Content() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("content"))
	return rv
}
// SetContent sets the value of the content property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/content
func (c_ CollectionView) SetContent(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setContent:"), value)
}
// An object that provides data for the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/dataSource
func (c_ CollectionView) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("dataSource"))
	return rv
}
// SetDataSource sets the value of the dataSource property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/dataSource
func (c_ CollectionView) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setDataSource:"), value)
}
// The collection view’s delegate object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/delegate
func (c_ CollectionView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/delegate
func (c_ CollectionView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setDelegate:"), value)
}
// A Boolean value indicating whether the collection view is the first responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/isFirstResponder
func (c_ CollectionView) FirstResponder() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("firstResponder"))
	return rv
}
// A Boolean value that indicates whether the user may select items in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/isSelectable
func (c_ CollectionView) Selectable() bool {
	rv := objc.Send[bool](c_.ID(), objc.RegisterName("selectable"))
	return rv
}
// SetSelectable sets the value of the selectable property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/isSelectable
func (c_ CollectionView) SetSelectable(value bool) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setSelectable:"), value)
}
// The receiver’s collection view item prototype. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/itemPrototype
func (c_ CollectionView) ItemPrototype() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("itemPrototype"))
	return rv
}
// SetItemPrototype sets the value of the itemPrototype property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/itemPrototype
func (c_ CollectionView) SetItemPrototype(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setItemPrototype:"), value)
}
// The maximum size (in points) of items in the collection view grid. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/maxItemSize
func (c_ CollectionView) MaxItemSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("maxItemSize"))
	return rv
}
// SetMaxItemSize sets the value of the maxItemSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/maxItemSize
func (c_ CollectionView) SetMaxItemSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setMaxItemSize:"), value)
}
// The maximum number of columns that the collection view displays. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/maxNumberOfColumns
func (c_ CollectionView) MaxNumberOfColumns() uint {
	rv := objc.Send[uint](c_.ID(), objc.RegisterName("maxNumberOfColumns"))
	return rv
}
// SetMaxNumberOfColumns sets the value of the maxNumberOfColumns property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/maxNumberOfColumns
func (c_ CollectionView) SetMaxNumberOfColumns(value uint) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setMaxNumberOfColumns:"), value)
}
// The maximum number of rows that the collection view displays. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/maxNumberOfRows
func (c_ CollectionView) MaxNumberOfRows() uint {
	rv := objc.Send[uint](c_.ID(), objc.RegisterName("maxNumberOfRows"))
	return rv
}
// SetMaxNumberOfRows sets the value of the maxNumberOfRows property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/maxNumberOfRows
func (c_ CollectionView) SetMaxNumberOfRows(value uint) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setMaxNumberOfRows:"), value)
}
// The minimum size (in points) of items in the collection view grid. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/minItemSize
func (c_ CollectionView) MinItemSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("minItemSize"))
	return rv
}
// SetMinItemSize sets the value of the minItemSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/minItemSize
func (c_ CollectionView) SetMinItemSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setMinItemSize:"), value)
}
// The number of sections in the collection view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/numberOfSections
func (c_ CollectionView) NumberOfSections() int {
	rv := objc.Send[int](c_.ID(), objc.RegisterName("numberOfSections"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/prefetchDataSource
func (c_ CollectionView) PrefetchDataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("prefetchDataSource"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/prefetchDataSource
func (c_ CollectionView) SetPrefetchDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setPrefetchDataSource:"), value)
}
// The set of index paths representing the currently selected items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/selectionIndexPaths
func (c_ CollectionView) SelectionIndexPaths() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("selectionIndexPaths"))
	return rv
}
// SetSelectionIndexPaths sets the value of the selectionIndexPaths property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/selectionIndexPaths
func (c_ CollectionView) SetSelectionIndexPaths(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setSelectionIndexPaths:"), value)
}
// The indexes of the currently selected items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/selectionIndexes
func (c_ CollectionView) SelectionIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID(), objc.RegisterName("selectionIndexes"))
	return rv
}
// SetSelectionIndexes sets the value of the selectionIndexes property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSCollectionView/selectionIndexes
func (c_ CollectionView) SetSelectionIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID(), objc.RegisterName("setSelectionIndexes:"), value)
}
