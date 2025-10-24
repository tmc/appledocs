// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSMatrix */


/* debug [class_header]: Header for NSMatrix */
// The class instance for the [Matrix] class.
var (
	MatrixClass     _MatrixClass
	MatrixClassOnce sync.Once
)

func getMatrixClass() _MatrixClass {
	MatrixClassOnce.Do(func() {
		MatrixClass = _MatrixClass{objc.GetClass("NSMatrix")}
	})
	return MatrixClass
}

type _MatrixClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Matrix */
// An interface definition for the [Matrix] class.
type IMatrix interface {
	IControl
	
/* debug [class_interface_properties]: Properties for Matrix */
	// properties:
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	AutorecalculatesCellSize() bool
	SetAutorecalculatesCellSize(value bool)
	AutosizesCells() bool
	SetAutosizesCells(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	CellBackgroundColor() IColor
	SetCellBackgroundColor(value IColor)
	CellClass() objc.Class
	SetCellClass(value objc.Class)
	CellSize() Size /* not a class type */
	SetCellSize(value Size /* not a class type */)
	Cells() []Cell
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	DrawsCellBackground() bool
	SetDrawsCellBackground(value bool)
	IntercellSpacing() Size /* not a class type */
	SetIntercellSpacing(value Size /* not a class type */)
	Autoscroll() bool
	SetAutoscroll(value bool)
	SelectionByRect() bool
	SetSelectionByRect(value bool)
	KeyCell() ICell
	SetKeyCell(value ICell)
	Mode() MatrixMode
	SetMode(value MatrixMode)
	MouseDownFlags() int
	NumberOfColumns() int
	NumberOfRows() int
	Prototype() ICell
	SetPrototype(value ICell)
	SelectedCell() ICell
	SelectedCells() []Cell
	SelectedColumn() int
	SelectedRow() int
	TabKeyTraversesCells() bool
	SetTabKeyTraversesCells(value bool)
	IsAutoscroll() bool
	SetIsAutoscroll(value bool)
	IsSelectionByRect() bool
	SetIsSelectionByRect(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Matrix */
	// methods:
	AcceptsFirstMouse(event IEvent) bool
	AddColumn()
	AddColumnWithCells(newCells []Cell)
	AddRow()
	AddRowWithCells(newCells []Cell)
	CellAtRowColumn(row int, col int) ICell
	CellWithTag(tag int) ICell
	CellFrameAtRowColumn(row int, col int) Rect /* not a class type */
	DeselectAllCells()
	DeselectSelectedCell()
	DrawCellAtRowColumn(row int, col int)
	GetNumberOfRowsColumns(rowCount int, colCount int)
	GetRowColumnForPoint(row int, col int, point vision.Point) bool
	GetRowColumnOfCell(row int, col int, cell ICell) bool
	HighlightCellAtRowColumn(flag bool, row int, col int)
	InsertColumn(column int)
	InsertColumnWithCells(column int, newCells []Cell)
	InsertRow(row int)
	InsertRowWithCells(row int, newCells []Cell)
	MakeCellAtRowColumn(row int, col int) ICell
	MouseDown(event IEvent)
	PerformKeyEquivalent(event IEvent) bool
	PutCellAtRowColumn(newCell ICell, row int, col int)
	RemoveColumn(col int)
	RemoveRow(row int)
	RenewRowsColumns(newRows int, newCols int)
	ResetCursorRects()
	ScrollCellToVisibleAtRowColumn(row int, col int)
	SelectAll(sender objc.IObject)
	SelectCellAtRowColumn(row int, col int)
	SelectCellWithTag(tag int) bool
	SelectText(sender objc.IObject)
	SelectTextAtRowColumn(row int, col int) ICell
	SendAction() bool
	SendActionToForAllCells(selector objc.SEL, object objc.IObject, flag bool)
	SendDoubleAction()
	SetScrollable(flag bool)
	SetSelectionFromToAnchorHighlight(startPos int, endPos int, anchorPos int, lit bool)
	SetStateAtRowColumn(value int, row int, col int)
	SetToolTipForCell(toolTipString objc.IObject /* cross-framework: NSString */, cell ICell)
	SetValidateSize(flag bool)
	SizeToCells()
	SortUsingSelector(comparator objc.SEL)
	SortUsingFunctionContext(compare objectivec.IObject, context objectivec.IObject)
	TextDidBeginEditing(notification foundation.Notification)
	TextDidChange(notification foundation.Notification)
	TextDidEndEditing(notification foundation.Notification)
	TextShouldBeginEditing(textObject IText) bool
	TextShouldEndEditing(textObject IText) bool
	ToolTipForCell(cell ICell) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Matrix */
// Alloc allocates a new instance without initialization.
func (mc _MatrixClass) Alloc() Matrix {
	rv := objc.Send[Matrix](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixClass) New() Matrix {
	rv := objc.Send[Matrix](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Matrix) Init() Matrix {
	rv := objc.Send[Matrix](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Matrix) Autorelease() Matrix {
	rv := objc.Send[Matrix](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrix creates a new Matrix instance.
func NewMatrix() Matrix {
	return getMatrixClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Matrix */
// A legacy interface for grouping radio buttons or other types of cells together.
//
// uses flipped coordinates by default. The cells in an object are numbered by row and column, each starting with 0; for example, the top left would be at (0, 0), and the that’s second down and third across would be at (1, 2). The class has the notion of a single selected cell, which is the cell that was most recently clicked or that was so designated by a or message. The selected cell is the cell chosen for action messages except for ( ), which is assigned to the key cell. (The key cell is generally identical to the selected cell, but can be given click focus while leaving the selected cell unchanged.) If the user has selected multiple cells, the selected cell is the one lowest and furthest to the right in the matrix of cells.


// A legacy interface for grouping radio buttons or other types of cells together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix
type Matrix struct {
	Control
}

// MatrixFrom constructs a [Matrix] from an unsafe.Pointer.
//
// A legacy interface for grouping radio buttons or other types of cells together.
func MatrixFrom(ptr unsafe.Pointer) Matrix {
	return Matrix{
		Control: ControlFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Matrix */

// Initializes a newly allocated matrix with the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/init(frame:)
func NewMatrixWithFrame(frameRect Rect /* not a class type */) Matrix {
	instance := getMatrixClass().Alloc()
	rv := objc.Send[Matrix](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixWithFrame */


// Initializes and returns a newly allocated matrix of the specified size using cells of the given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/init(frame:mode:cellClass:numberOfRows:numberOfColumns:)
func NewMatrixWithFrameModeCellClassNumberOfRowsNumberOfColumns(frameRect Rect /* not a class type */, mode MatrixMode, factoryId objc.Class, rowsHigh int, colsWide int) Matrix {
	instance := getMatrixClass().Alloc()
	rv := objc.Send[Matrix](instance.ID, objc.Sel("initWithFrame:mode:cellClass:numberOfRows:numberOfColumns:"), frameRect, mode, factoryId, rowsHigh, colsWide)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixWithFrameModeCellClassNumberOfRowsNumberOfColumns */


// Initializes and returns a newly allocated matrix of the specified size using the given cell as a prototype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/init(frame:mode:prototype:numberOfRows:numberOfColumns:)
func NewMatrixWithFrameModePrototypeNumberOfRowsNumberOfColumns(frameRect Rect /* not a class type */, mode MatrixMode, cell ICell, rowsHigh int, colsWide int) Matrix {
	instance := getMatrixClass().Alloc()
	rv := objc.Send[Matrix](instance.ID, objc.Sel("initWithFrame:mode:prototype:numberOfRows:numberOfColumns:"), frameRect, mode, cell, rowsHigh, colsWide)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixWithFrameModePrototypeNumberOfRowsNumberOfColumns */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Matrix */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Matrix */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Matrix */

// Returns a Boolean value indicating whether the receiver accepts the first mouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/acceptsFirstMouse(for:)
func (m_ Matrix) AcceptsFirstMouse(event IEvent) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("acceptsFirstMouse:"), event)
	return rv
}/* debug [instance_methods/method]: AcceptsFirstMouse */


// Adds a new column of cells to the right of the last column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/addColumn()
func (m_ Matrix) AddColumn() {
	objc.Send[objc.ID](m_.ID, objc.Sel("addColumn"))
}/* debug [instance_methods/method]: AddColumn */


// Adds a new column of cells to the right of the last column, using the given cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/addColumn(with:)
func (m_ Matrix) AddColumnWithCells(newCells []Cell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addColumnWithCells:"), newCells)
}/* debug [instance_methods/method]: AddColumnWithCells */


// Adds a new row of cells below the last row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/addRow()
func (m_ Matrix) AddRow() {
	objc.Send[objc.ID](m_.ID, objc.Sel("addRow"))
}/* debug [instance_methods/method]: AddRow */


// Adds a new row of cells below the last row, using the specified cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/addRow(with:)
func (m_ Matrix) AddRowWithCells(newCells []Cell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addRowWithCells:"), newCells)
}/* debug [instance_methods/method]: AddRowWithCells */


// Returns the cell at the specified row and column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cell(atRow:column:)
func (m_ Matrix) CellAtRowColumn(row int, col int) ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("cellAtRow:column:"), row, col)
	return rv
}/* debug [instance_methods/method]: CellAtRowColumn */


// Searches the receiver and returns the last cell matching the specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cell(withTag:)
func (m_ Matrix) CellWithTag(tag int) ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("cellWithTag:"), tag)
	return rv
}/* debug [instance_methods/method]: CellWithTag */


// Returns the frame rectangle of the cell that would be drawn at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cellFrame(atRow:column:)
func (m_ Matrix) CellFrameAtRowColumn(row int, col int) Rect /* not a class type */ {
	rv := objc.Send[Rect](m_.ID, objc.Sel("cellFrameAtRow:column:"), row, col)
	return rv
}/* debug [instance_methods/method]: CellFrameAtRowColumn */


// Deselects all cells in the receiver and, if necessary, redisplays the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/deselectAllCells()
func (m_ Matrix) DeselectAllCells() {
	objc.Send[objc.ID](m_.ID, objc.Sel("deselectAllCells"))
}/* debug [instance_methods/method]: DeselectAllCells */


// Deselects the selected cell or cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/deselectSelectedCell()
func (m_ Matrix) DeselectSelectedCell() {
	objc.Send[objc.ID](m_.ID, objc.Sel("deselectSelectedCell"))
}/* debug [instance_methods/method]: DeselectSelectedCell */


// Displays the cell at the specified row and column.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/drawCell(atRow:column:)
func (m_ Matrix) DrawCellAtRowColumn(row int, col int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("drawCellAtRow:column:"), row, col)
}/* debug [instance_methods/method]: DrawCellAtRowColumn */


// Obtains the number of rows and columns in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/getNumberOfRows(_:columns:)
func (m_ Matrix) GetNumberOfRowsColumns(rowCount int, colCount int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getNumberOfRows:columns:"), rowCount, colCount)
}/* debug [instance_methods/method]: GetNumberOfRowsColumns */


// Indicates whether the specified point lies within one of the cells of the matrix and returns the location of the cell within which the point lies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/getRow(_:column:for:)
func (m_ Matrix) GetRowColumnForPoint(row int, col int, point vision.Point) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("getRow:column:forPoint:"), row, col, point)
	return rv
}/* debug [instance_methods/method]: GetRowColumnForPoint */


// Searches the receiver for the specified cell and returns the row and column of the cell
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/getRow(_:column:of:)
func (m_ Matrix) GetRowColumnOfCell(row int, col int, cell ICell) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("getRow:column:ofCell:"), row, col, cell)
	return rv
}/* debug [instance_methods/method]: GetRowColumnOfCell */


// Highlights or unhighlights the cell at the specified row and column location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/highlightCell(_:atRow:column:)
func (m_ Matrix) HighlightCellAtRowColumn(flag bool, row int, col int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("highlightCell:atRow:column:"), flag, row, col)
}/* debug [instance_methods/method]: HighlightCellAtRowColumn */


// Inserts a new column of cells at the specified location. .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/insertColumn(_:)
func (m_ Matrix) InsertColumn(column int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertColumn:"), column)
}/* debug [instance_methods/method]: InsertColumn */


// Inserts a new column of cells before the specified column, using the given cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/insertColumn(_:with:)
func (m_ Matrix) InsertColumnWithCells(column int, newCells []Cell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertColumn:withCells:"), column, newCells)
}/* debug [instance_methods/method]: InsertColumnWithCells */


// Inserts a new row of cells before the specified row.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/insertRow(_:)
func (m_ Matrix) InsertRow(row int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertRow:"), row)
}/* debug [instance_methods/method]: InsertRow */


// Inserts a new row of cells before the specified row, using the given cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/insertRow(_:with:)
func (m_ Matrix) InsertRowWithCells(row int, newCells []Cell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertRow:withCells:"), row, newCells)
}/* debug [instance_methods/method]: InsertRowWithCells */


// Creates a new cell at the location specified by the given row and column in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/makeCell(atRow:column:)
func (m_ Matrix) MakeCellAtRowColumn(row int, col int) ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("makeCellAtRow:column:"), row, col)
	return rv
}/* debug [instance_methods/method]: MakeCellAtRowColumn */


// Responds to a mouse-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/mouseDown(with:)
func (m_ Matrix) MouseDown(event IEvent) {
	objc.Send[objc.ID](m_.ID, objc.Sel("mouseDown:"), event)
}/* debug [instance_methods/method]: MouseDown */


// Looks for a cell that has the given key equivalent and, if found, makes that cell respond as if clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/performKeyEquivalent(with:)
func (m_ Matrix) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("performKeyEquivalent:"), event)
	return rv
}/* debug [instance_methods/method]: PerformKeyEquivalent */


// Replaces the cell at the specified row and column with the new cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/putCell(_:atRow:column:)
func (m_ Matrix) PutCellAtRowColumn(newCell ICell, row int, col int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("putCell:atRow:column:"), newCell, row, col)
}/* debug [instance_methods/method]: PutCellAtRowColumn */


// Removes the specified column at from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/removeColumn(_:)
func (m_ Matrix) RemoveColumn(col int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeColumn:"), col)
}/* debug [instance_methods/method]: RemoveColumn */


// Removes the specified row from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/removeRow(_:)
func (m_ Matrix) RemoveRow(row int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeRow:"), row)
}/* debug [instance_methods/method]: RemoveRow */


// Changes the number of rows and columns in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/renewRows(_:columns:)
func (m_ Matrix) RenewRowsColumns(newRows int, newCols int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("renewRows:columns:"), newRows, newCols)
}/* debug [instance_methods/method]: RenewRowsColumns */


// Resets cursor rectangles so the cursor becomes an I-beam over text cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/resetCursorRects()
func (m_ Matrix) ResetCursorRects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("resetCursorRects"))
}/* debug [instance_methods/method]: ResetCursorRects */


// Scrolls the receiver so the specified cell is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/scrollCellToVisible(atRow:column:)
func (m_ Matrix) ScrollCellToVisibleAtRowColumn(row int, col int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("scrollCellToVisibleAtRow:column:"), row, col)
}/* debug [instance_methods/method]: ScrollCellToVisibleAtRowColumn */


