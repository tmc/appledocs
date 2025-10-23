// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionView] class.
var (
	CollectionViewClass     _CollectionViewClass
	CollectionViewClassOnce sync.Once
)

func getCollectionViewClass() _CollectionViewClass {
	CollectionViewClassOnce.Do(func() {
		CollectionViewClass = _CollectionViewClass{objc.GetClass("NSCollectionView")}
	})
	return CollectionViewClass
}

type _CollectionViewClass struct {
	class objc.Class
}

// An interface definition for the [CollectionView] class.
type ICollectionView interface {
	IView
	// properties:
	AllowsEmptySelection() bool /* primitive/slice/pointer. */
	SetAllowsEmptySelection(value bool /* primitive/slice/pointer. */)
	AllowsMultipleSelection() bool /* primitive/slice/pointer. */
	SetAllowsMultipleSelection(value bool /* primitive/slice/pointer. */)
	BackgroundColors() []Color /* primitive/slice/pointer. */
	SetBackgroundColors(value []Color /* primitive/slice/pointer. */)
	BackgroundView() IView
	SetBackgroundView(value IView)
	BackgroundViewScrollsWithContent() bool /* primitive/slice/pointer. */
	SetBackgroundViewScrollsWithContent(value bool /* primitive/slice/pointer. */)
	CollectionViewLayout() objc.IObject /* cross-framework: CollectionViewLayout */
	SetCollectionViewLayout(value objc.IObject /* cross-framework: CollectionViewLayout */)
	Content() []objc.ID /* already interface */
	SetContent(value []objc.ID /* already interface */)
	DataSource() objc.ID
	SetDataSource(value objc.ID)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	FirstResponder() bool /* primitive/slice/pointer. */
	Selectable() bool /* primitive/slice/pointer. */
	SetSelectable(value bool /* primitive/slice/pointer. */)
	ItemPrototype() ICollectionViewItem
	SetItemPrototype(value ICollectionViewItem)
	MaxItemSize() coregraphics.CGSize
	SetMaxItemSize(value coregraphics.CGSize)
	MaxNumberOfColumns() uint /* primitive/slice/pointer. */
	SetMaxNumberOfColumns(value uint /* primitive/slice/pointer. */)
	MaxNumberOfRows() uint /* primitive/slice/pointer. */
	SetMaxNumberOfRows(value uint /* primitive/slice/pointer. */)
	MinItemSize() coregraphics.CGSize
	SetMinItemSize(value coregraphics.CGSize)
	NumberOfSections() int /* primitive/slice/pointer. */
	PrefetchDataSource() objc.ID
	SetPrefetchDataSource(value objc.ID)
	SelectionIndexPaths() unsafe.Pointer
	SetSelectionIndexPaths(value unsafe.Pointer)
	SelectionIndexes() objc.IObject /* cross-framework: IndexSet */
	SetSelectionIndexes(value objc.IObject /* cross-framework: IndexSet */)
	IsFirstResponder() bool /* primitive/slice/pointer. */
	SetIsFirstResponder(value bool /* primitive/slice/pointer. */)
	IsSelectable() bool /* primitive/slice/pointer. */
	SetIsSelectable(value bool /* primitive/slice/pointer. */)
	// methods:
	DeleteItemsAtIndexPaths(indexPaths unsafe.Pointer)
	DeleteSections(sections objc.IObject /* cross-framework IndexSet */)
	DeselectAll(sender objectivec.IObject)
	DeselectItemsAtIndexPaths(indexPaths unsafe.Pointer)
	DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths unsafe.Pointer, event IEvent, dragImageOffset PointPointer /* not a class type */) IImage
	DraggingImageForItemsAtIndexesWithEventOffset(indexes objc.IObject /* cross-framework IndexSet */, event IEvent, dragImageOffset PointPointer /* not a class type */) IImage
	FrameForItemAtIndex(index uint /* primitive/slice/pointer. */) coregraphics.CGRect
	FrameForItemAtIndexWithNumberOfItems(index uint /* primitive/slice/pointer. */, numberOfItems uint /* primitive/slice/pointer. */) coregraphics.CGRect
	IndexPathForItem(item ICollectionViewItem) IndexPath /* not a class type */
	IndexPathForItemAtPoint(point coregraphics.CGPoint) IndexPath /* not a class type */
	IndexPathsForVisibleItems() unsafe.Pointer
	IndexPathsForVisibleSupplementaryElementsOfKind(elementKind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */) unsafe.Pointer
	InsertItemsAtIndexPaths(indexPaths unsafe.Pointer)
	InsertSections(sections objc.IObject /* cross-framework IndexSet */)
	ItemAtIndexPath(indexPath IndexPath /* not a class type */) ICollectionViewItem
	ItemAtIndex(index uint /* primitive/slice/pointer. */) ICollectionViewItem
	LayoutAttributesForItemAtIndexPath(indexPath IndexPath /* not a class type */) CollectionViewLayoutAttributes /* not a class type */
	LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, indexPath IndexPath /* not a class type */) CollectionViewLayoutAttributes /* not a class type */
	MakeItemWithIdentifierForIndexPath(identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */, indexPath IndexPath /* not a class type */) ICollectionViewItem
	MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */, indexPath IndexPath /* not a class type */) IView
	MoveItemAtIndexPathToIndexPath(indexPath IndexPath /* not a class type */, newIndexPath IndexPath /* not a class type */)
	MoveSectionToSection(section int /* primitive/slice/pointer. */, newSection int /* primitive/slice/pointer. */)
	NumberOfItemsInSection(section int /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */
	PerformBatchUpdatesCompletionHandler(updates unsafe.Pointer, completionHandler unsafe.Pointer)
	RegisterClassForItemWithIdentifier(itemClass objc.Class, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */)
	RegisterNibForItemWithIdentifier(nib INib, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */)
	RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */)
	RegisterNibForSupplementaryViewOfKindWithIdentifier(nib INib, kind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */)
	ReloadData()
	ReloadItemsAtIndexPaths(indexPaths unsafe.Pointer)
	ReloadSections(sections objc.IObject /* cross-framework IndexSet */)
	ScrollToItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition CollectionViewScrollPosition)
	SelectAll(sender objectivec.IObject)
	SelectItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition CollectionViewScrollPosition)
	SetDraggingSourceOperationMaskForLocal(dragOperationMask DragOperation, localDestination bool /* primitive/slice/pointer. */)
	SupplementaryViewForElementKindAtIndexPath(elementKind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, indexPath IndexPath /* not a class type */) unsafe.Pointer
	ToggleSectionCollapse(sender objectivec.IObject)
	VisibleItems() []CollectionViewItem /* primitive/slice/pointer. */
	VisibleSupplementaryViewsOfKind(elementKind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */) []View /* primitive/slice/pointer. */
}

