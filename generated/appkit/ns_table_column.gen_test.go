// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTableColumn

// ExampleNewTableColumnWithIdentifier demonstrates how to create a TableColumn instance using NewTableColumnWithIdentifier.
// Initializes a newly created table column with a string identifier.
func ExampleNewTableColumnWithIdentifier() {
	_ = appkit.NewTableColumnWithIdentifier(
		appkit.UserInterfaceItemIdentifier{}, // identifier UserInterfaceItemIdentifier
	)
	// Output:
}
// ExampleTableColumn_SizeToFit demonstrates using SizeToFit on a TableColumn instance.
// Resizes the table column to fit the width of its header cell.
func ExampleTableColumn_SizeToFit() {
	obj := appkit.NewTableColumn()
	obj.SizeToFit()
	// Output:
	}

