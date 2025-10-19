// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TableColumn] class.
var (
	tableColumnClass     _TableColumnClass
	tableColumnClassOnce sync.Once
)

func getTableColumnClass() _TableColumnClass {
	tableColumnClassOnce.Do(func() {
		tableColumnClass = _TableColumnClass{objc.GetClass("NSTableColumn")}
	})
	return tableColumnClass
}

type _TableColumnClass struct {
	class objc.Class
}

// An interface definition for the [TableColumn] class.
type ITableColumn interface {
	objectivec.IObject
}

// The display characteristics and identifier for a column in a table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn
type TableColumn struct {
	objectivec.Object
}

// TableColumnFrom constructs a [TableColumn] from an unsafe.Pointer.
//
// The display characteristics and identifier for a column in a table view.
func TableColumnFrom(ptr unsafe.Pointer) TableColumn {
	return TableColumn{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TableColumnClass) Alloc() TableColumn {
	rv := objc.Send[TableColumn](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TableColumnClass) New() TableColumn {
	rv := objc.Send[TableColumn](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableColumn) Init() TableColumn {
	rv := objc.Send[TableColumn](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableColumn) Autorelease() TableColumn {
	rv := objc.Send[TableColumn](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableColumn creates a new TableColumn instance.
func NewTableColumn() TableColumn {
	return getTableColumnClass().New()
}




