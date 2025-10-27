// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewMatrix

// ExampleMatrix_AddColumn demonstrates using AddColumn on a Matrix instance.
// Adds a new column of cells to the right of the last column.
func ExampleMatrix_AddColumn() {
	obj := appkit.NewMatrix()
	obj.AddColumn()
	// Output:
	}

// ExampleMatrix_AddRow demonstrates using AddRow on a Matrix instance.
// Adds a new row of cells below the last row.
func ExampleMatrix_AddRow() {
	obj := appkit.NewMatrix()
	obj.AddRow()
	// Output:
	}

// ExampleMatrix_DeselectAllCells demonstrates using DeselectAllCells on a Matrix instance.
// Deselects all cells in the receiver and, if necessary, redisplays the receiver.
func ExampleMatrix_DeselectAllCells() {
	obj := appkit.NewMatrix()
	obj.DeselectAllCells()
	// Output:
	}

// ExampleMatrix_DeselectSelectedCell demonstrates using DeselectSelectedCell on a Matrix instance.
// Deselects the selected cell or cells.
func ExampleMatrix_DeselectSelectedCell() {
	obj := appkit.NewMatrix()
	obj.DeselectSelectedCell()
	// Output:
	}

// ExampleMatrix_ResetCursorRects demonstrates using ResetCursorRects on a Matrix instance.
// Resets cursor rectangles so the cursor becomes an I-beam over text cells.
func ExampleMatrix_ResetCursorRects() {
	obj := appkit.NewMatrix()
	obj.ResetCursorRects()
	// Output:
	}

// ExampleMatrix_SendAction demonstrates using SendAction on a Matrix instance.
// If the selected cell has both an action and a target, sends its action to its target.
func ExampleMatrix_SendAction() {
	obj := appkit.NewMatrix()
	_ = obj.SendAction()
	// Output:
	}

// ExampleMatrix_SendDoubleAction demonstrates using SendDoubleAction on a Matrix instance.
// Sends the double-click action message to the target of the receiver.
func ExampleMatrix_SendDoubleAction() {
	obj := appkit.NewMatrix()
	obj.SendDoubleAction()
	// Output:
	}

// ExampleMatrix_SizeToCells demonstrates using SizeToCells on a Matrix instance.
// Changes the width and the height of the receiver’s frame so it exactly contains the cells.
func ExampleMatrix_SizeToCells() {
	obj := appkit.NewMatrix()
	obj.SizeToCells()
	// Output:
	}

