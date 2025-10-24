// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSOutlineView */


/* debug [class_header]: Header for NSOutlineView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OutlineView */
// An interface definition for the [OutlineView] class.
type IOutlineView interface {
	ITableView
	
/* debug [class_interface_properties]: Properties for OutlineView */
	// properties:
	AutoresizesOutlineColumn() bool
	SetAutoresizesOutlineColumn(value bool)
	AutosaveExpandedItems() bool
	SetAutosaveExpandedItems(value bool)
	DataSource() unsafe.Pointer
	SetDataSource(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OutlineView */
	// methods:
	ChildOfItem(index int, item objc.IObject) objc.ID
	ChildIndexForItem(item objc.IObject) int
	CollapseItem(item objc.IObject)
	CollapseItemCollapseChildren(item objc.IObject, collapseChildren bool)
	ExpandItem(item objc.IObject)
	ExpandItemExpandChildren(item objc.IObject, expandChildren bool)
	FrameOfOutlineCellAtRow(row int) Rect /* not a class type */
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
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OutlineView */
// Alloc allocates a new instance without initialization.
func (oc _OutlineViewClass) Alloc() OutlineView {
	rv := objc.Send[OutlineView](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OutlineView */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OutlineView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OutlineView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OutlineView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OutlineView */

// Returns the specified child of an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/child(_:ofItem:)
func (o_ OutlineView) ChildOfItem(index int, item objc.IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("child:ofItem:"), index, item)
	return rv
}/* debug [instance_methods/method]: ChildOfItem */


// Returns the child index of the specified item within its parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/childIndex(forItem:)
func (o_ OutlineView) ChildIndexForItem(item objc.IObject) int {
	rv := objc.Send[int](o_.ID, objc.Sel("childIndexForItem:"), item)
	return rv
}/* debug [instance_methods/method]: ChildIndexForItem */


// Collapses a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/collapseItem(_:)
func (o_ OutlineView) CollapseItem(item objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("collapseItem:"), item)
}/* debug [instance_methods/method]: CollapseItem */


// Collapses a given item and, optionally, its children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/collapseItem(_:collapseChildren:)
func (o_ OutlineView) CollapseItemCollapseChildren(item objc.IObject, collapseChildren bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("collapseItem:collapseChildren:"), item, collapseChildren)
}/* debug [instance_methods/method]: CollapseItemCollapseChildren */


// Expands a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/expandItem(_:)
func (o_ OutlineView) ExpandItem(item objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("expandItem:"), item)
}/* debug [instance_methods/method]: ExpandItem */


// Expands a specified item and, optionally, its children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/expandItem(_:expandChildren:)
func (o_ OutlineView) ExpandItemExpandChildren(item objc.IObject, expandChildren bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("expandItem:expandChildren:"), item, expandChildren)
}/* debug [instance_methods/method]: ExpandItemExpandChildren */


// Returns the frame of the outline cell for a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/frameOfOutlineCell(atRow:)
func (o_ OutlineView) FrameOfOutlineCellAtRow(row int) Rect /* not a class type */ {
	rv := objc.Send[Rect](o_.ID, objc.Sel("frameOfOutlineCellAtRow:"), row)
	return rv
}/* debug [instance_methods/method]: FrameOfOutlineCellAtRow */


// Inserts new items at the given indexes in the given parent with the specified optional animations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/insertItems(at:inParent:withAnimation:)
func (o_ OutlineView) InsertItemsAtIndexesInParentWithAnimation(indexes foundation.IndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	objc.Send[objc.ID](o_.ID, objc.Sel("insertItemsAtIndexes:inParent:withAnimation:"), indexes, parent, animationOptions)
}/* debug [instance_methods/method]: InsertItemsAtIndexesInParentWithAnimation */


// Returns a Boolean value that indicates whether a given item is expandable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/isExpandable(_:)
func (o_ OutlineView) IsExpandable(item objc.IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isExpandable:"), item)
	return rv
}/* debug [instance_methods/method]: IsExpandable */


// Returns a Boolean value that indicates whether a given item is expanded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/isItemExpanded(_:)
func (o_ OutlineView) IsItemExpanded(item objc.IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isItemExpanded:"), item)
	return rv
}/* debug [instance_methods/method]: IsItemExpanded */


// Returns the item associated with a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/item(atRow:)
func (o_ OutlineView) ItemAtRow(row int) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("itemAtRow:"), row)
	return rv
}/* debug [instance_methods/method]: ItemAtRow */


// Returns the indentation level for a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/level(forItem:)
func (o_ OutlineView) LevelForItem(item objc.IObject) int {
	rv := objc.Send[int](o_.ID, objc.Sel("levelForItem:"), item)
	return rv
}/* debug [instance_methods/method]: LevelForItem */


// Returns the indentation level for a given row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/level(forRow:)
func (o_ OutlineView) LevelForRow(row int) int {
	rv := objc.Send[int](o_.ID, objc.Sel("levelForRow:"), row)
	return rv
}/* debug [instance_methods/method]: LevelForRow */


// Moves an item at a given index in the given parent to a new index in a new parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/moveItem(at:inParent:to:inParent:)
func (o_ OutlineView) MoveItemAtIndexInParentToIndexInParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("moveItemAtIndex:inParent:toIndex:inParent:"), fromIndex, oldParent, toIndex, newParent)
}/* debug [instance_methods/method]: MoveItemAtIndexInParentToIndexInParent */


// Returns the number of children for the specified parent item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/numberOfChildren(ofItem:)
func (o_ OutlineView) NumberOfChildrenOfItem(item objc.IObject) int {
	rv := objc.Send[int](o_.ID, objc.Sel("numberOfChildrenOfItem:"), item)
	return rv
}/* debug [instance_methods/method]: NumberOfChildrenOfItem */


