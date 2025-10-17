
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableHeaderCell] class.
var TableHeaderCellClass _TableHeaderCellClass

func init() {
	TableHeaderCellClass = _TableHeaderCellClass{objc.GetClass("NSTableHeaderCell")}
}

type _TableHeaderCellClass struct {
	objc.Class
}

// An interface definition for the [TableHeaderCell] class.
type ITableHeaderCell interface {
	ID() objc.ID
}

type TableHeaderCell struct {
	id objc.ID
}

func TableHeaderCellFrom(ptr unsafe.Pointer) TableHeaderCell {
	return TableHeaderCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TableHeaderCell) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TableHeaderCellClass) Alloc() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TableHeaderCellClass) New() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTableHeaderCell creates and returns a new initialized instance.
func NewTableHeaderCell() TableHeaderCell {
	return TableHeaderCellClass.New()
}

// Init initializes the instance.
func (t_ TableHeaderCell) Init() TableHeaderCell {
	rv := objc.Send[TableHeaderCell](t_.ID(), selInit)
	return rv
}