// An ordered collection of data items displayed in a customizable layout.
//
// The simplest type of collection view displays its items in a grid, but you can define layouts to arrange items however you like. For example, you might create a layout where items are arranged in a circle. You can also change layouts dynamically at runtime whenever you need to present items differently. You can add collection views to your interface using Interface Builder or create them programmatically in your view controller or window controller code. It is recommended that you configure your collection view with a data source object, which is an object that conforms to the protocol. Data sources support multiple sections and the modern layout architecture and are the preferred way for specifying your data. In addition to displaying items, collection views support the display of supplementary and decoration views. Support for supplementary and decoration views is defined by the current layout object, but both types of views add to the visual presentation of your content. Supplementary views are associated with a specific section and can be used to create header and footer views for a related group of items. Decoration views are purely visual adornments and can be used to implement dynamic backgrounds or other types of configurable visual content. The layout of a collection view can be changed dynamically by assigning a new layout object to the property. Changing the layout object updates the appearance of the collection view without animating the changes.


// An ordered collection of data items displayed in a customizable layout.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deleteItems(at:)
func (c_ CollectionView) DeleteItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteItemsAtIndexPaths:"), indexPaths)
}


// Deletes the specified sections and their contained items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deleteSections(_:)
func (c_ CollectionView) DeleteSections(sections objc.IObject /* cross-framework IndexSet */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deleteSections:"), sections)
}


