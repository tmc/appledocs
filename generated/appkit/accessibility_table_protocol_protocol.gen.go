// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAccessibilityTable is the NSAccessibilityTable protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a table view.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityTable
type PAccessibilityTable interface {
	// Required methods
	AccessibilityLabel() foundation.String
	AccessibilityRows() []objc.ID
	// Optional methods
	AccessibilityColumnHeaderUIElements() foundation.Array
	HasAccessibilityColumnHeaderUIElements() bool
	AccessibilityColumns() foundation.Array
	HasAccessibilityColumns() bool
	AccessibilityHeaderGroup() foundation.String
	HasAccessibilityHeaderGroup() bool
	AccessibilityRowHeaderUIElements() foundation.Array
	HasAccessibilityRowHeaderUIElements() bool
	AccessibilitySelectedCells() foundation.Array
	HasAccessibilitySelectedCells() bool
	AccessibilitySelectedColumns() foundation.Array
	HasAccessibilitySelectedColumns() bool
	AccessibilitySelectedRows() []objc.ID
	HasAccessibilitySelectedRows() bool
	AccessibilityVisibleCells() foundation.Array
	HasAccessibilityVisibleCells() bool
	AccessibilityVisibleColumns() foundation.Array
	HasAccessibilityVisibleColumns() bool
	AccessibilityVisibleRows() []objc.ID
	HasAccessibilityVisibleRows() bool
	SetAccessibilitySelectedRows(selectedRows []objc.ID)
	HasSetAccessibilitySelectedRows() bool
}
