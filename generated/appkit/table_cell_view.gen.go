// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TableCellView] class.
var (
	tableCellViewClass     _TableCellViewClass
	tableCellViewClassOnce sync.Once
)

func getTableCellViewClass() _TableCellViewClass {
	tableCellViewClassOnce.Do(func() {
		tableCellViewClass = _TableCellViewClass{objc.GetClass("NSTableCellView")}
	})
	return tableCellViewClass
}

type _TableCellViewClass struct {
	class objc.Class
}

// An interface definition for the [TableCellView] class.
type ITableCellView interface {
	IView
}

// A reusable container view shown for a particular cell in a table view that uses rows for content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableCellView
type TableCellView struct {
	View
}

// TableCellViewFrom constructs a [TableCellView] from an unsafe.Pointer.
//
// A reusable container view shown for a particular cell in a table view that uses rows for content.
func TableCellViewFrom(ptr unsafe.Pointer) TableCellView {
	return TableCellView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TableCellViewClass) Alloc() TableCellView {
	rv := objc.Send[TableCellView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TableCellViewClass) New() TableCellView {
	rv := objc.Send[TableCellView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableCellView) Init() TableCellView {
	rv := objc.Send[TableCellView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableCellView) Autorelease() TableCellView {
	rv := objc.Send[TableCellView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableCellView creates a new TableCellView instance.
func NewTableCellView() TableCellView {
	return getTableCellViewClass().New()
}




