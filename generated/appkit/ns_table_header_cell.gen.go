// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TableHeaderCell] class.
var (
	TableHeaderCellClass     _TableHeaderCellClass
	TableHeaderCellClassOnce sync.Once
)

func getTableHeaderCellClass() _TableHeaderCellClass {
	TableHeaderCellClassOnce.Do(func() {
		TableHeaderCellClass = _TableHeaderCellClass{objc.GetClass("NSTableHeaderCell")}
	})
	return TableHeaderCellClass
}

type _TableHeaderCellClass struct {
	class objc.Class
}

// An interface definition for the [TableHeaderCell] class.
type ITableHeaderCell interface {
	ITextFieldCell
}

// An object that a table header view uses to draw the content of the column headers.
//
// Subclasses of the class can override the , , and methods to change the way headers appear. This specific subclass is responsible for drawing the sort indicators. See the class specification for information on overriding these methods. See the and for more information.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell
type TableHeaderCell struct {
	TextFieldCell
}

// TableHeaderCellFrom constructs a [TableHeaderCell] from an unsafe.Pointer.
//
// An object that a table header view uses to draw the content of the column headers.
func TableHeaderCellFrom(ptr unsafe.Pointer) TableHeaderCell {
	return TableHeaderCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TableHeaderCellClass) Alloc() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TableHeaderCellClass) New() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableHeaderCell) Init() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableHeaderCell) Autorelease() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableHeaderCell creates a new TableHeaderCell instance.
func NewTableHeaderCell() TableHeaderCell {
	return getTableHeaderCellClass().New()
}
