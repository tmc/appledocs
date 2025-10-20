// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewCell

// ExampleNewCell demonstrates how to create a Cell instance.
func ExampleNewCell() {
	_ = appkit.NewCell()
	// Output:
}

// ExampleNewCellTextCell demonstrates how to create a Cell instance using NewCellTextCell.
// Returns an NSCell object initialized with the specified string and set to have the cell’s default menu.
func ExampleNewCellTextCell() {
	_ = appkit.NewCellTextCell(
		"string", // string string
	)
	// Output:
}
