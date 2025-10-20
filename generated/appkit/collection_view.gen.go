// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [CollectionView] class.
var (
	collectionViewClass     _CollectionViewClass
	collectionViewClassOnce sync.Once
)

func getCollectionViewClass() _CollectionViewClass {
	collectionViewClassOnce.Do(func() {
		collectionViewClass = _CollectionViewClass{objc.GetClass("NSCollectionView")}
	})
	return collectionViewClass
}

type _CollectionViewClass struct {
	class objc.Class
}

// An interface definition for the [CollectionView] class.
type ICollectionView interface {
	IView
	DeleteItemsAtIndexPaths(indexPaths unsafe.Pointer)
	DeleteSections(sections unsafe.Pointer)
	DeselectAll(sender objc.ID)
	DeselectItemsAtIndexPaths(indexPaths unsafe.Pointer)
	DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer
	DraggingImageForItemsAtIndexesWithEventOffset(indexes unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer
	FrameForItemAtIndex(index uint) coregraphics.CGRect
	FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) coregraphics.CGRect
	IndexPathForItem(item unsafe.Pointer) unsafe.Pointer
	IndexPathForItemAtPoint(point coregraphics.CGPoint) unsafe.Pointer
	IndexPathsForVisibleItems() unsafe.Pointer
	IndexPathsForVisibleSupplementaryElementsOfKind(elementKind unsafe.Pointer) unsafe.Pointer
	InsertItemsAtIndexPaths(indexPaths unsafe.Pointer)
	InsertSections(sections unsafe.Pointer)
	ItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer
	ItemAtIndex(index uint) unsafe.Pointer
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
	RegisterNibForItemWithIdentifier(nib unsafe.Pointer, identifier unsafe.Pointer)
	RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind unsafe.Pointer, identifier unsafe.Pointer)
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
	VisibleItems() []CollectionViewItem
	VisibleSupplementaryViewsOfKind(elementKind unsafe.Pointer) []View<NSCollectionViewElement>
}

// An ordered collection of data items displayed in a customizable layout.
//
// The simplest type of collection view displays its items in a grid, but you can define layouts to arrange items however you like. For example, you might create a layout where items are arranged in a circle. You can also change layouts dynamically at runtime whenever you need to present items differently. You can add collection views to your interface using Interface Builder or create them programmatically in your view controller or window controller code. It is recommended that you configure your collection view with a data source object, which is an object that conforms to the protocol. Data sources support multiple sections and the modern layout architecture and are the preferred way for specifying your data. In addition to displaying items, collection views support the display of supplementary and decoration views. Support for supplementary and decoration views is defined by the current layout object, but both types of views add to the visual presentation of your content. Supplementary views are associated with a specific section and can be used to create header and footer views for a related group of items. Decoration views are purely visual adornments and can be used to implement dynamic backgrounds or other types of configurable visual content. The layout of a collection view can be changed dynamically by assigning a new layout object to the property. Changing the layout object updates the appearance of the collection view without animating the changes.
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

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewClass) Alloc() CollectionView {
	rv := objc.Send[CollectionView](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionViewClass) New() CollectionView {
	rv := objc.Send[CollectionView](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionView) Init() CollectionView {
	rv := objc.Send[CollectionView](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionView) Autorelease() CollectionView {
	rv := objc.Send[CollectionView](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionView creates a new CollectionView instance.
func NewCollectionView() CollectionView {
	return getCollectionViewClass().New()
}


// Deletes the items at the specified index paths.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deleteItems(at:)
func (c_ CollectionView) DeleteItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteItemsAtIndexPaths:"), indexPaths)
}

// Deletes the specified sections and their contained items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deleteSections(_:)
func (c_ CollectionView) DeleteSections(sections unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteSections:"), sections)
}

// Deselects all items in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deselectAll(_:)
func (c_ CollectionView) DeselectAll(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deselectAll:"), sender)
}

// Removes the specified items from the current selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deselectItems(at:)
func (c_ CollectionView) DeselectItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deselectItemsAtIndexPaths:"), indexPaths)
}

// Returns an image to use for dragging the specified items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-7rc4k
func (c_ CollectionView) DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("draggingImageForItemsAtIndexPaths:withEvent:offset:"), indexPaths, event, dragImageOffset)
	return rv
}

// This method computes and returns an image to use for dragging.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-951w7
func (c_ CollectionView) DraggingImageForItemsAtIndexesWithEventOffset(indexes unsafe.Pointer, event unsafe.Pointer, dragImageOffset unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("draggingImageForItemsAtIndexes:withEvent:offset:"), indexes, event, dragImageOffset)
	return rv
}

