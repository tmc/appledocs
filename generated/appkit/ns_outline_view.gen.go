// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OutlineView] class.
var (
	OutlineViewClass     _OutlineViewClass
	OutlineViewClassOnce sync.Once
)

func getOutlineViewClass() _OutlineViewClass {
	OutlineViewClassOnce.Do(func() {
		OutlineViewClass = _OutlineViewClass{objc.GetClass("NSOutlineView")}
	})
	return OutlineViewClass
}

type _OutlineViewClass struct {
	class objc.Class
}

// An interface definition for the [OutlineView] class.
type IOutlineView interface {
	ITableView
	// properties:
	AutoresizesOutlineColumn() bool
	SetAutoresizesOutlineColumn(value bool)
	AutosaveExpandedItems() bool
	SetAutosaveExpandedItems(value bool)
	DataSource() objc.ID
	SetDataSource(value objc.ID)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	IndentationMarkerFollowsCell() bool
	SetIndentationMarkerFollowsCell(value bool)
	IndentationPerLevel() float64
	SetIndentationPerLevel(value float64)
	OutlineTableColumn() ITableColumn
	SetOutlineTableColumn(value ITableColumn)
	StronglyReferencesItems() bool
	SetStronglyReferencesItems(value bool)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	// methods:
	ChildOfItem(index int, item objc.IObject) objc.ID
	ChildIndexForItem(item objc.IObject) int
	CollapseItem(item objc.IObject)
	CollapseItemCollapseChildren(item objc.IObject, collapseChildren bool)
	ExpandItem(item objc.IObject)
	ExpandItemExpandChildren(item objc.IObject, expandChildren bool)
	FrameOfOutlineCellAtRow(row int) objc.IObject /* cross-framework: Rect */
	InsertItemsAtIndexesInParentWithAnimation(indexes foundation.IndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	IsExpandable(item objc.IObject) bool
	IsItemExpanded(item objc.IObject) bool
	ItemAtRow(row int) objc.ID
	LevelForItem(item objc.IObject) int
	LevelForRow(row int) int
	MoveItemAtIndexInParentToIndexInParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject)
	NumberOfChildrenOfItem(item objc.IObject) int
	ParentForItem(item objc.IObject) objc.ID
	ReloadItem(item objc.IObject)
	ReloadItemReloadChildren(item objc.IObject, reloadChildren bool)
	RemoveItemsAtIndexesInParentWithAnimation(indexes foundation.IndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	RowForItem(item objc.IObject) int
	SetDropItemDropChildIndex(item objc.IObject, index int)
	ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool
}

// A view that uses a row-and-column format to display hierarchical data like directories and files that can be expanded and collapsed.
//
// Like a table view, an outline view does not store its own data, instead it retrieves data values as needed from a data source to which it has a weak reference (see ). See , which declares the methods that an object uses to access the contents of its data source object. An outline view has the following features: A user can expand and collapse rows, edit values, and resize and rearrange columns. Each item in the outline view must be unique. In order for the collapsed state to remain consistent between reloads the item’s pointer must remain the same and the item must maintain sameness. The view gets data from a data source (see ). The view retrieves only the data that needs to be displayed. For more information about using NSOutlineView in your app, see .


// A view that uses a row-and-column format to display hierarchical data like directories and files that can be expanded and collapsed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView
type OutlineView struct {
	TableView
}

// OutlineViewFrom constructs a [OutlineView] from an unsafe.Pointer.
//
// A view that uses a row-and-column format to display hierarchical data like directories and files that can be expanded and collapsed.
func OutlineViewFrom(ptr unsafe.Pointer) OutlineView {
	return OutlineView{
		TableView: TableViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc _OutlineViewClass) Alloc() OutlineView {
	rv := objc.Send[OutlineView](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OutlineViewClass) New() OutlineView {
	rv := objc.Send[OutlineView](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OutlineView) Init() OutlineView {
	rv := objc.Send[OutlineView](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OutlineView) Autorelease() OutlineView {
	rv := objc.Send[OutlineView](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOutlineView creates a new OutlineView instance.
func NewOutlineView() OutlineView {
	return getOutlineViewClass().New()
}



// Returns the specified child of an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/child(_:ofItem:)
func (o_ OutlineView) ChildOfItem(index int, item objc.IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("child:ofItem:"), index, item)
	return rv
}


// Returns the child index of the specified item within its parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/childIndex(forItem:)
func (o_ OutlineView) ChildIndexForItem(item objc.IObject) int {
	rv := objc.Send[int](o_.ID, objc.Sel("childIndexForItem:"), item)
	return rv
}


// Collapses a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/collapseItem(_:)
func (o_ OutlineView) CollapseItem(item objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("collapseItem:"), item)
}


// Collapses a given item and, optionally, its children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/collapseItem(_:collapseChildren:)
func (o_ OutlineView) CollapseItemCollapseChildren(item objc.IObject, collapseChildren bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("collapseItem:collapseChildren:"), item, collapseChildren)
}


// Expands a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/expandItem(_:)
func (o_ OutlineView) ExpandItem(item objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("expandItem:"), item)
}


// Expands a specified item and, optionally, its children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/expandItem(_:expandChildren:)
func (o_ OutlineView) ExpandItemExpandChildren(item objc.IObject, expandChildren bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("expandItem:expandChildren:"), item, expandChildren)
}


// Returns the frame of the outline cell for a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/frameOfOutlineCell(atRow:)
func (o_ OutlineView) FrameOfOutlineCellAtRow(row int) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](o_.ID, objc.Sel("frameOfOutlineCellAtRow:"), row)
	return rv
}


