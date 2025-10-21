// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A view that uses a row-and-column format to display hierarchical data like directories and files that can be expanded and collapsed.
//
// Like a table view, an outline view does not store its own data, instead it retrieves data values as needed from a data source to which it has a weak reference (see ). See , which declares the methods that an object uses to access the contents of its data source object. An outline view has the following features: A user can expand and collapse rows, edit values, and resize and rearrange columns. Each item in the outline view must be unique. In order for the collapsed state to remain consistent between reloads the item’s pointer must remain the same and the item must maintain sameness. The view gets data from a data source (see ). The view retrieves only the data that needs to be displayed. For more information about using NSOutlineView in your app, see .
//
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


// The per-level indentation, measured in points.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationPerLevel
func (o_ OutlineView) IndentationPerLevel() float64 {
	rv := objc.Send[float64](o_.ID, objc.Sel("indentationPerLevel"))
	return rv
}


// SetIndentationPerLevel sets the value of the indentationPerLevel property.
// The per-level indentation, measured in points.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationPerLevel
func (o_ OutlineView) SetIndentationPerLevel(value float64) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIndentationPerLevel:"), value)
}

// A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/autoresizesoutlinecolumn
func (o_ OutlineView) AutoresizesOutlineColumn() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("autoresizesOutlineColumn"))
	return rv
}


// SetAutoresizesOutlineColumn sets the value of the autoresizesOutlineColumn property.
// A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/autoresizesoutlinecolumn
func (o_ OutlineView) SetAutoresizesOutlineColumn(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutoresizesOutlineColumn:"), value)
}

// A Boolean value indicating whether the expanded items are automatically saved across launches of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/autosaveexpandeditems
func (o_ OutlineView) AutosaveExpandedItems() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("autosaveExpandedItems"))
	return rv
}


// SetAutosaveExpandedItems sets the value of the autosaveExpandedItems property.
// A Boolean value indicating whether the expanded items are automatically saved across launches of the app.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/autosaveexpandeditems
func (o_ OutlineView) SetAutosaveExpandedItems(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutosaveExpandedItems:"), value)
}

// The object that provides the data displayed by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/datasource
func (o_ OutlineView) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("dataSource"))
	return rv
}


// SetDataSource sets the value of the dataSource property.
// The object that provides the data displayed by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/datasource
func (o_ OutlineView) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDataSource:"), value)
}

// The outline view’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/delegate
func (o_ OutlineView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The outline view’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/delegate
func (o_ OutlineView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/indentationmarkerfollowscell
func (o_ OutlineView) IndentationMarkerFollowsCell() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("indentationMarkerFollowsCell"))
	return rv
}


// SetIndentationMarkerFollowsCell sets the value of the indentationMarkerFollowsCell property.
// A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/indentationmarkerfollowscell
func (o_ OutlineView) SetIndentationMarkerFollowsCell(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIndentationMarkerFollowsCell:"), value)
}

// The table column in which hierarchical data is displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/outlinetablecolumn
func (o_ OutlineView) OutlineTableColumn() NSTableColumn {
	rv := objc.Send[NSTableColumn](o_.ID, objc.Sel("outlineTableColumn"))
	return rv
}


// SetOutlineTableColumn sets the value of the outlineTableColumn property.
// The table column in which hierarchical data is displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/outlinetablecolumn
func (o_ OutlineView) SetOutlineTableColumn(value ITableColumn) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setOutlineTableColumn:"), value)
}

// A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/stronglyreferencesitems
func (o_ OutlineView) StronglyReferencesItems() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("stronglyReferencesItems"))
	return rv
}


// SetStronglyReferencesItems sets the value of the stronglyReferencesItems property.
// A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/stronglyreferencesitems
func (o_ OutlineView) SetStronglyReferencesItems(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setStronglyReferencesItems:"), value)
}

// The user interface layout direction.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/userinterfacelayoutdirection
func (o_ OutlineView) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](o_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// SetUserInterfaceLayoutDirection sets the value of the userInterfaceLayoutDirection property.
// The user interface layout direction.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/userinterfacelayoutdirection
func (o_ OutlineView) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}