// Selects and highlights all cells in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectAll(_:)
func (m_ Matrix) SelectAll(sender objc.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("selectAll:"), sender)
}/* debug [instance_methods/method]: SelectAll */


// Selects the cell at the specified row and column within the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectCell(atRow:column:)
func (m_ Matrix) SelectCellAtRowColumn(row int, col int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("selectCellAtRow:column:"), row, col)
}/* debug [instance_methods/method]: SelectCellAtRowColumn */


// Selects the last cell with the given tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectCell(withTag:)
func (m_ Matrix) SelectCellWithTag(tag int) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("selectCellWithTag:"), tag)
	return rv
}/* debug [instance_methods/method]: SelectCellWithTag */


// Selects text in the currently selected cell or in the key cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectText(_:)
func (m_ Matrix) SelectText(sender objc.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("selectText:"), sender)
}/* debug [instance_methods/method]: SelectText */


// Selects the text in the cell at the specified location and returns the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectText(atRow:column:)
func (m_ Matrix) SelectTextAtRowColumn(row int, col int) ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("selectTextAtRow:column:"), row, col)
	return rv
}/* debug [instance_methods/method]: SelectTextAtRowColumn */


// If the selected cell has both an action and a target, sends its action to its target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/sendAction()
func (m_ Matrix) SendAction() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("sendAction"))
	return rv
}/* debug [instance_methods/method]: SendAction */


