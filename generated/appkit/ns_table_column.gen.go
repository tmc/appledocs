// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TableColumn] class.
var (
	TableColumnClass     _TableColumnClass
	TableColumnClassOnce sync.Once
)

func getTableColumnClass() _TableColumnClass {
	TableColumnClassOnce.Do(func() {
		TableColumnClass = _TableColumnClass{objc.GetClass("NSTableColumn")}
	})
	return TableColumnClass
}

type _TableColumnClass struct {
	class objc.Class
}

// An interface definition for the [TableColumn] class.
type ITableColumn interface {
	objectivec.IObject
	// properties:
	DataCell() objc.ID
	SetDataCell(value objc.ID)
	HeaderCell() ITableHeaderCell
	SetHeaderCell(value ITableHeaderCell)
	HeaderToolTip() objc.IObject /* cross-framework: NSString */
	SetHeaderToolTip(value objc.IObject /* cross-framework: NSString */)
	Identifier() objc.IObject /* cross-framework: UserInterfaceItemIdentifier */
	SetIdentifier(value objc.IObject /* cross-framework: UserInterfaceItemIdentifier */)
	Editable() bool
	SetEditable(value bool)
	Hidden() bool
	SetHidden(value bool)
	MaxWidth() float64
	SetMaxWidth(value float64)
	MinWidth() float64
	SetMinWidth(value float64)
	ResizingMask() TableColumnResizingOptions
	SetResizingMask(value TableColumnResizingOptions)
	SortDescriptorPrototype() objectivec.IObject
	SetSortDescriptorPrototype(value objectivec.IObject)
	TableView() ITableView
	SetTableView(value ITableView)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	Width() float64
	SetWidth(value float64)
	IsEditable() bool
	SetIsEditable(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	// methods:
	DataCellForRow(row int) objc.ID
	SizeToFit()
}

// The display characteristics and identifier for a column in a table view.
//
// A table column object determines the width (including the maximum and minimum widths) of its column in the table view and specifies the column’s resizing and editing behavior. A table column stores two cell objects: the header cell, which is used to draw the column header, and the data cell, which is used to draw the values for each row. In a cell-based table, you can control the display of the column by specifying subclasses of to use and by setting the font and other display characteristics for these cells. For example, you can use an to display string values or substitute an to display pictures.


// The display characteristics and identifier for a column in a table view.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/init(coder:)
func NewTableColumnWithCoder(coder foundation.Coder) TableColumn {
	instance := getTableColumnClass().Alloc()
	rv := objc.Send[TableColumn](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes a newly created table column with a string identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/init(identifier:)
func NewTableColumnWithIdentifier(identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */) TableColumn {
	instance := getTableColumnClass().Alloc()
	rv := objc.Send[TableColumn](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}



// Returns the cell object used to display values in the specified row of the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell(forRow:)
func (t_ TableColumn) DataCellForRow(row int) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("dataCellForRow:"), row)
	return rv
}


// Resizes the table column to fit the width of its header cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/sizeToFit()
func (t_ TableColumn) SizeToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeToFit"))
}


// The cell prototype used by the table column to draw individual cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell
func (t_ TableColumn) DataCell() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("dataCell"))
	return rv
}


// The cell prototype used by the table column to draw individual cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell
func (t_ TableColumn) SetDataCell(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDataCell:"), value)
}


// The cell used to draw the table column’s header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerCell
func (t_ TableColumn) HeaderCell() ITableHeaderCell {
	rv := objc.Send[TableHeaderCell](t_.ID, objc.Sel("headerCell"))
	return rv
}


// The cell used to draw the table column’s header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerCell
func (t_ TableColumn) SetHeaderCell(value ITableHeaderCell) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeaderCell:"), value)
}


// The string that’s displayed in a help tag over the table column header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerToolTip
func (t_ TableColumn) HeaderToolTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("headerToolTip"))
	return rv
}


// The string that’s displayed in a help tag over the table column header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerToolTip
func (t_ TableColumn) SetHeaderToolTip(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeaderToolTip:"), value)
}


// The identifier string for the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/identifier
func (t_ TableColumn) Identifier() objc.IObject /* cross-framework: UserInterfaceItemIdentifier */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("identifier"))
	return rv
}


// The identifier string for the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/identifier
func (t_ TableColumn) SetIdentifier(value objc.IObject /* cross-framework: UserInterfaceItemIdentifier */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIdentifier:"), value)
}


// A Boolean that indicates whether a cell-based table’s column cells are user editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isEditable
func (t_ TableColumn) Editable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("editable"))
	return rv
}


// A Boolean that indicates whether a cell-based table’s column cells are user editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isEditable
func (t_ TableColumn) SetEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditable:"), value)
}


// A Boolean that indicates whether the table column is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isHidden
func (t_ TableColumn) Hidden() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hidden"))
	return rv
}


// A Boolean that indicates whether the table column is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isHidden
func (t_ TableColumn) SetHidden(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHidden:"), value)
}


// The table column’s maximum width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/maxWidth
func (t_ TableColumn) MaxWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("maxWidth"))
	return rv
}


// The table column’s maximum width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/maxWidth
func (t_ TableColumn) SetMaxWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxWidth:"), value)
}


// The table column’s minimum width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/minWidth
func (t_ TableColumn) MinWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("minWidth"))
	return rv
}


// The table column’s minimum width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/minWidth
func (t_ TableColumn) SetMinWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinWidth:"), value)
}


// The table column’s resizing mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/resizingMask
func (t_ TableColumn) ResizingMask() TableColumnResizingOptions {
	rv := objc.Send[TableColumnResizingOptions](t_.ID, objc.Sel("resizingMask"))
	return rv
}


// The table column’s resizing mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/resizingMask
func (t_ TableColumn) SetResizingMask(value TableColumnResizingOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResizingMask:"), value)
}


// The table column’s sort descriptor prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/sortDescriptorPrototype
func (t_ TableColumn) SortDescriptorPrototype() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("sortDescriptorPrototype"))
	return rv
}


// The table column’s sort descriptor prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/sortDescriptorPrototype
func (t_ TableColumn) SetSortDescriptorPrototype(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSortDescriptorPrototype:"), value)
}


// The table view that contains the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/tableView
func (t_ TableColumn) TableView() ITableView {
	rv := objc.Send[TableView](t_.ID, objc.Sel("tableView"))
	return rv
}


// The table view that contains the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/tableView
func (t_ TableColumn) SetTableView(value ITableView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTableView:"), value)
}


// The title of the table column’s header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/title
func (t_ TableColumn) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("title"))
	return rv
}


// The title of the table column’s header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/title
func (t_ TableColumn) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTitle:"), value)
}


// The table column’s width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/width
func (t_ TableColumn) Width() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("width"))
	return rv
}


// The table column’s width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/width
func (t_ TableColumn) SetWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:"), value)
}


// A Boolean that indicates whether a cell-based table’s column cells are user editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/iseditable
func (t_ TableColumn) IsEditable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean that indicates whether a cell-based table’s column cells are user editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/iseditable
func (t_ TableColumn) SetIsEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}


// A Boolean that indicates whether the table column is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/ishidden
func (t_ TableColumn) IsHidden() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isHidden"))
	return rv
}


// A Boolean that indicates whether the table column is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/ishidden
func (t_ TableColumn) SetIsHidden(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsHidden:"), value)
}