// Deselects all items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deselectAll(_:)
func (c_ CollectionView) DeselectAll(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deselectAll:"), sender)
}


// Removes the specified items from the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/deselectItems(at:)
func (c_ CollectionView) DeselectItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("deselectItemsAtIndexPaths:"), indexPaths)
}


// Returns an image to use for dragging the specified items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-7rc4k
func (c_ CollectionView) DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths unsafe.Pointer, event IEvent, dragImageOffset PointPointer /* not a class type */) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("draggingImageForItemsAtIndexPaths:withEvent:offset:"), indexPaths, event, dragImageOffset)
	return rv
}


// This method computes and returns an image to use for dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-951w7
func (c_ CollectionView) DraggingImageForItemsAtIndexesWithEventOffset(indexes objc.IObject /* cross-framework IndexSet */, event IEvent, dragImageOffset PointPointer /* not a class type */) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("draggingImageForItemsAtIndexes:withEvent:offset:"), indexes, event, dragImageOffset)
	return rv
}


// Returns the frame of the collection view item at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/frameForItem(at:)
func (c_ CollectionView) FrameForItemAtIndex(index uint /* primitive/slice/pointer. */) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("frameForItemAtIndex:"), index)
	return rv
}


// Returns the frame of an item based on the number of items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/frameForItem(at:withNumberOfItems:)
func (c_ CollectionView) FrameForItemAtIndexWithNumberOfItems(index uint /* primitive/slice/pointer. */, numberOfItems uint /* primitive/slice/pointer. */) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("frameForItemAtIndex:withNumberOfItems:"), index, numberOfItems)
	return rv
}


// Returns the index path of the specified item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPath(for:)
func (c_ CollectionView) IndexPathForItem(item ICollectionViewItem) IndexPath /* not a class type */ {
	rv := objc.Send[IndexPath](c_.ID, objc.Sel("indexPathForItem:"), item)
	return rv
}


// Returns the index path of the item at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathForItem(at:)
func (c_ CollectionView) IndexPathForItemAtPoint(point coregraphics.CGPoint) IndexPath /* not a class type */ {
	rv := objc.Send[IndexPath](c_.ID, objc.Sel("indexPathForItemAtPoint:"), point)
	return rv
}


// Returns the index paths of the currently active items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathsForVisibleItems()
func (c_ CollectionView) IndexPathsForVisibleItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsForVisibleItems"))
	return rv
}


// Returns the index paths of the currently active supplementary views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathsForVisibleSupplementaryElements(ofKind:)
func (c_ CollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("indexPathsForVisibleSupplementaryElementsOfKind:"), elementKind)
	return rv
}


// Inserts new items into the collection view at the specified locations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/insertItems(at:)
func (c_ CollectionView) InsertItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertItemsAtIndexPaths:"), indexPaths)
}


// Inserts new sections at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/insertSections(_:)
func (c_ CollectionView) InsertSections(sections objc.IObject /* cross-framework IndexSet */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("insertSections:"), sections)
}


// Returns the item associated with the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/item(at:)-2vx2h
func (c_ CollectionView) ItemAtIndexPath(indexPath IndexPath /* not a class type */) ICollectionViewItem {
	rv := objc.Send[CollectionViewItem](c_.ID, objc.Sel("itemAtIndexPath:"), indexPath)
	return rv
}


// Returns the collection view item for the represented object at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/item(at:)-80xze
func (c_ CollectionView) ItemAtIndex(index uint /* primitive/slice/pointer. */) ICollectionViewItem {
	rv := objc.Send[CollectionViewItem](c_.ID, objc.Sel("itemAtIndex:"), index)
	return rv
}


// Returns the layout information for the item at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/layoutAttributesForItem(at:)
func (c_ CollectionView) LayoutAttributesForItemAtIndexPath(indexPath IndexPath /* not a class type */) CollectionViewLayoutAttributes /* not a class type */ {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("layoutAttributesForItemAtIndexPath:"), indexPath)
	return rv
}


// Returns the layout information for the supplementary view at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/layoutAttributesForSupplementaryElement(ofKind:at:)
func (c_ CollectionView) LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, indexPath IndexPath /* not a class type */) CollectionViewLayoutAttributes /* not a class type */ {
	rv := objc.Send[CollectionViewLayoutAttributes](c_.ID, objc.Sel("layoutAttributesForSupplementaryElementOfKind:atIndexPath:"), kind, indexPath)
	return rv
}


