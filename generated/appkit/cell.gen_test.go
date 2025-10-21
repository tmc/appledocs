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
// ExampleNewCellImageCell demonstrates how to create a Cell instance using NewCellImageCell.
// Returns an   object initialized with the specified image and set to have the cell’s default menu.
func ExampleNewCellImageCell() {
	_ = appkit.NewCellImageCell(
		appkit.Image{}, // image Image
	)
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
// ExampleNewCellWithCoder demonstrates how to create a Cell instance using NewCellWithCoder.
func ExampleNewCellWithCoder() {
	_ = appkit.NewCellWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
