// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)


// ExampleNewCell demonstrates how to create a Cell instance.
func ExampleNewCell() {
	_ = appkit.NewCell()
	// Output:
}

// ExampleNewCellWithCoder demonstrates how to create a Cell instance using NewCellWithCoder.
func ExampleNewCellWithCoder() {
	_ = appkit.NewCellWithCoder(
		nil, // coder unsafe.Pointer
	)
	// Output:
}

// ExampleNewCellImageCell demonstrates how to create a Cell instance using NewCellImageCell.
// Returns an   object initialized with the specified image and set to have the cell’s default menu.
func ExampleNewCellImageCell() {
	_ = appkit.NewCellImageCell(
		nil, // image unsafe.Pointer
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