// Iterates through the cells in the receiver, sending the specified selector to an object for each cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/sendAction(_:to:forAllCells:)
func (m_ Matrix) SendActionToForAllCells(selector objc.SEL, object objc.IObject, flag bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sendAction:to:forAllCells:"), selector, object, flag)
}/* debug [instance_methods/method]: SendActionToForAllCells */


// Sends the double-click action message to the target of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/sendDoubleAction()
func (m_ Matrix) SendDoubleAction() {
	objc.Send[objc.ID](m_.ID, objc.Sel("sendDoubleAction"))
}/* debug [instance_methods/method]: SendDoubleAction */


// Specifies whether the cells in the matrix are scrollable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/setScrollable(_:)
func (m_ Matrix) SetScrollable(flag bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScrollable:"), flag)
}/* debug [instance_methods/method]: SetScrollable */


// Programmatically selects a range of cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/setSelectionFrom(_:to:anchor:highlight:)
func (m_ Matrix) SetSelectionFromToAnchorHighlight(startPos int, endPos int, anchorPos int, lit bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectionFrom:to:anchor:highlight:"), startPos, endPos, anchorPos, lit)
}/* debug [instance_methods/method]: SetSelectionFromToAnchorHighlight */


// Sets the state of the cell at specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/setState(_:atRow:column:)
func (m_ Matrix) SetStateAtRowColumn(value int, row int, col int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:atRow:column:"), value, row, col)
}/* debug [instance_methods/method]: SetStateAtRowColumn */