// Creates or returns a reusable item object of the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/makeItem(withIdentifier:for:)
func (c_ CollectionView) MakeItemWithIdentifierForIndexPath(identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */, indexPath IndexPath /* not a class type */) ICollectionViewItem {
	rv := objc.Send[CollectionViewItem](c_.ID, objc.Sel("makeItemWithIdentifier:forIndexPath:"), identifier, indexPath)
	return rv
}


// Creates or returns a reusable supplementary view of the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/makeSupplementaryView(ofKind:withIdentifier:for:)
func (c_ CollectionView) MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */, indexPath IndexPath /* not a class type */) IView {
	rv := objc.Send[View](c_.ID, objc.Sel("makeSupplementaryViewOfKind:withIdentifier:forIndexPath:"), elementKind, identifier, indexPath)
	return rv
}


// Moves an item from one location to another in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/moveItem(at:to:)
func (c_ CollectionView) MoveItemAtIndexPathToIndexPath(indexPath IndexPath /* not a class type */, newIndexPath IndexPath /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("moveItemAtIndexPath:toIndexPath:"), indexPath, newIndexPath)
}


// Moves a section from its current location to a new location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/moveSection(_:toSection:)
func (c_ CollectionView) MoveSectionToSection(section int /* primitive/slice/pointer. */, newSection int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("moveSection:toSection:"), section, newSection)
}


// Returns the number of items in the specified section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/numberOfItems(inSection:)
func (c_ CollectionView) NumberOfItemsInSection(section int /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfItemsInSection:"), section)
	return rv
}


// Encapsulates multiple insert, delete, reload, and move operations into a single animated operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/performBatchUpdates(_:completionHandler:)
func (c_ CollectionView) PerformBatchUpdatesCompletionHandler(updates unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performBatchUpdates:completionHandler:"), updates, completionHandler)
}


// Registers a class to use when creating new items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-6s4i
func (c_ CollectionView) RegisterClassForItemWithIdentifier(itemClass objc.Class, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerClass:forItemWithIdentifier:"), itemClass, identifier)
}


// Registers a nib file to use when creating items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-90h1i
func (c_ CollectionView) RegisterNibForItemWithIdentifier(nib INib, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerNib:forItemWithIdentifier:"), nib, identifier)
}


// Registers a class to use when creating new supplementary views in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-3dqa
func (c_ CollectionView) RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerClass:forSupplementaryViewOfKind:withIdentifier:"), viewClass, kind, identifier)
}


// Registers a nib file to use when creating supplementary views in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-7gvf2
func (c_ CollectionView) RegisterNibForSupplementaryViewOfKindWithIdentifier(nib INib, kind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, identifier objc.IObject /* cross-framework UserInterfaceItemIdentifier */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("registerNib:forSupplementaryViewOfKind:withIdentifier:"), nib, kind, identifier)
}


// Reloads all of the data for the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadData()
func (c_ CollectionView) ReloadData() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadData"))
}


// Reloads only the specified items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadItems(at:)
func (c_ CollectionView) ReloadItemsAtIndexPaths(indexPaths unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadItemsAtIndexPaths:"), indexPaths)
}


// Reloads the data in the specified sections of the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadSections(_:)
func (c_ CollectionView) ReloadSections(sections objc.IObject /* cross-framework IndexSet */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadSections:"), sections)
}


// Scrolls the collection view contents until the specified items are visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/scrollToItems(at:scrollPosition:)
func (c_ CollectionView) ScrollToItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition CollectionViewScrollPosition) {
	objc.Send[objc.ID](c_.ID, objc.Sel("scrollToItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}


// Selects all items in the collection view, if doing so is possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectAll(_:)
func (c_ CollectionView) SelectAll(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectAll:"), sender)
}


// Adds the specified items to the current selection and optionally scrolls the items into position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectItems(at:scrollPosition:)
func (c_ CollectionView) SelectItemsAtIndexPathsScrollPosition(indexPaths unsafe.Pointer, scrollPosition CollectionViewScrollPosition) {
	objc.Send[objc.ID](c_.ID, objc.Sel("selectItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}


// Configures the drag operation mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/setDraggingSourceOperationMask(_:forLocal:)
func (c_ CollectionView) SetDraggingSourceOperationMaskForLocal(dragOperationMask DragOperation, localDestination bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), dragOperationMask, localDestination)
}


// Returns the supplementary view associated with the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/supplementaryView(forElementKind:at:)
func (c_ CollectionView) SupplementaryViewForElementKindAtIndexPath(elementKind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */, indexPath IndexPath /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supplementaryViewForElementKind:atIndexPath:"), elementKind, indexPath)
	return rv
}


// Collapses the section in which the sender resides into a single horizontally scrollable row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/toggleSectionCollapse(_:)
func (c_ CollectionView) ToggleSectionCollapse(sender objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("toggleSectionCollapse:"), sender)
}


// Returns an array of the actively managed items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/visibleItems()
func (c_ CollectionView) VisibleItems() []CollectionViewItem /* primitive/slice/pointer. */ {
	rv := objc.Send[[]CollectionViewItem](c_.ID, objc.Sel("visibleItems"))
	return rv
}


// Returns an array of the actively managed supplementary views in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/visibleSupplementaryViews(ofKind:)
func (c_ CollectionView) VisibleSupplementaryViewsOfKind(elementKind objc.IObject /* cross-framework CollectionViewSupplementaryElementKind */) []View /* primitive/slice/pointer. */ {
	rv := objc.Send[[]View](c_.ID, objc.Sel("visibleSupplementaryViewsOfKind:"), elementKind)
	return rv
}


// A Boolean value indicating whether the collection view may have no selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsEmptySelection
func (c_ CollectionView) AllowsEmptySelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}


// A Boolean value indicating whether the collection view may have no selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsEmptySelection
func (c_ CollectionView) SetAllowsEmptySelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsEmptySelection:"), value)
}


// A Boolean value that indicates whether the user may select more than one item in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsMultipleSelection
func (c_ CollectionView) AllowsMultipleSelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// A Boolean value that indicates whether the user may select more than one item in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsMultipleSelection
func (c_ CollectionView) SetAllowsMultipleSelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}


// An array containing the collection view’s background colors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundColors
func (c_ CollectionView) BackgroundColors() []Color /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Color](c_.ID, objc.Sel("backgroundColors"))
	return rv
}


// An array containing the collection view’s background colors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundColors
func (c_ CollectionView) SetBackgroundColors(value []Color /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundColors:"), nsArray)
}


// The background view placed behind all items and supplementary views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundView
func (c_ CollectionView) BackgroundView() IView {
	rv := objc.Send[View](c_.ID, objc.Sel("backgroundView"))
	return rv
}


// The background view placed behind all items and supplementary views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundView
func (c_ CollectionView) SetBackgroundView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundView:"), value)
}


// A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundViewScrollsWithContent
func (c_ CollectionView) BackgroundViewScrollsWithContent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("backgroundViewScrollsWithContent"))
	return rv
}


// A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundViewScrollsWithContent
func (c_ CollectionView) SetBackgroundViewScrollsWithContent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBackgroundViewScrollsWithContent:"), value)
}


// The layout object used to organize the collection view’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/collectionViewLayout
func (c_ CollectionView) CollectionViewLayout() objc.IObject /* cross-framework: CollectionViewLayout */ {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("collectionViewLayout"))
	return rv
}


// The layout object used to organize the collection view’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/collectionViewLayout
func (c_ CollectionView) SetCollectionViewLayout(value objc.IObject /* cross-framework: CollectionViewLayout */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCollectionViewLayout:"), value)
}


// An array that provides data for the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/content
func (c_ CollectionView) Content() []objc.ID /* already interface */ {
	rv := objc.Send[[]objc.ID](c_.ID, objc.Sel("content"))
	return rv
}


// An array that provides data for the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/content
func (c_ CollectionView) SetContent(value []objc.ID /* already interface */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setContent:"), nsArray)
}


