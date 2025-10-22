// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextTableBlock] class.
var (
	TextTableBlockClass     _TextTableBlockClass
	TextTableBlockClassOnce sync.Once
)

func getTextTableBlockClass() _TextTableBlockClass {
	TextTableBlockClassOnce.Do(func() {
		TextTableBlockClass = _TextTableBlockClass{objc.GetClass("NSTextTableBlock")}
	})
	return TextTableBlockClass
}

type _TextTableBlockClass struct {
	class objc.Class
}

// An interface definition for the [TextTableBlock] class.
type ITextTableBlock interface {
	ITextBlock
	ColumnSpan() int
	SetColumnSpan(value int)
	RowSpan() int
	SetRowSpan(value int)
	StartingColumn() int
	SetStartingColumn(value int)
	StartingRow() int
	SetStartingRow(value int)
	Table() NSTextTable
	SetTable(value ITextTable)
}

// A text block that appears as a cell in a text table.


// A text block that appears as a cell in a text table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTableBlock

type TextTableBlock struct {
	TextBlock
}

// TextTableBlockFrom constructs a [TextTableBlock] from an unsafe.Pointer.
//
// A text block that appears as a cell in a text table.
func TextTableBlockFrom(ptr unsafe.Pointer) TextTableBlock {
	return TextTableBlock{
		TextBlock: TextBlockFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextTableBlockClass) Alloc() TextTableBlock {
	rv := objc.Send[TextTableBlock](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextTableBlockClass) New() TextTableBlock {
	rv := objc.Send[TextTableBlock](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextTableBlock) Init() TextTableBlock {
	rv := objc.Send[TextTableBlock](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextTableBlock) Autorelease() TextTableBlock {
	rv := objc.Send[TextTableBlock](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextTableBlock creates a new TextTableBlock instance.
func NewTextTableBlock() TextTableBlock {
	return getTextTableBlockClass().New()
}



// Returns the number of table columns spanned by this text table block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/columnspan

func (t_ TextTableBlock) ColumnSpan() int {
	rv := objc.Send[int](t_.ID, objc.Sel("columnSpan"))
	return rv
}


// Returns the number of table columns spanned by this text table block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/columnspan

func (t_ TextTableBlock) SetColumnSpan(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setColumnSpan:"), value)
}


// Returns the number of table rows spanned by this text table block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/rowspan

func (t_ TextTableBlock) RowSpan() int {
	rv := objc.Send[int](t_.ID, objc.Sel("rowSpan"))
	return rv
}


// Returns the number of table rows spanned by this text table block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/rowspan

func (t_ TextTableBlock) SetRowSpan(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRowSpan:"), value)
}


// Returns the table column at which this text table block starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/startingcolumn

func (t_ TextTableBlock) StartingColumn() int {
	rv := objc.Send[int](t_.ID, objc.Sel("startingColumn"))
	return rv
}


// Returns the table column at which this text table block starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/startingcolumn

func (t_ TextTableBlock) SetStartingColumn(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStartingColumn:"), value)
}


// Returns the table row at which this text table block starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/startingrow

func (t_ TextTableBlock) StartingRow() int {
	rv := objc.Send[int](t_.ID, objc.Sel("startingRow"))
	return rv
}


// Returns the table row at which this text table block starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/startingrow

func (t_ TextTableBlock) SetStartingRow(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStartingRow:"), value)
}


// Returns the table containing this text table block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/table

func (t_ TextTableBlock) Table() NSTextTable {
	rv := objc.Send[NSTextTable](t_.ID, objc.Sel("table"))
	return rv
}


// Returns the table containing this text table block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstexttableblock/table

func (t_ TextTableBlock) SetTable(value ITextTable) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTable:"), value)
}



