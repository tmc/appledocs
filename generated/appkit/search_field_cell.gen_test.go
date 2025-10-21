// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewSearchFieldCell

// ExampleNewSearchFieldCellTextCell demonstrates how to create a SearchFieldCell instance using NewSearchFieldCellTextCell.
func ExampleNewSearchFieldCellTextCell() {
	_ = appkit.NewSearchFieldCellTextCell(
		"string", // string string
	)
	// Output:
}
// ExampleNewSearchFieldCellWithCoder demonstrates how to create a SearchFieldCell instance using NewSearchFieldCellWithCoder.
func ExampleNewSearchFieldCellWithCoder() {
	_ = appkit.NewSearchFieldCellWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
