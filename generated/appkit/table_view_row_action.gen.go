
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableViewRowAction] class.
var TableViewRowActionClass _TableViewRowActionClass

func init() {
	TableViewRowActionClass = _TableViewRowActionClass{objc.GetClass("NSTableViewRowAction")}
}

type _TableViewRowActionClass struct {
	objc.Class
}

// An interface definition for the [TableViewRowAction] class.
type ITableViewRowAction interface {
	ID() objc.ID
}

type TableViewRowAction struct {
	id objc.ID
}

func TableViewRowActionFrom(ptr unsafe.Pointer) TableViewRowAction {
	return TableViewRowAction{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TableViewRowAction) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TableViewRowActionClass) Alloc() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TableViewRowActionClass) New() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTableViewRowAction creates and returns a new initialized instance.
func NewTableViewRowAction() TableViewRowAction {
	return TableViewRowActionClass.New()
}

// Init initializes the instance.
func (t_ TableViewRowAction) Init() TableViewRowAction {
	rv := objc.Send[TableViewRowAction](t_.ID(), selInit)
	return rv
}