// Returns the parent for a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/parent(forItem:)
func (o_ OutlineView) ParentForItem(item objc.IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("parentForItem:"), item)
	return rv
}/* debug [instance_methods/method]: ParentForItem */


// Reloads and redisplays the data for the given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/reloadItem(_:)
func (o_ OutlineView) ReloadItem(item objc.IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("reloadItem:"), item)
}/* debug [instance_methods/method]: ReloadItem */


// Reloads a given item and, optionally, its children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/reloadItem(_:reloadChildren:)
func (o_ OutlineView) ReloadItemReloadChildren(item objc.IObject, reloadChildren bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("reloadItem:reloadChildren:"), item, reloadChildren)
}/* debug [instance_methods/method]: ReloadItemReloadChildren */


// Removes items at the given indexes in the given parent with the specified optional animations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/removeItems(at:inParent:withAnimation:)
func (o_ OutlineView) RemoveItemsAtIndexesInParentWithAnimation(indexes foundation.IndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeItemsAtIndexes:inParent:withAnimation:"), indexes, parent, animationOptions)
}/* debug [instance_methods/method]: RemoveItemsAtIndexesInParentWithAnimation */


// Returns the row associated with a given item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/row(forItem:)
func (o_ OutlineView) RowForItem(item objc.IObject) int {
	rv := objc.Send[int](o_.ID, objc.Sel("rowForItem:"), item)
	return rv
}/* debug [instance_methods/method]: RowForItem */


// Used to “retarget” a proposed drop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/setDropItem(_:dropChildIndex:)
func (o_ OutlineView) SetDropItemDropChildIndex(item objc.IObject, index int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDropItem:dropChildIndex:"), item, index)
}/* debug [instance_methods/method]: SetDropItemDropChildIndex */


// Returns a Boolean value that indicates whether auto-expanded items should return to their original collapsed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/shouldCollapseAutoExpandedItems(forDeposited:)
func (o_ OutlineView) ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("shouldCollapseAutoExpandedItemsForDeposited:"), deposited)
	return rv
}/* debug [instance_methods/method]: ShouldCollapseAutoExpandedItemsForDeposited */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OutlineView */

// A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/autoresizesOutlineColumn
func (o_ OutlineView) AutoresizesOutlineColumn() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("autoresizesOutlineColumn"))
	return rv
}/* debug [instance_properties/getter]: autoresizesOutlineColumn */


// A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/autoresizesOutlineColumn
func (o_ OutlineView) SetAutoresizesOutlineColumn(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutoresizesOutlineColumn:"), value)
}/* debug [instance_properties/setter]: autoresizesOutlineColumn */


// A Boolean value indicating whether the expanded items are automatically saved across launches of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/autosaveExpandedItems
func (o_ OutlineView) AutosaveExpandedItems() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("autosaveExpandedItems"))
	return rv
}/* debug [instance_properties/getter]: autosaveExpandedItems */


// A Boolean value indicating whether the expanded items are automatically saved across launches of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/autosaveExpandedItems
func (o_ OutlineView) SetAutosaveExpandedItems(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutosaveExpandedItems:"), value)
}/* debug [instance_properties/setter]: autosaveExpandedItems */


// The object that provides the data displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/dataSource
func (o_ OutlineView) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// The object that provides the data displayed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/dataSource
func (o_ OutlineView) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// The outline view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/delegate
func (o_ OutlineView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The outline view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/delegate
func (o_ OutlineView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationMarkerFollowsCell
func (o_ OutlineView) IndentationMarkerFollowsCell() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("indentationMarkerFollowsCell"))
	return rv
}/* debug [instance_properties/getter]: indentationMarkerFollowsCell */


// A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationMarkerFollowsCell
func (o_ OutlineView) SetIndentationMarkerFollowsCell(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIndentationMarkerFollowsCell:"), value)
}/* debug [instance_properties/setter]: indentationMarkerFollowsCell */


// The per-level indentation, measured in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationPerLevel
func (o_ OutlineView) IndentationPerLevel() float64 {
	rv := objc.Send[float64](o_.ID, objc.Sel("indentationPerLevel"))
	return rv
}/* debug [instance_properties/getter]: indentationPerLevel */


// The per-level indentation, measured in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationPerLevel
func (o_ OutlineView) SetIndentationPerLevel(value float64) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIndentationPerLevel:"), value)
}/* debug [instance_properties/setter]: indentationPerLevel */


// The table column in which hierarchical data is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/outlineTableColumn
func (o_ OutlineView) OutlineTableColumn() ITableColumn {
	rv := objc.Send[TableColumn](o_.ID, objc.Sel("outlineTableColumn"))
	return rv
}/* debug [instance_properties/getter]: outlineTableColumn */


// The table column in which hierarchical data is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/outlineTableColumn
func (o_ OutlineView) SetOutlineTableColumn(value ITableColumn) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOutlineTableColumn:"), value)
}/* debug [instance_properties/setter]: outlineTableColumn */


// A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/stronglyReferencesItems
func (o_ OutlineView) StronglyReferencesItems() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("stronglyReferencesItems"))
	return rv
}/* debug [instance_properties/getter]: stronglyReferencesItems */


// A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/stronglyReferencesItems
func (o_ OutlineView) SetStronglyReferencesItems(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setStronglyReferencesItems:"), value)
}/* debug [instance_properties/setter]: stronglyReferencesItems */


// The user interface layout direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/userInterfaceLayoutDirection
func (o_ OutlineView) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](o_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}/* debug [instance_properties/getter]: userInterfaceLayoutDirection */


// The user interface layout direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/userInterfaceLayoutDirection
func (o_ OutlineView) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}/* debug [instance_properties/setter]: userInterfaceLayoutDirection */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSOutlineView */



