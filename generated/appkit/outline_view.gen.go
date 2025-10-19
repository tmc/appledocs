// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OutlineView] class.
var (
	outlineViewClass     _OutlineViewClass
	outlineViewClassOnce sync.Once
)

func getOutlineViewClass() _OutlineViewClass {
	outlineViewClassOnce.Do(func() {
		outlineViewClass = _OutlineViewClass{objc.GetClass("NSOutlineView")}
	})
	return outlineViewClass
}

type _OutlineViewClass struct {
	class objc.Class
}

// An interface definition for the [OutlineView] class.
type IOutlineView interface {
	ITableView
}

// A view that uses a row-and-column format to display hierarchical data like directories and files that can be expanded and collapsed. [Full Topic]
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