// Sets the tooltip for the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/setToolTip(_:for:)
func (m_ Matrix) SetToolTipForCell(toolTipString objc.IObject /* cross-framework: NSString */, cell ICell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setToolTip:forCell:"), toolTipString, cell)
}/* debug [instance_methods/method]: SetToolTipForCell */


// Specifies whether the receiver’s size information is validated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/setValidateSize(_:)
func (m_ Matrix) SetValidateSize(flag bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidateSize:"), flag)
}/* debug [instance_methods/method]: SetValidateSize */


// Changes the width and the height of the receiver’s frame so it exactly contains the cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/sizeToCells()
func (m_ Matrix) SizeToCells() {
	objc.Send[objc.ID](m_.ID, objc.Sel("sizeToCells"))
}/* debug [instance_methods/method]: SizeToCells */


// Sorts the receiver’s cells in ascending order as defined by the comparison method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/sort(using:)
func (m_ Matrix) SortUsingSelector(comparator objc.SEL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingSelector:"), comparator)
}/* debug [instance_methods/method]: SortUsingSelector */


// Sorts the receiver’s cells in ascending order as defined by the specified comparison function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/sort(using:context:)
func (m_ Matrix) SortUsingFunctionContext(compare objectivec.IObject, context objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("sortUsingFunction:context:"), compare, context)
}/* debug [instance_methods/method]: SortUsingFunctionContext */


