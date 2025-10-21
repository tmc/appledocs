// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit_test

import (
	"github.com/tmc/appledocs/generated/classkit"
)

// Suppress unused import errors
var _ = classkit.NewSQuantityItem

// ExampleNewSQuantityItemWithIdentifierTitle demonstrates how to create a SQuantityItem instance using NewSQuantityItemWithIdentifierTitle.
// Initializes an activity item that records a discrete quantity.
func ExampleNewSQuantityItemWithIdentifierTitle() {
	_ = classkit.NewSQuantityItemWithIdentifierTitle(
		"identifier", // identifier string
		"title", // title string
	)
	// Output:
}
