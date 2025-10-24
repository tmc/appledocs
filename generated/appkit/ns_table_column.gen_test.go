// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTableColumn

// ExampleTableColumn_SizeToFit demonstrates using SizeToFit on a TableColumn instance.
// Resizes the table column to fit the width of its header cell.
func ExampleTableColumn_SizeToFit() {
	obj := appkit.NewTableColumn()
	obj.SizeToFit()
	// Output:
	}

