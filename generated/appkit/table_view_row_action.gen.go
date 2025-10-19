// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TableViewRowAction] class.
var (
	tableViewRowActionClass     _TableViewRowActionClass
	tableViewRowActionClassOnce sync.Once
)

func getTableViewRowActionClass() _TableViewRowActionClass {
	tableViewRowActionClassOnce.Do(func() {
		tableViewRowActionClass = _TableViewRowActionClass{objc.GetClass("NSTableViewRowAction")}
	})
	return tableViewRowActionClass
}

type _TableViewRowActionClass struct {
	class objc.Class
}

// An interface definition for the [TableViewRowAction] class.
type ITableViewRowAction interface {
	objectivec.IObject
}

// A single action to present when the user swipes horizontally on a table row.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction
type TableViewRowAction struct {
	objectivec.Object
}

// TableViewRowActionFrom constructs a [TableViewRowAction] from an unsafe.Pointer.
//
// A single action to present when the user swipes horizontally on a table row.
func TableViewRowActionFrom(ptr unsafe.Pointer) TableViewRowAction {
	return TableViewRowAction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TableViewRowActionClass) Alloc() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TableViewRowActionClass) New() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableViewRowAction) Init() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableViewRowAction) Autorelease() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableViewRowAction creates a new TableViewRowAction instance.
func NewTableViewRowAction() TableViewRowAction {
	return getTableViewRowActionClass().New()
}




