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