// An object that provides data for the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/dataSource
func (c_ CollectionView) DataSource() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}


// An object that provides data for the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/dataSource
func (c_ CollectionView) SetDataSource(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}


// The collection view’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/delegate
func (c_ CollectionView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// The collection view’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/delegate
func (c_ CollectionView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value indicating whether the collection view is the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/isFirstResponder
func (c_ CollectionView) FirstResponder() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("firstResponder"))
	return rv
}


// A Boolean value that indicates whether the user may select items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/isSelectable
func (c_ CollectionView) Selectable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("selectable"))
	return rv
}


// A Boolean value that indicates whether the user may select items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/isSelectable
func (c_ CollectionView) SetSelectable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectable:"), value)
}


// The receiver’s collection view item prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/itemPrototype
func (c_ CollectionView) ItemPrototype() ICollectionViewItem {
	rv := objc.Send[CollectionViewItem](c_.ID, objc.Sel("itemPrototype"))
	return rv
}


// The receiver’s collection view item prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/itemPrototype
func (c_ CollectionView) SetItemPrototype(value ICollectionViewItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setItemPrototype:"), value)
}


// The maximum size (in points) of items in the collection view grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxItemSize
func (c_ CollectionView) MaxItemSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("maxItemSize"))
	return rv
}


// The maximum size (in points) of items in the collection view grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxItemSize
func (c_ CollectionView) SetMaxItemSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxItemSize:"), value)
}


// The maximum number of columns that the collection view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxNumberOfColumns
func (c_ CollectionView) MaxNumberOfColumns() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxNumberOfColumns"))
	return rv
}


// The maximum number of columns that the collection view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxNumberOfColumns
func (c_ CollectionView) SetMaxNumberOfColumns(value uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxNumberOfColumns:"), value)
}


// The maximum number of rows that the collection view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxNumberOfRows
func (c_ CollectionView) MaxNumberOfRows() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](c_.ID, objc.Sel("maxNumberOfRows"))
	return rv
}


// The maximum number of rows that the collection view displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/maxNumberOfRows
func (c_ CollectionView) SetMaxNumberOfRows(value uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxNumberOfRows:"), value)
}


// The minimum size (in points) of items in the collection view grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/minItemSize
func (c_ CollectionView) MinItemSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("minItemSize"))
	return rv
}


// The minimum size (in points) of items in the collection view grid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/minItemSize
func (c_ CollectionView) SetMinItemSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinItemSize:"), value)
}


// The number of sections in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/numberOfSections
func (c_ CollectionView) NumberOfSections() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfSections"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/prefetchDataSource
func (c_ CollectionView) PrefetchDataSource() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("prefetchDataSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/prefetchDataSource
func (c_ CollectionView) SetPrefetchDataSource(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrefetchDataSource:"), value)
}


// The set of index paths representing the currently selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexPaths
func (c_ CollectionView) SelectionIndexPaths() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("selectionIndexPaths"))
	return rv
}


// The set of index paths representing the currently selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexPaths
func (c_ CollectionView) SetSelectionIndexPaths(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectionIndexPaths:"), value)
}


// The indexes of the currently selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexes
func (c_ CollectionView) SelectionIndexes() objc.IObject /* cross-framework: IndexSet */ {
	rv := objc.Send[IndexSet](c_.ID, objc.Sel("selectionIndexes"))
	return rv
}


// The indexes of the currently selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexes
func (c_ CollectionView) SetSelectionIndexes(value objc.IObject /* cross-framework: IndexSet */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSelectionIndexes:"), value)
}


// A Boolean value indicating whether the collection view is the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/isfirstresponder
func (c_ CollectionView) IsFirstResponder() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFirstResponder"))
	return rv
}


// A Boolean value indicating whether the collection view is the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/isfirstresponder
func (c_ CollectionView) SetIsFirstResponder(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFirstResponder:"), value)
}


// A Boolean value that indicates whether the user may select items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/isselectable
func (c_ CollectionView) IsSelectable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSelectable"))
	return rv
}


// A Boolean value that indicates whether the user may select items in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionview/isselectable
func (c_ CollectionView) SetIsSelectable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSelectable:"), value)
}