// Invoked when there’s a change in the text after the receiver gains first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/textDidBeginEditing(_:)
func (m_ Matrix) TextDidBeginEditing(notification foundation.Notification) {
	objc.Send[objc.ID](m_.ID, objc.Sel("textDidBeginEditing:"), notification)
}/* debug [instance_methods/method]: TextDidBeginEditing */


// Invoked when a key-down event or paste operation occurs that changes the receiver’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/textDidChange(_:)
func (m_ Matrix) TextDidChange(notification foundation.Notification) {
	objc.Send[objc.ID](m_.ID, objc.Sel("textDidChange:"), notification)
}/* debug [instance_methods/method]: TextDidChange */


// Invoked when text editing ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/textDidEndEditing(_:)
func (m_ Matrix) TextDidEndEditing(notification foundation.Notification) {
	objc.Send[objc.ID](m_.ID, objc.Sel("textDidEndEditing:"), notification)
}/* debug [instance_methods/method]: TextDidEndEditing */


// Requests permission to begin editing text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/textShouldBeginEditing(_:)
func (m_ Matrix) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("textShouldBeginEditing:"), textObject)
	return rv
}/* debug [instance_methods/method]: TextShouldBeginEditing */


// Requests permission to end editing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/textShouldEndEditing(_:)
func (m_ Matrix) TextShouldEndEditing(textObject IText) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("textShouldEndEditing:"), textObject)
	return rv
}/* debug [instance_methods/method]: TextShouldEndEditing */


// Returns the tooltip for the specified cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/toolTip(for:)
func (m_ Matrix) ToolTipForCell(cell ICell) foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("toolTipForCell:"), cell)
	return rv
}/* debug [instance_methods/method]: ToolTipForCell */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Matrix */

// A Boolean that indicates whether a radio-mode matrix supports an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/allowsEmptySelection
func (m_ Matrix) AllowsEmptySelection() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}/* debug [instance_properties/getter]: allowsEmptySelection */


// A Boolean that indicates whether a radio-mode matrix supports an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/allowsEmptySelection
func (m_ Matrix) SetAllowsEmptySelection(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsEmptySelection:"), value)
}/* debug [instance_properties/setter]: allowsEmptySelection */


// A Boolean that indicates whether the matrix auto-recalculates its cell size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/autorecalculatesCellSize
func (m_ Matrix) AutorecalculatesCellSize() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("autorecalculatesCellSize"))
	return rv
}/* debug [instance_properties/getter]: autorecalculatesCellSize */


// A Boolean that indicates whether the matrix auto-recalculates its cell size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/autorecalculatesCellSize
func (m_ Matrix) SetAutorecalculatesCellSize(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutorecalculatesCellSize:"), value)
}/* debug [instance_properties/setter]: autorecalculatesCellSize */


// A Boolean that indicates whether the cell sizes change when the receiver is resized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/autosizesCells
func (m_ Matrix) AutosizesCells() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("autosizesCells"))
	return rv
}/* debug [instance_properties/getter]: autosizesCells */


// A Boolean that indicates whether the cell sizes change when the receiver is resized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/autosizesCells
func (m_ Matrix) SetAutosizesCells(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutosizesCells:"), value)
}/* debug [instance_properties/setter]: autosizesCells */


