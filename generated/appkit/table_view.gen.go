
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TableView] class.
var TableViewClass _TableViewClass

func init() {
	TableViewClass = _TableViewClass{objc.GetClass("NSTableView")}
}

type _TableViewClass struct {
	objc.Class
}

// An interface definition for the [TableView] class.
type ITableView interface {
	ID() objc.ID
	RectOfRow(row int) unsafe.Pointer
	RowAtPoint(point unsafe.Pointer) int
}

type TableView struct {
	id objc.ID
}

func TableViewFrom(ptr unsafe.Pointer) TableView {
	return TableView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TableView) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TableViewClass) Alloc() TableView {
	rv := objc.Send[TableView](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TableViewClass) New() TableView {
	rv := objc.Send[TableView](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTableView creates and returns a new initialized instance.
func NewTableView() TableView {
	return TableViewClass.New()
}

// Init initializes the instance.
func (t_ TableView) Init() TableView {
	rv := objc.Send[TableView](t_.ID(), selInit)
	return rv
}
// Returns the rectangle containing the row at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableView/rect(ofRow:)
func (t_ TableView) RectOfRow(row int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("rectOfRow:"), row)
	return rv
}
// Returns the index of the row the specified point lies in. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTableView/row(at:)
func (t_ TableView) RowAtPoint(point unsafe.Pointer) int {
	rv := objc.Send[int](t_.ID(), objc.RegisterName("rowAtPoint:"), point)
	return rv
}