// Returns the frame of the collection view item at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/frameForItem(at:)
func (c_ CollectionView) FrameForItemAtIndex(index uint) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("frameForItemAtIndex:"), index)
	return rv
}

// Returns the frame of an item based on the number of items in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/frameForItem(at:withNumberOfItems:)
func (c_ CollectionView) FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("frameForItemAtIndex:withNumberOfItems:"), index, numberOfItems)
	return rv
}

// Returns the index path of the specified item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPath(for:)
func (c_ CollectionView) IndexPathForItem(item unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathForItem:"), item)
	return rv
}

// Returns the index path of the item at the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathForItem(at:)
func (c_ CollectionView) IndexPathForItemAtPoint(point coregraphics.CGPoint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathForItemAtPoint:"), point)
	return rv
}

// Returns the index paths of the currently active items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathsForVisibleItems()
func (c_ CollectionView) IndexPathsForVisibleItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsForVisibleItems"))
	return rv
}

// Returns the index paths of the currently active supplementary views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathsForVisibleSupplementaryElements(ofKind:)
func (c_ CollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsForVisibleSupplementaryElementsOfKind:"), elementKind)
	return rv
}

// Inserts new items into the collection view at the specified locations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/insertItems(at:)
func (c_ CollectionView) InsertItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertItemsAtIndexPaths:"), indexPaths)
}

// Inserts new sections at the specified indexes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/insertSections(_:)
func (c_ CollectionView) InsertSections(sections unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertSections:"), sections)
}

// Returns the item associated with the specified index path.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/item(at:)-2vx2h
func (c_ CollectionView) ItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("itemAtIndexPath:"), indexPath)
	return rv
}

// Returns the collection view item for the represented object at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/item(at:)-80xze
func (c_ CollectionView) ItemAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}

// Returns the layout information for the item at the specified index path.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/layoutAttributesForItem(at:)
func (c_ CollectionView) LayoutAttributesForItemAtIndexPath(indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("layoutAttributesForItemAtIndexPath:"), indexPath)
	return rv
}

// Returns the layout information for the supplementary view at the specified index path.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/layoutAttributesForSupplementaryElement(ofKind:at:)
func (c_ CollectionView) LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("layoutAttributesForSupplementaryElementOfKind:atIndexPath:"), kind, indexPath)
	return rv
}

// Creates or returns a reusable item object of the specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/makeItem(withIdentifier:for:)
func (c_ CollectionView) MakeItemWithIdentifierForIndexPath(identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("makeItemWithIdentifier:forIndexPath:"), identifier, indexPath)
	return rv
}

// Creates or returns a reusable supplementary view of the specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/makeSupplementaryView(ofKind:withIdentifier:for:)
func (c_ CollectionView) MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind unsafe.Pointer, identifier unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("makeSupplementaryViewOfKind:withIdentifier:forIndexPath:"), elementKind, identifier, indexPath)
	return rv
}

// Moves an item from one location to another in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/moveItem(at:to:)
func (c_ CollectionView) MoveItemAtIndexPathToIndexPath(indexPath unsafe.Pointer, newIndexPath unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("moveItemAtIndexPath:toIndexPath:"), indexPath, newIndexPath)
}

// Moves a section from its current location to a new location.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/moveSection(_:toSection:)
func (c_ CollectionView) MoveSectionToSection(section int, newSection int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("moveSection:toSection:"), section, newSection)
}

// Returns the collection view item that is used for the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/newItem(forRepresentedObject:)
func (c_ CollectionView) NewItemForRepresentedObject(object objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("newItemForRepresentedObject:"), object)
	return rv
}

// Returns the number of items in the specified section.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/numberOfItems(inSection:)
func (c_ CollectionView) NumberOfItemsInSection(section int) int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfItemsInSection:"), section)
	return rv
}

// Encapsulates multiple insert, delete, reload, and move operations into a single animated operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/performBatchUpdates(_:completionHandler:)
func (c_ CollectionView) PerformBatchUpdatesCompletionHandler(updates unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performBatchUpdates:completionHandler:"), updates, completionHandler)
}

// Registers a class to use when creating new items in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-6s4i
func (c_ CollectionView) RegisterClassForItemWithIdentifier(itemClass objc.Class, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerClass:forItemWithIdentifier:"), itemClass, identifier)
}

// Registers a nib file to use when creating items in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-90h1i
func (c_ CollectionView) RegisterNibForItemWithIdentifier(nib unsafe.Pointer, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerNib:forItemWithIdentifier:"), nib, identifier)
}

// Registers a class to use when creating new supplementary views in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-3dqa
func (c_ CollectionView) RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind unsafe.Pointer, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerClass:forSupplementaryViewOfKind:withIdentifier:"), viewClass, kind, identifier)
}