// The background color of the matrix (the space between the cells).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/backgroundColor
func (m_ Matrix) BackgroundColor() IColor {
	rv := objc.Send[Color](m_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color of the matrix (the space between the cells).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/backgroundColor
func (m_ Matrix) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The background color of the matrix’s cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cellBackgroundColor
func (m_ Matrix) CellBackgroundColor() IColor {
	rv := objc.Send[Color](m_.ID, objc.Sel("cellBackgroundColor"))
	return rv
}/* debug [instance_properties/getter]: cellBackgroundColor */


// The background color of the matrix’s cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cellBackgroundColor
func (m_ Matrix) SetCellBackgroundColor(value IColor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCellBackgroundColor:"), value)
}/* debug [instance_properties/setter]: cellBackgroundColor */


// The subclass of that the matrix uses when creating new (empty) cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cellClass
func (m_ Matrix) CellClass() objc.Class {
	rv := objc.Send[objc.Class](m_.ID, objc.Sel("cellClass"))
	return rv
}/* debug [instance_properties/getter]: cellClass */


// The subclass of that the matrix uses when creating new (empty) cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cellClass
func (m_ Matrix) SetCellClass(value objc.Class) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCellClass:"), value)
}/* debug [instance_properties/setter]: cellClass */


// The size of each cell in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cellSize
func (m_ Matrix) CellSize() Size /* not a class type */ {
	rv := objc.Send[Size](m_.ID, objc.Sel("cellSize"))
	return rv
}/* debug [instance_properties/getter]: cellSize */


// The size of each cell in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cellSize
func (m_ Matrix) SetCellSize(value Size /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCellSize:"), value)
}/* debug [instance_properties/setter]: cellSize */


// An array containing the cells of the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/cells
func (m_ Matrix) Cells() []Cell {
	rv := objc.Send[[]Cell](m_.ID, objc.Sel("cells"))
	return rv
}/* debug [instance_properties/getter]: cells */


// The delegate for messages from the field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/delegate
func (m_ Matrix) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for messages from the field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/delegate
func (m_ Matrix) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The action sent to the target of the receiver when the user double-clicks a cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/doubleAction
func (m_ Matrix) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](m_.ID, objc.Sel("doubleAction"))
	return rv
}/* debug [instance_properties/getter]: doubleAction */


// The action sent to the target of the receiver when the user double-clicks a cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/doubleAction
func (m_ Matrix) SetDoubleAction(value objc.SEL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDoubleAction:"), value)
}/* debug [instance_properties/setter]: doubleAction */


// A Boolean that indicates whether the matrix draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/drawsBackground
func (m_ Matrix) DrawsBackground() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("drawsBackground"))
	return rv
}/* debug [instance_properties/getter]: drawsBackground */


// A Boolean that indicates whether the matrix draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/drawsBackground
func (m_ Matrix) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawsBackground:"), value)
}/* debug [instance_properties/setter]: drawsBackground */


// A Boolean that indicates whether the matrix draws the background within each of its cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/drawsCellBackground
func (m_ Matrix) DrawsCellBackground() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("drawsCellBackground"))
	return rv
}/* debug [instance_properties/getter]: drawsCellBackground */


// A Boolean that indicates whether the matrix draws the background within each of its cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/drawsCellBackground
func (m_ Matrix) SetDrawsCellBackground(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawsCellBackground:"), value)
}/* debug [instance_properties/setter]: drawsCellBackground */


// The vertical and horizontal spacing between cells in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/intercellSpacing
func (m_ Matrix) IntercellSpacing() Size /* not a class type */ {
	rv := objc.Send[Size](m_.ID, objc.Sel("intercellSpacing"))
	return rv
}/* debug [instance_properties/getter]: intercellSpacing */


// The vertical and horizontal spacing between cells in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/intercellSpacing
func (m_ Matrix) SetIntercellSpacing(value Size /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntercellSpacing:"), value)
}/* debug [instance_properties/setter]: intercellSpacing */


// A Boolean that indicates whether the receiver is automatically scrolled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/isAutoscroll
func (m_ Matrix) Autoscroll() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("autoscroll"))
	return rv
}/* debug [instance_properties/getter]: autoscroll */


// A Boolean that indicates whether the receiver is automatically scrolled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/isAutoscroll
func (m_ Matrix) SetAutoscroll(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutoscroll:"), value)
}/* debug [instance_properties/setter]: autoscroll */


// A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/isSelectionByRect
func (m_ Matrix) SelectionByRect() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("selectionByRect"))
	return rv
}/* debug [instance_properties/getter]: selectionByRect */


// A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/isSelectionByRect
func (m_ Matrix) SetSelectionByRect(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectionByRect:"), value)
}/* debug [instance_properties/setter]: selectionByRect */


// The cell that will be clicked when the user presses the Space bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/keyCell
func (m_ Matrix) KeyCell() ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("keyCell"))
	return rv
}/* debug [instance_properties/getter]: keyCell */


