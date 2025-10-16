
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableColumn] class.
var TableColumnClass _TableColumnClass

func init() {
	TableColumnClass = _TableColumnClass{objc.GetClass("NSTableColumn")}
}

type _TableColumnClass struct {
	objc.Class
}

// An interface definition for the [TableColumn] class.
type ITableColumn interface {
	ID() objc.ID
}

type TableColumn struct {
	id objc.ID
}

func TableColumnFrom(ptr unsafe.Pointer) TableColumn {
	return TableColumn{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TableColumn) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TableColumnClass) Alloc() TableColumn {
	rv := objc.Send[TableColumn](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TableColumnClass) New() TableColumn {
	rv := objc.Send[TableColumn](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTableColumn creates and returns a new initialized instance.
func NewTableColumn() TableColumn {
	return TableColumnClass.New()
}

// Init initializes the instance.
func (t_ TableColumn) Init() TableColumn {
	rv := objc.Send[TableColumn](t_.ID(), selInit)
	return rv
}