// Inserts new items at the given indexes in the given parent with the specified optional animations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/insertItems(at:inParent:withAnimation:)
func (o_ OutlineView) InsertItemsAtIndexesInParentWithAnimation(indexes foundation.IndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	objc.Send[objc.ID](o_.ID, objc.Sel("insertItemsAtIndexes:inParent:withAnimation:"), indexes, parent, animationOptions)
}


// Returns a Boolean value that indicates whether a given item is expandable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/isExpandable(_:)
func (o_ OutlineView) IsExpandable(item objc.IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isExpandable:"), item)
	return rv
}


// Returns a Boolean value that indicates whether a given item is expanded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/isItemExpanded(_:)
func (o_ OutlineView) IsItemExpanded(item objc.IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isItemExpanded:"), item)
	return rv
}


// Returns the item associated with a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/item(atRow:)
func (o_ OutlineView) ItemAtRow(row int) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("itemAtRow:"), row)
	return rv
}


// Returns the indentation level for a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/level(forItem:)
func (o_ OutlineView) LevelForItem(item objc.IObject) int {
	rv := objc.Send[int](o_.ID, objc.Sel("levelForItem:"), item)
	return rv
}


// Returns the indentation level for a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/level(forRow:)
func (o_ OutlineView) LevelForRow(row int) int {
	rv := objc.Send[int](o_.ID, objc.Sel("levelForRow:"), row)
	return rv
}


// Moves an item at a given index in the given parent to a new index in a new parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/moveItem(at:inParent:to:inParent:)
func (o_ OutlineView) MoveItemAtIndexInParentToIndexInParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("moveItemAtIndex:inParent:toIndex:inParent:"), fromIndex, oldParent, toIndex, newParent)
}


// Returns the number of children for the specified parent item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/numberOfChildren(ofItem:)
func (o_ OutlineView) NumberOfChildrenOfItem(item objc.IObject) int {
	rv := objc.Send[int](o_.ID, objc.Sel("numberOfChildrenOfItem:"), item)
	return rv
}


// Returns the parent for a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/parent(forItem:)
func (o_ OutlineView) ParentForItem(item objc.IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("parentForItem:"), item)
	return rv
}


// Reloads and redisplays the data for the given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/reloadItem(_:)
func (o_ OutlineView) ReloadItem(item objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("reloadItem:"), item)
}


