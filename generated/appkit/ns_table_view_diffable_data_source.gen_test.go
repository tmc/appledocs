// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTableViewDiffableDataSource

// ExampleTableViewDiffableDataSource_Snapshot demonstrates using Snapshot on a TableViewDiffableDataSource instance.
// Returns a representation of the current state of the data in the table view.
func ExampleTableViewDiffableDataSource_Snapshot() {
	obj := appkit.NewTableViewDiffableDataSource()
	_ = obj.Snapshot()
	// Output:
	}