// The cell that will be clicked when the user presses the Space bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/keyCell
func (m_ Matrix) SetKeyCell(value ICell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyCell:"), value)
}/* debug [instance_properties/setter]: keyCell */


// The selection mode of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/mode-swift.property
func (m_ Matrix) Mode() MatrixMode {
	rv := objc.Send[MatrixMode](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// The selection mode of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/mode-swift.property
func (m_ Matrix) SetMode(value MatrixMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */


// The flags in effect at the mouse-down event that started the current tracking session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/mouseDownFlags
func (m_ Matrix) MouseDownFlags() int {
	rv := objc.Send[int](m_.ID, objc.Sel("mouseDownFlags"))
	return rv
}/* debug [instance_properties/getter]: mouseDownFlags */


// The number of columns in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/numberOfColumns
func (m_ Matrix) NumberOfColumns() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfColumns"))
	return rv
}/* debug [instance_properties/getter]: numberOfColumns */


// The number of rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/numberOfRows
func (m_ Matrix) NumberOfRows() int {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfRows"))
	return rv
}/* debug [instance_properties/getter]: numberOfRows */


// The prototype cell that’s copied whenever the matrix creates a new cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/prototype
func (m_ Matrix) Prototype() ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("prototype"))
	return rv
}/* debug [instance_properties/getter]: prototype */


// The prototype cell that’s copied whenever the matrix creates a new cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/prototype
func (m_ Matrix) SetPrototype(value ICell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrototype:"), value)
}/* debug [instance_properties/setter]: prototype */


// The most recently selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectedCell
func (m_ Matrix) SelectedCell() ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("selectedCell"))
	return rv
}/* debug [instance_properties/getter]: selectedCell */


// An array containing all of the matrix’s highlighted cells plus its selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectedCells
func (m_ Matrix) SelectedCells() []Cell {
	rv := objc.Send[[]Cell](m_.ID, objc.Sel("selectedCells"))
	return rv
}/* debug [instance_properties/getter]: selectedCells */


// The column number of the selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectedColumn
func (m_ Matrix) SelectedColumn() int {
	rv := objc.Send[int](m_.ID, objc.Sel("selectedColumn"))
	return rv
}/* debug [instance_properties/getter]: selectedColumn */


// The row number of the selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectedRow
func (m_ Matrix) SelectedRow() int {
	rv := objc.Send[int](m_.ID, objc.Sel("selectedRow"))
	return rv
}/* debug [instance_properties/getter]: selectedRow */


// A Boolean that indicates whether pressing the Tab key advances the key cell to the next selectable cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/tabKeyTraversesCells
func (m_ Matrix) TabKeyTraversesCells() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("tabKeyTraversesCells"))
	return rv
}/* debug [instance_properties/getter]: tabKeyTraversesCells */


// A Boolean that indicates whether pressing the Tab key advances the key cell to the next selectable cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/tabKeyTraversesCells
func (m_ Matrix) SetTabKeyTraversesCells(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTabKeyTraversesCells:"), value)
}/* debug [instance_properties/setter]: tabKeyTraversesCells */


// A Boolean that indicates whether the receiver is automatically scrolled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/isautoscroll
func (m_ Matrix) IsAutoscroll() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAutoscroll"))
	return rv
}/* debug [instance_properties/getter]: isAutoscroll */


// A Boolean that indicates whether the receiver is automatically scrolled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/isautoscroll
func (m_ Matrix) SetIsAutoscroll(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAutoscroll:"), value)
}/* debug [instance_properties/setter]: isAutoscroll */


// A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/isselectionbyrect
func (m_ Matrix) IsSelectionByRect() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSelectionByRect"))
	return rv
}/* debug [instance_properties/getter]: isSelectionByRect */


// A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/isselectionbyrect
func (m_ Matrix) SetIsSelectionByRect(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSelectionByRect:"), value)
}/* debug [instance_properties/setter]: isSelectionByRect */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMatrix */


