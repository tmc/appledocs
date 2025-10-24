// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewButtonCell

// ExampleNewButtonCellImageCell demonstrates how to create a ButtonCell instance using NewButtonCellImageCell.
func ExampleNewButtonCellImageCell() {
	_ = appkit.NewButtonCellImageCell(
		appkit.Image{}, // image Image
	)
	// Output:
}

// ExampleNewButtonCellTextCell demonstrates how to create a ButtonCell instance using NewButtonCellTextCell.
func ExampleNewButtonCellTextCell() {
	_ = appkit.NewButtonCellTextCell(
		"string", // string string
	)
	// Output:
}
