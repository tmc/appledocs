// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTableView

// ExampleTableView_BeginUpdates demonstrates using BeginUpdates on a TableView instance.
// Begins a group of updates for the table view.
func ExampleTableView_BeginUpdates() {
	obj := appkit.NewTableView()
	obj.BeginUpdates()
	// Output:
	}

// ExampleTableView_EndUpdates demonstrates using EndUpdates on a TableView instance.
// Ends the group of updates for the table view.
func ExampleTableView_EndUpdates() {
	obj := appkit.NewTableView()
	obj.EndUpdates()
	// Output:
	}

// ExampleTableView_NoteNumberOfRowsChanged demonstrates using NoteNumberOfRowsChanged on a TableView instance.
// Informs the table view that the number of records in its data source has changed.
func ExampleTableView_NoteNumberOfRowsChanged() {
	obj := appkit.NewTableView()
	obj.NoteNumberOfRowsChanged()
	// Output:
	}

// ExampleTableView_ReloadData demonstrates using ReloadData on a TableView instance.
// Marks the table view as needing redisplay, so it will reload the data for visible cells and draw the new values.
func ExampleTableView_ReloadData() {
	obj := appkit.NewTableView()
	obj.ReloadData()
	// Output:
	}

// ExampleTableView_SizeLastColumnToFit demonstrates using SizeLastColumnToFit on a TableView instance.
// Resizes the last column so the table view fits exactly within its enclosing clip view.
func ExampleTableView_SizeLastColumnToFit() {
	obj := appkit.NewTableView()
	obj.SizeLastColumnToFit()
	// Output:
	}

// ExampleTableView_SizeToFit demonstrates using SizeToFit on a TableView instance.
// Sizes the  table view based on a uniform column autoresizing style.
func ExampleTableView_SizeToFit() {
	obj := appkit.NewTableView()
	obj.SizeToFit()
	// Output:
	}

// ExampleTableView_Tile demonstrates using Tile on a TableView instance.
// Properly sizes the table view and its header view and marks it as needing display.
func ExampleTableView_Tile() {
	obj := appkit.NewTableView()
	obj.Tile()
	// Output:
	}