// Reloads a given item and, optionally, its children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/reloadItem(_:reloadChildren:)
func (o_ OutlineView) ReloadItemReloadChildren(item objc.IObject, reloadChildren bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("reloadItem:reloadChildren:"), item, reloadChildren)
}


// Removes items at the given indexes in the given parent with the specified optional animations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/removeItems(at:inParent:withAnimation:)
func (o_ OutlineView) RemoveItemsAtIndexesInParentWithAnimation(indexes foundation.IndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeItemsAtIndexes:inParent:withAnimation:"), indexes, parent, animationOptions)
}


// Returns the row associated with a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/row(forItem:)
func (o_ OutlineView) RowForItem(item objc.IObject) int {
	rv := objc.Send[int](o_.ID, objc.Sel("rowForItem:"), item)
	return rv
}


// Used to “retarget” a proposed drop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/setDropItem(_:dropChildIndex:)
func (o_ OutlineView) SetDropItemDropChildIndex(item objc.IObject, index int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDropItem:dropChildIndex:"), item, index)
}


// Returns a Boolean value that indicates whether auto-expanded items should return to their original collapsed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/shouldCollapseAutoExpandedItems(forDeposited:)
func (o_ OutlineView) ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("shouldCollapseAutoExpandedItemsForDeposited:"), deposited)
	return rv
}


// A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/autoresizesOutlineColumn
func (o_ OutlineView) AutoresizesOutlineColumn() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("autoresizesOutlineColumn"))
	return rv
}


// A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/autoresizesOutlineColumn
func (o_ OutlineView) SetAutoresizesOutlineColumn(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutoresizesOutlineColumn:"), value)
}


// A Boolean value indicating whether the expanded items are automatically saved across launches of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/autosaveExpandedItems
func (o_ OutlineView) AutosaveExpandedItems() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("autosaveExpandedItems"))
	return rv
}


// A Boolean value indicating whether the expanded items are automatically saved across launches of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/autosaveExpandedItems
func (o_ OutlineView) SetAutosaveExpandedItems(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutosaveExpandedItems:"), value)
}


// The object that provides the data displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/dataSource
func (o_ OutlineView) DataSource() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("dataSource"))
	return rv
}


// The object that provides the data displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/dataSource
func (o_ OutlineView) SetDataSource(value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDataSource:"), value)
}


// The outline view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/delegate
func (o_ OutlineView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("delegate"))
	return rv
}


// The outline view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/delegate
func (o_ OutlineView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationMarkerFollowsCell
func (o_ OutlineView) IndentationMarkerFollowsCell() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("indentationMarkerFollowsCell"))
	return rv
}


// A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationMarkerFollowsCell
func (o_ OutlineView) SetIndentationMarkerFollowsCell(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIndentationMarkerFollowsCell:"), value)
}


// The per-level indentation, measured in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationPerLevel
func (o_ OutlineView) IndentationPerLevel() float64 {
	rv := objc.Send[float64](o_.ID, objc.Sel("indentationPerLevel"))
	return rv
}


// The per-level indentation, measured in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationPerLevel
func (o_ OutlineView) SetIndentationPerLevel(value float64) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIndentationPerLevel:"), value)
}


// The table column in which hierarchical data is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/outlineTableColumn
func (o_ OutlineView) OutlineTableColumn() ITableColumn {
	rv := objc.Send[TableColumn](o_.ID, objc.Sel("outlineTableColumn"))
	return rv
}


// The table column in which hierarchical data is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/outlineTableColumn
func (o_ OutlineView) SetOutlineTableColumn(value ITableColumn) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOutlineTableColumn:"), value)
}


// A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/stronglyReferencesItems
func (o_ OutlineView) StronglyReferencesItems() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("stronglyReferencesItems"))
	return rv
}


// A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/stronglyReferencesItems
func (o_ OutlineView) SetStronglyReferencesItems(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setStronglyReferencesItems:"), value)
}


// The user interface layout direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/userInterfaceLayoutDirection
func (o_ OutlineView) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](o_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// The user interface layout direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/userInterfaceLayoutDirection
func (o_ OutlineView) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}



