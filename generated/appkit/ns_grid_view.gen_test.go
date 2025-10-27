// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewGridView

// ExampleNewGridViewWithNumberOfColumnsRows demonstrates how to create a GridView instance using NewGridViewWithNumberOfColumnsRows.
// Creates a newly allocated grid view object with the specified number of columns and rows.
func ExampleNewGridViewWithNumberOfColumnsRows() {
	_ = appkit.NewGridViewWithNumberOfColumnsRows(
		10, // columnCount int
		10, // rowCount int
	)
	// Output:
}
