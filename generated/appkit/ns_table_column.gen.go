// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTableColumn */


/* debug [class_header]: Header for NSTableColumn */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TableColumn */
// An interface definition for the [TableColumn] class.
type ITableColumn interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TableColumn */
	// properties:
	DataCell() objc.ID
	SetDataCell(value objc.ID)
	HeaderCell() ITableHeaderCell
	SetHeaderCell(value ITableHeaderCell)
	HeaderToolTip() objc.IObject /* cross-framework: NSString */
	SetHeaderToolTip(value objc.IObject /* cross-framework: NSString */)
	Identifier() UserInterfaceItemIdentifier /* typedef */
	SetIdentifier(value UserInterfaceItemIdentifier /* typedef */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TableColumn */
	// methods:
	DataCellForRow(row int) objc.ID
	SizeToFit()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TableColumn */
// Alloc allocates a new instance without initialization.
func (tc _TableColumnClass) Alloc() TableColumn {
	rv := objc.Send[TableColumn](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TableColumn */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TableColumn */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/init(coder:)
func NewTableColumnWithCoder(coder foundation.Coder) TableColumn {
	instance := getTableColumnClass().Alloc()
	rv := objc.Send[TableColumn](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTableColumnWithCoder */


// Initializes a newly created table column with a string identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/init(identifier:)
func NewTableColumnWithIdentifier(identifier UserInterfaceItemIdentifier /* typedef */) TableColumn {
	instance := getTableColumnClass().Alloc()
	rv := objc.Send[TableColumn](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTableColumnWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TableColumn */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TableColumn */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TableColumn */

// Returns the cell object used to display values in the specified row of the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell(forRow:)
func (t_ TableColumn) DataCellForRow(row int) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("dataCellForRow:"), row)
	return rv
}/* debug [instance_methods/method]: DataCellForRow */


// Resizes the table column to fit the width of its header cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/sizeToFit()
func (t_ TableColumn) SizeToFit() {
	objc.Send[objc.ID](t_.ID, objc.Sel("sizeToFit"))
}/* debug [instance_methods/method]: SizeToFit */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TableColumn */

// The cell prototype used by the table column to draw individual cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell
func (t_ TableColumn) DataCell() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("dataCell"))
	return rv
}/* debug [instance_properties/getter]: dataCell */


// The cell prototype used by the table column to draw individual cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/dataCell
func (t_ TableColumn) SetDataCell(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDataCell:"), value)
}/* debug [instance_properties/setter]: dataCell */


// The cell used to draw the table column’s header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerCell
func (t_ TableColumn) HeaderCell() ITableHeaderCell {
	rv := objc.Send[TableHeaderCell](t_.ID, objc.Sel("headerCell"))
	return rv
}/* debug [instance_properties/getter]: headerCell */


// The cell used to draw the table column’s header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerCell
func (t_ TableColumn) SetHeaderCell(value ITableHeaderCell) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeaderCell:"), value)
}/* debug [instance_properties/setter]: headerCell */


// The string that’s displayed in a help tag over the table column header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerToolTip
func (t_ TableColumn) HeaderToolTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("headerToolTip"))
	return rv
}/* debug [instance_properties/getter]: headerToolTip */


// The string that’s displayed in a help tag over the table column header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/headerToolTip
func (t_ TableColumn) SetHeaderToolTip(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHeaderToolTip:"), value)
}/* debug [instance_properties/setter]: headerToolTip */


// The identifier string for the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/identifier
func (t_ TableColumn) Identifier() UserInterfaceItemIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The identifier string for the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/identifier
func (t_ TableColumn) SetIdentifier(value UserInterfaceItemIdentifier /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// A Boolean that indicates whether a cell-based table’s column cells are user editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isEditable
func (t_ TableColumn) Editable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("editable"))
	return rv
}/* debug [instance_properties/getter]: editable */


// A Boolean that indicates whether a cell-based table’s column cells are user editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isEditable
func (t_ TableColumn) SetEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditable:"), value)
}/* debug [instance_properties/setter]: editable */


// A Boolean that indicates whether the table column is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isHidden
func (t_ TableColumn) Hidden() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// A Boolean that indicates whether the table column is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/isHidden
func (t_ TableColumn) SetHidden(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHidden:"), value)
}/* debug [instance_properties/setter]: hidden */


// The table column’s maximum width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/maxWidth
func (t_ TableColumn) MaxWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("maxWidth"))
	return rv
}/* debug [instance_properties/getter]: maxWidth */


// The table column’s maximum width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/maxWidth
func (t_ TableColumn) SetMaxWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxWidth:"), value)
}/* debug [instance_properties/setter]: maxWidth */


// The table column’s minimum width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/minWidth
func (t_ TableColumn) MinWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("minWidth"))
	return rv
}/* debug [instance_properties/getter]: minWidth */


// The table column’s minimum width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/minWidth
func (t_ TableColumn) SetMinWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMinWidth:"), value)
}/* debug [instance_properties/setter]: minWidth */


// The table column’s resizing mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/resizingMask
func (t_ TableColumn) ResizingMask() TableColumnResizingOptions {
	rv := objc.Send[TableColumnResizingOptions](t_.ID, objc.Sel("resizingMask"))
	return rv
}/* debug [instance_properties/getter]: resizingMask */


// The table column’s resizing mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/resizingMask
func (t_ TableColumn) SetResizingMask(value TableColumnResizingOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResizingMask:"), value)
}/* debug [instance_properties/setter]: resizingMask */


// The table column’s sort descriptor prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/sortDescriptorPrototype
func (t_ TableColumn) SortDescriptorPrototype() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("sortDescriptorPrototype"))
	return rv
}/* debug [instance_properties/getter]: sortDescriptorPrototype */


// The table column’s sort descriptor prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/sortDescriptorPrototype
func (t_ TableColumn) SetSortDescriptorPrototype(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSortDescriptorPrototype:"), value)
}/* debug [instance_properties/setter]: sortDescriptorPrototype */


// The table view that contains the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/tableView
func (t_ TableColumn) TableView() ITableView {
	rv := objc.Send[TableView](t_.ID, objc.Sel("tableView"))
	return rv
}/* debug [instance_properties/getter]: tableView */


// The table view that contains the table column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/tableView
func (t_ TableColumn) SetTableView(value ITableView) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTableView:"), value)
}/* debug [instance_properties/setter]: tableView */


// The title of the table column’s header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/title
func (t_ TableColumn) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title of the table column’s header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/title
func (t_ TableColumn) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The table column’s width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/width
func (t_ TableColumn) Width() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// The table column’s width, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/width
func (t_ TableColumn) SetWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */


// A Boolean that indicates whether a cell-based table’s column cells are user editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/iseditable
func (t_ TableColumn) IsEditable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// A Boolean that indicates whether a cell-based table’s column cells are user editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/iseditable
func (t_ TableColumn) SetIsEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */


// A Boolean that indicates whether the table column is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/ishidden
func (t_ TableColumn) IsHidden() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// A Boolean that indicates whether the table column is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/ishidden
func (t_ TableColumn) SetIsHidden(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTableColumn */


