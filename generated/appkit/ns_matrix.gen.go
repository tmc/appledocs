// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [Matrix] class.
type IMatrix interface {
	IControl
	// properties:
	AutosizesCells() bool /* primitive/slice/pointer. */
	SetAutosizesCells(value bool /* primitive/slice/pointer. */)
	AllowsEmptySelection() bool /* primitive/slice/pointer. */
	SetAllowsEmptySelection(value bool /* primitive/slice/pointer. */)
	AutorecalculatesCellSize() bool /* primitive/slice/pointer. */
	SetAutorecalculatesCellSize(value bool /* primitive/slice/pointer. */)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	CellBackgroundColor() IColor
	SetCellBackgroundColor(value IColor)
	CellClass() unsafe.Pointer
	SetCellClass(value unsafe.Pointer)
	CellSize() objc.IObject /* cross-framework: Size */
	SetCellSize(value objc.IObject /* cross-framework: Size */)
	Cells() ICell
	SetCells(value ICell)
	Delegate() MatrixDelegate /* not a class type */
	SetDelegate(value MatrixDelegate /* not a class type */)
	DoubleAction() unsafe.Pointer
	SetDoubleAction(value unsafe.Pointer)
	DrawsBackground() bool /* primitive/slice/pointer. */
	SetDrawsBackground(value bool /* primitive/slice/pointer. */)
	DrawsCellBackground() bool /* primitive/slice/pointer. */
	SetDrawsCellBackground(value bool /* primitive/slice/pointer. */)
	IntercellSpacing() objc.IObject /* cross-framework: Size */
	SetIntercellSpacing(value objc.IObject /* cross-framework: Size */)
	IsAutoscroll() bool /* primitive/slice/pointer. */
	SetIsAutoscroll(value bool /* primitive/slice/pointer. */)
	IsSelectionByRect() bool /* primitive/slice/pointer. */
	SetIsSelectionByRect(value bool /* primitive/slice/pointer. */)
	KeyCell() ICell
	SetKeyCell(value ICell)
	Mode() unsafe.Pointer
	SetMode(value unsafe.Pointer)
	MouseDownFlags() int /* primitive/slice/pointer. */
	SetMouseDownFlags(value int /* primitive/slice/pointer. */)
	NumberOfColumns() int /* primitive/slice/pointer. */
	SetNumberOfColumns(value int /* primitive/slice/pointer. */)
	NumberOfRows() int /* primitive/slice/pointer. */
	SetNumberOfRows(value int /* primitive/slice/pointer. */)
	Prototype() ICell
	SetPrototype(value ICell)
	SelectedCells() ICell
	SetSelectedCells(value ICell)
	SelectedColumn() int /* primitive/slice/pointer. */
	SetSelectedColumn(value int /* primitive/slice/pointer. */)
	SelectedRow() int /* primitive/slice/pointer. */
	SetSelectedRow(value int /* primitive/slice/pointer. */)
	TabKeyTraversesCells() bool /* primitive/slice/pointer. */
	SetTabKeyTraversesCells(value bool /* primitive/slice/pointer. */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MatrixClass) Alloc() Matrix {
	rv := objc.Send[Matrix](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean that indicates whether the cell sizes change when the receiver is resized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/autosizesCells
func (m_ Matrix) AutosizesCells() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("autosizesCells"))
	return rv
}


// A Boolean that indicates whether the cell sizes change when the receiver is resized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/autosizesCells
func (m_ Matrix) SetAutosizesCells(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutosizesCells:"), value)
}


// A Boolean that indicates whether a radio-mode matrix supports an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/allowsemptyselection
func (m_ Matrix) AllowsEmptySelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}


// A Boolean that indicates whether a radio-mode matrix supports an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/allowsemptyselection
func (m_ Matrix) SetAllowsEmptySelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsEmptySelection:"), value)
}


// A Boolean that indicates whether the matrix auto-recalculates its cell size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/autorecalculatescellsize
func (m_ Matrix) AutorecalculatesCellSize() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("autorecalculatesCellSize"))
	return rv
}


// A Boolean that indicates whether the matrix auto-recalculates its cell size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/autorecalculatescellsize
func (m_ Matrix) SetAutorecalculatesCellSize(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutorecalculatesCellSize:"), value)
}


// The background color of the matrix (the space between the cells).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/backgroundcolor
func (m_ Matrix) BackgroundColor() IColor {
	rv := objc.Send[Color](m_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the matrix (the space between the cells).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/backgroundcolor
func (m_ Matrix) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The background color of the matrix’s cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/cellbackgroundcolor
func (m_ Matrix) CellBackgroundColor() IColor {
	rv := objc.Send[Color](m_.ID, objc.Sel("cellBackgroundColor"))
	return rv
}


// The background color of the matrix’s cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/cellbackgroundcolor
func (m_ Matrix) SetCellBackgroundColor(value IColor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCellBackgroundColor:"), value)
}


// The subclass of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/cellclass
func (m_ Matrix) CellClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cellClass"))
	return rv
}


// The subclass of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/cellclass
func (m_ Matrix) SetCellClass(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCellClass:"), value)
}


// The size of each cell in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/cellsize
func (m_ Matrix) CellSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](m_.ID, objc.Sel("cellSize"))
	return rv
}


// The size of each cell in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/cellsize
func (m_ Matrix) SetCellSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCellSize:"), value)
}


// An array containing the cells of the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/cells
func (m_ Matrix) Cells() ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("cells"))
	return rv
}


// An array containing the cells of the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/cells
func (m_ Matrix) SetCells(value ICell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCells:"), value)
}


// The delegate for messages from the field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/delegate
func (m_ Matrix) Delegate() MatrixDelegate /* not a class type */ {
	rv := objc.Send[MatrixDelegate](m_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for messages from the field editor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/delegate
func (m_ Matrix) SetDelegate(value MatrixDelegate /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}


// The action sent to the target of the receiver when the user double-clicks a cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/doubleaction
func (m_ Matrix) DoubleAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("doubleAction"))
	return rv
}


// The action sent to the target of the receiver when the user double-clicks a cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/doubleaction
func (m_ Matrix) SetDoubleAction(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDoubleAction:"), value)
}


// A Boolean that indicates whether the matrix draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/drawsbackground
func (m_ Matrix) DrawsBackground() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean that indicates whether the matrix draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/drawsbackground
func (m_ Matrix) SetDrawsBackground(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawsBackground:"), value)
}


// A Boolean that indicates whether the matrix draws the background within each of its cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/drawscellbackground
func (m_ Matrix) DrawsCellBackground() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("drawsCellBackground"))
	return rv
}


// A Boolean that indicates whether the matrix draws the background within each of its cells.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/drawscellbackground
func (m_ Matrix) SetDrawsCellBackground(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawsCellBackground:"), value)
}


// The vertical and horizontal spacing between cells in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/intercellspacing
func (m_ Matrix) IntercellSpacing() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](m_.ID, objc.Sel("intercellSpacing"))
	return rv
}


// The vertical and horizontal spacing between cells in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/intercellspacing
func (m_ Matrix) SetIntercellSpacing(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntercellSpacing:"), value)
}


// A Boolean that indicates whether the receiver is automatically scrolled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/isautoscroll
func (m_ Matrix) IsAutoscroll() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isAutoscroll"))
	return rv
}


// A Boolean that indicates whether the receiver is automatically scrolled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/isautoscroll
func (m_ Matrix) SetIsAutoscroll(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsAutoscroll:"), value)
}


// A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/isselectionbyrect
func (m_ Matrix) IsSelectionByRect() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSelectionByRect"))
	return rv
}


// A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/isselectionbyrect
func (m_ Matrix) SetIsSelectionByRect(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSelectionByRect:"), value)
}


// The cell that will be clicked when the user presses the Space bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/keycell
func (m_ Matrix) KeyCell() ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("keyCell"))
	return rv
}


// The cell that will be clicked when the user presses the Space bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/keycell
func (m_ Matrix) SetKeyCell(value ICell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyCell:"), value)
}


// The selection mode of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/mode-swift.property
func (m_ Matrix) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mode"))
	return rv
}


// The selection mode of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/mode-swift.property
func (m_ Matrix) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}


// The flags in effect at the mouse-down event that started the current tracking session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/mousedownflags
func (m_ Matrix) MouseDownFlags() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("mouseDownFlags"))
	return rv
}


// The flags in effect at the mouse-down event that started the current tracking session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/mousedownflags
func (m_ Matrix) SetMouseDownFlags(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMouseDownFlags:"), value)
}


// The number of columns in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/numberofcolumns
func (m_ Matrix) NumberOfColumns() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfColumns"))
	return rv
}


// The number of columns in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/numberofcolumns
func (m_ Matrix) SetNumberOfColumns(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfColumns:"), value)
}


// The number of rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/numberofrows
func (m_ Matrix) NumberOfRows() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("numberOfRows"))
	return rv
}


// The number of rows in the matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/numberofrows
func (m_ Matrix) SetNumberOfRows(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfRows:"), value)
}


// The prototype cell that’s copied whenever the matrix creates a new cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/prototype
func (m_ Matrix) Prototype() ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("prototype"))
	return rv
}


// The prototype cell that’s copied whenever the matrix creates a new cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/prototype
func (m_ Matrix) SetPrototype(value ICell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrototype:"), value)
}


// An array containing all of the matrix’s highlighted cells plus its selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/selectedcells
func (m_ Matrix) SelectedCells() ICell {
	rv := objc.Send[Cell](m_.ID, objc.Sel("selectedCells"))
	return rv
}


// An array containing all of the matrix’s highlighted cells plus its selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/selectedcells
func (m_ Matrix) SetSelectedCells(value ICell) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedCells:"), value)
}


// The column number of the selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/selectedcolumn
func (m_ Matrix) SelectedColumn() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("selectedColumn"))
	return rv
}


// The column number of the selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/selectedcolumn
func (m_ Matrix) SetSelectedColumn(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedColumn:"), value)
}


// The row number of the selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/selectedrow
func (m_ Matrix) SelectedRow() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("selectedRow"))
	return rv
}


// The row number of the selected cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/selectedrow
func (m_ Matrix) SetSelectedRow(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedRow:"), value)
}


// A Boolean that indicates whether pressing the Tab key advances the key cell to the next selectable cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/tabkeytraversescells
func (m_ Matrix) TabKeyTraversesCells() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("tabKeyTraversesCells"))
	return rv
}


// A Boolean that indicates whether pressing the Tab key advances the key cell to the next selectable cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/tabkeytraversescells
func (m_ Matrix) SetTabKeyTraversesCells(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTabKeyTraversesCells:"), value)
}