// Registers a nib file to use when creating supplementary views in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-7gvf2
func (c_ CollectionView) RegisterNibForSupplementaryViewOfKindWithIdentifier(nib unsafe.Pointer, kind unsafe.Pointer, identifier unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerNib:forSupplementaryViewOfKind:withIdentifier:"), nib, kind, identifier)
}

// Reloads all of the data for the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadData()
func (c_ CollectionView) ReloadData() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadData"))
}

// Reloads only the specified items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadItems(at:)
func (c_ CollectionView) ReloadItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadItemsAtIndexPaths:"), indexPaths)
}

// Reloads the data in the specified sections of the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadSections(_:)
func (c_ CollectionView) ReloadSections(sections unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadSections:"), sections)
}

// Scrolls the collection view contents until the specified items are visible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/scrollToItems(at:scrollPosition:)
func (c_ CollectionView) ScrollToItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollToItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}

// Selects all items in the collection view, if doing so is possible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectAll(_:)
func (c_ CollectionView) SelectAll(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectAll:"), sender)
}

// Adds the specified items to the current selection and optionally scrolls the items into position.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectItems(at:scrollPosition:)
func (c_ CollectionView) SelectItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}

// Configures the drag operation mask.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/setDraggingSourceOperationMask(_:forLocal:)
func (c_ CollectionView) SetDraggingSourceOperationMaskForLocal(dragOperationMask unsafe.Pointer, localDestination bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), dragOperationMask, localDestination)
}

// Returns the supplementary view associated with the specified index path.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/supplementaryView(forElementKind:at:)
func (c_ CollectionView) SupplementaryViewForElementKindAtIndexPath(elementKind unsafe.Pointer, indexPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supplementaryViewForElementKind:atIndexPath:"), elementKind, indexPath)
	return rv
}

// Collapses the section in which the sender resides into a single horizontally scrollable row.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/toggleSectionCollapse(_:)
func (c_ CollectionView) ToggleSectionCollapse(sender objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("toggleSectionCollapse:"), sender)
}

// Returns an array of the actively managed items in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/visibleItems()
func (c_ CollectionView) VisibleItems() []CollectionViewItem {
	rv := objc.Send[[]CollectionViewItem](c_.ID, objc.Sel("visibleItems"))
	return rv
}

// Returns an array of the actively managed supplementary views in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/visibleSupplementaryViews(ofKind:)
func (c_ CollectionView) VisibleSupplementaryViewsOfKind(elementKind unsafe.Pointer) []View<NSCollectionViewElement> {
	rv := objc.Send[[]View<NSCollectionViewElement>](c_.ID, objc.Sel("visibleSupplementaryViewsOfKind:"), elementKind)
	return rv
}

// A Boolean value indicating whether the collection view may have no selected items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsEmptySelection
func (c_ CollectionView) AllowsEmptySelection() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}

// SetAllowsEmptySelection sets the value of the allowsEmptySelection property.
// A Boolean value indicating whether the collection view may have no selected items.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsEmptySelection
func (c_ CollectionView) SetAllowsEmptySelection(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsEmptySelection:"), value)
}
// A Boolean value that indicates whether the user may select more than one item in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsMultipleSelection
func (c_ CollectionView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}

// SetAllowsMultipleSelection sets the value of the allowsMultipleSelection property.
// A Boolean value that indicates whether the user may select more than one item in the collection view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsMultipleSelection
func (c_ CollectionView) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}
// An array containing the collection view’s background colors.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundColors
func (c_ CollectionView) BackgroundColors() []Color {
	rv := objc.Send[[]Color](c_.ID, objc.Sel("backgroundColors"))
	return rv
}

// SetBackgroundColors sets the value of the backgroundColors property.
// An array containing the collection view’s background colors.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundColors
func (c_ CollectionView) SetBackgroundColors(value []Color) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundColors:"), value)
}
// The background view placed behind all items and supplementary views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundView
func (c_ CollectionView) BackgroundView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("backgroundView"))
	return rv
}

// SetBackgroundView sets the value of the backgroundView property.
// The background view placed behind all items and supplementary views.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundView
func (c_ CollectionView) SetBackgroundView(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundView:"), value)
}
// A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundViewScrollsWithContent
func (c_ CollectionView) BackgroundViewScrollsWithContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("backgroundViewScrollsWithContent"))
	return rv
}

// SetBackgroundViewScrollsWithContent sets the value of the backgroundViewScrollsWithContent property.
// A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundViewScrollsWithContent
func (c_ CollectionView) SetBackgroundViewScrollsWithContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundViewScrollsWithContent:"), value)
}
// The layout object used to organize the collection view’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/collectionViewLayout
func (c_ CollectionView) CollectionViewLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("collectionViewLayout"))
	return rv
}

// SetCollectionViewLayout sets the value of the collectionViewLayout property.
// The layout object used to organize the collection view’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/collectionViewLayout
func (c_ CollectionView) SetCollectionViewLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollectionViewLayout:"), value)
}
// An array that provides data for the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/content
func (c_ CollectionView) Content() []objc.ID {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("content"))
	return rv
}

// SetContent sets the value of the content property.
// An array that provides data for the collection view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/content
func (c_ CollectionView) SetContent(value []objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContent:"), value)
}
// An object that provides data for the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/dataSource
func (c_ CollectionView) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dataSource"))
	return rv
}

// SetDataSource sets the value of the dataSource property.
// An object that provides data for the collection view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/dataSource
func (c_ CollectionView) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}
// The collection view’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/delegate
func (c_ CollectionView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}

// SetDelegate sets the value of the delegate property.
// The collection view’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/delegate
func (c_ CollectionView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}
// A Boolean value indicating whether the collection view is the first responder.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/isFirstResponder
func (c_ CollectionView) FirstResponder() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("firstResponder"))
	return rv
}
// A Boolean value that indicates whether the user may select items in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/isSelectable
func (c_ CollectionView) Selectable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("selectable"))
	return rv
}

// SetSelectable sets the value of the selectable property.
// A Boolean value that indicates whether the user may select items in the collection view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/isSelectable
func (c_ CollectionView) SetSelectable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectable:"), value)
}
// The receiver’s collection view item prototype.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/itemPrototype
func (c_ CollectionView) ItemPrototype() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("itemPrototype"))
	return rv
}

// SetItemPrototype sets the value of the itemPrototype property.
// The receiver’s collection view item prototype.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/itemPrototype
func (c_ CollectionView) SetItemPrototype(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemPrototype:"), value)
}
// The maximum size (in points) of items in the collection view grid.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxItemSize
func (c_ CollectionView) MaxItemSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("maxItemSize"))
	return rv
}

// SetMaxItemSize sets the value of the maxItemSize property.
// The maximum size (in points) of items in the collection view grid.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxItemSize
func (c_ CollectionView) SetMaxItemSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxItemSize:"), value)
}
// The maximum number of columns that the collection view displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxNumberOfColumns
func (c_ CollectionView) MaxNumberOfColumns() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxNumberOfColumns"))
	return rv
}

// SetMaxNumberOfColumns sets the value of the maxNumberOfColumns property.
// The maximum number of columns that the collection view displays.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxNumberOfColumns
func (c_ CollectionView) SetMaxNumberOfColumns(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxNumberOfColumns:"), value)
}
// The maximum number of rows that the collection view displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxNumberOfRows
func (c_ CollectionView) MaxNumberOfRows() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxNumberOfRows"))
	return rv
}

// SetMaxNumberOfRows sets the value of the maxNumberOfRows property.
// The maximum number of rows that the collection view displays.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxNumberOfRows
func (c_ CollectionView) SetMaxNumberOfRows(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxNumberOfRows:"), value)
}
// The minimum size (in points) of items in the collection view grid.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/minItemSize
func (c_ CollectionView) MinItemSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("minItemSize"))
	return rv
}

// SetMinItemSize sets the value of the minItemSize property.
// The minimum size (in points) of items in the collection view grid.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/minItemSize
func (c_ CollectionView) SetMinItemSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinItemSize:"), value)
}
// The number of sections in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/numberOfSections
func (c_ CollectionView) NumberOfSections() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfSections"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/prefetchDataSource
func (c_ CollectionView) PrefetchDataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("prefetchDataSource"))
	return rv
}

// SetPrefetchDataSource sets the value of the prefetchDataSource property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/prefetchDataSource
func (c_ CollectionView) SetPrefetchDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrefetchDataSource:"), value)
}
// The set of index paths representing the currently selected items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexPaths
func (c_ CollectionView) SelectionIndexPaths() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("selectionIndexPaths"))
	return rv
}

// SetSelectionIndexPaths sets the value of the selectionIndexPaths property.
// The set of index paths representing the currently selected items.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexPaths
func (c_ CollectionView) SetSelectionIndexPaths(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectionIndexPaths:"), value)
}
// The indexes of the currently selected items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexes
func (c_ CollectionView) SelectionIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("selectionIndexes"))
	return rv
}

// SetSelectionIndexes sets the value of the selectionIndexes property.
// The indexes of the currently selected items.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexes
func (c_ CollectionView) SetSelectionIndexes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectionIndexes:"), value)
}


